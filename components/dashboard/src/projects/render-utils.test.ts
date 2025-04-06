/**
 * Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { toRemoteURL } from "./render-utils";

test("parse clone URL", () => {
    expect(toRemoteURL("https://gitlab.com/laushinka/my-node-project")).toEqual("gitlab.com/laushinka/my-node-project");
    expect(toRemoteURL("https://gitlab.gitlab.khulnasoft.com/test-group/test-project")).toEqual(
        "gitlab.gitlab.khulnasoft.com/test-group/test-project",
    );
});
