/**
 * Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import * as protocol from "@khulnasoft/khulnasoft-protocol/lib/workspace-cluster";

// This interface lives in protocol due to dependency issues but is re-exported here for consistency.
export const WorkspaceClusterDB = protocol.WorkspaceClusterDB;
export interface WorkspaceClusterDB extends protocol.WorkspaceClusterDB {}
