



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-07-13 (3x / 75.00%)


#### external (2x / 66.67%)

<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)
<details>
<summary>all...</summary>

* _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)
<details>
<summary>all...</summary>

* _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)

</details>

<hr/>
</details>

#### internal (1x / 33.33%)

<details>
<summary> KubeVirt deployment timeout (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)
<details>
<summary>all...</summary>

* _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)

</details>

<hr/>
</details>

### 2026-07-09 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-09 08:41:42 &#43;0000 UTC_: <code>08:43:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18394/pull-kubevirt-e2e-kind-1.34-sev/2075138204012384256#1:build-log.txt%3A547)
<details>
<summary>all...</summary>

* _2026-07-09 08:41:42 &#43;0000 UTC_: <code>08:43:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18394/pull-kubevirt-e2e-kind-1.34-sev/2075138204012384256#1:build-log.txt%3A547)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (3x / 75.00%)

<details>
<summary> download failure in context (2x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)
<details>
<summary>all...</summary>

* _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)
<details><summary>context</summary>
<pre>11:07:18: ERROR: Analysis of target &#39;//:push-virt-template-apiserver&#39; failed; build aborted:
11:07:18: INFO: Elapsed time: 2.099s
11:07:18: INFO: 0 processes.
11:07:18: ERROR: Build failed. Not running target
11:07:21: &#43; rm -f /tmp/kubevirt.deploy.thAt
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>

**1x**: _2026-07-09 08:41:42 &#43;0000 UTC_: <code>08:43:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18394/pull-kubevirt-e2e-kind-1.34-sev/2075138204012384256#1:build-log.txt%3A547)
<details>
<summary>all...</summary>

* _2026-07-09 08:41:42 &#43;0000 UTC_: <code>08:43:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18394/pull-kubevirt-e2e-kind-1.34-sev/2075138204012384256#1:build-log.txt%3A547)
<details><summary>context</summary>
<pre>08:43:35: Repository rule http_file defined at:
08:43:35:   /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:466:28: in &lt;toplevel&gt;
08:43:35: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:6d35dee4ad3f32402ae3e809aea0e74bf2d94d0a30fb633cb83a025d3e2dbaf3] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/6d35dee4ad3f32402ae3e809aea0e74bf2d94d0a30fb633cb83a025d3e2dbaf3: Bytes read 8388608 but wanted 22553707
08:43:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:43:35: INFO: Elapsed time: 3.127s
08:43:35: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)
<details>
<summary>all...</summary>

* _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)
<details><summary>context</summary>
<pre>14:06:04: &#43; sleep 1m
14:07:04: &#43; _kubectl wait -n kubevirt kv kubevirt --for condition=Available --timeout 5m
14:07:04: &#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.34/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.34/.kubeconfig wait -n kubevirt kv kubevirt --for condition=Available --timeout 5m
14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt
14:12:04: &#43; (( count&#43;&#43; ))
14:12:04: &#43; (( count == 5 ))
14:12:04: &#43; echo &#39;KubeVirt not ready in time&#39;</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 25.00%)

<details>
<summary> KubeVirt deployment timeout (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)
<details>
<summary>all...</summary>

* _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)
<details><summary>context</summary>
<pre>14:05:20: &#43; sleep 1m
14:06:20: &#43; _kubectl wait -n kubevirt kv kubevirt --for condition=Available --timeout 5m
14:06:20: &#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.34/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.34/.kubeconfig wait -n kubevirt kv kubevirt --for condition=Available --timeout 5m
14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt
14:11:20: &#43; (( count&#43;&#43; ))
14:11:20: &#43; (( count == 5 ))
14:11:20: &#43; echo &#39;KubeVirt not ready in time&#39;</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### release-1.9 (2x / 50.00%)


#### internal (1x / 50.00%)

<details>
<summary> KubeVirt deployment timeout (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)
<details>
<summary>all...</summary>

* _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)
<details>
<summary>all...</summary>

* _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)

</details>

<hr/>
</details>

### main (2x / 50.00%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (2x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)
<details>
<summary>all...</summary>

* _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)

</details>

<hr/>

**1x**: _2026-07-09 08:41:42 &#43;0000 UTC_: <code>08:43:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18394/pull-kubevirt-e2e-kind-1.34-sev/2075138204012384256#1:build-log.txt%3A547)
<details>
<summary>all...</summary>

* _2026-07-09 08:41:42 &#43;0000 UTC_: <code>08:43:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18394/pull-kubevirt-e2e-kind-1.34-sev/2075138204012384256#1:build-log.txt%3A547)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (4x / 100.00%)


#### external (3x / 75.00%)

<details>
<summary> download failure in context (2x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)
<details>
<summary>all...</summary>

* _2026-07-13 10:57:10 &#43;0000 UTC_: <code>11:07:18: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18351/pull-kubevirt-e2e-k8s-1.35-sig-compute/2076613329182265344#1:build-log.txt%3A4201)

</details>

<hr/>

**1x**: _2026-07-09 08:41:42 &#43;0000 UTC_: <code>08:43:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18394/pull-kubevirt-e2e-kind-1.34-sev/2075138204012384256#1:build-log.txt%3A547)
<details>
<summary>all...</summary>

* _2026-07-09 08:41:42 &#43;0000 UTC_: <code>08:43:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18394/pull-kubevirt-e2e-kind-1.34-sev/2075138204012384256#1:build-log.txt%3A547)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)
<details>
<summary>all...</summary>

* _2026-07-13 13:28:41 &#43;0000 UTC_: <code>14:12:04: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18445/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076659968248909824#1:build-log.txt%3A5189)

</details>

<hr/>
</details>

#### internal (1x / 25.00%)

<details>
<summary> KubeVirt deployment timeout (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)
<details>
<summary>all...</summary>

* _2026-07-13 13:29:22 &#43;0000 UTC_: <code>14:11:20: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18446/pull-kubevirt-e2e-kind-1.34-sev-1.9/2076660067666497536#1:build-log.txt%3A5137)

</details>

<hr/>
</details>

Last updated: 2026-07-15 06:50:34
