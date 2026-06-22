



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-06-21 (1x / 1.45%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)
<details>
<summary>all...</summary>

* _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)

</details>

<hr/>
</details>

### 2026-06-18 (65x / 94.20%)


#### external (65x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (64x / 98.46%) </summary>

<hr/>

**26x**: _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:31:34: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067524113043296256#1:build-log.txt%3A727)
<details>
<summary>all...</summary>

* _2026-06-18 09:49:09 &#43;0000 UTC_: <code>09:55:37: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067544976245395456#1:build-log.txt%3A729)

* _2026-06-18 09:49:06 &#43;0000 UTC_: <code>09:52:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067544975767244800#1:build-log.txt%3A744)

* _2026-06-18 09:49:04 &#43;0000 UTC_: <code>09:52:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067544975628832768#1:build-log.txt%3A741)

* _2026-06-18 09:49:01 &#43;0000 UTC_: <code>09:56:38: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067544975280705536#1:build-log.txt%3A828)

* _2026-06-18 09:48:57 &#43;0000 UTC_: <code>09:55:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067544975377174528#1:build-log.txt%3A755)

* _2026-06-18 09:48:57 &#43;0000 UTC_: <code>09:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067544975154876416#1:build-log.txt%3A755)

* _2026-06-18 09:34:30 &#43;0000 UTC_: <code>09:40:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067536885676249088#1:build-log.txt%3A815)

* _2026-06-18 09:31:41 &#43;0000 UTC_: <code>09:41:02: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067536883935612928#1:build-log.txt%3A779)

* _2026-06-18 09:28:37 &#43;0000 UTC_: <code>09:32:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067536883151278080#1:build-log.txt%3A772)

* _2026-06-18 09:28:12 &#43;0000 UTC_: <code>09:32:12: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067536882257891328#1:build-log.txt%3A806)

* _2026-06-18 09:17:38 &#43;0000 UTC_: <code>09:21:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067536881737797632#1:build-log.txt%3A731)

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067536881641328640#1:build-log.txt%3A769)

* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:22:17: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067536880865382400#1:build-log.txt%3A741)

* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:20:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067536881456779264#1:build-log.txt%3A742)

* _2026-06-18 09:16:49 &#43;0000 UTC_: <code>09:21:44: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067536880655667200#1:build-log.txt%3A740)

* _2026-06-18 09:16:47 &#43;0000 UTC_: <code>09:21:22: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067536880974434304#1:build-log.txt%3A827)

* _2026-06-18 08:27:17 &#43;0000 UTC_: <code>08:32:22: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067524389779279872#1:build-log.txt%3A805)

* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:34:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067524390601363456#1:build-log.txt%3A746)

* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:33:23: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067524389590536192#1:build-log.txt%3A807)

* _2026-06-18 08:27:13 &#43;0000 UTC_: <code>08:33:38: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067524389867360256#1:build-log.txt%3A813)

* _2026-06-18 08:27:12 &#43;0000 UTC_: <code>08:33:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067524389456318464#1:build-log.txt%3A813)

* _2026-06-18 08:26:13 &#43;0000 UTC_: <code>08:31:57: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067524111256522752#1:build-log.txt%3A722)

* _2026-06-18 08:26:11 &#43;0000 UTC_: <code>08:30:58: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067524112086994944#1:build-log.txt%3A729)

* _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:32:01: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067524113760522240#1:build-log.txt%3A756)

* _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:31:34: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067524113043296256#1:build-log.txt%3A727)

* _2026-06-18 08:26:05 &#43;0000 UTC_: <code>08:31:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067524107909468160#1:build-log.txt%3A811)

</details>

<hr/>

**21x**: _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:34:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067555554041008128#1:build-log.txt%3A767)
<details>
<summary>all...</summary>

* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:35:36: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067555549490188288#1:build-log.txt%3A822)

* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:35:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067555550215802880#1:build-log.txt%3A813)

* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:34:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067555554041008128#1:build-log.txt%3A767)

* _2026-06-18 10:31:06 &#43;0000 UTC_: <code>10:35:00: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067555550236774400#1:build-log.txt%3A742)

* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067555550224191488#1:build-log.txt%3A795)

* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:35:59: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067555549305638912#1:build-log.txt%3A756)

* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067555553206341632#1:build-log.txt%3A757)

* _2026-06-18 10:31:04 &#43;0000 UTC_: <code>10:36:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067555550635233280#1:build-log.txt%3A799)

* _2026-06-18 10:31:03 &#43;0000 UTC_: <code>10:36:48: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067555550232580096#1:build-log.txt%3A735)

* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067555554854703104#1:build-log.txt%3A759)

* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:31: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067555552426201088#1:build-log.txt%3A747)

* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067555551469899776#1:build-log.txt%3A824)

* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:35:39: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067555550228385792#1:build-log.txt%3A728)

* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:36:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067555549397913600#1:build-log.txt%3A891)

* _2026-06-18 09:58:25 &#43;0000 UTC_: <code>10:03:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067544978329964544#1:build-log.txt%3A731)

* _2026-06-18 09:58:25 &#43;0000 UTC_: <code>10:03:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067544979118493696#1:build-log.txt%3A744)

* _2026-06-18 09:58:23 &#43;0000 UTC_: <code>10:02:10: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067544980003491840#1:build-log.txt%3A723)

* _2026-06-18 09:57:15 &#43;0000 UTC_: <code>10:04:54: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067544977449160704#1:build-log.txt%3A750)

* _2026-06-18 09:53:46 &#43;0000 UTC_: <code>10:00:53: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067544976908095488#1:build-log.txt%3A785)

* _2026-06-18 08:27:25 &#43;0000 UTC_: <code>08:32:47: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067524389980606464#1:build-log.txt%3A756)

* _2026-06-18 08:27:18 &#43;0000 UTC_: <code>08:33:30: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067524389368238080#1:build-log.txt%3A799)

</details>

<hr/>

**11x**: _2026-06-18 09:49:11 &#43;0000 UTC_: <code>09:56:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067544975490420736#1:build-log.txt%3A748)
<details>
<summary>all...</summary>

* _2026-06-18 10:31:02 &#43;0000 UTC_: <code>10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067555549053980672#1:build-log.txt%3A836)

* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:43: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067555550219997184#1:build-log.txt%3A768)

* _2026-06-18 09:49:12 &#43;0000 UTC_: <code>09:57:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067544975045824512#1:build-log.txt%3A832)

* _2026-06-18 09:49:11 &#43;0000 UTC_: <code>09:56:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067544975490420736#1:build-log.txt%3A748)

* _2026-06-18 09:49:06 &#43;0000 UTC_: <code>09:54:00: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067544975918239744#1:build-log.txt%3A811)

* _2026-06-18 09:49:03 &#43;0000 UTC_: <code>09:59:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067544976639660032#1:build-log.txt%3A813)

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:20: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067536881171566592#1:build-log.txt%3A734)

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:21:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067536881553248256#1:build-log.txt%3A748)

* _2026-06-18 09:16:47 &#43;0000 UTC_: <code>09:21:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067536881284812800#1:build-log.txt%3A781)

* _2026-06-18 08:34:54 &#43;0000 UTC_: <code>08:38:26: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067526318710329344#1:build-log.txt%3A759)

* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:33:42: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067524389292740608#1:build-log.txt%3A742)

</details>

<hr/>

**3x**: _2026-06-18 09:31:48 &#43;0000 UTC_: <code>09:36:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067536884770279424#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-18 09:49:03 &#43;0000 UTC_: <code>09:53:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067544976035680256#1:build-log.txt%3A697)

* _2026-06-18 09:31:48 &#43;0000 UTC_: <code>09:36:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067536884770279424#1:build-log.txt%3A729)

* _2026-06-18 08:27:17 &#43;0000 UTC_: <code>08:33:02: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067524389678616576#1:build-log.txt%3A750)

</details>

<hr/>

**2x**: _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067536881062514688#1:build-log.txt%3A728)
<details>
<summary>all...</summary>

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067536881062514688#1:build-log.txt%3A728)

* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:22:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067536881393864704#1:build-log.txt%3A747)

</details>

<hr/>

**1x**: _2026-06-18 08:26:08 &#43;0000 UTC_: <code>08:32:27: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: Running nogo on //vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067524110417661952#1:build-log.txt%3A744)
<details>
<summary>all...</summary>

* _2026-06-18 08:26:08 &#43;0000 UTC_: <code>08:32:27: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: Running nogo on //vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067524110417661952#1:build-log.txt%3A744)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 1.54%) </summary>

<hr/>

**1x**: _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)
<details>
<summary>all...</summary>

* _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)

</details>

<hr/>
</details>

### 2026-06-17 (1x / 1.45%)


#### external (1x / 100.00%)

<details>
<summary> podman container removal timeout (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)
<details>
<summary>all...</summary>

* _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)

</details>

<hr/>
</details>

### 2026-06-16 (2x / 2.90%)


#### internal (1x / 50.00%)

<details>
<summary> kind cluster creation failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)
<details>
<summary>all...</summary>

* _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (68x / 98.55%)

<details>
<summary> bazel remote cache blob fetch failure (64x / 92.75%) </summary>

<hr/>

**26x**: _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:31:34: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067524113043296256#1:build-log.txt%3A727)
<details>
<summary>all...</summary>

* _2026-06-18 09:49:09 &#43;0000 UTC_: <code>09:55:37: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067544976245395456#1:build-log.txt%3A729)
<details><summary>context</summary>
<pre>09:55:07:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
09:55:07:       |   ^~~~~~~~
09:55:07:       |   vsprintf
09:55:37: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:55:37: INFO: Elapsed time: 40.446s, Critical Path: 21.99s
09:55:37: INFO: 6155 processes: 6126 remote cache hit, 29 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:49:06 &#43;0000 UTC_: <code>09:52:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067544975767244800#1:build-log.txt%3A744)
<details><summary>context</summary>
<pre>09:52:27: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:52:27: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:52:27: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:52:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:52:32: INFO: Elapsed time: 13.492s, Critical Path: 6.07s
09:52:32: INFO: 3512 processes: 3482 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:49:04 &#43;0000 UTC_: <code>09:52:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067544975628832768#1:build-log.txt%3A741)
<details><summary>context</summary>
<pre>09:52:18: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:52:18: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:52:18: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:52:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:52:24: INFO: Elapsed time: 12.891s, Critical Path: 5.45s
09:52:24: INFO: 3473 processes: 3442 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:49:01 &#43;0000 UTC_: <code>09:56:38: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067544975280705536#1:build-log.txt%3A828)
<details><summary>context</summary>
<pre>09:56:04: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:56:04: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:56:04: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:56:38: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:56:38: INFO: Elapsed time: 47.767s, Critical Path: 31.23s
09:56:38: INFO: 4099 processes: 4068 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:48:57 &#43;0000 UTC_: <code>09:55:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067544975377174528#1:build-log.txt%3A755)
<details><summary>context</summary>
<pre>09:55:40: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:55:40: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:55:40: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:55:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:55:42: INFO: Elapsed time: 48.059s, Critical Path: 32.22s
09:55:42: INFO: 7017 processes: 6982 remote cache hit, 31 internal, 4 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:48:57 &#43;0000 UTC_: <code>09:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067544975154876416#1:build-log.txt%3A755)
<details><summary>context</summary>
<pre>09:53:46: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:53:46: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:53:46: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:53:48: INFO: Elapsed time: 24.404s, Critical Path: 6.11s
09:53:48: INFO: 4779 processes: 4748 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:34:30 &#43;0000 UTC_: <code>09:40:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067536885676249088#1:build-log.txt%3A815)
<details><summary>context</summary>
<pre>09:40:14: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:40:14: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:40:14: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:40:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:40:17: INFO: Elapsed time: 46.113s, Critical Path: 26.49s
09:40:17: INFO: 7038 processes: 7001 remote cache hit, 31 internal, 6 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:31:41 &#43;0000 UTC_: <code>09:41:02: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067536883935612928#1:build-log.txt%3A779)
<details><summary>context</summary>
<pre>09:40:59: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:40:59: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:40:59: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:41:02: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:41:02: INFO: Elapsed time: 15.708s, Critical Path: 5.20s
09:41:02: INFO: 4573 processes: 4543 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:28:37 &#43;0000 UTC_: <code>09:32:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067536883151278080#1:build-log.txt%3A772)
<details><summary>context</summary>
<pre>09:32:10:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
09:32:10:       |   ^~~~~~~~
09:32:10:       |   vsprintf
09:32:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:32:25: INFO: Elapsed time: 19.395s, Critical Path: 4.65s
09:32:25: INFO: 3437 processes: 3408 remote cache hit, 29 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:28:12 &#43;0000 UTC_: <code>09:32:12: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067536882257891328#1:build-log.txt%3A806)
<details><summary>context</summary>
<pre>09:32:04: INFO: From GoLink cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client_/libvirt-hook-client:
09:32:04: /usr/bin/ld: /tmp/go-link-3209859693/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:32:04: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:32:12: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:32:12: INFO: Elapsed time: 14.108s, Critical Path: 6.35s
09:32:12: INFO: 1745 processes: 1714 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:17:38 &#43;0000 UTC_: <code>09:21:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067536881737797632#1:build-log.txt%3A731)
<details><summary>context</summary>
<pre>09:21:01:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
09:21:01:       |   ^~~~~~~~
09:21:01:       |   vsprintf
09:21:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:21:13: INFO: Elapsed time: 12.609s, Critical Path: 6.27s
09:21:13: INFO: 3632 processes: 3601 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067536881641328640#1:build-log.txt%3A769)
<details><summary>context</summary>
<pre>09:20:13: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:20:13: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:20:13: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:20:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:20:32: INFO: Elapsed time: 22.667s, Critical Path: 13.76s
09:20:32: INFO: 3483 processes: 3449 remote cache hit, 31 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:22:17: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067536880865382400#1:build-log.txt%3A741)
<details><summary>context</summary>
<pre>09:22:11: INFO: From GoLink cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client_/libvirt-hook-client:
09:22:11: /usr/bin/ld: /tmp/go-link-3209859693/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:22:11: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:22:17: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:22:17: INFO: Elapsed time: 18.588s, Critical Path: 10.87s
09:22:17: INFO: 3032 processes: 3004 remote cache hit, 28 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:20:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067536881456779264#1:build-log.txt%3A742)
<details><summary>context</summary>
<pre>09:20:34:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
09:20:34:       |   ^~~~~~~~
09:20:34:       |   vsprintf
09:20:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:20:43: INFO: Elapsed time: 9.551s, Critical Path: 6.68s
09:20:43: INFO: 2804 processes: 2773 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:16:49 &#43;0000 UTC_: <code>09:21:44: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067536880655667200#1:build-log.txt%3A740)
<details><summary>context</summary>
<pre>09:21:27:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
09:21:27:       |   ^~~~~~~~
09:21:27:       |   vsprintf
09:21:44: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:21:44: INFO: Elapsed time: 18.598s, Critical Path: 11.33s
09:21:44: INFO: 4376 processes: 4346 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:16:47 &#43;0000 UTC_: <code>09:21:22: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067536880974434304#1:build-log.txt%3A827)
<details><summary>context</summary>
<pre>09:21:13: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:21:13: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:21:13: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:21:22: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:21:22: INFO: Elapsed time: 15.737s, Critical Path: 8.60s
09:21:22: INFO: 4417 processes: 4386 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:27:17 &#43;0000 UTC_: <code>08:32:22: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067524389779279872#1:build-log.txt%3A805)
<details><summary>context</summary>
<pre>08:32:07:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
08:32:07:       |   ^~~~~~~~
08:32:07:       |   vsprintf
08:32:22: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:32:22: INFO: Elapsed time: 43.695s, Critical Path: 5.97s
08:32:22: INFO: 3487 processes: 3457 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:34:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067524390601363456#1:build-log.txt%3A746)
<details><summary>context</summary>
<pre>08:34:25: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:34:25: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:34:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:34:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:34:32: INFO: Elapsed time: 16.728s, Critical Path: 4.77s
08:34:32: INFO: 4911 processes: 4881 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:33:23: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067524389590536192#1:build-log.txt%3A807)
<details><summary>context</summary>
<pre>08:33:08: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:33:08: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:33:08: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:33:23: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:33:23: INFO: Elapsed time: 22.042s, Critical Path: 13.50s
08:33:23: INFO: 7143 processes: 7114 remote cache hit, 29 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:27:13 &#43;0000 UTC_: <code>08:33:38: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067524389867360256#1:build-log.txt%3A813)
<details><summary>context</summary>
<pre>08:33:35: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:33:35: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:33:35: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:33:38: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:33:38: INFO: Elapsed time: 36.191s, Critical Path: 20.89s
08:33:38: INFO: 7064 processes: 7025 remote cache hit, 30 internal, 9 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:27:12 &#43;0000 UTC_: <code>08:33:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067524389456318464#1:build-log.txt%3A813)
<details><summary>context</summary>
<pre>08:32:58: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:32:58: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:32:58: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:33:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:33:08: INFO: Elapsed time: 24.130s, Critical Path: 4.11s
08:33:08: INFO: 6034 processes: 6004 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:26:13 &#43;0000 UTC_: <code>08:31:57: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067524111256522752#1:build-log.txt%3A722)
<details><summary>context</summary>
<pre>08:31:15:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
08:31:15:       |   ^~~~~~~~
08:31:15:       |   vsprintf
08:31:57: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:31:57: INFO: Elapsed time: 59.791s, Critical Path: 7.08s
08:31:57: INFO: 4726 processes: 4696 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:26:11 &#43;0000 UTC_: <code>08:30:58: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067524112086994944#1:build-log.txt%3A729)
<details><summary>context</summary>
<pre>08:30:30:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
08:30:30:       |   ^~~~~~~~
08:30:30:       |   vsprintf
08:30:58: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:30:59: INFO: Elapsed time: 49.594s, Critical Path: 5.13s
08:30:59: INFO: 4062 processes: 4031 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:32:01: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067524113760522240#1:build-log.txt%3A756)
<details><summary>context</summary>
<pre>08:32:01: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:32:01: /usr/bin/ld: /tmp/go-link-1944123972/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:32:01: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:32:01: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:32:01: INFO: Elapsed time: 50.748s, Critical Path: 22.67s
08:32:01: INFO: 3432 processes: 3402 remote cache hit, 29 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:31:34: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067524113043296256#1:build-log.txt%3A727)
<details><summary>context</summary>
<pre>08:31:32: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:31:32: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:31:32: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:31:34: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:31:34: INFO: Elapsed time: 56.764s, Critical Path: 8.90s
08:31:34: INFO: 4090 processes: 4060 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:26:05 &#43;0000 UTC_: <code>08:31:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067524107909468160#1:build-log.txt%3A811)
<details><summary>context</summary>
<pre>08:30:35: INFO: From GoLink cmd/virt-launcher/libvirt-hook-client/libvirt-hook-client_/libvirt-hook-client:
08:30:35: /usr/bin/ld: /tmp/go-link-3209859693/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:30:35: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:31:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:31:26: INFO: Elapsed time: 77.244s, Critical Path: 34.09s
08:31:26: INFO: 3917 processes: 3885 remote cache hit, 30 internal, 2 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**21x**: _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:34:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067555554041008128#1:build-log.txt%3A767)
<details>
<summary>all...</summary>

* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:35:36: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067555549490188288#1:build-log.txt%3A822)
<details><summary>context</summary>
<pre>10:35:25: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
10:35:25: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:35:36: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:35:36: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:35:36: INFO: Elapsed time: 25.783s, Critical Path: 4.51s
10:35:36: INFO: 5270 processes: 5239 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:35:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067555550215802880#1:build-log.txt%3A813)
<details><summary>context</summary>
<pre>10:35:01: /usr/bin/ld: /tmp/go-link-475742333/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
10:35:01: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:35:07: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:35:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:35:08: INFO: Elapsed time: 37.365s, Critical Path: 16.13s
10:35:08: INFO: 4862 processes: 4828 remote cache hit, 31 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:34:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067555554041008128#1:build-log.txt%3A767)
<details><summary>context</summary>
<pre>10:34:11:       |   ^~~~~~~~
10:34:11:       |   vsprintf
10:34:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:34:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:34:21: INFO: Elapsed time: 13.570s, Critical Path: 5.50s
10:34:21: INFO: 2684 processes: 2653 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:31:06 &#43;0000 UTC_: <code>10:35:00: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067555550236774400#1:build-log.txt%3A742)
<details><summary>context</summary>
<pre>10:34:46:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
10:34:46:       |   ^~~~~~~~
10:34:46:       |   vsprintf
10:35:00: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:35:01: INFO: Elapsed time: 15.882s, Critical Path: 5.39s
10:35:01: INFO: 3454 processes: 3425 remote cache hit, 29 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067555550224191488#1:build-log.txt%3A795)
<details><summary>context</summary>
<pre>10:36:48: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
10:36:48: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:37:07: INFO: Elapsed time: 62.231s, Critical Path: 16.40s
10:37:07: INFO: 6821 processes: 6784 remote cache hit, 30 internal, 7 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:35:59: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067555549305638912#1:build-log.txt%3A756)
<details><summary>context</summary>
<pre>10:35:41: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
10:35:41: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:35:59: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:35:59: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:35:59: INFO: Elapsed time: 34.714s, Critical Path: 21.01s
10:35:59: INFO: 4615 processes: 4584 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067555553206341632#1:build-log.txt%3A757)
<details><summary>context</summary>
<pre>10:36:02:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
10:36:02:       |   ^~~~~~~~
10:36:02:       |   vsprintf
10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:15: INFO: Elapsed time: 16.543s, Critical Path: 5.61s
10:36:15: INFO: 4735 processes: 4707 remote cache hit, 28 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:31:04 &#43;0000 UTC_: <code>10:36:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067555550635233280#1:build-log.txt%3A799)
<details><summary>context</summary>
<pre>10:36:17: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:36:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x
10:36:25: INFO: Elapsed time: 61.157s, Critical Path: 27.12s
10:36:25: INFO: 6659 processes: 6620 remote cache hit, 31 internal, 8 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:31:03 &#43;0000 UTC_: <code>10:36:48: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067555550232580096#1:build-log.txt%3A735)
<details><summary>context</summary>
<pre>10:36:37:       |   ^~~~~~~~
10:36:37:       |   vsprintf
10:36:48: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:48: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:48: INFO: Elapsed time: 37.504s, Critical Path: 28.43s
10:36:48: INFO: 3672 processes: 3641 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067555554854703104#1:build-log.txt%3A759)
<details><summary>context</summary>
<pre>10:36:01:       |   ^~~~~~~~
10:36:01:       |   vsprintf
10:36:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:20: INFO: Elapsed time: 23.107s, Critical Path: 5.85s
10:36:20: INFO: 4180 processes: 4149 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:31: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067555552426201088#1:build-log.txt%3A747)
<details><summary>context</summary>
<pre>10:36:28: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:36:28: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
10:36:28: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:36:31: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:32: INFO: Elapsed time: 14.545s, Critical Path: 5.88s
10:36:32: INFO: 4082 processes: 4053 remote cache hit, 29 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067555551469899776#1:build-log.txt%3A824)
<details><summary>context</summary>
<pre>10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: Running nogo on //vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x
10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:15: INFO: Elapsed time: 60.603s, Critical Path: 19.49s
10:36:15: INFO: 7168 processes: 7129 remote cache hit, 31 internal, 8 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:35:39: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067555550228385792#1:build-log.txt%3A728)
<details><summary>context</summary>
<pre>10:35:25:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
10:35:25:       |   ^~~~~~~~
10:35:25:       |   vsprintf
10:35:39: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:35:39: INFO: Elapsed time: 19.644s, Critical Path: 6.15s
10:35:39: INFO: 4556 processes: 4527 remote cache hit, 29 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:36:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067555549397913600#1:build-log.txt%3A891)
<details><summary>context</summary>
<pre>10:36:19: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:36:19: /usr/bin/ld: /tmp/go-link-64481425/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
10:36:19: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:36:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:25: INFO: Elapsed time: 56.100s, Critical Path: 21.76s
10:36:25: INFO: 7029 processes: 6990 remote cache hit, 31 internal, 8 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:58:25 &#43;0000 UTC_: <code>10:03:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067544978329964544#1:build-log.txt%3A731)
<details><summary>context</summary>
<pre>10:02:53: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:02:53: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
10:02:53: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:03:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:03:21: INFO: Elapsed time: 33.294s, Critical Path: 16.64s
10:03:21: INFO: 5571 processes: 5540 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:58:25 &#43;0000 UTC_: <code>10:03:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067544979118493696#1:build-log.txt%3A744)
<details><summary>context</summary>
<pre>10:03:01: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
10:03:01: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:03:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:03:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:03:21: INFO: Elapsed time: 33.488s, Critical Path: 18.37s
10:03:21: INFO: 6654 processes: 6624 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:58:23 &#43;0000 UTC_: <code>10:02:10: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067544980003491840#1:build-log.txt%3A723)
<details><summary>context</summary>
<pre>10:02:00:    74 |   asprintf (&amp;err-&gt;error, &#34;%s: &#34;
10:02:00:       |   ^~~~~~~~
10:02:00:       |   vsprintf
10:02:10: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:02:10: INFO: Elapsed time: 12.015s, Critical Path: 5.91s
10:02:10: INFO: 4151 processes: 4121 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:57:15 &#43;0000 UTC_: <code>10:04:54: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067544977449160704#1:build-log.txt%3A750)
<details><summary>context</summary>
<pre>10:04:48: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:04:48: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
10:04:48: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:04:54: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:04:54: INFO: Elapsed time: 18.454s, Critical Path: 15.36s
10:04:54: INFO: 4345 processes: 4314 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:53:46 &#43;0000 UTC_: <code>10:00:53: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067544976908095488#1:build-log.txt%3A785)
<details><summary>context</summary>
<pre>10:00:51: /usr/bin/ld: /tmp/go-link-1535296558/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
10:00:51: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:00:53: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:00:53: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:00:53: INFO: Elapsed time: 52.285s, Critical Path: 17.27s
10:00:53: INFO: 5427 processes: 5393 remote cache hit, 31 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:27:25 &#43;0000 UTC_: <code>08:32:47: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067524389980606464#1:build-log.txt%3A756)
<details><summary>context</summary>
<pre>08:32:31: /usr/bin/ld: /tmp/go-link-2092130017/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:32:31: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:32:47: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:32:47: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:32:47: INFO: Elapsed time: 86.448s, Critical Path: 26.92s
08:32:47: INFO: 5776 processes: 5740 remote cache hit, 31 internal, 5 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:27:18 &#43;0000 UTC_: <code>08:33:30: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067524389368238080#1:build-log.txt%3A799)
<details><summary>context</summary>
<pre>08:33:25: /usr/bin/ld: /tmp/go-link-3209859693/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:33:25: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:33:30: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:33:30: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:33:30: INFO: Elapsed time: 12.908s, Critical Path: 5.37s
08:33:30: INFO: 2370 processes: 2341 remote cache hit, 29 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**11x**: _2026-06-18 09:49:11 &#43;0000 UTC_: <code>09:56:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067544975490420736#1:build-log.txt%3A748)
<details>
<summary>all...</summary>

* _2026-06-18 10:31:02 &#43;0000 UTC_: <code>10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067555549053980672#1:build-log.txt%3A836)
<details><summary>context</summary>
<pre>10:36:58: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:37:07: INFO: Elapsed time: 56.856s, Critical Path: 29.82s
10:37:07: INFO: 6150 processes: 6111 remote cache hit, 31 internal, 8 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:43: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067555550219997184#1:build-log.txt%3A768)
<details><summary>context</summary>
<pre>10:36:33: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
10:36:43: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:43: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:43: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
10:36:43: INFO: Elapsed time: 26.025s, Critical Path: 8.14s
10:36:43: INFO: 5882 processes: 5851 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:49:12 &#43;0000 UTC_: <code>09:57:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067544975045824512#1:build-log.txt%3A832)
<details><summary>context</summary>
<pre>09:57:23: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:57:23: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:57:30: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:57:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:57:30: INFO: Elapsed time: 103.248s, Critical Path: 48.84s
09:57:30: INFO: 7036 processes: 6989 remote cache hit, 31 internal, 16 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:49:11 &#43;0000 UTC_: <code>09:56:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067544975490420736#1:build-log.txt%3A748)
<details><summary>context</summary>
<pre>09:55:59: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:56:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:56:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:56:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:56:32: INFO: Elapsed time: 41.479s, Critical Path: 28.15s
09:56:32: INFO: 5053 processes: 5023 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:49:06 &#43;0000 UTC_: <code>09:54:00: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067544975918239744#1:build-log.txt%3A811)
<details><summary>context</summary>
<pre>09:53:55: /usr/bin/ld: /tmp/go-link-788969115/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:53:55: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:54:00: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:54:00: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:54:00: INFO: Elapsed time: 37.003s, Critical Path: 29.62s
09:54:00: INFO: 7214 processes: 7172 remote cache hit, 30 internal, 12 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:49:03 &#43;0000 UTC_: <code>09:59:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067544976639660032#1:build-log.txt%3A813)
<details><summary>context</summary>
<pre>09:59:32: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:59:32: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:59:37: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:59:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:59:37: INFO: Elapsed time: 63.420s, Critical Path: 4.49s
09:59:37: INFO: 3215 processes: 3184 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:20: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067536881171566592#1:build-log.txt%3A734)
<details><summary>context</summary>
<pre>09:20:07:       |   ^~~~~~~~
09:20:07:       |   vsprintf
09:20:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:20:20: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:20:20: INFO: Elapsed time: 24.151s, Critical Path: 11.59s
09:20:20: INFO: 3572 processes: 3541 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:21:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067536881553248256#1:build-log.txt%3A748)
<details><summary>context</summary>
<pre>09:21:18: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:21:30: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:21:30: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:21:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:21:30: INFO: Elapsed time: 23.568s, Critical Path: 10.23s
09:21:30: INFO: 6079 processes: 6048 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:16:47 &#43;0000 UTC_: <code>09:21:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067536881284812800#1:build-log.txt%3A781)
<details><summary>context</summary>
<pre>09:21:35: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:21:35: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:21:47: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:21:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:21:47: INFO: Elapsed time: 21.679s, Critical Path: 9.97s
09:21:47: INFO: 4561 processes: 4530 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:34:54 &#43;0000 UTC_: <code>08:38:26: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067526318710329344#1:build-log.txt%3A759)
<details><summary>context</summary>
<pre>08:38:17:       |   vsprintf
08:38:26: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:38:26: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:38:26: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:38:26: INFO: Elapsed time: 10.243s, Critical Path: 6.20s
08:38:26: INFO: 3493 processes: 3463 remote cache hit, 30 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:33:42: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067524389292740608#1:build-log.txt%3A742)
<details><summary>context</summary>
<pre>08:33:24:       |   ^~~~~~~~
08:33:24:       |   vsprintf
08:33:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:33:42: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
08:33:42: INFO: Elapsed time: 31.326s, Critical Path: 13.97s
08:33:42: INFO: 4477 processes: 4448 remote cache hit, 29 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**3x**: _2026-06-18 09:31:48 &#43;0000 UTC_: <code>09:36:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067536884770279424#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-18 09:49:03 &#43;0000 UTC_: <code>09:53:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067544976035680256#1:build-log.txt%3A697)
<details><summary>context</summary>
<pre>09:52:53: /usr/bin/ld: /tmp/go-link-3209859693/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:52:53: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:53:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:53:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x
09:53:13: INFO: Elapsed time: 21.219s, Critical Path: 4.40s
09:53:13: INFO: 4483 processes: 4452 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:31:48 &#43;0000 UTC_: <code>09:36:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067536884770279424#1:build-log.txt%3A729)
<details><summary>context</summary>
<pre>09:36:08: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:36:08: /usr/bin/ld: /tmp/go-link-398059591/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:36:08: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:36:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x
09:36:08: INFO: Elapsed time: 35.354s, Critical Path: 17.09s
09:36:08: INFO: 4539 processes: 4507 remote cache hit, 30 internal, 2 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 08:27:17 &#43;0000 UTC_: <code>08:33:02: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067524389678616576#1:build-log.txt%3A750)
<details><summary>context</summary>
<pre>08:32:51: cgo_lookup_cgo.cgo2.c:(.text&#43;0x85): warning: Using &#39;getpwuid_r&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:32:51: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:32:51: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:33:02: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x
08:33:02: INFO: Elapsed time: 32.686s, Critical Path: 5.25s
08:33:02: INFO: 5439 processes: 5405 remote cache hit, 31 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**2x**: _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067536881062514688#1:build-log.txt%3A728)
<details>
<summary>all...</summary>

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067536881062514688#1:build-log.txt%3A728)
<details><summary>context</summary>
<pre>09:20:04:       |   ^~~~~~~~
09:20:04:       |   vsprintf
09:20:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:20:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:20:15: INFO: Elapsed time: 12.138s, Critical Path: 5.29s
09:20:15: INFO: 2838 processes: 2807 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:22:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067536881393864704#1:build-log.txt%3A747)
<details><summary>context</summary>
<pre>09:22:07: /usr/bin/ld: /tmp/go-link-853841326/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
09:22:07: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
09:22:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:22:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x
09:22:16: INFO: Elapsed time: 19.422s, Critical Path: 12.47s
09:22:16: INFO: 3983 processes: 3952 remote cache hit, 31 internal.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**1x**: _2026-06-18 08:26:08 &#43;0000 UTC_: <code>08:32:27: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: Running nogo on //vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067524110417661952#1:build-log.txt%3A744)
<details>
<summary>all...</summary>

* _2026-06-18 08:26:08 &#43;0000 UTC_: <code>08:32:27: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: Running nogo on //vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067524110417661952#1:build-log.txt%3A744)
<details><summary>context</summary>
<pre>08:32:26: /usr/bin/ld: /tmp/go-link-3339655621/000004.o: in function `_cgo_77133bf98b3a_C2func_getaddrinfo&#39;:
08:32:26: cgo_unix_cgo.cgo2.c:(.text&#43;0x81): warning: Using &#39;getaddrinfo&#39; in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
08:32:27: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x
08:32:27: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: Running nogo on //vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x
08:32:30: INFO: Elapsed time: 89.468s, Critical Path: 25.90s
08:32:30: INFO: 5222 processes: 5189 remote cache hit, 30 internal, 3 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 1.45%) </summary>

<hr/>

**1x**: _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)
<details>
<summary>all...</summary>

* _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)
<details><summary>context</summary>
<pre>Error: cannot remove container 8d2ad360e3c6cbb7df3ce927bc2970362c7b6d0e538300b87b8f89a7e3f07d2e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 418d3eebf1f9cff01fa9e12d4bef1cfb14a8d7caaec3fa749d15525b64730a3a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 64299780286f8bb3d71dcf5c99da1840a93124aef44d0040e27f8bbd20f625db as it could not be stopped: given PID did not die within timeout
Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 21 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 1.45%) </summary>

<hr/>

**1x**: _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)
<details>
<summary>all...</summary>

* _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)
<details><summary>context</summary>
<pre>08:43:22: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted:
08:43:22: INFO: Elapsed time: 25.777s
08:43:22: INFO: 0 processes.
08:43:22: ERROR: Build failed. Not running target
make: *** [Makefile:25: bazel-build] Error 1
&#43; EXIT_VALUE=2
&#43; set &#43;o xtrace</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 1.45%) </summary>

<hr/>

**1x**: _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)
<details>
<summary>all...</summary>

* _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)
<details><summary>context</summary>
<pre>13:00:48: rsync: [sender] failed to connect to 127.0.0.1 (127.0.0.1): Connection refused (111)
13:00:48: rsync error: error in socket IO (code 10) at clientserver.c(139) [sender=3.4.1]
13:01:16: &#43; rm -f /tmp/kubevirt.deploy.7kv4
make: *** [Makefile:189: cluster-sync] Error 10
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 1.45%) </summary>

<hr/>

**1x**: _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)
<details>
<summary>all...</summary>

* _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)
<details><summary>context</summary>
<pre>22:46:53: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
22:46:53: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
22:46:58: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1416787248/2&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260621%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260621T224655Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=bb87a81e41fb2140fda6b6fe1f39c827129addef871d5ba06fb6376fca50a76b&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1782082915~hmac=7bd114a9b9115fd6386aa0672cdb3b2a0f43fbdd75fda521d5742b48c67b52ea&#34;: EOF)
make: *** [Makefile:174: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 1.45%)

<details>
<summary> kind cluster creation failure (1x / 1.45%) </summary>

<hr/>

**1x**: _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)
<details><summary>context</summary>
<pre>11:55:01:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
11:55:01:  • Preparing nodes 📦 📦   ...
11:55:03:  ✗ Preparing nodes 📦 📦
11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
11:55:03:
11:55:03: Stack Trace:
11:55:03: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (69x / 100.00%)


#### external (68x / 98.55%)

<details>
<summary> bazel remote cache blob fetch failure (64x / 92.75%) </summary>

<hr/>

**26x**: _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:31:34: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067524113043296256#1:build-log.txt%3A727)
<details>
<summary>all...</summary>

* _2026-06-18 09:49:09 &#43;0000 UTC_: <code>09:55:37: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067544976245395456#1:build-log.txt%3A729)

* _2026-06-18 09:49:06 &#43;0000 UTC_: <code>09:52:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067544975767244800#1:build-log.txt%3A744)

* _2026-06-18 09:49:04 &#43;0000 UTC_: <code>09:52:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067544975628832768#1:build-log.txt%3A741)

* _2026-06-18 09:49:01 &#43;0000 UTC_: <code>09:56:38: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067544975280705536#1:build-log.txt%3A828)

* _2026-06-18 09:48:57 &#43;0000 UTC_: <code>09:55:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067544975377174528#1:build-log.txt%3A755)

* _2026-06-18 09:48:57 &#43;0000 UTC_: <code>09:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067544975154876416#1:build-log.txt%3A755)

* _2026-06-18 09:34:30 &#43;0000 UTC_: <code>09:40:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067536885676249088#1:build-log.txt%3A815)

* _2026-06-18 09:31:41 &#43;0000 UTC_: <code>09:41:02: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067536883935612928#1:build-log.txt%3A779)

* _2026-06-18 09:28:37 &#43;0000 UTC_: <code>09:32:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067536883151278080#1:build-log.txt%3A772)

* _2026-06-18 09:28:12 &#43;0000 UTC_: <code>09:32:12: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067536882257891328#1:build-log.txt%3A806)

* _2026-06-18 09:17:38 &#43;0000 UTC_: <code>09:21:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067536881737797632#1:build-log.txt%3A731)

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067536881641328640#1:build-log.txt%3A769)

* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:22:17: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067536880865382400#1:build-log.txt%3A741)

* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:20:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067536881456779264#1:build-log.txt%3A742)

* _2026-06-18 09:16:49 &#43;0000 UTC_: <code>09:21:44: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067536880655667200#1:build-log.txt%3A740)

* _2026-06-18 09:16:47 &#43;0000 UTC_: <code>09:21:22: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067536880974434304#1:build-log.txt%3A827)

* _2026-06-18 08:27:17 &#43;0000 UTC_: <code>08:32:22: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067524389779279872#1:build-log.txt%3A805)

* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:34:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067524390601363456#1:build-log.txt%3A746)

* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:33:23: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067524389590536192#1:build-log.txt%3A807)

* _2026-06-18 08:27:13 &#43;0000 UTC_: <code>08:33:38: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067524389867360256#1:build-log.txt%3A813)

* _2026-06-18 08:27:12 &#43;0000 UTC_: <code>08:33:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067524389456318464#1:build-log.txt%3A813)

* _2026-06-18 08:26:13 &#43;0000 UTC_: <code>08:31:57: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067524111256522752#1:build-log.txt%3A722)

* _2026-06-18 08:26:11 &#43;0000 UTC_: <code>08:30:58: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067524112086994944#1:build-log.txt%3A729)

* _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:32:01: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067524113760522240#1:build-log.txt%3A756)

* _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:31:34: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067524113043296256#1:build-log.txt%3A727)

* _2026-06-18 08:26:05 &#43;0000 UTC_: <code>08:31:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067524107909468160#1:build-log.txt%3A811)

</details>

<hr/>

**21x**: _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:34:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067555554041008128#1:build-log.txt%3A767)
<details>
<summary>all...</summary>

* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:35:36: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067555549490188288#1:build-log.txt%3A822)

* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:35:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067555550215802880#1:build-log.txt%3A813)

* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:34:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067555554041008128#1:build-log.txt%3A767)

* _2026-06-18 10:31:06 &#43;0000 UTC_: <code>10:35:00: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067555550236774400#1:build-log.txt%3A742)

* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067555550224191488#1:build-log.txt%3A795)

* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:35:59: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067555549305638912#1:build-log.txt%3A756)

* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067555553206341632#1:build-log.txt%3A757)

* _2026-06-18 10:31:04 &#43;0000 UTC_: <code>10:36:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067555550635233280#1:build-log.txt%3A799)

* _2026-06-18 10:31:03 &#43;0000 UTC_: <code>10:36:48: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067555550232580096#1:build-log.txt%3A735)

* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067555554854703104#1:build-log.txt%3A759)

* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:31: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067555552426201088#1:build-log.txt%3A747)

* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067555551469899776#1:build-log.txt%3A824)

* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:35:39: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067555550228385792#1:build-log.txt%3A728)

* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:36:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067555549397913600#1:build-log.txt%3A891)

* _2026-06-18 09:58:25 &#43;0000 UTC_: <code>10:03:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067544978329964544#1:build-log.txt%3A731)

* _2026-06-18 09:58:25 &#43;0000 UTC_: <code>10:03:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067544979118493696#1:build-log.txt%3A744)

* _2026-06-18 09:58:23 &#43;0000 UTC_: <code>10:02:10: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067544980003491840#1:build-log.txt%3A723)

* _2026-06-18 09:57:15 &#43;0000 UTC_: <code>10:04:54: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067544977449160704#1:build-log.txt%3A750)

* _2026-06-18 09:53:46 &#43;0000 UTC_: <code>10:00:53: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067544976908095488#1:build-log.txt%3A785)

* _2026-06-18 08:27:25 &#43;0000 UTC_: <code>08:32:47: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067524389980606464#1:build-log.txt%3A756)

* _2026-06-18 08:27:18 &#43;0000 UTC_: <code>08:33:30: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067524389368238080#1:build-log.txt%3A799)

</details>

<hr/>

**11x**: _2026-06-18 09:49:11 &#43;0000 UTC_: <code>09:56:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067544975490420736#1:build-log.txt%3A748)
<details>
<summary>all...</summary>

* _2026-06-18 10:31:02 &#43;0000 UTC_: <code>10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067555549053980672#1:build-log.txt%3A836)

* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:43: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067555550219997184#1:build-log.txt%3A768)

* _2026-06-18 09:49:12 &#43;0000 UTC_: <code>09:57:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067544975045824512#1:build-log.txt%3A832)

* _2026-06-18 09:49:11 &#43;0000 UTC_: <code>09:56:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067544975490420736#1:build-log.txt%3A748)

* _2026-06-18 09:49:06 &#43;0000 UTC_: <code>09:54:00: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067544975918239744#1:build-log.txt%3A811)

* _2026-06-18 09:49:03 &#43;0000 UTC_: <code>09:59:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067544976639660032#1:build-log.txt%3A813)

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:20: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067536881171566592#1:build-log.txt%3A734)

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:21:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067536881553248256#1:build-log.txt%3A748)

* _2026-06-18 09:16:47 &#43;0000 UTC_: <code>09:21:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067536881284812800#1:build-log.txt%3A781)

* _2026-06-18 08:34:54 &#43;0000 UTC_: <code>08:38:26: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067526318710329344#1:build-log.txt%3A759)

* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:33:42: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067524389292740608#1:build-log.txt%3A742)

</details>

<hr/>

**3x**: _2026-06-18 09:31:48 &#43;0000 UTC_: <code>09:36:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067536884770279424#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-18 09:49:03 &#43;0000 UTC_: <code>09:53:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067544976035680256#1:build-log.txt%3A697)

* _2026-06-18 09:31:48 &#43;0000 UTC_: <code>09:36:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067536884770279424#1:build-log.txt%3A729)

* _2026-06-18 08:27:17 &#43;0000 UTC_: <code>08:33:02: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067524389678616576#1:build-log.txt%3A750)

</details>

<hr/>

**2x**: _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067536881062514688#1:build-log.txt%3A728)
<details>
<summary>all...</summary>

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067536881062514688#1:build-log.txt%3A728)

* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:22:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067536881393864704#1:build-log.txt%3A747)

</details>

<hr/>

**1x**: _2026-06-18 08:26:08 &#43;0000 UTC_: <code>08:32:27: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: Running nogo on //vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067524110417661952#1:build-log.txt%3A744)
<details>
<summary>all...</summary>

* _2026-06-18 08:26:08 &#43;0000 UTC_: <code>08:32:27: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: Running nogo on //vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067524110417661952#1:build-log.txt%3A744)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 1.45%) </summary>

<hr/>

**1x**: _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)
<details>
<summary>all...</summary>

* _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 1.45%) </summary>

<hr/>

**1x**: _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)
<details>
<summary>all...</summary>

* _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 1.45%) </summary>

<hr/>

**1x**: _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)
<details>
<summary>all...</summary>

* _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 1.45%) </summary>

<hr/>

**1x**: _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)
<details>
<summary>all...</summary>

* _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)

</details>

<hr/>
</details>

#### internal (1x / 1.45%)

<details>
<summary> kind cluster creation failure (1x / 1.45%) </summary>

<hr/>

**1x**: _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (40x / 57.97%)


#### external (39x / 97.50%)

<details>
<summary> bazel remote cache blob fetch failure (36x / 90.00%) </summary>

<hr/>

**16x**: _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:31:34: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067524113043296256#1:build-log.txt%3A727)
<details>
<summary>all...</summary>

* _2026-06-18 09:49:09 &#43;0000 UTC_: <code>09:55:37: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067544976245395456#1:build-log.txt%3A729)

* _2026-06-18 09:49:06 &#43;0000 UTC_: <code>09:52:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067544975767244800#1:build-log.txt%3A744)

* _2026-06-18 09:49:04 &#43;0000 UTC_: <code>09:52:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067544975628832768#1:build-log.txt%3A741)

* _2026-06-18 09:49:01 &#43;0000 UTC_: <code>09:56:38: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067544975280705536#1:build-log.txt%3A828)

* _2026-06-18 09:34:30 &#43;0000 UTC_: <code>09:40:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067536885676249088#1:build-log.txt%3A815)

* _2026-06-18 09:31:41 &#43;0000 UTC_: <code>09:41:02: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067536883935612928#1:build-log.txt%3A779)

* _2026-06-18 09:17:38 &#43;0000 UTC_: <code>09:21:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067536881737797632#1:build-log.txt%3A731)

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067536881641328640#1:build-log.txt%3A769)

* _2026-06-18 09:16:49 &#43;0000 UTC_: <code>09:21:44: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067536880655667200#1:build-log.txt%3A740)

* _2026-06-18 09:16:47 &#43;0000 UTC_: <code>09:21:22: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067536880974434304#1:build-log.txt%3A827)

* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:33:23: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067524389590536192#1:build-log.txt%3A807)

* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:34:32: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067524390601363456#1:build-log.txt%3A746)

* _2026-06-18 08:26:11 &#43;0000 UTC_: <code>08:30:58: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067524112086994944#1:build-log.txt%3A729)

* _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:31:34: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067524113043296256#1:build-log.txt%3A727)

* _2026-06-18 08:26:10 &#43;0000 UTC_: <code>08:32:01: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067524113760522240#1:build-log.txt%3A756)

* _2026-06-18 08:26:05 &#43;0000 UTC_: <code>08:31:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067524107909468160#1:build-log.txt%3A811)

</details>

<hr/>

**12x**: _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:34:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067555554041008128#1:build-log.txt%3A767)
<details>
<summary>all...</summary>

* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:34:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067555554041008128#1:build-log.txt%3A767)

* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067555553206341632#1:build-log.txt%3A757)

* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067555550224191488#1:build-log.txt%3A795)

* _2026-06-18 10:31:04 &#43;0000 UTC_: <code>10:36:24: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067555550635233280#1:build-log.txt%3A799)

* _2026-06-18 10:31:03 &#43;0000 UTC_: <code>10:36:48: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067555550232580096#1:build-log.txt%3A735)

* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067555554854703104#1:build-log.txt%3A759)

* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:35:39: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067555550228385792#1:build-log.txt%3A728)

* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:36:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2067555549397913600#1:build-log.txt%3A891)

* _2026-06-18 09:58:25 &#43;0000 UTC_: <code>10:03:21: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067544979118493696#1:build-log.txt%3A744)

* _2026-06-18 09:58:25 &#43;0000 UTC_: <code>10:03:20: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute/2067544978329964544#1:build-log.txt%3A731)

* _2026-06-18 09:58:23 &#43;0000 UTC_: <code>10:02:10: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-operator/2067544980003491840#1:build-log.txt%3A723)

* _2026-06-18 08:27:25 &#43;0000 UTC_: <code>08:32:47: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-compute/2067524389980606464#1:build-log.txt%3A756)

</details>

<hr/>

**5x**: _2026-06-18 09:49:03 &#43;0000 UTC_: <code>09:59:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067544976639660032#1:build-log.txt%3A813)
<details>
<summary>all...</summary>

* _2026-06-18 10:31:02 &#43;0000 UTC_: <code>10:37:07: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067555549053980672#1:build-log.txt%3A836)

* _2026-06-18 09:49:12 &#43;0000 UTC_: <code>09:57:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067544975045824512#1:build-log.txt%3A832)

* _2026-06-18 09:49:03 &#43;0000 UTC_: <code>09:59:37: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-operator/2067544976639660032#1:build-log.txt%3A813)

* _2026-06-18 09:16:47 &#43;0000 UTC_: <code>09:21:47: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-compute/2067536881284812800#1:build-log.txt%3A781)

* _2026-06-18 08:34:54 &#43;0000 UTC_: <code>08:38:26: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-windows2016/2067526318710329344#1:build-log.txt%3A759)

</details>

<hr/>

**2x**: _2026-06-18 09:31:48 &#43;0000 UTC_: <code>09:36:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067536884770279424#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-18 09:31:48 &#43;0000 UTC_: <code>09:36:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial/2067536884770279424#1:build-log.txt%3A729)

* _2026-06-18 08:27:17 &#43;0000 UTC_: <code>08:33:02: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067524389678616576#1:build-log.txt%3A750)

</details>

<hr/>

**1x**: _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:22:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067536881393864704#1:build-log.txt%3A747)
<details>
<summary>all...</summary>

* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:22:16: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-operator/2067536881393864704#1:build-log.txt%3A747)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 2.50%) </summary>

<hr/>

**1x**: _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)
<details>
<summary>all...</summary>

* _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 2.50%) </summary>

<hr/>

**1x**: _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)
<details>
<summary>all...</summary>

* _2026-06-17 07:01:38 &#43;0000 UTC_: <code>Error: cannot remove container e672de5d181574bc76d99016452be15084c22ab8f3d5895e5e5f8d0a986dc6f7 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17976/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2067140165721133056#1:build-log.txt%3A1868)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 2.50%) </summary>

<hr/>

**1x**: _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)
<details>
<summary>all...</summary>

* _2026-06-16 12:46:25 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 10</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2066864704415862784#1:build-log.txt%3A1383)

</details>

<hr/>
</details>

#### internal (1x / 2.50%)

<details>
<summary> kind cluster creation failure (1x / 2.50%) </summary>

<hr/>

**1x**: _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)
<details>
<summary>all...</summary>

* _2026-06-16 11:48:08 &#43;0000 UTC_: <code>11:55:03: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17983/pull-kubevirt-e2e-kind-1.34-sev/2066849993645363200#1:build-log.txt%3A729)

</details>

<hr/>
</details>

### sig-network (16x / 23.19%)


#### external (16x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (16x / 100.00%) </summary>

<hr/>

**6x**: _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:35:36: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067555549490188288#1:build-log.txt%3A822)
<details>
<summary>all...</summary>

* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:35:36: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067555549490188288#1:build-log.txt%3A822)

* _2026-06-18 10:31:06 &#43;0000 UTC_: <code>10:35:00: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067555550236774400#1:build-log.txt%3A742)

* _2026-06-18 10:31:05 &#43;0000 UTC_: <code>10:35:59: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067555549305638912#1:build-log.txt%3A756)

* _2026-06-18 10:30:59 &#43;0000 UTC_: <code>10:36:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067555551469899776#1:build-log.txt%3A824)

* _2026-06-18 09:53:46 &#43;0000 UTC_: <code>10:00:53: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067544976908095488#1:build-log.txt%3A785)

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:15: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: Running nogo on //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067536881062514688#1:build-log.txt%3A728)

</details>

<hr/>

**2x**: _2026-06-18 09:49:06 &#43;0000 UTC_: <code>09:54:00: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067544975918239744#1:build-log.txt%3A811)
<details>
<summary>all...</summary>

* _2026-06-18 09:49:06 &#43;0000 UTC_: <code>09:54:00: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067544975918239744#1:build-log.txt%3A811)

* _2026-06-18 08:27:16 &#43;0000 UTC_: <code>08:33:42: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067524389292740608#1:build-log.txt%3A742)

</details>

<hr/>

**6x**: _2026-06-18 09:48:57 &#43;0000 UTC_: <code>09:55:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067544975377174528#1:build-log.txt%3A755)
<details>
<summary>all...</summary>

* _2026-06-18 09:48:57 &#43;0000 UTC_: <code>09:55:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067544975377174528#1:build-log.txt%3A755)

* _2026-06-18 09:48:57 &#43;0000 UTC_: <code>09:53:48: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067544975154876416#1:build-log.txt%3A755)

* _2026-06-18 09:28:12 &#43;0000 UTC_: <code>09:32:12: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067536882257891328#1:build-log.txt%3A806)

* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:20:42: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067536881456779264#1:build-log.txt%3A742)

* _2026-06-18 09:16:50 &#43;0000 UTC_: <code>09:22:17: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2067536880865382400#1:build-log.txt%3A741)

* _2026-06-18 08:27:17 &#43;0000 UTC_: <code>08:32:22: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-network/2067524389779279872#1:build-log.txt%3A805)

</details>

<hr/>

**2x**: _2026-06-18 08:27:18 &#43;0000 UTC_: <code>08:33:30: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067524389368238080#1:build-log.txt%3A799)
<details>
<summary>all...</summary>

* _2026-06-18 08:27:18 &#43;0000 UTC_: <code>08:33:30: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-network/2067524389368238080#1:build-log.txt%3A799)

* _2026-06-18 08:26:08 &#43;0000 UTC_: <code>08:32:27: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: Running nogo on //vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-network/2067524110417661952#1:build-log.txt%3A744)

</details>

<hr/>
</details>

### sig-storage (12x / 17.39%)


#### external (12x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (12x / 100.00%) </summary>

<hr/>

**4x**: _2026-06-18 09:49:11 &#43;0000 UTC_: <code>09:56:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067544975490420736#1:build-log.txt%3A748)
<details>
<summary>all...</summary>

* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:43: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067555550219997184#1:build-log.txt%3A768)

* _2026-06-18 09:49:11 &#43;0000 UTC_: <code>09:56:32: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067544975490420736#1:build-log.txt%3A748)

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:20:20: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067536881171566592#1:build-log.txt%3A734)

* _2026-06-18 09:16:55 &#43;0000 UTC_: <code>09:21:30: ERROR: /root/go/src/kubevirt.io/kubevirt/tests/BUILD.bazel:43:8 Validating nogo output for //tests:go_default_test failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067536881553248256#1:build-log.txt%3A748)

</details>

<hr/>

**4x**: _2026-06-18 08:27:12 &#43;0000 UTC_: <code>08:33:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067524389456318464#1:build-log.txt%3A813)
<details>
<summary>all...</summary>

* _2026-06-18 09:28:37 &#43;0000 UTC_: <code>09:32:25: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067536883151278080#1:build-log.txt%3A772)

* _2026-06-18 08:27:13 &#43;0000 UTC_: <code>08:33:38: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067524389867360256#1:build-log.txt%3A813)

* _2026-06-18 08:27:12 &#43;0000 UTC_: <code>08:33:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067524389456318464#1:build-log.txt%3A813)

* _2026-06-18 08:26:13 &#43;0000 UTC_: <code>08:31:57: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067524111256522752#1:build-log.txt%3A722)

</details>

<hr/>

**3x**: _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:35:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067555550215802880#1:build-log.txt%3A813)
<details>
<summary>all...</summary>

* _2026-06-18 10:31:07 &#43;0000 UTC_: <code>10:35:08: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11 Validating nogo output for //vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.34-sig-storage/2067555550215802880#1:build-log.txt%3A813)

* _2026-06-18 10:31:01 &#43;0000 UTC_: <code>10:36:31: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067555552426201088#1:build-log.txt%3A747)

* _2026-06-18 09:57:15 &#43;0000 UTC_: <code>10:04:54: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/BUILD.bazel:3:11: GoCompilePkg vendor/kubevirt.io/qe-tools/pkg/ginkgo-reporters/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 63ff9f82f95febd57d3163e00bcc3832398c70ff31c4b7d660cf00a83b111756/5978 for bazel-out/k8-fastbuild/bin/vendor/kubevirt.io/qe-tools/pkg/polarion-xml/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.36-sig-storage/2067544977449160704#1:build-log.txt%3A750)

</details>

<hr/>

**1x**: _2026-06-18 09:49:03 &#43;0000 UTC_: <code>09:53:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067544976035680256#1:build-log.txt%3A697)
<details>
<summary>all...</summary>

* _2026-06-18 09:49:03 &#43;0000 UTC_: <code>09:53:13: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/rhobs/operator-observability-toolkit/pkg/testutil/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: 2dc5dd3e503bba0ed460cbde692789b79995fe03f7317c69d725681276a57d05/31872 for bazel-out/k8-fastbuild/bin/vendor/github.com/grafana/regexp/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17910/pull-kubevirt-e2e-k8s-1.35-sig-storage/2067544976035680256#1:build-log.txt%3A697)

</details>

<hr/>
</details>

### sig-performance (1x / 1.45%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)
<details>
<summary>all...</summary>

* _2026-06-18 08:42:15 &#43;0000 UTC_: <code>08:43:22: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17263/pull-kubevirt-e2e-k8s-1.34-sig-performance/2067528172320067584#1:build-log.txt%3A86)

</details>

<hr/>
</details>

Last updated: 2026-06-22 04:44:04
