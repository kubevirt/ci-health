# KubeVirt Flake Overview — 2026-05-20

## Summary

| Metric | 28-day | 7-day | Trend |
|--------|--------|-------|-------|
| Total failures | 3,071 | 1,219 | Stable (7d is ~40% of 28d, proportionally consistent) |
| Lanes with failures | 39 | 34 | Stable |

**Overall trend**: Stable. The 7-day failure volume is proportionally consistent with the 28-day baseline — no significant improvement or degradation.

**Top 3 worst SIGs by failure count (28d)**:
1. **sig-compute** — 2,007 failures (65.4%) across periodic + presubmit lanes
2. **sig-storage** — 499 failures (16.2%)
3. **Platform-specific (S390X)** — 288 failures (9.4%)

**Quarantine candidates identified**: 3 high priority, 2 medium priority, 1 low priority

**Infra-unstable lanes**: 8 presubmit lanes with Pod-level failure rates between 74–85%

## Navigation

- [Infrastructure](#infrastructure)
- [sig-compute](#sig-compute)
- [sig-storage](#sig-storage)
- [sig-network](#sig-network)
- [sig-monitoring](#sig-monitoring)
- [sig-operator](#sig-operator)
- [sig-performance](#sig-performance)
- [Platform-specific](#platform-specific)
- [Quarantine Candidate Summary](#quarantine-candidate-summary)
- [Quarantine Debt](#quarantine-debt)

---

## Infrastructure

### Presubmit Pod-level failure rates

Pod-level failures indicate infrastructure issues (cluster provisioning failures, resource exhaustion, etc.) rather than test flakiness. These should NOT be addressed by quarantining individual tests.

| Lane | 28d Success Rate | 7d Success Rate | Trend |
|------|-----------------|-----------------|-------|
| pull-kubevirt-e2e-k8s-1.35-sig-storage | 74.0% | — | Worst overall |
| pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations | 74.1% | 78.1% | Improving |
| pull-kubevirt-e2e-k8s-1.35-sig-compute | 78.4% | 84.6% | Improving |
| pull-kubevirt-e2e-k8s-1.35-sig-network | 78.1% | 85.0% | Improving |
| pull-kubevirt-e2e-k8s-1.35-sig-operator | 77.0% | 82.4% | Improving |
| pull-kubevirt-e2e-k8s-1.35-sig-compute-serial | 77.5% | 81.7% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-storage | 78.3% | 78.3% | Stable |
| pull-kubevirt-e2e-kind-1.35-sig-compute-arm64 | 80.5% | 84.4% | Improving |
| pull-kubevirt-e2e-k8s-1.36-sig-compute | 84.7% | 84.7% | Stable |
| pull-kubevirt-e2e-k8s-1.36-sig-compute-serial | 84.7% | 84.7% | Stable |
| pull-kubevirt-e2e-k8s-1.34-sig-performance | 79.2% | 85.6% | Improving |
| pull-kubevirt-e2e-k8s-1.34-sig-storage | 89.7% | 88.6% | Stable |

**Observation**: Most k8s-1.35 presubmit lanes have Pod-level success rates in the 74–85% range, meaning ~15–26% of all runs fail before tests even execute. The 7d trend shows improvement for most lanes, suggesting recent infra stabilization. The k8s-1.36-sig-storage lane remains the worst performer at a stable ~78%.

### Setup failures

| Lane | Issue | 28d Rate | 7d Rate |
|------|-------|----------|---------|
| periodic-kubevirt-e2e-test-S390X | BeforeSuite failures | 0% success (2 runs) | 0% success (2 runs) |

The S390X lane has intermittent complete setup failures. Low sample count (2 BeforeSuite failures across the window) but 100% failure rate when they occur.

---

## sig-compute

**Total failures (28d)**: ~2,007 | **Share**: 65.4%
**Total failures (7d)**: ~843 | **Trend**: Stable

### Periodic lanes

#### periodic-kubevirt-e2e-k8s-1.36-sig-compute

**28d**: 1,199 failures (39.0%) / 110 builds | **7d**: 662 failures (54.3%) / 26 builds | **Trend**: Worsening (share increased from 39% to 54%)

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern | Notes |
|------|----------|---------|-------|------------|---------|-------|
| virtiofs config volumes...ServiceAccount | 0% (105/105) | 0% (23/23) | Stable | 2 lanes (periodic+presubmit 1.36) | Consecutive | **Deterministic failure, not a flake** |
| Prometheus Endpoints [QUARANTINE] | 7.7% (96/104) | 13.0% (20/23) | Stable | 3 lanes (1.34, 1.35, 1.36 periodic) | Flaky | Already quarantined |
| VSOCK Live migration...retain CID | 46.2% (56/104) | — | — | 1 lane | Flaky | 1.36-only |
| VSOCK...TLS on both sides | 58.7% (43/104) | — | — | 1 lane | Flaky | 1.36-only |
| VSOCK...no app listeners | 59.6% (42/104) | — | — | 1 lane | Flaky | 1.36-only |
| USB Passthrough 1 device [QUARANTINE] | 60.0% (42/105) | 47.8% (12/23) | Worsening | 3 lanes (1.34, 1.35, 1.36) | Flaky | Already quarantined |
| USB Passthrough 2 devices [QUARANTINE] | 60.0% (42/105) | 47.8% (12/23) | Worsening | 3 lanes (1.34, 1.35, 1.36) | Flaky | Already quarantined |
| VSOCK...without TLS | 63.5% (38/104) | — | — | 1 lane | Flaky | 1.36-only |
| VSOCK...non-transitional | 66.3% (35/104) | — | — | 1 lane | Flaky | 1.36-only |
| VSOCK...transitional | 66.3% (35/104) | — | — | 1 lane | Flaky | 1.36-only |
| VSOCK...invalid port | 70.2% (31/104) | — | — | 1 lane | Flaky | 1.36-only |
| Guest console log [QUARANTINE] | 87.6% (13/105) | 91.3% (2/23) | Improving | 3 lanes | Flaky | Already quarantined |
| cluster profiler pprof | 88.5% (12/104) | 78.3% (5/23) | Worsening | 5+ lanes | Flaky | Cross-lane |
| CPU pinning...no pinning after pinning | 93.3% (7/105) | — | — | 2 lanes | Flaky | |

#### periodic-kubevirt-e2e-k8s-1.35-sig-compute

**28d**: 274 failures (8.9%) / 110 builds | **7d**: 68 failures (5.6%) / 27 builds | **Trend**: Improving

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| Prometheus Endpoints [QUARANTINE] | 4.6% (103/108) | 7.7% (24/26) | Stable | 3 lanes | Flaky |
| USB Passthrough 1 device [QUARANTINE] | 48.1% (56/108) | 46.2% (14/26) | Stable | 3 lanes | Flaky |
| USB Passthrough 2 devices [QUARANTINE] | 48.1% (56/108) | 46.2% (14/26) | Stable | 3 lanes | Flaky |
| Guest console log [QUARANTINE] | 82.4% (19/108) | 84.6% (4/26) | Stable | 3 lanes | Flaky |
| Force restart VM terminationGracePeriod | 93.5% (7/108) | 92.3% (2/26) | Stable | 2 lanes | Flaky |
| cluster profiler pprof | 94.4% (6/108) | 88.5% (3/26) | Worsening | 5+ lanes | Flaky |

#### periodic-kubevirt-e2e-k8s-1.34-sig-compute

**28d**: 250 failures (8.1%) / 109 builds | **7d**: 62 failures (5.1%) / 25 builds | **Trend**: Stable

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| Prometheus Endpoints [QUARANTINE] | 9.4% (96/106) | 8.7% (21/23) | Stable | 3 lanes | Flaky |
| USB Passthrough 1 device [QUARANTINE] | 54.7% (48/106) | 43.5% (13/23) | Worsening | 3 lanes | Flaky |
| USB Passthrough 2 devices [QUARANTINE] | 54.7% (48/106) | 43.5% (13/23) | Worsening | 3 lanes | Flaky |
| Guest console log [QUARANTINE] | 85.8% (15/106) | 91.3% (2/23) | Improving | 3 lanes | Flaky |
| cluster profiler pprof | 92.5% (8/106) | 87.0% (3/23) | Worsening | 5+ lanes | Flaky |

### Presubmit lanes

#### pull-kubevirt-e2e-k8s-1.36-sig-compute

**28d**: 14 failures (0.5%) / 140 builds | **7d**: 14 failures (1.2%) / 140 builds

| Test | 28d Rate | 7d Rate | Trend | Notes |
|------|----------|---------|-------|-------|
| virtiofs config volumes...ServiceAccount | 0% (92/92) | 0% (92/92) | Stable | Deterministic failure |
| cluster profiler pprof | 87.0% (12/92) | 87.0% (12/92) | Stable | Cross-lane |

#### pull-kubevirt-e2e-k8s-1.35-sig-compute-serial

**28d**: 30 failures (1.0%) / 427 builds | **7d**: 6 failures (0.5%) / 173 builds | **Trend**: Improving

| Test | 28d Rate | 7d Rate | Trend | Notes |
|------|----------|---------|-------|-------|
| cluster profiler pprof | 93.3% (19/282) | 92.3% (9/117) | Stable | Cross-lane |
| CPU pinning...no pinning after pinning | 94.3% (16/282) | 94.0% (7/117) | Stable | |

#### pull-kubevirt-e2e-k8s-1.36-sig-compute-serial

**28d**: 3 failures (0.1%) / 126 builds | **7d**: 3 failures (0.3%) / 126 builds

| Test | 28d Rate | 7d Rate | Trend | Notes |
|------|----------|---------|-------|-------|
| cluster profiler pprof | 87.8% (10/82) | 87.8% (10/82) | Stable | Cross-lane |
| CPU pinning...no pinning after pinning | 92.6% (6/81) | 92.6% (6/81) | Stable | |

### Compute — sig-compute-migrations

#### periodic-kubevirt-e2e-k8s-1.35-sig-compute-migrations

**28d**: 13 failures (0.4%) / 110 builds | **7d**: 3 failures (0.3%) / 27 builds | **Trend**: Stable

| Test | 28d Rate | 7d Rate | Trend |
|------|----------|---------|-------|
| Live Migrations with priority...live-migrate eviction | — | 88.0% (3/25) | — |

### Compute — SEV (kind-1.34-sev)

**28d**: 5 failures (0.2%) / 50 builds | **7d**: 1 failure (0.1%) / 21 builds | **Trend**: Improving

| Test | 28d Rate | 7d Rate | Trend |
|------|----------|---------|-------|
| SEV lifecycle [QUARANTINE] guest attestation | 90.0% (5/50) | — | — |

---

## sig-storage

**Total failures (28d)**: ~499 | **Share**: 16.2%
**Total failures (7d)**: ~290 | **Trend**: Worsening (share increased from 16.2% to 23.8%)

### Periodic lanes

#### periodic-kubevirt-e2e-k8s-1.35-sig-storage

**28d**: 253 failures (8.2%) / 109 builds | **7d**: 196 failures (16.1%) / 25 builds | **Trend**: Worsening (share nearly doubled)

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| error disk [QUARANTINE] pause VMI on IO error | 37.5% (65/104) | 31.8% (15/22) | Worsening | 3 lanes (1.34, 1.35, 1.36) | Flaky |
| cancel migration [QUARANTINE] delete source | 91.9% (7/86) | — | — | 2 lanes (1.35, 1.36) | Flaky |

#### periodic-kubevirt-e2e-k8s-1.36-sig-storage

**28d**: 116 failures (3.8%) / 112 builds | **7d**: 25 failures (2.1%) / 27 builds | **Trend**: Improving

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| error disk [QUARANTINE] pause VMI on IO error | 39.0% (64/105) | 39.1% (14/23) | Stable | 3 lanes | Flaky |
| volume migrate...block to filesystem | 75.2% (26/105) | 78.3% (5/23) | Stable | 2 lanes (periodic+presubmit 1.36) | Flaky |
| volume migrate...fs to fs | 85.7% (15/105) | 82.6% (4/23) | Stable | 2 lanes (periodic+presubmit 1.36) | Flaky |
| cancel migration [QUARANTINE] | 94.3% (5/87) | 91.3% (2/23) | Stable | 2 lanes | Flaky |

#### periodic-kubevirt-e2e-k8s-1.34-sig-storage

**28d**: 78 failures (2.5%) / 108 builds | **7d**: 30 failures (2.5%) / 24 builds | **Trend**: Stable

| Test | 28d Rate | 7d Rate | Trend | Dispersion |
|------|----------|---------|-------|------------|
| error disk [QUARANTINE] pause VMI on IO error | 40.4% (62/104) | 31.8% (15/22) | Worsening | 3 lanes |

### Presubmit lanes

#### pull-kubevirt-e2e-k8s-1.36-sig-storage

**28d**: 37 failures (1.2%) / 141 builds | **7d**: 37 failures (3.0%) / 141 builds | **Trend**: Worsening (all failures are recent)

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| volume migrate...block to fs | 61.3% (31/80) | 61.3% (31/80) | All recent | 2 lanes | Burst/regression |
| volume migrate...fs to fs | 61.3% (31/80) | 61.3% (31/80) | All recent | 2 lanes | Burst/regression |
| volume migrate...containerdisk + hotplug | 91.3% (7/80) | 91.3% (7/80) | All recent | 1 lane | Flaky |

---

## sig-network

**Total failures (28d)**: ~62 | **Share**: 2.0%
**Total failures (7d)**: ~21 | **Trend**: Stable

### Periodic lanes

#### periodic-kubevirt-e2e-k8s-1.35-ipv6-sig-network

**28d**: 26 failures (0.9%) / 111 builds | **7d**: 5 failures (0.4%) / 27 builds | **Trend**: Improving

| Test | 28d Rate | 7d Rate | Trend | Dispersion |
|------|----------|---------|-------|------------|
| passt binding migration IPv6 | 81.8% (20/110) | 80.8% (5/26) | Stable | 1 lane |

#### periodic-kubevirt-e2e-k8s-1.35-sig-network-with-dnc

**28d**: 2 failures (0.1%) / 28 builds | **7d**: — | **Trend**: —

| Test | 28d Rate | Notes |
|------|----------|-------|
| nic-hotplug...multiple interfaces In place | 92.9% (2/28) | DNC-specific |

### Presubmit lanes

#### pull-kubevirt-e2e-k8s-1.34-sig-network

**28d**: 13 failures (0.4%) / 146 builds | **7d**: 6 failures (0.5%) / 75 builds | **Trend**: Stable

| Test | 28d Rate | 7d Rate | Trend |
|------|----------|---------|-------|
| MultiQueue...correct number of queues | 94.2% (7/120) | 93.5% (4/62) | Stable |

---

## sig-monitoring

**Total failures (28d)**: ~99 | **Share**: 3.2%
**Total failures (7d)**: ~26 | **Trend**: Stable

### periodic-kubevirt-e2e-k8s-1.34-sig-monitoring

**28d**: 97 failures (3.2%) / 56 builds | **7d**: 24 failures (2.0%) / 14 builds | **Trend**: Stable

| Test | 28d Rate | 7d Rate | Trend | Dispersion | Pattern |
|------|----------|---------|-------|------------|---------|
| Prometheus scraped metrics [QUARANTINE] VMI namespace | 0% (53/53) | 0% (13/13) | Stable | 1 lane | Consecutive — deterministic |
| VM dirty rate metrics [QUARANTINE] | 69.8% (16/53) | 53.8% (6/13) | Worsening | 1 lane | Flaky |
| Prometheus metrics...transition time VM deletion | 80.4% (10/51) | 61.5% (5/13) | Worsening | 1 lane | Flaky |
| VM metrics guest agent [QUARANTINE] guest OS labels | 94.3% (3/53) | — | — | 1 lane | Flaky |

### pull-kubevirt-e2e-k8s-1.34-sig-monitoring

**28d**: 2 failures (0.1%) / 23 builds | **7d**: 2 failures (0.2%) / 19 builds

Low build count lane (23 builds over 28d). Many tests with exactly 1/17 failure (94.1% success) — likely infra-correlated from a single bad build rather than individual flakes. The Pod-level failure rate (78.3% 28d) confirms infrastructure instability in this lane.

---

## sig-operator

**Total failures (28d)**: ~97 | **Share**: 3.2%
**Total failures (7d)**: ~3 | **Trend**: Improving

### Periodic lanes

| Lane | 28d Failures | 7d Failures | Trend |
|------|-------------|-------------|-------|
| periodic-kubevirt-e2e-k8s-1.36-sig-operator | 47 | 2 | Improving |
| periodic-kubevirt-e2e-k8s-1.34-sig-operator | 46 | — | Improving |
| periodic-kubevirt-e2e-k8s-1.35-sig-operator | 4 | — | Stable |

No individual tests surfaced with high failure rates — failures are distributed. This suggests past infrastructure issues rather than test flakes.

---

## sig-performance

**Total failures (28d)**: ~117 | **Share**: 3.8%
**Total failures (7d)**: ~14 | **Trend**: Improving

### periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok-100

**28d**: 97 failures (3.2%) / 84 builds | **7d**: 7 failures (0.6%) / 21 builds | **Trend**: Improving significantly

No individual test data surfaced — failures appear to be infra-level.

### pull-kubevirt-e2e-k8s-1.34-sig-performance

**28d**: 10 failures (0.3%) / 271 builds | **7d**: 1 failure (0.1%) / 113 builds | **Trend**: Improving

| Test | 28d Rate | 7d Rate | Trend |
|------|----------|---------|-------|
| Density test...VMs with instancetype | 93.7% (12/191) | 88.4% (10/86) | Worsening |
| Density test...create all VMs | 94.8% (10/192) | 88.5% (10/87) | Worsening |

### periodic-kubevirt-e2e-k8s-1.34-sig-performance-kwok

**28d**: 5 failures (0.2%) / 84 builds | **7d**: 4 failures (0.3%) / 21 builds | **Trend**: Worsening

| Test | 28d Rate | 7d Rate | Trend |
|------|----------|---------|-------|
| kwok...create all fake VMs | — | 90.0% (2/20) | — |

---

## Platform-specific

### periodic-kubevirt-e2e-test-S390X

**28d**: 288 failures (9.4%) / 109 builds | **7d**: 28 failures (2.3%) / 28 builds | **Trend**: Improving significantly

| Test | 28d Rate | 7d Rate | Trend | Pattern |
|------|----------|---------|-------|---------|
| BeforeSuite | 0% (2/2) | 0% (2/2) | Stable | Infra setup failure |
| VM Live Migration...virtqemud restart | 73.7% (25/95) | — | — | Flaky |
| VM Live Migration...LiveMigrateIfPossible | 82.1% (17/95) | — | — | Flaky |
| Networkpolicy...ports 80 and 81 | 83.2% (16/95) | 92.0% (2/25) | Improving | Flaky |
| VM Live Migration...multiple times cloud-init | 84.2% (15/95) | 92.0% (2/25) | Improving | Flaky |
| VM Live Migration...persistent domain | 85.3% (14/95) | — | — | Flaky |
| bridge nic-hotunplug Migration based | 89.4% (10/94) | — | — | Flaky |
| Multiple Networkpolicy tests | 88–94% | 92%+ | Improving | Flaky |

S390X shows a clear improvement trend — the 7d failure rate is much lower than the 28d baseline. Many migration and network policy tests that were failing at 73–89% over 28d are now at 92%+ over 7d, suggesting a recent fix landed.

### pull-kubevirt-e2e-kind-1.35-sig-compute-arm64

**28d**: 28 failures (0.9%) / 485 builds | **7d**: 10 failures (0.8%) / 221 builds | **Trend**: Stable

| Test | 28d Rate | 7d Rate | Trend |
|------|----------|---------|-------|
| VMIlifecycle...should log libvirtd logs | 94.7% (19/358) | 93.7% (11/175) | Stable |

---

## Quarantine Candidate Summary

Non-quarantined tests that should be considered for quarantine, ranked by priority:

| Test | SIG | Lanes | Worst Rate (28d) | 7d Trend | Pattern | Dispersion | Priority |
|------|-----|-------|-------------------|----------|---------|------------|----------|
| virtiofs config volumes...ServiceAccount | sig-compute | 2 (periodic+presubmit 1.36) | 0% | Stable at 0% | Consecutive | Concentrated (1.36 only) | **HIGH** — deterministic failure, 100% fail rate |
| VSOCK Live migration...retain CID | sig-compute | 1 | 46.2% | — | Flaky | Concentrated (1.36) | **HIGH** — very low success rate |
| VSOCK...TLS on both sides | sig-compute | 1 | 58.7% | — | Flaky | Concentrated (1.36) | **HIGH** — very low success rate |
| VSOCK...no app listeners | sig-compute | 1 | 59.6% | — | Flaky | Concentrated (1.36) | **MEDIUM** — 1.36-only |
| VSOCK...without TLS | sig-compute | 1 | 63.5% | — | Flaky | Concentrated (1.36) | **MEDIUM** — 1.36-only |
| volume migrate...block to fs | sig-storage | 2 (periodic+presubmit 1.36) | 61.3% (presubmit) | All recent | Burst/regression | Concentrated (1.36) | **MEDIUM** — recent regression, investigate before quarantining |
| volume migrate...fs to fs | sig-storage | 2 (periodic+presubmit 1.36) | 61.3% (presubmit) | All recent | Burst/regression | Concentrated (1.36) | **MEDIUM** — same as above |
| cluster profiler pprof | sig-compute | 5+ lanes | 87.0–94.4% | Worsening | Flaky | Dispersed | **MEDIUM** — cross-lane but moderate rate |
| VSOCK...non-transitional | sig-compute | 1 | 66.3% | — | Flaky | Concentrated (1.36) | **LOW** — 1.36-only, may resolve with VSOCK fix |
| VSOCK...transitional | sig-compute | 1 | 66.3% | — | Flaky | Concentrated (1.36) | **LOW** — same as above |
| VSOCK...invalid port | sig-compute | 1 | 70.2% | — | Flaky | Concentrated (1.36) | **LOW** — same as above |
| Prometheus metrics...transition time | sig-monitoring | 1 | 80.4% | Worsening (61.5% 7d) | Flaky | Concentrated | **LOW** — worsening trend, monitor |
| passt binding migration IPv6 | sig-network | 1 | 81.8% | Stable (80.8%) | Flaky | Concentrated (IPv6) | **LOW** — IPv6-specific |
| Force restart VM terminationGracePeriod | sig-compute | 2 lanes | 93.5% | Stable (92.3%) | Flaky | Concentrated | **LOW** — moderate rate |

**Key observation**: The VSOCK tests all fail exclusively in the k8s-1.36 lanes, suggesting a k8s-1.36-specific issue with VSOCK support rather than individual test problems. A single root-cause investigation could address all 7 VSOCK test failures.

**Key observation**: The volume migration tests (`block to fs` and `fs to fs`) have identical 28d and 7d rates in the presubmit lane, meaning all 31 failures occurred in the last 7 days. This is a **recent regression**, not a long-standing flake. Investigate recent merges affecting volume migration before quarantining.

---

## Quarantine Debt

Already-quarantined tests that remain severely broken (success rate <50%):

| Test | SIG | Worst Rate (28d) | 7d Trend | Lanes |
|------|-----|-------------------|----------|-------|
| Prometheus Endpoints [QUARANTINE] should include correct labels | sig-compute | 4.6% (1.35) | Stable (~8-13%) | 3 periodic lanes (1.34, 1.35, 1.36) |
| Prometheus scraped metrics [QUARANTINE] VMI namespace | sig-monitoring | 0% | Stable at 0% | 1 periodic lane (1.34) |
| error disk [QUARANTINE] pause VMI on IO error | sig-storage | 37.5% (1.35) | Worsening (31.8%) | 3 periodic lanes (1.34, 1.35, 1.36) |
| USB Passthrough 1 device [QUARANTINE] | sig-compute | 43.5% (1.34 7d) | Worsening | 3 periodic lanes (1.34, 1.35, 1.36) |
| USB Passthrough 2 devices [QUARANTINE] | sig-compute | 43.5% (1.34 7d) | Worsening | 3 periodic lanes (1.34, 1.35, 1.36) |

These tests have been quarantined but never fixed. The Prometheus and IO error tests in particular are at near-zero or sub-40% success rates across all k8s versions, indicating persistent unfixed bugs. The USB passthrough tests are worsening over the 7d window.
