



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-05-26 (5x / 27.78%)


#### needs-investigation (3x / 60.00%)

<details>
<summary> no matching pattern (3x / 60.00%) </summary>

<hr/>

**3x**: _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)
<details>
<summary>all...</summary>

* _2026-05-26 04:07:33 &#43;0000 UTC_: <code>04:23:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059124119928049664#1:build-log.txt%3A6615)

* _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)

* _2026-05-26 00:07:20 &#43;0000 UTC_: <code>00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059063686630215680#1:build-log.txt%3A6633)

</details>

<hr/>
</details>

#### internal (2x / 40.00%)

<details>
<summary> make bazel-build target failure (2x / 40.00%) </summary>

<hr/>

**2x**: _2026-05-26 07:27:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059174345552433152#1:build-log.txt%3A203)
<details>
<summary>all...</summary>

* _2026-05-26 07:27:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059174345552433152#1:build-log.txt%3A203)

* _2026-05-26 01:07:29 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059078799525416960#1:build-log.txt%3A203)

</details>

<hr/>
</details>

### 2026-05-25 (6x / 33.33%)


#### needs-investigation (2x / 33.33%)

<details>
<summary> no matching pattern (2x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-25 22:06:15 &#43;0000 UTC_: <code>22:22:52: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059033219742830592#1:build-log.txt%3A6663)
<details>
<summary>all...</summary>

* _2026-05-25 22:06:15 &#43;0000 UTC_: <code>22:22:52: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059033219742830592#1:build-log.txt%3A6663)

</details>

<hr/>

**1x**: _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)

</details>

<hr/>
</details>

#### internal (3x / 50.00%)

<details>
<summary> make bazel-build target failure (3x / 50.00%) </summary>

<hr/>

**3x**: _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)
<details>
<summary>all...</summary>

* _2026-05-25 19:52:10 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058999436138254336#1:build-log.txt%3A212)

* _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)

* _2026-05-25 16:57:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058955558613946368#1:build-log.txt%3A212)

</details>

<hr/>
</details>

#### external (1x / 16.67%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)

</details>

<hr/>
</details>

### 2026-05-23 (1x / 5.56%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no matching pattern (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-23 00:05:29 &#43;0000 UTC_: <code>00:13:36: ERROR: failed to create cluster: failed to restart containerd after patching config: command &#34;podman exec --privileged kind-1.35-control-plane bash -c &#39;! pgrep --exact containerd || systemctl restart containerd&#39;&#34; failed with error: exit status 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17795/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057976065573261312#1:build-log.txt%3A1426)
<details>
<summary>all...</summary>

* _2026-05-23 00:05:29 &#43;0000 UTC_: <code>00:13:36: ERROR: failed to create cluster: failed to restart containerd after patching config: command &#34;podman exec --privileged kind-1.35-control-plane bash -c &#39;! pgrep --exact containerd || systemctl restart containerd&#39;&#34; failed with error: exit status 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17795/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057976065573261312#1:build-log.txt%3A1426)

</details>

<hr/>
</details>

### 2026-05-22 (3x / 16.67%)


#### external (3x / 100.00%)

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

### 2026-05-21 (3x / 16.67%)


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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (5x / 27.78%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 5.56%) </summary>

<hr/>

**1x**: _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)
<details><summary>context</summary>
<pre>09:16:06: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
09:16:07: [control-plane-check] kube-controller-manager is healthy after 1.092449818s
09:16:08: [control-plane-check] kube-scheduler is healthy after 1.978771755s
09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
09:16:10: [control-plane-check] kube-apiserver is healthy after 4.001529214s
09:16:10: I0525 05:16:10.566861    1606 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
09:16:10: I0525 05:16:10.569311    1606 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 5.56%) </summary>

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
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 5.56%) </summary>

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
<summary> API rate limiter timeout (1x / 5.56%) </summary>

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
<summary> API rate limiter timeout (from secondary snippet) (1x / 5.56%) </summary>

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

### internal (6x / 33.33%)

<details>
<summary> make bazel-build target failure (5x / 27.78%) </summary>

<hr/>

**5x**: _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)
<details>
<summary>all...</summary>

* _2026-05-26 07:27:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059174345552433152#1:build-log.txt%3A203)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
07:27:31: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-efe2b228ff2a1ce1d95aab1fb13854b583e8e74cebd8c3956352cd52d22d5414`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-26 01:07:29 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059078799525416960#1:build-log.txt%3A203)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
01:07:58: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-90a0783b2881336943496bef956125c62eaf5e7261df0077ef92209fc1315874`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-25 19:52:10 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058999436138254336#1:build-log.txt%3A212)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
19:52:39: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-3b6963538642576dc5e085a242ae8cf52a73f2621d3b14cf4115284e2b9e3de1`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
17:11:16: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-b0d249f272298390f6f04648e0394a2014389781537f15f35b3ee15c3c28c9e7`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


* _2026-05-25 16:57:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058955558613946368#1:build-log.txt%3A212)
<details><summary>context</summary>
<pre>&#43; make bazel-build-images
hack/dockerized &#34;export BUILD_ARCH= &amp;&amp; DOCKER_PREFIX= DOCKER_TAG= DOCKER_TAG_ALT= IMAGE_PREFIX= IMAGE_PREFIX_ALT=kv- ./hack/multi-arch.sh build-images&#34;
16:58:09: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-50696fa3ed6b5067c2a55c35d8faaec48ad10fd9f74941d9ddbd6820ae8cf086`: No space left on device
make: *** [Makefile:39: bazel-build-images] Error 126
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> control plane startup failure (1x / 5.56%) </summary>

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

### needs-investigation (7x / 38.89%)

<details>
<summary> no matching pattern (7x / 38.89%) </summary>

<hr/>

**4x**: _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)
<details>
<summary>all...</summary>

* _2026-05-26 04:07:33 &#43;0000 UTC_: <code>04:23:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059124119928049664#1:build-log.txt%3A6615)
<details><summary>context</summary>
<pre>04:23:46: 	- &#39;systemctl status kubelet&#39;
04:23:46: 	- &#39;journalctl -xeu kubelet&#39;
04:23:46:
04:23:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded
04:23:46:
04:23:46: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
04:23:46: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


* _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)
<details><summary>context</summary>
<pre>02:23:47: 	- &#39;systemctl status kubelet&#39;
02:23:47: 	- &#39;journalctl -xeu kubelet&#39;
02:23:47:
02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded
02:23:47:
02:23:47: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
02:23:47: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


* _2026-05-26 00:07:20 &#43;0000 UTC_: <code>00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059063686630215680#1:build-log.txt%3A6633)
<details><summary>context</summary>
<pre>00:23:57: 	- &#39;systemctl status kubelet&#39;
00:23:57: 	- &#39;journalctl -xeu kubelet&#39;
00:23:57:
00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused
00:23:57:
00:23:57: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
00:23:57: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


* _2026-05-25 22:06:15 &#43;0000 UTC_: <code>22:22:52: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059033219742830592#1:build-log.txt%3A6663)
<details><summary>context</summary>
<pre>22:22:52: 	- &#39;systemctl status kubelet&#39;
22:22:52: 	- &#39;journalctl -xeu kubelet&#39;
22:22:52:
22:22:52: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded
22:22:52:
22:22:52: k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow.(*Runner).Run.func1
22:22:52: 	k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow/runner.go:262</pre>
</details>


</details>

<hr/>

**1x**: _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)
<details><summary>context</summary>
<pre>08:25:53:  ✗ Preparing nodes 📦 📦
08:25:53: ERROR: failed to create cluster: command &#34;podman run --name kind-1.34-control-plane --hostname kind-1.34-control-plane --label io.x-k8s.kind.role=control-plane --privileged --tmpfs /tmp --tmpfs /run --volume e972a0936de444675a134034826b6ab581cc5cb5273dd45e9252f096913c1e7b:/var:suid,exec,dev --volume /lib/modules:/lib/modules:ro -e KIND_EXPERIMENTAL_CONTAINERD_SNAPSHOTTER --detach --tty --net kind --label io.x-k8s.kind.cluster=kind-1.34 -e container=podman --cgroupns=private --volume /dev/mapper:/dev/mapper --volume=/var/log/audit:/var/log/audit:ro --publish=127.0.0.1:34443:6443/tcp -e KUBECONFIG=/etc/kubernetes/admin.conf docker.io/kindest/node@sha256:08497ee19eace7b4b5348db5c6a1591d7752b164530a36f855cb0f2bdcbadd48&#34; failed with error: exit status 126
08:25:53:
08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device
08:25:53:
08:25:53: Stack Trace:
08:25:53: sigs.k8s.io/kind/pkg/errors.WithStack</pre>
</details>


</details>

<hr/>

**1x**: _2026-05-23 00:05:29 &#43;0000 UTC_: <code>00:13:36: ERROR: failed to create cluster: failed to restart containerd after patching config: command &#34;podman exec --privileged kind-1.35-control-plane bash -c &#39;! pgrep --exact containerd || systemctl restart containerd&#39;&#34; failed with error: exit status 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17795/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057976065573261312#1:build-log.txt%3A1426)
<details>
<summary>all...</summary>

* _2026-05-23 00:05:29 &#43;0000 UTC_: <code>00:13:36: ERROR: failed to create cluster: failed to restart containerd after patching config: command &#34;podman exec --privileged kind-1.35-control-plane bash -c &#39;! pgrep --exact containerd || systemctl restart containerd&#39;&#34; failed with error: exit status 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17795/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057976065573261312#1:build-log.txt%3A1426)
<details><summary>context</summary>
<pre>00:13:34: kind: KubeProxyConfiguration
00:13:34: mode: iptables
00:13:36:  ✗ Writing configuration 📜
00:13:36: ERROR: failed to create cluster: failed to restart containerd after patching config: command &#34;podman exec --privileged kind-1.35-control-plane bash -c &#39;! pgrep --exact containerd || systemctl restart containerd&#39;&#34; failed with error: exit status 1
00:13:36:
00:13:36: Command Output: 109
00:13:36: Job for containerd.service failed because the control process exited with error code.</pre>
</details>


</details>

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


### release-1.8 (4x / 22.22%)


#### needs-investigation (3x / 75.00%)

<details>
<summary> no matching pattern (3x / 75.00%) </summary>

<hr/>

**2x**: _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)
<details>
<summary>all...</summary>

* _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)

* _2026-05-25 22:06:15 &#43;0000 UTC_: <code>22:22:52: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059033219742830592#1:build-log.txt%3A6663)

</details>

<hr/>

**1x**: _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)

</details>

<hr/>
</details>

#### internal (1x / 25.00%)

<details>
<summary> make bazel-build target failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)
<details>
<summary>all...</summary>

* _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)

</details>

<hr/>
</details>

### main (13x / 72.22%)


#### external (5x / 38.46%)

<details>
<summary> download failure in context (1x / 7.69%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 7.69%) </summary>

<hr/>

**1x**: _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)
<details>
<summary>all...</summary>

* _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)

</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (1x / 7.69%) </summary>

<hr/>

**1x**: _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)
<details>
<summary>all...</summary>

* _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)

</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (from secondary snippet) (1x / 7.69%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 7.69%) </summary>

<hr/>

**1x**: _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)

</details>

<hr/>
</details>

#### internal (5x / 38.46%)

<details>
<summary> make bazel-build target failure (4x / 30.77%) </summary>

<hr/>

**4x**: _2026-05-25 19:52:10 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058999436138254336#1:build-log.txt%3A212)
<details>
<summary>all...</summary>

* _2026-05-26 07:27:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059174345552433152#1:build-log.txt%3A203)

* _2026-05-26 01:07:29 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059078799525416960#1:build-log.txt%3A203)

* _2026-05-25 19:52:10 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058999436138254336#1:build-log.txt%3A212)

* _2026-05-25 16:57:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058955558613946368#1:build-log.txt%3A212)

</details>

<hr/>
</details>
<details>
<summary> control plane startup failure (1x / 7.69%) </summary>

<hr/>

**1x**: _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)
<details>
<summary>all...</summary>

* _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)

</details>

<hr/>
</details>

#### needs-investigation (3x / 23.08%)

<details>
<summary> no matching pattern (3x / 23.08%) </summary>

<hr/>

**2x**: _2026-05-26 00:07:20 &#43;0000 UTC_: <code>00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059063686630215680#1:build-log.txt%3A6633)
<details>
<summary>all...</summary>

* _2026-05-26 04:07:33 &#43;0000 UTC_: <code>04:23:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059124119928049664#1:build-log.txt%3A6615)

* _2026-05-26 00:07:20 &#43;0000 UTC_: <code>00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059063686630215680#1:build-log.txt%3A6633)

</details>

<hr/>

**1x**: _2026-05-23 00:05:29 &#43;0000 UTC_: <code>00:13:36: ERROR: failed to create cluster: failed to restart containerd after patching config: command &#34;podman exec --privileged kind-1.35-control-plane bash -c &#39;! pgrep --exact containerd || systemctl restart containerd&#39;&#34; failed with error: exit status 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17795/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057976065573261312#1:build-log.txt%3A1426)
<details>
<summary>all...</summary>

* _2026-05-23 00:05:29 &#43;0000 UTC_: <code>00:13:36: ERROR: failed to create cluster: failed to restart containerd after patching config: command &#34;podman exec --privileged kind-1.35-control-plane bash -c &#39;! pgrep --exact containerd || systemctl restart containerd&#39;&#34; failed with error: exit status 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17795/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057976065573261312#1:build-log.txt%3A1426)

</details>

<hr/>
</details>

### release-1.7 (1x / 5.56%)


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


### sig-compute (17x / 94.44%)


#### external (4x / 23.53%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 5.88%) </summary>

<hr/>

**1x**: _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)
<details>
<summary>all...</summary>

* _2026-05-21 21:03:25 &#43;0000 UTC_: <code>21:23:59: timed out waiting for the condition on pods/coredns-7d764666f9-wzncv</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057567382322286592#1:build-log.txt%3A1224)

</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (1x / 5.88%) </summary>

<hr/>

**1x**: _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)
<details>
<summary>all...</summary>

* _2026-05-22 06:01:24 &#43;0000 UTC_: <code>06:17:49: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: client rate limiter Wait returned an error: context deadline exceeded, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057703244842405888#1:build-log.txt%3A10629)

</details>

<hr/>
</details>
<details>
<summary> API rate limiter timeout (from secondary snippet) (1x / 5.88%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:29 &#43;0000 UTC_: <code>00:22:48: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-apiserver check failed at https://10.89.0.5:6443/livez: Get &#34;https://10.89.0.5:6443/livez?timeout=10s&#34;: dial tcp 10.89.0.5:6443: connect: connection refused]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057612824133242880#1:build-log.txt%3A9756)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 5.88%) </summary>

<hr/>

**1x**: _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)
<details>
<summary>all...</summary>

* _2026-05-22 00:02:20 &#43;0000 UTC_: <code>00:13:39: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17866/pull-kubevirt-e2e-kind-1.34-sev/2057612816566718464#1:build-log.txt%3A4457)

</details>

<hr/>
</details>

#### internal (6x / 35.29%)

<details>
<summary> make bazel-build target failure (5x / 29.41%) </summary>

<hr/>

**5x**: _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)
<details>
<summary>all...</summary>

* _2026-05-26 07:27:02 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059174345552433152#1:build-log.txt%3A203)

* _2026-05-26 01:07:29 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059078799525416960#1:build-log.txt%3A203)

* _2026-05-25 19:52:10 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058999436138254336#1:build-log.txt%3A212)

* _2026-05-25 17:10:43 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058958831135756288#1:build-log.txt%3A204)

* _2026-05-25 16:57:40 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 126</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2058955558613946368#1:build-log.txt%3A212)

</details>

<hr/>
</details>
<details>
<summary> control plane startup failure (1x / 5.88%) </summary>

<hr/>

**1x**: _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)
<details>
<summary>all...</summary>

* _2026-05-21 18:17:18 &#43;0000 UTC_: <code>18:38:10: error: error execution phase wait-control-plane: failed while waiting for the control plane to start: [kube-controller-manager check failed at https://127.0.0.1:10257/healthz: Get &#34;https://127.0.0.1:10257/healthz&#34;: dial tcp 127.0.0.1:10257: connect: connection refused, kube-scheduler check failed at https://127.0.0.1:10259/livez: Get &#34;https://127.0.0.1:10259/livez&#34;: dial tcp 127.0.0.1:10259: connect: connection refused, kube-apiserver check failed at https://10.89.0.4:6443/livez: client rate limiter Wait returned an error: rate: Wait(n=1) would exceed context deadline]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17874/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057525899489710080#1:build-log.txt%3A5710)

</details>

<hr/>
</details>

#### needs-investigation (7x / 41.18%)

<details>
<summary> no matching pattern (7x / 41.18%) </summary>

<hr/>

**4x**: _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)
<details>
<summary>all...</summary>

* _2026-05-26 04:07:33 &#43;0000 UTC_: <code>04:23:46: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17100/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059124119928049664#1:build-log.txt%3A6615)

* _2026-05-26 02:07:30 &#43;0000 UTC_: <code>02:23:47: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059093907123212288#1:build-log.txt%3A6643)

* _2026-05-26 00:07:20 &#43;0000 UTC_: <code>00:23:57: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: dial tcp 127.0.0.1:10248: connect: connection refused</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2059063686630215680#1:build-log.txt%3A6633)

* _2026-05-25 22:06:15 &#43;0000 UTC_: <code>22:22:52: error: error execution phase wait-control-plane: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2059033219742830592#1:build-log.txt%3A6663)

</details>

<hr/>

**1x**: _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)
<details>
<summary>all...</summary>

* _2026-05-25 08:14:21 &#43;0000 UTC_: <code>08:25:53: Command Output: Error: OCI runtime error: crun: create `/sys/fs/cgroup/libpod_parent/libpod-e2bf97930c84c036cb4bc89deb6708e780df3ddcdb4f359553f73f210835653b`: No space left on device</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17876/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.8/2058823736756277248#1:build-log.txt%3A652)

</details>

<hr/>

**1x**: _2026-05-23 00:05:29 &#43;0000 UTC_: <code>00:13:36: ERROR: failed to create cluster: failed to restart containerd after patching config: command &#34;podman exec --privileged kind-1.35-control-plane bash -c &#39;! pgrep --exact containerd || systemctl restart containerd&#39;&#34; failed with error: exit status 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17795/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057976065573261312#1:build-log.txt%3A1426)
<details>
<summary>all...</summary>

* _2026-05-23 00:05:29 &#43;0000 UTC_: <code>00:13:36: ERROR: failed to create cluster: failed to restart containerd after patching config: command &#34;podman exec --privileged kind-1.35-control-plane bash -c &#39;! pgrep --exact containerd || systemctl restart containerd&#39;&#34; failed with error: exit status 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17795/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2057976065573261312#1:build-log.txt%3A1426)

</details>

<hr/>

**1x**: _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)
<details>
<summary>all...</summary>

* _2026-05-21 23:01:27 &#43;0000 UTC_: <code>23:18:41: error: error execution phase kubelet-wait-bootstrap: failed while waiting for the kubelet to start: The HTTP call equal to &#39;curl -sSL http://127.0.0.1:10248/healthz&#39; returned error: Get &#34;http://127.0.0.1:10248/healthz&#34;: context deadline exceeded</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17728/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2057597520984412160#1:build-log.txt%3A1943)

</details>

<hr/>
</details>

### sig-storage (1x / 5.56%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)
<details>
<summary>all...</summary>

* _2026-05-25 09:12:21 &#43;0000 UTC_: <code>09:16:10: I0525 05:16:10.064112    1606 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17825/pull-kubevirt-e2e-k8s-1.34-sig-storage/2058830816263278592#1:build-log.txt%3A762)

</details>

<hr/>
</details>

Last updated: 2026-05-27 16:26:18
