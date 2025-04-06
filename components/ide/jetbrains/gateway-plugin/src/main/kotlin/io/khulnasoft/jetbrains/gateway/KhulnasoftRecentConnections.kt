// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.jetbrains.gateway

import com.intellij.ui.dsl.builder.AlignX
import com.intellij.ui.dsl.builder.AlignY
import com.intellij.ui.dsl.builder.panel
import com.jetbrains.gateway.api.GatewayRecentConnections
import com.jetbrains.rd.util.lifetime.Lifetime
import io.khulnasoft.jetbrains.icons.KhulnasoftIcons
import javax.swing.JComponent

class KhulnasoftRecentConnections : GatewayRecentConnections {

    override val recentsIcon = KhulnasoftIcons.Logo

    private lateinit var view: KhulnasoftWorkspacesView
    override fun createRecentsView(lifetime: Lifetime): JComponent {
        this.view = KhulnasoftWorkspacesView(lifetime)
        return panel {
            row {
                resizableRow()
                cell(view.component)
                    .resizableColumn()
                    .align(AlignX.FILL)
                    .align(AlignY.FILL)
                cell()
            }
        }
    }

    override fun getRecentsTitle(): String {
        return "Khulnasoft"
    }

    override fun updateRecentView() {
        if (this::view.isInitialized) {
            this.view.refresh()
        }
    }
}
