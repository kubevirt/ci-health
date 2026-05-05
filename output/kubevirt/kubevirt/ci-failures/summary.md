



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-04-29 (6x / 66.67%)


#### external (6x / 100.00%)

<details>
<summary> download failure in context (6x / 100.00%) </summary>

<hr/>

**4x**: _2026-04-29 16:24:52 &#43;0000 UTC_: <code>16:26:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17652/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049525179276595200#1:build-log.txt%3A307)
<details>
<summary>all...</summary>

* _2026-04-29 16:24:52 &#43;0000 UTC_: <code>16:26:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17652/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049525179276595200#1:build-log.txt%3A307)

* _2026-04-29 12:01:51 &#43;0000 UTC_: <code>12:02:55: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17647/pull-kubevirt-e2e-kind-sriov/2049458981394452480#1:build-log.txt%3A307)

* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:33:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2049451505487974400#1:build-log.txt%3A307)

* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:53:16: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049451503768309760#1:build-log.txt%3A5396)

</details>

<hr/>

**1x**: _2026-04-29 11:32:06 &#43;0000 UTC_: <code>11:34:30: ERROR: Analysis of target &#39;//cmd/cniplugins/passt-binding/cmd:network-passt-binding-cni-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2049451511481634816#1:build-log.txt%3A564)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:06 &#43;0000 UTC_: <code>11:34:30: ERROR: Analysis of target &#39;//cmd/cniplugins/passt-binding/cmd:network-passt-binding-cni-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2049451511481634816#1:build-log.txt%3A564)

</details>

<hr/>

**1x**: _2026-04-29 11:32:03 &#43;0000 UTC_: <code>11:34:28: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-kind-sriov-1.8/2049451494389846016#1:build-log.txt%3A529)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:03 &#43;0000 UTC_: <code>11:34:28: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-kind-sriov-1.8/2049451494389846016#1:build-log.txt%3A529)

</details>

<hr/>
</details>

### 2026-04-28 (3x / 33.33%)


#### external (2x / 66.67%)

<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-04-28 09:53:40 &#43;0000 UTC_: <code>10:11:07: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17049/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049064376685563904#1:build-log.txt%3A4739)
<details>
<summary>all...</summary>

* _2026-04-28 09:53:40 &#43;0000 UTC_: <code>10:11:07: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17049/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049064376685563904#1:build-log.txt%3A4739)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)
<details>
<summary>all...</summary>

* _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)

</details>

<hr/>
</details>

#### internal (1x / 33.33%)

<details>
<summary> make cluster lifecycle target failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-04-28 15:26:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17631/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049148231257952256#1:build-log.txt%3A1217)
<details>
<summary>all...</summary>

* _2026-04-28 15:26:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17631/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049148231257952256#1:build-log.txt%3A1217)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (8x / 88.89%)

<details>
<summary> download failure in context (7x / 77.78%) </summary>

<hr/>

**5x**: _2026-04-29 16:24:52 &#43;0000 UTC_: <code>16:26:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17652/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049525179276595200#1:build-log.txt%3A307)
<details>
<summary>all...</summary>

* _2026-04-29 16:24:52 &#43;0000 UTC_: <code>16:26:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17652/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049525179276595200#1:build-log.txt%3A307)
<details><summary>context</summary>
<pre>16:26:02: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted:
16:26:02: INFO: Elapsed time: 0.338s
16:26:02: INFO: 0 processes.
16:26:02: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


* _2026-04-29 12:01:51 &#43;0000 UTC_: <code>12:02:55: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17647/pull-kubevirt-e2e-kind-sriov/2049458981394452480#1:build-log.txt%3A307)
<details><summary>context</summary>
<pre>12:02:55: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted:
12:02:55: INFO: Elapsed time: 0.522s
12:02:55: INFO: 0 processes.
12:02:55: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:33:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2049451505487974400#1:build-log.txt%3A307)
<details><summary>context</summary>
<pre>11:33:02: ERROR: Analysis of target &#39;//:gazelle&#39; failed; build aborted:
11:33:02: INFO: Elapsed time: 0.529s
11:33:02: INFO: 0 processes.
11:33:02: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:53:16: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049451503768309760#1:build-log.txt%3A5396)
<details><summary>context</summary>
<pre>11:53:16: ERROR: Analysis of target &#39;//:buildifier&#39; failed; build aborted:
11:53:16: INFO: Elapsed time: 0.239s
11:53:16: INFO: 0 processes.
11:53:16: ERROR: Build failed. Not running target
make: *** [Makefile:28: bazel-build-functests] Error 1
&#43; ret=2
&#43; check_for_panics</pre>
</details>


* _2026-04-28 09:53:40 &#43;0000 UTC_: <code>10:11:07: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17049/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049064376685563904#1:build-log.txt%3A4739)
<details><summary>context</summary>
<pre>10:11:07: ERROR: Analysis of target &#39;//:buildifier&#39; failed; build aborted:
10:11:07: INFO: Elapsed time: 0.667s
10:11:07: INFO: 0 processes.
10:11:07: ERROR: Build failed. Not running target
make: *** [Makefile:28: bazel-build-functests] Error 1
&#43; ret=2
&#43; check_for_panics</pre>
</details>


</details>

<hr/>

**1x**: _2026-04-29 11:32:06 &#43;0000 UTC_: <code>11:34:30: ERROR: Analysis of target &#39;//cmd/cniplugins/passt-binding/cmd:network-passt-binding-cni-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2049451511481634816#1:build-log.txt%3A564)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:06 &#43;0000 UTC_: <code>11:34:30: ERROR: Analysis of target &#39;//cmd/cniplugins/passt-binding/cmd:network-passt-binding-cni-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2049451511481634816#1:build-log.txt%3A564)
<details><summary>context</summary>
<pre>11:34:30: Repository rule rpm defined at:
11:34:30:   /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazeldnf/internal/rpm.bzl:48:22: in &lt;toplevel&gt;
11:34:30: ERROR: /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/busybox_single/BUILD.bazel:5:18: @busybox_single//:busybox_single depends on @copy_to_directory_linux_amd64//:copy_to_directory_toolchain in repository @copy_to_directory_linux_amd64 which failed to fetch. no such package &#39;@copy_to_directory_linux_amd64//&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/copy_to_directory-linux_amd64] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/copy_to_directory_linux_amd64/copy_to_directory: GET returned 502 Bad Gateway or Proxy Error
11:34:30: ERROR: Analysis of target &#39;//cmd/cniplugins/passt-binding/cmd:network-passt-binding-cni-image&#39; failed; build aborted:
11:34:30: INFO: Elapsed time: 1.200s
11:34:30: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**1x**: _2026-04-29 11:32:03 &#43;0000 UTC_: <code>11:34:28: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-kind-sriov-1.8/2049451494389846016#1:build-log.txt%3A529)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:03 &#43;0000 UTC_: <code>11:34:28: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-kind-sriov-1.8/2049451494389846016#1:build-log.txt%3A529)
<details><summary>context</summary>
<pre>11:34:28: Repository rule rpm defined at:
11:34:28:   /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazeldnf/internal/rpm.bzl:48:22: in &lt;toplevel&gt;
11:34:28: ERROR: /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/alpine_with_test_tooling_single/BUILD.bazel:5:18: @alpine_with_test_tooling_single//:alpine_with_test_tooling_single depends on @copy_to_directory_linux_amd64//:copy_to_directory_toolchain in repository @copy_to_directory_linux_amd64 which failed to fetch. no such package &#39;@copy_to_directory_linux_amd64//&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/copy_to_directory-linux_amd64] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/copy_to_directory_linux_amd64/copy_to_directory: GET returned 502 Bad Gateway or Proxy Error
11:34:28: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted:
11:34:28: INFO: Elapsed time: 0.483s
11:34:28: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)
<details>
<summary>all...</summary>

* _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)
<details><summary>context</summary>
<pre>09:55:54: 	File &#34;/tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl&#34;, line 132, column 45, in _http_archive_impl
09:55:54: 		download_info = ctx.download_and_extract(
09:55:54: Error in download_and_extract: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error
09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 11.11%)

<details>
<summary> make cluster lifecycle target failure (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-04-28 15:26:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17631/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049148231257952256#1:build-log.txt%3A1217)
<details>
<summary>all...</summary>

* _2026-04-28 15:26:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17631/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049148231257952256#1:build-log.txt%3A1217)
<details><summary>context</summary>
<pre>15:35:18: gzip: stdin: not in gzip format
15:35:18: tar: Child returned status 1
15:35:18: tar: Error is not recoverable: exiting now
make: *** [Makefile:173: cluster-up] Error 2
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (4x / 44.44%)


#### external (4x / 100.00%)

<details>
<summary> download failure in context (3x / 75.00%) </summary>

<hr/>

**3x**: _2026-04-29 16:24:52 &#43;0000 UTC_: <code>16:26:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17652/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049525179276595200#1:build-log.txt%3A307)
<details>
<summary>all...</summary>

* _2026-04-29 16:24:52 &#43;0000 UTC_: <code>16:26:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17652/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049525179276595200#1:build-log.txt%3A307)

* _2026-04-29 12:01:51 &#43;0000 UTC_: <code>12:02:55: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17647/pull-kubevirt-e2e-kind-sriov/2049458981394452480#1:build-log.txt%3A307)

* _2026-04-28 09:53:40 &#43;0000 UTC_: <code>10:11:07: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17049/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049064376685563904#1:build-log.txt%3A4739)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)
<details>
<summary>all...</summary>

* _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)

</details>

<hr/>
</details>

### release-1.8 (5x / 55.56%)


#### external (4x / 80.00%)

<details>
<summary> download failure in context (4x / 80.00%) </summary>

<hr/>

**2x**: _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:33:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2049451505487974400#1:build-log.txt%3A307)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:33:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2049451505487974400#1:build-log.txt%3A307)

* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:53:16: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049451503768309760#1:build-log.txt%3A5396)

</details>

<hr/>

**1x**: _2026-04-29 11:32:06 &#43;0000 UTC_: <code>11:34:30: ERROR: Analysis of target &#39;//cmd/cniplugins/passt-binding/cmd:network-passt-binding-cni-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2049451511481634816#1:build-log.txt%3A564)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:06 &#43;0000 UTC_: <code>11:34:30: ERROR: Analysis of target &#39;//cmd/cniplugins/passt-binding/cmd:network-passt-binding-cni-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2049451511481634816#1:build-log.txt%3A564)

</details>

<hr/>

**1x**: _2026-04-29 11:32:03 &#43;0000 UTC_: <code>11:34:28: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-kind-sriov-1.8/2049451494389846016#1:build-log.txt%3A529)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:03 &#43;0000 UTC_: <code>11:34:28: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-kind-sriov-1.8/2049451494389846016#1:build-log.txt%3A529)

</details>

<hr/>
</details>

#### internal (1x / 20.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-28 15:26:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17631/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049148231257952256#1:build-log.txt%3A1217)
<details>
<summary>all...</summary>

* _2026-04-28 15:26:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17631/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049148231257952256#1:build-log.txt%3A1217)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (6x / 66.67%)


#### external (5x / 83.33%)

<details>
<summary> download failure in context (4x / 66.67%) </summary>

<hr/>

**3x**: _2026-04-29 16:24:52 &#43;0000 UTC_: <code>16:26:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17652/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049525179276595200#1:build-log.txt%3A307)
<details>
<summary>all...</summary>

* _2026-04-29 16:24:52 &#43;0000 UTC_: <code>16:26:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17652/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049525179276595200#1:build-log.txt%3A307)

* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:53:16: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049451503768309760#1:build-log.txt%3A5396)

* _2026-04-28 09:53:40 &#43;0000 UTC_: <code>10:11:07: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17049/pull-kubevirt-e2e-k8s-1.35-sig-compute/2049064376685563904#1:build-log.txt%3A4739)

</details>

<hr/>

**1x**: _2026-04-29 11:32:06 &#43;0000 UTC_: <code>11:34:30: ERROR: Analysis of target &#39;//cmd/cniplugins/passt-binding/cmd:network-passt-binding-cni-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2049451511481634816#1:build-log.txt%3A564)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:06 &#43;0000 UTC_: <code>11:34:30: ERROR: Analysis of target &#39;//cmd/cniplugins/passt-binding/cmd:network-passt-binding-cni-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2049451511481634816#1:build-log.txt%3A564)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)
<details>
<summary>all...</summary>

* _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)

</details>

<hr/>
</details>

#### internal (1x / 16.67%)

<details>
<summary> make cluster lifecycle target failure (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-04-28 15:26:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17631/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049148231257952256#1:build-log.txt%3A1217)
<details>
<summary>all...</summary>

* _2026-04-28 15:26:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17631/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049148231257952256#1:build-log.txt%3A1217)

</details>

<hr/>
</details>

### sig-network (3x / 33.33%)


#### external (3x / 100.00%)

<details>
<summary> download failure in context (3x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-29 11:32:03 &#43;0000 UTC_: <code>11:34:28: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-kind-sriov-1.8/2049451494389846016#1:build-log.txt%3A529)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:03 &#43;0000 UTC_: <code>11:34:28: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-kind-sriov-1.8/2049451494389846016#1:build-log.txt%3A529)

</details>

<hr/>

**2x**: _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:33:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2049451505487974400#1:build-log.txt%3A307)
<details>
<summary>all...</summary>

* _2026-04-29 12:01:51 &#43;0000 UTC_: <code>12:02:55: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17647/pull-kubevirt-e2e-kind-sriov/2049458981394452480#1:build-log.txt%3A307)

* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:33:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2049451505487974400#1:build-log.txt%3A307)

</details>

<hr/>
</details>

Last updated: 2026-05-05 09:41:52
