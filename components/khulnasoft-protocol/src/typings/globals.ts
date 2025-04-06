/**
 * Copyright (c) 2020 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// eslint-disable-next-line @typescript-eslint/no-unused-vars
interface Window {
    // khulnasoft is used by supervisor/frontend and vscode browser https://github.com/khulnasoft-com/openvscode-server/blob/dadd40deda5959c0efdaa0c7d02ab64b6f8b5ad0/src/vs/khulnasoft/browser/workbench/workbench.ts#L7
    khulnasoft: {
        ideService?: import("../ide-frontend-service").IDEFrontendService;
        loggedUserID?: string;

        openDesktopIDE(url: string): void;

        encrypt(content: string): string;
        decrypt(content: string): string;
        isEncryptedData(content: string): boolean;
    };
}
