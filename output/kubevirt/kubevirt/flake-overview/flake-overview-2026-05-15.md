# KubeVirt CI Flake Overview -- 2026-05-15

## Summary

| | 28-day | 7-day |
|---|--------|-------|
| **Period** | 2026-04-17 to 2026-05-15 | 2026-05-08 to 2026-05-15 |
| **Total failures** | 2,221 | 559 |
| **Lanes with failures** | 35 | 29 |

**Overall trend: Stable** -- 7d failures are 25.2% of 28d (expected 25.0% for proportional).

**Top 3 worst SIGs by failure count (28d):**

1. **sig-compute** -- 1,144 failures (51.5%)
2. **Platform-specific** (S390X/arm64/SEV) -- 375 failures (16.9%)
3. **sig-storage** -- 260 failures (11.7%)

**Quarantine candidates identified:** 1 high, 3 medium, 1 low priority

**Infra-unstable lanes:** 1 completely broken (kwok-100), 9 presubmit lanes with elevated .Pod failure rates (71--85%)

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

### Completely broken lane

| Lane | 28d Failures | 7d Failures | Status |
|------|-------------|-------------|--------|
| [periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok-100](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok-100) | 104 (BeforeSuite 0%, AfterSuite 0%) | 28 (BeforeSuite 0%, AfterSuite 0%) | Broken -- no tests run |

The kwok-100 lane has been unable to set up the test environment for the entire 28-day window. All 71 BeforeSuite attempts in 28d and all 17 in 7d failed. This is an infrastructure/configuration issue, not a test problem.

### Presubmit .Pod failure rates

Pod-level failures indicate infrastructure problems (provisioning, scheduling, timeouts) rather than test code issues. These are **infra-correlated** and should not be addressed by quarantining individual tests.

| Lane | 28d Rate | 7d Rate | Trend |
|------|----------|---------|-------|
| [pull-kubevirt-e2e-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-storage) | 74.1% (100/386) | 71.0% (58/200) | Worsening |
| [pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | 75.1% (98/393) | 73.5% (54/204) | Worsening |
| [pull-kubevirt-e2e-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-network) | 76.9% (86/372) | 73.9% (49/188) | Worsening |
| [pull-kubevirt-e2e-k8s-1.35-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-serial) | 77.1% (90/393) | 74.4% (50/195) | Worsening |
| [pull-kubevirt-e2e-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-operator) | 77.3% (84/370) | 74.9% (47/187) | Worsening |
| [pull-kubevirt-e2e-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute) | 78.6% (81/379) | 76.6% (45/192) | Slightly worsening |
| [pull-kubevirt-e2e-k8s-1.34-sig-performance](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-performance) | 79.4% (52/253) | 80.3% (25/127) | Stable |
| [pull-kubevirt-e2e-kind-1.35-sig-compute-arm64](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-kind-1.35-sig-compute-arm64) | 79.7% (87/429) | 78.2% (50/229) | Stable |
| [pull-kubevirt-e2e-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-storage) | 91.1% (11/124) | 85.1% (10/67) | Worsening |

**Trend:** Most k8s-1.35 presubmit lanes are worsening (2--4pp drop over 7d). Approximately 1 in 4 presubmit runs fails at the Pod level before any tests execute.

---

## sig-compute

**28d total: 1,144 failures (51.5%) | 7d total: 298 failures (53.3%) | Trend: Stable**

### Lanes

| Lane | Type | 28d Fail | 28d Share | 28d Builds | 7d Fail | 7d Share | Trend |
|------|------|----------|-----------|------------|---------|----------|-------|
| [periodic-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-compute) | periodic | 558 | 25.12% | 91 | 133 | 23.79% | Stable |
| [periodic-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute) | periodic | 258 | 11.62% | 111 | 79 | 14.13% | Stable |
| [periodic-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-compute) | periodic | 236 | 10.63% | 110 | 63 | 11.27% | Stable |
| [pull-k8s-1.35-sig-compute-serial](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-serial) | presubmit | 32 | 1.44% | 396 | 9 | 1.61% | Stable |
| [pull-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | presubmit | 25 | 1.13% | 397 | 5 | 0.89% | Stable |
| [periodic-k8s-1.35-sig-compute-migrations](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-compute-migrations) | periodic | 14 | 0.63% | 110 | 5 | 0.89% | Stable |
| [pull-k8s-1.35-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-compute) | presubmit | 12 | 0.54% | 381 | 2 | 0.36% | Stable |
| [pull-k8s-1.34-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-compute) | presubmit | 8 | 0.36% | 112 | 1 | 0.18% | Stable |
| [pull-k8s-1.36-sig-compute](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.36-sig-compute) | presubmit | 1 | 0.05% | 38 | 1 | 0.18% | Stable |

### Failing tests -- periodic-kubevirt-e2e-k8s-1.36-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flag |
|------|----------|---------|-------|------------|---------|------|
| virtiofs ServiceAccount | 0.0% (0/87) | 0.0% (0/24) | Stable | Concentrated (2 lanes, k8s-1.36 only) | Consecutive -- 100% failure | |
| Prometheus [test_id:4145] | 5.8% (5/86) | 4.2% (1/24) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| VSOCK Live migration retain CID | 34.9% (30/86) | 66.7% (16/24) | Improving | Concentrated (1 lane) | Flaky | |
| VSOCK TLS on both sides | 50.0% (43/86) | 70.8% (17/24) | Improving | Concentrated (1 lane) | Flaky | |
| VSOCK no app listens | 51.2% (44/86) | 62.5% (15/24) | Improving | Concentrated (1 lane) | Flaky | |
| VSOCK without TLS | 55.8% (48/86) | 66.7% (16/24) | Improving | Concentrated (1 lane) | Flaky | |
| VSOCK virtio non-transitional | 59.3% (51/86) | 66.7% (16/24) | Stable | Concentrated (1 lane) | Flaky | |
| VSOCK virtio transitional | 59.3% (51/86) | 66.7% (16/24) | Stable | Concentrated (1 lane) | Flaky | |
| VSOCK invalid port | 64.0% (55/86) | 79.2% (19/24) | Improving | Concentrated (1 lane) | Flaky | |
| USB passthrough 1 device | 64.4% (56/87) | 66.7% (16/24) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| USB passthrough 2 devices | 64.4% (56/87) | 66.7% (16/24) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| Guest console log | 85.1% (74/87) | 87.5% (21/24) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| ClusterProfiler pprof | 88.4% (76/86) | 70.8% (17/24) | **Worsening** | Dispersed (4 lanes) | Flaky | |
| CPU pinning [829] no-pin after pin | 92.0% (80/87) | -- | -- | Dispersed (3 lanes) | Flaky | |
| CPU pinning [832] pin after no-pin | 94.3% (82/87) | 92.6% (25/27, k8s-1.35) | Stable | Dispersed (4 lanes) | Flaky | |

### Failing tests -- periodic-kubevirt-e2e-k8s-1.35-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flag |
|------|----------|---------|-------|------------|---------|------|
| Prometheus [test_id:4145] | 17.4% (19/109) | 3.7% (1/27) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| USB passthrough 1 device | 49.5% (54/109) | 44.4% (12/27) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| USB passthrough 2 devices | 49.5% (54/109) | 44.4% (12/27) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| Guest console log | 81.7% (89/109) | 74.1% (20/27) | Slightly worsening | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| VM force restart [3007] | 93.6% (102/109) | 88.9% (24/27) | Slightly worsening | Dispersed (3 lanes) | Flaky | |
| ClusterProfiler pprof | 94.5% (103/109) | 92.6% (25/27) | Stable | Dispersed (4 lanes) | Flaky | |
| CPU pinning [832] | 94.5% (103/109) | 92.6% (25/27) | Stable | Dispersed (4 lanes) | Flaky | |

### Failing tests -- periodic-kubevirt-e2e-k8s-1.34-sig-compute

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flag |
|------|----------|---------|-------|------------|---------|------|
| Prometheus [test_id:4145] | 22.2% (24/108) | 4.0% (1/25) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| USB passthrough 1 device | 55.6% (60/108) | 52.0% (13/25) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| USB passthrough 2 devices | 55.6% (60/108) | 52.0% (13/25) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| Guest console log | 84.3% (91/108) | 84.0% (21/25) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| ClusterProfiler pprof | 92.6% (100/108) | 84.0% (21/25) | **Worsening** | Dispersed (4 lanes) | Flaky | |
| CPU pinning [832] | 93.5% (101/108) | 88.0% (22/25) | Slightly worsening | Dispersed (4 lanes) | Flaky | |
| CPU pinning [829] | 94.4% (102/108) | -- | -- | Dispersed (3 lanes) | Flaky | |

### Failing tests -- presubmit sig-compute lanes

| Test | Lane | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|------|----------|---------|-------|------------|---------|
| virtiofs ServiceAccount | pull-k8s-1.36 | 0.0% (0/26) | 0.0% (0/25) | Stable | Concentrated (k8s-1.36) | Consecutive |
| ClusterProfiler pprof | pull-k8s-1.36 | 84.6% (22/26) | 84.0% (21/25) | Stable | Dispersed (4 lanes) | Flaky |
| CPU pinning [829] | pull-k8s-1.35-serial | 94.7% (249/263) | -- | -- | Dispersed (3 lanes) | Flaky |
| CPU pinning [832] | pull-k8s-1.35-serial | 94.7% (249/263) | -- | -- | Dispersed (4 lanes) | Flaky |
| Post-copy migration liveness probe | pull-k8s-1.35-migrations | 94.2% (243/258) | 94.5% (121/128) | Stable | Concentrated (1 lane) | Flaky |

---

## sig-storage

**28d total: 260 failures (11.7%) | 7d total: 70 failures (12.5%) | Trend: Stable**

### Lanes

| Lane | Type | 28d Fail | 28d Share | 28d Builds | 7d Fail | 7d Share | Trend |
|------|------|----------|-----------|------------|---------|----------|-------|
| [periodic-k8s-1.36-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-storage) | periodic | 94 | 4.23% | 94 | 23 | 4.11% | Stable |
| [periodic-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-storage) | periodic | 77 | 3.47% | 109 | 29 | 5.19% | Slightly worsening |
| [periodic-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-storage) | periodic | 75 | 3.38% | 112 | 16 | 2.86% | Stable |
| [pull-k8s-1.34-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-storage) | presubmit | 8 | 0.36% | 128 | 1 | 0.18% | Stable |
| [pull-k8s-1.35-sig-storage](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-storage) | presubmit | 6 | 0.27% | 391 | 1 | 0.18% | Stable |

### Failing tests -- periodic sig-storage lanes

| Test | Lane | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flag |
|------|------|----------|---------|-------|------------|---------|------|
| should pause VMI on IO error | k8s-1.36 | 39.8% (35/88) | 58.3% (14/24) | Improving | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| should pause VMI on IO error | k8s-1.35 | 42.1% (45/107) | 48.0% (12/25) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| should pause VMI on IO error | k8s-1.34 | 42.9% (45/105) | 41.7% (10/24) | Stable | Dispersed (3 lanes) | Flaky | [QUARANTINE] |
| Hotplug volumes block-to-fs | k8s-1.36 | 75.0% (66/88) | 70.8% (17/24) | Stable | Concentrated (1 lane) | Flaky | |
| Hotplug volumes fs-to-fs | k8s-1.36 | 85.2% (75/88) | 70.8% (17/24) | **Worsening** | Concentrated (1 lane) | Flaky | |
| Live Migration cancel source | k8s-1.35 | 90.9% (80/88) | 88.0% (22/25) | Stable | Concentrated (2 lanes) | Flaky | [QUARANTINE] |
| Live Migration cancel source | k8s-1.36 | 94.3% (66/70) | 91.7% (22/24) | Stable | Concentrated (2 lanes) | Flaky | [QUARANTINE] |

---

## sig-network

**28d total: 56 failures (2.5%) | 7d total: 9 failures (1.6%) | Trend: Slightly improving**

### Lanes

| Lane | Type | 28d Fail | 28d Share | 28d Builds | 7d Fail | 7d Share | Trend |
|------|------|----------|-----------|------------|---------|----------|-------|
| [periodic-k8s-1.35-ipv6-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-ipv6-sig-network) | periodic | 26 | 1.17% | 111 | 4 | 0.72% | Slightly improving |
| [pull-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-network) | presubmit | 9 | 0.41% | 375 | 1 | 0.18% | Stable |
| [pull-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-network) | presubmit | 8 | 0.36% | 119 | 1 | 0.18% | Stable |
| [periodic-k8s-1.34-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-network) | periodic | 6 | 0.27% | 112 | 2 | 0.36% | Stable |
| [periodic-k8s-1.36-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-network) | periodic | 3 | 0.14% | 93 | -- | -- | Improving |
| [periodic-k8s-1.35-sig-network](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network) | periodic | 2 | 0.09% | 111 | 1 | 0.18% | Stable |
| [periodic-k8s-1.35-sig-network-with-dnc](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-network-with-dnc) | periodic | 2 | 0.09% | 24 | -- | -- | Improving |

### Failing tests

| Test | Lane | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|------|----------|---------|-------|------------|---------|
| passt migration connectivity [IPv6] | k8s-1.35-ipv6 | 82.0% (91/111) | 88.9% (24/27) | Slightly improving | Concentrated (1 lane) | Flaky |
| MultiQueue report queues [4599] | pull-k8s-1.35 | 94.7% (230/243) | -- | -- | Concentrated (1 lane) | Flaky |
| nic-hotplug multiple interfaces In place | k8s-1.35-dnc | 91.7% (22/24) | -- | -- | Concentrated (1 lane) | Flaky |

---

## sig-operator

**28d total: 142 failures (6.4%) | 7d total: 46 failures (8.2%) | Trend: Slightly worsening (driven by presubmit .Pod failures)**

### Lanes

| Lane | Type | 28d Fail | 28d Share | 28d Builds | 7d Fail | 7d Share | Trend |
|------|------|----------|-----------|------------|---------|----------|-------|
| [periodic-k8s-1.34-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-operator) | periodic | 47 | 2.12% | 112 | -- | -- | Improving |
| [pull-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.35-sig-operator) | presubmit | 46 | 2.07% | 373 | 45 | 8.05% | **Worsening** |
| [periodic-k8s-1.36-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.36-sig-operator) | periodic | 45 | 2.03% | 93 | 1 | 0.18% | Improving |
| [periodic-k8s-1.35-sig-operator](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.35-sig-operator) | periodic | 4 | 0.18% | 111 | -- | -- | Stable |

The pull-k8s-1.35-sig-operator lane's worsening trend is entirely driven by .Pod infrastructure failures (74.9% success rate in 7d, down from 77.3% in 28d). No individual operator tests are failing -- the lane itself is unstable.

---

## sig-monitoring

**28d total: 100 failures (4.5%) | 7d total: 22 failures (3.9%) | Trend: Stable**

### Lanes

| Lane | Type | 28d Fail | 28d Share | 28d Builds | 7d Fail | 7d Share | Trend |
|------|------|----------|-----------|------------|---------|----------|-------|
| [periodic-k8s-1.34-sig-monitoring](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-monitoring) | periodic | 100 | 4.50% | 55 | 22 | 3.94% | Stable |

### Failing tests

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Flag |
|------|----------|---------|-------|------------|---------|------|
| Prometheus scraped metrics [4135] | 0.0% (0/53) | 0.0% (0/13) | Stable | Concentrated (1 lane) | Consecutive -- 100% failure | [QUARANTINE] |
| VM dirty rate metrics | 73.6% (39/53) | 84.6% (11/13) | Slightly improving | Concentrated (1 lane) | Flaky | [QUARANTINE] |
| VM deletion transition time metric | 88.2% (45/51) | -- | -- | Concentrated (1 lane) | Flaky | |
| guest agent OS labels [11267] | 92.5% (49/53) | 92.3% (12/13) | Stable | Concentrated (1 lane) | Flaky | [QUARANTINE] |
| VirtAPIDown alert | 92.9% (26/28) | -- | -- | Concentrated (1 lane) | Flaky | |
| KubeVirtDeprecatedAPIRequested | 94.3% (50/53) | -- | -- | Concentrated (1 lane) | Flaky | [QUARANTINE] |

---

## sig-performance

**28d total: 140 failures (6.3%) | 7d total: 35 failures (6.3%) | Trend: Stable**

### Lanes

| Lane | Type | 28d Fail | 28d Share | 28d Builds | 7d Fail | 7d Share | Trend |
|------|------|----------|-----------|------------|---------|----------|-------|
| [periodic-k8s-1.34-kwok-100](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok-100) | periodic | 104 | 4.68% | 83 | 28 | 5.01% | Broken (infra) |
| [pull-k8s-1.34-sig-performance](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-performance) | presubmit | 12 | 0.54% | 253 | 1 | 0.18% | Improving |
| [pull-k8s-1.34-sig-performance-kwok](https://testgrid.k8s.io/kubevirt-presubmits#pull-kubevirt-e2e-k8s-1.34-sig-performance-kwok) | presubmit | 11 | 0.50% | 1 | -- | -- | -- |
| [periodic-k8s-1.34-sig-performance](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance) | periodic | 8 | 0.36% | 84 | 2 | 0.36% | Stable |
| [periodic-k8s-1.34-sig-performance-kwok](https://testgrid.k8s.io/kubevirt-periodics#periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok) | periodic | 5 | 0.23% | 84 | 4 | 0.72% | Slightly worsening |

### Failing tests (excluding kwok-100 infra failures)

| Test | Lane | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|------|----------|---------|-------|------------|---------|
| Density 100 VMs instancetype | pull-k8s-1.34-perf | 93.3% (167/179) | 90.6% (77/85) | Slightly worsening | Concentrated (1 lane) | Flaky |
| Density 100 VMIs | periodic-k8s-1.34-perf | 92.8% (77/83) | -- | -- | Concentrated (1 lane) | Flaky |
| kwok fake VMs | periodic-k8s-1.34-kwok | -- | 89.5% (17/19) | -- | Concentrated (1 lane) | Flaky |

---

## Platform-specific

**28d total: 375 failures (16.9%) | 7d total: 79 failures (14.1%) | Trend: Slightly improving**

### S390X -- periodic-kubevirt-e2e-test-S390X

**28d: 343 failures (15.44%) | 7d: 73 failures (13.06%) | Builds: 104 (28d) / 27 (7d) | Trend: Slightly improving**

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| Live Migration virtqemud restart [4746] | 67.4% (58/86) | 63.0% (17/27) | Stable | Flaky |
| Live Migration eviction [2353] | 77.9% (67/86) | 81.5% (22/27) | Stable | Flaky |
| Networkpolicy default-deny [1512] clientVMI from serverVMI | 77.9% (67/86) | 88.9% (24/27) | Slightly improving | Flaky |
| Networkpolicy alt-ns [1514] fail pinging | 79.1% (68/86) | 81.5% (22/27) | Stable | Flaky |
| Live Migration cloud-init [1783] | 80.2% (69/86) | 85.2% (23/27) | Stable | Flaky |
| Networkpolicy [2774] allow port 80+81 | 81.4% (70/86) | 70.4% (19/27) | **Worsening** | Flaky |
| Live Migration persistent domain [6972] | 82.6% (71/86) | 92.6% (25/27) | Improving | Flaky |
| Networkpolicy [2775] allow 80 deny 81 | 82.6% (71/86) | 88.9% (24/27) | Stable | Flaky |
| Networkpolicy default-deny [369] | 82.6% (71/86) | 88.9% (24/27) | Stable | Flaky |
| Networkpolicy alt-ns [1515] serverVMI from alt | 82.6% (71/86) | 88.9% (24/27) | Stable | Flaky |
| Networkpolicy default-deny [1511] serverVMI from clientVMI | 83.7% (72/86) | 92.6% (25/27) | Improving | Flaky |
| Networkpolicy default-ns [1515] serverVMI from clientVMI | 84.9% (73/86) | -- | -- | Flaky |
| Networkpolicy alt-ns [1517] clientVMI from alt | 87.2% (75/86) | 88.9% (24/27) | Stable | Flaky |
| nic-hotunplug Migration based | 88.2% (75/85) | -- | -- | Flaky |
| nic-hotplug can be hotplugged Migration based | 89.4% (76/85) | -- | -- | Flaky |
| nic-hotplug can migrate hotplugged Migration based | 89.4% (76/85) | -- | -- | Flaky |
| nic-hotplug connectivity Migration based | 89.4% (76/85) | -- | -- | Flaky |
| nic-hotplug multiple Migration based | 89.4% (76/85) | -- | -- | Flaky |
| Networkpolicy same-ns [1513] | 89.5% (77/86) | 88.9% (24/27) | Stable | Flaky |
| Live Migration paused vmi [6973] | 91.9% (79/86) | 85.2% (23/27) | Slightly worsening | Flaky |
| MTU verification IPv4 | 93.0% (80/86) | 92.6% (25/27) | Stable | Flaky |

**Analysis:** The S390X lane has widespread test instability with 22 failing tests. Many failures cluster around Live Migration and Networkpolicy tests. The large number of Networkpolicy tests failing together at similar rates suggests **infra-correlated** failures on this platform rather than individual test flakes. The Live Migration tests may be affected by S390X-specific virtualization behavior. Overall the lane is slightly improving in 7d.

### arm64 -- pull-kubevirt-e2e-kind-1.35-sig-compute-arm64

**28d: 28 failures (1.26%) | 7d: 4 failures (0.72%) | Builds: 433 (28d) / 233 (7d) | Trend: Stable**

Only .Pod infrastructure failures (79.7% 28d, 78.2% 7d). No individual test failures.

### SEV -- periodic-kubevirt-e2e-kind-1.34-sev

**28d: 4 failures (0.18%) | 7d: 2 failures (0.36%) | Builds: 38 (28d) / 20 (7d) | Trend: Stable**

| Test | 28d Rate | 7d Rate | Trend | Flag |
|------|----------|---------|-------|------|
| SEV guest attestation | 89.5% (34/38) | 90.0% (18/20) | Stable | [QUARANTINE] |

---

## Quarantine Candidate Summary

### New quarantine candidates (not yet quarantined)

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority |
|------|-----|-------|-------------------|----------|---------|------------|----------|
| ClusterProfiler pprof subresource access | sig-compute | 4 (k8s-1.36 periodic, k8s-1.35 periodic, k8s-1.34 periodic, pull-k8s-1.36) | 84.6% | **Worsening** (88.4%->70.8% on k8s-1.36, 92.6%->84% on k8s-1.34) | Flaky | Dispersed | **HIGH** |
| virtiofs ServiceAccount | sig-compute | 2 (k8s-1.36 periodic, pull-k8s-1.36) | 0.0% | Stable (0% both windows) | Consecutive | Concentrated (k8s-1.36 only) | **MEDIUM** |
| Hotplug volumes block-to-filesystem | sig-storage | 1 (k8s-1.36 periodic) | 75.0% | Stable (75%->70.8%) | Flaky | Concentrated | **MEDIUM** |
| VSOCK suite (7 tests) | sig-compute | 1 (k8s-1.36 periodic) | 34.9--64.0% | Improving (all tests improved 5--30pp) | Flaky | Concentrated | **MEDIUM** |
| CPU pinning [829]/[832] | sig-compute | 4 (3 periodic + 1 presubmit) | 92.0% | Slightly worsening (93.5%->88% on k8s-1.34) | Flaky | Dispersed | **LOW** |

### Quarantine debt (quarantined but still severely broken, success rate <50%)

| Test | SIG | Lanes | Worst Rate (28d) | 7d Rate | Status |
|------|-----|-------|-------------------|---------|--------|
| Prometheus [test_id:4145] | sig-compute | 3 | 5.8% | 3.7--4.2% | Stable -- never fixed |
| Prometheus scraped metrics [test_id:4135] | sig-monitoring | 1 | 0.0% | 0.0% | Stable -- never fixed |
| USB passthrough 1 device | sig-compute | 3 | 49.5% | 44.4--66.7% | Stable -- never fixed |
| USB passthrough 2 devices | sig-compute | 3 | 49.5% | 44.4--66.7% | Stable -- never fixed |
| should pause VMI on IO error | sig-storage | 3 | 39.8% | 41.7--58.3% | Stable -- never fixed |
