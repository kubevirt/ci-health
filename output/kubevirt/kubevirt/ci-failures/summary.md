



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-05-22 (3x / 33.33%)


#### external (3x / 100.00%)

<details>
<summary> API rate limiter timeout (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)
<details>
<summary>all...</summary>

* _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)

</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)

</details>

<hr/>
</details>

### 2026-05-21 (3x / 33.33%)


#### internal (1x / 33.33%)

<details>
<summary> control plane startup failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)
<details>
<summary>all...</summary>

* _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)

</details>

<hr/>
</details>

#### needs-investigation (1x / 33.33%)

<details>
<summary> no matching pattern (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)
<details>
<summary>all...</summary>

* _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)

</details>

<hr/>
</details>

#### external (1x / 33.33%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)
<details>
<summary>all...</summary>

* _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)

</details>

<hr/>
</details>

### 2026-05-19 (3x / 33.33%)


#### external (3x / 100.00%)

<details>
<summary> bazel remote cache IO failure (from secondary snippet) (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-05-19 15:07:21 &#43;0000 UTC_: <code>15:33:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17833/pull-kubevirt-e2e-k8s-1.34-sig-compute/2056706870336294912#1:build-log.txt%3A4672)
<details>
<summary>all...</summary>

* _2026-05-19 15:07:21 &#43;0000 UTC_: <code>15:33:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17833/pull-kubevirt-e2e-k8s-1.34-sig-compute/2056706870336294912#1:build-log.txt%3A4672)

* _2026-05-19 10:30:44 &#43;0000 UTC_: <code>10:45:00: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16759/pull-kubevirt-e2e-k8s-1.35-sig-operator/2056649475119648768#1:build-log.txt%3A4707)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache blob fetch failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-19 15:39:19 &#43;0000 UTC_: <code>15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11 Validating nogo output for //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16500/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2056690058336931840#1:build-log.txt%3A414)
<details>
<summary>all...</summary>

* _2026-05-19 15:39:19 &#43;0000 UTC_: <code>15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11 Validating nogo output for //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16500/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2056690058336931840#1:build-log.txt%3A414)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (7x / 77.78%)

<details>
<summary> bazel remote cache IO failure (from secondary snippet) (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-05-19 15:07:21 &#43;0000 UTC_: <code>15:33:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17833/pull-kubevirt-e2e-k8s-1.34-sig-compute/2056706870336294912#1:build-log.txt%3A4672)
<details>
<summary>all...</summary>

* _2026-05-19 15:07:21 &#43;0000 UTC_: <code>15:33:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17833/pull-kubevirt-e2e-k8s-1.34-sig-compute/2056706870336294912#1:build-log.txt%3A4672)
<details><summary>context</summary>
<pre>15:33:10: java.net.UnknownHostException: bazel-cache.kubevirt-prow
15:33:10: INFO: Elapsed time: 0.726s, Critical Path: 0.25s
15:33:10: INFO: 10 processes: 6 internal, 4 processwrapper-sandbox.
15:33:10: ERROR: Build failed. Not running target
make: *** [Makefile:28: bazel-build-functests] Error 1
&#43; ret=2
&#43; check_for_panics</pre>
</details>


* _2026-05-19 10:30:44 &#43;0000 UTC_: <code>10:45:00: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16759/pull-kubevirt-e2e-k8s-1.35-sig-operator/2056649475119648768#1:build-log.txt%3A4707)
<details><summary>context</summary>
<pre>10:45:00: io.netty.channel.AbstractChannel$AnnotatedConnectException: Connection refused: bazel-cache.kubevirt-prow/172.30.239.131:8080
10:45:00: INFO: Elapsed time: 0.395s, Critical Path: 0.07s
10:45:00: INFO: 15 processes: 15 internal.
10:45:00: ERROR: Build failed. Not running target
make: *** [Makefile:28: bazel-build-functests] Error 1
&#43; ret=2
&#43; check_for_panics</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> bazel remote cache blob fetch failure (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-05-19 15:39:19 &#43;0000 UTC_: <code>15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11 Validating nogo output for //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16500/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2056690058336931840#1:build-log.txt%3A414)
<details>
<summary>all...</summary>

* _2026-05-19 15:39:19 &#43;0000 UTC_: <code>15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11 Validating nogo output for //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16500/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2056690058336931840#1:build-log.txt%3A414)
<details><summary>context</summary>
<pre>15:53:35: INFO: Analyzed 30 targets (0 packages loaded, 0 targets configured).
15:53:35: INFO: Found 30 targets...
15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11: Running nogo on //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x
15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11 Validating nogo output for //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x
15:53:38: INFO: Elapsed time: 3.489s, Critical Path: 2.92s
15:53:38: INFO: 7 processes: 1 remote cache hit, 5 internal, 1 processwrapper-sandbox.
make: *** [Makefile:39: bazel-build-images] Error 1</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)
<details><summary>context</summary>
<pre>00:13:39: ERROR: Analysis of target &#39;//:push-virt-template-controller&#39; failed; build aborted:
00:13:39: INFO: Elapsed time: 2.275s
00:13:39: INFO: 0 processes.
00:13:39: ERROR: Build failed. Not running target
00:13:40: &#43; rm -f /tmp/kubevirt.deploy.2EHQ
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)
<details>
<summary>all...</summary>

* _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)
<details><summary>context</summary>
<pre>21:11:59: &#43;&#43; _kubectl wait -n kube-system --timeout=12m --for=condition=Ready -l k8s-app=kube-dns pods
21:11:59: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.35/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.35/.kubeconfig wait -n kube-system --timeout=12m --for=condition=Ready -l k8s-app=kube-dns pods
21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-m6h9d
21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv
make: *** [Makefile:174: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)
<details>
<summary>all...</summary>

* _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)
<details><summary>context</summary>
<pre>06:17:49: 	Once you have found the failing container, you can inspect its logs with:
06:17:49: 	- &#39;crictl --runtime-endpoint unix:///run/containerd/containerd.sock logs CONTAINERID&#39;
06:17:49:
06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]
06:17:49: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
06:17:49: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262
06:17:49: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).visitAll</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (from secondary snippet) (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)
<details><summary>context</summary>
<pre>00:22:48: 	Once you have found the failing container, you can inspect its logs with:
00:22:48: 	- &#39;crictl --runtime-endpoint unix:///run/containerd/containerd.sock logs CONTAINERID&#39;
00:22:48:
00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]
00:22:48: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
00:22:48: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262
00:22:48: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).visitAll</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 11.11%)

<details>
<summary> control plane startup failure (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)
<details>
<summary>all...</summary>

* _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)
<details><summary>context</summary>
<pre>18:38:10: 	Once you have found the failing container, you can inspect its logs with:
18:38:10: 	- &#39;crictl --runtime-endpoint unix:///run/containerd/containerd.sock logs CONTAINERID&#39;
18:38:10:
18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]
18:38:10: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
18:38:10: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262
18:38:10: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).visitAll</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (1x / 11.11%)

<details>
<summary> no matching pattern (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)
<details>
<summary>all...</summary>

* _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)
<details><summary>context</summary>
<pre>23:18:41: 	- &#39;systemctl status kubelet&#39;
23:18:41: 	- &#39;journalctl -xeu kubelet&#39;
23:18:41:
23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded
23:18:41:
23:18:41: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
23:18:41: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (8x / 88.89%)


#### external (7x / 87.50%)

<details>
<summary> bazel remote cache IO failure (from secondary snippet) (2x / 25.00%) </summary>

<hr/>

**2x**: _2026-05-19 15:07:21 &#43;0000 UTC_: <code>15:33:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17833/pull-kubevirt-e2e-k8s-1.34-sig-compute/2056706870336294912#1:build-log.txt%3A4672)
<details>
<summary>all...</summary>

* _2026-05-19 15:07:21 &#43;0000 UTC_: <code>15:33:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17833/pull-kubevirt-e2e-k8s-1.34-sig-compute/2056706870336294912#1:build-log.txt%3A4672)

* _2026-05-19 10:30:44 &#43;0000 UTC_: <code>10:45:00: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16759/pull-kubevirt-e2e-k8s-1.35-sig-operator/2056649475119648768#1:build-log.txt%3A4707)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)
<details>
<summary>all...</summary>

* _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)

</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)
<details>
<summary>all...</summary>

* _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)

</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (from secondary snippet) (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache blob fetch failure (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-19 15:39:19 &#43;0000 UTC_: <code>15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11 Validating nogo output for //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16500/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2056690058336931840#1:build-log.txt%3A414)
<details>
<summary>all...</summary>

* _2026-05-19 15:39:19 &#43;0000 UTC_: <code>15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11 Validating nogo output for //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16500/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2056690058336931840#1:build-log.txt%3A414)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)

</details>

<hr/>
</details>

#### internal (1x / 12.50%)

<details>
<summary> control plane startup failure (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)
<details>
<summary>all...</summary>

* _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)

</details>

<hr/>
</details>

### release-1.7 (1x / 11.11%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no matching pattern (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)
<details>
<summary>all...</summary>

* _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (8x / 88.89%)


#### external (6x / 75.00%)

<details>
<summary> bazel remote cache IO failure (from secondary snippet) (2x / 25.00%) </summary>

<hr/>

**2x**: _2026-05-19 15:07:21 &#43;0000 UTC_: <code>15:33:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17833/pull-kubevirt-e2e-k8s-1.34-sig-compute/2056706870336294912#1:build-log.txt%3A4672)
<details>
<summary>all...</summary>

* _2026-05-19 15:07:21 &#43;0000 UTC_: <code>15:33:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17833/pull-kubevirt-e2e-k8s-1.34-sig-compute/2056706870336294912#1:build-log.txt%3A4672)

* _2026-05-19 10:30:44 &#43;0000 UTC_: <code>10:45:00: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16759/pull-kubevirt-e2e-k8s-1.35-sig-operator/2056649475119648768#1:build-log.txt%3A4707)

</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (from secondary snippet) (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)
<details>
<summary>all...</summary>

* _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)

</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)
<details>
<summary>all...</summary>

* _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)

</details>

<hr/>
</details>

#### needs-investigation (1x / 12.50%)

<details>
<summary> no matching pattern (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)
<details>
<summary>all...</summary>

* _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)

</details>

<hr/>
</details>

#### internal (1x / 12.50%)

<details>
<summary> control plane startup failure (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)
<details>
<summary>all...</summary>

* _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)

</details>

<hr/>
</details>

### sig-network (1x / 11.11%)


#### external (1x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-19 15:39:19 &#43;0000 UTC_: <code>15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11 Validating nogo output for //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16500/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2056690058336931840#1:build-log.txt%3A414)
<details>
<summary>all...</summary>

* _2026-05-19 15:39:19 &#43;0000 UTC_: <code>15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11 Validating nogo output for //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16500/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2056690058336931840#1:build-log.txt%3A414)

</details>

<hr/>
</details>

Last updated: 2026-05-26 00:52:21
