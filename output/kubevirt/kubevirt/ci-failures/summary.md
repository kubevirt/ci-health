



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-04-29 (5x / 35.71%)


#### external (5x / 100.00%)

<details>
<summary> download failure in context (5x / 100.00%) </summary>

<hr/>

**3x**: _2026-04-29 12:01:51 &#43;0000 UTC_: <code>12:02:55: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17647/pull-kubevirt-e2e-kind-sriov/2049458981394452480#1:build-log.txt%3A307)
<details>
<summary>all...</summary>

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

### 2026-04-28 (6x / 42.86%)


#### external (2x / 33.33%)

<details>
<summary> download failure in context (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-04-28 09:06:28 &#43;0000 UTC_: <code>09:21:58: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049052505228185600#1:build-log.txt%3A5334)
<details>
<summary>all...</summary>

* _2026-04-28 09:06:28 &#43;0000 UTC_: <code>09:21:58: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049052505228185600#1:build-log.txt%3A5334)

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

#### internal (2x / 33.33%)

<details>
<summary> kind cluster creation failure (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-04-28 07:39:01 &#43;0000 UTC_: <code>07:53:14: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049030373865689088#1:build-log.txt%3A1092)
<details>
<summary>all...</summary>

* _2026-04-28 07:39:01 &#43;0000 UTC_: <code>07:53:14: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049030373865689088#1:build-log.txt%3A1092)

</details>

<hr/>
</details>
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

#### needs-investigation (2x / 33.33%)

<details>
<summary> no error snippets (2x / 33.33%) </summary>

<hr/>

**1x**: _2026-04-28 08:05:45 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17599/pull-kubevirt-e2e-kind-1.34-sev/2049037224124616704)
<details>
<summary>all...</summary>

* _2026-04-28 08:05:45 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17599/pull-kubevirt-e2e-kind-1.34-sev/2049037224124616704)

</details>

<hr/>

**1x**: _2026-04-28 07:57:44 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-kind-1.34-sev/2049035197751496704)
<details>
<summary>all...</summary>

* _2026-04-28 07:57:44 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-kind-1.34-sev/2049035197751496704)

</details>

<hr/>
</details>

### 2026-04-27 (1x / 7.14%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)
<details>
<summary>all...</summary>

* _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)

</details>

<hr/>
</details>

### 2026-04-26 (1x / 7.14%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-26 18:34:25 &#43;0000 UTC_: <code>18:35:51: ERROR: Analysis of target &#39;//cmd/virt-controller:virt-controller-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2048470635738304512#1:build-log.txt%3A560)
<details>
<summary>all...</summary>

* _2026-04-26 18:34:25 &#43;0000 UTC_: <code>18:35:51: ERROR: Analysis of target &#39;//cmd/virt-controller:virt-controller-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2048470635738304512#1:build-log.txt%3A560)

</details>

<hr/>
</details>

### 2026-04-24 (1x / 7.14%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-24 08:18:40 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17591/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2047590808902701056#1:build-log.txt%3A1605)
<details>
<summary>all...</summary>

* _2026-04-24 08:18:40 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17591/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2047590808902701056#1:build-log.txt%3A1605)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### internal (3x / 21.43%)

<details>
<summary> make cluster lifecycle target failure (2x / 14.29%) </summary>

<hr/>

**2x**: _2026-04-24 08:18:40 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17591/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2047590808902701056#1:build-log.txt%3A1605)
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


* _2026-04-24 08:18:40 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17591/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2047590808902701056#1:build-log.txt%3A1605)
<details><summary>context</summary>
<pre>08:34:54: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
08:34:54: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
08:35:43: &#43; rm -f /tmp/kubevirt.deploy.NwzY
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> kind cluster creation failure (1x / 7.14%) </summary>

<hr/>

**1x**: _2026-04-28 07:39:01 &#43;0000 UTC_: <code>07:53:14: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049030373865689088#1:build-log.txt%3A1092)
<details>
<summary>all...</summary>

* _2026-04-28 07:39:01 &#43;0000 UTC_: <code>07:53:14: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049030373865689088#1:build-log.txt%3A1092)
<details><summary>context</summary>
<pre>07:53:12:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
07:53:12:  • Preparing nodes 📦 📦   ...
07:53:14:  ✗ Preparing nodes 📦 📦
07:53:14: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
07:53:14:
07:53:14: Stack Trace:
07:53:14: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


</details>

<hr/>
</details>

### external (9x / 64.29%)

<details>
<summary> download failure in context (8x / 57.14%) </summary>

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

**4x**: _2026-04-28 09:06:28 &#43;0000 UTC_: <code>09:21:58: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049052505228185600#1:build-log.txt%3A5334)
<details>
<summary>all...</summary>

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


* _2026-04-28 09:06:28 &#43;0000 UTC_: <code>09:21:58: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049052505228185600#1:build-log.txt%3A5334)
<details><summary>context</summary>
<pre>09:21:58: ERROR: Analysis of target &#39;//:buildifier&#39; failed; build aborted:
09:21:58: INFO: Elapsed time: 0.323s
09:21:58: INFO: 0 processes.
09:21:58: ERROR: Build failed. Not running target
make: *** [Makefile:28: bazel-build-functests] Error 1
&#43; ret=2
&#43; check_for_panics</pre>
</details>


</details>

<hr/>

**1x**: _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)
<details>
<summary>all...</summary>

* _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)
<details><summary>context</summary>
<pre>07:21:17: Repository rule oci_alias defined at:
07:21:17:   /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/rules_oci/oci/private/pull.bzl:417:28: in &lt;toplevel&gt;
07:21:17: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-operator/BUILD.bazel:47:10: //cmd/virt-operator:virt-operator-image depends on @zstd_linux_amd64//:zstd_toolchain in repository @zstd_linux_amd64 which failed to fetch. no such package &#39;@zstd_linux_amd64//&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/zstd-prebuilt/releases/download/v1.5.6/zstd_linux_amd64] to /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/zstd_linux_amd64/zstd: GET returned 502 Bad Gateway or Proxy Error
07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed
07:21:17: INFO: Elapsed time: 0.253s
07:21:17: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>

**1x**: _2026-04-26 18:34:25 &#43;0000 UTC_: <code>18:35:51: ERROR: Analysis of target &#39;//cmd/virt-controller:virt-controller-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2048470635738304512#1:build-log.txt%3A560)
<details>
<summary>all...</summary>

* _2026-04-26 18:34:25 &#43;0000 UTC_: <code>18:35:51: ERROR: Analysis of target &#39;//cmd/virt-controller:virt-controller-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2048470635738304512#1:build-log.txt%3A560)
<details><summary>context</summary>
<pre>18:35:51: Repository rule oci_alias defined at:
18:35:51:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/rules_oci/oci/private/pull.bzl:472:28: in &lt;toplevel&gt;
18:35:51: ERROR: /root/go/src/kubevirt.io/kubevirt/cmd/virt-controller/BUILD.bazel:44:10: //cmd/virt-controller:virt-controller-image depends on @oci_regctl_linux_amd64//:regctl_toolchain in repository @oci_regctl_linux_amd64 which failed to fetch. no such package &#39;@oci_regctl_linux_amd64//&#39;: java.io.IOException: Error downloading [https://github.com/regclient/regclient/releases/download/v0.8.0/regctl-linux-amd64] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/oci_regctl_linux_amd64/regctl: GET returned 502 Bad Gateway or Proxy Error
18:35:51: ERROR: Analysis of target &#39;//cmd/virt-controller:virt-controller-image&#39; failed; build aborted: Analysis failed
18:35:51: INFO: Elapsed time: 0.265s
18:35:51: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 7.14%) </summary>

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

### needs-investigation (2x / 14.29%)

<details>
<summary> no error snippets (2x / 14.29%) </summary>

<hr/>

**1x**: _2026-04-28 08:05:45 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17599/pull-kubevirt-e2e-kind-1.34-sev/2049037224124616704)
<details>
<summary>all...</summary>

* _2026-04-28 08:05:45 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17599/pull-kubevirt-e2e-kind-1.34-sev/2049037224124616704)

</details>

<hr/>

**1x**: _2026-04-28 07:57:44 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-kind-1.34-sev/2049035197751496704)
<details>
<summary>all...</summary>

* _2026-04-28 07:57:44 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-kind-1.34-sev/2049035197751496704)

</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (5x / 35.71%)


#### external (3x / 60.00%)

<details>
<summary> download failure in context (2x / 40.00%) </summary>

<hr/>

**1x**: _2026-04-29 12:01:51 &#43;0000 UTC_: <code>12:02:55: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17647/pull-kubevirt-e2e-kind-sriov/2049458981394452480#1:build-log.txt%3A307)
<details>
<summary>all...</summary>

* _2026-04-29 12:01:51 &#43;0000 UTC_: <code>12:02:55: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17647/pull-kubevirt-e2e-kind-sriov/2049458981394452480#1:build-log.txt%3A307)

</details>

<hr/>

**1x**: _2026-04-26 18:34:25 &#43;0000 UTC_: <code>18:35:51: ERROR: Analysis of target &#39;//cmd/virt-controller:virt-controller-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2048470635738304512#1:build-log.txt%3A560)
<details>
<summary>all...</summary>

* _2026-04-26 18:34:25 &#43;0000 UTC_: <code>18:35:51: ERROR: Analysis of target &#39;//cmd/virt-controller:virt-controller-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2048470635738304512#1:build-log.txt%3A560)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)
<details>
<summary>all...</summary>

* _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)

</details>

<hr/>
</details>

#### needs-investigation (2x / 40.00%)

<details>
<summary> no error snippets (2x / 40.00%) </summary>

<hr/>

**1x**: _2026-04-28 08:05:45 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17599/pull-kubevirt-e2e-kind-1.34-sev/2049037224124616704)
<details>
<summary>all...</summary>

* _2026-04-28 08:05:45 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17599/pull-kubevirt-e2e-kind-1.34-sev/2049037224124616704)

</details>

<hr/>

**1x**: _2026-04-28 07:57:44 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-kind-1.34-sev/2049035197751496704)
<details>
<summary>all...</summary>

* _2026-04-28 07:57:44 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-kind-1.34-sev/2049035197751496704)

</details>

<hr/>
</details>

### release-1.8 (8x / 57.14%)


#### internal (2x / 25.00%)

<details>
<summary> kind cluster creation failure (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-04-28 07:39:01 &#43;0000 UTC_: <code>07:53:14: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049030373865689088#1:build-log.txt%3A1092)
<details>
<summary>all...</summary>

* _2026-04-28 07:39:01 &#43;0000 UTC_: <code>07:53:14: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049030373865689088#1:build-log.txt%3A1092)

</details>

<hr/>
</details>
<details>
<summary> make cluster lifecycle target failure (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-04-28 15:26:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17631/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049148231257952256#1:build-log.txt%3A1217)
<details>
<summary>all...</summary>

* _2026-04-28 15:26:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17631/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049148231257952256#1:build-log.txt%3A1217)

</details>

<hr/>
</details>

#### external (6x / 75.00%)

<details>
<summary> download failure in context (6x / 75.00%) </summary>

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

**3x**: _2026-04-28 09:06:28 &#43;0000 UTC_: <code>09:21:58: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049052505228185600#1:build-log.txt%3A5334)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:33:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2049451505487974400#1:build-log.txt%3A307)

* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:53:16: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049451503768309760#1:build-log.txt%3A5396)

* _2026-04-28 09:06:28 &#43;0000 UTC_: <code>09:21:58: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049052505228185600#1:build-log.txt%3A5334)

</details>

<hr/>

**1x**: _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)
<details>
<summary>all...</summary>

* _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)

</details>

<hr/>
</details>

### release-1.7 (1x / 7.14%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-04-24 08:18:40 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17591/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2047590808902701056#1:build-log.txt%3A1605)
<details>
<summary>all...</summary>

* _2026-04-24 08:18:40 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17591/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2047590808902701056#1:build-log.txt%3A1605)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (3x / 21.43%)


#### external (3x / 100.00%)

<details>
<summary> download failure in context (3x / 100.00%) </summary>

<hr/>

**2x**: _2026-04-29 12:01:51 &#43;0000 UTC_: <code>12:02:55: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17647/pull-kubevirt-e2e-kind-sriov/2049458981394452480#1:build-log.txt%3A307)
<details>
<summary>all...</summary>

* _2026-04-29 12:01:51 &#43;0000 UTC_: <code>12:02:55: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17647/pull-kubevirt-e2e-kind-sriov/2049458981394452480#1:build-log.txt%3A307)

* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:33:02: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.33-sig-network-1.8/2049451505487974400#1:build-log.txt%3A307)

</details>

<hr/>

**1x**: _2026-04-29 11:32:03 &#43;0000 UTC_: <code>11:34:28: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-kind-sriov-1.8/2049451494389846016#1:build-log.txt%3A529)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:03 &#43;0000 UTC_: <code>11:34:28: ERROR: Analysis of target &#39;//containerimages:alpine-with-test-tooling&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-kind-sriov-1.8/2049451494389846016#1:build-log.txt%3A529)

</details>

<hr/>
</details>

### sig-compute (11x / 78.57%)


#### internal (3x / 27.27%)

<details>
<summary> make cluster lifecycle target failure (2x / 18.18%) </summary>

<hr/>

**2x**: _2026-04-24 08:18:40 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17591/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2047590808902701056#1:build-log.txt%3A1605)
<details>
<summary>all...</summary>

* _2026-04-28 15:26:50 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 2</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17631/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049148231257952256#1:build-log.txt%3A1217)

* _2026-04-24 08:18:40 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17591/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2047590808902701056#1:build-log.txt%3A1605)

</details>

<hr/>
</details>
<details>
<summary> kind cluster creation failure (1x / 9.09%) </summary>

<hr/>

**1x**: _2026-04-28 07:39:01 &#43;0000 UTC_: <code>07:53:14: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049030373865689088#1:build-log.txt%3A1092)
<details>
<summary>all...</summary>

* _2026-04-28 07:39:01 &#43;0000 UTC_: <code>07:53:14: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049030373865689088#1:build-log.txt%3A1092)

</details>

<hr/>
</details>

#### external (6x / 54.55%)

<details>
<summary> download failure in context (5x / 45.45%) </summary>

<hr/>

**1x**: _2026-04-29 11:32:06 &#43;0000 UTC_: <code>11:34:30: ERROR: Analysis of target &#39;//cmd/cniplugins/passt-binding/cmd:network-passt-binding-cni-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2049451511481634816#1:build-log.txt%3A564)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:06 &#43;0000 UTC_: <code>11:34:30: ERROR: Analysis of target &#39;//cmd/cniplugins/passt-binding/cmd:network-passt-binding-cni-image&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-operator-1.8/2049451511481634816#1:build-log.txt%3A564)

</details>

<hr/>

**2x**: _2026-04-28 09:06:28 &#43;0000 UTC_: <code>09:21:58: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049052505228185600#1:build-log.txt%3A5334)
<details>
<summary>all...</summary>

* _2026-04-29 11:32:02 &#43;0000 UTC_: <code>11:53:16: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17646/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2049451503768309760#1:build-log.txt%3A5396)

* _2026-04-28 09:06:28 &#43;0000 UTC_: <code>09:21:58: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17590/pull-kubevirt-e2e-kind-1.34-sev-1.8/2049052505228185600#1:build-log.txt%3A5334)

</details>

<hr/>

**1x**: _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)
<details>
<summary>all...</summary>

* _2026-04-27 07:18:06 &#43;0000 UTC_: <code>07:21:17: ERROR: Analysis of target &#39;//cmd/virt-operator:virt-operator-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17612/pull-kubevirt-e2e-k8s-1.34-sig-compute-1.8/2048662822543429632#1:build-log.txt%3A498)

</details>

<hr/>

**1x**: _2026-04-26 18:34:25 &#43;0000 UTC_: <code>18:35:51: ERROR: Analysis of target &#39;//cmd/virt-controller:virt-controller-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2048470635738304512#1:build-log.txt%3A560)
<details>
<summary>all...</summary>

* _2026-04-26 18:34:25 &#43;0000 UTC_: <code>18:35:51: ERROR: Analysis of target &#39;//cmd/virt-controller:virt-controller-image&#39; failed; build aborted: Analysis failed</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-k8s-1.35-sig-compute-migrations/2048470635738304512#1:build-log.txt%3A560)

</details>

<hr/>
</details>
<details>
<summary> download failure from external URL (1x / 9.09%) </summary>

<hr/>

**1x**: _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)
<details>
<summary>all...</summary>

* _2026-04-28 09:55:02 &#43;0000 UTC_: <code>09:55:54: ERROR: Error computing the main repository mapping: no such package &#39;@aspect_bazel_lib//lib&#39;: java.io.IOException: Error downloading [https://github.com/aspect-build/bazel-lib/releases/download/v2.7.2/bazel-lib-v2.7.2.tar.gz] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/aspect_bazel_lib/temp7199907394090039910/bazel-lib-v2.7.2.tar.gz: GET returned 502 Bad Gateway or Proxy Error</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17481/pull-kubevirt-e2e-kind-1.34-sev/2049064726003978240#1:build-log.txt%3A289)

</details>

<hr/>
</details>

#### needs-investigation (2x / 18.18%)

<details>
<summary> no error snippets (2x / 18.18%) </summary>

<hr/>

**1x**: _2026-04-28 08:05:45 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17599/pull-kubevirt-e2e-kind-1.34-sev/2049037224124616704)
<details>
<summary>all...</summary>

* _2026-04-28 08:05:45 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17599/pull-kubevirt-e2e-kind-1.34-sev/2049037224124616704)

</details>

<hr/>

**1x**: _2026-04-28 07:57:44 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-kind-1.34-sev/2049035197751496704)
<details>
<summary>all...</summary>

* _2026-04-28 07:57:44 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17603/pull-kubevirt-e2e-kind-1.34-sev/2049035197751496704)

</details>

<hr/>
</details>

Last updated: 2026-05-01 04:32:42
