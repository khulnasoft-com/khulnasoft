/**
 * Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

{
  grafanaDashboards+:: {
    // Import raw json files here.
    // Example:
    // 'my-new-dashboard.json': (import 'dashboards/components/new-component.json'),
    'khulnasoft-component-dashboard.json': (import 'dashboards/components/dashboard.json'),
    'khulnasoft-component-db.json': (import 'dashboards/components/db.json'),
    'khulnasoft-component-ws-manager-bridge.json': (import 'dashboards/components/ws-manager-bridge.json'),
    'khulnasoft-component-proxy.json': (import 'dashboards/components/proxy.json'),
    'khulnasoft-component-server.json': (import 'dashboards/components/server.json'),
    'khulnasoft-component-server-garbage-collector.json': (import 'dashboards/components/server-garbage-collector.json'),
    'khulnasoft-component-usage.json': (import 'dashboards/components/usage.json'),
    'khulnasoft-slo-login.json': (import 'dashboards/SLOs/login.json'),
    'khulnasoft-meta-overview.json': (import 'dashboards/components/meta-overview.json'),
    'khulnasoft-meta-services.json': (import 'dashboards/components/meta-services.json'),
    'khulnasoft-components-spicedb.json': (import 'dashboards/components/spicedb.json'),
  },
}
