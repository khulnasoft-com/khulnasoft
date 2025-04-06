// Copyright (c) 2025 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddKhulnasoftTrailer(t *testing.T) {
	tests := []struct {
		Name        string
		CommitMsg   string
		HostName    string
		Expected    string
		ExpectError bool
	}{
		{
			Name:        "adds trailer to simple message",
			CommitMsg:   "Initial commit",
			HostName:    "khulnasoft.com",
			Expected:    "Initial commit\n\nTool: khulnasoft/khulnasoft.com\n",
			ExpectError: false,
		},
		{
			Name:        "doesn't duplicate existing trailer",
			CommitMsg:   "Initial commit\n\nTool: khulnasoft/khulnasoft.com\n",
			HostName:    "khulnasoft.com",
			Expected:    "Initial commit\n\nTool: khulnasoft/khulnasoft.com\n",
			ExpectError: false,
		},
		{
			Name:        "preserves other trailers",
			CommitMsg:   "Initial commit\n\nSigned-off-by: Kyle <john@example.com>\n",
			HostName:    "khulnasoft.com",
			Expected:    "Initial commit\n\nSigned-off-by: Kyle <john@example.com>\nTool: khulnasoft/khulnasoft.com\n",
			ExpectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tmpfile, err := os.CreateTemp("", "commit-msg-*")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if err := os.WriteFile(tmpfile.Name(), []byte(tt.CommitMsg), 0644); err != nil {
				t.Fatal(err)
			}

			err = addKhulnasoftTrailer(tmpfile.Name(), tt.HostName)
			if (err != nil) != tt.ExpectError {
				t.Errorf("addKhulnasoftTrailer() error = %v, wantErr %v", err, tt.ExpectError)
				return
			}

			got, err := os.ReadFile(tmpfile.Name())
			if err != nil {
				t.Fatal(err)
			}

			equal := cmp.Equal(string(got), tt.Expected)
			if !equal {
				t.Fatalf(`Detected git command info was incorrect, got: %v, expected: %v.`, string(got), tt.Expected)
			}
		})
	}
}
