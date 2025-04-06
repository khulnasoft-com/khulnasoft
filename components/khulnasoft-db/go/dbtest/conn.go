// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package dbtest

import (
	"net"
	"os"
	"sync"
	"testing"

	db "github.com/khulnasoft-com/khulnasoft/components/khulnasoft-db/go"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var (
	connLock = sync.Mutex{}
	conn     *gorm.DB
)

func ConnectForTests(t *testing.T) *gorm.DB {
	t.Helper()

	connLock.Lock()
	defer connLock.Unlock()

	if conn != nil {
		return conn
	}

	// These are static connection details for tests, started by `blazedock build components/khulnasoft-db/go:init-testdb`.
	// We use the same static credentials for CI & local instance of MySQL Server.
	var err error
	conn, err = db.Connect(db.ConnectionParams{
		User:     "root",
		Password: "test",
		Host:     net.JoinHostPort(os.Getenv("DB_HOST"), "23306"),
		Database: "khulnasoft",
	})
	require.NoError(t, err, "Failed to establish connection to  In a workspace, run `blazedock run components/khulnasoft-db:init-testdb` once to bootstrap the db")

	return conn
}
