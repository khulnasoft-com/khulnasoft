/**
 * Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { Button } from "@podkit/buttons/Button";
import InfoBox from "../components/InfoBox";
import { khulnasoftHostUrl } from "../service/service";
import { Heading1, Subheading } from "@podkit/typography/Headings";

export default function InstallGitHubApp() {
    const goToApp = () => (window.location.href = khulnasoftHostUrl.toString());

    return (
        <>
            <div className="app-container flex flex-col space-y-2">
                <div className="px-6 py-3 flex justify-between space-x-2">
                    <div className="flex flex-col items-center m-auto max-w-lg mt-40">
                        <Heading1 className="text-center pb-3">GitHub App 🌅</Heading1>
                        <Subheading className="text-center pb-6">
                            You likely tried to install the GitHub App for Khulnasoft.
                        </Subheading>
                        <InfoBox>Khulnasoft no longer requires to install the GitHub App on repositories.</InfoBox>
                        <div className="mt-6">
                            <Button onClick={goToApp}>Go to Dashboard</Button>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}
