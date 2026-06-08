# KubeVirt Flake Overview — 2026-06-08

## Summary

| Metric | 28-day | 7-day | Trend |
|--------|--------|-------|-------|
| **Total failures** | 2641 | 414 | |
| **Active lanes** | 45 | 34 | |
| **Daily failure rate** | ~94/day | ~59/day | Improving (-) |
| **Quarantine candidates** | 3 high, 2 medium | | |
| **Infra-unstable lanes** | ~12 (Pod failure rate <85%) | 7 | |

**Overall trend: Improving.** The 7-day daily failure rate (~59/day) is significantly lower than the prior 21-day rate (~106/day). Total 28-day failures dropped from 2778 (June 5) to 2641. Several previously failing tests (VSOCK, ClusterProfiler, virtiofs non-quarantine, volume migration) show zero failures in the 7-day window, suggesting recent fixes or stabilization.

**Top 3 worst SIGs (28d):**
1. **sig-compute** — HostDevices PCI deterministic failure, USB quarantine debt, DRA test failures (new)
2. **sig-storage** — IO error quarantine debt persists across all k8s versions
3. **Platform-specific (S390X)** — eviction migration continues to degrade

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
| pull-kubevirt-e2e-k8s-1.36-sig-storage | 74.3% | 75.0% | Stable |
| pull-kubevirt-e2e-kind-1.35-sig-compute-arm64 | 67.9% | 76.1% | **Improving** |
| pull-kubevirt-e2e-k8s-1.36-sig-compute | 75.1% | 76.6% | Stable |
| pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations | 76.2% | 76.9% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-compute-serial | 76.0% | 77.4% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-network | 75.9% | 78.3% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-operator | 77.1% | 78.7% | Stable |
| pull-kubevirt-check-tests-for-flakes | 92.2% | 90.6% | Stable |
| pull-kubevirt-e2e-k8s-1.35-sig-storage | 79.6% | 94.0% | **Improving** |

**Key observations:**
- The k8s-1.36 presubmit lanes remain at ~75–79% Pod success — roughly 1 in 4 runs fail at the infrastructure level before tests even execute.
- The arm64 lane improved significantly from 68% to 76%.
- k8s-1.35-sig-storage improved dramatically (80% → 94%).
- Most lanes remain stable, indicating persistent but not worsening infra issues.

---

## sig-compute

**Total 7d failures: ~211 (51.0%)**

### Periodic lanes

| Lane | 7d Failures | 7d Share | Trend vs June 5 |
|------|-------------|----------|-----------------|
| [periodic-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-compute) | 125 | 30.2% | **Worsening** |
| [periodic-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute) | 38 | 9.2% | **Improving** |
| [periodic-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute) | 28 | 6.8% | Stable |

#### periodic-kubevirt-e2e-k8s-1.34-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|----------|---------|-------|------------|---------|-------|
| [HostDevices PCI 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | 31.7% | 0.0% | **Worsening** | 1 lane | consecutive | |
| [HostDevices PCI 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | 31.7% | 0.0% | **Worsening** | 1 lane | consecutive | |
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | 12.5% | 0.0% | **Worsening** | dispersed (3 lanes) | consecutive | `[QUARANTINE]` |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | 12.5% | 0.0% | **Worsening** | dispersed (3 lanes) | consecutive | `[QUARANTINE]` |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) | ~84% | 70.4% | **Worsening** | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [VMI libvirtd logs](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_lifecycle_test.go#L178) | ~94% | 92.6% | Stable | 1 lane | flaky | |

#### periodic-kubevirt-e2e-k8s-1.36-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|----------|---------|-------|------------|---------|-------|
| [virtiofs ServiceAccount (quarantine)](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L173) | 0.0% | — | — | 1 lane | consecutive | `[QUARANTINE]` |
| [virtiofs ServiceAccount](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L232) | 23.4% | — | **Improving** | 2 lanes | consecutive | |
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | ~59% | 59.3% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | ~59% | 59.3% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) | ~91% | 88.9% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |

Several tests that appeared in the June 5 report (VSOCK tests, ClusterProfiler pprof) show zero 7d failures, indicating recent improvement.

#### periodic-kubevirt-e2e-k8s-1.35-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|----------|---------|-------|------------|---------|-------|
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | ~50% | 53.8% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | ~50% | 53.8% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) | ~85% | 84.6% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |

### Presubmit lanes

| Lane | 7d Failures | 7d Share | Trend |
|------|-------------|----------|-------|
| [pull-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute) | 9 | 2.2% | Stable |
| [pull-kind-1.35-arm64](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-sig-compute-arm64) | 3 | 0.7% | Stable |
| [pull-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute) | 3 | 0.7% | Stable |
| [pull-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-compute) | 2 | 0.5% | Stable |

#### pull-kubevirt-e2e-k8s-1.34-sig-compute

| Test | 7d Rate | Trend | Dispersion | Pattern |
|------|---------|-------|------------|---------|
| [VMI libvirtd logs](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_lifecycle_test.go#L178) | 94.8% | Stable | 1 lane | flaky |

### DRA tests (new)

DRA (Dynamic Resource Allocation) tests are new in this reporting period. They appear exclusively in presubmit lanes with minimal sample size.

| Lane | Tests | 7d Runs | 7d Rate | Assessment |
|------|-------|---------|---------|------------|
| [pull-kubevirt-check-tests-for-flakes](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-check-tests-for-flakes) | 11 DRA tests | 2 each | 0% | infra-correlated |
| [pull-k8s-1.36-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute-serial) | 11 DRA tests | 1 each | 0% | infra-correlated |

All 11 DRA tests failed in every run across both lanes, with only 1–2 runs per test. The tests cover ResourceClaim management, CEL-based device selection, and matchAttribute-based allocation for vfio-gpu devices. Since all tests failed together in the same builds, this is likely a lane/configuration issue (infra-correlated), not individual test problems. The test source files were not found on the `main` branch, suggesting these tests are being introduced via PRs. Related issues: [#16481](https://github.com/kubevirt/kubevirt/issues/16481) (create e2e test for DRA), [#17674](https://github.com/kubevirt/kubevirt/issues/17674) (virt-launcher missing capabilities for DRA VFIO passthrough).

<details>
<summary>DRA test list (11 tests)</summary>

1. `[sig-compute]DRA create 8 VMIs and pre-created ResourceClaims individually`
2. `[sig-compute]DRA create a VMI with a matchattribute based ResourceClaimTemplate should allocate vfio-gpu devices sharing vendorID via matchAttribute`
3. `[sig-compute]DRA create a VMI with a strict CEL based ResourceClaimTemplate should allocate a vfio-gpu device matching the CEL selector`
4. `[sig-compute]DRA create a VMI with a strict CEL based ResourceClaimTemplate should allocate a vfio-gpu device matching the index`
5. `[sig-compute]DRA create a VMI with multiple device requests in one ResourceClaim should allocate three vfio-gpu devices`
6. `[sig-compute]DRA create eight VMIs backed by the ResourceClaimTemplate`
7. `[sig-compute]DRA failing scenarios create a VMI with a matchattribute based ResourceClaimTemplate should fail when the matchattribute has different values`
8. `[sig-compute]DRA failing scenarios should fail when a ResourceClaimTemplate requests more than one device per request`
9. `[sig-compute]DRA failing scenarios should reject a VMI when the host device references an unknown resourceClaim`
10. `[sig-compute]DRA failing scenarios should remain unschedulable when the CEL selector matches no device`
11. `[sig-compute]DRA failing scenarios should remain unschedulable when the CEL selector uses invalid attribute access`

</details>

**Not quarantine candidates** — insufficient sample size and infra-correlated. Monitor for persistence.

### SEV lane

| Lane | 7d Failures | Trend |
|------|-------------|-------|
| [periodic-kind-1.34-sev](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-kind-1.34-sev) | 1 | Improving |

No individual test failures in 7d data. SEV guest attestation `[QUARANTINE]` test appears stable.

---

## sig-storage

**Total 7d failures: 54 (13.0%)**

### Periodic lanes

| Lane | 7d Failures | 7d Share | Trend vs June 5 |
|------|-------------|----------|-----------------|
| [periodic-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage) | 20 | 4.8% | Stable |
| [periodic-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-storage) | 15 | 3.6% | Stable |
| [periodic-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-storage) | 14 | 3.4% | Stable |

#### Failing tests — periodic storage lanes (all 3 k8s versions)

| Test | Lane | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|------|----------|---------|-------|------------|---------|-------|
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | k8s-1.36 | 44.3% | 33.3% | **Worsening** | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | k8s-1.34 | 35.6% | 44.4% | Improving | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | k8s-1.35 | ~45% | 53.6% | Improving | dispersed (3 lanes) | flaky | `[QUARANTINE]` |

Volume migration tests (block→fs, fs→fs, containerdisk+hotplug) that appeared in the June 5 report show zero 7d failures, indicating improvement.

### Presubmit lanes

| Lane | 7d Failures | Trend |
|------|-------------|-------|
| [pull-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-storage) | 2 | Improving |
| [pull-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-storage) | 2 | Stable |
| [pull-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-storage) | 1 | Stable |

Presubmit storage failures are entirely Pod-level infrastructure issues, not test flakes.

---

## sig-network

**Total 7d failures: ~43 (10.4%)**

### Periodic lanes

| Lane | 7d Failures | Trend vs June 5 |
|------|-------------|-----------------|
| [periodic-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network) | 9 | Stable |
| [periodic-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network) | 6 | Stable |
| [periodic-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-network) | 3 | Stable |

#### Failing tests — periodic network lanes

| Test | Lane | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|------|----------|---------|-------|------------|---------|
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | k8s-1.35 | ~85% | 75.0% | **Worsening** | dispersed (5+ lanes) | flaky |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | k8s-1.34 | ~89% | 88.9% | Stable | dispersed (5+ lanes) | flaky |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | k8s-1.36 | ~93% | 92.6% | Stable | dispersed (5+ lanes) | flaky |
| [bridge nic-hotplug Migration based](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L215) | k8s-1.36 | — | 92.6% | New in 7d | 1 lane | flaky |

### Presubmit lanes

| Lane | 7d Failures | Trend |
|------|-------------|-------|
| [pull-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-network) | 9 | Stable |
| [pull-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-network) | 7 | Worsening |
| [pull-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-network) | 5 | Stable |

#### Failing tests — presubmit network lanes

| Test | Lane | 7d Rate | Dispersion | Pattern |
|------|------|---------|------------|---------|
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | pull-k8s-1.35 | 83.8% | dispersed (5+ lanes) | flaky |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | pull-k8s-1.36 | 84.2% | dispersed (5+ lanes) | flaky |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | pull-k8s-1.34 | 90.3% | dispersed (5+ lanes) | flaky |

### Other network lanes

| Lane | 7d Failures | Notes |
|------|-------------|-------|
| periodic-k8s-1.36-sig-network-with-dnc | 1 | [bridge nic-hotplug In place](https://github.com/kubevirt/kubevirt/blob/main/tests/network/hotplug_bridge.go#L215) at 80% |
| periodic-k8s-1.36-ipv6-sig-network | 1 | — |
| periodic-k8s-1.35-sig-network-with-dnc | 1 | — |
| pull-k8s-1.35-ipv6-sig-network | 1 | — |
| pull-k8s-1.36-ipv6-sig-network | 1 | — |

---

## sig-operator

**Total 7d failures: 45 (10.9%)**

| Lane | 7d Failures | Trend |
|------|-------------|-------|
| [pull-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-operator) | 44 | **Worsening** |
| [periodic-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-operator) | 1 | Stable |

sig-operator failures are almost entirely Pod-level infrastructure failures, not test flakes. The pull-k8s-1.36-sig-operator lane has a 78.7% Pod success rate (7d), meaning ~21% of runs fail before tests execute.

---

## sig-monitoring

**Total 7d failures: ~6 (1.4%)**

| Lane | 7d Failures | Trend |
|------|-------------|-------|
| [periodic-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-monitoring) | 5 | **Improving** |

#### periodic-kubevirt-e2e-k8s-1.34-sig-monitoring

| Test | 7d Rate | Trend vs June 5 | Pattern |
|------|---------|-----------------|---------|
| [VirtApiRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L329) | 91.7% | Stable | flaky |
| [LowReadyVirtControllersCount](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L254) | 92.3% | Stable | flaky |
| [VirtControllerDown](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L149) | 92.9% | — | flaky |
| [kubevirt_vmi_info guest OS](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/vm_monitoring.go#L248) | 92.9% | Stable | flaky |

VM dirty rate metrics `[QUARANTINE]` showed zero failures in the 7d window, improving from 79.2% in June 5. The quarantine appears to be working — the test has stabilized.

---

## sig-compute-migrations

**Total 7d failures: 17 (4.1%)**

| Lane | 7d Failures | Trend |
|------|-------------|-------|
| [pull-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | 11 | Stable |
| [periodic-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | 6 | Improving |

#### pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations

| Test | 7d Rate | Trend | Pattern |
|------|---------|-------|---------|
| [Migration multiple times cloud-init](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L547) | 93.8% | Stable | flaky |

Bandwidth limitations and Backup checkpoints tests from the June 5 7d window show zero failures, suggesting those were transient.

---

## Platform-specific

### S390X

**Total 7d failures: 37 (8.9%)**

| Lane | 7d Failures | Trend |
|------|-------------|-------|
| [periodic-S390X](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-test-S390X) | 37 | Stable |

#### periodic-kubevirt-e2e-test-S390X

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [Eviction API migrate (cluster-wide)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | 50.7% | 40.0% | **Worsening** | flaky |
| [Migration multiple times cloud-init](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L547) | ~81% | 72.0% | **Worsening** | flaky |
| [LiveMigrateIfPossible](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L342) | ~91% | 92.0% | Stable | flaky |
| [Config cfgMap+secret layout](https://github.com/kubevirt/kubevirt/blob/main/tests/config_test.go#L451) | ~94% | 92.0% | Stable | flaky |
| [Networkpolicy ports 80+81](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L242) | ~89% | 92.0% | Improving | flaky |
| [Networkpolicy cross-namespace reach](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L208) | ~94% | 92.0% | Stable | flaky |

### ARM64

| Lane | 7d Failures | Trend |
|------|-------------|-------|
| [pull-kind-1.35-arm64](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-sig-compute-arm64) | 3 | Improving |

arm64 failures remain entirely Pod-level (76.1% success). No individual test failures.

---

## Quarantine Status Changes (since 2026-06-05)

Compared against the [prior report (2026-06-05)](flake-overview-2026-06-05.md).

### Stale quarantine candidates

Tests flagged as quarantine candidates in prior reports that are still not quarantined:

| Test | SIG | First Flagged | Current Rate | Issue |
|------|-----|--------------|--------------|-------|
| [HostDevices PCI 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | compute | 2026-06-02 | 0.0% (7d) | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [HostDevices PCI 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | compute | 2026-06-02 | 0.0% (7d) | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | network | 2026-06-02 | 75.0% (7d periodic k8s-1.35) | none |
| [Eviction API migrate (S390X)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | S390X | 2026-06-05 | 40.0% (7d) | none |
| [virtiofs ServiceAccount](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L232) | compute | 2026-06-05 | — (no 7d failures) | none |
| [Volume migrate block→fs](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1274) | storage | 2026-06-05 | — (no 7d failures) | none |
| [VSOCK TLS both sides](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L223) | compute | 2026-06-02 | — (no 7d failures) | none |
| [Transition time from VM deletion](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L182) | monitoring | 2026-06-05 | — (no 7d failures) | none |

All 8 quarantine candidates from the June 5 report remain unaddressed. The HostDevices PCI tests are the most critical — now at 0% for 3+ weeks with [#17929](https://github.com/kubevirt/kubevirt/issues/17929) open but no quarantine action taken.

Positive note: 4 of 8 stale candidates (virtiofs non-quarantine, VSOCK TLS, Volume migrate, Transition time) show zero failures in the 7-day window, suggesting recent improvement. These may no longer warrant quarantine but should be monitored for another cycle.

---

## Quarantine Candidate Summary

### Quarantine candidates (not yet quarantined)

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority | Issue |
|------|-----|-------|-------------------|----------|---------|------------|----------|-------|
| [HostDevices PCI 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | compute | 1 (k8s-1.34) | 31.7% | 0.0% (worsening) | consecutive | concentrated | **High** | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [HostDevices PCI 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | compute | 1 (k8s-1.34) | 31.7% | 0.0% (worsening) | consecutive | concentrated | **High** | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | network | 5+ lanes | ~85% | 75.0% (worsening) | flaky | **dispersed** | **High** | none |
| [Eviction API migrate (S390X)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | S390X | 1 | 50.7% | 40.0% (worsening) | flaky | concentrated | **Medium** | none |
| [virtiofs ServiceAccount](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L232) | compute | 2 (periodic+presubmit 1.36) | 23.4% | — (no 7d failures, improving) | consecutive | concentrated | **Medium** → monitor | none |

virtiofs non-quarantine, Volume migrate block→fs, VSOCK TLS, and Transition time from VM deletion were candidates in June 5 but show zero 7d failures. These should be monitored for another cycle before downgrading.

### Quarantine debt — quarantined tests still severely broken (<50% success rate)

These tests were quarantined but never fixed. They continue to consume CI resources and should be prioritized for repair or permanent removal.

| Test | SIG | Worst 28d Rate | 7d Rate | Lanes | Issue |
|------|-----|---------------|---------|-------|-------|
| [virtiofs ServiceAccount (quarantine)](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L173) | compute | 0.0% | — | 1 (k8s-1.36) | none |
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | compute | 12.5% (k8s-1.34) | 0.0% (k8s-1.34), 53.8% (k8s-1.35), 59.3% (k8s-1.36) | 3 | [#15907](https://github.com/kubevirt/kubevirt/issues/15907) |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | compute | 12.5% (k8s-1.34) | 0.0% (k8s-1.34), 53.8% (k8s-1.35), 59.3% (k8s-1.36) | 3 | [#15907](https://github.com/kubevirt/kubevirt/issues/15907) |
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | storage | 35.6% (k8s-1.34) | 33.3% (k8s-1.36), 44.4% (k8s-1.34), 53.6% (k8s-1.35) | 3 | none |

**SIG-level quarantine trackers:** [#17720 sig-compute](https://github.com/kubevirt/kubevirt/issues/17720), [#17722 sig-network](https://github.com/kubevirt/kubevirt/issues/17722), [#17723 sig-storage](https://github.com/kubevirt/kubevirt/issues/17723)

---

## Key Takeaways

1. **Overall trajectory is improving.** Daily failure rate dropped from ~106/day (prior 21d) to ~59/day (7d). Total 28-day failures decreased from 2778 (June 5) to 2641.

2. **HostDevices PCI on k8s-1.34 remains at 0% for 3+ weeks** — tracked by [#17929](https://github.com/kubevirt/kubevirt/issues/17929) but still not quarantined. This has been the top action item since June 2.

3. **Several previously failing tests have stabilized:** VSOCK tests, ClusterProfiler pprof, virtiofs non-quarantine, volume migration, and Transition time metrics all show zero 7d failures. These may have been recently fixed or may be in a quiet period.

4. **passt migration ipv4 continues to worsen on k8s-1.35** — now at 75.0% success (7d periodic), down from ~85% (28d). Still dispersed across 5+ lanes. No tracking issue exists.

5. **S390X eviction continues to degrade** — now at 40.0% (7d), down from 50.7% (28d). No tracking issue exists.

6. **DRA tests are new but likely infra-correlated** — 11 tests at 0% with only 1–2 runs each, all failing together. Not quarantine candidates until sample size increases. Related: [#16481](https://github.com/kubevirt/kubevirt/issues/16481), [#17674](https://github.com/kubevirt/kubevirt/issues/17674).

7. **Quarantine debt persists unchanged** — USB passthrough (0–13%, [#15907](https://github.com/kubevirt/kubevirt/issues/15907)), virtiofs (0%), and IO error (34–54%) remain broken. These tests consume CI resources without providing signal.

8. **No quarantine status changes since June 2** — 8 candidates remain stale (4 showing 7d improvement), 4 debt tests unchanged. No tests were newly quarantined or removed from quarantine.
