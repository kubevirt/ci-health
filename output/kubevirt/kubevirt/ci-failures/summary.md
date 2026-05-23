



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-05-22 (3x / 30.00%)


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

### 2026-05-21 (1x / 10.00%)


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

### 2026-05-19 (3x / 30.00%)


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

### 2026-05-18 (3x / 30.00%)


#### pr-build (2x / 66.67%)

<details>
<summary> ginkgo test failure marker (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)
<details>
<summary>all...</summary>

* _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)

* _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)

</details>

<hr/>
</details>

#### internal (1x / 33.33%)

<details>
<summary> kind cluster creation failure (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)
<details>
<summary>all...</summary>

* _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (6x / 60.00%)

<details>
<summary> bazel remote cache IO failure (from secondary snippet) (2x / 20.00%) </summary>

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
<summary> download failure in context (1x / 10.00%) </summary>

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
<summary> API rate limiter timeout (1x / 10.00%) </summary>

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
<summary> API rate limiter timeout (from secondary snippet) (1x / 10.00%) </summary>

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
<details>
<summary> bazel remote cache blob fetch failure (1x / 10.00%) </summary>

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

### pr-build (2x / 20.00%)

<details>
<summary> ginkgo test failure marker (2x / 20.00%) </summary>

<hr/>

**2x**: _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)
<details>
<summary>all...</summary>

* _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)
<details><summary>context</summary>
<pre>11:00:00:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;system is in sync with kubevirt config resource version 206657&#34;,&#34;pos&#34;:&#34;kvconfig.go:101&#34;,&#34;timestamp&#34;:&#34;2026-05-18T11:00:00.128573Z&#34;}
11:00:00:   &lt;&lt; Captured StdOut/StdErr Output
11:00:00: ------------------------------
11:09:10: • [FAILED] [549.986 seconds]
11:09:10: [sig-storage] Hotplug with PCI hostdev [It] should restart a VM after hotplugging a block volume [sig-storage, Serial, RequiresBlockStorage]
11:09:10: tests/storage/hotplug.go:2216
11:09:10:</pre>
</details>


* _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)
<details><summary>context</summary>
<pre>11:00:05:   {&#34;component&#34;:&#34;portforward&#34;,&#34;level&#34;:&#34;info&#34;,&#34;msg&#34;:&#34;system is in sync with kubevirt config resource version 200331&#34;,&#34;pos&#34;:&#34;kvconfig.go:102&#34;,&#34;timestamp&#34;:&#34;2026-05-18T11:00:05.879219Z&#34;}
11:00:05:   &lt;&lt; Captured StdOut/StdErr Output
11:00:05: ------------------------------
11:09:36: • [FAILED] [570.279 seconds]
11:09:36: [sig-storage] SCSI persistent reservation [BeforeEach] with PersistentReservation feature gate toggled should delete and recreate virt-handler [sig-storage, Serial]
11:09:36:   [BeforeEach] tests/storage/reservation.go:186
11:09:36:   [It] tests/storage/reservation.go:344</pre>
</details>


</details>

<hr/>
</details>

### needs-investigation (1x / 10.00%)

<details>
<summary> no matching pattern (1x / 10.00%) </summary>

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

### internal (1x / 10.00%)

<details>
<summary> kind cluster creation failure (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)
<details>
<summary>all...</summary>

* _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)
<details><summary>context</summary>
<pre>10:29:27:  ✓ Ensuring node image (kindest/node:v1.34.3) 🖼
10:29:27:  • Preparing nodes 📦 📦   ...
10:29:29:  ✗ Preparing nodes 📦 📦
10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;
10:29:29:
10:29:29: Stack Trace:
10:29:29: sigs.k8s.io/kind/pkg/errors.Errorf</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (8x / 80.00%)


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

#### pr-build (1x / 12.50%)

<details>
<summary> ginkgo test failure marker (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)
<details>
<summary>all...</summary>

* _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)

</details>

<hr/>
</details>

#### internal (1x / 12.50%)

<details>
<summary> kind cluster creation failure (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)
<details>
<summary>all...</summary>

* _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)

</details>

<hr/>
</details>

### release-1.7 (1x / 10.00%)


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

### release-1.8 (1x / 10.00%)


#### pr-build (1x / 100.00%)

<details>
<summary> ginkgo test failure marker (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)
<details>
<summary>all...</summary>

* _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (7x / 70.00%)


#### external (5x / 71.43%)

<details>
<summary> bazel remote cache IO failure (from secondary snippet) (2x / 28.57%) </summary>

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
<summary> download failure in context (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)

</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)
<details>
<summary>all...</summary>

* _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)

</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (from secondary snippet) (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)

</details>

<hr/>
</details>

#### needs-investigation (1x / 14.29%)

<details>
<summary> no matching pattern (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)
<details>
<summary>all...</summary>

* _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)

</details>

<hr/>
</details>

#### internal (1x / 14.29%)

<details>
<summary> kind cluster creation failure (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)
<details>
<summary>all...</summary>

* _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)

</details>

<hr/>
</details>

### sig-network (1x / 10.00%)


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

### sig-storage (2x / 20.00%)


#### pr-build (2x / 100.00%)

<details>
<summary> ginkgo test failure marker (2x / 100.00%) </summary>

<hr/>

**2x**: _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)
<details>
<summary>all...</summary>

* _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)

* _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)

</details>

<hr/>
</details>

Last updated: 2026-05-23 00:36:57
