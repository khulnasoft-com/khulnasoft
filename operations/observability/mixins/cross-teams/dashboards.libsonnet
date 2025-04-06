/**
 * Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

(import './dashboards/SLOs/workspace-startup-time.libsonnet') +
{
  grafanaDashboards+:: {
    // Import raw json files here.
    // Example:
    // 'my-new-dashboard.json': (import 'dashboards/components/new-component.json'),
    'khulnasoft-cluster-autoscaler-k3s.json': (import 'dashboards/khulnasoft-cluster-autoscaler-k3s.json'),
    'khulnasoft-node-resource-metrics.json': (import 'dashboards/khulnasoft-node-resource-metrics.json'),
    'khulnasoft-grpc-server.json': (import 'dashboards/khulnasoft-grpc-server.json'),
    'khulnasoft-grpc-client.json': (import 'dashboards/khulnasoft-grpc-client.json'),
    'khulnasoft-connect-server.json': (import 'dashboards/khulnasoft-connect-server.json'),
    'khulnasoft-overview.json': (import 'dashboards/khulnasoft-overview.json'),
    'khulnasoft-nodes-overview.json': (import 'dashboards/khulnasoft-nodes-overview.json'),
    'khulnasoft-admin-node.json': (import 'dashboards/khulnasoft-admin-node.json'),
    'khulnasoft-admin-workspace.json': (import 'dashboards/khulnasoft-admin-workspace.json'),
    'khulnasoft-applications.json': (import 'dashboards/khulnasoft-applications.json'),
    'redis.json': (import 'dashboards/redis.json')
  },
}
