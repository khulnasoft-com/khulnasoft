// Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package ide_proxy

import "github.com/khulnasoft-com/khulnasoft/installer/pkg/common"

const (
	Component     = common.IDEProxyComponent
	ContainerPort = common.IDEProxyPort
	PortName      = "http"
	ServicePort   = common.IDEProxyPort
	ReadinessPort = 8080
)
