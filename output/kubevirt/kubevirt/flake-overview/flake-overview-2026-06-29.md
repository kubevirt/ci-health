# Flake Overview Report — 2026-06-29

## Summary

| Window | Total Failures | Lanes | Failures/Day |
|--------|---------------|-------|--------------|
| 28d (Jun 1–29) | 2,058 | 45 | 73.5 |
| 7d (Jun 22–29) | 633 | 33 | 90.4 |

**Overall trend: Worsening** — daily failure rate increased ~23% in the last 7 days compared to the 28-day average.

**Top 3 worst SIGs by failure count (28d):**
1. **sig-compute** — 896 failures (43.5%) across periodic + presubmit compute/migration lanes
2. **sig-storage** — 403 failures (19.6%) — dominated by migration-cancel quarantine debt
3. **Platform-specific (S390X/arm64)** — 244 failures (11.9%)

**Quarantine candidates identified:** 2 high priority, 1 medium priority
**Quarantine debt (broken quarantined tests):** 5 tests at 0% success in 7d
**Infra-unstable lanes:** k8s-1.36 presubmit lanes show elevated Pod-level failure rates (73–79%)

### Key changes this week

- **Live Migration cancel migration tests dropped to 0% success** across all lanes in 7d — already quarantined but now a deterministic regression, not a flake
- **KWOK density tests worsened sharply** from ~88% to ~47-57% success rate — emerging regression, not yet quarantined
- **k8s-1.36 presubmit infrastructure instability** — Pod-level failures at 73-80% across most k8s-1.36 presubmit lanes

## Navigation

- [Infrastructure](#infrastructure)
- [sig-compute](#sig-compute)
- [sig-storage](#sig-storage)
- [sig-network](#sig-network)
- [sig-operator](#sig-operator)
- [sig-monitoring](#sig-monitoring)
- [sig-performance](#sig-performance)
- [Platform-specific](#platform-specific)
- [Quarantine Candidate Summary](#quarantine-candidate-summary)

---

## Infrastructure

### Presubmit Pod-Level Failure Rates

Pod-level failures represent infrastructure instability — builds that fail before any test code runs.

| Lane | 28d Success | 7d Success | Trend | Builds (28d) |
|------|------------|------------|-------|-------------|
| pull-kubevirt-e2e-k8s-1.36-sig-storage.Pod | 72.3% | 78.6% | Improving | 511 |
| pull-kubevirt-e2e-k8s-1.36-sig-network.Pod | 72.8% | 79.9% | Improving | 513 |
| pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations.Pod | 73.6% | 79.6% | Improving | 569 |
| pull-kubevirt-e2e-k8s-1.36-sig-operator.Pod | 73.6% | 79.0% | Improving | 489 |
| pull-kubevirt-e2e-k8s-1.36-sig-compute-serial.Pod | 74.3% | N/A | — | 484 |
| pull-kubevirt-e2e-k8s-1.36-sig-compute.Pod | 76.3% | 82.8% | Improving | 500 |
| pull-kubevirt-e2e-kind-1.35-sig-compute-arm64.Pod | 82.6% | 88.9% | Improving | 573 |
| pull-kubevirt-check-tests-for-flakes.Pod | 84.9% | 93.3% | Improving | 554 |
| pull-kubevirt-e2e-k8s-1.34-sig-monitoring.Pod | 70.3% | 81.3% | Improving | 75 |

**Assessment:** The k8s-1.36 presubmit lanes have the worst infrastructure reliability at 72-80% Pod-level success. While the 7d trend shows slight improvement across the board, one in five builds still fails at the infrastructure level on k8s-1.36 presubmit lanes. This wastes significant CI capacity and creates false-negative pressure on all test results in those lanes.

### Periodic Pod-Level Failure Rates

| Lane | 28d Success | 7d Success | Builds (28d) |
|------|------------|------------|-------------|
| periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations.Pod | 82.4% | N/A | 74 |
| periodic-kubevirt-e2e-k8s-1.36-sig-storage.Pod | 86.4% | N/A | 111 |
| periodic-kubevirt-e2e-k8s-1.36-ipv6-sig-network.Pod | 86.8% | N/A | 107 |
| periodic-kubevirt-e2e-k8s-1.34-sig-compute.Pod | 86.9% | N/A | 108 |
| periodic-kubevirt-e2e-k8s-1.34-sig-network.Pod | 87.9% | N/A | 108 |
| periodic-kubevirt-e2e-test-S390X.Pod | 88.2% | N/A | 111 |

Periodic Pod-level failures are moderate (82-89% success), with k8s-1.36 sig-compute-migrations being the worst.

---

## sig-compute

**Total 28d failures:** 896 (43.5%) | **7d failures:** 278 (43.9%) | **Trend:** Stable share

### Periodic Lanes

| Lane | 28d Failures | 28d Share | 7d Failures | 7d Share | Builds (28d) | Trend |
|------|-------------|-----------|-------------|----------|-------------|-------|
| [periodic-kubevirt-e2e-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-compute) | 344 | 16.72% | 32 | 5.06% | 108 | Improving |
| [periodic-kubevirt-e2e-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute) | 139 | 6.75% | 33 | 5.21% | 107 | Stable |
| [periodic-kubevirt-e2e-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute) | 131 | 6.37% | 31 | 4.90% | 108 | Stable |
| [periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations) | 116 | 5.64% | 63 | 9.95% | 74 | **Worsening** |

### Presubmit Lanes

| Lane | 28d Failures | 7d Failures | Builds (28d) | Trend |
|------|-------------|-------------|-------------|-------|
| [pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations) | 61 | 32 | 569 | **Worsening** |
| [pull-kubevirt-e2e-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute) | 45 | 13 | 500 | Stable |
| [pull-kubevirt-e2e-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute) | 16 | 7 | 189 | Stable |
| [pull-kubevirt-e2e-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-compute) | 10 | 4 | 180 | Stable |
| [pull-kubevirt-e2e-k8s-1.36-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute-serial) | 6 | N/A | 484 | — |
| [pull-kubevirt-check-tests-for-flakes](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-check-tests-for-flakes) | 6 | 4 | 554 | Stable |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.34-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [USB 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) [QUARANTINE] | 18.1% | 52.0% | Improving | Flaky | Dispersed (3 lanes) |
| [USB 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) [QUARANTINE] | 18.1% | 52.0% | Improving | Flaky | Dispersed (3 lanes) |
| [PCI passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L112) | 34.3% | N/A | Improving | Flaky | Concentrated (1 lane) |
| [PCI passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L111) | 34.3% | N/A | Improving | Flaky | Concentrated (1 lane) |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) [QUARANTINE] | 85.7% | 88.0% | Stable | Flaky | Dispersed (3 lanes) |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.36-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [USB 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) [QUARANTINE] | 48.5% | 45.8% | Stable | Flaky | Dispersed (3 lanes) |
| [USB 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) [QUARANTINE] | 48.5% | 45.8% | Stable | Flaky | Dispersed (3 lanes) |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) [QUARANTINE] | 87.4% | 91.7% | Stable | Flaky | Dispersed (3 lanes) |
| [ClusterProfiler pprof](https://github.com/kubevirt/kubevirt/blob/main/tests/infrastructure/cluster-profiler.go#L62) [QUARANTINE] | 93.2% | 91.7% | Stable | Flaky | Concentrated (1 lane) |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.35-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [USB 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) [QUARANTINE] | 45.2% | 46.2% | Stable | Flaky | Dispersed (3 lanes) |
| [USB 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) [QUARANTINE] | 45.2% | 46.2% | Stable | Flaky | Dispersed (3 lanes) |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) [QUARANTINE] | 89.4% | 92.3% | Stable | Flaky | Dispersed (3 lanes) |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [Cancel migration (delete source)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) [QUARANTINE] | 30.0% | **0.0%** | **Worsening** | Consecutive | Dispersed (4+ lanes) |
| [Cancel migration (delete target)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) [QUARANTINE] | 30.0% | **0.0%** | **Worsening** | Consecutive | Dispersed (4+ lanes) |
| [Post-copy migration liveness probe](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/postcopy.go#L176) [QUARANTINE] | 86.4% | 85.2% | Stable | Flaky | Concentrated (1 lane) |
| [Abort a vmi migration](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L1499) | 94.3% | 92.6% | Stable | Flaky | Dispersed (2 lanes) |

### Failing Tests — pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [Abort a vmi migration](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L1499) | 93.3% | 88.6% | **Worsening** | Flaky | Dispersed (2 lanes) |
| [Live Migrations with priority](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/priority.go#L76) | N/A | 94.0% | New | Flaky | Concentrated (1 lane) |

### Failing Tests — pull-kubevirt-check-tests-for-flakes

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [PCI Port Allocation (1Gi, 0 devs)](https://github.com/kubevirt/kubevirt/blob/main/tests/hotplug/pci-ports.go#L46) | 75.0% | N/A | — | Flaky | Concentrated (1 lane) |
| [PCI Port Allocation (2.1Gi, 6 devs)](https://github.com/kubevirt/kubevirt/blob/main/tests/hotplug/pci-ports.go#L46) | 75.0% | N/A | — | Flaky | Concentrated (1 lane) |

---

## sig-storage

**Total 28d failures:** 403 (19.6%) | **7d failures:** 166 (26.2%) | **Trend:** Worsening share

### Periodic Lanes

| Lane | 28d Failures | 28d Share | 7d Failures | 7d Share | Builds (28d) | Trend |
|------|-------------|-----------|-------------|----------|-------------|-------|
| [periodic-kubevirt-e2e-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage) | 130 | 6.32% | 57 | 9.00% | 111 | **Worsening** |
| [periodic-kubevirt-e2e-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-storage) | 129 | 6.27% | 50 | 7.90% | 109 | **Worsening** |
| [periodic-kubevirt-e2e-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-storage) | 120 | 5.83% | 50 | 7.90% | 108 | **Worsening** |

### Presubmit Lanes

| Lane | 28d Failures | 7d Failures | Builds (28d) | Trend |
|------|-------------|-------------|-------------|-------|
| [pull-kubevirt-e2e-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-storage) | 10 | 4 | 511 | Stable |
| [pull-kubevirt-e2e-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-storage) | 8 | 3 | 195 | Stable |
| [pull-kubevirt-e2e-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-storage) | 6 | 2 | 190 | Stable |

### Failing Tests — All sig-storage periodic lanes (identical pattern)

The sig-storage lanes run the migration cancel tests and these dominate all failures.

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [Cancel migration (delete source)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) [QUARANTINE] | 53.8% | **0.0%** | **Worsening** | Consecutive | Dispersed (4+ lanes) |
| [Cancel migration (delete target)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) [QUARANTINE] | 53.8–54.8% | **0.0%** | **Worsening** | Consecutive | Dispersed (4+ lanes) |

---

## sig-network

**Total 28d failures:** 133 (6.5%) | **7d failures:** 15 (2.4%) | **Trend:** Improving

### Periodic Lanes

| Lane | 28d Failures | 7d Failures | Builds (28d) | Trend |
|------|-------------|-------------|-------------|-------|
| [periodic-kubevirt-e2e-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network) | 21 | 3 | 109 | Stable |
| [periodic-kubevirt-e2e-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network) | 20 | 1 | 111 | Improving |
| [periodic-kubevirt-e2e-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-network) | 10 | 2 | 108 | Stable |
| [periodic-kubevirt-e2e-k8s-1.36-sig-network-with-dnc](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network-with-dnc) | 5 | 2 | 26 | Stable |
| [periodic-kubevirt-e2e-k8s-1.36-ipv6-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-ipv6-sig-network) | 2 | 1 | 107 | Stable |

### Presubmit Lanes

| Lane | 28d Failures | 7d Failures | Builds (28d) | Trend |
|------|-------------|-------------|-------------|-------|
| [pull-kubevirt-e2e-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-network) | 31 | 3 | 186 | Stable |
| [pull-kubevirt-e2e-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-network) | 26 | 3 | 513 | Stable |
| [pull-kubevirt-e2e-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-network) | 25 | N/A | 194 | — |
| [pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network) | 10 | N/A | 204 | — |

### Failing Tests

| Test | Lane(s) | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|---------|----------|---------|-------|---------|------------|
| [Passt binding migration connectivity IPv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) [QUARANTINE] | periodic 1.34/1.35/1.36 | 88.6–92.9% | 91.7–92.3% | Stable | Flaky | Dispersed (3 lanes) |
| [Bridge nic-hotplug multiple interfaces In place](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L215) | periodic 1.36-with-dnc | 84.6% | 71.4% | **Worsening** | Flaky | Concentrated (1 lane) |

---

## sig-operator

**Total 28d failures:** 239 (11.6%) | **7d failures:** 95 (15.0%) | **Trend:** Worsening share

### Lanes

| Lane | 28d Failures | 7d Failures | Builds (28d) | Trend |
|------|-------------|-------------|-------------|-------|
| [pull-kubevirt-e2e-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-operator) | 100 | 51 | 489 | **Worsening** |
| [pull-kubevirt-e2e-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-operator) | 47 | N/A | 180 | — |
| [periodic-kubevirt-e2e-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-operator) | 45 | N/A | 109 | — |
| [pull-kubevirt-e2e-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-operator) | 44 | 44 | 178 | Stable |
| [periodic-kubevirt-e2e-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-operator) | 3 | N/A | 108 | — |

**Assessment:** sig-operator failures are exclusively Pod-level (infrastructure) — no individual test failures were reported. The pull-kubevirt-e2e-k8s-1.36-sig-operator lane has the worst Pod-level success rate (73.6% 28d / 79.0% 7d), accounting for half of all sig-operator failures.

---

## sig-monitoring

**Total 28d failures:** 78 (3.8%) | **7d failures:** 37 (5.8%) | **Trend:** Worsening share

### Lanes

| Lane | 28d Failures | 7d Failures | Builds (28d) | Trend |
|------|-------------|-------------|-------------|-------|
| [periodic-kubevirt-e2e-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-monitoring) | 61 | 20 | 56 | Stable |
| [pull-kubevirt-e2e-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-monitoring) | 17 | 17 | 75 | Stable |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.34-sig-monitoring

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [VirtOperatorRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L340) [QUARANTINE] | 30.0% | 30.0% | Stable | Flaky | Concentrated (1 lane) |
| [Virt components metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L119) | 60.7% | 71.4% | Improving | Flaky | Dispersed (2 lanes) |
| [VM dirty rate metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/vm_monitoring.go#L334) [QUARANTINE] | 92.9% | 92.9% | Stable | Flaky | Concentrated (1 lane) |
| [LowReadyVirtHandlerCount](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L272) | N/A | 91.7% | New in 7d | Flaky | Concentrated (1 lane) |
| [VirtAPIDown](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L171) | N/A | 92.3% | New in 7d | Flaky | Concentrated (1 lane) |
| [VirtControllerDown](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L149) | N/A | 92.9% | New in 7d | Flaky | Concentrated (1 lane) |
| [VM count metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/vm_monitoring.go#L393) | N/A | 92.9% | New in 7d | Flaky | Concentrated (1 lane) |
| [Orphaned VMIs alert](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/vm_monitoring.go#L510) | N/A | 92.9% | New in 7d | Flaky | Concentrated (1 lane) |

### Failing Tests — pull-kubevirt-e2e-k8s-1.34-sig-monitoring

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [Virt components metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L119) | 29.1% | 54.2% | Improving | Flaky | Dispersed (2 lanes) |
| [VNIC metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L132) | 94.1% | 92.9% | Stable | Flaky | Concentrated (1 lane) |

**Note:** The "virt components metrics" test has been quarantined in code ([QUARANTINE] tag added) but may appear without the tag in older flakefinder data. Its success rate is improving from 29-61% to 54-71%.

---

## sig-performance

**Total 28d failures:** 36 (1.7%) | **7d failures:** 32 (5.1%) | **Trend:** Worsening

### Lanes

| Lane | 28d Failures | 7d Failures | Builds (28d) | Trend |
|------|-------------|-------------|-------------|-------|
| [periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok) | 20 | 16 | 78 | **Worsening** |
| [periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok-100](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok-100) | 16 | 16 | 84 | **Worsening** |

### Failing Tests

| Test | Lane | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|------|----------|---------|-------|---------|------------|
| [KWOK fake VMIs](https://github.com/kubevirt/kubevirt/blob/main/tests/performance/density-kwok.go#L89) | kwok | 87.8% | **47.1%** | **Worsening** | Burst/regression | Dispersed (2 lanes) |
| [KWOK fake VMs](https://github.com/kubevirt/kubevirt/blob/main/tests/performance/density-kwok.go#L105) | kwok | 87.8% | **47.1%** | **Worsening** | Burst/regression | Dispersed (2 lanes) |
| [KWOK fake VMIs](https://github.com/kubevirt/kubevirt/blob/main/tests/performance/density-kwok.go#L89) | kwok-100 | 89.2% | **57.1%** | **Worsening** | Burst/regression | Dispersed (2 lanes) |
| [KWOK fake VMs](https://github.com/kubevirt/kubevirt/blob/main/tests/performance/density-kwok.go#L105) | kwok-100 | 89.2% | **57.1%** | **Worsening** | Burst/regression | Dispersed (2 lanes) |

**Assessment:** The KWOK density tests have degraded sharply this week — dropping from ~88-89% to 47-57% success rate. Both the VMI and VM variants fail together, and both kwok and kwok-100 lanes are affected, suggesting a regression in the KWOK density test infrastructure or the control plane performance that these tests measure. These tests are **not quarantined** and should be investigated for a recent regression. 9 out of 17-21 runs are failing in the 7d window.

---

## Platform-specific

### S390X

**28d failures:** 166 (8.07%) | **7d failures:** 37 (5.85%) | **Trend:** Improving

| Lane | 28d Failures | 7d Failures | Builds (28d) | Trend |
|------|-------------|-------------|-------------|-------|
| [periodic-kubevirt-e2e-test-S390X](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-test-S390X) | 166 | 37 | 111 | Improving |

#### Failing Tests

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [Eviction strategy block and migrate](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | 56.1% | N/A | Improving | Burst |
| [Virt-handler crash recovery](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_lifecycle_test.go#L492) | N/A | 84.6% | New in 7d | Flaky |
| [VMPool scale-in offline](https://github.com/kubevirt/kubevirt/blob/main/tests/pool_test.go#L881) | N/A | 84.6% | New in 7d | Flaky |
| [Multi-migration cloud-init](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L551) | 84.1% | 88.5% | Stable | Flaky |
| [Virtqemud restart migration](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L589) | 91.6% | N/A | Improving | Flaky |
| [CfgMap/Secret fs layout](https://github.com/kubevirt/kubevirt/blob/main/tests/config_test.go#L520) | 94.4% | N/A | Improving | Flaky |
| [ReplicaSet server-side printing](https://github.com/kubevirt/kubevirt/blob/main/tests/replicaset_test.go#L236) | N/A | 92.3% | New in 7d | Flaky |
| [Masquerade MTU IPv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/vmi_networking.go#L720) | N/A | 92.3% | New in 7d | Flaky |

**Assessment:** S390X is improving overall. The worst 28d test (eviction strategy at 56%) did not appear in 7d top failures, suggesting a fix or transient burst. Several new low-severity flakes appeared in 7d (84-92% range) — worth monitoring but not quarantine-worthy yet since they're all concentrated in the S390X lane (platform-specific).

### ARM64

| Lane | 28d Failures | 7d Failures | Builds (28d) | Trend |
|------|-------------|-------------|-------------|-------|
| [pull-kubevirt-e2e-kind-1.35-sig-compute-arm64](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-sig-compute-arm64) | 65 | 25 | 573 | Improving |

ARM64 failures are exclusively Pod-level (82.6% 28d → 88.9% 7d, improving).

### SEV

| Lane | 28d Failures | 7d Failures | Builds (28d) | Trend |
|------|-------------|-------------|-------------|-------|
| [periodic-kubevirt-e2e-kind-1.34-sev](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-kind-1.34-sev) | 5 | 3 | 83 | Stable |

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [SEV guest attestation](https://github.com/kubevirt/kubevirt/blob/main/tests/launchsecurity/sev.go#L349) [QUARANTINE] | 93.9% | 85.0% | **Worsening** | Flaky |

### Other Platform Lanes

| Lane | 28d Failures | 7d Failures | Builds (28d) |
|------|-------------|-------------|-------------|
| [pull-kubevirt-e2e-kind-sriov](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-sriov) | 7 | 3 | 198 |
| [pull-kubevirt-e2e-kind-1.35-vgpu](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-vgpu) | 1 | N/A | 219 |
| [pull-kubevirt-e2e-k8s-1.34-windows2016](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-windows2016) | 1 | N/A | 221 |

---

## Quarantine Candidate Summary

### New quarantine candidates (not yet quarantined)

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority |
|------|-----|-------|-------------------|----------|---------|------------|----------|
| [KWOK fake VMIs](https://github.com/kubevirt/kubevirt/blob/main/tests/performance/density-kwok.go#L89) | sig-performance | kwok, kwok-100 | 87.8% | **Worsening to 47%** | Burst/regression | Dispersed (2) | **High** |
| [KWOK fake VMs](https://github.com/kubevirt/kubevirt/blob/main/tests/performance/density-kwok.go#L105) | sig-performance | kwok, kwok-100 | 87.8% | **Worsening to 47%** | Burst/regression | Dispersed (2) | **High** |
| [PCI passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L112) | sig-compute | periodic 1.34 | 34.3% | Improving (absent in 7d) | Flaky | Concentrated (1) | Medium |
| [PCI passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L111) | sig-compute | periodic 1.34 | 34.3% | Improving (absent in 7d) | Flaky | Concentrated (1) | Medium |
| [Abort a vmi migration](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L1499) | sig-compute | periodic + presubmit migrations | 93.3% | **Worsening to 88.6%** | Flaky | Dispersed (2) | Low |

### Quarantine debt (already quarantined, still severely broken)

These tests have been quarantined but remain at critically low success rates — they have never been fixed.

| Test | SIG | Worst Rate (7d) | Lanes Affected | Status |
|------|-----|-----------------|----------------|--------|
| [Cancel migration (delete source)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) | sig-compute/storage | **0.0%** | 4+ lanes | Deterministic regression — needs urgent investigation |
| [Cancel migration (delete target)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) | sig-compute/storage | **0.0%** | 4+ lanes | Deterministic regression — needs urgent investigation |
| [USB 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) | sig-compute | 45.8–52.0% | 3 lanes | Long-standing, stable |
| [USB 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) | sig-compute | 45.8–52.0% | 3 lanes | Long-standing, stable |
| [VirtOperatorRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L340) | sig-monitoring | 30.0% | 1 lane | Long-standing, stable |

---

## Suggested Follow-Up

- **lane-failure-rate** on [periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations) — the cancel-migration tests dropped to 0% in 7d; deeper per-build analysis will confirm if this is a deterministic regression or correlated infra failure
- **lane-failure-rate** on [periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok) — KWOK density tests worsened sharply; flip-rate analysis will distinguish regression from infra instability
- **analyze-build** on recent failing builds from the k8s-1.36 presubmit lanes — high Pod-level failure rates (73-80%) suggest infrastructure issues worth root-causing
