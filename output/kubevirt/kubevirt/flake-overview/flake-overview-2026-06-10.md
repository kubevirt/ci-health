# KubeVirt Flake Overview — 2026-06-10

## Summary

| Metric | 28-day | 7-day | Trend |
|--------|--------|-------|-------|
| **Total failures** | 2598 | 348 | |
| **Active lanes** | 45 | 28 | |
| **Daily failure rate** | ~93/day | ~50/day | **Improving** (-53%) |
| **Quarantine candidates** | 2 high, 2 medium | | |
| **Infra-unstable lanes** | ~10 (Pod failure rate <85%) | 5 | |

**Overall trend: Improving.** The 7-day daily failure rate (~50/day) is significantly lower than the prior 21-day rate (~107/day). Total 28-day failures dropped from 2641 (June 8) to 2598. The virtiofs ServiceAccount test shows zero 7d failures for two consecutive report cycles, confirming the fix has held. VSOCK tests and DRA tests that appeared in prior reports have dropped out entirely — zero failures. The S390X eviction test continues to worsen (26% success in 7d, down from 46% in 28d).

**Top 3 worst SIGs (28d):**
1. **sig-compute** — 1569 failures (60.4%) — HostDevices PCI deterministic failure, USB quarantine debt
2. **sig-storage** — 508 failures (19.6%) — IO error quarantine debt persists across all k8s versions
3. **Platform-specific** — 274 failures (10.5%) — S390X eviction continues to degrade

## Navigation

- [Infrastructure](#infrastructure)
- [sig-compute](#sig-compute)
- [sig-storage](#sig-storage)
- [sig-network](#sig-network)
- [sig-operator](#sig-operator)
- [sig-monitoring](#sig-monitoring)
- [sig-compute-migrations](#sig-compute-migrations)
- [Platform-specific](#platform-specific)
- [Quarantine Status Changes](#quarantine-status-changes)
- [Quarantine Candidate Summary](#quarantine-candidate-summary)

---

## Infrastructure

Pod-level setup failures indicate infrastructure instability rather than individual test flakes. These affect all tests in a lane when they occur.

| Lane | 28d Pod Rate | 7d Pod Rate | Trend |
|------|-------------|-------------|-------|
| pull-kubevirt-e2e-kind-1.35-sig-compute-arm64 | 68.2% | 83.9% | **Improving** |
| pull-kubevirt-e2e-k8s-1.36-sig-storage | 74.2% | 75.8% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-compute | 75.4% | 77.2% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-network | 76.3% | 79.6% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-compute-serial | 77.1% | 79.6% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-operator | 78.2% | — | — |
| pull-kubevirt-e2e-k8s-1.34-sig-performance | 78.6% | — | — |
| pull-kubevirt-e2e-k8s-1.34-sig-monitoring | 79.4% | — | — |
| pull-kubevirt-e2e-k8s-1.35-sig-storage | 81.3% | — | — |
| pull-kubevirt-e2e-k8s-1.35-sig-network | 85.7% | — | — |
| pull-kubevirt-e2e-k8s-1.35-sig-operator | 85.9% | — | — |
| pull-kubevirt-e2e-k8s-1.35-sig-compute | 86.6% | — | — |
| pull-kubevirt-check-tests-for-flakes | 91.3% | — | — |
| pull-kubevirt-e2e-k8s-1.34-sig-storage | 93.3% | — | — |

**Key observations:**
- The arm64 lane improved significantly from 68% to 84% — the strongest infra improvement.
- The k8s-1.36 presubmit lanes remain at ~75–80% Pod success — roughly 1 in 4 runs fail at the infrastructure level before tests even execute.
- k8s-1.35 lanes are generally healthier (81–87%) than k8s-1.36 lanes.
- Lanes marked "—" in the 7d column had too few flakefinder-reported failures to appear in the 7d output; their Pod rates could not be computed.

---

## sig-compute

**Total 28d failures: 1569 (60.4%) | 7d failures: 191 (54.9%)**

### Periodic lanes

| Lane | 28d Failures | 7d Failures | 28d Share | Trend |
|------|-------------|-------------|-----------|-------|
| [periodic-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute) | 801 | 32 | 30.8% | **Improving** |
| [periodic-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-compute) | 403 | 117 | 15.5% | **Worsening** |
| [periodic-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute) | 156 | 24 | 6.0% | Stable |

#### periodic-kubevirt-e2e-k8s-1.34-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|----------|---------|-------|------------|---------|-------|
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | 9.7% | 0.0% | **Worsening** | dispersed (3 lanes) | consecutive | `[QUARANTINE]` |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | 9.7% | 0.0% | **Worsening** | dispersed (3 lanes) | consecutive | `[QUARANTINE]` |
| [HostDevices PCI 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | 24.3% | 0.0% | **Worsening** | concentrated (1 lane) | consecutive | |
| [HostDevices PCI 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | 24.3% | 0.0% | **Worsening** | concentrated (1 lane) | consecutive | |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) | 83.5% | 80.8% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [VMI libvirtd logs](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_lifecycle_test.go#L178) | 94.2% | 92.3% | Stable | concentrated (1 lane) | flaky | |

The k8s-1.34 lane is now the worst sig-compute lane. HostDevices PCI and USB passthrough are at 0% for the 7-day window — completely broken. All 4 tests show consecutive failure patterns (not intermittent), indicating deterministic breakage rather than flakiness.

#### periodic-kubevirt-e2e-k8s-1.36-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|----------|---------|-------|------------|---------|-------|
| [virtiofs ServiceAccount](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L232) | 32.9% | — | **Improving** | concentrated (2 lanes) | consecutive | |
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | 57.3% | 53.8% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | 57.3% | 53.8% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) | 89.3% | 88.5% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |

virtiofs ServiceAccount has zero 7d failures for two consecutive report cycles (since June 8), confirming a recent fix. The 28d rate improved from 23.4% (June 8) to 32.9% as old failures age out. VSOCK tests and ClusterProfiler that appeared in prior reports are completely absent — also fixed.

#### periodic-kubevirt-e2e-k8s-1.35-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|----------|---------|-------|------------|---------|-------|
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | 51.9% | 63.0% | **Improving** | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | 51.9% | 63.0% | **Improving** | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) | 85.6% | 81.5% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |

### Presubmit lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute) | 102 | 6 | **Improving** |
| [pull-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-compute) | 11 | 1 | Stable |
| [pull-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute) | 8 | 0 | Stable |
| [pull-k8s-1.36-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute-serial) | 8 | 1 | Stable |
| [pull-k8s-1.35-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-serial) | 18 | 1 | **Improving** |

#### pull-kubevirt-e2e-k8s-1.36-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| [virtiofs ServiceAccount](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L232) | 78.6% | — | **Improving** | concentrated (2 lanes) | consecutive |

The virtiofs ServiceAccount failure that was the top presubmit sig-compute issue in previous reports now shows zero 7d failures, consistent with the periodic lane improvement.

DRA tests that appeared in the June 8 report (11 tests at 0% across check-tests-for-flakes and compute-serial lanes) are no longer in the data — either the tests were removed or the infra issue was resolved.

### SEV lane

| Lane | 28d Failures | Trend |
|------|-------------|-------|
| [periodic-kind-1.34-sev](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-kind-1.34-sev) | 6 | Stable |

| Test | 28d Rate | Trend | Flags |
|------|----------|-------|-------|
| [SEV guest attestation](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/sev_test.go) | 92.8% | Stable | `[QUARANTINE]` |

SEV guest attestation remains stable at ~93% success. De-quarantine tracking: [#15378](https://github.com/kubevirt/kubevirt/issues/15378).

---

## sig-storage

**Total 28d failures: 508 (19.6%) | 7d failures: 57 (16.4%)**

### Periodic lanes

| Lane | 28d Failures | 7d Failures | 28d Share | Trend |
|------|-------------|-------------|-----------|-------|
| [periodic-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-storage) | 242 | 16 | 9.3% | **Improving** |
| [periodic-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage) | 98 | 20 | 3.8% | Stable |
| [periodic-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-storage) | 80 | 17 | 3.1% | Stable |

#### Failing tests — periodic storage lanes (all 3 k8s versions)

| Test | Lane | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|------|----------|---------|-------|------------|---------|-------|
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | k8s-1.34 | 37.9% | 38.5% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | k8s-1.36 | 43.9% | 37.0% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | k8s-1.35 | 43.8% | 48.1% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |

IO error pause VMI remains consistently broken across all 3 k8s versions at 37–48% success. This test has been in quarantine debt since at least May and shows no improvement trend.

#### Volume migration tests (k8s-1.36 only)

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [Volume migrate block→fs](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1274) | 85.0% | — | **Improving** | flaky |
| [Volume migrate containerdisk+hotplug](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go) | 90.7% | — | **Improving** | flaky |
| [Volume migrate fs→fs](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go) | 92.5% | — | **Improving** | flaky |

Volume migration tests show zero 7d failures for two consecutive cycles (since June 8), suggesting stabilization. The block→fs test was a quarantine candidate in June 5 and can likely be downgraded.

### Presubmit lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-storage) | 75 | 1 | **Improving** |
| [pull-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-storage) | 7 | 2 | Stable |
| [pull-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-storage) | 6 | 1 | Stable |

Presubmit storage failures are predominantly Pod-level infrastructure issues. One individual test appeared: [Backup FS Freeze](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/) in pull-k8s-1.34 with 1 run / 1 failure — insufficient sample to assess.

---

## sig-network

**Total 28d failures: 101 (3.9%) | 7d failures: 35 (10.1%)**

The 7d share increased from 3.9% to 10.1% because other SIGs improved faster while network failures remained steady.

### Periodic lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [periodic-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network) | 19 | 7 | Stable |
| [periodic-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network) | 13 | 5 | Stable |
| [periodic-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-network) | 7 | 3 | Stable |

#### Failing tests — periodic network lanes

| Test | Lane | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|------|----------|---------|-------|------------|---------|
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | k8s-1.35 | 83.3% | 81.5% | Stable | dispersed (5+ lanes) | flaky |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | k8s-1.34 | — | 88.5% | — | dispersed (5+ lanes) | flaky |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | k8s-1.36 | 94.9% | 92.9% | Stable | dispersed (5+ lanes) | flaky |
| [bridge nic-hotplug Migration based](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L215) | k8s-1.36 | — | 92.9% | New in 7d | concentrated (1 lane) | flaky |

### Other periodic network lanes

| Lane | 28d Failures | 7d Failures | Notes |
|------|-------------|-------------|-------|
| periodic-k8s-1.35-ipv6-sig-network | 9 | 0 | Testgrid 404 — lane may be retired |
| periodic-k8s-1.35-sig-network-with-dnc | 4 | 0 | Testgrid 404 — lane may be retired |
| periodic-k8s-1.36-sig-network-with-dnc | 1 | 1 | [bridge nic-hotplug In place](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L215) at 85.7% |
| periodic-k8s-1.36-ipv6-sig-network | 1 | 1 | — |

### Presubmit lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-network) | 19 | 9 | Stable |
| [pull-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-network) | 15 | 3 | **Improving** |
| [pull-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-network) | 7 | 4 | Stable |

#### Failing tests — presubmit network lanes

| Test | Lane | 28d Rate | 7d Rate | Dispersion | Pattern |
|------|------|----------|---------|------------|---------|
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | pull-k8s-1.35 | 90.0% | 83.1% | dispersed (5+ lanes) | flaky |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | pull-k8s-1.36 | 88.8% | 87.4% | dispersed (5+ lanes) | flaky |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | pull-k8s-1.34 | 89.6% | 87.7% | dispersed (5+ lanes) | flaky |

---

## sig-operator

**Total 28d failures: 68 (2.6%) | 7d failures: 1 (0.3%)**

### Periodic lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [periodic-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-operator) | 4 | 0 | Stable |
| [periodic-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-operator) | 1 | 1 | Stable |

### Presubmit lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-operator) | 45 | 0 | **Improving** |
| [pull-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-operator) | 16 | 0 | **Improving** |
| [pull-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-operator) | 2 | 0 | Stable |

sig-operator failures were almost entirely Pod-level infrastructure issues, not test flakes. The dramatic improvement from 68 (28d) to 1 (7d) reflects infra stabilization.

---

## sig-monitoring

**Total 28d failures: 60 (2.3%) | 7d failures: 4 (1.1%)**

### Periodic lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [periodic-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-monitoring) | 57 | 4 | **Improving** |

#### periodic-kubevirt-e2e-k8s-1.34-sig-monitoring

| Test | 28d Rate | 7d Rate | Trend | Pattern | Flags |
|------|----------|---------|-------|---------|-------|
| [VM dirty rate metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/vm_monitoring.go) | 77.4% | 85.7% | **Improving** | flaky | `[QUARANTINE]` |
| [Transition time from VM deletion](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L182) | 86.8% | — | **Improving** | flaky | |
| [VirtApiRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go) | — | 91.7% | — | flaky | |
| [LowReadyVirtHandlerCount](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go) | — | 92.3% | — | flaky | |
| [LowReadyVirtControllersCount](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go) | — | 92.9% | — | flaky | |

VM dirty rate metrics `[QUARANTINE]` improved from 77.4% (28d) to 85.7% (7d), continuing the positive trend from June 8. Transition time from VM deletion shows zero 7d failures for two consecutive cycles.

### Presubmit lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-monitoring) | 3 | 0 | **Improving** |

The presubmit monitoring lane had only 34 builds in 28d with various low-count test failures. With zero 7d failures, these appear transient.

---

## sig-compute-migrations

**Total 28d failures: 56 (2.2%) | 7d failures: 9 (2.6%)**

### Lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | 32 | 7 | Stable |
| [periodic-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | 24 | 2 | **Improving** |

Both lanes returned testgrid 404, so per-test analysis is unavailable. Failures are tracked via flakefinder only.

---

## Platform-specific

### S390X

**Total 28d failures: 244 (9.4%) | 7d failures: 59 (17.0%)**

The 7d share rose from 9.4% to 17.0% because other SIGs improved while S390X degraded.

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [periodic-S390X](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-test-S390X) | 244 | 59 | Stable |

#### periodic-kubevirt-e2e-test-S390X

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [Eviction API migrate (cluster-wide)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | 45.8% | 25.9% | **Worsening** | flaky |
| [Migration multiple times cloud-init](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L547) | 78.5% | 70.4% | **Worsening** | flaky |
| [MemBalloon period 0](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L160) | 88.7% | — | **Improving** | flaky |
| [Networkpolicy cross-ns reach](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L208) | — | 88.9% | — | flaky |
| [Migration virtqemud restart](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L585) | 93.5% | 88.9% | **Worsening** | flaky |
| [MemBalloon period 12](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L159) | 90.6% | — | **Improving** | flaky |
| [MemBalloon present in domain](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L92) | 90.7% | — | **Improving** | flaky |
| [Config cfgMap+secret layout](https://github.com/kubevirt/kubevirt/blob/main/tests/config_test.go#L451) | 92.5% | 92.6% | Stable | flaky |
| [host-model cpu MigrationSelectorLabel](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go) | 94.0% | 92.6% | Stable | flaky |
| [Networkpolicy ports 80+81](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L242) | 94.4% | — | — | flaky |
| [Networkpolicy port 80/81 deny](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L271) | — | 92.6% | — | flaky |
| [Networkpolicy pinging same ns](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go) | — | 92.6% | — | flaky |
| [LiveMigrateIfPossible](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L342) | — | 92.6% | — | flaky |

**Key observations:**
- The eviction API migration test continues to degrade sharply: 54.7% (June 5) → 50.7% (June 5 28d) → 45.8% (current 28d) → 25.9% (current 7d). This is now failing 3 out of 4 runs.
- Migration multiple times cloud-init also worsening: 78.5% → 70.4%.
- MemBalloon tests appeared in the 28d data but have zero 7d failures — likely a burst that has resolved.
- S390X has a broad spread of ~10 different flaky tests, typical of platform-level instability.

### ARM64

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-kind-1.35-arm64](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-sig-compute-arm64) | 19 | 1 | **Improving** |

arm64 failures remain entirely Pod-level (68.2% → 83.9% success). No individual test failures — the improvement is purely infrastructure.

### Other specialty lanes

| Lane | 28d Failures | 7d Failures | Notes |
|------|-------------|-------------|-------|
| pull-kubevirt-e2e-kind-sriov | 6 | 0 | No test-level data |
| pull-kubevirt-e2e-kind-1.35-vgpu | 3 | 0 | No test-level data |
| pull-kubevirt-check-tests-for-flakes | 2 | 0 | Pod failures only (91.3%) |

---

<a id="quarantine-status-changes"></a>

## Quarantine Status Changes (since 2026-06-08)

Compared against the [prior report (2026-06-08)](flake-overview-2026-06-08.md).

### Quarantine debt resolved

| Test | SIG | Prior Rate | Resolution |
|------|-----|-----------|------------|
| [virtiofs ServiceAccount (quarantine)](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L173) | compute | 0.0% | No longer in data — test may have been removed or is no longer running |

### Stale quarantine candidates

Tests flagged as quarantine candidates in prior reports that are still not quarantined:

| Test | SIG | First Flagged | Current Rate | Issue |
|------|-----|--------------|--------------|-------|
| [HostDevices PCI 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | compute | 2026-06-02 | 0.0% (7d) | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [HostDevices PCI 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | compute | 2026-06-02 | 0.0% (7d) | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | network | 2026-06-02 | 81.5% (7d periodic k8s-1.35) | none |
| [Eviction API migrate (S390X)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | S390X | 2026-06-05 | 25.9% (7d) | none |
| [virtiofs ServiceAccount](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L232) | compute | 2026-06-05 | — (no 7d failures, improving) | none |
| [Volume migrate block→fs](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1274) | storage | 2026-06-05 | — (no 7d failures, improving) | none |
| [VSOCK TLS both sides](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L223) | compute | 2026-06-02 | — (no failures in current data) | none |
| [Transition time from VM deletion](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L182) | monitoring | 2026-06-05 | — (no 7d failures, improving) | none |

All 8 quarantine candidates from the June 8 report remain unaddressed. The HostDevices PCI tests have been at 0% for 3+ weeks with [#17929](https://github.com/kubevirt/kubevirt/issues/17929) open but no quarantine action taken. The S390X eviction test has worsened significantly (40% → 26%).

Positive note: 4 of 8 stale candidates (virtiofs, Volume migrate, VSOCK TLS, Transition time) show zero failures for 2+ consecutive cycles. These can likely be dropped from the candidate list in the next report if the trend holds.

---

<a id="quarantine-candidate-summary"></a>

## Quarantine Candidate Summary

### Quarantine candidates (not yet quarantined)

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority | Issue |
|------|-----|-------|-------------------|----------|---------|------------|----------|-------|
| [HostDevices PCI 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | compute | 1 (k8s-1.34) | 24.3% | 0.0% (worsening) | consecutive | concentrated | **High** | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [HostDevices PCI 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | compute | 1 (k8s-1.34) | 24.3% | 0.0% (worsening) | consecutive | concentrated | **High** | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | network | 5+ lanes | ~83% | 81.5% (stable) | flaky | **dispersed** | **High** | none |
| [Eviction API migrate (S390X)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | S390X | 1 | 45.8% | 25.9% (worsening) | flaky | concentrated | **Medium** | none |
| [Migration multiple times (S390X)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L547) | S390X | 1 | 78.5% | 70.4% (worsening) | flaky | concentrated | **Medium** | none |

Tests downgraded from prior reports (zero failures for 2+ cycles):
- virtiofs ServiceAccount — was Medium, now stable (no 7d failures since June 5)
- Volume migrate block→fs — was Medium, now stable
- VSOCK TLS both sides — was Low, now zero failures in data
- Transition time from VM deletion — was Low, now stable

### Quarantine debt — quarantined tests still severely broken (<50% success rate)

These tests were quarantined but never fixed. They continue to consume CI resources and should be prioritized for repair or permanent removal.

| Test | SIG | Worst 28d Rate | 7d Rate | Lanes | Issue |
|------|-----|---------------|---------|-------|-------|
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | compute | 9.7% (k8s-1.34) | 0.0% (k8s-1.34), 63.0% (k8s-1.35), 53.8% (k8s-1.36) | 3 | [#15907](https://github.com/kubevirt/kubevirt/issues/15907) |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | compute | 9.7% (k8s-1.34) | 0.0% (k8s-1.34), 63.0% (k8s-1.35), 53.8% (k8s-1.36) | 3 | [#15907](https://github.com/kubevirt/kubevirt/issues/15907) |
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | storage | 37.9% (k8s-1.34) | 37.0% (k8s-1.36), 38.5% (k8s-1.34), 48.1% (k8s-1.35) | 3 | none |

**SIG-level quarantine trackers:** [#17720 sig-compute](https://github.com/kubevirt/kubevirt/issues/17720), [#17722 sig-network](https://github.com/kubevirt/kubevirt/issues/17722), [#17723 sig-storage](https://github.com/kubevirt/kubevirt/issues/17723)

---

## Key Takeaways

1. **Overall trajectory is improving.** Daily failure rate dropped from ~107/day (prior 21d) to ~50/day (7d). Total 28-day failures decreased from 2641 (June 8) to 2598.

2. **HostDevices PCI on k8s-1.34 remains at 0% for 3+ weeks** — tracked by [#17929](https://github.com/kubevirt/kubevirt/issues/17929) but still not quarantined despite being flagged since June 2. This remains the top action item.

3. **S390X eviction is accelerating downward** — now at 25.9% (7d), down from 40.0% (June 8 7d) and 45.8% (28d). At this rate, the test will be near-deterministic failure within days. No tracking issue exists.

4. **Multiple previously failing tests have stabilized:** virtiofs ServiceAccount (non-quarantine), VSOCK tests, volume migration, and transition time metrics all show zero failures for 2+ consecutive cycles. The virtiofs quarantine version has disappeared from the data entirely (possibly removed).

5. **passt migration ipv4 remains dispersed across 5+ lanes** at 81–93% success with no improvement trend and no tracking issue. Still the highest-priority network quarantine candidate.

6. **Quarantine debt persists unchanged** — USB passthrough (0–10% on k8s-1.34, [#15907](https://github.com/kubevirt/kubevirt/issues/15907)) and IO error (37–48%, no issue) remain broken. These tests consume CI resources without providing signal.

7. **No quarantine status changes since June 2** — 8 candidates remain stale (4 improving, 4 still problematic). No tests were newly quarantined or removed from quarantine in any report since June 2.

8. **DRA tests from June 8 are gone** — the 11 DRA tests that appeared at 0% with 1–2 runs each are no longer in the data, suggesting the underlying infra issue was resolved or the tests were removed.
