



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-03-29 (1x / 0.72%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-03-29 10:49:43 &#43;0000 UTC_: <code>10:53:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17227/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2038206823877578752#1:build-log.txt%3A309)
<details>
<summary>all...</summary>

* _2026-03-29 10:49:43 &#43;0000 UTC_: <code>10:53:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17227/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2038206823877578752#1:build-log.txt%3A309)

</details>

<hr/>
</details>

### 2026-03-26 (97x / 69.78%)


#### external (85x / 87.63%)

<details>
<summary> bazel remote cache blob fetch failure (69x / 71.13%) </summary>

<hr/>

**14x**: _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:53:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037088891659358208#1:build-log.txt%3A518)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:34 &#43;0000 UTC_: <code>08:55:31: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-windows2016/2037088890090688512#1:build-log.txt%3A489)

* _2026-03-26 08:47:27 &#43;0000 UTC_: <code>08:56:00: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-network/2037088890619170816#1:build-log.txt%3A516)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:53:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037088891659358208#1:build-log.txt%3A518)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:52:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037088891520946176#1:build-log.txt%3A515)

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:52:38: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037088891369951232#1:build-log.txt%3A524)

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:53:12: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2037088890333958144#1:build-log.txt%3A516)

* _2026-03-26 07:56:52 &#43;0000 UTC_: <code>08:01:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037076170310160384#1:build-log.txt%3A511)

* _2026-03-26 07:47:39 &#43;0000 UTC_: <code>07:53:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037073833353023488#1:build-log.txt%3A523)

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:51:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-network/2037073834191884288#1:build-log.txt%3A524)

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:35: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-windows2016/2037073830886772736#1:build-log.txt%3A549)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:51:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-network/2037073831499141120#1:build-log.txt%3A540)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037073835894771712#1:build-log.txt%3A511)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037073831637553152#1:build-log.txt%3A529)

* _2026-03-26 06:47:31 &#43;0000 UTC_: <code>06:51:22: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2037058711255519232#1:build-log.txt%3A424)

</details>

<hr/>

**6x**: _2026-03-26 07:47:48 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037073832539328512#1:build-log.txt%3A502)
<details>
<summary>all...</summary>

* _2026-03-26 07:47:48 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037073832539328512#1:build-log.txt%3A502)

* _2026-03-26 07:47:47 &#43;0000 UTC_: <code>07:53:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037073831079710720#1:build-log.txt%3A595)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2037073831012601856#1:build-log.txt%3A540)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037073831281037312#1:build-log.txt%3A530)

* _2026-03-26 05:47:21 &#43;0000 UTC_: <code>05:51:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17276/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037043575102902272#1:build-log.txt%3A550)

* _2026-03-26 04:47:33 &#43;0000 UTC_: <code>04:53:53: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037028524841242624#1:build-log.txt%3A499)

</details>

<hr/>

**27x**: _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036967849838252032#1:build-log.txt%3A446)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:52:26: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037088894213689344#1:build-log.txt%3A518)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:55:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037088892464664576#1:build-log.txt%3A517)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-operator/2037088891055378432#1:build-log.txt%3A482)

* _2026-03-26 07:47:40 &#43;0000 UTC_: <code>07:51:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-operator/2037073831419449344#1:build-log.txt%3A529)

* _2026-03-26 07:47:39 &#43;0000 UTC_: <code>07:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037073831914377216#1:build-log.txt%3A525)

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037073836737826816#1:build-log.txt%3A529)

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-network/2037073831209734144#1:build-log.txt%3A540)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-compute/2037073831356534784#1:build-log.txt%3A460)

* _2026-03-26 07:47:36 &#43;0000 UTC_: <code>07:54:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037073835026550784#1:build-log.txt%3A532)

* _2026-03-26 02:47:25 &#43;0000 UTC_: <code>02:52:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036998292239290368#1:build-log.txt%3A481)

* _2026-03-26 00:46:37 &#43;0000 UTC_: <code>00:55:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036967851490807808#1:build-log.txt%3A510)

* _2026-03-26 00:46:32 &#43;0000 UTC_: <code>00:54:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-network/2036967848978419712#1:build-log.txt%3A451)

* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036967847292309504#1:build-log.txt%3A508)

* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036967848152141824#1:build-log.txt%3A478)

* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036967849838252032#1:build-log.txt%3A446)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:52:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036967845736222720#1:build-log.txt%3A484)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:52:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036967850643558400#1:build-log.txt%3A505)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:51:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036967846004658176#1:build-log.txt%3A506)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:51:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-network/2036967846151458816#1:build-log.txt%3A521)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:55:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036967845870440448#1:build-log.txt%3A504)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:52:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036967845358735360#1:build-log.txt%3A588)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:53:15: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-network/2036967845614587904#1:build-log.txt%3A520)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:54:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036967845245489152#1:build-log.txt%3A517)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:51:42: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036967846302453760#1:build-log.txt%3A509)

* _2026-03-26 00:46:25 &#43;0000 UTC_: <code>00:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-monitoring/2036967845480370176#1:build-log.txt%3A508)

* _2026-03-26 00:46:25 &#43;0000 UTC_: <code>00:52:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036967846478614528#1:build-log.txt%3A510)

* _2026-03-26 00:46:23 &#43;0000 UTC_: <code>00:54:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-windows2016/2036967844863807488#1:build-log.txt%3A520)

</details>

<hr/>

**12x**: _2026-03-26 05:04:28 &#43;0000 UTC_: <code>05:08:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037032781342576640#1:build-log.txt%3A534)
<details>
<summary>all...</summary>

* _2026-03-26 05:04:37 &#43;0000 UTC_: <code>05:09:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-network/2037032781271273472#1:build-log.txt%3A572)

* _2026-03-26 05:04:29 &#43;0000 UTC_: <code>05:08:54: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037032781178998784#1:build-log.txt%3A529)

* _2026-03-26 05:04:29 &#43;0000 UTC_: <code>05:08:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037032781422268416#1:build-log.txt%3A541)

* _2026-03-26 05:04:28 &#43;0000 UTC_: <code>05:08:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037032781342576640#1:build-log.txt%3A534)

* _2026-03-26 05:04:26 &#43;0000 UTC_: <code>05:10:23: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037032781506154496#1:build-log.txt%3A522)

* _2026-03-26 05:04:26 &#43;0000 UTC_: <code>05:10:15: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037032781090918400#1:build-log.txt%3A632)

* _2026-03-26 01:49:19 &#43;0000 UTC_: <code>01:59:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036983575923920896#1:build-log.txt%3A480)

* _2026-03-26 01:49:19 &#43;0000 UTC_: <code>01:59:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-windows2016/2036983572635586560#1:build-log.txt%3A561)

* _2026-03-26 01:49:16 &#43;0000 UTC_: <code>01:59:08: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036983573491224576#1:build-log.txt%3A496)

* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:53:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036983573973569536#1:build-log.txt%3A531)

* _2026-03-26 01:48:56 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-network/2036983573851934720#1:build-log.txt%3A538)

* _2026-03-26 01:48:56 &#43;0000 UTC_: <code>01:54:56: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-network/2036983573382172672#1:build-log.txt%3A538)

</details>

<hr/>

**1x**: _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036983575089254400#1:build-log.txt%3A505)
<details>
<summary>all...</summary>

* _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036983575089254400#1:build-log.txt%3A505)

</details>

<hr/>

**9x**: _2026-03-26 01:48:57 &#43;0000 UTC_: <code>01:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036983577610031104#1:build-log.txt%3A481)
<details>
<summary>all...</summary>

* _2026-03-26 01:49:28 &#43;0000 UTC_: <code>01:59:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036983573042434048#1:build-log.txt%3A534)

* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036983574225227776#1:build-log.txt%3A534)

* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036983579300335616#1:build-log.txt%3A541)

* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036983573143097344#1:build-log.txt%3A615)

* _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:19: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036983573713522688#1:build-log.txt%3A535)

* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:55:04: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036983578432114688#1:build-log.txt%3A527)

* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:56:07: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-network/2036983576750198784#1:build-log.txt%3A549)

* _2026-03-26 01:48:57 &#43;0000 UTC_: <code>01:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036983577610031104#1:build-log.txt%3A481)

* _2026-03-26 01:48:57 &#43;0000 UTC_: <code>01:55:06: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036983573587693568#1:build-log.txt%3A540)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (7x / 7.22%) </summary>

<hr/>

**5x**: _2026-03-26 08:52:44 &#43;0000 UTC_: <code>08:57:27: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 33 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037090118291951616#1:build-log.txt%3A439)
<details>
<summary>all...</summary>

* _2026-03-26 08:52:44 &#43;0000 UTC_: <code>08:57:27: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 33 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037090118291951616#1:build-log.txt%3A439)

* _2026-03-26 08:52:43 &#43;0000 UTC_: <code>08:57:26: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-launcher/libvirt-hook-client/BUILD.bazel:10:10: GoCompilePkg cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client.a failed: Exec failed due to IOException: 351 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-network/2037090113296535552#1:build-log.txt%3A760)

* _2026-03-26 08:52:30 &#43;0000 UTC_: <code>08:58:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tools/csv-generator/BUILD.bazel:16:10 Validating nogo output for //tools/csv-generator:csv-generator failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037090114504495104#1:build-log.txt%3A490)

* _2026-03-26 08:52:28 &#43;0000 UTC_: <code>08:57:24: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 73 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-network/2037090114290585600#1:build-log.txt%3A725)

* _2026-03-26 08:52:20 &#43;0000 UTC_: <code>08:57:23: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportserver/BUILD.bazel:20:10: GoLink cmd/virt-exportserver/virt-exportserver_/virt-exportserver failed: Exec failed due to IOException: 161 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037090113539805184#1:build-log.txt%3A479)

</details>

<hr/>

**1x**: _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:58:56: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/virt-handler/vsock/system/BUILD.bazel:3:11 Validating nogo output for //pkg/virt-handler/vsock/system:go_default_library failed: Exec failed due to IOException: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037088895878828032#1:build-log.txt%3A883)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:58:56: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/virt-handler/vsock/system/BUILD.bazel:3:11 Validating nogo output for //pkg/virt-handler/vsock/system:go_default_library failed: Exec failed due to IOException: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037088895878828032#1:build-log.txt%3A883)

</details>

<hr/>

**1x**: _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:55:51: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/pierrec/lz4/v4/internal/xxh32/BUILD.bazel:3:11: Validating nogo output for //vendor/github.com/pierrec/lz4/v4/internal/xxh32:go_default_library failed: Exec failed due to IOException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-network/2037088893345468416#1:build-log.txt%3A467)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:55:51: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/pierrec/lz4/v4/internal/xxh32/BUILD.bazel:3:11: Validating nogo output for //vendor/github.com/pierrec/lz4/v4/internal/xxh32:go_default_library failed: Exec failed due to IOException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-network/2037088893345468416#1:build-log.txt%3A467)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache connection timeout (5x / 5.15%) </summary>

<hr/>

**5x**: _2026-03-26 08:52:28 &#43;0000 UTC_: <code>15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037090115150417920#1:build-log.txt%3A1004)
<details>
<summary>all...</summary>

* _2026-03-26 08:53:02 &#43;0000 UTC_: <code>12:20:45: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037090114382860288#1:build-log.txt%3A1122)

* _2026-03-26 08:52:28 &#43;0000 UTC_: <code>15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037090115150417920#1:build-log.txt%3A1004)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>09:02:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037088890476564480#1:build-log.txt%3A2362)

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>09:00:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037088895111270400#1:build-log.txt%3A1119)

* _2026-03-26 08:47:24 &#43;0000 UTC_: <code>08:59:51: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037088890786942976#1:build-log.txt%3A971)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (from secondary snippet) (1x / 1.03%) </summary>

<hr/>

**1x**: _2026-03-26 08:34:18 &#43;0000 UTC_: <code>08:57:20: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-k8s-1.34-sig-network/2037085594298880000#1:build-log.txt%3A4832)
<details>
<summary>all...</summary>

* _2026-03-26 08:34:18 &#43;0000 UTC_: <code>08:57:20: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-k8s-1.34-sig-network/2037085594298880000#1:build-log.txt%3A4832)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache download timeout (1x / 1.03%) </summary>

<hr/>

**1x**: _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:54:51: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportproxy/BUILD.bazel:27:10: GoLink cmd/virt-exportproxy/virt-exportproxy_/virt-exportproxy failed: Exec failed due to IOException: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/cas/1c6882fc7400bd46e51aa4130768869fb62c82d129bba7bf916ca68ab6a2ac1c&#39; timed out. Received 21508688 of 24245460 bytes.</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-network/2037088891227344896#1:build-log.txt%3A466)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:54:51: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportproxy/BUILD.bazel:27:10: GoLink cmd/virt-exportproxy/virt-exportproxy_/virt-exportproxy failed: Exec failed due to IOException: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/cas/1c6882fc7400bd46e51aa4130768869fb62c82d129bba7bf916ca68ab6a2ac1c&#39; timed out. Received 21508688 of 24245460 bytes.</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-network/2037088891227344896#1:build-log.txt%3A466)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 1.03%) </summary>

<hr/>

**1x**: _2026-03-26 03:48:42 &#43;0000 UTC_: <code>04:06:36: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037013641730985984#1:build-log.txt%3A4356)
<details>
<summary>all...</summary>

* _2026-03-26 03:48:42 &#43;0000 UTC_: <code>04:06:36: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037013641730985984#1:build-log.txt%3A4356)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 1.03%) </summary>

<hr/>

**1x**: _2026-03-26 06:47:21 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2037058711163244544#1:build-log.txt%3A5065)
<details>
<summary>all...</summary>

* _2026-03-26 06:47:21 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2037058711163244544#1:build-log.txt%3A5065)

</details>

<hr/>
</details>

#### internal (4x / 4.12%)

<details>
<summary> kind cluster creation failure (3x / 3.09%) </summary>

<hr/>

**3x**: _2026-03-26 09:22:19 &#43;0000 UTC_: <code>09:29:01: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037097667556806656#1:build-log.txt%3A816)
<details>
<summary>all...</summary>

* _2026-03-26 09:22:19 &#43;0000 UTC_: <code>09:29:01: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037097667556806656#1:build-log.txt%3A816)

* _2026-03-26 08:47:18 &#43;0000 UTC_: <code>08:54:22: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-kind-1.34-sev/2037088890237489152#1:build-log.txt%3A709)

* _2026-03-26 01:49:22 &#43;0000 UTC_: <code>01:55:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2036983572786581504#1:build-log.txt%3A789)

</details>

<hr/>
</details>
<details>
<summary> control plane startup failure (1x / 1.03%) </summary>

<hr/>

**1x**: _2026-03-26 03:48:17 &#43;0000 UTC_: <code>04:03:20: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037013641584185344#1:build-log.txt%3A2850)
<details>
<summary>all...</summary>

* _2026-03-26 03:48:17 &#43;0000 UTC_: <code>04:03:20: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037013641584185344#1:build-log.txt%3A2850)

</details>

<hr/>
</details>

#### needs-investigation (6x / 6.19%)

<details>
<summary> no error snippets (6x / 6.19%) </summary>

<hr/>

**1x**: _2026-03-26 08:52:44 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2037090112856133632)
<details>
<summary>all...</summary>

* _2026-03-26 08:52:44 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2037090112856133632)

</details>

<hr/>

**1x**: _2026-03-26 08:50:21 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037089601964740608)
<details>
<summary>all...</summary>

* _2026-03-26 08:50:21 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037089601964740608)

</details>

<hr/>

**1x**: _2026-03-26 07:47:29 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-kind-1.34-sev/2037073830949687296)
<details>
<summary>all...</summary>

* _2026-03-26 07:47:29 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-kind-1.34-sev/2037073830949687296)

</details>

<hr/>

**1x**: _2026-03-26 05:04:21 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037032780927340544)
<details>
<summary>all...</summary>

* _2026-03-26 05:04:21 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037032780927340544)

</details>

<hr/>

**1x**: _2026-03-26 02:47:17 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036998292079906816)
<details>
<summary>all...</summary>

* _2026-03-26 02:47:17 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036998292079906816)

</details>

<hr/>

**1x**: _2026-03-26 00:46:17 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-kind-1.34-sev/2036967845056745472)
<details>
<summary>all...</summary>

* _2026-03-26 00:46:17 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-kind-1.34-sev/2036967845056745472)

</details>

<hr/>
</details>

#### pr-build (2x / 2.06%)

<details>
<summary> panic detected in test output (2x / 2.06%) </summary>

<hr/>

**2x**: _2026-03-26 07:35:38 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037070842281594880#1:build-log.txt%3A5081)
<details>
<summary>all...</summary>

* _2026-03-26 08:34:11 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-kind-1.34-sev/2037085594152079360#1:build-log.txt%3A5063)

* _2026-03-26 07:35:38 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037070842281594880#1:build-log.txt%3A5081)

</details>

<hr/>
</details>

### 2026-03-25 (41x / 29.50%)


#### external (36x / 87.80%)

<details>
<summary> bazel remote cache blob fetch failure (35x / 85.37%) </summary>

<hr/>

**8x**: _2026-03-25 22:30:47 &#43;0000 UTC_: <code>22:34:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036892368967307264#1:build-log.txt%3A512)
<details>
<summary>all...</summary>

* _2026-03-25 22:46:26 &#43;0000 UTC_: <code>22:51:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036937636135833600#1:build-log.txt%3A503)

* _2026-03-25 22:30:47 &#43;0000 UTC_: <code>22:34:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036892368967307264#1:build-log.txt%3A512)

* _2026-03-25 22:29:24 &#43;0000 UTC_: <code>22:33:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036892368111669248#1:build-log.txt%3A461)

* _2026-03-25 22:28:59 &#43;0000 UTC_: <code>22:35:02: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036892367205699584#1:build-log.txt%3A517)

* _2026-03-25 22:28:15 &#43;0000 UTC_: <code>22:33:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-network/2036892367029538816#1:build-log.txt%3A529)

* _2026-03-25 22:27:03 &#43;0000 UTC_: <code>22:32:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036892366903709696#1:build-log.txt%3A503)

* _2026-03-25 21:43:49 &#43;0000 UTC_: <code>21:50:13: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036892369781002240#1:build-log.txt%3A459)

* _2026-03-25 20:39:55 &#43;0000 UTC_: <code>20:44:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036904750133284864#1:build-log.txt%3A445)

</details>

<hr/>

**7x**: _2026-03-25 17:27:42 &#43;0000 UTC_: <code>19:27:04: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036873472784732160#1:build-log.txt%3A385)
<details>
<summary>all...</summary>

* _2026-03-25 22:08:13 &#43;0000 UTC_: <code>22:12:17: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036857991180849152#1:build-log.txt%3A346)

* _2026-03-25 21:08:14 &#43;0000 UTC_: <code>21:12:54: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036862229206274048#1:build-log.txt%3A433)

* _2026-03-25 19:08:48 &#43;0000 UTC_: <code>19:16:10: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036882243925839872#1:build-log.txt%3A394)

* _2026-03-25 18:30:52 &#43;0000 UTC_: <code>18:36:56: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036873262981451776#1:build-log.txt%3A500)

* _2026-03-25 18:30:51 &#43;0000 UTC_: <code>18:37:58: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-network/2036873262641713152#1:build-log.txt%3A469)

* _2026-03-25 18:25:43 &#43;0000 UTC_: <code>18:30:50: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-windows2016/2036871960771694592#1:build-log.txt%3A468)

* _2026-03-25 17:27:42 &#43;0000 UTC_: <code>19:27:04: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036873472784732160#1:build-log.txt%3A385)

</details>

<hr/>

**4x**: _2026-03-25 21:47:10 &#43;0000 UTC_: <code>21:51:33: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17276/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036922551527018496#1:build-log.txt%3A509)
<details>
<summary>all...</summary>

* _2026-03-25 23:46:23 &#43;0000 UTC_: <code>23:50:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036952734418079744#1:build-log.txt%3A498)

* _2026-03-25 21:47:42 &#43;0000 UTC_: <code>21:52:56: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036892374516371456#1:build-log.txt%3A439)

* _2026-03-25 21:47:10 &#43;0000 UTC_: <code>21:51:33: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17276/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036922551527018496#1:build-log.txt%3A509)

* _2026-03-25 21:46:12 &#43;0000 UTC_: <code>21:51:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-network/2036892370359816192#1:build-log.txt%3A487)

</details>

<hr/>

**3x**: _2026-03-25 21:36:01 &#43;0000 UTC_: <code>21:40:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036909661638103040#1:build-log.txt%3A555)
<details>
<summary>all...</summary>

* _2026-03-25 21:37:27 &#43;0000 UTC_: <code>21:42:31: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-network/2036909662514712576#1:build-log.txt%3A542)

* _2026-03-25 21:36:01 &#43;0000 UTC_: <code>21:40:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036909661638103040#1:build-log.txt%3A555)

* _2026-03-25 21:35:18 &#43;0000 UTC_: <code>21:41:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036909659150880768#1:build-log.txt%3A648)

</details>

<hr/>

**1x**: _2026-03-25 20:46:36 &#43;0000 UTC_: <code>20:51:59: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036906547623235584#1:build-log.txt%3A519)
<details>
<summary>all...</summary>

* _2026-03-25 20:46:36 &#43;0000 UTC_: <code>20:51:59: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036906547623235584#1:build-log.txt%3A519)

</details>

<hr/>

**4x**: _2026-03-25 19:56:13 &#43;0000 UTC_: <code>20:02:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036894133938819072#1:build-log.txt%3A522)
<details>
<summary>all...</summary>

* _2026-03-25 21:43:46 &#43;0000 UTC_: <code>21:50:52: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036909664993546240#1:build-log.txt%3A565)

* _2026-03-25 21:08:47 &#43;0000 UTC_: <code>21:13:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036912517216735232#1:build-log.txt%3A557)

* _2026-03-25 19:56:46 &#43;0000 UTC_: <code>20:08:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036892366803046400#1:build-log.txt%3A506)

* _2026-03-25 19:56:13 &#43;0000 UTC_: <code>20:02:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036894133938819072#1:build-log.txt%3A522)

</details>

<hr/>

**4x**: _2026-03-25 19:46:36 &#43;0000 UTC_: <code>19:51:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-windows2016/2036892366194872320#1:build-log.txt%3A512)
<details>
<summary>all...</summary>

* _2026-03-25 19:52:50 &#43;0000 UTC_: <code>19:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-network/2036892366677217280#1:build-log.txt%3A529)

* _2026-03-25 19:52:47 &#43;0000 UTC_: <code>19:56:55: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036892366597525504#1:build-log.txt%3A586)

* _2026-03-25 19:49:08 &#43;0000 UTC_: <code>19:53:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036892366492667904#1:build-log.txt%3A532)

* _2026-03-25 19:46:36 &#43;0000 UTC_: <code>19:51:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-windows2016/2036892366194872320#1:build-log.txt%3A512)

</details>

<hr/>

**4x**: _2026-03-25 17:38:42 &#43;0000 UTC_: <code>17:42:45: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036860131525791744#1:build-log.txt%3A423)
<details>
<summary>all...</summary>

* _2026-03-25 21:48:21 &#43;0000 UTC_: <code>21:51:40: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036892400621719552#1:build-log.txt%3A438)

* _2026-03-25 17:45:13 &#43;0000 UTC_: <code>17:48:29: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-network/2036861683703484416#1:build-log.txt%3A432)

* _2026-03-25 17:38:42 &#43;0000 UTC_: <code>17:42:45: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036860131525791744#1:build-log.txt%3A423)

* _2026-03-25 17:29:54 &#43;0000 UTC_: <code>17:36:31: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036857570827702272#1:build-log.txt%3A458)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 2.44%) </summary>

<hr/>

**1x**: _2026-03-25 11:11:30 &#43;0000 UTC_: <code>11:21:44: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17069/pull-kubevirt-e2e-kind-sriov/2036723122534617088#1:build-log.txt%3A1464)
<details>
<summary>all...</summary>

* _2026-03-25 11:11:30 &#43;0000 UTC_: <code>11:21:44: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17069/pull-kubevirt-e2e-kind-sriov/2036723122534617088#1:build-log.txt%3A1464)

</details>

<hr/>
</details>

#### needs-investigation (2x / 4.88%)

<details>
<summary> no error snippets (2x / 4.88%) </summary>

<hr/>

**1x**: _2026-03-25 20:55:09 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2036909651676631040)
<details>
<summary>all...</summary>

* _2026-03-25 20:55:09 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2036909651676631040)

</details>

<hr/>

**1x**: _2026-03-25 18:46:20 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17207/pull-kubevirt-e2e-kind-1.34-sev/2036877259465297920)
<details>
<summary>all...</summary>

* _2026-03-25 18:46:20 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17207/pull-kubevirt-e2e-kind-1.34-sev/2036877259465297920)

</details>

<hr/>
</details>

#### internal (3x / 7.32%)

<details>
<summary> kind cluster creation failure (3x / 7.32%) </summary>

<hr/>

**3x**: _2026-03-25 17:05:27 &#43;0000 UTC_: <code>17:12:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17291/pull-kubevirt-e2e-kind-1.34-sev-1.7/2036851449563975680#1:build-log.txt%3A609)
<details>
<summary>all...</summary>

* _2026-03-25 22:46:16 &#43;0000 UTC_: <code>22:52:30: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036937635984838656#1:build-log.txt%3A697)

* _2026-03-25 17:07:07 &#43;0000 UTC_: <code>17:15:35: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17290/pull-kubevirt-e2e-kind-1.34-sev-1.8/2036852207432765440#1:build-log.txt%3A834)

* _2026-03-25 17:05:27 &#43;0000 UTC_: <code>17:12:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17291/pull-kubevirt-e2e-kind-1.34-sev-1.7/2036851449563975680#1:build-log.txt%3A609)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (122x / 87.77%)

<details>
<summary> bazel remote cache blob fetch failure (104x / 74.82%) </summary>

<hr/>

**21x**: _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:53:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037088891659358208#1:build-log.txt%3A518)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:34 &#43;0000 UTC_: <code>08:55:31: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-windows2016/2037088890090688512#1:build-log.txt%3A489)
<details><summary>context</summary>
<pre>08:55:16: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:55:16: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:55:16: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:55:31: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:55:31: INFO: Elapsed time: 58.186s, Critical Path: 37.91s
08:55:31: INFO: 6421 processes: 6391 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 08:47:27 &#43;0000 UTC_: <code>08:56:00: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-network/2037088890619170816#1:build-log.txt%3A516)
<details><summary>context</summary>
<pre>08:55:51: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:55:51: /usr/bin/ld: /tmp/go-link-4052249609/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:55:51: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:56:00: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:56:00: INFO: Elapsed time: 85.128s, Critical Path: 58.10s
08:56:00: INFO: 7659 processes: 7627 remote cache hit, 31 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:52:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037088891520946176#1:build-log.txt%3A515)
<details><summary>context</summary>
<pre>08:52:15: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:52:15: /usr/bin/ld: /tmp/go-link-1904356845/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:52:15: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:52:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:52:18: INFO: Elapsed time: 63.941s, Critical Path: 22.50s
08:52:18: INFO: 7503 processes: 7465 remote cache hit, 30 internal, 8 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:53:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037088891659358208#1:build-log.txt%3A518)
<details><summary>context</summary>
<pre>08:53:32: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:53:32: /usr/bin/ld: /tmp/go-link-2025919358/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:53:32: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:53:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:53:51: INFO: Elapsed time: 90.892s, Critical Path: 38.62s
08:53:51: INFO: 7551 processes: 7515 remote cache hit, 31 internal, 5 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:53:12: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2037088890333958144#1:build-log.txt%3A516)
<details><summary>context</summary>
<pre>08:52:20: /usr/bin/ld: /tmp/go-link-3408989329/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:52:20: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:53:12: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:53:12: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:53:12: INFO: Elapsed time: 53.656s, Critical Path: 29.02s
08:53:12: INFO: 6762 processes: 6732 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:52:38: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037088891369951232#1:build-log.txt%3A524)
<details><summary>context</summary>
<pre>08:52:14:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
08:52:14:       |   ^~~~~~~~
08:52:14:       |   vsprintf
08:52:38: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:52:39: INFO: Elapsed time: 49.050s, Critical Path: 14.33s
08:52:39: INFO: 7686 processes: 7655 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:56:52 &#43;0000 UTC_: <code>08:01:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037076170310160384#1:build-log.txt%3A511)
<details><summary>context</summary>
<pre>08:01:04: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:01:04: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:01:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:01:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:01:18: INFO: Elapsed time: 40.496s, Critical Path: 6.53s
08:01:18: INFO: 7642 processes: 7611 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:39 &#43;0000 UTC_: <code>07:53:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037073833353023488#1:build-log.txt%3A523)
<details><summary>context</summary>
<pre>07:53:25: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:25: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:53:25: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:37: INFO: Elapsed time: 49.860s, Critical Path: 30.72s
07:53:37: INFO: 6652 processes: 6623 remote cache hit, 29 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:35: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-windows2016/2037073830886772736#1:build-log.txt%3A549)
<details><summary>context</summary>
<pre>07:53:11: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:11: /usr/bin/ld: /tmp/go-link-4288142663/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:53:11: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:35: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:35: INFO: Elapsed time: 65.969s, Critical Path: 22.11s
07:53:35: INFO: 7683 processes: 7648 remote cache hit, 31 internal, 4 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:51:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-network/2037073834191884288#1:build-log.txt%3A524)
<details><summary>context</summary>
<pre>07:51:16: INFO: From GoLink cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client_/libvirt-hook-client:
07:51:16: /usr/bin/ld: /tmp/go-link-3408989329/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:51:16: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:51:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:51:30: INFO: Elapsed time: 26.018s, Critical Path: 5.26s
07:51:30: INFO: 4428 processes: 4399 remote cache hit, 29 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:51:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-network/2037073831499141120#1:build-log.txt%3A540)
<details><summary>context</summary>
<pre>07:51:16: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:51:16: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:51:16: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:51:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:51:29: INFO: Elapsed time: 36.425s, Critical Path: 17.42s
07:51:29: INFO: 7243 processes: 7211 remote cache hit, 31 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037073835894771712#1:build-log.txt%3A511)
<details><summary>context</summary>
<pre>07:53:12: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:12: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:53:12: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:48: INFO: Elapsed time: 82.156s, Critical Path: 65.13s
07:53:48: INFO: 7124 processes: 7090 remote cache hit, 31 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037073831637553152#1:build-log.txt%3A529)
<details><summary>context</summary>
<pre>07:53:16: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:16: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:53:16: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:29: INFO: Elapsed time: 60.002s, Critical Path: 46.11s
07:53:29: INFO: 7410 processes: 7375 remote cache hit, 31 internal, 4 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 06:47:31 &#43;0000 UTC_: <code>06:51:22: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2037058711255519232#1:build-log.txt%3A424)
<details><summary>context</summary>
<pre>06:51:07: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
06:51:07: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
06:51:07: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
06:51:22: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
06:51:22: INFO: Elapsed time: 47.159s, Critical Path: 27.48s
06:51:22: INFO: 7405 processes: 7374 remote cache hit, 30 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 22:46:26 &#43;0000 UTC_: <code>22:51:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036937636135833600#1:build-log.txt%3A503)
<details><summary>context</summary>
<pre>22:50:46: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
22:50:46: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
22:50:46: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
22:51:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
22:51:14: INFO: Elapsed time: 41.820s, Critical Path: 12.71s
22:51:14: INFO: 7605 processes: 7574 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 22:30:47 &#43;0000 UTC_: <code>22:34:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036892368967307264#1:build-log.txt%3A512)
<details><summary>context</summary>
<pre>22:34:39: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
22:34:39: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
22:34:39: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
22:34:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
22:34:48: INFO: Elapsed time: 27.669s, Critical Path: 11.73s
22:34:48: INFO: 5531 processes: 5502 remote cache hit, 29 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 22:29:24 &#43;0000 UTC_: <code>22:33:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036892368111669248#1:build-log.txt%3A461)
<details><summary>context</summary>
<pre>22:33:39: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
22:33:39: /usr/bin/ld: /tmp/go-link-624699758/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
22:33:39: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
22:33:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
22:33:46: INFO: Elapsed time: 32.323s, Critical Path: 19.90s
22:33:46: INFO: 7122 processes: 7088 remote cache hit, 31 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 22:28:59 &#43;0000 UTC_: <code>22:35:02: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036892367205699584#1:build-log.txt%3A517)
<details><summary>context</summary>
<pre>22:34:34: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
22:34:34: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
22:34:34: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
22:35:02: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
22:35:02: INFO: Elapsed time: 46.781s, Critical Path: 13.15s
22:35:02: INFO: 7683 processes: 7652 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 22:28:15 &#43;0000 UTC_: <code>22:33:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-network/2036892367029538816#1:build-log.txt%3A529)
<details><summary>context</summary>
<pre>22:33:18: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
22:33:18: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
22:33:18: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
22:33:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
22:33:40: INFO: Elapsed time: 28.118s, Critical Path: 15.86s
22:33:40: INFO: 7374 processes: 7342 remote cache hit, 31 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 21:43:49 &#43;0000 UTC_: <code>21:50:13: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036892369781002240#1:build-log.txt%3A459)
<details><summary>context</summary>
<pre>21:49:55: INFO: From GoLink cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client_/libvirt-hook-client:
21:49:55: /usr/bin/ld: /tmp/go-link-3408989329/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
21:49:55: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:50:13: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
21:50:13: INFO: Elapsed time: 31.154s, Critical Path: 12.76s
21:50:13: INFO: 5674 processes: 5643 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 20:39:55 &#43;0000 UTC_: <code>20:44:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036904750133284864#1:build-log.txt%3A445)
<details><summary>context</summary>
<pre>20:44:48: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
20:44:48: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
20:44:48: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
20:44:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
20:44:57: INFO: Elapsed time: 21.277s, Critical Path: 6.57s
20:44:57: INFO: 5057 processes: 5026 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**7x**: _2026-03-26 07:47:48 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037073832539328512#1:build-log.txt%3A502)
<details>
<summary>all...</summary>

* _2026-03-26 07:47:48 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037073832539328512#1:build-log.txt%3A502)
<details><summary>context</summary>
<pre>07:53:25: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:53:25: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:48: INFO: Elapsed time: 52.757s, Critical Path: 19.07s
07:53:48: INFO: 7417 processes: 7387 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:47 &#43;0000 UTC_: <code>07:53:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037073831079710720#1:build-log.txt%3A595)
<details><summary>context</summary>
<pre>07:53:12: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:53:12: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:34: INFO: Elapsed time: 40.587s, Critical Path: 29.62s
07:53:34: INFO: 7307 processes: 7276 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2037073831012601856#1:build-log.txt%3A540)
<details><summary>context</summary>
<pre>07:52:58: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:52:58: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:37: INFO: Elapsed time: 55.086s, Critical Path: 24.21s
07:53:37: INFO: 7678 processes: 7647 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037073831281037312#1:build-log.txt%3A530)
<details><summary>context</summary>
<pre>07:53:04: /usr/bin/ld: /tmp/go-link-4101630891/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:53:04: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:44: INFO: Elapsed time: 79.794s, Critical Path: 40.72s
07:53:44: INFO: 7678 processes: 7644 remote cache hit, 31 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 05:47:21 &#43;0000 UTC_: <code>05:51:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17276/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037043575102902272#1:build-log.txt%3A550)
<details><summary>context</summary>
<pre>05:51:33: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
05:51:33: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
05:51:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
05:51:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
05:51:47: INFO: Elapsed time: 38.372s, Critical Path: 5.53s
05:51:47: INFO: 7623 processes: 7592 remote cache hit, 29 internal, 2 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 04:47:33 &#43;0000 UTC_: <code>04:53:53: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037028524841242624#1:build-log.txt%3A499)
<details><summary>context</summary>
<pre>04:53:46: /usr/bin/ld: /tmp/go-link-3131746748/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
04:53:46: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
04:53:53: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
04:53:53: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
04:53:53: INFO: Elapsed time: 79.884s, Critical Path: 77.79s
04:53:53: INFO: 7590 processes: 7552 remote cache hit, 31 internal, 7 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 20:46:36 &#43;0000 UTC_: <code>20:51:59: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036906547623235584#1:build-log.txt%3A519)
<details><summary>context</summary>
<pre>20:51:38: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
20:51:38: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
20:51:59: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
20:51:59: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
20:51:59: INFO: Elapsed time: 37.913s, Critical Path: 13.67s
20:51:59: INFO: 7703 processes: 7672 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**32x**: _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036967849838252032#1:build-log.txt%3A446)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:52:26: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037088894213689344#1:build-log.txt%3A518)
<details><summary>context</summary>
<pre>08:52:23: /usr/bin/ld: /tmp/go-link-1282931798/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:52:23: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:52:26: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:52:26: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:52:26: INFO: Elapsed time: 47.519s, Critical Path: 30.87s
08:52:26: INFO: 7707 processes: 7672 remote cache hit, 31 internal, 4 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:55:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037088892464664576#1:build-log.txt%3A517)
<details><summary>context</summary>
<pre>08:54:49: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:55:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:55:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:55:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:55:34: INFO: Elapsed time: 61.383s, Critical Path: 30.20s
08:55:34: INFO: 7703 processes: 7672 remote cache hit, 30 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-operator/2037088891055378432#1:build-log.txt%3A482)
<details><summary>context</summary>
<pre>08:51:36: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:51:36: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
08:52:11: INFO: Elapsed time: 42.328s, Critical Path: 26.26s
08:52:11: INFO: 7698 processes: 7666 remote cache hit, 31 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:40 &#43;0000 UTC_: <code>07:51:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-operator/2037073831419449344#1:build-log.txt%3A529)
<details><summary>context</summary>
<pre>07:51:47: /usr/bin/ld: /tmp/go-link-347866404/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:51:47: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:51:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:51:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:51:51: INFO: Elapsed time: 45.409s, Critical Path: 30.60s
07:51:51: INFO: 7421 processes: 7382 remote cache hit, 31 internal, 8 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:39 &#43;0000 UTC_: <code>07:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037073831914377216#1:build-log.txt%3A525)
<details><summary>context</summary>
<pre>07:52:07: /usr/bin/ld: /tmp/go-link-3848556535/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:52:07: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:52:11: INFO: Elapsed time: 60.867s, Critical Path: 21.28s
07:52:11: INFO: 7697 processes: 7659 remote cache hit, 31 internal, 7 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-network/2037073831209734144#1:build-log.txt%3A540)
<details><summary>context</summary>
<pre>07:52:53: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:21: INFO: Elapsed time: 50.137s, Critical Path: 21.74s
07:53:21: INFO: 7703 processes: 7672 remote cache hit, 30 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037073836737826816#1:build-log.txt%3A529)
<details><summary>context</summary>
<pre>07:53:29: /usr/bin/ld: /tmp/go-link-945778869/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:53:29: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:32: INFO: Elapsed time: 80.390s, Critical Path: 37.38s
07:53:32: INFO: 7695 processes: 7657 remote cache hit, 31 internal, 7 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-compute/2037073831356534784#1:build-log.txt%3A460)
<details><summary>context</summary>
<pre>07:52:55: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:52:55: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:53:32: INFO: Elapsed time: 78.630s, Critical Path: 44.83s
07:53:32: INFO: 7698 processes: 7663 remote cache hit, 31 internal, 4 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 07:47:36 &#43;0000 UTC_: <code>07:54:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037073835026550784#1:build-log.txt%3A532)
<details><summary>context</summary>
<pre>07:53:56: /usr/bin/ld: /tmp/go-link-855134256/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
07:53:56: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
07:54:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:54:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
07:54:01: INFO: Elapsed time: 70.441s, Critical Path: 22.48s
07:54:01: INFO: 7699 processes: 7661 remote cache hit, 31 internal, 7 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 02:47:25 &#43;0000 UTC_: <code>02:52:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036998292239290368#1:build-log.txt%3A481)
<details><summary>context</summary>
<pre>02:52:03: /usr/bin/ld: /tmp/go-link-797515234/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
02:52:03: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
02:52:04: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
02:52:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
02:52:05: INFO: Elapsed time: 54.403s, Critical Path: 23.28s
02:52:05: INFO: 7679 processes: 7645 remote cache hit, 31 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:37 &#43;0000 UTC_: <code>00:55:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036967851490807808#1:build-log.txt%3A510)
<details><summary>context</summary>
<pre>00:55:16: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:55:16: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:55:16: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:55:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:55:44: INFO: Elapsed time: 46.422s, Critical Path: 25.27s
00:55:44: INFO: 7586 processes: 7553 remote cache hit, 31 internal, 2 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:32 &#43;0000 UTC_: <code>00:54:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-network/2036967848978419712#1:build-log.txt%3A451)
<details><summary>context</summary>
<pre>00:53:46: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:53:46: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:53:46: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:54:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:54:34: INFO: Elapsed time: 82.160s, Critical Path: 62.89s
00:54:34: INFO: 7407 processes: 7372 remote cache hit, 31 internal, 4 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036967848152141824#1:build-log.txt%3A478)
<details><summary>context</summary>
<pre>00:52:35: /usr/bin/ld: /tmp/go-link-1280386486/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:52:35: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:52:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:52:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:52:37: INFO: Elapsed time: 72.171s, Critical Path: 41.13s
00:52:37: INFO: 7682 processes: 7648 remote cache hit, 31 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036967849838252032#1:build-log.txt%3A446)
<details><summary>context</summary>
<pre>00:52:05: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:52:05: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:52:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:52:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:52:46: INFO: Elapsed time: 67.010s, Critical Path: 37.88s
00:52:46: INFO: 7603 processes: 7573 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036967847292309504#1:build-log.txt%3A508)
<details><summary>context</summary>
<pre>00:51:39: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:51:39: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:52:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:52:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:52:01: INFO: Elapsed time: 47.271s, Critical Path: 15.08s
00:52:01: INFO: 7693 processes: 7662 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:51:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-network/2036967846151458816#1:build-log.txt%3A521)
<details><summary>context</summary>
<pre>00:51:18: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:51:18: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:51:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:51:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:51:57: INFO: Elapsed time: 63.470s, Critical Path: 16.31s
00:51:57: INFO: 7681 processes: 7650 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:52:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036967845736222720#1:build-log.txt%3A484)
<details><summary>context</summary>
<pre>00:52:03: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:52:03: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:52:03: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:52:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:52:45: INFO: Elapsed time: 62.248s, Critical Path: 31.43s
00:52:45: INFO: 7434 processes: 7403 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:51:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036967846004658176#1:build-log.txt%3A506)
<details><summary>context</summary>
<pre>00:51:18: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:51:18: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:51:18: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:51:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:51:29: INFO: Elapsed time: 36.177s, Critical Path: 29.57s
00:51:29: INFO: 5541 processes: 5508 remote cache hit, 30 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:52:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036967850643558400#1:build-log.txt%3A505)
<details><summary>context</summary>
<pre>00:51:53: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:51:53: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:51:53: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:52:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:52:36: INFO: Elapsed time: 51.584s, Critical Path: 23.00s
00:52:36: INFO: 7686 processes: 7655 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:51:42: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036967846302453760#1:build-log.txt%3A509)
<details><summary>context</summary>
<pre>00:51:27: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:51:27: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:51:27: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:51:42: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:51:42: INFO: Elapsed time: 54.798s, Critical Path: 11.73s
00:51:42: INFO: 7513 processes: 7481 remote cache hit, 31 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:52:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036967845358735360#1:build-log.txt%3A588)
<details><summary>context</summary>
<pre>00:51:40: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:51:40: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:52:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:52:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:52:15: INFO: Elapsed time: 51.100s, Critical Path: 25.70s
00:52:15: INFO: 7673 processes: 7642 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:53:15: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-network/2036967845614587904#1:build-log.txt%3A520)
<details><summary>context</summary>
<pre>00:52:52: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:52:52: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:52:52: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:53:15: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:53:15: INFO: Elapsed time: 70.296s, Critical Path: 33.38s
00:53:15: INFO: 7698 processes: 7665 remote cache hit, 31 internal, 2 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:55:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036967845870440448#1:build-log.txt%3A504)
<details><summary>context</summary>
<pre>00:55:21: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:55:21: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:55:21: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:55:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:55:50: INFO: Elapsed time: 52.623s, Critical Path: 18.73s
00:55:50: INFO: 7686 processes: 7655 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:54:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036967845245489152#1:build-log.txt%3A517)
<details><summary>context</summary>
<pre>00:53:56:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
00:53:56:       |   ^~~~~~~~
00:53:56:       |   vsprintf
00:54:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:54:11: INFO: Elapsed time: 37.048s, Critical Path: 16.59s
00:54:11: INFO: 6728 processes: 6697 remote cache hit, 30 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:25 &#43;0000 UTC_: <code>00:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-monitoring/2036967845480370176#1:build-log.txt%3A508)
<details><summary>context</summary>
<pre>00:53:47: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:53:47: /usr/bin/ld: /tmp/go-link-2529653012/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:53:47: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:54:47: INFO: Elapsed time: 93.058s, Critical Path: 45.42s
00:54:47: INFO: 7692 processes: 7660 remote cache hit, 31 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:25 &#43;0000 UTC_: <code>00:52:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036967846478614528#1:build-log.txt%3A510)
<details><summary>context</summary>
<pre>00:52:21: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:52:21: /usr/bin/ld: /tmp/go-link-1071709154/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:52:21: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:52:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:52:25: INFO: Elapsed time: 66.264s, Critical Path: 25.85s
00:52:25: INFO: 7686 processes: 7647 remote cache hit, 31 internal, 8 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 00:46:23 &#43;0000 UTC_: <code>00:54:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-windows2016/2036967844863807488#1:build-log.txt%3A520)
<details><summary>context</summary>
<pre>00:54:11: /usr/bin/ld: /tmp/go-link-2529653012/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
00:54:11: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
00:54:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:54:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
00:54:50: INFO: Elapsed time: 95.692s, Critical Path: 63.61s
00:54:50: INFO: 7700 processes: 7669 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 23:46:23 &#43;0000 UTC_: <code>23:50:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036952734418079744#1:build-log.txt%3A498)
<details><summary>context</summary>
<pre>23:50:33: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
23:50:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
23:50:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
23:50:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
23:50:45: INFO: Elapsed time: 28.463s, Critical Path: 5.74s
23:50:45: INFO: 7207 processes: 7176 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 22:27:03 &#43;0000 UTC_: <code>22:32:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036892366903709696#1:build-log.txt%3A503)
<details><summary>context</summary>
<pre>22:32:31: /usr/bin/ld: /tmp/go-link-4032574104/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
22:32:31: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
22:32:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
22:32:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
22:32:37: INFO: Elapsed time: 75.246s, Critical Path: 32.43s
22:32:37: INFO: 7701 processes: 7663 remote cache hit, 31 internal, 7 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 21:47:42 &#43;0000 UTC_: <code>21:52:56: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036892374516371456#1:build-log.txt%3A439)
<details><summary>context</summary>
<pre>21:52:30: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
21:52:30: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:52:56: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
21:52:56: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
21:52:56: INFO: Elapsed time: 29.806s, Critical Path: 14.63s
21:52:56: INFO: 7438 processes: 7407 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 21:47:10 &#43;0000 UTC_: <code>21:51:33: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17276/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036922551527018496#1:build-log.txt%3A509)
<details><summary>context</summary>
<pre>21:51:21: /usr/bin/ld: /tmp/go-link-3369357513/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
21:51:21: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:51:33: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
21:51:33: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
21:51:33: INFO: Elapsed time: 52.569s, Critical Path: 25.96s
21:51:33: INFO: 7607 processes: 7565 remote cache hit, 31 internal, 11 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-25 21:46:12 &#43;0000 UTC_: <code>21:51:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-network/2036892370359816192#1:build-log.txt%3A487)
<details><summary>context</summary>
<pre>21:51:17: /usr/bin/ld: /tmp/go-link-4253825822/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
21:51:17: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:51:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
21:51:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
21:51:18: INFO: Elapsed time: 58.055s, Critical Path: 20.79s
21:51:18: INFO: 7694 processes: 7659 remote cache hit, 31 internal, 4 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**15x**: _2026-03-26 05:04:28 &#43;0000 UTC_: <code>05:08:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037032781342576640#1:build-log.txt%3A534)
<details>
<summary>all...</summary>

* _2026-03-26 05:04:37 &#43;0000 UTC_: <code>05:09:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-network/2037032781271273472#1:build-log.txt%3A572)
<details><summary>context</summary>
<pre>05:09:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
05:09:45: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
05:09:45: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
05:09:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
05:09:45: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
05:09:45: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
05:09:45: INFO: Elapsed time: 41.707s, Critical Path: 21.86s</pre>
</details>


* _2026-03-26 05:04:29 &#43;0000 UTC_: <code>05:08:54: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037032781178998784#1:build-log.txt%3A529)
<details><summary>context</summary>
<pre>05:08:36: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
05:08:36: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
05:08:36: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
05:08:54: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
05:08:54: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
05:08:54: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
05:08:54: INFO: Elapsed time: 37.336s, Critical Path: 13.72s</pre>
</details>


* _2026-03-26 05:04:29 &#43;0000 UTC_: <code>05:08:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037032781422268416#1:build-log.txt%3A541)
<details><summary>context</summary>
<pre>05:08:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
05:08:51: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
05:08:51: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
05:08:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
05:08:51: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
05:08:51: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
05:08:51: INFO: Elapsed time: 51.678s, Critical Path: 22.73s</pre>
</details>


* _2026-03-26 05:04:28 &#43;0000 UTC_: <code>05:08:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037032781342576640#1:build-log.txt%3A534)
<details><summary>context</summary>
<pre>05:08:05: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
05:08:05: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
05:08:05: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
05:08:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
05:08:25: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
05:08:25: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
05:08:25: INFO: Elapsed time: 30.357s, Critical Path: 11.82s</pre>
</details>


* _2026-03-26 05:04:26 &#43;0000 UTC_: <code>05:10:23: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037032781506154496#1:build-log.txt%3A522)
<details><summary>context</summary>
<pre>05:10:23: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
05:10:23: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
05:10:23: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
05:10:23: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
05:10:23: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
05:10:23: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
05:10:23: INFO: Elapsed time: 57.533s, Critical Path: 23.46s</pre>
</details>


* _2026-03-26 05:04:26 &#43;0000 UTC_: <code>05:10:15: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037032781090918400#1:build-log.txt%3A632)
<details><summary>context</summary>
<pre>05:10:15: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
05:10:15: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
05:10:15: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
05:10:15: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
05:10:15: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
05:10:15: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
05:10:15: INFO: Elapsed time: 52.140s, Critical Path: 28.77s</pre>
</details>


* _2026-03-26 01:49:19 &#43;0000 UTC_: <code>01:59:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036983575923920896#1:build-log.txt%3A480)
<details><summary>context</summary>
<pre>01:58:02: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
01:58:02: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
01:58:02: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
01:59:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:59:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:59:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:59:15: INFO: Elapsed time: 97.126s, Critical Path: 56.30s</pre>
</details>


* _2026-03-26 01:49:19 &#43;0000 UTC_: <code>01:59:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-windows2016/2036983572635586560#1:build-log.txt%3A561)
<details><summary>context</summary>
<pre>01:58:02:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
01:58:02:       |   ^~~~~~~~
01:58:02:       |   vsprintf
01:59:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:59:05: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:59:05: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:59:05: INFO: Elapsed time: 89.207s, Critical Path: 58.44s</pre>
</details>


* _2026-03-26 01:49:16 &#43;0000 UTC_: <code>01:59:08: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036983573491224576#1:build-log.txt%3A496)
<details><summary>context</summary>
<pre>01:58:02: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
01:58:02: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
01:58:02: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
01:59:08: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:59:08: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:59:08: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:59:08: INFO: Elapsed time: 90.354s, Critical Path: 59.12s</pre>
</details>


* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:53:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036983573973569536#1:build-log.txt%3A531)
<details><summary>context</summary>
<pre>01:53:29: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
01:53:29: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
01:53:29: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
01:53:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:53:34: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:53:34: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:53:34: INFO: Elapsed time: 34.693s, Critical Path: 13.69s</pre>
</details>


* _2026-03-26 01:48:56 &#43;0000 UTC_: <code>01:54:56: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-network/2036983573382172672#1:build-log.txt%3A538)
<details><summary>context</summary>
<pre>01:54:03: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
01:54:03: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
01:54:03: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
01:54:56: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:54:56: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:54:56: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:54:56: INFO: Elapsed time: 59.946s, Critical Path: 27.62s</pre>
</details>


* _2026-03-26 01:48:56 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-network/2036983573851934720#1:build-log.txt%3A538)
<details><summary>context</summary>
<pre>01:54:30:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
01:54:30:       |   ^~~~~~~~
01:54:30:       |   vsprintf
01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:55:16: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:55:16: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:55:16: INFO: Elapsed time: 56.180s, Critical Path: 18.88s</pre>
</details>


* _2026-03-25 21:37:27 &#43;0000 UTC_: <code>21:42:31: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-network/2036909662514712576#1:build-log.txt%3A542)
<details><summary>context</summary>
<pre>21:42:19: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:42:19: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
21:42:19: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:42:31: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
21:42:31: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
21:42:31: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
21:42:32: INFO: Elapsed time: 34.574s, Critical Path: 6.51s</pre>
</details>


* _2026-03-25 21:36:01 &#43;0000 UTC_: <code>21:40:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036909661638103040#1:build-log.txt%3A555)
<details><summary>context</summary>
<pre>21:39:51: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:39:51: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
21:39:51: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:40:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
21:40:21: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
21:40:21: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
21:40:21: INFO: Elapsed time: 44.479s, Critical Path: 13.17s</pre>
</details>


* _2026-03-25 21:35:18 &#43;0000 UTC_: <code>21:41:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036909659150880768#1:build-log.txt%3A648)
<details><summary>context</summary>
<pre>21:41:33: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:41:33: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
21:41:33: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:41:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
21:41:47: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
21:41:47: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
21:41:47: INFO: Elapsed time: 39.956s, Critical Path: 14.05s</pre>
</details>


</details>

<hr/>

**6x**: _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036983575089254400#1:build-log.txt%3A505)
<details>
<summary>all...</summary>

* _2026-03-26 01:49:28 &#43;0000 UTC_: <code>01:59:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036983573042434048#1:build-log.txt%3A534)
<details><summary>context</summary>
<pre>01:59:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:59:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:59:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:59:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:59:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:59:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:59:14: INFO: Elapsed time: 121.072s, Critical Path: 74.67s</pre>
</details>


* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036983574225227776#1:build-log.txt%3A534)
<details><summary>context</summary>
<pre>01:57:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:57:46: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:57:46: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:57:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:57:46: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:57:46: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:57:46: INFO: Elapsed time: 84.158s, Critical Path: 40.43s</pre>
</details>


* _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036983575089254400#1:build-log.txt%3A505)
<details><summary>context</summary>
<pre>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:55:16: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:55:16: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:55:16: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:55:16: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:55:16: INFO: Elapsed time: 74.554s, Critical Path: 35.38s</pre>
</details>


* _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:19: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036983573713522688#1:build-log.txt%3A535)
<details><summary>context</summary>
<pre>01:55:19: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:55:19: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:55:19: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:55:19: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:55:19: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:55:19: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:55:20: INFO: Elapsed time: 60.787s, Critical Path: 29.80s</pre>
</details>


* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:55:04: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036983578432114688#1:build-log.txt%3A527)
<details><summary>context</summary>
<pre>01:55:04: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:55:04: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:55:04: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:55:04: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:55:04: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:55:04: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:55:04: INFO: Elapsed time: 51.737s, Critical Path: 20.20s</pre>
</details>


* _2026-03-26 01:48:57 &#43;0000 UTC_: <code>01:55:06: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036983573587693568#1:build-log.txt%3A540)
<details><summary>context</summary>
<pre>01:55:06: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:55:06: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:55:06: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:55:06: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:55:06: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:55:06: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:55:06: INFO: Elapsed time: 61.230s, Critical Path: 42.19s</pre>
</details>


</details>

<hr/>

**8x**: _2026-03-25 19:56:13 &#43;0000 UTC_: <code>20:02:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036894133938819072#1:build-log.txt%3A522)
<details>
<summary>all...</summary>

* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036983579300335616#1:build-log.txt%3A541)
<details><summary>context</summary>
<pre>01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:57:14: INFO: Elapsed time: 57.683s, Critical Path: 28.90s</pre>
</details>


* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036983573143097344#1:build-log.txt%3A615)
<details><summary>context</summary>
<pre>01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:57:14: INFO: Elapsed time: 56.981s, Critical Path: 27.81s</pre>
</details>


* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:56:07: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-network/2036983576750198784#1:build-log.txt%3A549)
<details><summary>context</summary>
<pre>01:56:07: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:56:07: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:56:07: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:56:07: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:56:07: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:56:07: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:56:07: INFO: Elapsed time: 62.506s, Critical Path: 33.50s</pre>
</details>


* _2026-03-26 01:48:57 &#43;0000 UTC_: <code>01:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036983577610031104#1:build-log.txt%3A481)
<details><summary>context</summary>
<pre>01:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:54:47: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:54:47: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
01:54:47: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
01:54:47: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
01:54:47: INFO: Elapsed time: 42.126s, Critical Path: 18.22s</pre>
</details>


* _2026-03-25 21:43:46 &#43;0000 UTC_: <code>21:50:52: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036909664993546240#1:build-log.txt%3A565)
<details><summary>context</summary>
<pre>21:50:52: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
21:50:52: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
21:50:52: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
21:50:52: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
21:50:52: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
21:50:52: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
21:50:52: INFO: Elapsed time: 69.978s, Critical Path: 25.19s</pre>
</details>


* _2026-03-25 21:08:47 &#43;0000 UTC_: <code>21:13:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036912517216735232#1:build-log.txt%3A557)
<details><summary>context</summary>
<pre>21:13:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
21:13:18: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
21:13:18: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
21:13:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
21:13:18: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: ad00773da5a5af3427cc60043da295cc3948fa0e7d75a9c5b30cd73d744605f2/22302 for bazel-out/k8-fastbuild/bin/vendor/github.com/mitchellh/go-vnc/go_default_library.x
21:13:18: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4781398cf1ae263e2808fe390f4a931ba696e3ed541e68e4a87992bdedea1cc9/48934 for bazel-out/k8-fastbuild/bin/tests/libssh/go_default_library.x
21:13:18: INFO: Elapsed time: 33.704s, Critical Path: 6.41s</pre>
</details>


* _2026-03-25 19:56:46 &#43;0000 UTC_: <code>20:08:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036892366803046400#1:build-log.txt%3A506)
<details><summary>context</summary>
<pre>20:08:23: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
20:08:23: /usr/bin/ld: /tmp/go-link-4191254286/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
20:08:23: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
20:08:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
20:08:24: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d28cd0c48ee89103d13e6e7f79d4a0dea1133a032079012c434887177586f5d4/4163532 for bazel-out/k8-fastbuild/bin/tests/storage/go_default_library.x
20:08:24: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
20:08:24: INFO: Elapsed time: 157.194s, Critical Path: 18.80s</pre>
</details>


* _2026-03-25 19:56:13 &#43;0000 UTC_: <code>20:02:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036894133938819072#1:build-log.txt%3A522)
<details><summary>context</summary>
<pre>20:02:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
20:02:40: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
20:02:40: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d28cd0c48ee89103d13e6e7f79d4a0dea1133a032079012c434887177586f5d4/4163532 for bazel-out/k8-fastbuild/bin/tests/storage/go_default_library.x
20:02:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
20:02:40: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
20:02:40: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d28cd0c48ee89103d13e6e7f79d4a0dea1133a032079012c434887177586f5d4/4163532 for bazel-out/k8-fastbuild/bin/tests/storage/go_default_library.x
20:02:40: INFO: Elapsed time: 114.520s, Critical Path: 15.74s</pre>
</details>


</details>

<hr/>

**4x**: _2026-03-25 19:46:36 &#43;0000 UTC_: <code>19:51:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-windows2016/2036892366194872320#1:build-log.txt%3A512)
<details>
<summary>all...</summary>

* _2026-03-25 19:52:50 &#43;0000 UTC_: <code>19:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-network/2036892366677217280#1:build-log.txt%3A529)
<details><summary>context</summary>
<pre>19:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
19:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d28cd0c48ee89103d13e6e7f79d4a0dea1133a032079012c434887177586f5d4/4163532 for bazel-out/k8-fastbuild/bin/tests/storage/go_default_library.x
19:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
19:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
19:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d28cd0c48ee89103d13e6e7f79d4a0dea1133a032079012c434887177586f5d4/4163532 for bazel-out/k8-fastbuild/bin/tests/storage/go_default_library.x
19:57:14: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
19:57:15: INFO: Elapsed time: 39.917s, Critical Path: 12.72s</pre>
</details>


* _2026-03-25 19:52:47 &#43;0000 UTC_: <code>19:56:55: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036892366597525504#1:build-log.txt%3A586)
<details><summary>context</summary>
<pre>19:56:53: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
19:56:53: /usr/bin/ld: /tmp/go-link-1800051056/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
19:56:53: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
19:56:55: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
19:56:55: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d28cd0c48ee89103d13e6e7f79d4a0dea1133a032079012c434887177586f5d4/4163532 for bazel-out/k8-fastbuild/bin/tests/storage/go_default_library.x
19:56:55: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
19:56:55: INFO: Elapsed time: 46.178s, Critical Path: 15.41s</pre>
</details>


* _2026-03-25 19:49:08 &#43;0000 UTC_: <code>19:53:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036892366492667904#1:build-log.txt%3A532)
<details><summary>context</summary>
<pre>19:53:44: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
19:53:44: /usr/bin/ld: /tmp/go-link-1614599599/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
19:53:44: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
19:53:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
19:53:50: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d28cd0c48ee89103d13e6e7f79d4a0dea1133a032079012c434887177586f5d4/4163532 for bazel-out/k8-fastbuild/bin/tests/storage/go_default_library.x
19:53:50: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
19:53:50: INFO: Elapsed time: 49.098s, Critical Path: 17.60s</pre>
</details>


* _2026-03-25 19:46:36 &#43;0000 UTC_: <code>19:51:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-windows2016/2036892366194872320#1:build-log.txt%3A512)
<details><summary>context</summary>
<pre>19:51:09: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
19:51:09: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
19:51:09: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
19:51:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
19:51:36: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: d28cd0c48ee89103d13e6e7f79d4a0dea1133a032079012c434887177586f5d4/4163532 for bazel-out/k8-fastbuild/bin/tests/storage/go_default_library.x
19:51:36: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x
19:51:37: INFO: Elapsed time: 47.631s, Critical Path: 17.44s</pre>
</details>


</details>

<hr/>

**7x**: _2026-03-25 17:27:42 &#43;0000 UTC_: <code>19:27:04: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036873472784732160#1:build-log.txt%3A385)
<details>
<summary>all...</summary>

* _2026-03-25 22:08:13 &#43;0000 UTC_: <code>22:12:17: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036857991180849152#1:build-log.txt%3A346)
<details><summary>context</summary>
<pre>22:12:08: INFO: Invocation ID: 26a3e0f6-e07f-42c4-8d7e-5f30a963ad90
22:12:08: INFO: Analyzed 30 targets (0 packages loaded, 0 targets configured).
22:12:08: INFO: Found 30 targets...
22:12:17: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
22:12:17: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07/1384 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07
22:12:17: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d/1669 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d
22:12:17: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 10bdee2f647d1c55da8c3b7dd34c7ebfe19e8ebd53db457aa74b201bc19ed932/308 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/index.json</pre>
</details>


* _2026-03-25 21:08:14 &#43;0000 UTC_: <code>21:12:54: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036862229206274048#1:build-log.txt%3A433)
<details><summary>context</summary>
<pre>21:12:36: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:12:36: /usr/bin/ld: /tmp/go-link-2942130853/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
21:12:36: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:12:54: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
21:12:54: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 10bdee2f647d1c55da8c3b7dd34c7ebfe19e8ebd53db457aa74b201bc19ed932/308 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/index.json
21:12:54: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d/1669 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d
21:12:54: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07/1384 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07</pre>
</details>


* _2026-03-25 19:08:48 &#43;0000 UTC_: <code>19:16:10: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036882243925839872#1:build-log.txt%3A394)
<details><summary>context</summary>
<pre>19:16:01: INFO: Invocation ID: 59934a59-cbe1-41af-99d4-ae852453ecb4
19:16:01: INFO: Analyzed 30 targets (0 packages loaded, 0 targets configured).
19:16:01: INFO: Found 30 targets...
19:16:10: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
19:16:10: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07/1384 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07
19:16:10: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d/1669 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d
19:16:10: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 10bdee2f647d1c55da8c3b7dd34c7ebfe19e8ebd53db457aa74b201bc19ed932/308 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/index.json</pre>
</details>


* _2026-03-25 18:30:52 &#43;0000 UTC_: <code>18:36:56: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036873262981451776#1:build-log.txt%3A500)
<details><summary>context</summary>
<pre>18:36:35: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
18:36:35: /usr/bin/ld: /tmp/go-link-738077025/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
18:36:35: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
18:36:56: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
18:36:56: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07/1384 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07
18:36:56: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d/1669 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d
18:36:56: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 10bdee2f647d1c55da8c3b7dd34c7ebfe19e8ebd53db457aa74b201bc19ed932/308 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/index.json</pre>
</details>


* _2026-03-25 18:30:51 &#43;0000 UTC_: <code>18:37:58: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-network/2036873262641713152#1:build-log.txt%3A469)
<details><summary>context</summary>
<pre>18:37:50: INFO: Invocation ID: 31eed13a-fc49-49cc-bffd-9dc29f46161e
18:37:50: INFO: Analyzed 30 targets (0 packages loaded, 0 targets configured).
18:37:50: INFO: Found 30 targets...
18:37:58: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
18:37:58: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07/1384 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07
18:37:58: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 10bdee2f647d1c55da8c3b7dd34c7ebfe19e8ebd53db457aa74b201bc19ed932/308 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/index.json
18:37:58: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d/1669 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d</pre>
</details>


* _2026-03-25 18:25:43 &#43;0000 UTC_: <code>18:30:50: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-windows2016/2036871960771694592#1:build-log.txt%3A468)
<details><summary>context</summary>
<pre>18:30:32: INFO: Invocation ID: 4fd89832-76c0-4f21-ad61-a27d9abc9ac7
18:30:32: INFO: Analyzed 30 targets (0 packages loaded, 0 targets configured).
18:30:32: INFO: Found 30 targets...
18:30:50: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
18:30:50: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07/1384 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07
18:30:50: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d/1669 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d
18:30:50: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 10bdee2f647d1c55da8c3b7dd34c7ebfe19e8ebd53db457aa74b201bc19ed932/308 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/index.json</pre>
</details>


* _2026-03-25 17:27:42 &#43;0000 UTC_: <code>19:27:04: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036873472784732160#1:build-log.txt%3A385)
<details><summary>context</summary>
<pre>19:26:57: INFO: Invocation ID: c4562b4d-c07f-4583-8fa9-4b1da3ff3ac6
19:26:57: INFO: Analyzed 30 targets (0 packages loaded, 0 targets configured).
19:26:57: INFO: Found 30 targets...
19:27:04: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
19:27:04: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07/1384 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07
19:27:04: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 10bdee2f647d1c55da8c3b7dd34c7ebfe19e8ebd53db457aa74b201bc19ed932/308 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/index.json
19:27:04: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d/1669 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d</pre>
</details>


</details>

<hr/>

**4x**: _2026-03-25 17:38:42 &#43;0000 UTC_: <code>17:42:45: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036860131525791744#1:build-log.txt%3A423)
<details>
<summary>all...</summary>

* _2026-03-25 21:48:21 &#43;0000 UTC_: <code>21:51:40: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036892400621719552#1:build-log.txt%3A438)
<details><summary>context</summary>
<pre>21:51:36: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:51:36: /usr/bin/ld: /tmp/go-link-827060483/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
21:51:36: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
21:51:40: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
21:51:40: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 897af945d1c58366086d5933ae4f341a5f1413b88e6c7f2b659436adc5d0f522/573 for bazel-out/k8-fastbuild/bin/external/fedora_with_test_tooling_single/layout/blobs/sha256/897af945d1c58366086d5933ae4f341a5f1413b88e6c7f2b659436adc5d0f522
21:51:40: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 99247de7916ab228ca6ceaed914e22930ae06350ea9612631f0354517058e36c/307 for bazel-out/k8-fastbuild/bin/external/fedora_with_test_tooling_single/layout/index.json
21:51:40: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: e277c9b07fac83163eae97c509b0b054ad1e1dc360b0d9df6dcbe18562c0364d/519 for bazel-out/k8-fastbuild/bin/external/fedora_with_test_tooling_single/layout/blobs/sha256/e277c9b07fac83163eae97c509b0b054ad1e1dc360b0d9df6dcbe18562c0364d</pre>
</details>


* _2026-03-25 17:45:13 &#43;0000 UTC_: <code>17:48:29: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-network/2036861683703484416#1:build-log.txt%3A432)
<details><summary>context</summary>
<pre>17:48:24: INFO: From GoLink cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client_/libvirt-hook-client:
17:48:24: /usr/bin/ld: /tmp/go-link-3408989329/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
17:48:24: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
17:48:29: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
17:48:29: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 10bdee2f647d1c55da8c3b7dd34c7ebfe19e8ebd53db457aa74b201bc19ed932/308 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/index.json
17:48:29: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07/1384 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/6f6ee6da8a45d3d9afb18347e508a5aebbfb2849d87da625726d1c37d9860b07
17:48:29: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d/1669 for bazel-out/k8-fastbuild/bin/external/alpine-ext-kernel-boot-demo-container-base_single/layout/blobs/sha256/bccd990554f55623d96fa70bc7efc553dd617523ebca76919b917ad3ee616c1d</pre>
</details>


* _2026-03-25 17:38:42 &#43;0000 UTC_: <code>17:42:45: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036860131525791744#1:build-log.txt%3A423)
<details><summary>context</summary>
<pre>17:42:35: INFO: From GoLink cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client_/libvirt-hook-client:
17:42:35: /usr/bin/ld: /tmp/go-link-3408989329/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
17:42:35: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
17:42:45: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
17:42:45: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: e277c9b07fac83163eae97c509b0b054ad1e1dc360b0d9df6dcbe18562c0364d/519 for bazel-out/k8-fastbuild/bin/external/fedora_with_test_tooling_single/layout/blobs/sha256/e277c9b07fac83163eae97c509b0b054ad1e1dc360b0d9df6dcbe18562c0364d
17:42:45: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 897af945d1c58366086d5933ae4f341a5f1413b88e6c7f2b659436adc5d0f522/573 for bazel-out/k8-fastbuild/bin/external/fedora_with_test_tooling_single/layout/blobs/sha256/897af945d1c58366086d5933ae4f341a5f1413b88e6c7f2b659436adc5d0f522
17:42:45: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 99247de7916ab228ca6ceaed914e22930ae06350ea9612631f0354517058e36c/307 for bazel-out/k8-fastbuild/bin/external/fedora_with_test_tooling_single/layout/index.json</pre>
</details>


* _2026-03-25 17:29:54 &#43;0000 UTC_: <code>17:36:31: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036857570827702272#1:build-log.txt%3A458)
<details><summary>context</summary>
<pre>17:36:27: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
17:36:27: /usr/bin/ld: /tmp/go-link-1286366182/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
17:36:27: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
17:36:31: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:
17:36:31: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 99247de7916ab228ca6ceaed914e22930ae06350ea9612631f0354517058e36c/307 for bazel-out/k8-fastbuild/bin/external/fedora_with_test_tooling_single/layout/index.json
17:36:31: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: e277c9b07fac83163eae97c509b0b054ad1e1dc360b0d9df6dcbe18562c0364d/519 for bazel-out/k8-fastbuild/bin/external/fedora_with_test_tooling_single/layout/blobs/sha256/e277c9b07fac83163eae97c509b0b054ad1e1dc360b0d9df6dcbe18562c0364d
17:36:31: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 897af945d1c58366086d5933ae4f341a5f1413b88e6c7f2b659436adc5d0f522/573 for bazel-out/k8-fastbuild/bin/external/fedora_with_test_tooling_single/layout/blobs/sha256/897af945d1c58366086d5933ae4f341a5f1413b88e6c7f2b659436adc5d0f522</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (7x / 5.04%) </summary>

<hr/>

**5x**: _2026-03-26 08:52:44 &#43;0000 UTC_: <code>08:57:27: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 33 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037090118291951616#1:build-log.txt%3A439)
<details>
<summary>all...</summary>

* _2026-03-26 08:52:44 &#43;0000 UTC_: <code>08:57:27: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 33 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037090118291951616#1:build-log.txt%3A439)
<details><summary>context</summary>
<pre>08:57:27: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:27: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:27: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:27: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 33 errors during bulk transfer:
08:57:27: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:27: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:27: java.net.UnknownHostException: bazel-cache.kubevirt-prow</pre>
</details>


* _2026-03-26 08:52:43 &#43;0000 UTC_: <code>08:57:26: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-launcher/libvirt-hook-client/BUILD.bazel:10:10: GoCompilePkg cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client.a failed: Exec failed due to IOException: 351 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-network/2037090113296535552#1:build-log.txt%3A760)
<details><summary>context</summary>
<pre>08:57:26: INFO: Analyzed 30 targets (0 packages loaded, 0 targets configured).
08:57:26: INFO: Found 30 targets...
08:57:26: WARNING: Remote Cache: bazel-cache.kubevirt-prow
08:57:26: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-launcher/libvirt-hook-client/BUILD.bazel:10:10: GoCompilePkg cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client.a failed: Exec failed due to IOException: 351 errors during bulk transfer:
08:57:26: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:26: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:26: java.net.UnknownHostException: bazel-cache.kubevirt-prow</pre>
</details>


* _2026-03-26 08:52:30 &#43;0000 UTC_: <code>08:58:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tools/csv-generator/BUILD.bazel:16:10 Validating nogo output for //tools/csv-generator:csv-generator failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037090114504495104#1:build-log.txt%3A490)
<details><summary>context</summary>
<pre>08:58:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/go.yaml.in/yaml/v2/BUILD.bazel:3:11: Running nogo on //vendor/go.yaml.in/yaml/v2:go_default_library failed: Exec failed due to IOException: 2 errors during bulk transfer:
08:58:24: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:58:24: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:58:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tools/csv-generator/BUILD.bazel:16:10 Validating nogo output for //tools/csv-generator:csv-generator failed: Exec failed due to IOException: 2 errors during bulk transfer:
08:58:24: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:58:24: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:58:24: INFO: Elapsed time: 0.719s, Critical Path: 0.42s</pre>
</details>


* _2026-03-26 08:52:28 &#43;0000 UTC_: <code>08:57:24: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 73 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-network/2037090114290585600#1:build-log.txt%3A725)
<details><summary>context</summary>
<pre>08:57:24: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:24: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:24: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:24: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 73 errors during bulk transfer:
08:57:24: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:24: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:24: java.net.UnknownHostException: bazel-cache.kubevirt-prow</pre>
</details>


* _2026-03-26 08:52:20 &#43;0000 UTC_: <code>08:57:23: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportserver/BUILD.bazel:20:10: GoLink cmd/virt-exportserver/virt-exportserver_/virt-exportserver failed: Exec failed due to IOException: 161 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037090113539805184#1:build-log.txt%3A479)
<details><summary>context</summary>
<pre>08:57:22: INFO: Analyzed 30 targets (0 packages loaded, 0 targets configured).
08:57:22: INFO: Found 30 targets...
08:57:23: WARNING: Remote Cache: bazel-cache.kubevirt-prow
08:57:23: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportserver/BUILD.bazel:20:10: GoLink cmd/virt-exportserver/virt-exportserver_/virt-exportserver failed: Exec failed due to IOException: 161 errors during bulk transfer:
08:57:23: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:23: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:23: java.net.UnknownHostException: bazel-cache.kubevirt-prow</pre>
</details>


</details>

<hr/>

**1x**: _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:58:56: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/virt-handler/vsock/system/BUILD.bazel:3:11 Validating nogo output for //pkg/virt-handler/vsock/system:go_default_library failed: Exec failed due to IOException: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037088895878828032#1:build-log.txt%3A883)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:58:56: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/virt-handler/vsock/system/BUILD.bazel:3:11 Validating nogo output for //pkg/virt-handler/vsock/system:go_default_library failed: Exec failed due to IOException: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037088895878828032#1:build-log.txt%3A883)
<details><summary>context</summary>
<pre>08:58:56: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:58:56: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:58:56: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:58:56: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/virt-handler/vsock/system/BUILD.bazel:3:11 Validating nogo output for //pkg/virt-handler/vsock/system:go_default_library failed: Exec failed due to IOException: 3 errors during bulk transfer:
08:58:56: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:58:56: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:58:56: java.net.UnknownHostException: bazel-cache.kubevirt-prow</pre>
</details>


</details>

<hr/>

**1x**: _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:55:51: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/pierrec/lz4/v4/internal/xxh32/BUILD.bazel:3:11: Validating nogo output for //vendor/github.com/pierrec/lz4/v4/internal/xxh32:go_default_library failed: Exec failed due to IOException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-network/2037088893345468416#1:build-log.txt%3A467)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:55:51: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/pierrec/lz4/v4/internal/xxh32/BUILD.bazel:3:11: Validating nogo output for //vendor/github.com/pierrec/lz4/v4/internal/xxh32:go_default_library failed: Exec failed due to IOException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-network/2037088893345468416#1:build-log.txt%3A467)
<details><summary>context</summary>
<pre>08:54:51: WARNING: Remote Cache: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/ac/2887e295c86357bcf9b07c50908dc03b9d6e3ad2ca4e291374f4cc02592abb4d&#39; timed out. Received 0 bytes.
08:54:51: WARNING: Remote Cache: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/ac/e7dfc8ec01b7a8ba26df8ad987f881d5736b79a2324bed261cbb5deb9065322e&#39; timed out. Received 0 bytes.
08:54:51: WARNING: Remote Cache: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/ac/34715ba2482362a34320130322fcb06c5f753fa518b6037acb723ace24ab2d8b&#39; timed out. Received 0 bytes.
08:55:51: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/pierrec/lz4/v4/internal/xxh32/BUILD.bazel:3:11: Validating nogo output for //vendor/github.com/pierrec/lz4/v4/internal/xxh32:go_default_library failed: Exec failed due to IOException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
08:55:51: INFO: Elapsed time: 125.098s, Critical Path: 120.65s
08:55:51: INFO: 438 processes: 426 remote cache hit, 12 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> bazel remote cache connection timeout (5x / 3.60%) </summary>

<hr/>

**5x**: _2026-03-26 08:52:28 &#43;0000 UTC_: <code>15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037090115150417920#1:build-log.txt%3A1004)
<details>
<summary>all...</summary>

* _2026-03-26 08:53:02 &#43;0000 UTC_: <code>12:20:45: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037090114382860288#1:build-log.txt%3A1122)
<details><summary>context</summary>
<pre>12:20:45: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080
12:20:45: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080
12:20:45: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080
12:20:45: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080
{&#34;component&#34;:&#34;entrypoint&#34;,&#34;file&#34;:&#34;sigs.k8s.io/prow/pkg/entrypoint/run.go:169&#34;,&#34;func&#34;:&#34;sigs.k8s.io/prow/pkg/entrypoint.Options.ExecuteProcess&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Process did not finish before 4h0m0s timeout&#34;,&#34;severity&#34;:&#34;error&#34;,&#34;time&#34;:&#34;2026-03-26T12:53:09Z&#34;}
{&#34;component&#34;:&#34;entrypoint&#34;,&#34;file&#34;:&#34;sigs.k8s.io/prow/pkg/entrypoint/run.go:267&#34;,&#34;func&#34;:&#34;sigs.k8s.io/prow/pkg/entrypoint.gracefullyTerminate&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Process did not exit before 5m0s grace period&#34;,&#34;severity&#34;:&#34;error&#34;,&#34;time&#34;:&#34;2026-03-26T12:58:09Z&#34;}
</pre>
</details>


* _2026-03-26 08:52:28 &#43;0000 UTC_: <code>15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037090115150417920#1:build-log.txt%3A1004)
<details><summary>context</summary>
<pre>15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080
15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080
15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080
15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080
{&#34;component&#34;:&#34;entrypoint&#34;,&#34;file&#34;:&#34;sigs.k8s.io/prow/pkg/entrypoint/run.go:169&#34;,&#34;func&#34;:&#34;sigs.k8s.io/prow/pkg/entrypoint.Options.ExecuteProcess&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Process did not finish before 7h0m0s timeout&#34;,&#34;severity&#34;:&#34;error&#34;,&#34;time&#34;:&#34;2026-03-26T15:52:30Z&#34;}
{&#34;component&#34;:&#34;entrypoint&#34;,&#34;file&#34;:&#34;sigs.k8s.io/prow/pkg/entrypoint/run.go:267&#34;,&#34;func&#34;:&#34;sigs.k8s.io/prow/pkg/entrypoint.gracefullyTerminate&#34;,&#34;level&#34;:&#34;error&#34;,&#34;msg&#34;:&#34;Process did not exit before 5m0s grace period&#34;,&#34;severity&#34;:&#34;error&#34;,&#34;time&#34;:&#34;2026-03-26T15:57:30Z&#34;}
</pre>
</details>


* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>09:02:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037088890476564480#1:build-log.txt%3A2362)
<details><summary>context</summary>
<pre>09:02:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
09:02:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
09:02:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
09:02:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
09:02:52: INFO: Elapsed time: 563.252s, Critical Path: 542.14s
09:02:52: INFO: 2858 processes: 2827 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>09:00:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037088895111270400#1:build-log.txt%3A1119)
<details><summary>context</summary>
<pre>09:00:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
09:00:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
09:00:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
09:00:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
09:00:52: INFO: Elapsed time: 468.685s, Critical Path: 433.52s
09:00:52: INFO: 4931 processes: 4900 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-03-26 08:47:24 &#43;0000 UTC_: <code>08:59:51: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037088890786942976#1:build-log.txt%3A971)
<details><summary>context</summary>
<pre>08:59:51: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
08:59:51: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
08:59:51: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
08:59:51: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080
08:59:51: INFO: Elapsed time: 410.413s, Critical Path: 371.97s
08:59:51: INFO: 6190 processes: 6159 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (2x / 1.44%) </summary>

<hr/>

**2x**: _2026-03-29 10:49:43 &#43;0000 UTC_: <code>10:53:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17227/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2038206823877578752#1:build-log.txt%3A309)
<details>
<summary>all...</summary>

* _2026-03-29 10:49:43 &#43;0000 UTC_: <code>10:53:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17227/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2038206823877578752#1:build-log.txt%3A309)
<details><summary>context</summary>
<pre>10:53:31: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted:
10:53:31: INFO: Elapsed time: 0.330s
10:53:31: INFO: 0 processes.
10:53:31: ERROR: Build failed. Not running target
make: *** [Makefile:37: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


* _2026-03-26 03:48:42 &#43;0000 UTC_: <code>04:06:36: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037013641730985984#1:build-log.txt%3A4356)
<details><summary>context</summary>
<pre>04:06:36: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted:
04:06:36: INFO: Elapsed time: 0.537s
04:06:36: INFO: 0 processes.
04:06:36: ERROR: Build failed. Not running target
04:06:38: &#43; rm -f /tmp/kubevirt.deploy.AsqX
make: *** [Makefile:188: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 1.44%) </summary>

<hr/>

**1x**: _2026-03-26 06:47:21 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2037058711163244544#1:build-log.txt%3A5065)
<details>
<summary>all...</summary>

* _2026-03-26 06:47:21 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2037058711163244544#1:build-log.txt%3A5065)
<details><summary>context</summary>
<pre>&#43; set &#43;x

================================
ERROR: Found panic in test output
Files:
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2037058711163244544/artifacts/pods/0_kubevirt_virt-handler-kqhcc-virt-handler.log
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2037058711163244544/artifacts/pods/0_kubevirt_virt-handler-kqhcc-virt-handler_previous.log</pre>
</details>


</details>

<hr/>

**1x**: _2026-03-25 11:11:30 &#43;0000 UTC_: <code>11:21:44: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17069/pull-kubevirt-e2e-kind-sriov/2036723122534617088#1:build-log.txt%3A1464)
<details>
<summary>all...</summary>

* _2026-03-25 11:11:30 &#43;0000 UTC_: <code>11:21:44: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17069/pull-kubevirt-e2e-kind-sriov/2036723122534617088#1:build-log.txt%3A1464)
<details><summary>context</summary>
<pre>11:21:44: &#43;&#43; for node in ${master_nodes[@]}
11:21:44: &#43;&#43; _kubectl taint nodes sriov-control-plane node-role.kubernetes.io/master:NoSchedule-
11:21:44: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubeconfig taint nodes sriov-control-plane node-role.kubernetes.io/master:NoSchedule-
11:21:44: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
11:21:44: &#43;&#43; _kubectl taint nodes sriov-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
11:21:44: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-sriov/.kubeconfig taint nodes sriov-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
11:21:44: node/sriov-control-plane untainted</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> bazel remote cache download timeout (1x / 0.72%) </summary>

<hr/>

**1x**: _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:54:51: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportproxy/BUILD.bazel:27:10: GoLink cmd/virt-exportproxy/virt-exportproxy_/virt-exportproxy failed: Exec failed due to IOException: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/cas/1c6882fc7400bd46e51aa4130768869fb62c82d129bba7bf916ca68ab6a2ac1c&#39; timed out. Received 21508688 of 24245460 bytes.</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-network/2037088891227344896#1:build-log.txt%3A466)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:54:51: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportproxy/BUILD.bazel:27:10: GoLink cmd/virt-exportproxy/virt-exportproxy_/virt-exportproxy failed: Exec failed due to IOException: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/cas/1c6882fc7400bd46e51aa4130768869fb62c82d129bba7bf916ca68ab6a2ac1c&#39; timed out. Received 21508688 of 24245460 bytes.</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-network/2037088891227344896#1:build-log.txt%3A466)
<details><summary>context</summary>
<pre>08:53:23: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:54:51: WARNING: Remote Cache: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/ac/59efdb9893882a3aaf96e05a77690725741eae117fec4570da5cb0502ce25e86&#39; timed out. Received 0 bytes.
08:54:51: WARNING: Remote Cache: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/ac/eaa74f1d3a46ef47cf84169e12eaf41c6bdf9edae1755cf53347c4d95273647e&#39; timed out. Received 0 bytes.
08:54:51: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportproxy/BUILD.bazel:27:10: GoLink cmd/virt-exportproxy/virt-exportproxy_/virt-exportproxy failed: Exec failed due to IOException: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/cas/1c6882fc7400bd46e51aa4130768869fb62c82d129bba7bf916ca68ab6a2ac1c&#39; timed out. Received 21508688 of 24245460 bytes.
08:54:51: INFO: Elapsed time: 108.432s, Critical Path: 80.10s
08:54:51: INFO: 4676 processes: 4645 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (from secondary snippet) (1x / 0.72%) </summary>

<hr/>

**1x**: _2026-03-26 08:34:18 &#43;0000 UTC_: <code>08:57:20: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-k8s-1.34-sig-network/2037085594298880000#1:build-log.txt%3A4832)
<details>
<summary>all...</summary>

* _2026-03-26 08:34:18 &#43;0000 UTC_: <code>08:57:20: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-k8s-1.34-sig-network/2037085594298880000#1:build-log.txt%3A4832)
<details><summary>context</summary>
<pre>08:57:20: java.net.UnknownHostException: bazel-cache.kubevirt-prow
08:57:20: INFO: Elapsed time: 0.497s, Critical Path: 0.06s
08:57:20: INFO: 6 processes: 6 internal.
08:57:20: ERROR: Build failed. Not running target
make: *** [Makefile:28: bazel-build-functests] Error 1
&#43; ret=2
&#43; check_for_panics</pre>
</details>


</details>

<hr/>
</details>

### internal (7x / 5.04%)

<details>
<summary> kind cluster creation failure (6x / 4.32%) </summary>

<hr/>

**6x**: _2026-03-25 17:05:27 &#43;0000 UTC_: <code>17:12:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17291/pull-kubevirt-e2e-kind-1.34-sev-1.7/2036851449563975680#1:build-log.txt%3A609)
<details>
<summary>all...</summary>

* _2026-03-26 09:22:19 &#43;0000 UTC_: <code>09:29:01: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037097667556806656#1:build-log.txt%3A816)
<details><summary>context</summary>
<pre>09:28:58:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
09:28:58:  • Preparing nodes 📦 📦   ...
09:29:01:  ✗ Preparing nodes 📦 📦
09:29:01: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
09:29:01:
09:29:01: Stack Trace:
09:29:01: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


* _2026-03-26 08:47:18 &#43;0000 UTC_: <code>08:54:22: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-kind-1.34-sev/2037088890237489152#1:build-log.txt%3A709)
<details><summary>context</summary>
<pre>08:54:19:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
08:54:19:  • Preparing nodes 📦 📦   ...
08:54:22:  ✗ Preparing nodes 📦 📦
08:54:22: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
08:54:22:
08:54:22: Stack Trace:
08:54:22: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


* _2026-03-26 01:49:22 &#43;0000 UTC_: <code>01:55:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2036983572786581504#1:build-log.txt%3A789)
<details><summary>context</summary>
<pre>01:55:40:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
01:55:40:  • Preparing nodes 📦 📦   ...
01:55:41:  ✗ Preparing nodes 📦 📦
01:55:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
01:55:41:
01:55:41: Stack Trace:
01:55:41: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


* _2026-03-25 22:46:16 &#43;0000 UTC_: <code>22:52:30: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036937635984838656#1:build-log.txt%3A697)
<details><summary>context</summary>
<pre>22:52:28:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
22:52:28:  • Preparing nodes 📦 📦   ...
22:52:30:  ✗ Preparing nodes 📦 📦
22:52:30: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
22:52:30:
22:52:30: Stack Trace:
22:52:30: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


* _2026-03-25 17:07:07 &#43;0000 UTC_: <code>17:15:35: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17290/pull-kubevirt-e2e-kind-1.34-sev-1.8/2036852207432765440#1:build-log.txt%3A834)
<details><summary>context</summary>
<pre>17:15:04:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
17:15:04:  • Preparing nodes 📦 📦   ...
17:15:35:  ✗ Preparing nodes 📦 📦
17:15:35: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
17:15:35:
17:15:35: Stack Trace:
17:15:35: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


* _2026-03-25 17:05:27 &#43;0000 UTC_: <code>17:12:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17291/pull-kubevirt-e2e-kind-1.34-sev-1.7/2036851449563975680#1:build-log.txt%3A609)
<details><summary>context</summary>
<pre>17:12:39:  ✓ Ensuring node image (kindest/node:v1.34.0) 🖼
17:12:39:  • Preparing nodes 📦 📦   ...
17:12:41:  ✗ Preparing nodes 📦 📦
17:12:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
17:12:41:
17:12:41: Stack Trace:
17:12:41: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> control plane startup failure (1x / 0.72%) </summary>

<hr/>

**1x**: _2026-03-26 03:48:17 &#43;0000 UTC_: <code>04:03:20: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037013641584185344#1:build-log.txt%3A2850)
<details>
<summary>all...</summary>

* _2026-03-26 03:48:17 &#43;0000 UTC_: <code>04:03:20: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037013641584185344#1:build-log.txt%3A2850)
<details><summary>context</summary>
<pre>04:03:20: 	Once you have found the failing container, you can inspect its logs with:
04:03:20: 	- &#39;crictl --runtime-endpoint unix:///run/containerd/containerd.sock logs CONTAINERID&#39;
04:03:20:
04:03:20: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]
04:03:20: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
04:03:20: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262
04:03:20: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).visitAll</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (8x / 5.76%)

<details>
<summary> no error snippets (8x / 5.76%) </summary>

<hr/>

**1x**: _2026-03-26 08:52:44 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2037090112856133632)
<details>
<summary>all...</summary>

* _2026-03-26 08:52:44 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2037090112856133632)

</details>

<hr/>

**1x**: _2026-03-26 08:50:21 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037089601964740608)
<details>
<summary>all...</summary>

* _2026-03-26 08:50:21 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037089601964740608)

</details>

<hr/>

**1x**: _2026-03-26 07:47:29 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-kind-1.34-sev/2037073830949687296)
<details>
<summary>all...</summary>

* _2026-03-26 07:47:29 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-kind-1.34-sev/2037073830949687296)

</details>

<hr/>

**1x**: _2026-03-26 05:04:21 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037032780927340544)
<details>
<summary>all...</summary>

* _2026-03-26 05:04:21 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037032780927340544)

</details>

<hr/>

**1x**: _2026-03-26 02:47:17 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036998292079906816)
<details>
<summary>all...</summary>

* _2026-03-26 02:47:17 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036998292079906816)

</details>

<hr/>

**1x**: _2026-03-26 00:46:17 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-kind-1.34-sev/2036967845056745472)
<details>
<summary>all...</summary>

* _2026-03-26 00:46:17 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-kind-1.34-sev/2036967845056745472)

</details>

<hr/>

**1x**: _2026-03-25 20:55:09 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2036909651676631040)
<details>
<summary>all...</summary>

* _2026-03-25 20:55:09 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2036909651676631040)

</details>

<hr/>

**1x**: _2026-03-25 18:46:20 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17207/pull-kubevirt-e2e-kind-1.34-sev/2036877259465297920)
<details>
<summary>all...</summary>

* _2026-03-25 18:46:20 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17207/pull-kubevirt-e2e-kind-1.34-sev/2036877259465297920)

</details>

<hr/>
</details>

### pr-build (2x / 1.44%)

<details>
<summary> panic detected in test output (2x / 1.44%) </summary>

<hr/>

**2x**: _2026-03-26 07:35:38 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037070842281594880#1:build-log.txt%3A5081)
<details>
<summary>all...</summary>

* _2026-03-26 08:34:11 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-kind-1.34-sev/2037085594152079360#1:build-log.txt%3A5063)
<details><summary>context</summary>
<pre>&#43; set &#43;x

================================
ERROR: Found panic in test output
Files:
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-kind-1.34-sev/2037085594152079360/artifacts/pods/0_kubevirt_virt-handler-cv29m-virt-handler.log
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-kind-1.34-sev/2037085594152079360/artifacts/pods/0_kubevirt_virt-handler-cv29m-virt-handler_previous.log</pre>
</details>


* _2026-03-26 07:35:38 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037070842281594880#1:build-log.txt%3A5081)
<details><summary>context</summary>
<pre>&#43; set &#43;x

================================
ERROR: Found panic in test output
Files:
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037070842281594880/artifacts/pods/0_kubevirt_virt-handler-fntcb-virt-handler.log
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037070842281594880/artifacts/pods/0_kubevirt_virt-handler-fntcb-virt-handler_previous.log</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### release-1.7 (2x / 1.44%)


#### external (1x / 50.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-03-29 10:49:43 &#43;0000 UTC_: <code>10:53:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17227/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2038206823877578752#1:build-log.txt%3A309)
<details>
<summary>all...</summary>

* _2026-03-29 10:49:43 &#43;0000 UTC_: <code>10:53:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17227/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2038206823877578752#1:build-log.txt%3A309)

</details>

<hr/>
</details>

#### internal (1x / 50.00%)

<details>
<summary> kind cluster creation failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-03-25 17:05:27 &#43;0000 UTC_: <code>17:12:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17291/pull-kubevirt-e2e-kind-1.34-sev-1.7/2036851449563975680#1:build-log.txt%3A609)
<details>
<summary>all...</summary>

* _2026-03-25 17:05:27 &#43;0000 UTC_: <code>17:12:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17291/pull-kubevirt-e2e-kind-1.34-sev-1.7/2036851449563975680#1:build-log.txt%3A609)

</details>

<hr/>
</details>

### main (136x / 97.84%)


#### external (121x / 88.97%)

<details>
<summary> bazel remote cache blob fetch failure (104x / 76.47%) </summary>

<hr/>

**21x**: _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:53:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037088891659358208#1:build-log.txt%3A518)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:34 &#43;0000 UTC_: <code>08:55:31: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-windows2016/2037088890090688512#1:build-log.txt%3A489)

* _2026-03-26 08:47:27 &#43;0000 UTC_: <code>08:56:00: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-network/2037088890619170816#1:build-log.txt%3A516)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:52:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037088891520946176#1:build-log.txt%3A515)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:53:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037088891659358208#1:build-log.txt%3A518)

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:53:12: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2037088890333958144#1:build-log.txt%3A516)

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:52:38: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037088891369951232#1:build-log.txt%3A524)

* _2026-03-26 07:56:52 &#43;0000 UTC_: <code>08:01:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037076170310160384#1:build-log.txt%3A511)

* _2026-03-26 07:47:39 &#43;0000 UTC_: <code>07:53:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037073833353023488#1:build-log.txt%3A523)

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:35: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-windows2016/2037073830886772736#1:build-log.txt%3A549)

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:51:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-network/2037073834191884288#1:build-log.txt%3A524)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:51:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-network/2037073831499141120#1:build-log.txt%3A540)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037073835894771712#1:build-log.txt%3A511)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037073831637553152#1:build-log.txt%3A529)

* _2026-03-26 06:47:31 &#43;0000 UTC_: <code>06:51:22: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2037058711255519232#1:build-log.txt%3A424)

* _2026-03-25 22:46:26 &#43;0000 UTC_: <code>22:51:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036937636135833600#1:build-log.txt%3A503)

* _2026-03-25 22:30:47 &#43;0000 UTC_: <code>22:34:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036892368967307264#1:build-log.txt%3A512)

* _2026-03-25 22:29:24 &#43;0000 UTC_: <code>22:33:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036892368111669248#1:build-log.txt%3A461)

* _2026-03-25 22:28:59 &#43;0000 UTC_: <code>22:35:02: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036892367205699584#1:build-log.txt%3A517)

* _2026-03-25 22:28:15 &#43;0000 UTC_: <code>22:33:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-network/2036892367029538816#1:build-log.txt%3A529)

* _2026-03-25 21:43:49 &#43;0000 UTC_: <code>21:50:13: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036892369781002240#1:build-log.txt%3A459)

* _2026-03-25 20:39:55 &#43;0000 UTC_: <code>20:44:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036904750133284864#1:build-log.txt%3A445)

</details>

<hr/>

**7x**: _2026-03-26 07:47:48 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037073832539328512#1:build-log.txt%3A502)
<details>
<summary>all...</summary>

* _2026-03-26 07:47:48 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037073832539328512#1:build-log.txt%3A502)

* _2026-03-26 07:47:47 &#43;0000 UTC_: <code>07:53:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037073831079710720#1:build-log.txt%3A595)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2037073831012601856#1:build-log.txt%3A540)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037073831281037312#1:build-log.txt%3A530)

* _2026-03-26 05:47:21 &#43;0000 UTC_: <code>05:51:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17276/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037043575102902272#1:build-log.txt%3A550)

* _2026-03-26 04:47:33 &#43;0000 UTC_: <code>04:53:53: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037028524841242624#1:build-log.txt%3A499)

* _2026-03-25 20:46:36 &#43;0000 UTC_: <code>20:51:59: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036906547623235584#1:build-log.txt%3A519)

</details>

<hr/>

**32x**: _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036967849838252032#1:build-log.txt%3A446)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:52:26: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037088894213689344#1:build-log.txt%3A518)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:55:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037088892464664576#1:build-log.txt%3A517)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-operator/2037088891055378432#1:build-log.txt%3A482)

* _2026-03-26 07:47:40 &#43;0000 UTC_: <code>07:51:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-operator/2037073831419449344#1:build-log.txt%3A529)

* _2026-03-26 07:47:39 &#43;0000 UTC_: <code>07:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037073831914377216#1:build-log.txt%3A525)

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-network/2037073831209734144#1:build-log.txt%3A540)

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037073836737826816#1:build-log.txt%3A529)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-compute/2037073831356534784#1:build-log.txt%3A460)

* _2026-03-26 07:47:36 &#43;0000 UTC_: <code>07:54:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037073835026550784#1:build-log.txt%3A532)

* _2026-03-26 02:47:25 &#43;0000 UTC_: <code>02:52:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036998292239290368#1:build-log.txt%3A481)

* _2026-03-26 00:46:37 &#43;0000 UTC_: <code>00:55:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036967851490807808#1:build-log.txt%3A510)

* _2026-03-26 00:46:32 &#43;0000 UTC_: <code>00:54:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-network/2036967848978419712#1:build-log.txt%3A451)

* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036967848152141824#1:build-log.txt%3A478)

* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036967849838252032#1:build-log.txt%3A446)

* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036967847292309504#1:build-log.txt%3A508)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:51:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-network/2036967846151458816#1:build-log.txt%3A521)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:52:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036967845736222720#1:build-log.txt%3A484)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:51:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036967846004658176#1:build-log.txt%3A506)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:52:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036967850643558400#1:build-log.txt%3A505)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:51:42: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036967846302453760#1:build-log.txt%3A509)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:52:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036967845358735360#1:build-log.txt%3A588)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:53:15: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-network/2036967845614587904#1:build-log.txt%3A520)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:55:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036967845870440448#1:build-log.txt%3A504)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:54:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036967845245489152#1:build-log.txt%3A517)

* _2026-03-26 00:46:25 &#43;0000 UTC_: <code>00:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-monitoring/2036967845480370176#1:build-log.txt%3A508)

* _2026-03-26 00:46:25 &#43;0000 UTC_: <code>00:52:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036967846478614528#1:build-log.txt%3A510)

* _2026-03-26 00:46:23 &#43;0000 UTC_: <code>00:54:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-windows2016/2036967844863807488#1:build-log.txt%3A520)

* _2026-03-25 23:46:23 &#43;0000 UTC_: <code>23:50:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036952734418079744#1:build-log.txt%3A498)

* _2026-03-25 22:27:03 &#43;0000 UTC_: <code>22:32:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036892366903709696#1:build-log.txt%3A503)

* _2026-03-25 21:47:42 &#43;0000 UTC_: <code>21:52:56: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036892374516371456#1:build-log.txt%3A439)

* _2026-03-25 21:47:10 &#43;0000 UTC_: <code>21:51:33: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17276/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036922551527018496#1:build-log.txt%3A509)

* _2026-03-25 21:46:12 &#43;0000 UTC_: <code>21:51:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-network/2036892370359816192#1:build-log.txt%3A487)

</details>

<hr/>

**15x**: _2026-03-26 05:04:28 &#43;0000 UTC_: <code>05:08:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037032781342576640#1:build-log.txt%3A534)
<details>
<summary>all...</summary>

* _2026-03-26 05:04:37 &#43;0000 UTC_: <code>05:09:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-network/2037032781271273472#1:build-log.txt%3A572)

* _2026-03-26 05:04:29 &#43;0000 UTC_: <code>05:08:54: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037032781178998784#1:build-log.txt%3A529)

* _2026-03-26 05:04:29 &#43;0000 UTC_: <code>05:08:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037032781422268416#1:build-log.txt%3A541)

* _2026-03-26 05:04:28 &#43;0000 UTC_: <code>05:08:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037032781342576640#1:build-log.txt%3A534)

* _2026-03-26 05:04:26 &#43;0000 UTC_: <code>05:10:23: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037032781506154496#1:build-log.txt%3A522)

* _2026-03-26 05:04:26 &#43;0000 UTC_: <code>05:10:15: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037032781090918400#1:build-log.txt%3A632)

* _2026-03-26 01:49:19 &#43;0000 UTC_: <code>01:59:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036983575923920896#1:build-log.txt%3A480)

* _2026-03-26 01:49:19 &#43;0000 UTC_: <code>01:59:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-windows2016/2036983572635586560#1:build-log.txt%3A561)

* _2026-03-26 01:49:16 &#43;0000 UTC_: <code>01:59:08: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036983573491224576#1:build-log.txt%3A496)

* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:53:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036983573973569536#1:build-log.txt%3A531)

* _2026-03-26 01:48:56 &#43;0000 UTC_: <code>01:54:56: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-network/2036983573382172672#1:build-log.txt%3A538)

* _2026-03-26 01:48:56 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-network/2036983573851934720#1:build-log.txt%3A538)

* _2026-03-25 21:37:27 &#43;0000 UTC_: <code>21:42:31: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-network/2036909662514712576#1:build-log.txt%3A542)

* _2026-03-25 21:36:01 &#43;0000 UTC_: <code>21:40:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036909661638103040#1:build-log.txt%3A555)

* _2026-03-25 21:35:18 &#43;0000 UTC_: <code>21:41:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036909659150880768#1:build-log.txt%3A648)

</details>

<hr/>

**6x**: _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036983575089254400#1:build-log.txt%3A505)
<details>
<summary>all...</summary>

* _2026-03-26 01:49:28 &#43;0000 UTC_: <code>01:59:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036983573042434048#1:build-log.txt%3A534)

* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036983574225227776#1:build-log.txt%3A534)

* _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036983575089254400#1:build-log.txt%3A505)

* _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:19: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036983573713522688#1:build-log.txt%3A535)

* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:55:04: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036983578432114688#1:build-log.txt%3A527)

* _2026-03-26 01:48:57 &#43;0000 UTC_: <code>01:55:06: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036983573587693568#1:build-log.txt%3A540)

</details>

<hr/>

**8x**: _2026-03-25 19:56:13 &#43;0000 UTC_: <code>20:02:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036894133938819072#1:build-log.txt%3A522)
<details>
<summary>all...</summary>

* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036983579300335616#1:build-log.txt%3A541)

* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036983573143097344#1:build-log.txt%3A615)

* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:56:07: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-network/2036983576750198784#1:build-log.txt%3A549)

* _2026-03-26 01:48:57 &#43;0000 UTC_: <code>01:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036983577610031104#1:build-log.txt%3A481)

* _2026-03-25 21:43:46 &#43;0000 UTC_: <code>21:50:52: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036909664993546240#1:build-log.txt%3A565)

* _2026-03-25 21:08:47 &#43;0000 UTC_: <code>21:13:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036912517216735232#1:build-log.txt%3A557)

* _2026-03-25 19:56:46 &#43;0000 UTC_: <code>20:08:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036892366803046400#1:build-log.txt%3A506)

* _2026-03-25 19:56:13 &#43;0000 UTC_: <code>20:02:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036894133938819072#1:build-log.txt%3A522)

</details>

<hr/>

**4x**: _2026-03-25 19:46:36 &#43;0000 UTC_: <code>19:51:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-windows2016/2036892366194872320#1:build-log.txt%3A512)
<details>
<summary>all...</summary>

* _2026-03-25 19:52:50 &#43;0000 UTC_: <code>19:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-network/2036892366677217280#1:build-log.txt%3A529)

* _2026-03-25 19:52:47 &#43;0000 UTC_: <code>19:56:55: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036892366597525504#1:build-log.txt%3A586)

* _2026-03-25 19:49:08 &#43;0000 UTC_: <code>19:53:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036892366492667904#1:build-log.txt%3A532)

* _2026-03-25 19:46:36 &#43;0000 UTC_: <code>19:51:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-windows2016/2036892366194872320#1:build-log.txt%3A512)

</details>

<hr/>

**7x**: _2026-03-25 17:27:42 &#43;0000 UTC_: <code>19:27:04: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036873472784732160#1:build-log.txt%3A385)
<details>
<summary>all...</summary>

* _2026-03-25 22:08:13 &#43;0000 UTC_: <code>22:12:17: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036857991180849152#1:build-log.txt%3A346)

* _2026-03-25 21:08:14 &#43;0000 UTC_: <code>21:12:54: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036862229206274048#1:build-log.txt%3A433)

* _2026-03-25 19:08:48 &#43;0000 UTC_: <code>19:16:10: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036882243925839872#1:build-log.txt%3A394)

* _2026-03-25 18:30:52 &#43;0000 UTC_: <code>18:36:56: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036873262981451776#1:build-log.txt%3A500)

* _2026-03-25 18:30:51 &#43;0000 UTC_: <code>18:37:58: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-network/2036873262641713152#1:build-log.txt%3A469)

* _2026-03-25 18:25:43 &#43;0000 UTC_: <code>18:30:50: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-windows2016/2036871960771694592#1:build-log.txt%3A468)

* _2026-03-25 17:27:42 &#43;0000 UTC_: <code>19:27:04: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036873472784732160#1:build-log.txt%3A385)

</details>

<hr/>

**4x**: _2026-03-25 17:38:42 &#43;0000 UTC_: <code>17:42:45: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036860131525791744#1:build-log.txt%3A423)
<details>
<summary>all...</summary>

* _2026-03-25 21:48:21 &#43;0000 UTC_: <code>21:51:40: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036892400621719552#1:build-log.txt%3A438)

* _2026-03-25 17:45:13 &#43;0000 UTC_: <code>17:48:29: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-network/2036861683703484416#1:build-log.txt%3A432)

* _2026-03-25 17:38:42 &#43;0000 UTC_: <code>17:42:45: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036860131525791744#1:build-log.txt%3A423)

* _2026-03-25 17:29:54 &#43;0000 UTC_: <code>17:36:31: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036857570827702272#1:build-log.txt%3A458)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (7x / 5.15%) </summary>

<hr/>

**5x**: _2026-03-26 08:52:44 &#43;0000 UTC_: <code>08:57:27: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 33 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037090118291951616#1:build-log.txt%3A439)
<details>
<summary>all...</summary>

* _2026-03-26 08:52:44 &#43;0000 UTC_: <code>08:57:27: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 33 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037090118291951616#1:build-log.txt%3A439)

* _2026-03-26 08:52:43 &#43;0000 UTC_: <code>08:57:26: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-launcher/libvirt-hook-client/BUILD.bazel:10:10: GoCompilePkg cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client.a failed: Exec failed due to IOException: 351 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-network/2037090113296535552#1:build-log.txt%3A760)

* _2026-03-26 08:52:30 &#43;0000 UTC_: <code>08:58:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tools/csv-generator/BUILD.bazel:16:10 Validating nogo output for //tools/csv-generator:csv-generator failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037090114504495104#1:build-log.txt%3A490)

* _2026-03-26 08:52:28 &#43;0000 UTC_: <code>08:57:24: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 73 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-network/2037090114290585600#1:build-log.txt%3A725)

* _2026-03-26 08:52:20 &#43;0000 UTC_: <code>08:57:23: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportserver/BUILD.bazel:20:10: GoLink cmd/virt-exportserver/virt-exportserver_/virt-exportserver failed: Exec failed due to IOException: 161 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037090113539805184#1:build-log.txt%3A479)

</details>

<hr/>

**1x**: _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:58:56: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/virt-handler/vsock/system/BUILD.bazel:3:11 Validating nogo output for //pkg/virt-handler/vsock/system:go_default_library failed: Exec failed due to IOException: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037088895878828032#1:build-log.txt%3A883)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:58:56: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/virt-handler/vsock/system/BUILD.bazel:3:11 Validating nogo output for //pkg/virt-handler/vsock/system:go_default_library failed: Exec failed due to IOException: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037088895878828032#1:build-log.txt%3A883)

</details>

<hr/>

**1x**: _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:55:51: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/pierrec/lz4/v4/internal/xxh32/BUILD.bazel:3:11: Validating nogo output for //vendor/github.com/pierrec/lz4/v4/internal/xxh32:go_default_library failed: Exec failed due to IOException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-network/2037088893345468416#1:build-log.txt%3A467)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:55:51: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/pierrec/lz4/v4/internal/xxh32/BUILD.bazel:3:11: Validating nogo output for //vendor/github.com/pierrec/lz4/v4/internal/xxh32:go_default_library failed: Exec failed due to IOException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-network/2037088893345468416#1:build-log.txt%3A467)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache connection timeout (5x / 3.68%) </summary>

<hr/>

**5x**: _2026-03-26 08:52:28 &#43;0000 UTC_: <code>15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037090115150417920#1:build-log.txt%3A1004)
<details>
<summary>all...</summary>

* _2026-03-26 08:53:02 &#43;0000 UTC_: <code>12:20:45: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037090114382860288#1:build-log.txt%3A1122)

* _2026-03-26 08:52:28 &#43;0000 UTC_: <code>15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037090115150417920#1:build-log.txt%3A1004)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>09:02:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037088890476564480#1:build-log.txt%3A2362)

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>09:00:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037088895111270400#1:build-log.txt%3A1119)

* _2026-03-26 08:47:24 &#43;0000 UTC_: <code>08:59:51: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037088890786942976#1:build-log.txt%3A971)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 1.47%) </summary>

<hr/>

**1x**: _2026-03-26 06:47:21 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2037058711163244544#1:build-log.txt%3A5065)
<details>
<summary>all...</summary>

* _2026-03-26 06:47:21 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2037058711163244544#1:build-log.txt%3A5065)

</details>

<hr/>

**1x**: _2026-03-25 11:11:30 &#43;0000 UTC_: <code>11:21:44: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17069/pull-kubevirt-e2e-kind-sriov/2036723122534617088#1:build-log.txt%3A1464)
<details>
<summary>all...</summary>

* _2026-03-25 11:11:30 &#43;0000 UTC_: <code>11:21:44: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17069/pull-kubevirt-e2e-kind-sriov/2036723122534617088#1:build-log.txt%3A1464)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (from secondary snippet) (1x / 0.74%) </summary>

<hr/>

**1x**: _2026-03-26 08:34:18 &#43;0000 UTC_: <code>08:57:20: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-k8s-1.34-sig-network/2037085594298880000#1:build-log.txt%3A4832)
<details>
<summary>all...</summary>

* _2026-03-26 08:34:18 &#43;0000 UTC_: <code>08:57:20: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-k8s-1.34-sig-network/2037085594298880000#1:build-log.txt%3A4832)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache download timeout (1x / 0.74%) </summary>

<hr/>

**1x**: _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:54:51: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportproxy/BUILD.bazel:27:10: GoLink cmd/virt-exportproxy/virt-exportproxy_/virt-exportproxy failed: Exec failed due to IOException: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/cas/1c6882fc7400bd46e51aa4130768869fb62c82d129bba7bf916ca68ab6a2ac1c&#39; timed out. Received 21508688 of 24245460 bytes.</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-network/2037088891227344896#1:build-log.txt%3A466)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:54:51: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportproxy/BUILD.bazel:27:10: GoLink cmd/virt-exportproxy/virt-exportproxy_/virt-exportproxy failed: Exec failed due to IOException: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/cas/1c6882fc7400bd46e51aa4130768869fb62c82d129bba7bf916ca68ab6a2ac1c&#39; timed out. Received 21508688 of 24245460 bytes.</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-network/2037088891227344896#1:build-log.txt%3A466)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 0.74%) </summary>

<hr/>

**1x**: _2026-03-26 03:48:42 &#43;0000 UTC_: <code>04:06:36: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037013641730985984#1:build-log.txt%3A4356)
<details>
<summary>all...</summary>

* _2026-03-26 03:48:42 &#43;0000 UTC_: <code>04:06:36: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037013641730985984#1:build-log.txt%3A4356)

</details>

<hr/>
</details>

#### internal (5x / 3.68%)

<details>
<summary> kind cluster creation failure (4x / 2.94%) </summary>

<hr/>

**4x**: _2026-03-26 09:22:19 &#43;0000 UTC_: <code>09:29:01: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037097667556806656#1:build-log.txt%3A816)
<details>
<summary>all...</summary>

* _2026-03-26 09:22:19 &#43;0000 UTC_: <code>09:29:01: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037097667556806656#1:build-log.txt%3A816)

* _2026-03-26 08:47:18 &#43;0000 UTC_: <code>08:54:22: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-kind-1.34-sev/2037088890237489152#1:build-log.txt%3A709)

* _2026-03-26 01:49:22 &#43;0000 UTC_: <code>01:55:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2036983572786581504#1:build-log.txt%3A789)

* _2026-03-25 22:46:16 &#43;0000 UTC_: <code>22:52:30: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036937635984838656#1:build-log.txt%3A697)

</details>

<hr/>
</details>
<details>
<summary> control plane startup failure (1x / 0.74%) </summary>

<hr/>

**1x**: _2026-03-26 03:48:17 &#43;0000 UTC_: <code>04:03:20: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037013641584185344#1:build-log.txt%3A2850)
<details>
<summary>all...</summary>

* _2026-03-26 03:48:17 &#43;0000 UTC_: <code>04:03:20: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037013641584185344#1:build-log.txt%3A2850)

</details>

<hr/>
</details>

#### needs-investigation (8x / 5.88%)

<details>
<summary> no error snippets (8x / 5.88%) </summary>

<hr/>

**1x**: _2026-03-26 08:52:44 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2037090112856133632)
<details>
<summary>all...</summary>

* _2026-03-26 08:52:44 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2037090112856133632)

</details>

<hr/>

**1x**: _2026-03-26 08:50:21 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037089601964740608)
<details>
<summary>all...</summary>

* _2026-03-26 08:50:21 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037089601964740608)

</details>

<hr/>

**1x**: _2026-03-26 07:47:29 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-kind-1.34-sev/2037073830949687296)
<details>
<summary>all...</summary>

* _2026-03-26 07:47:29 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-kind-1.34-sev/2037073830949687296)

</details>

<hr/>

**1x**: _2026-03-26 05:04:21 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037032780927340544)
<details>
<summary>all...</summary>

* _2026-03-26 05:04:21 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037032780927340544)

</details>

<hr/>

**1x**: _2026-03-26 02:47:17 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036998292079906816)
<details>
<summary>all...</summary>

* _2026-03-26 02:47:17 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036998292079906816)

</details>

<hr/>

**1x**: _2026-03-26 00:46:17 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-kind-1.34-sev/2036967845056745472)
<details>
<summary>all...</summary>

* _2026-03-26 00:46:17 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-kind-1.34-sev/2036967845056745472)

</details>

<hr/>

**1x**: _2026-03-25 20:55:09 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2036909651676631040)
<details>
<summary>all...</summary>

* _2026-03-25 20:55:09 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2036909651676631040)

</details>

<hr/>

**1x**: _2026-03-25 18:46:20 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17207/pull-kubevirt-e2e-kind-1.34-sev/2036877259465297920)
<details>
<summary>all...</summary>

* _2026-03-25 18:46:20 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17207/pull-kubevirt-e2e-kind-1.34-sev/2036877259465297920)

</details>

<hr/>
</details>

#### pr-build (2x / 1.47%)

<details>
<summary> panic detected in test output (2x / 1.47%) </summary>

<hr/>

**2x**: _2026-03-26 07:35:38 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037070842281594880#1:build-log.txt%3A5081)
<details>
<summary>all...</summary>

* _2026-03-26 08:34:11 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-kind-1.34-sev/2037085594152079360#1:build-log.txt%3A5063)

* _2026-03-26 07:35:38 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037070842281594880#1:build-log.txt%3A5081)

</details>

<hr/>
</details>

### release-1.8 (1x / 0.72%)


#### internal (1x / 100.00%)

<details>
<summary> kind cluster creation failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-03-25 17:07:07 &#43;0000 UTC_: <code>17:15:35: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17290/pull-kubevirt-e2e-kind-1.34-sev-1.8/2036852207432765440#1:build-log.txt%3A834)
<details>
<summary>all...</summary>

* _2026-03-25 17:07:07 &#43;0000 UTC_: <code>17:15:35: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17290/pull-kubevirt-e2e-kind-1.34-sev-1.8/2036852207432765440#1:build-log.txt%3A834)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-storage (27x / 19.42%)


#### external (27x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (22x / 81.48%) </summary>

<hr/>

**3x**: _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:52:38: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037088891369951232#1:build-log.txt%3A524)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:52:38: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037088891369951232#1:build-log.txt%3A524)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037073831637553152#1:build-log.txt%3A529)

* _2026-03-25 22:28:59 &#43;0000 UTC_: <code>22:35:02: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036892367205699584#1:build-log.txt%3A517)

</details>

<hr/>

**5x**: _2026-03-26 05:47:21 &#43;0000 UTC_: <code>05:51:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17276/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037043575102902272#1:build-log.txt%3A550)
<details>
<summary>all...</summary>

* _2026-03-26 07:56:52 &#43;0000 UTC_: <code>08:01:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037076170310160384#1:build-log.txt%3A511)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037073831281037312#1:build-log.txt%3A530)

* _2026-03-26 05:47:21 &#43;0000 UTC_: <code>05:51:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17276/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037043575102902272#1:build-log.txt%3A550)

* _2026-03-26 04:47:33 &#43;0000 UTC_: <code>04:53:53: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037028524841242624#1:build-log.txt%3A499)

* _2026-03-25 20:46:36 &#43;0000 UTC_: <code>20:51:59: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036906547623235584#1:build-log.txt%3A519)

</details>

<hr/>

**7x**: _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036967849838252032#1:build-log.txt%3A446)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:52:26: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037088894213689344#1:build-log.txt%3A518)

* _2026-03-26 07:47:36 &#43;0000 UTC_: <code>07:54:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037073835026550784#1:build-log.txt%3A532)

* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036967849838252032#1:build-log.txt%3A446)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:52:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036967845736222720#1:build-log.txt%3A484)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:51:42: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036967846302453760#1:build-log.txt%3A509)

* _2026-03-25 23:46:23 &#43;0000 UTC_: <code>23:50:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17087/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036952734418079744#1:build-log.txt%3A498)

* _2026-03-25 21:47:10 &#43;0000 UTC_: <code>21:51:33: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17276/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036922551527018496#1:build-log.txt%3A509)

</details>

<hr/>

**3x**: _2026-03-26 05:04:28 &#43;0000 UTC_: <code>05:08:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037032781342576640#1:build-log.txt%3A534)
<details>
<summary>all...</summary>

* _2026-03-26 05:04:28 &#43;0000 UTC_: <code>05:08:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037032781342576640#1:build-log.txt%3A534)

* _2026-03-26 01:49:16 &#43;0000 UTC_: <code>01:59:08: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036983573491224576#1:build-log.txt%3A496)

* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:53:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036983573973569536#1:build-log.txt%3A531)

</details>

<hr/>

**1x**: _2026-03-25 21:48:21 &#43;0000 UTC_: <code>21:51:40: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036892400621719552#1:build-log.txt%3A438)
<details>
<summary>all...</summary>

* _2026-03-25 21:48:21 &#43;0000 UTC_: <code>21:51:40: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-storage/2036892400621719552#1:build-log.txt%3A438)

</details>

<hr/>

**2x**: _2026-03-25 19:56:13 &#43;0000 UTC_: <code>20:02:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036894133938819072#1:build-log.txt%3A522)
<details>
<summary>all...</summary>

* _2026-03-26 01:48:57 &#43;0000 UTC_: <code>01:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036983577610031104#1:build-log.txt%3A481)

* _2026-03-25 19:56:13 &#43;0000 UTC_: <code>20:02:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-storage/2036894133938819072#1:build-log.txt%3A522)

</details>

<hr/>

**1x**: _2026-03-25 17:27:42 &#43;0000 UTC_: <code>19:27:04: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036873472784732160#1:build-log.txt%3A385)
<details>
<summary>all...</summary>

* _2026-03-25 17:27:42 &#43;0000 UTC_: <code>19:27:04: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-storage/2036873472784732160#1:build-log.txt%3A385)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (2x / 7.41%) </summary>

<hr/>

**2x**: _2026-03-26 08:52:44 &#43;0000 UTC_: <code>08:57:27: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 33 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037090118291951616#1:build-log.txt%3A439)
<details>
<summary>all...</summary>

* _2026-03-26 08:52:44 &#43;0000 UTC_: <code>08:57:27: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 33 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-storage/2037090118291951616#1:build-log.txt%3A439)

* _2026-03-26 08:52:20 &#43;0000 UTC_: <code>08:57:23: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportserver/BUILD.bazel:20:10: GoLink cmd/virt-exportserver/virt-exportserver_/virt-exportserver failed: Exec failed due to IOException: 161 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037090113539805184#1:build-log.txt%3A479)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache connection timeout (2x / 7.41%) </summary>

<hr/>

**2x**: _2026-03-26 08:53:02 &#43;0000 UTC_: <code>12:20:45: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037090114382860288#1:build-log.txt%3A1122)
<details>
<summary>all...</summary>

* _2026-03-26 08:53:02 &#43;0000 UTC_: <code>12:20:45: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-storage/2037090114382860288#1:build-log.txt%3A1122)

* _2026-03-26 08:47:24 &#43;0000 UTC_: <code>08:59:51: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-storage/2037088890786942976#1:build-log.txt%3A971)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 3.70%) </summary>

<hr/>

**1x**: _2026-03-29 10:49:43 &#43;0000 UTC_: <code>10:53:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17227/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2038206823877578752#1:build-log.txt%3A309)
<details>
<summary>all...</summary>

* _2026-03-29 10:49:43 &#43;0000 UTC_: <code>10:53:31: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17227/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2038206823877578752#1:build-log.txt%3A309)

</details>

<hr/>
</details>

### sig-network (29x / 20.86%)


#### external (29x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (23x / 79.31%) </summary>

<hr/>

**1x**: _2026-03-26 05:04:37 &#43;0000 UTC_: <code>05:09:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-network/2037032781271273472#1:build-log.txt%3A572)
<details>
<summary>all...</summary>

* _2026-03-26 05:04:37 &#43;0000 UTC_: <code>05:09:45: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-network/2037032781271273472#1:build-log.txt%3A572)

</details>

<hr/>

**5x**: _2026-03-26 01:48:56 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-network/2036983573851934720#1:build-log.txt%3A538)
<details>
<summary>all...</summary>

* _2026-03-26 01:49:28 &#43;0000 UTC_: <code>01:59:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036983573042434048#1:build-log.txt%3A534)

* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:56:07: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-network/2036983576750198784#1:build-log.txt%3A549)

* _2026-03-26 01:48:56 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-network/2036983573851934720#1:build-log.txt%3A538)

* _2026-03-26 01:48:56 &#43;0000 UTC_: <code>01:54:56: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-network/2036983573382172672#1:build-log.txt%3A538)

* _2026-03-25 21:37:27 &#43;0000 UTC_: <code>21:42:31: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-network/2036909662514712576#1:build-log.txt%3A542)

</details>

<hr/>

**6x**: _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:51:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-network/2036967846151458816#1:build-log.txt%3A521)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:53:12: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2037088890333958144#1:build-log.txt%3A516)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2037073831012601856#1:build-log.txt%3A540)

* _2026-03-26 00:46:32 &#43;0000 UTC_: <code>00:54:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-network/2036967848978419712#1:build-log.txt%3A451)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:51:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-network/2036967846151458816#1:build-log.txt%3A521)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:53:15: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-network/2036967845614587904#1:build-log.txt%3A520)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:54:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036967845245489152#1:build-log.txt%3A517)

</details>

<hr/>

**4x**: _2026-03-25 22:28:15 &#43;0000 UTC_: <code>22:33:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-network/2036892367029538816#1:build-log.txt%3A529)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:27 &#43;0000 UTC_: <code>08:56:00: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-network/2037088890619170816#1:build-log.txt%3A516)

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:51:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-network/2037073834191884288#1:build-log.txt%3A524)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:51:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-network/2037073831499141120#1:build-log.txt%3A540)

* _2026-03-25 22:28:15 &#43;0000 UTC_: <code>22:33:40: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-network/2036892367029538816#1:build-log.txt%3A529)

</details>

<hr/>

**2x**: _2026-03-25 21:46:12 &#43;0000 UTC_: <code>21:51:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-network/2036892370359816192#1:build-log.txt%3A487)
<details>
<summary>all...</summary>

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-network/2037073831209734144#1:build-log.txt%3A540)

* _2026-03-25 21:46:12 &#43;0000 UTC_: <code>21:51:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-network/2036892370359816192#1:build-log.txt%3A487)

</details>

<hr/>

**2x**: _2026-03-25 19:52:50 &#43;0000 UTC_: <code>19:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-network/2036892366677217280#1:build-log.txt%3A529)
<details>
<summary>all...</summary>

* _2026-03-25 19:52:50 &#43;0000 UTC_: <code>19:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-network/2036892366677217280#1:build-log.txt%3A529)

* _2026-03-25 19:49:08 &#43;0000 UTC_: <code>19:53:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036892366492667904#1:build-log.txt%3A532)

</details>

<hr/>

**2x**: _2026-03-25 17:45:13 &#43;0000 UTC_: <code>17:48:29: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-network/2036861683703484416#1:build-log.txt%3A432)
<details>
<summary>all...</summary>

* _2026-03-25 18:30:51 &#43;0000 UTC_: <code>18:37:58: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-network/2036873262641713152#1:build-log.txt%3A469)

* _2026-03-25 17:45:13 &#43;0000 UTC_: <code>17:48:29: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-network/2036861683703484416#1:build-log.txt%3A432)

</details>

<hr/>

**1x**: _2026-03-25 17:29:54 &#43;0000 UTC_: <code>17:36:31: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036857570827702272#1:build-log.txt%3A458)
<details>
<summary>all...</summary>

* _2026-03-25 17:29:54 &#43;0000 UTC_: <code>17:36:31: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2036857570827702272#1:build-log.txt%3A458)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (3x / 10.34%) </summary>

<hr/>

**2x**: _2026-03-26 08:52:28 &#43;0000 UTC_: <code>08:57:24: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 73 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-network/2037090114290585600#1:build-log.txt%3A725)
<details>
<summary>all...</summary>

* _2026-03-26 08:52:43 &#43;0000 UTC_: <code>08:57:26: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-launcher/libvirt-hook-client/BUILD.bazel:10:10: GoCompilePkg cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client.a failed: Exec failed due to IOException: 351 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-network/2037090113296535552#1:build-log.txt%3A760)

* _2026-03-26 08:52:28 &#43;0000 UTC_: <code>08:57:24: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:16:10 Validating nogo output for //cmd/virt-controller:virt-controller failed: Exec failed due to IOException: 73 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-network/2037090114290585600#1:build-log.txt%3A725)

</details>

<hr/>

**1x**: _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:55:51: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/pierrec/lz4/v4/internal/xxh32/BUILD.bazel:3:11: Validating nogo output for //vendor/github.com/pierrec/lz4/v4/internal/xxh32:go_default_library failed: Exec failed due to IOException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-network/2037088893345468416#1:build-log.txt%3A467)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:55:51: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/pierrec/lz4/v4/internal/xxh32/BUILD.bazel:3:11: Validating nogo output for //vendor/github.com/pierrec/lz4/v4/internal/xxh32:go_default_library failed: Exec failed due to IOException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-network/2037088893345468416#1:build-log.txt%3A467)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (from secondary snippet) (1x / 3.45%) </summary>

<hr/>

**1x**: _2026-03-26 08:34:18 &#43;0000 UTC_: <code>08:57:20: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-k8s-1.34-sig-network/2037085594298880000#1:build-log.txt%3A4832)
<details>
<summary>all...</summary>

* _2026-03-26 08:34:18 &#43;0000 UTC_: <code>08:57:20: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-k8s-1.34-sig-network/2037085594298880000#1:build-log.txt%3A4832)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache download timeout (1x / 3.45%) </summary>

<hr/>

**1x**: _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:54:51: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportproxy/BUILD.bazel:27:10: GoLink cmd/virt-exportproxy/virt-exportproxy_/virt-exportproxy failed: Exec failed due to IOException: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/cas/1c6882fc7400bd46e51aa4130768869fb62c82d129bba7bf916ca68ab6a2ac1c&#39; timed out. Received 21508688 of 24245460 bytes.</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-network/2037088891227344896#1:build-log.txt%3A466)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>08:54:51: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-exportproxy/BUILD.bazel:27:10: GoLink cmd/virt-exportproxy/virt-exportproxy_/virt-exportproxy failed: Exec failed due to IOException: Download of &#39;/kubevirt/kubevirt,6b868b264dd2c4a4788a96bbe7462b63/cas/1c6882fc7400bd46e51aa4130768869fb62c82d129bba7bf916ca68ab6a2ac1c&#39; timed out. Received 21508688 of 24245460 bytes.</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-network/2037088891227344896#1:build-log.txt%3A466)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 3.45%) </summary>

<hr/>

**1x**: _2026-03-25 11:11:30 &#43;0000 UTC_: <code>11:21:44: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17069/pull-kubevirt-e2e-kind-sriov/2036723122534617088#1:build-log.txt%3A1464)
<details>
<summary>all...</summary>

* _2026-03-25 11:11:30 &#43;0000 UTC_: <code>11:21:44: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17069/pull-kubevirt-e2e-kind-sriov/2036723122534617088#1:build-log.txt%3A1464)

</details>

<hr/>
</details>

### sig-compute (82x / 58.99%)


#### external (65x / 79.27%)

<details>
<summary> bazel remote cache blob fetch failure (58x / 70.73%) </summary>

<hr/>

**12x**: _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:53:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037088891659358208#1:build-log.txt%3A518)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:34 &#43;0000 UTC_: <code>08:55:31: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-windows2016/2037088890090688512#1:build-log.txt%3A489)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:53:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037088891659358208#1:build-log.txt%3A518)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:52:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037088891520946176#1:build-log.txt%3A515)

* _2026-03-26 07:47:39 &#43;0000 UTC_: <code>07:53:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037073833353023488#1:build-log.txt%3A523)

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:35: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-windows2016/2037073830886772736#1:build-log.txt%3A549)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037073835894771712#1:build-log.txt%3A511)

* _2026-03-26 06:47:31 &#43;0000 UTC_: <code>06:51:22: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2037058711255519232#1:build-log.txt%3A424)

* _2026-03-25 22:46:26 &#43;0000 UTC_: <code>22:51:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036937636135833600#1:build-log.txt%3A503)

* _2026-03-25 22:30:47 &#43;0000 UTC_: <code>22:34:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036892368967307264#1:build-log.txt%3A512)

* _2026-03-25 22:29:24 &#43;0000 UTC_: <code>22:33:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036892368111669248#1:build-log.txt%3A461)

* _2026-03-25 21:43:49 &#43;0000 UTC_: <code>21:50:13: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036892369781002240#1:build-log.txt%3A459)

* _2026-03-25 20:39:55 &#43;0000 UTC_: <code>20:44:57: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036904750133284864#1:build-log.txt%3A445)

</details>

<hr/>

**5x**: _2026-03-26 05:04:29 &#43;0000 UTC_: <code>05:08:54: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037032781178998784#1:build-log.txt%3A529)
<details>
<summary>all...</summary>

* _2026-03-26 05:04:29 &#43;0000 UTC_: <code>05:08:54: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037032781178998784#1:build-log.txt%3A529)

* _2026-03-26 05:04:26 &#43;0000 UTC_: <code>05:10:23: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037032781506154496#1:build-log.txt%3A522)

* _2026-03-26 05:04:26 &#43;0000 UTC_: <code>05:10:15: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037032781090918400#1:build-log.txt%3A632)

* _2026-03-25 21:36:01 &#43;0000 UTC_: <code>21:40:21: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036909661638103040#1:build-log.txt%3A555)

* _2026-03-25 21:35:18 &#43;0000 UTC_: <code>21:41:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036909659150880768#1:build-log.txt%3A648)

</details>

<hr/>

**8x**: _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036983575089254400#1:build-log.txt%3A505)
<details>
<summary>all...</summary>

* _2026-03-26 05:04:29 &#43;0000 UTC_: <code>05:08:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037032781422268416#1:build-log.txt%3A541)

* _2026-03-26 01:49:19 &#43;0000 UTC_: <code>01:59:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036983575923920896#1:build-log.txt%3A480)

* _2026-03-26 01:49:19 &#43;0000 UTC_: <code>01:59:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: GoCompilePkg tests/virtctl/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-windows2016/2036983572635586560#1:build-log.txt%3A561)

* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:46: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036983574225227776#1:build-log.txt%3A534)

* _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:16: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11 Validating nogo output for //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036983575089254400#1:build-log.txt%3A505)

* _2026-03-26 01:48:59 &#43;0000 UTC_: <code>01:55:19: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036983573713522688#1:build-log.txt%3A535)

* _2026-03-26 01:48:58 &#43;0000 UTC_: <code>01:55:04: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036983578432114688#1:build-log.txt%3A527)

* _2026-03-26 01:48:57 &#43;0000 UTC_: <code>01:55:06: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/virtctl/BUILD.bazel:3:11: Running nogo on //tests/virtctl:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036983573587693568#1:build-log.txt%3A540)

</details>

<hr/>

**2x**: _2026-03-26 07:47:48 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037073832539328512#1:build-log.txt%3A502)
<details>
<summary>all...</summary>

* _2026-03-26 07:47:48 &#43;0000 UTC_: <code>07:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037073832539328512#1:build-log.txt%3A502)

* _2026-03-26 07:47:47 &#43;0000 UTC_: <code>07:53:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037073831079710720#1:build-log.txt%3A595)

</details>

<hr/>

**18x**: _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036967847292309504#1:build-log.txt%3A508)
<details>
<summary>all...</summary>

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.33-sig-operator/2037088891055378432#1:build-log.txt%3A482)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>08:55:34: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2037088892464664576#1:build-log.txt%3A517)

* _2026-03-26 07:47:40 &#43;0000 UTC_: <code>07:51:51: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-operator/2037073831419449344#1:build-log.txt%3A529)

* _2026-03-26 07:47:39 &#43;0000 UTC_: <code>07:52:11: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037073831914377216#1:build-log.txt%3A525)

* _2026-03-26 07:47:38 &#43;0000 UTC_: <code>07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037073836737826816#1:build-log.txt%3A529)

* _2026-03-26 07:47:37 &#43;0000 UTC_: <code>07:53:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-k8s-1.33-sig-compute/2037073831356534784#1:build-log.txt%3A460)

* _2026-03-26 02:47:25 &#43;0000 UTC_: <code>02:52:05: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036998292239290368#1:build-log.txt%3A481)

* _2026-03-26 00:46:37 &#43;0000 UTC_: <code>00:55:44: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036967851490807808#1:build-log.txt%3A510)

* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute-serial/2036967848152141824#1:build-log.txt%3A478)

* _2026-03-26 00:46:28 &#43;0000 UTC_: <code>00:52:01: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036967847292309504#1:build-log.txt%3A508)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:52:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036967850643558400#1:build-log.txt%3A505)

* _2026-03-26 00:46:27 &#43;0000 UTC_: <code>00:51:29: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036967846004658176#1:build-log.txt%3A506)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:55:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036967845870440448#1:build-log.txt%3A504)

* _2026-03-26 00:46:26 &#43;0000 UTC_: <code>00:52:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: Running nogo on //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036967845358735360#1:build-log.txt%3A588)

* _2026-03-26 00:46:25 &#43;0000 UTC_: <code>00:52:25: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036967846478614528#1:build-log.txt%3A510)

* _2026-03-26 00:46:23 &#43;0000 UTC_: <code>00:54:50: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-windows2016/2036967844863807488#1:build-log.txt%3A520)

* _2026-03-25 22:27:03 &#43;0000 UTC_: <code>22:32:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036892366903709696#1:build-log.txt%3A503)

* _2026-03-25 21:47:42 &#43;0000 UTC_: <code>21:52:56: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036892374516371456#1:build-log.txt%3A439)

</details>

<hr/>

**5x**: _2026-03-25 22:08:13 &#43;0000 UTC_: <code>22:12:17: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036857991180849152#1:build-log.txt%3A346)
<details>
<summary>all...</summary>

* _2026-03-25 22:08:13 &#43;0000 UTC_: <code>22:12:17: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036857991180849152#1:build-log.txt%3A346)

* _2026-03-25 21:08:14 &#43;0000 UTC_: <code>21:12:54: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-compute/2036862229206274048#1:build-log.txt%3A433)

* _2026-03-25 19:08:48 &#43;0000 UTC_: <code>19:16:10: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036882243925839872#1:build-log.txt%3A394)

* _2026-03-25 18:30:52 &#43;0000 UTC_: <code>18:36:56: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.33-sig-operator/2036873262981451776#1:build-log.txt%3A500)

* _2026-03-25 18:25:43 &#43;0000 UTC_: <code>18:30:50: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:113:10: OCI Image //containerimages:alpine-ext-kernel-boot-demo-container failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-windows2016/2036871960771694592#1:build-log.txt%3A468)

</details>

<hr/>

**4x**: _2026-03-25 21:08:47 &#43;0000 UTC_: <code>21:13:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036912517216735232#1:build-log.txt%3A557)
<details>
<summary>all...</summary>

* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036983579300335616#1:build-log.txt%3A541)

* _2026-03-26 01:49:01 &#43;0000 UTC_: <code>01:57:14: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036983573143097344#1:build-log.txt%3A615)

* _2026-03-25 21:43:46 &#43;0000 UTC_: <code>21:50:52: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-operator/2036909664993546240#1:build-log.txt%3A565)

* _2026-03-25 21:08:47 &#43;0000 UTC_: <code>21:13:18: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-k8s-1.35-sig-compute/2036912517216735232#1:build-log.txt%3A557)

</details>

<hr/>

**3x**: _2026-03-25 19:56:46 &#43;0000 UTC_: <code>20:08:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036892366803046400#1:build-log.txt%3A506)
<details>
<summary>all...</summary>

* _2026-03-25 19:56:46 &#43;0000 UTC_: <code>20:08:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.33-sig-compute/2036892366803046400#1:build-log.txt%3A506)

* _2026-03-25 19:52:47 &#43;0000 UTC_: <code>19:56:55: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2036892366597525504#1:build-log.txt%3A586)

* _2026-03-25 19:46:36 &#43;0000 UTC_: <code>19:51:36: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16889/pull-kubevirt-e2e-k8s-1.34-windows2016/2036892366194872320#1:build-log.txt%3A512)

</details>

<hr/>

**1x**: _2026-03-25 17:38:42 &#43;0000 UTC_: <code>17:42:45: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036860131525791744#1:build-log.txt%3A423)
<details>
<summary>all...</summary>

* _2026-03-25 17:38:42 &#43;0000 UTC_: <code>17:42:45: ERROR: /root/go/src/kubevirt.io/kubevirt/containerimages/BUILD.bazel:93:10: OCI Image //containerimages:fedora-with-test-tooling failed: Failed to fetch blobs because they do not exist remotely.: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17138/pull-kubevirt-e2e-k8s-1.34-sig-operator/2036860131525791744#1:build-log.txt%3A423)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache connection timeout (3x / 3.66%) </summary>

<hr/>

**3x**: _2026-03-26 08:52:28 &#43;0000 UTC_: <code>15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037090115150417920#1:build-log.txt%3A1004)
<details>
<summary>all...</summary>

* _2026-03-26 08:52:28 &#43;0000 UTC_: <code>15:28:16: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.210.196:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-operator/2037090115150417920#1:build-log.txt%3A1004)

* _2026-03-26 08:47:26 &#43;0000 UTC_: <code>09:02:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037088890476564480#1:build-log.txt%3A2362)

* _2026-03-26 08:47:25 &#43;0000 UTC_: <code>09:00:52: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.173.157:8080</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-compute/2037088895111270400#1:build-log.txt%3A1119)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (2x / 2.44%) </summary>

<hr/>

**2x**: _2026-03-26 08:52:30 &#43;0000 UTC_: <code>08:58:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tools/csv-generator/BUILD.bazel:16:10 Validating nogo output for //tools/csv-generator:csv-generator failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037090114504495104#1:build-log.txt%3A490)
<details>
<summary>all...</summary>

* _2026-03-26 08:52:30 &#43;0000 UTC_: <code>08:58:24: ERROR: /root/go/src/kubevirt.io/kubevirt/tools/csv-generator/BUILD.bazel:16:10 Validating nogo output for //tools/csv-generator:csv-generator failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-k8s-1.34-sig-compute/2037090114504495104#1:build-log.txt%3A490)

* _2026-03-26 08:47:28 &#43;0000 UTC_: <code>08:58:56: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/virt-handler/vsock/system/BUILD.bazel:3:11 Validating nogo output for //pkg/virt-handler/vsock/system:go_default_library failed: Exec failed due to IOException: 3 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-k8s-1.35-sig-operator/2037088895878828032#1:build-log.txt%3A883)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 1.22%) </summary>

<hr/>

**1x**: _2026-03-26 06:47:21 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2037058711163244544#1:build-log.txt%3A5065)
<details>
<summary>all...</summary>

* _2026-03-26 06:47:21 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2037058711163244544#1:build-log.txt%3A5065)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 1.22%) </summary>

<hr/>

**1x**: _2026-03-26 03:48:42 &#43;0000 UTC_: <code>04:06:36: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037013641730985984#1:build-log.txt%3A4356)
<details>
<summary>all...</summary>

* _2026-03-26 03:48:42 &#43;0000 UTC_: <code>04:06:36: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2037013641730985984#1:build-log.txt%3A4356)

</details>

<hr/>
</details>

#### internal (7x / 8.54%)

<details>
<summary> kind cluster creation failure (6x / 7.32%) </summary>

<hr/>

**6x**: _2026-03-25 17:05:27 &#43;0000 UTC_: <code>17:12:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17291/pull-kubevirt-e2e-kind-1.34-sev-1.7/2036851449563975680#1:build-log.txt%3A609)
<details>
<summary>all...</summary>

* _2026-03-26 09:22:19 &#43;0000 UTC_: <code>09:29:01: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037097667556806656#1:build-log.txt%3A816)

* _2026-03-26 08:47:18 &#43;0000 UTC_: <code>08:54:22: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17042/pull-kubevirt-e2e-kind-1.34-sev/2037088890237489152#1:build-log.txt%3A709)

* _2026-03-26 01:49:22 &#43;0000 UTC_: <code>01:55:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2036983572786581504#1:build-log.txt%3A789)

* _2026-03-25 22:46:16 &#43;0000 UTC_: <code>22:52:30: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036937635984838656#1:build-log.txt%3A697)

* _2026-03-25 17:07:07 &#43;0000 UTC_: <code>17:15:35: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17290/pull-kubevirt-e2e-kind-1.34-sev-1.8/2036852207432765440#1:build-log.txt%3A834)

* _2026-03-25 17:05:27 &#43;0000 UTC_: <code>17:12:41: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17291/pull-kubevirt-e2e-kind-1.34-sev-1.7/2036851449563975680#1:build-log.txt%3A609)

</details>

<hr/>
</details>
<details>
<summary> control plane startup failure (1x / 1.22%) </summary>

<hr/>

**1x**: _2026-03-26 03:48:17 &#43;0000 UTC_: <code>04:03:20: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037013641584185344#1:build-log.txt%3A2850)
<details>
<summary>all...</summary>

* _2026-03-26 03:48:17 &#43;0000 UTC_: <code>04:03:20: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037013641584185344#1:build-log.txt%3A2850)

</details>

<hr/>
</details>

#### needs-investigation (8x / 9.76%)

<details>
<summary> no error snippets (8x / 9.76%) </summary>

<hr/>

**1x**: _2026-03-26 08:52:44 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2037090112856133632)
<details>
<summary>all...</summary>

* _2026-03-26 08:52:44 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16514/pull-kubevirt-e2e-kind-1.34-sev/2037090112856133632)

</details>

<hr/>

**1x**: _2026-03-26 08:50:21 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037089601964740608)
<details>
<summary>all...</summary>

* _2026-03-26 08:50:21 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037089601964740608)

</details>

<hr/>

**1x**: _2026-03-26 07:47:29 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-kind-1.34-sev/2037073830949687296)
<details>
<summary>all...</summary>

* _2026-03-26 07:47:29 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16842/pull-kubevirt-e2e-kind-1.34-sev/2037073830949687296)

</details>

<hr/>

**1x**: _2026-03-26 05:04:21 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037032780927340544)
<details>
<summary>all...</summary>

* _2026-03-26 05:04:21 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2037032780927340544)

</details>

<hr/>

**1x**: _2026-03-26 02:47:17 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036998292079906816)
<details>
<summary>all...</summary>

* _2026-03-26 02:47:17 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17273/pull-kubevirt-e2e-kind-1.34-sev/2036998292079906816)

</details>

<hr/>

**1x**: _2026-03-26 00:46:17 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-kind-1.34-sev/2036967845056745472)
<details>
<summary>all...</summary>

* _2026-03-26 00:46:17 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-kind-1.34-sev/2036967845056745472)

</details>

<hr/>

**1x**: _2026-03-25 20:55:09 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2036909651676631040)
<details>
<summary>all...</summary>

* _2026-03-25 20:55:09 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17296/pull-kubevirt-e2e-kind-1.34-sev/2036909651676631040)

</details>

<hr/>

**1x**: _2026-03-25 18:46:20 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17207/pull-kubevirt-e2e-kind-1.34-sev/2036877259465297920)
<details>
<summary>all...</summary>

* _2026-03-25 18:46:20 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17207/pull-kubevirt-e2e-kind-1.34-sev/2036877259465297920)

</details>

<hr/>
</details>

#### pr-build (2x / 2.44%)

<details>
<summary> panic detected in test output (2x / 2.44%) </summary>

<hr/>

**2x**: _2026-03-26 07:35:38 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037070842281594880#1:build-log.txt%3A5081)
<details>
<summary>all...</summary>

* _2026-03-26 08:34:11 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17240/pull-kubevirt-e2e-kind-1.34-sev/2037085594152079360#1:build-log.txt%3A5063)

* _2026-03-26 07:35:38 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17251/pull-kubevirt-e2e-kind-1.34-sev/2037070842281594880#1:build-log.txt%3A5081)

</details>

<hr/>
</details>

### sig-monitoring (1x / 0.72%)


#### external (1x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-03-26 00:46:25 &#43;0000 UTC_: <code>00:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-monitoring/2036967845480370176#1:build-log.txt%3A508)
<details>
<summary>all...</summary>

* _2026-03-26 00:46:25 &#43;0000 UTC_: <code>00:54:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8: GoCompilePkg tests/go_default_test.internal.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: a72f131812f04bc6dd583d392bb3fb253b5db26a4c0694fda7cb33bc77ac4e83/139666 for bazel-out/k8-fastbuild/bin/tests/virtctl/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io//view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17253/pull-kubevirt-e2e-k8s-1.34-sig-monitoring/2036967845480370176#1:build-log.txt%3A508)

</details>

<hr/>
</details>

Last updated: 2026-04-01 09:35:17
