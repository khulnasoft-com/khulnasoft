/**
 * Copyright (c) 2023 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import * as crypto from "crypto";
import { DBKhulnasoftToken, UserDB } from "@khulnasoft/khulnasoft-db/lib";
import { KhulnasoftToken, KhulnasoftTokenType } from "@khulnasoft/khulnasoft-protocol";
import { log } from "@khulnasoft/khulnasoft-protocol/lib/util/logging";
import { inject, injectable } from "inversify";
import { Authorizer } from "../authorization/authorizer";

@injectable()
export class KhulnasoftTokenService {
    constructor(
        @inject(UserDB) private readonly userDB: UserDB,
        @inject(Authorizer) private readonly auth: Authorizer,
    ) {}

    async getKhulnasoftTokens(requestorId: string, userId: string): Promise<KhulnasoftToken[]> {
        await this.auth.checkPermissionOnUser(requestorId, "read_tokens", userId);
        const khulnasoftTokens = await this.userDB.findAllKhulnasoftTokensOfUser(userId);
        return khulnasoftTokens;
    }

    async generateNewKhulnasoftToken(
        requestorId: string,
        userId: string,
        options: { name?: string; type: KhulnasoftTokenType; scopes?: string[] },
        oldPermissionCheck?: (dbToken: DBKhulnasoftToken) => Promise<void>, // @deprecated
    ): Promise<string> {
        await this.auth.checkPermissionOnUser(requestorId, "write_tokens", userId);
        const token = crypto.randomBytes(30).toString("hex");
        const tokenHash = crypto.createHash("sha256").update(token, "utf8").digest("hex");
        const dbToken: DBKhulnasoftToken = {
            tokenHash,
            name: options.name,
            type: options.type,
            userId,
            scopes: options.scopes || [],
            created: new Date().toISOString(),
        };
        if (oldPermissionCheck) {
            await oldPermissionCheck(dbToken);
        }
        await this.userDB.storeKhulnasoftToken(dbToken);
        return token;
    }

    async findKhulnasoftToken(requestorId: string, userId: string, tokenHash: string): Promise<KhulnasoftToken | undefined> {
        await this.auth.checkPermissionOnUser(requestorId, "read_tokens", userId);
        let token: KhulnasoftToken | undefined;
        try {
            token = await this.userDB.findKhulnasoftTokensOfUser(userId, tokenHash);
        } catch (error) {
            log.error({ userId }, "failed to resolve khulnasoft token: ", error);
        }
        return token;
    }

    async deleteKhulnasoftToken(
        requestorId: string,
        userId: string,
        tokenHash: string,
        oldPermissionCheck?: (token: KhulnasoftToken) => Promise<void>, // @deprecated
    ): Promise<void> {
        await this.auth.checkPermissionOnUser(requestorId, "write_tokens", userId);
        const existingTokens = await this.getKhulnasoftTokens(requestorId, userId);
        const tkn = existingTokens.find((token) => token.tokenHash === tokenHash);
        if (!tkn) {
            throw new Error(`User ${requestorId} tries to delete a token ${tokenHash} that does not exist.`);
        }
        if (oldPermissionCheck) {
            await oldPermissionCheck(tkn);
        }
        await this.userDB.deleteKhulnasoftToken(tokenHash);
    }
}
