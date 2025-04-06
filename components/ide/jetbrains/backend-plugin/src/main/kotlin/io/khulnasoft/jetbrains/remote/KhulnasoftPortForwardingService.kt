// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.jetbrains.remote

import com.intellij.openapi.Disposable
import java.util.Optional
import com.jetbrains.rd.util.URI

interface KhulnasoftPortForwardingService : Disposable {
    /** Returns the localhost URI if the given port is forwarded on client. */
    fun getLocalHostUriFromHostPort(hostPort: Int): Optional<URI>
}
