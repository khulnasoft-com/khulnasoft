// Copyright (c) 2025 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.jetbrains.remote.stable

import com.intellij.openapi.client.ClientProjectSession
import com.intellij.openapi.client.ClientSessionsManager
import com.intellij.openapi.project.Project
import io.khulnasoft.jetbrains.remote.AbstractKhulnasoftClientProjectSessionTracker

@Suppress("UnstableApiUsage")
class KhulnasoftClientProjectSessionTracker(val project: Project) : AbstractKhulnasoftClientProjectSessionTracker(project) {
    override val session: ClientProjectSession? = ClientSessionsManager.getProjectSession(project)
}
