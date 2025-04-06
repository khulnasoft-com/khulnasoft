#!/bin/bash
# Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.

set -e

JB_GP_VERSION=${1:-debug}

if [ "${NO_VERIFY_JB_PLUGIN}" == "true" ]; then
    echo "build.sh: skip verify plugin step"
else
    ./gradlew -PsupervisorApiProjectPath=components-supervisor-api-java--lib/ -PkhulnasoftProtocolProjectPath=components-khulnasoft-protocol-java--lib/ -PenvironmentName="$JB_QUALIFIER" -Dgradle.user.home="/workspace/.gradle-$JB_QUALIFIER" -Dplugin.verifier.home.dir="$HOME/.cache/pluginVerifier-$JB_QUALIFIER" -PkhulnasoftVersion="$JB_GP_VERSION" runPluginVerifier
fi
./gradlew -PsupervisorApiProjectPath=components-supervisor-api-java--lib/ -PkhulnasoftProtocolProjectPath=components-khulnasoft-protocol-java--lib/ -PenvironmentName="$JB_QUALIFIER" -Dgradle.user.home="/workspace/.gradle-$JB_QUALIFIER" -Dplugin.verifier.home.dir="$HOME/.cache/pluginVerifier-$JB_QUALIFIER" -PkhulnasoftVersion="$JB_GP_VERSION" buildPlugin
unzip ./build/distributions/khulnasoft-remote.zip -d ./build
