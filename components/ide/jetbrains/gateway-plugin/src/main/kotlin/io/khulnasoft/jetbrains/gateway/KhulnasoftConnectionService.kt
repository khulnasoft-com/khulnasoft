// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.jetbrains.gateway

import com.intellij.ide.plugins.PluginManagerCore
import com.intellij.openapi.components.Service
import com.intellij.openapi.diagnostic.thisLogger
import com.intellij.openapi.extensions.PluginId
import com.intellij.remoteDev.util.onTerminationOrNow
import com.intellij.util.ExceptionUtil
import com.intellij.util.io.DigestUtil
import com.intellij.util.net.JdkProxyProvider
import com.intellij.util.net.ssl.CertificateManager
import com.jetbrains.rd.util.concurrentMapOf
import com.jetbrains.rd.util.lifetime.Lifetime
import io.khulnasoft.khulnasoftprotocol.api.KhulnasoftServerLauncher
import io.khulnasoft.jetbrains.auth.KhulnasoftAuthService
import kotlinx.coroutines.*
import kotlinx.coroutines.future.await
import org.eclipse.jetty.websocket.api.UpgradeException
import java.net.URL
import java.nio.charset.StandardCharsets

@Service
class KhulnasoftConnectionService {

    private val clients = concurrentMapOf<String, GatewayKhulnasoftClient>()

    fun obtainClient(khulnasoftHost: String): GatewayKhulnasoftClient {
        return clients.getOrPut(khulnasoftHost) {
            val lifetime = Lifetime.Eternal.createNested()
            val client = GatewayKhulnasoftClient(lifetime, khulnasoftHost)
            val launcher = KhulnasoftServerLauncher.create(client)
            val job = GlobalScope.launch {
                var accessToken = KhulnasoftAuthService.getAccessToken(khulnasoftHost)
                val authorize = suspend {
                    ensureActive()
                    accessToken = KhulnasoftAuthService.authorize(khulnasoftHost)
                }
                if (accessToken == null) {
                    authorize()
                }

                val plugin = PluginManagerCore.getPlugin(PluginId.getId("io.khulnasoft.jetbrains.gateway"))!!
                val connect = suspend {
                    ensureActive()

                    val originalClassLoader = Thread.currentThread().contextClassLoader
                    val connection = try {
                        val origin = "https://$khulnasoftHost/"
                        val proxies = JdkProxyProvider.getInstance().proxySelector.select(URL(origin).toURI())
                        val sslContext = CertificateManager.getInstance().sslContext

                        // see https://intellij-support.jetbrains.com/hc/en-us/community/posts/360003146180/comments/360000376240
                        Thread.currentThread().contextClassLoader = KhulnasoftConnectionProvider::class.java.classLoader

                        launcher.listen(
                                "wss://$khulnasoftHost/api/v1",
                                origin,
                                plugin.pluginId.idString,
                                plugin.version,
                                accessToken,
                                proxies,
                                sslContext
                        )
                    } catch (t: Throwable) {
                        val badUpgrade = ExceptionUtil.findCause(t, UpgradeException::class.java)
                        if (badUpgrade?.responseStatusCode == 401 || badUpgrade?.responseStatusCode == 403) {
                            throw InvalidTokenException("failed web socket handshake (${badUpgrade.responseStatusCode})")
                        }
                        throw t
                    } finally {
                        Thread.currentThread().contextClassLoader = originalClassLoader
                    }
                    val tokenHash = DigestUtil.sha256Hex(accessToken!!.toByteArray(StandardCharsets.UTF_8))
                    val tokenScopes = client.server.getKhulnasoftTokenScopes(tokenHash).await()
                    for (scope in KhulnasoftAuthService.scopes) {
                        if (!tokenScopes.contains(scope)) {
                            connection.cancel(false)
                            throw InvalidTokenException("$scope scope is not granted")
                        }
                    }
                    connection
                }

                val minReconnectionDelay = 2 * 1000L
                val maxReconnectionDelay = 30 * 1000L
                val reconnectionDelayGrowFactor = 1.5
                var reconnectionDelay = minReconnectionDelay
                while (isActive) {
                    try {
                        val connection = try {
                            connect()
                        } catch (t: Throwable) {
                            val e = ExceptionUtil.findCause(t, InvalidTokenException::class.java) ?: throw t
                            thisLogger().warn("$khulnasoftHost: invalid token, authorizing again and reconnecting:", e)
                            authorize()
                            connect()
                        }
                        reconnectionDelay = minReconnectionDelay
                        thisLogger().info("$khulnasoftHost: connected")
                        val reason = connection.await()
                        if (isActive) {
                            thisLogger().warn("$khulnasoftHost: connection closed, reconnecting after $reconnectionDelay milliseconds: $reason")
                        } else {
                            thisLogger().info("$khulnasoftHost: connection permanently closed: $reason")
                        }
                    } catch (t: Throwable) {
                        if (isActive) {
                            thisLogger().warn(
                                    "$khulnasoftHost: failed to connect, trying again after $reconnectionDelay milliseconds:",
                                    t
                            )
                        } else {
                            thisLogger().error("$khulnasoftHost: connection permanently closed:", t)
                        }
                    }
                    delay(reconnectionDelay)
                    reconnectionDelay = (reconnectionDelay * reconnectionDelayGrowFactor).toLong()
                    if (reconnectionDelay > maxReconnectionDelay) {
                        reconnectionDelay = maxReconnectionDelay
                    }
                }
            }
            lifetime.onTerminationOrNow {
                clients.remove(khulnasoftHost)
                job.cancel()
            }
            return@getOrPut client
        }
    }

    private class InvalidTokenException(message: String) : Exception(message)
}
