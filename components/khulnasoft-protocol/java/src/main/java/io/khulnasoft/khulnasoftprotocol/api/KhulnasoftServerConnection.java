// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.khulnasoftprotocol.api;

import javax.websocket.CloseReason;
import java.util.concurrent.CompletionStage;
import java.util.concurrent.Future;

public interface KhulnasoftServerConnection extends Future<CloseReason>, CompletionStage<CloseReason> {
}
