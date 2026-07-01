# Flake Overview Report — 2026-06-22

## Summary

| Metric | 28-day | 7-day | Trend |
|--------|--------|-------|-------|
| Total failures | 1903 | 598 | Worsening (7d projects to ~2392/28d) |
| Lanes with failures | 47 | 32 | — |
| Analysis window | May 25 – Jun 22 | Jun 15 – Jun 22 | — |

**Overall trend: Worsening.** The 7-day failure count (598) projects to ~2392 over 28 days, exceeding the actual 28-day total (1903) by ~26%. The recent week is disproportionately noisy.

**Top 3 worst SIGs by failure count (28d):**
1. **sig-compute** — 1148 failures (60.3%) across periodic + presubmit lanes
2. **sig-storage** — 303 failures (15.9%)
3. **sig-network** — 149 failures (7.8%)

**Quarantine candidates identified:** 3 high priority, 3 medium priority, 1 low priority

**Infra-unstable lanes:** Pod-level failures affect nearly every lane, with presubmit lanes particularly bad (52–77% Pod success rates)

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

Pod-level failures represent infrastructure instability (cluster provisioning, setup failures) rather than test code issues. These should **not** be addressed by quarantining individual tests.

### Presubmit Pod Failure Rates (28d)

| Lane | Pod Success Rate (28d) | Pod Success Rate (7d) | Trend |
|------|------------------------|-----------------------|-------|
| pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations | 52.6% | 76.0% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-network | 63.6% | 72.3% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-storage | 64.5% | 75.7% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-compute | 65.9% | 76.4% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-operator | 65.3% | 73.9% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-compute-serial | 65.4% | 74.6% | Improving |
| pull-kubevirt-e2e-kind-1.35-sig-compute-arm64 | 71.6% | 85.0% | Improving |
| pull-kubevirt-e2e-k8s-1.35-sig-storage | 72.0% | 93.2% | Improving |
| pull-kubevirt-e2e-k8s-1.35-sig-operator | 73.3% | n/a | — |
| pull-kubevirt-e2e-k8s-1.34-sig-network | 73.5% | 91.3% | Improving |
| pull-kubevirt-e2e-k8s-1.34-sig-compute | 73.7% | n/a | — |
| pull-kubevirt-e2e-k8s-1.35-sig-network | 74.1% | n/a | — |
| pull-kubevirt-e2e-kind-sriov | 74.5% | n/a | — |
| pull-kubevirt-e2e-k8s-1.34-sig-storage | 74.9% | 94.9% | Improving |
| pull-kubevirt-e2e-k8s-1.34-sig-operator | 75.4% | n/a | — |
| pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network | 76.0% | n/a | — |
| pull-kubevirt-e2e-k8s-1.34-windows2016 | 77.5% | n/a | — |
| pull-kubevirt-check-tests-for-flakes | 78.1% | 85.2% | Improving |
| pull-kubevirt-e2e-kind-1.35-vgpu | 78.4% | n/a | — |

**Key observation:** Pod failure rates in presubmit lanes are **improving** in the 7d window vs 28d, particularly for k8s-1.36 lanes and 1.35-sig-storage. However, k8s-1.36 lanes still have the worst rates overall (65–76% success).

### Periodic Pod Failure Rates (28d)

| Lane | Pod Success Rate (28d) | Pod Success Rate (7d) | Trend |
|------|------------------------|-----------------------|-------|
| periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations | 64.7% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.36-sig-network-with-dnc | 81.3% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.36-ipv6-sig-network | 80.6% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.36-sig-storage | 85.0% | 87.5% | Stable |
| periodic-kubevirt-e2e-k8s-1.34-sig-compute | 86.6% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.34-sig-performance | 86.7% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.34-sig-storage | 86.9% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.35-sig-network | 86.9% | 93.8% | Improving |
| periodic-kubevirt-e2e-k8s-1.34-sig-network | 86.9% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.35-sig-compute | 87.5% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok | 87.5% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.35-sig-operator | 87.8% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.36-sig-compute | 87.9% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.36-sig-network | 87.9% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.36-sig-operator | 87.9% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.35-sig-storage | 87.9% | n/a | — |
| periodic-kubevirt-e2e-test-S390X | 88.0% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.34-sig-monitoring | 88.0% | n/a | — |
| periodic-kubevirt-e2e-kind-1.34-sev | 88.0% | n/a | — |
| periodic-kubevirt-e2e-k8s-1.34-sig-operator | 88.7% | n/a | — |

### AfterSuite Failures

AfterSuite failures indicate test framework teardown issues:

| Lane | AfterSuite Success Rate (28d) | 7d |
|------|-------------------------------|----|
| pull-kubevirt-e2e-k8s-1.35-sig-operator | 0% (2/2 runs) | 0% (2/2) |
| pull-kubevirt-e2e-k8s-1.35-sig-storage | 0% (2/2 runs) | 0% (2/2) |
| pull-kubevirt-e2e-k8s-1.34-sig-operator | 0% (1/1 runs) | n/a |

These are very low-sample entries and likely represent individual bad builds rather than persistent issues.

---

## sig-compute

**28d total failures: 1148** (60.3% of all failures)
**7d total failures: 384** (64.2% of all failures)
**Trend: Worsening** — sig-compute's share of failures is growing

### Periodic Lanes

#### [periodic-kubevirt-e2e-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-compute)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 421 | 22.12% | 97 | 75 | 12.54% | Stable |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [Should successfully passthrough 2 emulated PCI devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L112) | 0% | 0% | Stable | Concentrated (1 lane) | Consecutive — deterministic breakage |
| [Should successfully passthrough an emulated PCI device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L111) | 0% | 0% | Stable | Concentrated (1 lane + serial) | Consecutive — deterministic breakage |
| [USB] [QUARANTINE] Should successfully passthrough 1 emulated USB device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) `[QUARANTINE]` | 0% | 0% | Stable | Dispersed (3 lanes) | Consecutive — deterministic breakage |
| [[QUARANTINE] Should successfully passthrough 2 emulated USB devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) `[QUARANTINE]` | 0% | 0% | Stable | Dispersed (3 lanes) | Consecutive — deterministic breakage |
| [[QUARANTINE] Guest console log flood serial console](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) `[QUARANTINE]` | 85.3% | 92.9% | Improving | Dispersed (3 lanes) | Flaky |
| [[QUARANTINE] ClusterProfiler subresource access](https://github.com/kubevirt/kubevirt/blob/main/tests/infrastructure/cluster-profiler.go#L62) `[QUARANTINE]` | n/a (28d) | 92.9% | — | Concentrated (2 lanes) | Flaky |

**Analysis:** The emulated PCI device tests (0% success, k8s-1.34 only) are **deterministic breakage** on this lane, not flakes. The k8s-1.34 environment likely lacks the required device emulation support. USB passthrough tests are similarly broken at 0% here but show 52–58% on k8s-1.35/1.36, indicating a genuine flake on newer versions but complete breakage on k8s-1.34.

#### [periodic-kubevirt-e2e-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 122 | 6.41% | 97 | 47 | 7.86% | Worsening |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [[QUARANTINE] USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) `[QUARANTINE]` | 52.8% | 15.4% | Worsening | Dispersed (3 lanes) | Flaky |
| [[QUARANTINE] USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) `[QUARANTINE]` | 52.8% | 15.4% | Worsening | Dispersed (3 lanes) | Flaky |
| [[QUARANTINE] Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) `[QUARANTINE]` | 87.6% | 84.6% | Stable | Dispersed (3 lanes) | Flaky |

#### [periodic-kubevirt-e2e-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 159 | 8.36% | 99 | 34 | 5.69% | Improving |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [[QUARANTINE] USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) `[QUARANTINE]` | 58.3% | 56.3% | Stable | Dispersed (3 lanes) | Flaky |
| [[QUARANTINE] USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) `[QUARANTINE]` | 58.3% | 56.3% | Stable | Dispersed (3 lanes) | Flaky |
| [[QUARANTINE] Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) `[QUARANTINE]` | 85.4% | 81.3% | Stable | Dispersed (3 lanes) | Flaky |
| [[QUARANTINE] ClusterProfiler subresource access](https://github.com/kubevirt/kubevirt/blob/main/tests/infrastructure/cluster-profiler.go#L62) `[QUARANTINE]` | 92.7% | 93.8% | Stable | Concentrated (2 lanes) | Flaky |

#### [periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 53 | 2.79% | 35 | 48 | 8.03% | Worsening significantly |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [[QUARANTINE] Cancel migration — delete source](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L650) `[QUARANTINE]` | 70.0% | 30.8% | Worsening | Dispersed (4+ lanes) | Flaky — high flip rate |
| [[QUARANTINE] Cancel migration — delete target](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L651) `[QUARANTINE]` | 70.0% | 30.8% | Worsening | Dispersed (4+ lanes) | Flaky — high flip rate |
| [Post-copy migration with liveness probe](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/postcopy.go#L176) | 88.5% | 92.3% | Stable | Concentrated (1 lane) | Flaky |
| [Evacuation backoff should be cleared](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/evacuation.go#L185) | 93.3% | n/a | — | Concentrated (1 lane) | Flaky |

### Presubmit Lanes (sig-compute)

| Lane | Failures (28d) | Pod Rate (28d) | Pod Rate (7d) | Trend |
|------|----------------|----------------|---------------|-------|
| [pull-kubevirt-e2e-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute) | 84 | 65.9% | 76.4% | Improving |
| [pull-kubevirt-e2e-k8s-1.36-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute-serial) | 7 | 65.4% | 74.6% | Improving |
| [pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations) | 29 | 52.6% | 76.0% | Improving |
| [pull-kubevirt-e2e-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute) | 13 | 74.0% | n/a | — |
| [pull-kubevirt-e2e-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-compute) | 9 | 73.7% | n/a | — |

Presubmit sig-compute failures are almost entirely Pod-level (infra). Actual test failures in presubmit lanes are minimal.

---

## sig-storage

**28d total failures: 303** (15.9% of all failures)
**7d total failures: 142** (23.7% of all failures)
**Trend: Worsening significantly** — sig-storage's share nearly doubled in the last 7 days

### Periodic Lanes

#### [periodic-kubevirt-e2e-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-storage)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 96 | 5.04% | 99 | 48 | 8.03% | Worsening |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [[QUARANTINE] Should pause VMI on IO error](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) `[QUARANTINE]` | 48.9% | 87.5% | Improving | Dispersed (3 lanes) | Flaky |
| [[QUARANTINE] Cancel migration — delete source](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L650) `[QUARANTINE]` | 87.2% | 25.0% | Worsening | Dispersed (4+ lanes) | Flaky — recent regression |
| [[QUARANTINE] Cancel migration — delete target](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L651) `[QUARANTINE]` | 87.2% | 25.0% | Worsening | Dispersed (4+ lanes) | Flaky — recent regression |

#### [periodic-kubevirt-e2e-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-storage)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 85 | 4.47% | 100 | 44 | 7.36% | Worsening |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [[QUARANTINE] Should pause VMI on IO error](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) `[QUARANTINE]` | 55.2% | 93.8% | Improving | Dispersed (3 lanes) | Flaky |
| [[QUARANTINE] Cancel migration — delete source](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L650) `[QUARANTINE]` | 87.5% | 31.3% | Worsening | Dispersed (4+ lanes) | Flaky — recent regression |
| [[QUARANTINE] Cancel migration — delete target](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L651) `[QUARANTINE]` | 88.5% | 31.3% | Worsening | Dispersed (4+ lanes) | Flaky — recent regression |

#### [periodic-kubevirt-e2e-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 92 | 4.83% | 100 | 43 | 7.19% | Worsening |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [[QUARANTINE] Should pause VMI on IO error](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) `[QUARANTINE]` | 52.1% | 92.9% | Improving | Dispersed (3 lanes) | Flaky |
| [[QUARANTINE] Cancel migration — delete source](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L650) `[QUARANTINE]` | 89.4% | 28.6% | Worsening | Dispersed (4+ lanes) | Flaky — recent regression |
| [[QUARANTINE] Cancel migration — delete target](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L651) `[QUARANTINE]` | 89.4% | 28.6% | Worsening | Dispersed (4+ lanes) | Flaky — recent regression |

### Presubmit Lanes (sig-storage)

| Lane | Failures (28d) | Pod Rate (28d) | Pod Rate (7d) | Trend |
|------|----------------|----------------|---------------|-------|
| [pull-kubevirt-e2e-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-storage) | 18 | 64.5% | 75.7% | Improving |
| [pull-kubevirt-e2e-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-storage) | 7 | 74.9% | 94.9% | Improving |
| [pull-kubevirt-e2e-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-storage) | 5 | 72.0% | 93.2% | Improving |

---

## sig-network

**28d total failures: 149** (7.8% of all failures)
**7d total failures: 37** (6.2% of all failures)
**Trend: Improving** — sig-network's share is decreasing

### Periodic Lanes

#### [periodic-kubevirt-e2e-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 24 | 1.26% | 100 | 5 | 0.84% | Improving |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [Passt migration connectivity ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | 82.4% | 81.8% | Stable | Dispersed (3+ lanes) | Flaky |
| [NIC hotplug migrate VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L145) | n/a | 93.3% | — | Concentrated (1 lane) | Flaky |

#### [periodic-kubevirt-e2e-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 21 | 1.10% | 99 | 8 | 1.34% | Stable |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [[QUARANTINE] Passt migration connectivity ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) `[QUARANTINE]` | 80.0% | 80.0% | Stable | Dispersed (3+ lanes) | Flaky |
| [Passt migration connectivity ipv4 (non-quarantine)](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | 92.4% | 90.9% | Stable | Dispersed (3+ lanes) | Flaky |

#### [periodic-kubevirt-e2e-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-network)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 8 | 0.42% | 99 | 3 | 0.50% | Stable |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [[QUARANTINE] Passt migration connectivity ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) `[QUARANTINE]` | 80.0% | 80.0% | Stable | Dispersed (3+ lanes) | Flaky |
| [Passt migration connectivity ipv4 (non-quarantine)](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | 92.4% | 90.9% | Stable | Dispersed (3+ lanes) | Flaky |

#### [periodic-kubevirt-e2e-k8s-1.36-sig-network-with-dnc](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network-with-dnc)

| Failures (28d) | Share | Builds |
|-----------------|-------|--------|
| 3 | 0.16% | 16 |

| Test | Success Rate (28d) | Dispersion | Pattern |
|------|---------------------|------------|---------|
| [NIC hotplug multiple interfaces](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L145) | 87.5% | Concentrated (1 lane) | Flaky |
| [Passt migration connectivity ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | 93.3% | Dispersed (3+ lanes) | Flaky |

### Presubmit Lanes (sig-network)

| Lane | Failures (28d) | Pod Rate (28d) | Pod Rate (7d) | Trend |
|------|----------------|----------------|---------------|-------|
| [pull-kubevirt-e2e-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-network) | 23 | 63.6% | 72.3% | Improving |
| [pull-kubevirt-e2e-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-network) | 32 | 74.1% | n/a | — |
| [pull-kubevirt-e2e-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-network) | 27 | 73.5% | 91.3% | Improving |
| [pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network) | 10 | 76.0% | n/a | — |

---

## sig-operator

**28d total failures: 161** (8.5% of all failures)
**7d total failures: 91** (15.2% of all failures)
**Trend: Worsening** — sig-operator's share nearly doubled

### Periodic Lanes

#### [periodic-kubevirt-e2e-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-operator)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 45 | 2.36% | 97 | 44 | 7.36% | Worsening significantly |

No failing tests reported (28d only has Pod failures). The 7d spike to 44 failures with only 14 builds and no test details suggests a cluster-level outage or infra issue during this period.

#### [periodic-kubevirt-e2e-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-operator)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 3 | 0.16% | 98 | 1 | 0.17% | Stable |

#### [periodic-kubevirt-e2e-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-operator)

| Failures (28d) | Share | Builds | 7d Failures |
|-----------------|-------|--------|-------------|
| 2 | 0.11% | 99 | n/a |

### Presubmit Lanes (sig-operator)

| Lane | Failures (28d) | Pod Rate (28d) | Pod Rate (7d) | Trend |
|------|----------------|----------------|---------------|-------|
| [pull-kubevirt-e2e-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-operator) | 60 | 73.3% | n/a | — |
| [pull-kubevirt-e2e-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-operator) | 50 | 65.3% | 73.9% | Improving |
| [pull-kubevirt-e2e-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-operator) | 1 | 75.4% | n/a | — |

Notable presubmit test:
| Test | Success Rate (7d) | Lane |
|------|-------------------|------|
| [Should be able to delete and re-create kubevirt install](https://github.com/kubevirt/kubevirt/blob/main/tests/operator/operator.go#L1032) | 95.0% | pull-kubevirt-e2e-k8s-1.36-sig-operator |

---

## sig-monitoring

**28d total failures: 54** (2.8% of all failures)
**7d total failures: 24** (4.0% of all failures)
**Trend: Worsening** — sig-monitoring share increased

### [periodic-kubevirt-e2e-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-monitoring)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 53 | 2.79% | 50 | 24 | 4.01% | Worsening |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [Should contain virt components metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L119) | 75.0% | 0% | Worsening | Concentrated (1 periodic + 1 presubmit) | Recent regression? |
| [VirtOperatorRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L340) | 75.0% | 28.6% | Worsening | Concentrated (1 periodic + 1 presubmit) | Recent regression? |
| [[QUARANTINE] VM dirty rate metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/vm_monitoring.go#L334) `[QUARANTINE]` | 89.6% | n/a | — | Concentrated (1 lane) | Flaky |

### [pull-kubevirt-e2e-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-monitoring)

| Test | Success Rate (28d) | Severity |
|------|---------------------|----------|
| [Should contain virt components metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L119) | 30.8% | likely-flaky |
| [VirtOperatorRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L340) | 28.1% | likely-flaky |
| [Should have all required annotations](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L119) | 87.2% | inconclusive |

**Analysis:** The monitoring tests "virt components metrics" and "VirtOperatorRESTErrorsBurst" are severely broken in the presubmit lane (28–31% success) and getting worse in periodics (7d: 0% and 28.6%). This looks like a recent regression rather than long-standing flakiness.

---

## sig-performance

**28d total failures: 8** (0.4% of all failures)
**7d total failures: 0**
**Trend: Stable/Minimal**

### [periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok)

| Failures (28d) | Share | Builds |
|-----------------|-------|--------|
| 5 | 0.26% | 72 |

Pod-level failures only (87.5% success rate). No test-level issues.

### [periodic-kubevirt-e2e-k8s-1.34-sig-performance](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance)

| Failures (28d) | Share | Builds |
|-----------------|-------|--------|
| 3 | 0.16% | 75 |

Pod-level failures only (86.7% success rate). No test-level issues.

---

## Platform-specific

### [periodic-kubevirt-e2e-test-S390X](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-test-S390X)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 209 | 10.98% | 100 | 39 | 6.52% | Improving |

| Test | Success Rate (28d) | Success Rate (7d) | Trend | Dispersion | Pattern |
|------|--------------------|--------------------|-------|------------|---------|
| [Should block eviction api and migrate](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | 45.9% | 37.5% | Worsening | Concentrated (S390X only) | Flaky — high flip rate |
| [Migrate multiple times with cloud-init](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L551) | 80.6% | 75.0% | Worsening | Concentrated (S390X only) | Flaky |
| [Should migrate even if virtqemud restarted](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L589) | 92.9% | 93.8% | Stable | Concentrated (S390X only) | Flaky |
| [CPU hotplug — plug vCPUs](https://github.com/kubevirt/kubevirt/blob/main/tests/hotplug/cpu.go#L106) | n/a | 87.5% | — | Concentrated (S390X only) | Flaky |

**Analysis:** S390X migration tests are the worst non-quarantined flakes by success rate. The eviction test at 37.5% (7d) is a strong quarantine candidate. These are S390X-specific and don't appear in other lanes.

### [pull-kubevirt-e2e-kind-1.35-sig-compute-arm64](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-sig-compute-arm64)

| Failures (28d) | Share | Builds | 7d Failures | 7d Share | Trend |
|-----------------|-------|--------|-------------|----------|-------|
| 40 | 2.10% | 540 | 18 | 3.01% | Worsening |

Pod-level failures only (71.6% → 85.0%, improving in 7d).

### [periodic-kubevirt-e2e-kind-1.34-sev](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-kind-1.34-sev)

| Failures (28d) | Share | Builds |
|-----------------|-------|--------|
| 4 | 0.21% | 75 |

| Test | Success Rate (28d) | Success Rate (7d) | Pattern |
|------|--------------------|--------------------|---------|
| [[QUARANTINE] SEV guest attestation](https://github.com/kubevirt/kubevirt/blob/main/tests/launchsecurity/sev.go#L349) `[QUARANTINE]` | 94.6% | 91.7% | Flaky |

### [pull-kubevirt-e2e-kind-1.35-vgpu](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-vgpu)

| Failures (28d) | Share | Builds |
|-----------------|-------|--------|
| 2 | 0.11% | 191 |

| Test | Success Rate (7d) | Pattern |
|------|-------------------|---------|
| [Passthrough mediated device](https://github.com/kubevirt/kubevirt/blob/main/tests/mdev_configuration_allocation_test.go#L269) | 94.0% | Flaky |

### [pull-kubevirt-e2e-kind-sriov](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-sriov)

| Failures (28d) | Share | Builds |
|-----------------|-------|--------|
| 9 | 0.47% | 185 |

Pod-level failures only (74.5% success rate).

---

## Quarantine Candidate Summary

### Non-quarantined tests that should be quarantined

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority |
|------|-----|-------|-------------------|----------|---------|------------|----------|
| [Should successfully passthrough 2 emulated PCI devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L112) | compute | 1 periodic | 0% | Stable (0%) | Consecutive | Concentrated | **High** |
| [Should successfully passthrough an emulated PCI device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L111) | compute | 1 periodic + serial | 0% | Stable (0%) | Consecutive | Concentrated | **High** |
| [Should block eviction api and migrate (S390X)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | compute | 1 (S390X) | 45.9% | Worsening (37.5%) | Flaky | Concentrated | **High** |
| [Should contain virt components metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L119) | monitoring | 1 periodic + 1 presubmit | 30.8% (presubmit) | Worsening (0% periodic 7d) | Regression | Concentrated | **Medium** |
| [VirtOperatorRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L340) | monitoring | 1 periodic + 1 presubmit | 28.1% (presubmit) | Worsening (28.6% periodic 7d) | Regression | Concentrated | **Medium** |
| [Passt migration connectivity ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | network | 3+ lanes | 82.4% | Stable | Flaky | Dispersed | **Medium** |
| [Migrate multiple times cloud-init (S390X)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L551) | compute | 1 (S390X) | 80.6% | Worsening (75.0%) | Flaky | Concentrated | **Low** |

### Quarantine debt — already quarantined but still severely broken

These tests are quarantined but remain unfixed with very low success rates:

| Test | SIG | Worst Rate (28d) | 7d Trend | Lanes |
|------|-----|-------------------|----------|-------|
| [[QUARANTINE] USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) | compute | 0% (k8s-1.34), 52.8% (k8s-1.35) | Worsening (15.4% 7d on 1.35) | 3 periodic |
| [[QUARANTINE] USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) | compute | 0% (k8s-1.34), 52.8% (k8s-1.35) | Worsening (15.4% 7d on 1.35) | 3 periodic |
| [[QUARANTINE] Cancel migration — delete source](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L650) | compute/storage | 70.0% (migrations) | Worsening (25–31% 7d) | 4+ lanes |
| [[QUARANTINE] Cancel migration — delete target](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L651) | compute/storage | 70.0% (migrations) | Worsening (25–31% 7d) | 4+ lanes |
| [[QUARANTINE] Should pause VMI on IO error](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | storage | 48.9% (k8s-1.34) | Improving (87.5–93.8% 7d) | 3 periodic |
| [[QUARANTINE] Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) | compute | 85.3% | Improving (81–93% 7d) | 3 periodic |

### Key findings

1. **Cancel migration tests are regressing rapidly.** The `delete source/target migration` quarantined tests dropped from ~70–89% (28d) to 25–31% (7d) across 4+ lanes. This is the most urgent issue — something changed recently to make these much worse.

2. **Monitoring tests may have a recent regression.** "Virt components metrics" and "VirtOperatorRESTErrorsBurst" are at 0–28% success in the 7d window for periodics. Investigate recent changes to monitoring infrastructure.

3. **PCI device passthrough is deterministically broken on k8s-1.34.** The two emulated PCI device tests have 0% success across all 28 days on the 1.34 lane only. This is not a flake — it's a broken test/environment combination that should be quarantined or the lane configuration fixed.

4. **USB passthrough flakiness is worsening.** Already quarantined but dropping from ~53% to ~15% on k8s-1.35 in the last 7 days.

5. **S390X migration tests are platform-specific flakes.** The eviction test at 37.5% (7d) is the worst non-quarantined flake by rate and should be quarantined.

6. **Infra stability is improving.** Pod failure rates across presubmit lanes are generally better in the 7d window, especially for k8s-1.36 lanes. The k8s-1.36 compute-migrations lane improved dramatically (52.6% → 76.0%).

7. **IO error test is recovering.** The quarantined "should pause VMI on IO error" test improved from ~49–55% (28d) to 87.5–93.8% (7d) — a potential fix may have landed.
