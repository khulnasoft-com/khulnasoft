// Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.khulnasoftprotocol.testclient;

import io.khulnasoft.khulnasoftprotocol.api.KhulnasoftClient;
import io.khulnasoft.khulnasoftprotocol.api.KhulnasoftServer;
import io.khulnasoft.khulnasoftprotocol.api.KhulnasoftServerLauncher;
import io.khulnasoft.khulnasoftprotocol.api.entities.SendHeartBeatOptions;
import io.khulnasoft.khulnasoftprotocol.api.entities.User;

import java.util.Collections;

public class TestClient {
    public static void main(String[] args) throws Exception {
        String uri = "wss://khulnasoft.com/api/v1";
        String token = "CHANGE-ME";
        String origin = "https://CHANGE-ME.khulnasoft.com/";

        KhulnasoftClient client = new KhulnasoftClient();
        KhulnasoftServerLauncher.create(client).listen(uri, origin, token, "Test", "Test", Collections.emptyList(), null);
        KhulnasoftServer khulnasoftServer = client.getServer();
        User user = khulnasoftServer.getLoggedInUser().join();
        System.out.println("logged in user:" + user);

        Void result = khulnasoftServer
                .sendHeartBeat(new SendHeartBeatOptions("CHANGE-ME", false)).join();
        System.out.println("send heart beat:" + result);
    }
}
