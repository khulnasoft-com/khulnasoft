/**
 * Copyright (c) 2020 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import * as chai from "chai";
import { suite, test } from "@testdeck/mocha";
import { KhulnasoftHostUrl } from "./khulnasoft-host-url";
const expect = chai.expect;

@suite
export class KhulnasoftHostUrlTest {
    @test public parseWorkspaceId_hosts_withEnvVarsInjected() {
        const actual = new KhulnasoftHostUrl(
            "https://gray-grasshopper-nfbitfia.ws-eu02.khulnasoft-staging.com/#passedin=test%20value/https://github.com/khulnasoft-com/khulnasoft-test-repo",
        ).workspaceId;
        expect(actual).to.equal("gray-grasshopper-nfbitfia");
    }

    @test public async testWithoutWorkspacePrefix() {
        expect(
            new KhulnasoftHostUrl("https://3000-moccasin-ferret-155799b3.ws-eu02.khulnasoft-staging.com/")
                .withoutWorkspacePrefix()
                .toString(),
        ).to.equal("https://khulnasoft-staging.com/");
    }

    @test public async testWithoutWorkspacePrefix2() {
        expect(new KhulnasoftHostUrl("https://khulnasoft-staging.com/").withoutWorkspacePrefix().toString()).to.equal(
            "https://khulnasoft-staging.com/",
        );
    }

    @test public async testWithoutWorkspacePrefix3() {
        expect(
            new KhulnasoftHostUrl("https://gray-rook-5523v5d8.ws-dev.my-branch-1234.staging.khulnasoft-dev.com/")
                .withoutWorkspacePrefix()
                .toString(),
        ).to.equal("https://my-branch-1234.staging.khulnasoft-dev.com/");
    }

    @test public async testWithoutWorkspacePrefix4() {
        expect(
            new KhulnasoftHostUrl("https://my-branch-1234.staging.khulnasoft-dev.com/").withoutWorkspacePrefix().toString(),
        ).to.equal("https://my-branch-1234.staging.khulnasoft-dev.com/");
    }

    @test public async testWithoutWorkspacePrefix5() {
        expect(
            new KhulnasoftHostUrl("https://abc-nice-brunch-4224.staging.khulnasoft-dev.com/")
                .withoutWorkspacePrefix()
                .toString(),
        ).to.equal("https://abc-nice-brunch-4224.staging.khulnasoft-dev.com/");
    }

    @test public async testWithoutWorkspacePrefix6() {
        expect(
            new KhulnasoftHostUrl("https://gray-rook-5523v5d8.ws-dev.abc-nice-brunch-4224.staging.khulnasoft-dev.com/")
                .withoutWorkspacePrefix()
                .toString(),
        ).to.equal("https://abc-nice-brunch-4224.staging.khulnasoft-dev.com/");
    }
}
module.exports = new KhulnasoftHostUrlTest();
