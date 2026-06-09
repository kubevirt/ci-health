# KubeVirt Flake Overview — 2026-06-05

## Summary

| Metric | 28-day | 7-day | Trend |
|--------|--------|-------|-------|
| **Total failures** | 2778 | 436 | |
| **Active lanes** | 42 | 35 | |
| **Daily failure rate** | ~99/day | ~62/day | Improving (-) |
| **Quarantine candidates** | 3 high, 2 medium | | |
| **Infra-unstable lanes** | 12 (Pod failure rate <85%) | | |

**Overall trend: Improving.** The 7-day daily failure rate (~62/day) is lower than the prior 21-day rate (~112/day), indicating a positive trend. However, sig-compute and sig-storage continue to dominate with persistent, long-standing flakes.

**Top 3 worst SIGs (28d):**
1. **sig-compute** — 1696 failures (61.1%) — virtiofs, USB, PCI, VSOCK flakes
2. **sig-storage** — 515 failures (18.5%) — IO error and volume migration flakes
3. **Platform-specific (S390X)** — 256 failures (9.2%) — live migration instability

## Navigation

- [Infrastructure](#infrastructure)
- [sig-compute](#sig-compute)
- [sig-storage](#sig-storage)
- [sig-network](#sig-network)
- [sig-operator](#sig-operator)
- [sig-monitoring](#sig-monitoring)
- [sig-performance](#sig-performance)
- [sig-compute-migrations](#sig-compute-migrations)
- [Platform-specific](#platform-specific)
- [Quarantine Status Changes](#quarantine-status-changes)
- [Quarantine Candidate Summary](#quarantine-candidate-summary)

---

## Infrastructure

Pod-level setup failures indicate infrastructure instability rather than individual test flakes. These affect all tests in a lane when they occur.

| Lane | 28d Pod Rate | 7d Pod Rate | Trend |
|------|-------------|-------------|-------|
| pull-kubevirt-e2e-kind-1.35-sig-compute-arm64 | 67.9% | 74.4% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-storage | 74.3% | 72.1% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-compute | 75.1% | 73.3% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-network | 75.9% | 76.3% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-compute-serial | 76.0% | — | — |
| pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations | 76.2% | 75.2% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-operator | 77.1% | 75.9% | Stable |
| pull-kubevirt-e2e-k8s-1.34-sig-performance | 79.4% | — | — |
| pull-kubevirt-e2e-k8s-1.35-sig-storage | 79.6% | 85.4% | Improving |
| pull-kubevirt-e2e-k8s-1.34-sig-monitoring | 81.8% | 80.0% | Stable |
| pull-kubevirt-e2e-k8s-1.35-sig-network | 83.1% | 89.7% | Improving |
| pull-kubevirt-e2e-k8s-1.35-sig-operator | 83.2% | — | — |
| pull-kubevirt-e2e-k8s-1.35-sig-compute | 85.0% | 89.6% | Improving |
| pull-kubevirt-check-tests-for-flakes | 92.2% | 89.2% | Worsening |
| pull-kubevirt-e2e-k8s-1.34-sig-storage | 94.7% | — | — |

**Key observations:**
- The k8s-1.36 presubmit lanes consistently show Pod success rates in the 72–77% range — roughly 1 in 4 runs fail at the infrastructure level before tests even execute.
- The arm64 lane is the worst at 68–74% Pod success.
- k8s-1.35 lanes generally improved over the last 7 days.

---

## sig-compute

**Total 28d failures: 1696 (61.1%) | 7d failures: 219 (50.2%)**

### Periodic lanes

| Lane | 28d Failures | 28d Share | 7d Failures | 7d Share | Trend |
|------|-------------|-----------|-------------|----------|-------|
| [periodic-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute) | 895 | 32.2% | 56 | 12.8% | Improving |
| [periodic-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-compute) | 366 | 13.2% | 122 | 28.0% | Worsening |
| [periodic-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute) | 189 | 6.8% | 18 | 4.1% | Improving |

#### periodic-kubevirt-e2e-k8s-1.36-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|----------|---------|-------|------------|---------|-------|
| [virtiofs ServiceAccount (quarantine)](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L173) | 0.0% | 0.0% | Stable | 1 lane | consecutive | `[QUARANTINE]` |
| [virtiofs ServiceAccount](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L232) | 10.3% | — | — | 2 lanes (presubmit) | consecutive | |
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | 59.0% | 60.7% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | 59.0% | 60.7% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) | 90.5% | 92.9% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [VSOCK TLS both sides](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L223) | 90.5% | 92.9% | Stable | 1 lane | flaky | |
| [VSOCK no app listener](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L240) | 91.4% | — | — | 1 lane | flaky | |
| [VSOCK migration CID](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L121) | 92.4% | — | — | 1 lane | flaky | |
| [VSOCK non-transitional](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L95) | 92.4% | — | — | 1 lane | flaky | |
| [VSOCK transitional](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L94) | 92.4% | — | — | 1 lane | flaky | |
| [VSOCK without TLS](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L224) | 92.4% | — | — | 1 lane | flaky | |
| [ClusterProfiler pprof](https://github.com/kubevirt/kubevirt/blob/main/tests/infrastructure/cluster-profiler.go#L62) | 94.9% | — | — | 1 lane | flaky | `[QUARANTINE]` |

#### periodic-kubevirt-e2e-k8s-1.34-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|----------|---------|-------|------------|---------|-------|
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | 20.2% | 0.0% | **Worsening** | dispersed (3 lanes) | consecutive (7d) | `[QUARANTINE]` |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | 20.2% | 0.0% | **Worsening** | dispersed (3 lanes) | consecutive (7d) | `[QUARANTINE]` |
| [HostDevices PCI 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | 42.3% | 0.0% | **Worsening** | 1 lane | consecutive (7d) | |
| [HostDevices PCI 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | 42.3% | 0.0% | **Worsening** | 1 lane | consecutive (7d) | |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) | 83.7% | 77.8% | Worsening | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [VMI libvirtd logs](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_lifecycle_test.go#L178) | 94.2% | 92.6% | Stable | 1 lane | flaky | |

#### periodic-kubevirt-e2e-k8s-1.35-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|----------|---------|-------|------------|---------|-------|
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | 50.0% | 66.7% | Improving | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | 50.0% | 66.7% | Improving | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [Guest console log flood](https://github.com/kubevirt/kubevirt/blob/main/tests/guestlog/guestlog.go#L159) | 84.6% | 87.5% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |

### Presubmit lanes

| Lane | 28d Failures | 28d Share | 7d Failures | 7d Share | Trend |
|------|-------------|-----------|-------------|----------|-------|
| [pull-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute) | 98 | 3.5% | 12 | 2.8% | Stable |
| [pull-kind-1.35-arm64](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-sig-compute-arm64) | 22 | 0.8% | 3 | 0.7% | Stable |
| [pull-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute) | 9 | 0.3% | 5 | 1.2% | Worsening |
| [pull-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-compute) | 12 | 0.4% | 2 | 0.5% | Stable |
| [pull-k8s-1.36-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute-serial) | 7 | 0.3% | — | — | — |

#### pull-kubevirt-e2e-k8s-1.36-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|----------|---------|-------|------------|---------|-------|
| [virtiofs ServiceAccount](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L232) | 40.0% | — | — | 2 lanes (periodic+presubmit) | consecutive | |

### SEV lane

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [periodic-kind-1.34-sev](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-kind-1.34-sev) | 8 | 1 | Improving |

| Test | 28d Rate | 7d Rate | Trend | Pattern | Flags |
|------|----------|---------|-------|---------|-------|
| [SEV guest attestation](https://github.com/kubevirt/kubevirt/blob/main/tests/launchsecurity/sev.go#L349) | 90.4% | — | — | flaky | `[QUARANTINE]` |

---

## sig-storage

**Total 28d failures: 515 (18.5%) | 7d failures: 37 (8.5%)**

### Periodic lanes

| Lane | 28d Failures | 28d Share | 7d Failures | 7d Share | Trend |
|------|-------------|-----------|-------------|----------|-------|
| [periodic-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-storage) | 241 | 8.7% | 15 | 3.4% | Improving |
| [periodic-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage) | 106 | 3.8% | 15 | 3.4% | Stable |
| [periodic-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-storage) | 81 | 2.9% | 15 | 3.4% | Stable |

#### Failing tests — periodic storage lanes (all 3 k8s versions)

| Test | Lane | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flags |
|------|------|----------|---------|-------|------------|---------|-------|
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | k8s-1.34 | 37.5% | 48.1% | Worsening | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | k8s-1.35 | 45.2% | 51.9% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | k8s-1.36 | 47.2% | 48.1% | Stable | dispersed (3 lanes) | flaky | `[QUARANTINE]` |
| [Volume migrate block→fs](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1274) | k8s-1.36 | 79.2% | — | — | 2 lanes | flaky | |
| [Volume migrate fs→fs](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1272) | k8s-1.36 | 87.7% | — | — | 2 lanes | flaky | |
| [Volume migrate containerdisk+hotplug](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1131) | k8s-1.36 | 90.6% | — | — | 2 lanes | flaky | |

### Presubmit lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-storage) | 75 | 5 | Improving |
| [pull-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-storage) | 6 | 1 | Stable |
| [pull-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-storage) | 6 | 1 | Stable |

#### pull-kubevirt-e2e-k8s-1.36-sig-storage

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| [Volume migrate block→fs](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1274) | 91.7% | — | — | 2 lanes (periodic+presubmit) | flaky |
| [Volume migrate containerdisk+hotplug](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1131) | 92.1% | — | — | 2 lanes | flaky |
| [Volume migrate fs→fs](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1272) | 92.9% | — | — | 2 lanes | flaky |

---

## sig-network

**Total 28d failures: 84 (3.0%) | 7d failures: 34 (7.8%)**

### Periodic lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [periodic-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network) | 14 | 10 | **Worsening** |
| [periodic-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network) | 8 | 5 | **Worsening** |
| [periodic-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-network) | 5 | 1 | Stable |
| periodic-k8s-1.35-ipv6-sig-network | 12 | 3 | Improving |

#### Failing tests — periodic network lanes

| Test | Lane | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|------|----------|---------|-------|------------|---------|
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | k8s-1.35 | 85.0% | 70.4% | **Worsening** | dispersed (5+ lanes) | flaky |

### Presubmit lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-network) | 16 | 8 | Worsening |
| [pull-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-network) | 14 | 5 | Stable |
| [pull-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-network) | 5 | 5 | **Worsening** |

#### Failing tests — presubmit network lanes

| Test | Lane | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|------|----------|---------|-------|------------|---------|
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | pull-k8s-1.35 | 93.1% | 83.1% | Worsening | dispersed (5+ lanes) | flaky |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | pull-k8s-1.34 | 91.1% | 85.5% | Worsening | dispersed (5+ lanes) | flaky |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | pull-k8s-1.36 | 90.5% | 84.8% | Worsening | dispersed (5+ lanes) | flaky |

---

## sig-operator

**Total 28d failures: 114 (4.1%) | 7d failures: 46 (10.6%)**

### Lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-operator) | 61 | — | — |
| [pull-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-operator) | 45 | 44 | **Worsening** |
| [periodic-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-operator) | 5 | 1 | Stable |
| [periodic-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-operator) | 1 | 1 | Stable |
| [pull-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-operator) | 2 | — | — |

**Note:** sig-operator failures are almost entirely Pod-level infrastructure failures, not test flakes. No individual test failures were recorded — only `.Pod` entries. The pull-k8s-1.36-sig-operator lane has a 77% Pod success rate (28d) and 76% (7d), indicating persistent infra instability.

---

## sig-monitoring

**Total 28d failures: 73 (2.6%) | 7d failures: 9 (2.1%)**

### Lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [periodic-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-monitoring) | 70 | 8 | Improving |
| [pull-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-monitoring) | 3 | 1 | Stable |

#### periodic-kubevirt-e2e-k8s-1.34-sig-monitoring

| Test | 28d Rate | 7d Rate | Trend | Pattern | Flags |
|------|----------|---------|-------|---------|-------|
| [VM dirty rate metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/vm_monitoring.go#L334) | 79.2% | 85.7% | Improving | flaky | `[QUARANTINE]` |
| [Transition time from VM deletion](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L182) | 86.8% | — | — | flaky | |
| [kubevirt_vmi_info guest OS labels](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/vm_monitoring.go#L248) | 94.1% | 92.9% | Stable | flaky | |

#### pull-kubevirt-e2e-k8s-1.34-sig-monitoring

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [Kubevirt alert rules annotations](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go) | 80.0% | 55.6% | **Worsening** | flaky |
| [Virt components metrics](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L119) | 75.0% | — | — | flaky |
| [LowReadyVirtAPICount](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L263) | 90.0% | 77.8% | Worsening | flaky |
| [VirtOperatorRESTErrorsBurst](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/component_monitoring.go#L340) | 88.2% | 85.7% | Stable | flaky |

---

## sig-performance

**Total 28d failures: 40 (1.4%) | 7d failures: 0**

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [periodic-k8s-1.34-sig-performance-kwok-100](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok-100) | 28 | — | — |
| [periodic-k8s-1.34-sig-performance](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance) | 5 | — | — |
| [periodic-k8s-1.34-sig-performance-kwok](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok) | 5 | — | — |
| [pull-k8s-1.34-sig-performance](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-performance) | 2 | — | — |

**Note:** sig-performance had 0 failures in the last 7 days. The 28d failures were concentrated in the kwok-100 lane, likely infrastructure-related (setup failures).

---

## sig-compute-migrations

**Total 28d failures: 57 (2.1%) | 7d failures: 21 (4.8%)**

### Lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | 31 | 10 | Worsening |
| [periodic-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | 26 | 11 | Worsening |

#### periodic-kubevirt-e2e-k8s-1.35-sig-compute-migrations

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [Live Migrations with priority](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/priority.go#L76) | 92.5% | — | — | flaky |
| [Bandwidth limitations migration durations](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L262) | — | 84.6% | New in 7d | flaky |
| [Backup checkpoints after migration](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/backup.go#L1167) | — | 88.5% | New in 7d | flaky |
| [Migration multiple times cloud-init](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L547) | — | 92.3% | New in 7d | flaky |

---

## Platform-specific

### S390X

**Total 28d failures: 256 (9.2%) | 7d failures: 43 (9.9%)**

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [periodic-S390X](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-test-S390X) | 256 | 43 | Stable |

#### periodic-kubevirt-e2e-test-S390X

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| [Eviction API migrate (cluster-wide)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | 54.7% | 48.0% | Worsening | flaky |
| [Migration multiple times cloud-init](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L547) | 81.3% | 80.0% | Stable | flaky |
| [Migration virtqemud restart](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L585) | 87.9% | — | — | flaky |
| [MemBalloon period 0](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L160) | 88.7% | — | — | flaky |
| [Networkpolicy ports 80+81](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L242) | 88.8% | 92.0% | Improving | flaky |
| [MemBalloon period 12](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L159) | 90.6% | — | — | flaky |
| [LiveMigrateIfPossible](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/migration.go#L342) | 90.7% | 88.0% | Stable | flaky |
| [MemBalloon present in domain](https://github.com/kubevirt/kubevirt/blob/main/tests/compute/vmidefaults.go#L92) | 90.7% | — | — | flaky |
| [Networkpolicy cross-namespace reach](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L208) | 94.4% | 88.0% | Worsening | flaky |
| [Networkpolicy port 80 allow, 81 deny](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L271) | 94.4% | — | — | flaky |
| [Networkpolicy cross-namespace ping fail](https://github.com/kubevirt/kubevirt/blob/main/tests/network/networkpolicy.go#L138) | 94.4% | — | — | flaky |
| [Config cfgMap+secret layout](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go) | 94.4% | — | — | flaky |

### ARM64

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| [pull-kind-1.35-arm64](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-sig-compute-arm64) | 22 | 3 | Improving |

**Note:** arm64 failures are entirely Pod-level (68% 28d, 74% 7d success). No individual test failures — purely infrastructure.

---

<a id="quarantine-status-changes"></a>

## Quarantine Status Changes (since 2026-06-02)

Compared against the [prior report (2026-06-02)](flake-overview-2026-06-02.md).

### Stale quarantine candidates

Tests flagged as quarantine candidates in the prior report that are still not quarantined:

| Test | SIG | First Flagged | Current Rate | Issue |
|------|-----|--------------|--------------|-------|
| [HostDevices PCI 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | compute | 2026-06-02 | 0.0% (7d) | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [HostDevices PCI 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | compute | 2026-06-02 | 0.0% (7d) | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | network | 2026-06-02 | 70.4% (7d) | none |
| [VSOCK TLS both sides](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L223) | compute | 2026-06-02 | 92.9% (7d) | none |
| [VSOCK no app listener](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L240) | compute | 2026-06-02 | 91.4% (28d) | none |
| [VSOCK retain CID](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L121) | compute | 2026-06-02 | 92.4% (28d) | none |

All 6 quarantine candidates from the June 2 report remain unaddressed. The HostDevices PCI tests are the most critical — still at 0% success rate with a tracking issue ([#17929](https://github.com/kubevirt/kubevirt/issues/17929)) open but no quarantine action taken.

---

## Quarantine Candidate Summary

### New quarantine candidates (not yet quarantined)

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority | Issue |
|------|-----|-------|-------------------|----------|---------|------------|----------|-------|
| [HostDevices PCI 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L103) | compute | 1 (k8s-1.34) | 42.3% | 0.0% (worsening) | consecutive | concentrated | **High** | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [HostDevices PCI 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_hostdev_test.go#L102) | compute | 1 (k8s-1.34) | 42.3% | 0.0% (worsening) | consecutive | concentrated | **High** | [#17929](https://github.com/kubevirt/kubevirt/issues/17929) |
| [virtiofs ServiceAccount](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L232) | compute | 2 (periodic+presubmit 1.36) | 10.3% / 40.0% | — | consecutive | concentrated | **High** | none |
| [passt migration ipv4](https://github.com/kubevirt/kubevirt/blob/main/tests/network/passt.go#L280) | network | 5+ lanes | 85.0% | 70.4% (worsening) | flaky | **dispersed** | **High** | none |
| [Eviction API migrate (S390X)](https://github.com/kubevirt/kubevirt/blob/main/tests/migration/eviction_strategy.go#L520) | S390X | 1 | 54.7% | 48.0% (worsening) | flaky | concentrated | **Medium** | none |
| [Volume migrate block→fs](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/migration.go#L1274) | storage | 2 (periodic+presubmit 1.36) | 79.2% | — | flaky | concentrated | **Medium** | none |
| [VSOCK TLS both sides](https://github.com/kubevirt/kubevirt/blob/main/tests/vmi_vsock_test.go#L223) | compute | 1 (k8s-1.36) | 90.5% | 92.9% (stable) | flaky | concentrated | Low | none |
| [Transition time from VM deletion](https://github.com/kubevirt/kubevirt/blob/main/tests/monitoring/metrics.go#L182) | monitoring | 1 | 86.8% | — | flaky | concentrated | Low | none |

### Quarantine debt — quarantined tests still severely broken (<50% success rate)

These tests were quarantined but never fixed. They continue to consume CI resources and should be prioritized for repair or permanent removal.

| Test | SIG | Worst 28d Rate | 7d Rate | Lanes | Issue |
|------|-----|---------------|---------|-------|-------|
| [virtiofs ServiceAccount (quarantine)](https://github.com/kubevirt/kubevirt/blob/main/tests/virtiofs/config.go#L173) | compute | 0.0% | 0.0% | 1 (k8s-1.36) | none |
| [USB passthrough 1 device](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L120) | compute | 20.2% (k8s-1.34) | 0.0% (k8s-1.34) | 3 | [#15907](https://github.com/kubevirt/kubevirt/issues/15907) |
| [USB passthrough 2 devices](https://github.com/kubevirt/kubevirt/blob/main/tests/usb/usb.go#L121) | compute | 20.2% (k8s-1.34) | 0.0% (k8s-1.34) | 3 | [#15907](https://github.com/kubevirt/kubevirt/issues/15907) |
| [IO error pause VMI](https://github.com/kubevirt/kubevirt/blob/main/tests/storage/storage.go#L167) | storage | 37.5% (k8s-1.34) | 48.1% (k8s-1.34) | 3 | none |

**SIG-level quarantine trackers:** [#17720 sig-compute](https://github.com/kubevirt/kubevirt/issues/17720), [#17722 sig-network](https://github.com/kubevirt/kubevirt/issues/17722), [#17723 sig-storage](https://github.com/kubevirt/kubevirt/issues/17723)

---

## Key Takeaways

1. **HostDevices PCI passthrough on k8s-1.34 has been at 0% success for 2+ weeks** — tracked by [#17929](https://github.com/kubevirt/kubevirt/issues/17929) but still not quarantined despite being flagged in the June 2 report. This is the highest-priority action item.

2. **virtiofs ServiceAccount test is broken across periodic and presubmit on k8s-1.36** — the quarantine version is at 0%, and the non-quarantine copy is at 10–40%. The non-quarantine copy should also be quarantined. No tracking issue exists.

3. **passt migration connectivity is worsening across all network lanes** — 28d rates of 85–93% have degraded to 70–85% in 7d. This is dispersed across 5+ lanes, making it a strong quarantine candidate. No tracking issue exists.

4. **S390X eviction migration test continues to degrade** — now at 48% 7d success, down from 55% 28d. Platform-specific issue needing investigation. No tracking issue exists.

5. **Infrastructure Pod success rates remain poor on k8s-1.36 presubmits** — ~25% of all runs fail before tests execute. This is not a flake problem but an infrastructure problem.

6. **Quarantine debt is significant** — USB passthrough (0–20%, tracked by [#15907](https://github.com/kubevirt/kubevirt/issues/15907)), virtiofs (0%, no issue), and IO error (37–50%, no issue) tests have been quarantined but remain broken. These should be prioritized for repair.

7. **No quarantine status changes since June 2** — all 6 candidates from the prior report remain unaddressed, and all 4 debt tests are unchanged. No tests were newly quarantined or removed from quarantine.
