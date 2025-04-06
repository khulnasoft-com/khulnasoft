/**
 * Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { BillingTier } from "../protocol";

export const Client = Symbol("Client");

// Attributes define attributes which can be used to segment audiences.
// Set the attributes which you want to use to group audiences into.
export interface Attributes {
    // user.id is mapped to ConfigCat's "identifier" + "custom.user_id"
    user?: { id: string; email?: string };

    // The BillingTier of this particular user
    billingTier?: BillingTier;

    // Currently selected Khulnasoft Project ID (mapped to "custom.project_id")
    projectId?: string;

    // Currently selected Khulnasoft Team ID (mapped to "custom.team_id")
    teamId?: string;

    // Currently selected Khulnasoft Team Name (mapped to "custom.team_name")
    teamName?: string;

    // Host name of the Khulnasoft installation (mapped to "custom.khulnasoft_host")
    khulnasoftHost?: string;
}

export interface Client {
    getValueAsync<T>(experimentName: string, defaultValue: T, attributes: Attributes): Promise<T>;

    // dispose will dispose of the client, no longer retrieving flags
    dispose(): void;
}
