// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.jetbrains.remote.actions

import com.intellij.openapi.actionSystem.AnAction
import com.intellij.openapi.actionSystem.AnActionEvent
import com.intellij.openapi.components.service
import io.khulnasoft.jetbrains.remote.KhulnasoftManager

class DashboardAction : AnAction() {
    private val manager = service<KhulnasoftManager>()

    override fun actionPerformed(event: AnActionEvent) {
        manager.pendingInfo.thenAccept { workspaceInfo ->
            manager.openUrlFromAction(workspaceInfo.khulnasoftHost)
        }
    }
}
