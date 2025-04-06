// Copyright (c) 2024 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package ide

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"
)

type testLogWriter struct {
	t *testing.T
}

var _ io.Writer = &testLogWriter{}

func (t *testLogWriter) Write(p []byte) (n int, err error) {
	t.t.Log(string(p))
	return len(p), nil
}

const localDebug = false

func testWithoutGithubAction(ctx context.Context, gatewayLink, khulnasoftAccessToken, secretEndpoint string, useLatest bool) {
	scriptName := "dev/jetbrains-test:test-stable"
	if useLatest {
		scriptName = "dev/jetbrains-test:test-latest"
	}

	if localDebug {
		fmt.Printf("Exec command below to run UI tests:\n\nexport DISPLAY=:0\nexport GATEWAY_LINK=\"%s\"\nexport KHULNASOFT_TEST_ACCESSTOKEN=\"%s\"\nexport WS_ENDPOINT=%s\nblazedock run %s -Dversion=integration-test -DpublishToJBMarketplace=false", gatewayLink, khulnasoftAccessToken, secretEndpoint, scriptName)
		os.Exit(1)
	}
	cmdEnv := os.Environ()
	cmdEnv = append(cmdEnv, "GATEWAY_LINK="+gatewayLink)
	cmdEnv = append(cmdEnv, "KHULNASOFT_TEST_ACCESSTOKEN="+khulnasoftAccessToken)
	cmdEnv = append(cmdEnv, "WS_ENDPOINT="+secretEndpoint)
	cmd := exec.CommandContext(ctx, "blazedock", "run", scriptName, "-Dversion=integration-test", "-DpublishToJBMarketplace=false")
	cmd.Env = cmdEnv
	// writer := &testLogWriter{t: t}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
}
