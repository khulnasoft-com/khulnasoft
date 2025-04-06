// Copyright (c) 2020 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package khulnasoft

import (
	"context"

	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/khulnasoft-com/khulnasoft/common-go/util"
	serverapi "github.com/khulnasoft-com/khulnasoft/khulnasoft-protocol"
	supervisor "github.com/khulnasoft-com/khulnasoft/supervisor/api"
)

var (
	// Version - set during build
	Version = "dev"
)

func GetWSInfo(ctx context.Context) (*supervisor.WorkspaceInfoResponse, error) {
	supervisorConn, err := grpc.Dial(util.GetSupervisorAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, xerrors.Errorf("failed connecting to supervisor: %w", err)
	}
	defer supervisorConn.Close()
	wsinfo, err := supervisor.NewInfoServiceClient(supervisorConn).WorkspaceInfo(ctx, &supervisor.WorkspaceInfoRequest{})
	if err != nil {
		return nil, xerrors.Errorf("failed getting workspace info from supervisor: %w", err)
	}
	return wsinfo, nil
}

func ConnectToServer(ctx context.Context, wsInfo *supervisor.WorkspaceInfoResponse, scope []string) (*serverapi.APIoverJSONRPC, error) {
	supervisorConn, err := grpc.Dial(util.GetSupervisorAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, xerrors.Errorf("failed connecting to supervisor: %w", err)
	}
	defer supervisorConn.Close()
	clientToken, err := supervisor.NewTokenServiceClient(supervisorConn).GetToken(ctx, &supervisor.GetTokenRequest{
		Host:  wsInfo.KhulnasoftApi.Host,
		Kind:  "khulnasoft",
		Scope: scope,
	})
	if err != nil {
		return nil, xerrors.Errorf("failed getting token from supervisor: %w", err)
	}

	client, err := serverapi.ConnectToServer(wsInfo.KhulnasoftApi.Endpoint, serverapi.ConnectToServerOpts{
		Token:   clientToken.Token,
		Context: ctx,
		Log:     log.NewEntry(log.StandardLogger()),
		ExtraHeaders: map[string]string{
			"User-Agent":       "khulnasoft/cli",
			"X-Client-Version": Version,
		},
	})
	if err != nil {
		return nil, xerrors.Errorf("failed connecting to server: %w", err)
	}
	return client, nil
}
