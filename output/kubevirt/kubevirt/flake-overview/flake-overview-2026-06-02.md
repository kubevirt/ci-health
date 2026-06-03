# Flake Overview — 2026-06-02

## Summary

| Metric | 28-day | 7-day | Trend |
|--------|--------|-------|-------|
| **Total failures** | 2842 | 441 | **Improving** |
| **Weekly average** | ~710 | 441 | -37.9% |
| **Lanes with failures** | 42 | 35 | — |

**Overall trend: Improving.** The 7-day failure count (441) is 38% lower than the 28-day weekly average (~710). Most SIGs show reduced failure counts, with sig-storage seeing the biggest improvement.

**Top 3 worst SIGs by total failures (28d):**
1. **sig-compute** — 1700 failures (59.8%), driven by USB passthrough, PCI device, and VSOCK tests
2. **sig-storage** — 523 failures (18.4%), dominated by the quarantined IO error test
3. **Platform-specific (S390X)** — 273 failures (9.6%), live migration and MemBalloon flakes

**Quarantine candidates identified:** 1 high priority, 1 medium (regression — needs investigation, not quarantine)

**Infra-unstable lanes:** 12 presubmit lanes have Pod-level failure rates of 17–29%, indicating significant infrastructure instability across the board.

**Critical regression detected:** HostDevices PCI passthrough tests on k8s-1.34 dropped from 54% to **0% success** in the last 7 days — 25 consecutive failures indicate deterministic breakage, not flakiness.

## Navigation

- [Infrastructure](#infrastructure)
- [sig-compute](#sig-compute)
- [sig-storage](#sig-storage)
- [sig-network](#sig-network)
- [sig-operator](#sig-operator)
- [sig-monitoring](#sig-monitoring)
- [sig-performance](#sig-performance)
- [Platform-specific (S390X)](#platform-specific-s390x)
- [Quarantine Candidate Summary](#quarantine-candidate-summary)
- [Quarantine Debt](#quarantine-debt)

---

## Infrastructure

### Presubmit Pod-Level Failure Rates

Pod-level failures indicate infrastructure problems (cluster provisioning, resource exhaustion, timeouts) rather than test-specific issues. These failures prevent any tests from running.

| Lane | 28d Success | 7d Success | Trend | Failed/Total (28d) |
|------|-------------|------------|-------|---------------------|
| pull-kubevirt-e2e-kind-1.35-sig-compute-arm64 | 71.3% | — | — | 142/494 |
| pull-kubevirt-e2e-k8s-1.36-sig-storage | 75.3% | 73.6% | Worsening | 86/348 |
| pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations | 75.8% | 75.3% | Stable | 102/422 |
| pull-kubevirt-e2e-k8s-1.36-sig-compute | 76.1% | 73.3% | Worsening | 84/351 |
| pull-kubevirt-e2e-k8s-1.35-sig-storage | 76.8% | 76.2% | Stable | 88/379 |
| pull-kubevirt-e2e-k8s-1.36-sig-compute-serial | 77.0% | 75.8% | Stable | 74/322 |
| pull-kubevirt-e2e-k8s-1.36-sig-operator | 78.5% | 76.0% | Worsening | 65/302 |
| pull-kubevirt-e2e-k8s-1.35-sig-operator | 80.8% | 81.7% | Stable | 66/344 |
| pull-kubevirt-e2e-k8s-1.34-sig-monitoring | 81.4% | 76.9% | Worsening | 8/43 |
| pull-kubevirt-e2e-k8s-1.35-sig-network | 81.6% | 80.6% | Stable | 71/385 |
| pull-kubevirt-e2e-k8s-1.34-sig-performance | 83.0% | — | — | 42/247 |
| pull-kubevirt-e2e-k8s-1.35-sig-compute | 83.1% | 82.9% | Stable | 59/349 |
| pull-kubevirt-e2e-k8s-1.34-sig-storage | 90.9% | 90.7% | Stable | 12/132 |
| pull-kubevirt-check-tests-for-flakes | 93.5% | 92.8% | Stable | 26/402 |
| pull-kubevirt-e2e-k8s-1.34-sig-network | 94.5% | — | — | 7/127 |

**Key observation:** Most k8s-1.36 and k8s-1.35 presubmit lanes hover around 73–83% Pod success rate — roughly 1 in 4–5 job runs fail at infrastructure level before tests execute. The k8s-1.36 lanes are slightly worsening. This systemic infrastructure instability is the single largest contributor to CI noise.

---

## sig-compute

**28d total: 1700 failures (59.8%) | 7d total: 273 failures (61.9%)**

### Lanes

| Lane | Type | 28d Failures | 28d Share | 7d Failures | 7d Share | Trend |
|------|------|-------------|-----------|-------------|----------|-------|
| [periodic-kubevirt-e2e-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute) | Periodic | 941 | 33.1% | 50 | 11.3% | Improving |
| [periodic-kubevirt-e2e-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-compute) | Periodic | 343 | 12.1% | 112 | 25.4% | **Worsening** |
| [periodic-kubevirt-e2e-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute) | Periodic | 208 | 7.3% | 17 | 3.9% | Improving |
| [pull-kubevirt-e2e-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute) | Presubmit | 90 | 3.2% | 53 | 12.0% | Worsening |
| [pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | Presubmit | 33 | 1.2% | 13 | 3.0% | Stable |
| [pull-kubevirt-e2e-k8s-1.35-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-serial) | Presubmit | 29 | 1.0% | 4 | 0.9% | Stable |
| [periodic-kubevirt-e2e-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | Periodic | 23 | 0.8% | 13 | 3.0% | Worsening |
| [pull-kubevirt-e2e-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-compute) | Presubmit | 10 | 0.4% | 3 | 0.7% | Stable |
| [periodic-kubevirt-e2e-kind-1.34-sev](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-kind-1.34-sev) | Periodic | 9 | 0.3% | 3 | 0.7% | Stable |
| [pull-kubevirt-e2e-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute) | Presubmit | 7 | 0.2% | 4 | 0.9% | Stable |
| [pull-kubevirt-e2e-k8s-1.36-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute-serial) | Presubmit | 7 | 0.2% | 1 | 0.2% | Stable |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.34-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [HostDevices — Should successfully passthrough an emulated PCI device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | 54.4% | **0%** | **Regression** | Consecutive | Concentrated (1 lane) |
| [HostDevices — Should successfully passthrough 2 emulated PCI devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | 54.4% | **0%** | **Regression** | Consecutive | Concentrated (1 lane) |
| [USB Passthrough — 1 emulated USB device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) `[QUARANTINE]` | 26.2% | **0%** | Worsening | Consecutive | Dispersed (3 lanes) |
| [USB Passthrough — 2 emulated USB devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) `[QUARANTINE]` | 26.2% | **0%** | Worsening | Consecutive | Dispersed (3 lanes) |
| [Guest console log flood QOSGuaranteed](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) `[QUARANTINE]` | 84.5% | 80.0% | Stable | Flaky | Dispersed (3 lanes) |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.36-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [virtiofs ServiceAccount — namespace and token](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L173) `[QUARANTINE]` | 0% | 0% | Broken | Consecutive | Concentrated (1 lane) |
| [USB Passthrough — 1 emulated USB device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) `[QUARANTINE]` | 62.1% | 70.4% | Improving | Flaky | Dispersed (3 lanes) |
| [USB Passthrough — 2 emulated USB devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) `[QUARANTINE]` | 62.1% | 70.4% | Improving | Flaky | Dispersed (3 lanes) |
| [VSOCK — TLS on both sides](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L223) | 84.5% | 92.6% | Improving | Flaky | Concentrated (1 lane) |
| [VSOCK — err if no app listens](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L240) | 84.5% | — | Improving | Flaky | Concentrated (1 lane) |
| [VSOCK — retain CID for migration target](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L121) | 85.4% | — | Improving | Flaky | Concentrated (1 lane) |
| [VSOCK — non-transitional](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L65) | 88.3% | — | Improving | Flaky | Concentrated (1 lane) |
| [VSOCK — without TLS](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L224) | 88.3% | — | Improving | Flaky | Concentrated (1 lane) |
| [VSOCK — transitional](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L65) | 89.3% | — | Improving | Flaky | Concentrated (1 lane) |
| [Guest console log flood QOSGuaranteed](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) `[QUARANTINE]` | 90.3% | 92.6% | Stable | Flaky | Dispersed (3 lanes) |
| [VSOCK — invalid port](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L227) | 92.2% | — | Improving | Flaky | Concentrated (1 lane) |
| [ClusterProfiler — subresource access](https://github.com/kubevirt/kubevirt/blob/main/tests/infrastructure/cluster-profiler.go#L62) `[QUARANTINE]` | 93.5% | 92.6% | Stable | Flaky | Concentrated (1 lane) |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.35-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [USB Passthrough — 1 emulated USB device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) `[QUARANTINE]` | 45.2% | 63.6% | Improving | Flaky | Dispersed (3 lanes) |
| [USB Passthrough — 2 emulated USB devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) `[QUARANTINE]` | 45.2% | 63.6% | Improving | Flaky | Dispersed (3 lanes) |
| [Guest console log flood QOSGuaranteed](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) `[QUARANTINE]` | 83.7% | — | Improving | Flaky | Dispersed (3 lanes) |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.35-sig-compute-migrations

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [Migration bandwidth limitations — different durations](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L262) | — | 80.8% | New in 7d | Flaky | Concentrated (1 lane) |
| [Backup — preserve checkpoints after migration](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/backup.go#L1167) | — | 92.3% | New in 7d | Flaky | Concentrated (1 lane) |
| [Live Migrations with priority](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/priority.go#L76) | 93.5% | 92.3% | Stable | Flaky | Concentrated (1 lane) |

### Failing Tests — periodic-kubevirt-e2e-kind-1.34-sev

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [SEV — guest attestation](https://github.com/kubevirt/kubevirt/blob/main/tests/launchsecurity/sev.go#L349) `[QUARANTINE]` | 89.0% | 85.0% | Worsening | Flaky | Concentrated (1 lane) |

---

## sig-storage

**28d total: 523 failures (18.4%) | 7d total: 63 failures (14.3%) | Trend: Improving**

### Lanes

| Lane | Type | 28d Failures | 28d Share | 7d Failures | 7d Share | Trend |
|------|------|-------------|-----------|-------------|----------|-------|
| [periodic-kubevirt-e2e-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-storage) | Periodic | 244 | 8.6% | 16 | 3.6% | Improving |
| [periodic-kubevirt-e2e-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage) | Periodic | 111 | 3.9% | 15 | 3.4% | Stable |
| [periodic-kubevirt-e2e-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-storage) | Periodic | 81 | 2.9% | 17 | 3.9% | Stable |
| [pull-kubevirt-e2e-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-storage) | Presubmit | 73 | 2.6% | 12 | 2.7% | Stable |
| [pull-kubevirt-e2e-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-storage) | Presubmit | 8 | 0.3% | 2 | 0.5% | Stable |
| [pull-kubevirt-e2e-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-storage) | Presubmit | 6 | 0.2% | 1 | 0.2% | Stable |

### Failing Tests — periodic lanes (all k8s versions)

| Test | Lane | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|------|----------|---------|-------|---------|------------|
| [Storage error disk — pause VMI on IO error](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) `[QUARANTINE]` | k8s-1.34 | 35.6% | 28.0% | **Worsening** | Flaky | Dispersed (3 lanes) |
| [Storage error disk — pause VMI on IO error](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) `[QUARANTINE]` | k8s-1.35 | 44.7% | 48.1% | Stable | Flaky | Dispersed (3 lanes) |
| [Storage error disk — pause VMI on IO error](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) `[QUARANTINE]` | k8s-1.36 | 50.0% | 57.1% | Improving | Flaky | Dispersed (3 lanes) |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.36-sig-storage (non-quarantined)

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [Volume migration — block to filesystem](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1275) | 76.4% | — | Improving | Flaky | Concentrated (2 lanes) |
| [Volume migration — filesystem to filesystem](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1273) | 85.8% | — | Improving | Flaky | Concentrated (2 lanes) |
| [Volume migration — containerdisk + hotplugged](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1132) | 90.6% | — | Improving | Flaky | Concentrated (2 lanes) |

### Failing Tests — pull-kubevirt-e2e-k8s-1.36-sig-storage (non-quarantined)

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| Volume migration — block to filesystem | 80.5% | — | Improving | Flaky | Concentrated (2 lanes) |
| Volume migration — filesystem to filesystem | 81.0% | — | Improving | Flaky | Concentrated (2 lanes) |
| Volume migration — containerdisk + hotplugged | 87.8% | — | Improving | Flaky | Concentrated (2 lanes) |

---

## sig-network

**28d total: 74 failures (2.6%) | 7d total: 30 failures (6.8%) | Trend: Worsening (proportionally)**

### Lanes

| Lane | Type | 28d Failures | 7d Failures | Trend |
|------|------|-------------|-------------|-------|
| [periodic-kubevirt-e2e-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network) | Periodic | 11 | 8 | Worsening |
| [periodic-kubevirt-e2e-k8s-1.35-ipv6-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-ipv6-sig-network) | Periodic | 13 | 3 | Stable |
| [periodic-kubevirt-e2e-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network) | Periodic | 7 | 2 | Stable |
| [periodic-kubevirt-e2e-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-network) | Periodic | 5 | — | — |
| [periodic-kubevirt-e2e-k8s-1.35-sig-network-with-dnc](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network-with-dnc) | Periodic | 3 | 2 | Stable |
| [pull-kubevirt-e2e-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-network) | Presubmit | 13 | 4 | Stable |
| [pull-kubevirt-e2e-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-network) | Presubmit | 11 | 4 | Stable |
| [pull-kubevirt-e2e-kind-sriov](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-sriov) | Presubmit | 6 | 5 | Worsening |
| [pull-kubevirt-check-tests-for-flakes](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-check-tests-for-flakes) | Presubmit | 2 | 2 | Stable |

### Cross-Lane Failing Test: passt network binding migration ipv4

This test fails across 5+ lanes — a genuine cross-environment flake.

| Lane | 28d Rate | 7d Rate | Trend |
|------|----------|---------|-------|
| [periodic-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network) | 87.2% | **77.8%** | **Worsening** |
| [periodic-k8s-1.35-sig-network-with-dnc](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network-with-dnc) | 90.0% | 83.3% | Worsening |
| [pull-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-network) | 93.5% | 90.4% | Stable |
| [periodic-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network) | — | 92.0% | — |
| [pull-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-network) | — | 94.2% | — |

Source: [tests/network/passt.go#L286](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L286)

**Analysis:** Worsening trend in 7d, particularly in periodic k8s-1.35 (87% → 78%). The test is dispersed across both periodic and presubmit lanes, across k8s versions and including DNC — genuine cross-environment flake. **High priority quarantine candidate.**

### Failing Tests — pull-kubevirt-check-tests-for-flakes (bridge nic-hotplug)

| Test | 28d Rate | 7d Rate | Runs | Pattern |
|------|----------|---------|------|---------|
| [bridge nic-hotplug — can be hotplugged](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L118) | 0% | 0% | 2 | Insufficient data |
| [bridge nic-hotplug — migrate with hotplugged](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L145) | 0% | 0% | 2 | Insufficient data |
| [bridge nic-hotplug — secondary network connectivity](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L159) | 0% | 0% | 2 | Insufficient data |
| [bridge nic-hotplug — multiple interfaces](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L215) | 0% | 0% | 2 | Insufficient data |
| [bridge nic-hotunplug — succeed](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L335) | 0% | 0% | 2 | Insufficient data |

All 5 bridge nic-hotplug "In place" tests failed in both their runs (100% failure). The extremely low sample size (2 runs total over 28 days) means these failures could be from a single bad build. Monitor — if consistent, these tests have a systemic issue.

### DRA-SRIOV Tests (pull-kubevirt-e2e-kind-sriov)

10 DRA-SRIOV tests each show 66.7% (28d, 1/3) or 50% (7d, 1/2) success rate. These tests are too new (only 2–3 runs) to draw conclusions. The DRA-SRIOV test source files are not present in the local kubevirt checkout — likely very recent additions. Monitor as test volume increases.

---

## sig-operator

**28d total: 114 failures (4.0%) | 7d total: 17 failures (3.9%) | Trend: Stable**

### Lanes

| Lane | Type | 28d Failures | 7d Failures | Trend |
|------|------|-------------|-------------|-------|
| [pull-kubevirt-e2e-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-operator) | Presubmit | 61 | 13 | Stable |
| [periodic-kubevirt-e2e-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-operator) | Periodic | 48 | 2 | Improving |
| [pull-kubevirt-e2e-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-operator) | Presubmit | 2 | 1 | Stable |
| [pull-kubevirt-e2e-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-operator) | Presubmit | 1 | 1 | Stable |
| [periodic-kubevirt-e2e-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-operator) | Periodic | 1 | — | — |
| [periodic-kubevirt-e2e-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-operator) | Periodic | 1 | — | — |

No test-level failures reported for sig-operator lanes (only Pod-level infrastructure failures). The sig-operator failure count is entirely driven by Pod provisioning issues, not test flakiness.

---

## sig-monitoring

**28d total: 77 failures (2.7%) | 7d total: 10 failures (2.3%) | Trend: Stable**

### Lanes

| Lane | Type | 28d Failures | 7d Failures | Trend |
|------|------|-------------|-------------|-------|
| [periodic-kubevirt-e2e-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-monitoring) | Periodic | 74 | 9 | Stable |
| [pull-kubevirt-e2e-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-monitoring) | Presubmit | 3 | 1 | Stable |

### Failing Tests — periodic-kubevirt-e2e-k8s-1.34-sig-monitoring

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [VM dirty rate metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/vm_monitoring.go#L334) `[QUARANTINE]` | 76.9% | 75.0% | Stable | Flaky | Concentrated (1 lane) |
| [Prometheus metrics — transition time from VM deletion](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L179) | 86.3% | — | Improving | Flaky | Concentrated (1 lane) |
| [VirtOperatorRESTErrorsBurst — triggered when failing](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L340) | 88.5% | 88.9% | Stable | Flaky | Concentrated (1 lane) |

### Failing Tests — pull-kubevirt-e2e-k8s-1.34-sig-monitoring

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [Prometheus metrics — contain virt components metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L116) | 72.4% | 69.2% | Stable | Flaky |
| VM reboot metrics — expose kubevirt_vm_reboot_total | 0% | 0% | — | Insufficient data (1 run) |
| VM reboot metrics — load VMPVPanicReboot alert rule | 0% | 0% | — | Insufficient data (1 run) |

---

## sig-performance

**28d total: 53 failures (1.9%) | 7d total: 4 failures (0.9%) | Trend: Improving**

### Lanes

| Lane | Type | 28d Failures | 7d Failures | Trend |
|------|------|-------------|-------------|-------|
| [periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok-100](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok-100) | Periodic | 41 | — | Improving |
| [periodic-kubevirt-e2e-k8s-1.34-sig-performance](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance) | Periodic | 5 | 3 | Stable |
| [periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok) | Periodic | 5 | 1 | Stable |
| [pull-kubevirt-e2e-k8s-1.34-sig-performance](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-performance) | Presubmit | 2 | — | — |

### Failing Tests

| Test | Lane | 28d Rate | 7d Rate | Trend |
|------|------|----------|---------|-------|
| [Density test — create 100 running VMs](https://github.com/kubevirt/kubevirt/blob/main/tests/performance/density.go#L102) | periodic-perf | — | 94.7% | New in 7d |
| [KWOK density — create fake VMs](https://github.com/kubevirt/kubevirt/blob/main/tests/performance/density-kwok.go#L105) | periodic-perf-kwok | — | 94.7% | New in 7d |

Both performance failures are at >94% success rate with very low failure counts. No action needed.

---

## Platform-specific (S390X)

**28d total: 273 failures (9.6%) | 7d total: 43 failures (9.8%) | Trend: Stable**

### Failing Tests — periodic-kubevirt-e2e-test-S390X

| Test | 28d Rate | 7d Rate | Trend | Pattern | Dispersion |
|------|----------|---------|-------|---------|------------|
| [Live Migration eviction — block eviction api and migrate](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | 57.7% | 65.4% | Improving | Flaky | Concentrated (S390X) |
| [VM Live Migration Alpine — multiple times with cloud-init](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L547) | 83.2% | 80.8% | Stable | Flaky | Concentrated (S390X) |
| [VM Live Migration Alpine — virtqemud restarted](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L585) | 85.0% | — | Improving | Flaky | Concentrated (S390X) |
| [VMIDefaults MemBalloon — period 0](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L121) | 88.7% | — | Improving | Flaky | Concentrated (S390X) |
| [VM Live Migration Alpine — LiveMigrateIfPossible](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L342) | 88.8% | — | Improving | Flaky | Concentrated (S390X) |
| [Networkpolicy — allow traffic ports 80 and 81](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L242) | 88.8% | — | Improving | Flaky | Concentrated (S390X) |
| [VMIDefaults MemBalloon — period 12](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L121) | 90.6% | — | Improving | Flaky | Concentrated (S390X) |
| [VMIDefaults MemBalloon — present in domain](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L92) | 90.7% | — | Improving | Flaky | Concentrated (S390X) |
| [Networkpolicy — fail pinging different namespaces](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L138) | 92.5% | — | Improving | Flaky | Concentrated (S390X) |
| [Networkpolicy — allow port 80 deny port 81](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L271) | 94.4% | 92.3% | Stable | Flaky | Concentrated (S390X) |
| [Node-labeller — update cpu model vendor label](https://github.com/kubevirt/kubevirt/blob/main/tests/infrastructure/node-labeller.go#L255) | 94.2% | — | Improving | Flaky | Concentrated (S390X) |
| [VM Live Migration — persistent libvirt domain](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L630) | 94.4% | — | Improving | Flaky | Concentrated (S390X) |
| [Host-model migration — no SupportedHostModelMigrationCPU label](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/host_model.go#L202) | 94.2% | — | Improving | Flaky | Concentrated (S390X) |
| [Networkpolicy — succeed pinging same namespace](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L123) | 94.4% | — | Improving | Flaky | Concentrated (S390X) |
| [Networkpolicy — reach from alternative namespace](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L208) | 94.4% | — | Improving | Flaky | Concentrated (S390X) |

**Analysis:** S390X failures are concentrated on this single platform lane with no cross-lane dispersion. The worst offender is the live migration eviction test at 57.7% (28d), but it's improving to 65.4% in the 7d window. Most other tests are >85% and improving. The overall failure pattern is S390X-specific — likely related to the platform's slower I/O and timing differences rather than test logic issues.

---

## Quarantine Candidate Summary

### New Quarantine Candidates

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority |
|------|-----|-------|-------------------|----------|---------|------------|----------|
| [passt migration — ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L286) | sig-network | 5+ | 87.2% | **Worsening** (→78%) | Flaky | Dispersed | **High** |
| [VSOCK — TLS on both sides](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L223) | sig-compute | 1 | 84.5% | Improving (→93%) | Flaky | Concentrated | Low |
| [VSOCK — err if no app listens](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L240) | sig-compute | 1 | 84.5% | Improving | Flaky | Concentrated | Low |
| [VSOCK — retain CID for migration](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L121) | sig-compute | 1 | 85.4% | Improving | Flaky | Concentrated | Low |

### Regression — Needs Investigation (NOT Quarantine)

| Test | SIG | Lane | 28d Rate | 7d Rate | Pattern |
|------|-----|------|----------|---------|---------|
| [HostDevices — passthrough emulated PCI device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | sig-compute | k8s-1.34 periodic | 54.4% | **0%** | **25 consecutive failures** |
| [HostDevices — passthrough 2 emulated PCI devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | sig-compute | k8s-1.34 periodic | 54.4% | **0%** | **25 consecutive failures** |

These PCI device passthrough tests went from 54% success to **0% in the last 7 days** on k8s-1.34 — all 25 runs failed consecutively. This is deterministic breakage, not a flake. Investigate recent changes to the k8s-1.34 cluster configuration or PCI device emulation setup.

---

## Quarantine Debt

These tests are already quarantined but remain severely broken (success rate <50%). They have been quarantined but never fixed:

| Test | Worst Lane | 28d Rate | 7d Rate | Status |
|------|-----------|----------|---------|--------|
| [virtiofs ServiceAccount](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L173) | k8s-1.36 periodic | **0%** | **0%** | Never passes (22/22 failed) |
| [USB Passthrough — 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | k8s-1.34 periodic | 26.2% | **0%** | Worsening on k8s-1.34 |
| [USB Passthrough — 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | k8s-1.34 periodic | 26.2% | **0%** | Worsening on k8s-1.34 |
| [Storage error disk — IO error](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | k8s-1.34 periodic | 35.6% | **28.0%** | Worsening |
| [USB Passthrough — 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | k8s-1.35 periodic | 45.2% | 63.6% | Improving on k8s-1.35 |
| [USB Passthrough — 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | k8s-1.35 periodic | 45.2% | 63.6% | Improving on k8s-1.35 |
