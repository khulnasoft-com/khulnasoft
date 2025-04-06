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
    'khulnasoft-component-agent-smith.json': (import 'dashboards/components/agent-smith.json'),
    'khulnasoft-component-content-service.json': (import 'dashboards/components/content-service.json'),
    'khulnasoft-component-registry-facade.json': (import 'dashboards/components/registry-facade.json'),
    'khulnasoft-component-ws-daemon.json': (import 'dashboards/components/ws-daemon.json'),
    'khulnasoft-component-ws-manager-mk2.json': (import 'dashboards/components/ws-manager-mk2.json'),
    'khulnasoft-component-ws-proxy.json': (import 'dashboards/components/ws-proxy.json'),
    'khulnasoft-workspace-success-criteria.json': (import 'dashboards/success-criteria.json'),
    'khulnasoft-workspace-coredns.json': (import 'dashboards/coredns.json'),
    'khulnasoft-node-swap.json': (import 'dashboards/node-swap.json'),
    'khulnasoft-node-ephemeral-storage.json': (import 'dashboards/ephemeral-storage.json'),
    'khulnasoft-node-problem-detector.json': (import 'dashboards/node-problem-detector.json'),
    'khulnasoft-network-limiting.json': (import 'dashboards/network-limiting.json'),
    'khulnasoft-component-image-builder.json': (import 'dashboards/components/image-builder.json'),
    'khulnasoft-psi.json': (import 'dashboards/node-psi.json'),
    'khulnasoft-workspace-psi.json': (import 'dashboards/workspace-psi.json'),
    'khulnasoft-workspace-registry-facade-blobsource.json': (import 'dashboards/registry-facade-blobsource.json'),
  },
}
