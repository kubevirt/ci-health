# KubeVirt Flake Overview - 2026-06-24

## Summary

| Metric | 28-day | 7-day | Trend |
|--------|--------|-------|-------|
| Analysis window | May 27 - Jun 24 | Jun 17 - Jun 24 | |
| Total failures | 1,895 | 561 | |
| Lanes with failures | 47 | 34 | |
| Daily failure rate | ~67.7/day | ~80.1/day | **Worsening (+18%)** |

The daily failure rate has increased 18% in the most recent 7-day window compared to the 28-day baseline. The primary drivers are:

- **Cancel migration tests** (quarantined) have regressed from ~72% success to **0%** across all storage and compute-migrations lanes — this is a complete breakage, not flakiness
- **sig-monitoring metrics tests** (NOT quarantined) have worsened from ~59% to **0%** in the periodic lane — highest-priority quarantine candidate
- **USB passthrough** and **HostDevices PCI** (sig-compute) remain consistently broken on k8s-1.34, though PCI devices are improving significantly (17% -> 72%)

**Top 3 worst SIGs by total failure count (28d):**
1. **sig-compute**: 1,128 failures (59.5%) — dominated by USB/PCI passthrough + migrations
2. **sig-storage**: 318 failures (16.8%) — driven by cancel migration + error disk
3. **Platform-specific (S390X)**: 174 failures (9.2%) — eviction migration flake

**Quarantine candidates:** 2 high priority, 1 medium priority
**Quarantine debt (broken despite quarantine):** 4 tests at 0% or near-0% success
**Infra-unstable lanes:** Pod-level failure rates range 60-82%; k8s-1.36 presubmit lanes are the worst

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

Pod-level failures represent infrastructure instability — test environment setup failures, cluster provisioning issues, or resource exhaustion. These are not test-level problems and should not be addressed by quarantining individual tests.

| Lane | 28d Success Rate | 7d Success Rate | Trend |
|------|------------------|-----------------|-------|
| pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations | 60.1% | 79.4% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-network | 64.5% | 76.7% | Improving |
| pull-kubevirt-e2e-k8s-1.34-sig-monitoring | 65.4% | 82.1% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-compute-serial | 65.7% | 77.0% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-operator | 66.1% | 78.7% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-storage | 66.3% | 80.1% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-compute | 67.5% | 78.9% | Improving |
| pull-kubevirt-e2e-kind-1.35-sig-compute-arm64 | 73.3% | 84.4% | Improving |
| pull-kubevirt-e2e-k8s-1.35-sig-storage | 73.9% | — | — |
| pull-kubevirt-e2e-k8s-1.34-sig-compute | 74.5% | — | — |
| pull-kubevirt-e2e-k8s-1.35-sig-operator | 74.4% | 94.7% | Improving |
| pull-kubevirt-e2e-k8s-1.34-sig-network | 74.6% | — | — |
| pull-kubevirt-e2e-k8s-1.35-sig-network | 75.5% | — | — |
| pull-kubevirt-e2e-k8s-1.34-sig-operator | 76.6% | — | — |
| pull-kubevirt-e2e-k8s-1.35-sig-compute | 76.5% | — | — |
| pull-kubevirt-e2e-kind-sriov | 77.1% | — | — |
| pull-kubevirt-e2e-k8s-1.34-sig-storage | 77.4% | — | — |
| pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network | 77.7% | — | — |
| pull-kubevirt-check-tests-for-flakes | 78.3% | 85.7% | Improving |
| pull-kubevirt-e2e-k8s-1.34-windows2016 | 80.3% | — | — |
| pull-kubevirt-e2e-kind-1.35-vgpu | 81.7% | — | — |

**Observation:** k8s-1.36 presubmit lanes have the worst Pod-level success rates over 28 days but show a strong improving trend in the 7d window. Infrastructure stability appears to be recovering. The k8s-1.36 lanes still have the lowest absolute rates even in 7d.

### Periodic Pod-Level Failure Rates

| Lane | 28d Success Rate | 7d Note |
|------|------------------|---------|
| periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations | 77.8% | Not in 7d top |
| periodic-kubevirt-e2e-k8s-1.36-ipv6-sig-network | 85.1% | — |
| periodic-kubevirt-e2e-k8s-1.36-sig-network-with-dnc | 85.7% | — |
| periodic-kubevirt-e2e-k8s-1.36-sig-storage | 86.4% | — |
| periodic-kubevirt-e2e-k8s-1.34-sig-compute | 88.0% | — |
| periodic-kubevirt-e2e-k8s-1.34-sig-storage | 88.0% | — |
| periodic-kubevirt-e2e-k8s-1.35-sig-network | 88.1% | — |
| periodic-kubevirt-e2e-k8s-1.36-sig-compute | 88.8% | — |
| periodic-kubevirt-e2e-k8s-1.35-sig-compute | 88.8% | — |
| periodic-kubevirt-e2e-k8s-1.35-sig-storage | 88.9% | — |
| periodic-kubevirt-e2e-k8s-1.36-sig-network | 89.0% | — |
| periodic-kubevirt-e2e-k8s-1.35-sig-operator | 88.8% | — |
| periodic-kubevirt-e2e-k8s-1.36-sig-operator | 88.9% | — |
| periodic-kubevirt-e2e-k8s-1.34-sig-monitoring | 89.3% | — |
| periodic-kubevirt-e2e-test-S390X | 89.3% | — |
| periodic-kubevirt-e2e-k8s-1.34-sig-operator | 89.9% | — |

### BeforeSuite / Complete Setup Failures

No lanes show complete BeforeSuite/AfterSuite failures at 0% success rate. All setup failures are captured within the Pod-level metrics above.

---

## sig-compute

**Total 28d failures:** 1,128 (59.5% of all failures)
**Lanes:** 15 (periodic + presubmit across k8s-1.34/1.35/1.36 + arm64 + S390X)

### Periodic Lanes

| Lane | 28d Failures | Share | Builds | 7d Failures | Trend |
|------|-------------|-------|--------|-------------|-------|
| [periodic-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-compute) | 401 | 21.2% | 108 | 57 | Stable |
| [periodic-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute) | 155 | 8.2% | 108 | 35 | Stable |
| [periodic-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute) | 119 | 6.3% | 108 | 41 | Stable |
| [periodic-k8s-1.36-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations) | 70 | 3.7% | 55 | 58 | **Worsening** |

### Presubmit Lanes

| Lane | 28d Failures | Share | Builds | 7d Failures | Trend |
|------|-------------|-------|--------|-------------|-------|
| [pull-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute) | 81 | 4.3% | 519 | 10 | Improving |
| [pull-kind-1.35-sig-compute-arm64](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-sig-compute-arm64) | 45 | 2.4% | 598 | 15 | Stable |
| [pull-k8s-1.36-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations) | 33 | 1.7% | 636 | 19 | Worsening |
| [pull-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute) | 14 | 0.7% | 206 | 1 | Improving |
| [pull-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-compute) | 9 | 0.5% | 198 | 1 | Improving |
| [pull-k8s-1.36-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute-serial) | 7 | 0.4% | 491 | 2 | Stable |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.34-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) `[QUARANTINE]` | 9.5% | 40% | Improving | 3 lanes (1.34/1.35/1.36) | Flaky |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) `[QUARANTINE]` | 9.5% | 40% | Improving | 3 lanes (1.34/1.35/1.36) | Flaky |
| [PCI passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L112) | 17.1% | 72% | **Improving** | 1 lane (1.34 only) | Burst → recovering |
| [PCI passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L111) | 17.1% | 72% | **Improving** | 1 lane (1.34 only) | Burst → recovering |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) `[QUARANTINE]` | 83.8% | 84% | Stable | 3 lanes (1.34/1.35/1.36) | Flaky |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.36-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) `[QUARANTINE]` | 51% | 41.7% | Worsening | 3 lanes | Flaky |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) `[QUARANTINE]` | 51% | 41.7% | Worsening | 3 lanes | Flaky |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) `[QUARANTINE]` | 87.5% | 87.5% | Stable | 3 lanes | Flaky |
| [ClusterProfiler subresource access](https://github.com/kubevirt/kubevirt/blob/main/tests/infrastructure/cluster-profiler.go#L62) `[QUARANTINE]` | 93.3% | — | Not in 7d | 1 lane | Flaky |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.35-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) `[QUARANTINE]` | 51% | 37.5% | Worsening | 3 lanes | Flaky |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) `[QUARANTINE]` | 51% | 37.5% | Worsening | 3 lanes | Flaky |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) `[QUARANTINE]` | 89% | — | Not in 7d | 3 lanes | Flaky |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| [Cancel migration — delete source](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) `[QUARANTINE]` | 42% | **0%** | **Severely worsening** | 4+ lanes | Consecutive |
| [Cancel migration — delete target](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) `[QUARANTINE]` | 42% | **0%** | **Severely worsening** | 4+ lanes | Consecutive |
| [Post-copy migration liveness probe](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/postcopy.go#L176) `[QUARANTINE]` | 87.5% | 87.5% | Stable | 1 lane | Flaky |
| [Abort VMI migration](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L1499) | — | 92.3% | — | 1 lane | Flaky |

### Failing Tests — pull-kubevirt-e2e-k8s-1.36-sig-compute-serial (DRA tests)

All DRA tests fail at 0% success rate with 1-2 total runs each. These are **infra-correlated** — they share the same build IDs and all require vfio-gpu hardware that appears unavailable in the test environment. Not quarantine candidates; the test infrastructure needs DRA/vfio-gpu support.

| Test | 28d Rate | 7d Rate | Pattern |
|------|----------|---------|---------|
| [DRA — allocate vfio-gpu by pciBusId](https://github.com/kubevirt/kubevirt/blob/main/tests/) | 0% (1 run) | 0% (1 run) | Infra-correlated |
| DRA — allocate vfio-gpu by vendorID | 0% (1 run) | 0% (1 run) | Infra-correlated |
| DRA — allocate three vfio-gpu devices | 0% (2 runs) | 0% (2 runs) | Infra-correlated |
| DRA — matchAttribute vendorID | 0% (2 runs) | 0% (2 runs) | Infra-correlated |
| DRA — pre-created ResourceClaim | 0% (1 run) | 0% (1 run) | Infra-correlated |
| DRA — four VMIs individual claims | 0% (2 runs) | 0% (2 runs) | Infra-correlated |
| DRA — four VMIs ResourceClaimTemplate | 0% (2 runs) | 0% (2 runs) | Infra-correlated |
| DRA — more than one device per request | 0% (2 runs) | 0% (2 runs) | Infra-correlated |
| DRA — matchattribute different values | 0% (2 runs) | 0% (2 runs) | Infra-correlated |
| DRA — unknown resourceClaim | 0% (1 run) | 0% (1 run) | Infra-correlated |
| DRA — CEL selector no device | 0% (2 runs) | 0% (2 runs) | Infra-correlated |
| DRA — CEL invalid attribute access | 0% (2 runs) | 0% (2 runs) | Infra-correlated |

### Failing Tests — pull-kubevirt-e2e-kind-1.35-sig-compute-arm64

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| [libvirtd logs](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_lifecycle_test.go#L178) | — | 94.8% | — | 2 lanes (arm64 + 1.36 periodic) | Flaky |

---

## sig-storage

**Total 28d failures:** 318 (16.8% of all failures)
**Lanes:** 6 (periodic + presubmit across k8s-1.34/1.35/1.36)

### Periodic Lanes

| Lane | 28d Failures | Share | Builds | 7d Failures | Trend |
|------|-------------|-------|--------|-------------|-------|
| [periodic-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-storage) | 103 | 5.4% | 108 | 50 | Worsening |
| [periodic-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage) | 97 | 5.1% | 110 | 53 | Worsening |
| [periodic-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-storage) | 95 | 5.0% | 109 | 52 | Worsening |

### Presubmit Lanes

| Lane | 28d Failures | Share | Builds | 7d Failures | Trend |
|------|-------------|-------|--------|-------------|-------|
| [pull-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-storage) | 18 | 0.9% | 510 | 4 | Improving |
| [pull-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-storage) | 8 | 0.4% | 213 | 2 | Stable |
| [pull-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-storage) | 6 | 0.3% | 215 | 2 | Stable |

### Failing Tests — periodic sig-storage (all k8s versions)

The cancel migration tests dominate storage lane failures. The error disk test has improved.

| Test | 28d Rate (1.34) | 28d Rate (1.35) | 28d Rate (1.36) | 7d Rate | Trend | Dispersion | Pattern |
|------|-----------------|-----------------|-----------------|---------|-------|------------|---------|
| [Cancel migration — delete source](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) `[QUARANTINE]` | 71.8% | 72.1% | 72.1% | **0%** (all) | **Severely worsening** | 6+ lanes | Consecutive |
| [Cancel migration — delete target](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) `[QUARANTINE]` | 71.8% | 73.1% | 72.1% | **0%** (all) | **Severely worsening** | 6+ lanes | Consecutive |
| [Error disk IO pause](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) `[QUARANTINE]` | 54.1% | 62.6% | 62.2% | — | **Not in 7d** | 3 lanes (1.34/1.35/1.36) | Improving |

---

## sig-network

**Total 28d failures:** 162 (8.6% of all failures)
**Lanes:** 10 (periodic + presubmit across k8s versions + IPv6 + DNC)

### Periodic Lanes

| Lane | 28d Failures | Share | Builds | 7d Failures | Trend |
|------|-------------|-------|--------|-------------|-------|
| [periodic-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network) | 23 | 1.2% | 110 | 3 | Stable |
| [periodic-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network) | 21 | 1.1% | 109 | 9 | Worsening |
| [periodic-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-network) | 9 | 0.5% | 108 | 2 | Stable |
| [periodic-k8s-1.36-sig-network-with-dnc](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network-with-dnc) | 4 | 0.2% | 21 | 1 | Stable |
| [periodic-k8s-1.36-ipv6-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-ipv6-sig-network) | 1 | 0.1% | 87 | — | Stable |

### Presubmit Lanes

| Lane | 28d Failures | Share | Builds | 7d Failures | Trend |
|------|-------------|-------|--------|-------------|-------|
| [pull-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-network) | 30 | 1.6% | 211 | 4 | Improving |
| [pull-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-network) | 27 | 1.4% | 216 | 8 | Stable |
| [pull-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-network) | 24 | 1.3% | 526 | 3 | Improving |
| [pull-k8s-1.36-ipv6-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network) | 10 | 0.5% | 222 | 8 | Worsening |

### Failing Tests

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| [Passt migration connectivity IPv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) `[QUARANTINE]` | 79-91% | 79-91% | Stable | 2 lanes (1.34/1.36 periodic) | Flaky |
| [Bridge nic-hotplug multiple interfaces](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L275) | 85.7% | 85.7% | Stable | 1 lane (DNC only) | Flaky |
| Passt TCP connectivity IPv4 | — | 92.3% | — | 1 lane (1.36 periodic) | Flaky |

---

## sig-operator

**Total 28d failures:** 162 (8.6% of all failures)
**Lanes:** 6 (periodic + presubmit across k8s-1.34/1.35/1.36)

### Lanes

| Lane | 28d Failures | Share | Builds | 7d Failures | Trend |
|------|-------------|-------|--------|-------------|-------|
| [pull-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-operator) | 59 | 3.1% | 198 | 1 | **Improving** |
| [pull-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-operator) | 49 | 2.6% | 491 | 1 | **Improving** |
| [periodic-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-operator) | 45 | 2.4% | 109 | 44 | **Worsening** |
| [periodic-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-operator) | 3 | 0.2% | 108 | 1 | Stable |
| [periodic-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-operator) | 2 | 0.1% | 108 | — | Stable |
| [pull-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-operator) | 1 | 0.1% | 191 | — | Stable |

No test-level failures identified in sig-operator lanes — all failures are Pod-level infrastructure issues. The periodic-k8s-1.34-sig-operator lane shows anomalous 7d failure count (44 of 45 total 28d failures occurred in the last 7 days), suggesting a recent infrastructure regression on k8s-1.34 for this lane. Presubmit lanes are improving significantly.

---

## sig-monitoring

**Total 28d failures:** 61 (3.2% of all failures)
**Lanes:** 2 (periodic-k8s-1.34-sig-monitoring + pull-k8s-1.34-sig-monitoring)

### Lanes

| Lane | 28d Failures | Share | Builds | 7d Failures | Trend |
|------|-------------|-------|--------|-------------|-------|
| [periodic-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-monitoring) | 57 | 3.0% | 56 | 28 | **Worsening** |
| [pull-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-monitoring) | 4 | 0.2% | 82 | 3 | Worsening |

### Failing Tests

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| [Prometheus virt components metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L119) | 59.3% (periodic) / 30% (presubmit) | **0% (periodic)** / 23.3% (presubmit) | **Severely worsening** | 3 lanes (periodic + presubmit + check-for-flakes) | Consecutive |
| [VirtOperatorRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L340) | 66.7% (periodic) / 32.6% (presubmit) | 30% (periodic) / 36.4% (presubmit) | Worsening | 2 lanes | Flaky |
| [VirtOperatorRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L340) `[QUARANTINE]` | 33.3% (periodic) | 33.3% | Stable | 1 lane | Flaky |
| [VM dirty rate metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/vm_monitoring.go#L334) `[QUARANTINE]` | 87% | 85.7% | Stable | 1 lane | Flaky |
| [Alert rule annotations](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/monitoring.go#L68) | 91.8% (presubmit) | 92.9% (periodic) | Stable | 2 lanes | Flaky |
| [VNIC metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L132) | 94.7% (presubmit) | 87.5% (presubmit) | Slightly worsening | 1 lane | Flaky |
| VirtControllerDown bad image | 0% (1 run) | 0% (1 run) | — | 1 lane | Low sample |
| [VirtOperatorDown](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L133) | — | 92.9% | — | 1 lane | Flaky |

**Key finding:** The `should contain virt components metrics` test has deteriorated from 59% success to **0% in the periodic lane** over 7 days (14 consecutive failures). This is the highest-priority quarantine candidate in the project. The `VirtOperatorRESTErrorsBurst` test exists in both quarantined and non-quarantined forms — the non-quarantined version should also be quarantined.

---

## sig-performance

**Total 28d failures:** 12 (0.6% of all failures)
**Lanes:** 3 (periodic-k8s-1.34-sig-performance, periodic-k8s-1.34-sig-performance-kwok, periodic-kind-1.34-sev)

All failures are Pod-level infrastructure issues only. No test-level failures detected. This SIG has the lowest failure rate across the project.

| Lane | 28d Failures | Pod Success Rate |
|------|-------------|-----------------|
| [periodic-k8s-1.34-sig-performance-kwok](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok) | 5 | 88.5% |
| [periodic-kind-1.34-sev](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-kind-1.34-sev) | 4 | 89.2% |
| [periodic-k8s-1.34-sig-performance](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance) | 3 | 88.0% |

---

## Platform-specific

### S390X — periodic-kubevirt-e2e-test-S390X

**28d failures:** 174 (9.2%) | **7d failures:** 39 (7.0%) | **Builds:** 112 (28d), 28 (7d)

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [Eviction strategy — block eviction and migrate](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L86) | 48.6% | 63% | Improving | Flaky |
| [Migration multiple times cloud-init](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L551) | 81.7% | 88.9% | Improving | Flaky |
| [Migrate after virtqemud restart](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L589) | 91.7% | 92.6% | Stable | Flaky |
| [CfgMap/secret fs layout](https://github.com/kubevirt/kubevirt/blob/main/tests/config_test.go#L520) | 94.5% | 88.9% | Slightly worsening | Flaky |
| [CPU hotplug vCPUs](https://github.com/kubevirt/kubevirt/blob/main/tests/hotplug/cpu.go#L106) | — | 92.6% | — | Flaky |
| Port-forward VMI IPv4 | — | 92.6% | — | Flaky |
| MTU verification IPv4 | — | 92.6% | — | Flaky |

The S390X eviction migration test at 48.6% (28d) is the worst non-quarantined test on this platform. It's improving (63% in 7d) but remains unreliable. It is **concentrated to S390X** — other platforms don't show this failure, making it a platform-specific issue rather than a test code problem.

### ARM64 — pull-kubevirt-e2e-kind-1.35-sig-compute-arm64

**28d failures:** 45 (2.4%) | **7d failures:** 15 (2.7%) | Pod-only in 28d, one test failure in 7d.

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [libvirtd logs](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_lifecycle_test.go#L178) | — | 94.8% | — | Flaky |

### Other Presubmit Lanes

| Lane | 28d Failures | Pod Success Rate | 7d |
|------|-------------|------------------|----|
| [pull-kind-sriov](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-sriov) | 11 | 77.1% | 2 |
| [pull-kind-1.35-vgpu](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-vgpu) | 2 | 81.7% | 1 |
| [pull-k8s-1.34-windows2016](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-windows2016) | 1 | 80.3% | — |

All failures in these lanes are Pod-level infrastructure issues only.

---

## Quarantine Candidate Summary

### New Quarantine Candidates

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority |
|------|-----|-------|-------------------|----------|---------|------------|----------|
| [Prometheus virt components metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L119) | sig-monitoring | 3 (periodic + presubmit + check-for-flakes) | 8.3% | Worsening (0% periodic) | Consecutive | Dispersed | **High** |
| [VirtOperatorRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L340) (non-quarantined) | sig-monitoring | 2 (periodic + presubmit) | 32.6% | Worsening | Flaky | Concentrated | **High** |
| [HostDevices PCI passthrough](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L111) (2 tests) | sig-compute | 1 (k8s-1.34 periodic) | 17.1% | Improving (72% in 7d) | Burst → recovering | Concentrated | **Medium** |

### Quarantine Debt (quarantined but still severely broken)

These tests were quarantined but never fixed. They continue to consume CI resources and mask real signal.

| Test | SIG | Worst Rate (28d) | 7d Rate | Lanes | Status |
|------|-----|-------------------|---------|-------|--------|
| [Cancel migration — delete source](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) | sig-compute | 42% | **0%** | 6+ (storage + migrations, all k8s versions) | **Regression — completely broken** |
| [Cancel migration — delete target](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/namespace.go#L595) | sig-compute | 42% | **0%** | 6+ (storage + migrations, all k8s versions) | **Regression — completely broken** |
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L133) | sig-compute | 9.5% | 37.5-41.7% | 3 (1.34/1.35/1.36 periodic) | Improving on 1.34, worsening on 1.35/1.36 |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L134) | sig-compute | 9.5% | 37.5-41.7% | 3 (1.34/1.35/1.36 periodic) | Same as above |
| [Error disk IO pause](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | sig-storage | 54.1% | Not in 7d | 3 (all k8s versions) | May be improving |
| [VirtOperatorRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L340) (quarantined variant) | sig-monitoring | 33.3% | 33.3% | 1 | Stable — not being fixed |

---

## Suggested Follow-Up Analysis

- **lane-failure-rate** on `periodic-kubevirt-e2e-k8s-1.34-sig-operator` — the 7d failure spike (44 of 45 total 28d failures in last 7 days) suggests a recent infrastructure regression that warrants investigation
- **lane-failure-rate** on `periodic-kubevirt-e2e-k8s-1.34-sig-monitoring` — the `virt components metrics` test going to 0% in 7d needs root cause analysis
- **analyze-build** on recent `periodic-kubevirt-e2e-k8s-1.36-sig-compute-migrations` builds — the cancel migration tests hitting 0% success rate suggests a code regression, not flakiness
- **lane-failure-rate** on `periodic-kubevirt-e2e-test-S390X` — the eviction migration test at ~50% needs flip-rate analysis to determine if it's a genuine flake or intermittent infra issue
