// Copyright (c) 2020 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"context"
	"time"

	"github.com/khulnasoft-com/khulnasoft/khulnasoft-cli/pkg/khulnasoft"
	"github.com/khulnasoft-com/khulnasoft/khulnasoft-cli/pkg/utils"
	"github.com/spf13/cobra"
)

// stopWorkspaceCmd represents the stopWorkspaceCmd command
var stopWorkspaceCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop current workspace",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TOOD: currently, we skip sending analysis because we cannot send it while the workspace stopping.
		utils.TrackCommandUsageEvent.Command = nil

		ctx, cancel := context.WithTimeout(cmd.Context(), 5*time.Second)
		defer cancel()
		wsInfo, err := khulnasoft.GetWSInfo(ctx)
		if err != nil {
			return err
		}
		client, err := khulnasoft.ConnectToServer(ctx, wsInfo, []string{
			"function:stopWorkspace",
			"resource:workspace::" + wsInfo.WorkspaceId + "::get/update",
		})
		if err != nil {
			return err
		}
		defer client.Close()
		return client.StopWorkspace(ctx, wsInfo.WorkspaceId)
	},
}

func init() {
	rootCmd.AddCommand(stopWorkspaceCmd)
}
