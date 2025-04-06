// Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package wsmanagermk2

import "github.com/khulnasoft-com/khulnasoft/installer/pkg/common"

const (
	Component                  = common.WSManagerMk2Component
	RPCPort                    = 8080
	RPCPortName                = "rpc"
	HealthPort                 = 9090
	TLSSecretNameSecret        = "ws-manager-mk2-tls"
	TLSSecretNameClient        = "ws-manager-mk2-client-tls"
	VolumeConfig               = "config"
	VolumeTLSCerts             = "tls-certs"
	VolumeWorkspaceTemplate    = "workspace-template"
	WorkspaceTemplatePath      = "/workspace-templates"
	WorkspaceTemplateConfigMap = "workspace-templates"
)
