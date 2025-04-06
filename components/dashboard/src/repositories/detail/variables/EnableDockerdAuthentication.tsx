/**
 * Copyright (c) 2025 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { SwitchInputField } from "@podkit/switch/Switch";
import { Heading3, Subheading } from "@podkit/typography/Headings";
import { FC, useCallback } from "react";
import { InputField } from "../../../components/forms/InputField";
import { useToast } from "../../../components/toasts/Toasts";
import { useId } from "../../../hooks/useId";
import { ConfigurationSettingsField } from "../ConfigurationSettingsField";
import { Configuration } from "@khulnasoft/public-api/lib/khulnasoft/v1/configuration_pb";
import { SquareArrowOutUpRight } from "lucide-react";
import { useConfiguration, useConfigurationMutation } from "../../../data/configurations/configuration-queries";
import Alert from "../../../components/Alert";

type Props = {
    configuration: Configuration;
};
export const EnableDockerdAuthentication: FC<Props> = ({ configuration }) => {
    const { data } = useConfiguration(configuration.id);
    const configurationMutation = useConfigurationMutation();
    const { toast } = useToast();

    const updateEnableDockerdAuthentication = useCallback(
        async (enable: boolean) => {
            await configurationMutation.mutateAsync(
                {
                    configurationId: configuration.id,
                    workspaceSettings: {
                        enableDockerdAuthentication: enable,
                    },
                },
                {
                    onError: (error) => {
                        toast(`Failed to update dockerd authentication: ${error.message}`);
                    },
                },
            );
        },
        [configurationMutation, configuration.id, toast],
    );

    const inputId = useId({ prefix: "enable-dockerd-authentication" });
    const isEnabled = data?.workspaceSettings?.enableDockerdAuthentication;

    return (
        <ConfigurationSettingsField>
            <Heading3 className="flex flex-row items-center gap-2">Docker registry authentication</Heading3>
            <Subheading className="max-w-lg flex flex-col gap-2">
                <span className="flex-1 text-left">
                    Enable authentication with Docker registries inside of workspaces based on the{" "}
                    <code>KHULNASOFT_IMAGE_AUTH</code> environment variable.
                </span>

                <Alert type={"warning"} closable={false} showIcon={true} className="flex rounded p-2 mb-2 w-full">
                    By enabling this, credentials specified in <code>KHULNASOFT_IMAGE_AUTH</code> will be visible inside all
                    workspaces on this project.
                </Alert>
                <a
                    className="gp-link flex flex-row items-center gap-1"
                    href="https://www.khulnasoft.com/docs/configure/repositories/environment-variables#docker-registry-authentication"
                    target="_blank"
                    rel="noreferrer"
                >
                    Learn about using private Docker images with Khulnasoft
                    <SquareArrowOutUpRight size={12} />
                </a>
            </Subheading>
            <InputField id={inputId}>
                <SwitchInputField
                    id={inputId}
                    checked={isEnabled}
                    disabled={configurationMutation.isLoading}
                    onCheckedChange={updateEnableDockerdAuthentication}
                    label={isEnabled ? "Auto-login enabled" : "Auto-login disabled"}
                />
            </InputField>
        </ConfigurationSettingsField>
    );
};
