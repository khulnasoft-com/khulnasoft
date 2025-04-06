// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	_ "embed"
	"fmt"

	"github.com/khulnasoft-com/khulnasoft/khulnasoft-cli/pkg/khulnasoft"

	"github.com/spf13/cobra"
)

// urlCmd represents the url command
var versionCmd = &cobra.Command{
	Use:    "version",
	Hidden: false,
	Short:  "Prints the version of the CLI",
	Args:   cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(khulnasoft.Version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
