// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	t.Run("with all options", func(t *testing.T) {
		expectedOptions := &options{
			url:         "https://foo.bar.com",
			client:      &http.Client{},
			credentials: "my_awesome_credentials",
		}
		khulnasoft, err := New(
			WithURL(expectedOptions.url),
			WithCredentials(expectedOptions.credentials),
			WithHTTPClient(expectedOptions.client),
		)
		require.NoError(t, err)
		require.Equal(t, expectedOptions, khulnasoft.cfg)

		require.NotNil(t, khulnasoft.PersonalAccessTokens)
		require.NotNil(t, khulnasoft.Workspaces)
		require.NotNil(t, khulnasoft.Projects)
		require.NotNil(t, khulnasoft.PersonalAccessTokens)
		require.NotNil(t, khulnasoft.User)
	})

	t.Run("fails when no credentials specified", func(t *testing.T) {
		_, err := New()
		require.Error(t, err)
	})

	t.Run("defaults to https://api.khulnasoft.com", func(t *testing.T) {
		khulnasoft, err := New(WithCredentials("foo"))
		require.NoError(t, err)

		require.Equal(t, "https://api.khulnasoft.com", khulnasoft.cfg.url)
	})

}
