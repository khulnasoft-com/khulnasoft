// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.khulnasoftprotocol.api;

import javax.websocket.CloseReason;
import javax.websocket.Session;
import java.io.IOException;
import java.util.concurrent.CompletableFuture;
import java.util.logging.Level;
import java.util.logging.Logger;

public class KhulnasoftServerConnectionImpl extends CompletableFuture<CloseReason> implements KhulnasoftServerConnection {

    public static final Logger LOG = Logger.getLogger(KhulnasoftServerConnectionImpl.class.getName());

    private final String khulnasoftHost;

    private Session session;

    public KhulnasoftServerConnectionImpl(String khulnasoftHost) {
        this.khulnasoftHost = khulnasoftHost;
    }

    public void setSession(Session session) {
        this.session = session;
    }

    @Override
    public boolean cancel(boolean mayInterruptIfRunning) {
        Session session = this.session;
        this.session = null;
        if (session != null) {
            try {
                session.close();
            } catch (IOException e) {
                LOG.log(Level.WARNING, khulnasoftHost + ": failed to close connection:", e);
            }
        }
        return super.cancel(mayInterruptIfRunning);
    }
}
