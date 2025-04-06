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
    'khulnasoft-component-blobserve.json': (import 'dashboards/components/blobserve.json'),
    'khulnasoft-component-openvsx-proxy.json': (import 'dashboards/components/openvsx-proxy.json'),
    'khulnasoft-component-openvsx-mirror.json': (import 'dashboards/components/openvsx-mirror.json'),
    'khulnasoft-component-ssh-gateway.json': (import 'dashboards/components/ssh-gateway.json'),
    'khulnasoft-component-supervisor.json': (import 'dashboards/components/supervisor.json'),
    'khulnasoft-component-jb.json': (import 'dashboards/components/jb.json'),
    'khulnasoft-component-browser-overview.json': (import 'dashboards/components/browser-overview.json'),
    'khulnasoft-component-code-browser.json': (import 'dashboards/components/code-browser.json'),
    'khulnasoft-component-ide-startup-time.json': (import 'dashboards/components/ide-startup-time.json'),
    'khulnasoft-component-ide-service.json': (import 'dashboards/components/ide-service.json'),
    'khulnasoft-component-local-ssh.json': (import 'dashboards/components/local-ssh.json'),
  },
}
