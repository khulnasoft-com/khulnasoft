// Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

rootProject.name = "jetbrains-gateway-khulnasoft-plugin"

include(":khulnasoft-protocol")
val khulnasoftProtocolProjectPath: String by settings
project(":khulnasoft-protocol").projectDir = File(khulnasoftProtocolProjectPath)
