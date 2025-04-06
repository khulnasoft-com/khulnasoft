// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package analytics

import (
	"github.com/khulnasoft-com/khulnasoft/common-go/analytics"
)

var writer analytics.Writer = analytics.NewFromEnvironment()

func Identify(message analytics.IdentifyMessage) {
	writer.Identify(message)
}
func Track(message analytics.TrackMessage) {
	writer.Track(message)
}
func Close() error {
	return writer.Close()
}
