// Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package server

import (
	"github.com/khulnasoft-com/khulnasoft/installer/pkg/common"
)

const (
	Component                              = common.ServerComponent
	ContainerPort                          = 3000
	ContainerPortName                      = "http"
	authProviderFilePath                   = "/khulnasoft/auth-providers"
	licenseFilePath                        = "/khulnasoft/license"
	stripeSecretMountPath                  = "/stripe-secret"
	linkedInSecretMountPath                = "/linkedin-secret"
	githubAppCertSecret                    = "github-app-cert-secret"
	IAMSessionPort                         = common.ServerIAMSessionPort
	IAMSessionPortName                     = "session"
	InstallationAdminPort                  = common.ServerInstallationAdminPort
	InstallationAdminName                  = "install-admin"
	DebugPortName                          = "debug"
	DebugNodePortName                      = "debugnode"
	ServicePort                            = 3000
	personalAccessTokenSigningKeyMountPath = "/secrets/personal-access-token-signing-key"
	ProbesPort                             = 9400
	ProbesPortName                         = "probes"

	AdminCredentialsSecretName      = "admin-credentials"
	AdminCredentialsSecretMountPath = "/credentials/admin"
	AdminCredentialsSecretKey       = "admin.json"

	GRPCAPIName = "grpc"
	GRPCAPIPort = common.ServerGRPCAPIPort

	PublicAPIName = "public-api"
	PublicAPIPort = common.ServerPublicAPIPort
)
