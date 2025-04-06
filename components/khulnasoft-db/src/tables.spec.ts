/**
 * Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import * as chai from "chai";
const expect = chai.expect;
import { suite, test, timeout } from "@testdeck/mocha";

import { KhulnasoftTableDescriptionProvider } from "./tables";

@suite.only
class TablesSpec {
    async before() {}

    async after() {}

    @test(timeout(10000))
    public async createAndFindATeam() {
        const thing = new KhulnasoftTableDescriptionProvider();
        expect(() => thing.getSortedTables()).to.not.throw();
    }
}

module.exports = new TablesSpec();
