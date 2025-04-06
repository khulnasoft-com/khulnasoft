// Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package ide_proxy

import "github.com/khulnasoft-com/khulnasoft/installer/pkg/common"

var Objects = common.CompositeRenderFunc(
	deployment,
	rolebinding,
	service,
	common.DefaultServiceAccount(Component),
)
