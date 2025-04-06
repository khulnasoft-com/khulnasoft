// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package examples

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/khulnasoft-com/khulnasoft/components/public-api/go/client"
	v1 "github.com/khulnasoft-com/khulnasoft/components/public-api/go/experimental/v1"
	"os"
)

func ExampleClient() {
	token := "khulnasoft_pat_example.personal-access-token"
	khulnasoft, err := client.New(client.WithCredentials(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to construct khulnasoft client %v", err)
		return
	}

	// use the khulnasoft client to access resources
	khulnasoft.Teams.ListTeams(context.Background(), connect.NewRequest(&v1.ListTeamsRequest{}))
}
