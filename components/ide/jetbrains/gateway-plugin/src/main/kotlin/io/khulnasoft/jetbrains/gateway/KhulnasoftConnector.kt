// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.jetbrains.gateway

import com.jetbrains.gateway.api.GatewayConnector
import com.jetbrains.gateway.api.GatewayConnectorDocumentationPage
import com.jetbrains.rd.util.lifetime.Lifetime
import io.khulnasoft.jetbrains.icons.KhulnasoftIcons
import java.awt.Component

class KhulnasoftConnector : GatewayConnector {
    override val icon = KhulnasoftIcons.Logo

    override fun createView(lifetime: Lifetime) = KhulnasoftConnectorView(lifetime)

    override fun getActionText() = "Connect to Khulnasoft"

    override fun getDescription() = "Connect to Khulnasoft workspaces"

    override fun getDocumentationAction() = GatewayConnectorDocumentationPage("https://www.khulnasoft.com/docs/ides-and-editors/jetbrains-gateway")

    override fun getConnectorId() = "khulnasoft.connector"

    override fun getRecentConnections(setContentCallback: (Component) -> Unit) = KhulnasoftRecentConnections()

    override fun getTitle() = "Khulnasoft"

    @Deprecated("Not used", ReplaceWith("null"))
    override fun getTitleAdornment() = null

    override fun initProcedure() {}
}
