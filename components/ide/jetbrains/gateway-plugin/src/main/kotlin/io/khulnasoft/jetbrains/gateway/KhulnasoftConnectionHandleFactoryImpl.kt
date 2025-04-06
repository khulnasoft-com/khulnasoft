// Copyright (c) 2024 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.jetbrains.gateway

import com.intellij.openapi.components.Service
import com.jetbrains.gateway.api.GatewayConnectionHandle
import com.jetbrains.gateway.ssh.ClientOverSshTunnelConnector
import com.jetbrains.gateway.ssh.HostTunnelConnector
import com.jetbrains.gateway.thinClientLink.ThinClientHandle
import com.jetbrains.rd.util.lifetime.Lifetime
import io.khulnasoft.jetbrains.gateway.KhulnasoftConnectionProvider.ConnectParams
import io.khulnasoft.jetbrains.gateway.common.KhulnasoftConnectionHandle
import io.khulnasoft.jetbrains.gateway.common.KhulnasoftConnectionHandleFactory
import java.net.URI
import javax.swing.JComponent

@Suppress("UnstableApiUsage")
class KhulnasoftConnectionHandleFactoryImpl: KhulnasoftConnectionHandleFactory {
    override fun createKhulnasoftConnectionHandle(
        lifetime: Lifetime,
        component: JComponent,
        params: ConnectParams
    ): GatewayConnectionHandle {
        return KhulnasoftConnectionHandle(lifetime, component, params)
    }

    override suspend fun connect(lifetime: Lifetime, connector: HostTunnelConnector, tcpJoinLink: URI): ThinClientHandle {
        return ClientOverSshTunnelConnector(
            lifetime,
            connector
        ).connect(tcpJoinLink, null)
    }
}
