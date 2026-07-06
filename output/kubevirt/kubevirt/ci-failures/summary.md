



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-07-02 (2x / 7.69%)


#### needs-investigation (1x / 50.00%)

<details>
<summary> no matching pattern (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details>
<summary>all...</summary>

* _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)

</details>

<hr/>
</details>

#### internal (1x / 50.00%)

<details>
<summary> kind cluster creation failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)
<details>
<summary>all...</summary>

* _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)

</details>

<hr/>
</details>

### 2026-07-01 (24x / 92.31%)


#### external (23x / 95.83%)

<details>
<summary> download failure in context (23x / 95.83%) </summary>

<hr/>

**1x**: _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details>
<summary>all...</summary>

* _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)

</details>

<hr/>

**22x**: _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)
<details>
<summary>all...</summary>

* _2026-07-01 23:00:19 &#43;0000 UTC_: <code>23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072414264844357632#1:build-log.txt%3A557)

* _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)

* _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)

* _2026-07-01 15:41:18 &#43;0000 UTC_: <code>15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072291187003232256#1:build-log.txt%3A591)

* _2026-07-01 15:41:06 &#43;0000 UTC_: <code>15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072291186906763264#1:build-log.txt%3A569)

* _2026-07-01 13:49:45 &#43;0000 UTC_: <code>13:51:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-k8s-1.36-sig-operator/2072293741271453696#1:build-log.txt%3A562)

* _2026-07-01 12:34:52 &#43;0000 UTC_: <code>12:37:06: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-windows2016/2072297118487285760#1:build-log.txt%3A613)

* _2026-07-01 12:18:33 &#43;0000 UTC_: <code>12:20:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072292085007912960#1:build-log.txt%3A573)

* _2026-07-01 12:14:12 &#43;0000 UTC_: <code>12:15:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072292084718505984#1:build-log.txt%3A620)

* _2026-07-01 11:33:46 &#43;0000 UTC_: <code>11:36:21: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.35-sig-operator/2072238776180019200#1:build-log.txt%3A665)

* _2026-07-01 11:30:18 &#43;0000 UTC_: <code>11:31:57: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.34-sig-network/2072238774993031168#1:build-log.txt%3A671)

* _2026-07-01 11:29:20 &#43;0000 UTC_: <code>11:31:11: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072238774904950784#1:build-log.txt%3A651)

* _2026-07-01 09:28:19 &#43;0000 UTC_: <code>09:29:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072238855989235712#1:build-log.txt%3A573)

* _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)

* _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)

* _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)

* _2026-07-01 08:42:43 &#43;0000 UTC_: <code>08:44:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-compute/2072235893535543296#1:build-log.txt%3A577)

* _2026-07-01 08:39:41 &#43;0000 UTC_: <code>08:41:54: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072235893455851520#1:build-log.txt%3A545)

* _2026-07-01 08:39:26 &#43;0000 UTC_: <code>08:41:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072235893300662272#1:build-log.txt%3A576)

* _2026-07-01 08:36:56 &#43;0000 UTC_: <code>08:38:52: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072235893233553408#1:build-log.txt%3A592)

* _2026-07-01 08:31:01 &#43;0000 UTC_: <code>08:32:48: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-windows2016/2072235893103529984#1:build-log.txt%3A578)

* _2026-07-01 07:11:39 &#43;0000 UTC_: <code>07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18275/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2072215977843494912#1:build-log.txt%3A696)

</details>

<hr/>
</details>

#### internal (1x / 4.17%)

<details>
<summary> KubeVirt deployment timeout (1x / 4.17%) </summary>

<hr/>

**1x**: _2026-07-01 11:25:03 &#43;0000 UTC_: <code>12:06:57: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-kind-1.34-sev/2072280205950455808#1:build-log.txt%3A5129)
<details>
<summary>all...</summary>

* _2026-07-01 11:25:03 &#43;0000 UTC_: <code>12:06:57: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-kind-1.34-sev/2072280205950455808#1:build-log.txt%3A5129)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### internal (2x / 7.69%)

<details>
<summary> kind cluster creation failure (1x / 3.85%) </summary>

<hr/>

**1x**: _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)
<details>
<summary>all...</summary>

* _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)
<details><summary>context</summary>
<pre>13:14:10:  ✓ Ensuring node image (kindest/node:v1.34.0) 🖼
13:14:10:  • Preparing nodes 📦 📦   ...
13:14:13:  ✗ Preparing nodes 📦 📦
13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
13:14:13:
13:14:13: Stack Trace:
13:14:13: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> KubeVirt deployment timeout (1x / 3.85%) </summary>

<hr/>

**1x**: _2026-07-01 11:25:03 &#43;0000 UTC_: <code>12:06:57: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-kind-1.34-sev/2072280205950455808#1:build-log.txt%3A5129)
<details>
<summary>all...</summary>

* _2026-07-01 11:25:03 &#43;0000 UTC_: <code>12:06:57: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-kind-1.34-sev/2072280205950455808#1:build-log.txt%3A5129)
<details><summary>context</summary>
<pre>12:00:57: &#43; sleep 1m
12:01:57: &#43; _kubectl wait -n kubevirt kv kubevirt --for condition=Available --timeout 5m
12:01:57: &#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.34/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.34/.kubeconfig wait -n kubevirt kv kubevirt --for condition=Available --timeout 5m
12:06:57: error: timed out waiting for the condition on kubevirts/kubevirt
12:06:57: &#43; (( count&#43;&#43; ))
12:06:57: &#43; (( count == 5 ))
12:06:57: &#43; echo &#39;KubeVirt not ready in time&#39;</pre>
</details>


</details>

<hr/>
</details>

### external (23x / 88.46%)

<details>
<summary> download failure in context (23x / 88.46%) </summary>

<hr/>

**1x**: _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details>
<summary>all...</summary>

* _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details><summary>context</summary>
<pre>13:39:54: ERROR: Analysis of target &#39;//:push-alpine-ext-kernel-boot-demo&#39; failed; build aborted:
13:39:54: INFO: Elapsed time: 1.124s
13:39:54: INFO: 0 processes.
13:39:54: ERROR: Build failed. Not running target
13:39:58: &#43; rm -f /tmp/kubevirt.deploy.azEg
make: *** [Makefile:189: cluster-sync] Error 1
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>

**22x**: _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)
<details>
<summary>all...</summary>

* _2026-07-01 23:00:19 &#43;0000 UTC_: <code>23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072414264844357632#1:build-log.txt%3A557)
<details><summary>context</summary>
<pre>23:02:17:  * Hash &#39;7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212
23:02:17: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
23:02:17: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
23:02:17: INFO: Elapsed time: 0.829s
23:02:17: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)
<details><summary>context</summary>
<pre>16:18:01:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
16:18:01: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
16:18:01: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
16:18:01: INFO: Elapsed time: 0.933s
16:18:01: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)
<details><summary>context</summary>
<pre>15:45:00:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
15:45:00: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
15:45:00: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
15:45:00: INFO: Elapsed time: 1.118s
15:45:00: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 15:41:18 &#43;0000 UTC_: <code>15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072291187003232256#1:build-log.txt%3A591)
<details><summary>context</summary>
<pre>15:43:41:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
15:43:41: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
15:43:41: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
15:43:41: INFO: Elapsed time: 0.808s
15:43:41: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 15:41:06 &#43;0000 UTC_: <code>15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072291186906763264#1:build-log.txt%3A569)
<details><summary>context</summary>
<pre>15:42:53:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
15:42:53: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
15:42:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
15:42:53: INFO: Elapsed time: 0.958s
15:42:53: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 13:49:45 &#43;0000 UTC_: <code>13:51:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-k8s-1.36-sig-operator/2072293741271453696#1:build-log.txt%3A562)
<details><summary>context</summary>
<pre>13:51:35:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
13:51:35: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
13:51:35: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
13:51:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
13:51:35: INFO: Elapsed time: 1.937s
13:51:35: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 12:34:52 &#43;0000 UTC_: <code>12:37:06: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-windows2016/2072297118487285760#1:build-log.txt%3A613)
<details><summary>context</summary>
<pre>12:37:06:  * Hash &#39;7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212
12:37:06: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
12:37:06: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
12:37:06: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
12:37:06: INFO: Elapsed time: 1.173s
12:37:06: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 12:18:33 &#43;0000 UTC_: <code>12:20:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072292085007912960#1:build-log.txt%3A573)
<details><summary>context</summary>
<pre>12:20:20:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
12:20:20: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
12:20:20: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
12:20:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
12:20:20: INFO: Elapsed time: 0.602s
12:20:20: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 12:14:12 &#43;0000 UTC_: <code>12:15:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072292084718505984#1:build-log.txt%3A620)
<details><summary>context</summary>
<pre>12:15:55:  * Hash &#39;ca3f629c350fd1cb653caee7c5aed8967be4e43d9577fbb1cb4ed66c260e224e&#39; for https://gcr.io/v2/distroless/base-debian12/manifests/sha256:ca3f629c350fd1cb653caee7c5aed8967be4e43d9577fbb1cb4ed66c260e224e
12:15:55: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
12:15:55: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
12:15:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
12:15:55: INFO: Elapsed time: 0.881s
12:15:55: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 11:33:46 &#43;0000 UTC_: <code>11:36:21: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.35-sig-operator/2072238776180019200#1:build-log.txt%3A665)
<details><summary>context</summary>
<pre>11:36:21:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
11:36:21: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
11:36:21: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
11:36:21: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
11:36:21: INFO: Elapsed time: 1.137s
11:36:21: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 11:30:18 &#43;0000 UTC_: <code>11:31:57: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.34-sig-network/2072238774993031168#1:build-log.txt%3A671)
<details><summary>context</summary>
<pre>11:31:57:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
11:31:57: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
11:31:57: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
11:31:57: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
11:31:57: INFO: Elapsed time: 1.086s
11:31:57: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 11:29:20 &#43;0000 UTC_: <code>11:31:11: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072238774904950784#1:build-log.txt%3A651)
<details><summary>context</summary>
<pre>11:31:11:  * Hash &#39;7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212
11:31:11: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
11:31:11: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
11:31:11: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
11:31:11: INFO: Elapsed time: 2.058s
11:31:11: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 09:28:19 &#43;0000 UTC_: <code>09:29:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072238855989235712#1:build-log.txt%3A573)
<details><summary>context</summary>
<pre>09:29:53:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
09:29:53: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
09:29:53: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
09:29:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
09:29:53: INFO: Elapsed time: 0.834s
09:29:53: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)
<details><summary>context</summary>
<pre>09:24:45:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
09:24:45: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
09:24:45: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
09:24:45: INFO: Elapsed time: 0.768s
09:24:45: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)
<details><summary>context</summary>
<pre>08:58:05:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:58:05: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:58:05: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:58:05: INFO: Elapsed time: 0.863s
08:58:05: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)
<details><summary>context</summary>
<pre>08:53:13:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:53:13: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:53:13: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:53:13: INFO: Elapsed time: 0.707s
08:53:13: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:42:43 &#43;0000 UTC_: <code>08:44:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-compute/2072235893535543296#1:build-log.txt%3A577)
<details><summary>context</summary>
<pre>08:44:20:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:44:20: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:44:20: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:44:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:44:20: INFO: Elapsed time: 0.923s
08:44:20: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:39:41 &#43;0000 UTC_: <code>08:41:54: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072235893455851520#1:build-log.txt%3A545)
<details><summary>context</summary>
<pre>08:41:54:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:41:54: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:41:54: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:41:54: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:41:54: INFO: Elapsed time: 0.599s
08:41:54: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:39:26 &#43;0000 UTC_: <code>08:41:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072235893300662272#1:build-log.txt%3A576)
<details><summary>context</summary>
<pre>08:41:00:  * Hash &#39;a53fd982787799c2d8cfaa37a2b6fbac4f416437768a25d2eb246dff46bb9d79&#39; for https://quay.io/v2/kubevirtci/fedora-with-test-tooling/manifests/sha256:a53fd982787799c2d8cfaa37a2b6fbac4f416437768a25d2eb246dff46bb9d79
08:41:00: If the definition of &#39;repository @fedora_with_test_tooling_single&#39; was updated, verify that the hashes were also updated.
08:41:00: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:41:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:41:00: INFO: Elapsed time: 0.930s
08:41:00: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:36:56 &#43;0000 UTC_: <code>08:38:52: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072235893233553408#1:build-log.txt%3A592)
<details><summary>context</summary>
<pre>08:38:52:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:38:52: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:38:52: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:38:52: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:38:52: INFO: Elapsed time: 0.752s
08:38:52: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 08:31:01 &#43;0000 UTC_: <code>08:32:48: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-windows2016/2072235893103529984#1:build-log.txt%3A578)
<details><summary>context</summary>
<pre>08:32:48:  * Hash &#39;990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e&#39; for https://gcr.io/v2/distroless/base-debian12/blobs/sha256:990a9c434e5e0f11549a8d4a41a1991e621b04e30cd63269adbc97b1dc38fd7e
08:32:48: If the definition of &#39;repository @go_image_base_single&#39; was updated, verify that the hashes were also updated.
08:32:48: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
08:32:48: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
08:32:48: INFO: Elapsed time: 0.637s
08:32:48: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


* _2026-07-01 07:11:39 &#43;0000 UTC_: <code>07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18275/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2072215977843494912#1:build-log.txt%3A696)
<details><summary>context</summary>
<pre>07:13:55:  * Hash &#39;5aed89bce5d6bb3ece39298fe92e2900d15dcecf58ce152e7c7d9f97c6dd43fe&#39; for https://quay.io/v2/kubevirt/fedora-realtime-container-disk/blobs/sha256:5aed89bce5d6bb3ece39298fe92e2900d15dcecf58ce152e7c7d9f97c6dd43fe
07:13:55: If the definition of &#39;repository @fedora_realtime_single&#39; was updated, verify that the hashes were also updated.
07:13:55: ERROR: /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base/BUILD.bazel:7:6: @alpine-ext-kernel-boot-demo-container-base//:alpine-ext-kernel-boot-demo-container-base depends on @alpine-ext-kernel-boot-demo-container-base_single//:alpine-ext-kernel-boot-demo-container-base_single in repository @alpine-ext-kernel-boot-demo-container-base_single which failed to fetch. no such package &#39;@alpine-ext-kernel-boot-demo-container-base_single//&#39;: java.io.IOException: Error downloading [https://quay.io/v2/kubevirt/alpine-ext-kernel-boot-demo/blobs/sha256:1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e] to /tmp/cache/bazel/6f347497f91c9a385dcd9294645b76e0/external/alpine-ext-kernel-boot-demo-container-base_single/blobs/sha256/1451e69a7ebf16911ee0ac3d8dd930dbb9c78f88b1d762f10a3992cf10e6147e: Bytes read 2097153 but wanted 54212734
07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:
07:13:55: INFO: Elapsed time: 0.826s
07:13:55: INFO: 0 processes.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (1x / 3.85%)

<details>
<summary> no matching pattern (1x / 3.85%) </summary>

<hr/>

**1x**: _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details>
<summary>all...</summary>

* _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details><summary>context</summary>
<pre>time=&#34;2026-07-02T07:38:37Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f\&#34;, deleting it&#34;
time=&#34;2026-07-02T07:38:39Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f\&#34;, deleting it&#34;
time=&#34;2026-07-02T07:38:39Z&#34; level=error msg=&#34;cleaning up storage: removing container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer \&#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f\&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-183572116/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link&#34;
Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link
time=&#34;2026-07-02T07:38:39Z&#34; level=warning msg=&#34;Found incomplete layer \&#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f\&#34;, deleting it&#34;
/usr/local/bin/runner.sh: line 50: wait: pid 1220 is not a child of this shell
================================================================================</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### release-1.7 (1x / 3.85%)


#### internal (1x / 100.00%)

<details>
<summary> kind cluster creation failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)
<details>
<summary>all...</summary>

* _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)

</details>

<hr/>
</details>

### main (25x / 96.15%)


#### external (23x / 92.00%)

<details>
<summary> download failure in context (23x / 92.00%) </summary>

<hr/>

**1x**: _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details>
<summary>all...</summary>

* _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)

</details>

<hr/>

**22x**: _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)
<details>
<summary>all...</summary>

* _2026-07-01 23:00:19 &#43;0000 UTC_: <code>23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072414264844357632#1:build-log.txt%3A557)

* _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)

* _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)

* _2026-07-01 15:41:18 &#43;0000 UTC_: <code>15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072291187003232256#1:build-log.txt%3A591)

* _2026-07-01 15:41:06 &#43;0000 UTC_: <code>15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072291186906763264#1:build-log.txt%3A569)

* _2026-07-01 13:49:45 &#43;0000 UTC_: <code>13:51:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-k8s-1.36-sig-operator/2072293741271453696#1:build-log.txt%3A562)

* _2026-07-01 12:34:52 &#43;0000 UTC_: <code>12:37:06: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-windows2016/2072297118487285760#1:build-log.txt%3A613)

* _2026-07-01 12:18:33 &#43;0000 UTC_: <code>12:20:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072292085007912960#1:build-log.txt%3A573)

* _2026-07-01 12:14:12 &#43;0000 UTC_: <code>12:15:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072292084718505984#1:build-log.txt%3A620)

* _2026-07-01 11:33:46 &#43;0000 UTC_: <code>11:36:21: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.35-sig-operator/2072238776180019200#1:build-log.txt%3A665)

* _2026-07-01 11:30:18 &#43;0000 UTC_: <code>11:31:57: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.34-sig-network/2072238774993031168#1:build-log.txt%3A671)

* _2026-07-01 11:29:20 &#43;0000 UTC_: <code>11:31:11: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072238774904950784#1:build-log.txt%3A651)

* _2026-07-01 09:28:19 &#43;0000 UTC_: <code>09:29:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072238855989235712#1:build-log.txt%3A573)

* _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)

* _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)

* _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)

* _2026-07-01 08:42:43 &#43;0000 UTC_: <code>08:44:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-compute/2072235893535543296#1:build-log.txt%3A577)

* _2026-07-01 08:39:41 &#43;0000 UTC_: <code>08:41:54: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072235893455851520#1:build-log.txt%3A545)

* _2026-07-01 08:39:26 &#43;0000 UTC_: <code>08:41:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072235893300662272#1:build-log.txt%3A576)

* _2026-07-01 08:36:56 &#43;0000 UTC_: <code>08:38:52: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072235893233553408#1:build-log.txt%3A592)

* _2026-07-01 08:31:01 &#43;0000 UTC_: <code>08:32:48: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-windows2016/2072235893103529984#1:build-log.txt%3A578)

* _2026-07-01 07:11:39 &#43;0000 UTC_: <code>07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18275/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2072215977843494912#1:build-log.txt%3A696)

</details>

<hr/>
</details>

#### needs-investigation (1x / 4.00%)

<details>
<summary> no matching pattern (1x / 4.00%) </summary>

<hr/>

**1x**: _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details>
<summary>all...</summary>

* _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)

</details>

<hr/>
</details>

#### internal (1x / 4.00%)

<details>
<summary> KubeVirt deployment timeout (1x / 4.00%) </summary>

<hr/>

**1x**: _2026-07-01 11:25:03 &#43;0000 UTC_: <code>12:06:57: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-kind-1.34-sev/2072280205950455808#1:build-log.txt%3A5129)
<details>
<summary>all...</summary>

* _2026-07-01 11:25:03 &#43;0000 UTC_: <code>12:06:57: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-kind-1.34-sev/2072280205950455808#1:build-log.txt%3A5129)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (12x / 46.15%)


#### internal (2x / 16.67%)

<details>
<summary> kind cluster creation failure (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)
<details>
<summary>all...</summary>

* _2026-07-02 13:04:52 &#43;0000 UTC_: <code>13:14:13: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18323/pull-kubevirt-e2e-kind-1.34-sev-1.7/2072667534082969600#1:build-log.txt%3A635)

</details>

<hr/>
</details>
<details>
<summary> KubeVirt deployment timeout (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-07-01 11:25:03 &#43;0000 UTC_: <code>12:06:57: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-kind-1.34-sev/2072280205950455808#1:build-log.txt%3A5129)
<details>
<summary>all...</summary>

* _2026-07-01 11:25:03 &#43;0000 UTC_: <code>12:06:57: error: timed out waiting for the condition on kubevirts/kubevirt</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-kind-1.34-sev/2072280205950455808#1:build-log.txt%3A5129)

</details>

<hr/>
</details>

#### needs-investigation (1x / 8.33%)

<details>
<summary> no matching pattern (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)
<details>
<summary>all...</summary>

* _2026-07-02 07:28:38 &#43;0000 UTC_: <code>Error: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: cleaning up container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 storage: unmounting container 9281d0bb90970ba371489628fe55e9bff93169e3e7349072bd4bf5b3ace70da1 root filesystem: deleting layer &#34;45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f&#34;: failed to add to stage directory: rename /var/lib/shared-images/overlay/45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f /var/lib/containers/storage/overlay/tempdirs/temp-dir-210287877/1-45d857b8e4ea6d442ea23c8a411b09ed0d58cd82ff905895244fdcea6659261f: invalid cross-device link</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072583064571088896#1:build-log.txt%3A640)

</details>

<hr/>
</details>

#### external (9x / 75.00%)

<details>
<summary> download failure in context (9x / 75.00%) </summary>

<hr/>

**9x**: _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)
<details>
<summary>all...</summary>

* _2026-07-01 13:49:45 &#43;0000 UTC_: <code>13:51:35: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18308/pull-kubevirt-e2e-k8s-1.36-sig-operator/2072293741271453696#1:build-log.txt%3A562)

* _2026-07-01 12:34:52 &#43;0000 UTC_: <code>12:37:06: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-windows2016/2072297118487285760#1:build-log.txt%3A613)

* _2026-07-01 11:33:46 &#43;0000 UTC_: <code>11:36:21: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.35-sig-operator/2072238776180019200#1:build-log.txt%3A665)

* _2026-07-01 09:28:19 &#43;0000 UTC_: <code>09:29:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072238855989235712#1:build-log.txt%3A573)

* _2026-07-01 09:23:08 &#43;0000 UTC_: <code>09:24:45: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.35-sig-compute/2072241915692060672#1:build-log.txt%3A547)

* _2026-07-01 08:51:37 &#43;0000 UTC_: <code>08:53:13: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18298/pull-kubevirt-e2e-k8s-1.34-sig-operator/2072238855586582528#1:build-log.txt%3A567)

* _2026-07-01 08:42:43 &#43;0000 UTC_: <code>08:44:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-compute/2072235893535543296#1:build-log.txt%3A577)

* _2026-07-01 08:31:01 &#43;0000 UTC_: <code>08:32:48: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-windows2016/2072235893103529984#1:build-log.txt%3A578)

* _2026-07-01 07:11:39 &#43;0000 UTC_: <code>07:13:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18275/pull-kubevirt-e2e-k8s-1.36-sig-compute-migrations/2072215977843494912#1:build-log.txt%3A696)

</details>

<hr/>
</details>

### sig-storage (3x / 11.54%)


#### external (3x / 100.00%)

<details>
<summary> download failure in context (3x / 100.00%) </summary>

<hr/>

**3x**: _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)
<details>
<summary>all...</summary>

* _2026-07-01 16:16:17 &#43;0000 UTC_: <code>16:18:01: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18314/pull-kubevirt-e2e-k8s-1.36-sig-storage/2072353016559702016#1:build-log.txt%3A561)

* _2026-07-01 12:18:33 &#43;0000 UTC_: <code>12:20:20: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072292085007912960#1:build-log.txt%3A573)

* _2026-07-01 08:39:41 &#43;0000 UTC_: <code>08:41:54: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.34-sig-storage/2072235893455851520#1:build-log.txt%3A545)

</details>

<hr/>
</details>

### sig-network (10x / 38.46%)


#### external (10x / 100.00%)

<details>
<summary> download failure in context (10x / 100.00%) </summary>

<hr/>

**10x**: _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)
<details>
<summary>all...</summary>

* _2026-07-01 23:00:19 &#43;0000 UTC_: <code>23:02:17: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072414264844357632#1:build-log.txt%3A557)

* _2026-07-01 15:43:17 &#43;0000 UTC_: <code>15:45:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.34-sig-network/2072291187099701248#1:build-log.txt%3A588)

* _2026-07-01 15:41:18 &#43;0000 UTC_: <code>15:43:41: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072291187003232256#1:build-log.txt%3A591)

* _2026-07-01 15:41:06 &#43;0000 UTC_: <code>15:42:53: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18038/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072291186906763264#1:build-log.txt%3A569)

* _2026-07-01 12:14:12 &#43;0000 UTC_: <code>12:15:55: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18256/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072292084718505984#1:build-log.txt%3A620)

* _2026-07-01 11:30:18 &#43;0000 UTC_: <code>11:31:57: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.34-sig-network/2072238774993031168#1:build-log.txt%3A671)

* _2026-07-01 11:29:20 &#43;0000 UTC_: <code>11:31:11: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18240/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072238774904950784#1:build-log.txt%3A651)

* _2026-07-01 08:56:09 &#43;0000 UTC_: <code>08:58:05: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18299/pull-kubevirt-e2e-k8s-1.34-sig-network/2072241915125829632#1:build-log.txt%3A577)

* _2026-07-01 08:39:26 &#43;0000 UTC_: <code>08:41:00: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2072235893300662272#1:build-log.txt%3A576)

* _2026-07-01 08:36:56 &#43;0000 UTC_: <code>08:38:52: ERROR: Analysis of target &#39;//containerimages:alpine-ext-kernel-boot-demo-container&#39; failed; build aborted:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18288/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2072235893233553408#1:build-log.txt%3A592)

</details>

<hr/>
</details>

### sig-performance (1x / 3.85%)


#### external (1x / 100.00%)

<details>
<summary> download failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)
<details>
<summary>all...</summary>

* _2026-07-01 13:22:49 &#43;0000 UTC_: <code>13:39:54: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17955/pull-kubevirt-e2e-k8s-1.34-sig-performance/2072309073948184576#1:build-log.txt%3A2181)

</details>

<hr/>
</details>

Last updated: 2026-07-06 18:44:48
