// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.jetbrains.remote.services

import com.intellij.openapi.Disposable
import com.intellij.openapi.components.Service
import com.intellij.openapi.components.service
import com.intellij.openapi.diagnostic.thisLogger
import io.khulnasoft.khulnasoftprotocol.api.entities.SendHeartBeatOptions
import io.khulnasoft.jetbrains.remote.KhulnasoftManager
import io.khulnasoft.jetbrains.remote.services.ControllerStatusService.ControllerStatus
import kotlinx.coroutines.*
import kotlinx.coroutines.future.await
import kotlin.random.Random.Default.nextInt

@Service
class HeartbeatService : Disposable {

    private val manager = service<KhulnasoftManager>()

    private val job = GlobalScope.launch {
        val info = manager.pendingInfo.await()
        val intervalInSeconds = 30
        var current = ControllerStatus(
                connected = false,
                secondsSinceLastActivity = 0
        )
        while (isActive) {
            try {
                val previous = current;
                current = ControllerStatusService.fetch()

                val maxIntervalInSeconds = intervalInSeconds + nextInt(5, 15)
                val wasClosed: Boolean? = when {
                    current.connected != previous.connected -> !current.connected
                    current.connected && current.secondsSinceLastActivity <= maxIntervalInSeconds -> false
                    else -> null
                }

                if (wasClosed != null) {
                    manager.client.server.sendHeartBeat(SendHeartBeatOptions(info.instanceId, wasClosed)).await()
                    if (wasClosed) {
                        manager.trackEvent("ide_close_signal", mapOf(
                                "clientKind" to "jetbrains"
                        ))
                    }
                }
            } catch (t: Throwable) {
                thisLogger().error("khulnasoft: failed to check activity:", t)
            }
            delay(intervalInSeconds * 1000L)
        }
    }

    override fun dispose() = job.cancel()
}
