// Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

rootProject.name = "khulnasoft-remote"

include(":supervisor-api")
val supervisorApiProjectPath: String by settings
project(":supervisor-api").projectDir = File(supervisorApiProjectPath)

include(":khulnasoft-protocol")
val khulnasoftProtocolProjectPath: String by settings
project(":khulnasoft-protocol").projectDir = File(khulnasoftProtocolProjectPath)
