// Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package client

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	khulnasoft_experimental_v1connect "github.com/khulnasoft-com/khulnasoft/components/public-api/go/experimental/v1/v1connect"
)

type Khulnasoft struct {
	cfg *options

	Workspaces           khulnasoft_experimental_v1connect.WorkspacesServiceClient
	Editors              khulnasoft_experimental_v1connect.EditorServiceClient
	Teams                khulnasoft_experimental_v1connect.TeamsServiceClient
	Projects             khulnasoft_experimental_v1connect.ProjectsServiceClient
	PersonalAccessTokens khulnasoft_experimental_v1connect.TokensServiceClient
	IdentityProvider     khulnasoft_experimental_v1connect.IdentityProviderServiceClient
	User                 khulnasoft_experimental_v1connect.UserServiceClient
}

func New(options ...Option) (*Khulnasoft, error) {
	opts, err := evaluateOptions(defaultOptions(), options...)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate client options: %w", err)
	}

	if opts.credentials == "" {
		return nil, errors.New("no authentication credentials specified")
	}

	client := opts.client
	url := opts.url

	serviceOpts := []connect.ClientOption{
		connect.WithInterceptors(
			AuthorizationInterceptor(opts.credentials),
		),
	}

	return &Khulnasoft{
		cfg:                  opts,
		Teams:                khulnasoft_experimental_v1connect.NewTeamsServiceClient(client, url, serviceOpts...),
		Projects:             khulnasoft_experimental_v1connect.NewProjectsServiceClient(client, url, serviceOpts...),
		PersonalAccessTokens: khulnasoft_experimental_v1connect.NewTokensServiceClient(client, url, serviceOpts...),
		Workspaces:           khulnasoft_experimental_v1connect.NewWorkspacesServiceClient(client, url, serviceOpts...),
		Editors:              khulnasoft_experimental_v1connect.NewEditorServiceClient(client, url, serviceOpts...),
		IdentityProvider:     khulnasoft_experimental_v1connect.NewIdentityProviderServiceClient(client, url, serviceOpts...),
		User:                 khulnasoft_experimental_v1connect.NewUserServiceClient(client, url, serviceOpts...),
	}, nil
}

type Option func(opts *options) error

func WithURL(url string) Option {
	return func(opts *options) error {
		opts.url = url
		return nil
	}
}

func WithCredentials(token string) Option {
	return func(opts *options) error {
		opts.credentials = token
		return nil
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(opts *options) error {
		opts.client = client
		return nil
	}
}

type options struct {
	url         string
	client      *http.Client
	credentials string
}

func defaultOptions() *options {
	return &options{
		url:    "https://api.khulnasoft.com",
		client: http.DefaultClient,
	}
}

func evaluateOptions(base *options, opts ...Option) (*options, error) {
	for _, opt := range opts {
		if err := opt(base); err != nil {
			return nil, fmt.Errorf("failed to evaluate options: %w", err)
		}
	}

	return base, nil
}
