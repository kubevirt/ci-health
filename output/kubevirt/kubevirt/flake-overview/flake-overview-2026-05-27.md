# Flake Overview Report - 2026-05-27

## Summary

| Metric | 28-day | 7-day | Trend |
|--------|--------|-------|-------|
| **Total failures** | 2949 | 523 | Stable (7d proportionally ~25% of 28d, as expected) |
| **Lanes with failures** | 42 | 31 | Stable |
| **Analysis window** | Apr 29 - May 27 | May 20 - May 27 | |

**Overall trend**: Stable. The 7-day failure count (523) is proportionally consistent with the 28-day count (2949), showing no significant improvement or degradation.

**Top 3 worst SIGs by failure count (28d)**:
1. **sig-compute** - 1,764 failures (59.8%) - dominated by virtiofs, USB passthrough, VSOCK, and PCI device tests
2. **sig-storage** - 538 failures (18.2%) - dominated by IO error test and volume hotplug migration tests
3. **Platform-specific (S390X)** - 301 failures (10.2%) - broad instability across migration, network, and MemBalloon tests

**Quarantine candidates identified**: 4 high priority, 3 medium priority
**Infra-unstable lanes**: Multiple presubmit lanes show Pod-level failure rates of 15-22%

## Navigation

- [Infrastructure](#infrastructure)
- [sig-compute](#sig-compute)
- [sig-storage](#sig-storage)
- [sig-monitoring](#sig-monitoring)
- [sig-network](#sig-network)
- [sig-operator](#sig-operator)
- [sig-performance](#sig-performance)
- [Platform-specific](#platform-specific)
- [Quarantine Candidate Summary](#quarantine-candidate-summary)
- [Quarantine Debt](#quarantine-debt)

---

## Infrastructure

### Presubmit Pod-Level Failure Rates

Pod-level failures indicate infrastructure problems (cluster provisioning, node readiness, test framework setup) rather than individual test bugs.

| Lane | 28d Pod Rate | 7d Pod Rate | Trend |
|------|-------------|-------------|-------|
| pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations | 78.0% | 78.0% | Stable |
| pull-kubevirt-e2e-k8s-1.35-sig-storage | 77.1% | 77.9% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-storage | 77.5% | 77.4% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-compute-serial | 80.8% | 77.8% | Worsening |
| pull-kubevirt-e2e-k8s-1.36-sig-compute | 81.2% | 78.1% | Worsening |
| pull-kubevirt-e2e-k8s-1.35-sig-network | 81.6% | 79.6% | Worsening |
| pull-kubevirt-e2e-k8s-1.35-sig-compute-serial | 80.8% | 79.9% | Stable |
| pull-kubevirt-e2e-k8s-1.35-sig-operator | 81.2% | 80.3% | Stable |
| pull-kubevirt-e2e-k8s-1.35-sig-compute | 82.8% | N/A | N/A |
| pull-kubevirt-e2e-k8s-1.36-sig-operator | 82.6% | 78.3% | Worsening |
| pull-kubevirt-e2e-kind-1.35-sig-compute-arm64 | 84.8% | 87.9% | Improving |
| pull-kubevirt-e2e-k8s-1.34-sig-performance | 85.4% | N/A | N/A |
| pull-kubevirt-e2e-k8s-1.34-sig-monitoring | 86.7% | N/A | N/A |
| pull-kubevirt-e2e-k8s-1.34-sig-storage | 90.4% | N/A | N/A |
| pull-kubevirt-e2e-k8s-1.34-sig-operator | 93.9% | N/A | N/A |
| pull-kubevirt-e2e-k8s-1.34-sig-network | 94.1% | N/A | N/A |

Several presubmit lanes have Pod success rates below 80%, meaning ~1 in 5 runs fail before tests even start. The k8s-1.35 and k8s-1.36 lanes are consistently worse than k8s-1.34, suggesting infrastructure issues with newer cluster versions.

### Setup Failures

| Lane | Issue | 28d | 7d |
|------|-------|-----|-----|
| periodic-kubevirt-e2e-test-S390X | AfterSuite at 0% (1/1 run) | Present | Present |

---

## sig-compute

**Total failures (28d)**: 1,764 (59.8% of all failures)
**Total failures (7d)**: ~316 (~60.4%)

### Lanes

| Lane | 28d Failures | 28d Share | 7d Failures | 7d Share | Trend |
|------|-------------|-----------|-------------|----------|-------|
| [periodic-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute) | 1062 | 36.01% | 52 | 9.94% | Improving |
| [periodic-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-compute) | 298 | 10.11% | 110 | 21.03% | Worsening |
| [periodic-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute) | 243 | 8.24% | 41 | 7.84% | Stable |
| [pull-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute) | 41 | 1.39% | 27 | 5.16% | Worsening |
| [pull-k8s-1.35-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-serial) | 32 | 1.09% | 6 | 1.15% | Stable |
| [pull-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | 24 | 0.81% | 6 | 1.15% | Stable |
| [periodic-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | 11 | 0.37% | 4 | 0.76% | Stable |
| [pull-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-compute) | 12 | 0.41% | 3 | 0.57% | Stable |
| [pull-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute) | 6 | 0.20% | N/A | N/A | N/A |
| [pull-k8s-1.36-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute-serial) | 6 | 0.20% | 3 | 0.57% | Stable |
| [periodic-kind-1.34-sev](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-kind-1.34-sev) | 6 | 0.20% | 2 | 0.38% | Stable |

### Failing Tests

#### virtiofs - ServiceAccount token (100% failure rate)

| Test | 28d Rate | 7d Rate | Trend | Lanes | Pattern | Dispersion |
|------|----------|---------|-------|-------|---------|------------|
| [Should be the namespace and token the same for a pod and vmi with virtiofs](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L232) | 0% (104/104 failed) | 0% (25/25 failed) | Stable | periodic-1.36-compute, pull-1.36-compute | **Consecutive** | 2 lanes |

This test has a **100% failure rate** across the entire 28-day window in both periodic and presubmit k8s-1.36 lanes. This is a deterministic regression, not a flake. It does not appear in k8s-1.34 or k8s-1.35, so this is specific to k8s-1.36. Presubmit data also confirms: 188/188 failed (28d), 92/92 failed (7d).

#### VSOCK tests (65-84% success rate)

| Test | 28d Rate | 7d Rate | Trend | Lanes | Dispersion |
|------|----------|---------|-------|-------|------------|
| [should retain the CID for migration target](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L121) | 65.4% | N/A | N/A | periodic-1.36-compute | 1 lane |
| [should succeed with TLS on both sides](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L223) | 72.1% | N/A | N/A | periodic-1.36-compute | 1 lane |
| [should return err if no app listerns on the port](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L240) | 73.1% | N/A | N/A | periodic-1.36-compute | 1 lane |
| [should succeed without TLS on both sides](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L224) | 76.0% | N/A | N/A | periodic-1.36-compute | 1 lane |
| [VM creation - Use virtio non-transitional](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L95) | 76.9% | N/A | N/A | periodic-1.36-compute | 1 lane |
| [VM creation - Use virtio transitional](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L94) | 77.9% | N/A | N/A | periodic-1.36-compute | 1 lane |
| [should return err if the port is invalid](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L227) | 83.7% | N/A | N/A | periodic-1.36-compute | 1 lane |

All VSOCK failures are concentrated in `periodic-kubevirt-e2e-k8s-1.36-sig-compute` only. They do not appear in the 7-day flakefinder data as individual entries. These are likely a common root cause — a VSOCK infrastructure issue specific to k8s-1.36 periodic lane. The shared lane concentration suggests **infra-correlated** failures rather than individual test bugs.

#### HostDevices - PCI passthrough (worsening)

| Test | 28d Rate | 7d Rate | Trend | Lanes | Pattern | Dispersion |
|------|----------|---------|-------|-------|---------|------------|
| [Should successfully passthrough an emulated PCI device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | 76.2% | 3.8% | **Worsening** | periodic-1.34-compute | 1 lane | Concentrated |
| [Should successfully passthrough 2 emulated PCI devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | 76.2% | 3.8% | **Worsening** | periodic-1.34-compute, pull-1.36-compute | 2 lanes | Concentrated |

These have dramatically worsened in the last 7 days on k8s-1.34 (from 76% to 4% success). The 7-day data shows 25/26 failures. This pattern (sudden near-total failure) indicates a recent regression rather than a flake. Also appearing in pull-1.36-compute at 94.6% success (7d), so the 1.34 lane is uniquely affected.

#### Guest console log flood [QUARANTINE]

| Test | 28d Rate | 7d Rate | Trend | Lanes | Pattern | Dispersion |
|------|----------|---------|-------|-------|---------|------------|
| [should not skip any log line even trying to flood the serial console for QOSGuaranteed VMs](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) | 81.7-87.5% | 84-88.5% | Stable | periodic-1.34/1.35/1.36-compute | **Flaky** | 3 lanes (dispersed) |

Already quarantined. Dispersed across all 3 k8s versions. Genuine flake at ~85% success rate.

#### USB Passthrough [QUARANTINE]

| Test | 28d Rate | 7d Rate | Trend | Lanes | Pattern | Dispersion |
|------|----------|---------|-------|-------|---------|------------|
| [Should successfully passthrough 1 emulated USB device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | 40-62.5% | 0-64% | Worsening in 1.34 | periodic-1.34/1.35/1.36-compute | **Flaky** | 3 lanes (dispersed) |
| [Should successfully passthrough 2 emulated USB devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | 40-62.5% | 0-64% | Worsening in 1.34 | periodic-1.34/1.35/1.36-compute | **Flaky** | 3 lanes (dispersed) |

Already quarantined. Severely broken across all k8s versions, with k8s-1.34 dropping to 0% success in the 7-day window.

#### SEV guest attestation [QUARANTINE]

| Test | 28d Rate | 7d Rate | Trend | Lanes | Dispersion |
|------|----------|---------|-------|-------|------------|
| [should run guest attestation](https://github.com/kubevirt/kubevirt/blob/main/tests/launchsecurity/sev.go#L349) | 89.2% | 85.7% | Stable | periodic-kind-1.34-sev | 1 lane |

Already quarantined. Concentrated in one specialized lane. Low impact.

#### Post-copy migration liveness probe (7d only)

| Test | 28d Rate | 7d Rate | Trend | Lanes | Dispersion |
|------|----------|---------|-------|-------|------------|
| [should complete post-copy migration without the liveness probe killing the source pod](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/postcopy.go#L174) | N/A | 94.1% | N/A | pull-1.35-sig-compute-migrations | 1 lane |

Only appeared in the 7d presubmit data. 6/102 failures. Inconclusive — monitor.

---

## sig-storage

**Total failures (28d)**: 538 (18.2% of all failures)
**Total failures (7d)**: ~103 (~19.7%)

### Lanes

| Lane | 28d Failures | 28d Share | 7d Failures | 7d Share | Trend |
|------|-------------|-----------|-------------|----------|-------|
| [periodic-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-storage) | 252 | 8.55% | 15 | 2.87% | Improving |
| [periodic-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage) | 132 | 4.48% | 41 | 7.84% | Worsening |
| [periodic-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-storage) | 80 | 2.71% | 16 | 3.06% | Stable |
| [pull-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-storage) | 62 | 2.10% | 25 | 4.78% | Worsening |
| [pull-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-storage) | 6 | 0.20% | 2 | 0.38% | Stable |
| [pull-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-storage) | 6 | 0.20% | 4 | 0.76% | Stable |

### Failing Tests

#### IO error pause [QUARANTINE]

| Test | 28d Rate | 7d Rate | Trend | Lanes | Pattern | Dispersion |
|------|----------|---------|-------|-------|---------|------------|
| [should pause VMI on IO error](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | 37.1-44.8% | 40.7-44.4% | Stable | periodic-1.34/1.35/1.36-storage | **Flaky** | 3 lanes (dispersed) |

Already quarantined. Consistently failing at ~40% success across all k8s versions. Genuine flake — long-standing, no trend improvement.

#### Volume hotplug migration (worsening in presubmit)

| Test | 28d Rate | 7d Rate | Trend | Lanes | Pattern | Dispersion |
|------|----------|---------|-------|-------|---------|------------|
| [block to filesystem](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1275) | 64.8% (periodic), 65.9% (presubmit) | 63.0% (periodic), 70.5% (presubmit) | Stable | periodic-1.36-storage, pull-1.36-storage | **Flaky** | 2 lanes |
| [filesystem to filesystem](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1273) | 81.9% (periodic), 65.9% (presubmit) | 85.2% (periodic), 69.3% (presubmit) | Stable | periodic-1.36-storage, pull-1.36-storage | **Flaky** | 2 lanes |
| [containerdisk and hotplugged volume](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1132) | 89.5% (periodic), 82.1% (presubmit) | 70.4% (periodic), 75.0% (presubmit) | **Worsening** | periodic-1.36-storage, pull-1.36-storage | **Flaky** | 2 lanes |

These three volume migration hotplug tests are concentrated in k8s-1.36 only (both periodic and presubmit). Not quarantined. The "containerdisk and hotplugged volume" test is actively worsening (89.5% -> 70.4% periodic). All three are strong quarantine candidates.

#### Migration cancel [QUARANTINE] (7d only absent from flakefinder)

| Test | 28d Rate | 7d Rate | Trend | Lanes | Dispersion |
|------|----------|---------|-------|-------|------------|
| [should be able to cancel a migration by deleting the migration resource - delete source migration](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L591) | 92.3% | N/A | N/A | periodic-1.35-storage | 1 lane |

Already quarantined. Low failure rate (92.3%), absent from 7d data. May be improving.

---

## sig-monitoring

**Total failures (28d)**: 96 (3.3% of all failures)
**Total failures (7d)**: 20 (3.8%)

### Lanes

| Lane | 28d Failures | 28d Share | 7d Failures | 7d Share | Trend |
|------|-------------|-----------|-------------|----------|-------|
| [periodic-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-monitoring) | 94 | 3.19% | 20 | 3.82% | Stable |
| [pull-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-monitoring) | 2 | 0.07% | N/A | N/A | N/A |

### Failing Tests

#### Prometheus VMI namespace metric [QUARANTINE]

| Test | 28d Rate | 7d Rate | Trend | Lanes | Pattern | Dispersion |
|------|----------|---------|-------|-------|---------|------------|
| [should find VMI namespace on namespace label of the metric](https://github.com/kubevirt/kubevirt/blob/main/tests/infrastructure/prometheus.go#L95) | 0% (54/54) | 0% (14/14) | Stable | periodic-1.34-monitoring | **Consecutive** | 1 lane |

Already quarantined. 100% failure rate across the entire 28d window. This is a deterministic breakage, not a flake. Only in the k8s-1.34 monitoring lane.

#### VM dirty rate [QUARANTINE]

| Test | 28d Rate | 7d Rate | Trend | Lanes | Dispersion |
|------|----------|---------|-------|-------|------------|
| [should ensure a stress VM has high dirty rate than a stale VM](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/vm_monitoring.go#L334) | 72.2% | 92.9% | **Improving** | periodic-1.34-monitoring | 1 lane |

Already quarantined. Significantly improved in 7d (72% -> 93%). A fix may have landed.

#### VM deletion transition time

| Test | 28d Rate | 7d Rate | Trend | Lanes | Dispersion |
|------|----------|---------|-------|-------|------------|
| [should contain transition time from VM deletion](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L179) | 83.0% | 85.7% | Stable | periodic-1.34-monitoring | 1 lane |

Not quarantined. Concentrated in one lane. Low failure count — monitor.

---

## sig-network

**Total failures (28d)**: ~55 (~1.9% of all failures)
**Total failures (7d)**: ~13 (~2.5%)

### Lanes

| Lane | 28d Failures | 28d Share | 7d Failures | 7d Share | Trend |
|------|-------------|-----------|-------------|----------|-------|
| [periodic-k8s-1.35-ipv6-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-ipv6-sig-network) | 16 | 0.54% | 1 | 0.19% | Improving |
| [pull-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-network) | 12 | 0.41% | 5 | 0.96% | Stable |
| [pull-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-network) | 9 | 0.31% | 1 | 0.19% | Stable |
| [periodic-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-network) | 6 | 0.20% | N/A | N/A | N/A |
| [periodic-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network) | 5 | 0.17% | 2 | 0.38% | Stable |
| [periodic-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network) | 4 | 0.14% | 1 | 0.19% | Stable |
| [pull-k8s-1.35-ipv6-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network) | 3 | 0.10% | 2 | 0.38% | Stable |
| [periodic-k8s-1.35-sig-network-with-dnc](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network-with-dnc) | 1 | 0.03% | 1 | 0.19% | Stable |

sig-network has low failure counts overall. No individual test stands out with a low success rate in the non-S390X lanes. The Networkpolicy and bridge nic-hotplug failures are primarily on S390X (see Platform-specific section).

---

## sig-operator

**Total failures (28d)**: ~104 (~3.5% of all failures)
**Total failures (7d)**: ~8 (~1.5%)

### Lanes

| Lane | 28d Failures | 28d Share | 7d Failures | 7d Share | Trend |
|------|-------------|-----------|-------------|----------|-------|
| [pull-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-operator) | 50 | 1.70% | 3 | 0.57% | Improving |
| [periodic-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-operator) | 46 | 1.56% | N/A | N/A | N/A |
| [periodic-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-operator) | 4 | 0.14% | N/A | N/A | N/A |
| [periodic-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-operator) | 2 | 0.07% | N/A | N/A | N/A |
| [pull-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-operator) | 1 | 0.03% | 1 | 0.19% | Stable |
| [pull-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-operator) | 1 | 0.03% | 1 | 0.19% | Stable |

No individual test failures reported from operator lanes — all failures are Pod-level infrastructure issues. Improving trend overall.

---

## sig-performance

**Total failures (28d)**: ~76 (~2.6% of all failures)
**Total failures (7d)**: N/A (no 7d data for performance lanes)

### Lanes

| Lane | 28d Failures | 28d Share | 7d Failures | 7d Share | Trend |
|------|-------------|-----------|-------------|----------|-------|
| [periodic-k8s-1.34-sig-performance-kwok-100](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok-100) | 68 | 2.31% | N/A | N/A | N/A |
| [periodic-k8s-1.34-sig-performance-kwok](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok) | 4 | 0.14% | N/A | N/A | N/A |
| [periodic-k8s-1.34-sig-performance](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance) | 2 | 0.07% | N/A | N/A | N/A |
| [pull-k8s-1.34-sig-performance](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-performance) | 2 | 0.07% | N/A | N/A | N/A |

The kwok-100 lane has no per-test failure data (68 total failures, no `failed_tests` entries), suggesting all failures are at the infrastructure/Pod level. Performance lanes show relatively low failure rates overall.

### Failing Tests (presubmit only)

| Test | 28d Rate | Lanes | Dispersion |
|------|----------|-------|------------|
| Density test - create batch of 100 running VMs with instancetype/preference | 94.0% | pull-1.34-sig-performance | 1 lane |
| Density test - create batch of 100 running VMs | 94.5% | pull-1.34-sig-performance | 1 lane |

Low failure rates, concentrated in one presubmit lane. Not actionable.

---

## Platform-specific

### S390X

**Total failures (28d)**: 301 (10.2%)
**Total failures (7d)**: 120 (22.9%)

Lane: [periodic-kubevirt-e2e-test-S390X](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-test-S390X)

| 28d Builds | 7d Builds | 28d Failures | 7d Failures | Trend |
|------------|-----------|-------------|-------------|-------|
| 112 | 28 | 301 | 120 | **Worsening** (7d failure share jumped from 10.2% to 22.9%) |

#### Failing Tests

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [should block the eviction api and migrate](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | 56.7% (13/30) | 57.1% (12/28) | Stable | Flaky |
| [MemBalloon period 0](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L160) | 88.9% (12/108) | 60.7% (11/28) | **Worsening** | Burst |
| [MemBalloon period 12](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L159) | 90.7% (10/108) | 67.9% (9/28) | **Worsening** | Burst |
| [MemBalloon Should be present in domain](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L92) | 90.8% (10/109) | 67.9% (9/28) | **Worsening** | Burst |
| [should migrate even if virtqemud has restarted](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L585) | 78.9% (23/109) | 92.9% (2/28) | **Improving** | Flaky |
| [should be successfully migrated multiple times](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L547) | 82.6% (19/109) | 71.4% (8/28) | **Worsening** | Flaky |
| [Networkpolicy allow ports 80 and 81](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L242) | 85.3% (16/109) | 92.9% (2/28) | Improving | Flaky |
| [should migrate with LiveMigrateIfPossible](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L342) | 87.2% (14/109) | 92.9% (2/28) | Improving | Flaky |
| [CPU Hotplug - should successfully plug vCPUs](https://github.com/kubevirt/kubevirt/blob/main/tests/hotplug/cpu.go#L106) | N/A | 89.3% (3/28) | New in 7d | Burst |
| [Configurations - resolved QEMU machine type](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_configuration_test.go#L877) | N/A | 89.3% (3/28) | New in 7d | Burst |
| [Configurations - machine type from kubevirt-config](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_configuration_test.go#L889) | N/A | 89.3% (3/28) | New in 7d | Burst |

**Analysis**: S390X is showing significant instability. The MemBalloon defaults tests are **actively worsening** (88-91% -> 61-68%), suggesting a recent regression on S390X. The burst pattern (failures clustered in recent builds) and the emergence of new failures (CPU Hotplug, Configurations) in the 7d window point to a platform-level regression, not individual test flakes.

The Networkpolicy tests on S390X (multiple tests at 90-95% success rate) appear to be **infra-correlated** — several unrelated network tests all fail at similar rates, suggesting shared infrastructure instability during certain S390X builds.

### ARM64

Lane: [pull-kubevirt-e2e-kind-1.35-sig-compute-arm64](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-sig-compute-arm64)

| 28d Failures | 7d Failures | Trend |
|-------------|-------------|-------|
| 31 | 6 | Stable |

All failures are Pod-level (84.8% success 28d, 87.9% success 7d). No individual test failures. Infrastructure only.

---

## Quarantine Candidate Summary

Non-quarantined tests that should be considered for quarantine, ranked by priority:

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority |
|------|-----|-------|-------------------|----------|---------|------------|----------|
| [virtiofs ServiceAccount token](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L232) | compute | periodic-1.36, pull-1.36 | 0% | Stable (0%) | Consecutive | 2 lanes | **High** |
| [volume migrate block to filesystem](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1275) | storage | periodic-1.36, pull-1.36 | 64.8% | Stable | Flaky | 2 lanes | **High** |
| [volume migrate filesystem to filesystem](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1273) | storage | periodic-1.36, pull-1.36 | 65.9% | Stable | Flaky | 2 lanes | **High** |
| [HostDevices PCI passthrough (1 device)](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | compute | periodic-1.34 | 76.2% -> 3.8% | Worsening | Burst | 1 lane | **High** |
| [HostDevices PCI passthrough (2 devices)](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | compute | periodic-1.34, pull-1.36 | 76.2% -> 3.8% | Worsening | Burst | 2 lanes | **High** |
| [volume migrate containerdisk + hotplugged](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1132) | storage | periodic-1.36, pull-1.36 | 82.1% | Worsening (70.4%) | Flaky | 2 lanes | **Medium** |
| [VSOCK retain CID for migration target](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L121) | compute | periodic-1.36 | 65.4% | N/A | Infra-correlated | 1 lane | **Medium** |
| [VSOCK TLS both sides](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L223) | compute | periodic-1.36 | 72.1% | N/A | Infra-correlated | 1 lane | **Medium** |
| [VSOCK no app listeners](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L240) | compute | periodic-1.36 | 73.1% | N/A | Infra-correlated | 1 lane | **Low** |
| [VSOCK without TLS](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L224) | compute | periodic-1.36 | 76.0% | N/A | Infra-correlated | 1 lane | **Low** |
| [VSOCK virtio non-transitional](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L95) | compute | periodic-1.36 | 76.9% | N/A | Infra-correlated | 1 lane | **Low** |
| [VSOCK virtio transitional](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L94) | compute | periodic-1.36 | 77.9% | N/A | Infra-correlated | 1 lane | **Low** |
| [VSOCK invalid port](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L227) | compute | periodic-1.36 | 83.7% | N/A | Infra-correlated | 1 lane | **Low** |
| [VM deletion transition time](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L179) | monitoring | periodic-1.34 | 83.0% | 85.7% | Flaky | 1 lane | **Low** |
| [S390X MemBalloon period 0](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L160) | compute (S390X) | periodic-S390X | 88.9% -> 60.7% | Worsening | Burst | 1 lane | **Medium** |
| [S390X eviction api and migrate](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | compute (S390X) | periodic-S390X | 56.7% | Stable | Flaky | 1 lane | **Medium** |

**Key recommendations**:
1. **virtiofs ServiceAccount token**: Not a flake — this is a 100% deterministic failure on k8s-1.36 only. Needs a code fix, not quarantine. Investigate what changed in k8s-1.36 that breaks virtiofs ServiceAccount mounting.
2. **Volume hotplug migration tests**: Genuine flakes concentrated in k8s-1.36. Strong quarantine candidates.
3. **HostDevices PCI passthrough**: Recent regression on k8s-1.34 (dropped from ~76% to ~4%). Investigate before quarantining — likely a code or infra change in the last week.
4. **VSOCK tests**: All concentrated in one lane. Investigate the k8s-1.36 periodic VSOCK infrastructure before quarantining individual tests.
5. **S390X MemBalloon**: Actively worsening. Investigate as a platform regression.

---

## Quarantine Debt

Already-quarantined tests that remain severely broken (success rate <50%):

| Test | SIG | Worst 28d Rate | 7d Rate | Status |
|------|-----|----------------|---------|--------|
| [Prometheus VMI namespace metric](https://github.com/kubevirt/kubevirt/blob/main/tests/infrastructure/prometheus.go#L95) | monitoring | 0% | 0% | 100% failure for 28+ days |
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | storage | 37.1% | 40.7% | Stable at ~40% success |
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | compute | 40.0% | 0% (k8s-1.34) | Worsening |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | compute | 40.0% | 0% (k8s-1.34) | Worsening |

These tests have been quarantined but never fixed. The Prometheus metric test is a 100% failure for the entire analysis window. The USB passthrough tests are getting worse, not better. These represent accumulated tech debt that should be prioritized for investigation or permanent removal.
