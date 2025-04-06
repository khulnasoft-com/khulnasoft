// Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.khulnasoftprotocol.api;

import io.khulnasoft.khulnasoftprotocol.api.entities.WorkspaceInstance;
import org.eclipse.lsp4j.jsonrpc.services.JsonNotification;

public class KhulnasoftClient {

    private KhulnasoftServer server;

    public void connect(KhulnasoftServer server) {
        this.server = server;
    }

    public KhulnasoftServer getServer() {
        if (this.server == null) {
            throw new IllegalStateException("not connected");
        }
        return this.server;
    }

    public void notifyConnect() {
    }

    @JsonNotification
    public void onInstanceUpdate(WorkspaceInstance instance) {

    }
}
