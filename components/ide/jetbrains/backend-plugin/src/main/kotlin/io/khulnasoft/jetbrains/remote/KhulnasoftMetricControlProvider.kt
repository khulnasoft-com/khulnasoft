// Copyright (c) 2024 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.khulnasoft.jetbrains.remote

import com.jetbrains.ide.model.uiautomation.BeControl
import com.jetbrains.ide.model.uiautomation.DefiniteProgress
import com.jetbrains.rd.platform.codeWithMe.unattendedHost.metrics.Metric
import com.jetbrains.rd.ui.bedsl.dsl.VerticalGridBuilder
import com.jetbrains.rd.ui.bedsl.dsl.withMargin
import com.jetbrains.rd.util.lifetime.Lifetime
import com.jetbrains.rd.util.reactive.Property
import com.jetbrains.rdserver.unattendedHost.customization.controlCenter.performance.createProgressRow
import com.jetbrains.rdserver.unattendedHost.diagnostics.BackendDiagnosticsService
import io.khulnasoft.jetbrains.remote.AbstractKhulnasoftMetricControlProvider
import io.khulnasoft.jetbrains.remote.IBackendDiagnosticsService

class KhulnasoftMetricControlProvider: AbstractKhulnasoftMetricControlProvider() {
    override fun setMargin(element: BeControl, left: Int, top: Int, right: Int, bottom: Int) = element.withMargin(left, top, right, bottom)

    override fun getBackendDiagnosticsService(): IBackendDiagnosticsService {
        val obj = BackendDiagnosticsService.Companion.getInstance()
        return object : IBackendDiagnosticsService {
            override fun getMetric(name: String): Metric {
                return obj.getMetric(name)
            }
        }
    }

    override fun createProgressControl(ctx: VerticalGridBuilder, lifetime: Lifetime, label: String, cpuPercentage: Metric, labelProperty: Property<String>, cpuPercentageProperty: Property<String>, progressBar: DefiniteProgress) {
        createProgressRow(ctx, id, lifetime, label, cpuPercentage.statusProperty, labelProperty, cpuPercentageProperty, progressBar)
    }

}
