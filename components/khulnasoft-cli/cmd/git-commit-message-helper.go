// Copyright (c) 2025 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/khulnasoft-com/khulnasoft/khulnasoft-cli/pkg/khulnasoft"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var gitCommitMessageHelperOpts struct {
	CommitMessageFile string
}

func addKhulnasoftTrailer(commitMsgFile string, hostName string) error {
	trailerCmd := exec.Command("git", "interpret-trailers",
		"--if-exists", "addIfDifferent",
		"--trailer", fmt.Sprintf("Tool: khulnasoft/%s", hostName),
		commitMsgFile)

	output, err := trailerCmd.Output()
	if err != nil {
		return fmt.Errorf("error adding trailer: %w", err)
	}

	err = os.WriteFile(commitMsgFile, output, 0644)
	if err != nil {
		return fmt.Errorf("error writing commit message file: %w", err)
	}

	return nil
}

var gitCommitMessageHelper = &cobra.Command{
	Use:    "git-commit-message-helper",
	Short:  "Khulnasoft's Git commit message helper",
	Long:   "Automatically adds Tool information to Git commit messages",
	Args:   cobra.ExactArgs(0),
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), 5*time.Second)
		defer cancel()

		wsInfo, err := khulnasoft.GetWSInfo(ctx)
		if err != nil {
			log.WithError(err).Fatal("error getting workspace info")
			return nil // don't block commit
		}

		if err := addKhulnasoftTrailer(gitCommitMessageHelperOpts.CommitMessageFile, wsInfo.KhulnasoftApi.Host); err != nil {
			log.WithError(err).Fatal("failed to add khulnasoft trailer")
			return nil // don't block commit
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(gitCommitMessageHelper)
	gitCommitMessageHelper.Flags().StringVarP(&gitCommitMessageHelperOpts.CommitMessageFile, "file", "f", "", "Path to the commit message file")
	_ = gitCommitMessageHelper.MarkFlagRequired("file")
}
