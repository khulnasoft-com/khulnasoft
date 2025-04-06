// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.jetbrains.remote

import com.jetbrains.rdserver.unattendedHost.customization.GatewayClientCustomizationProvider
import com.jetbrains.rdserver.unattendedHost.customization.controlCenter.DefaultGatewayControlCenterProvider
import com.jetbrains.rdserver.unattendedHost.customization.controlCenter.GatewayControlCenterProvider
import com.jetbrains.rdserver.unattendedHost.customization.controlCenter.GatewayHostnameDisplayKind
import io.khulnasoft.jetbrains.remote.icons.KhulnasoftIcons
import javax.swing.Icon

class KhulnasoftGatewayClientCustomizationProvider : GatewayClientCustomizationProvider {
    override val icon: Icon = KhulnasoftIcons.Logo
    override val title: String = System.getenv("JETBRAINS_KHULNASOFT_WORKSPACE_HOST") ?: DefaultGatewayControlCenterProvider().getHostnameShort()

    override val controlCenter: GatewayControlCenterProvider = object : GatewayControlCenterProvider {
        override fun getHostnameDisplayKind() = GatewayHostnameDisplayKind.ShowHostnameOnNavbar
        override fun getHostnameShort() = System.getenv("KHULNASOFT_WORKSPACE_NAME") ?: title
        override fun getHostnameLong() = title
    }
}
