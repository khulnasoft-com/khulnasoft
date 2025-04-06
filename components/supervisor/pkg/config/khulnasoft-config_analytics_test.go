// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package config

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/khulnasoft-com/khulnasoft/common-go/log"
	khulnasoft "github.com/khulnasoft-com/khulnasoft/khulnasoft-protocol"
)

func TestAnalyzeKhulnasoftConfig(t *testing.T) {
	tests := []struct {
		Desc    string
		Prev    *khulnasoft.KhulnasoftConfig
		Current *khulnasoft.KhulnasoftConfig
		Fields  []string
	}{
		{
			Desc: "change",
			Prev: &khulnasoft.KhulnasoftConfig{
				CheckoutLocation: "foo",
			},
			Current: &khulnasoft.KhulnasoftConfig{
				CheckoutLocation: "bar",
			},
			Fields: []string{"CheckoutLocation"},
		},
		{
			Desc: "add",
			Prev: &khulnasoft.KhulnasoftConfig{},
			Current: &khulnasoft.KhulnasoftConfig{
				CheckoutLocation: "bar",
			},
			Fields: []string{"CheckoutLocation"},
		},
		{
			Desc: "remove",
			Prev: &khulnasoft.KhulnasoftConfig{
				CheckoutLocation: "bar",
			},
			Current: &khulnasoft.KhulnasoftConfig{},
			Fields:  []string{"CheckoutLocation"},
		},
		{
			Desc: "none",
			Prev: &khulnasoft.KhulnasoftConfig{
				CheckoutLocation: "bar",
			},
			Current: &khulnasoft.KhulnasoftConfig{
				CheckoutLocation: "bar",
			},
			Fields: nil,
		},
		{
			Desc:    "fie created",
			Current: &khulnasoft.KhulnasoftConfig{},
			Fields:  nil,
		},
	}
	for _, test := range tests {
		t.Run(test.Desc, func(t *testing.T) {
			var fields []string
			analyzer := NewConfigAnalyzer(log.Log, 100*time.Millisecond, func(field string) {
				fields = append(fields, field)
			}, test.Prev)
			<-analyzer.Analyse(test.Current)
			if diff := cmp.Diff(test.Fields, fields); diff != "" {
				t.Errorf("unexpected output (-want +got):\n%s", diff)
			}
		})
	}
}
