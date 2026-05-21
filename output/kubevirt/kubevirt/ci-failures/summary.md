



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-05-19 (3x / 25.00%)


#### external (3x / 100.00%)

<details>
<summary> bazel remote cache IO failure (from secondary snippet) (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-05-19 10:30:44 &#43;0000 UTC_: <code>10:45:00: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16759/pull-kubevirt-e2e-k8s-1.35-sig-operator/2056649475119648768#1:build-log.txt%3A4707)
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

### 2026-05-18 (3x / 25.00%)


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

#### pr-build (2x / 66.67%)

<details>
<summary> ginkgo test failure marker (2x / 66.67%) </summary>

<hr/>

**2x**: _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)
<details>
<summary>all...</summary>

* _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)

* _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)

</details>

<hr/>
</details>

### 2026-05-15 (5x / 41.67%)


#### needs-investigation (4x / 80.00%)

<details>
<summary> no error snippets (4x / 80.00%) </summary>

<hr/>

**1x**: _2026-05-15 11:34:55 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)
<details>
<summary>all...</summary>

* _2026-05-15 11:34:55 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)

</details>

<hr/>

**1x**: _2026-05-15 08:44:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)
<details>
<summary>all...</summary>

* _2026-05-15 08:44:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)

</details>

<hr/>

**1x**: _2026-05-15 07:29:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)

</details>

<hr/>

**1x**: _2026-05-15 07:29:22 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:22 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)

</details>

<hr/>
</details>

#### external (1x / 20.00%)

<details>
<summary> transient kube-apiserver body decode noise (1x / 20.00%) </summary>

<hr/>

**1x**: _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)
<details>
<summary>all...</summary>

* _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)

</details>

<hr/>
</details>

### 2026-05-14 (1x / 8.33%)


#### pr-build (1x / 100.00%)

<details>
<summary> ginkgo test failure marker (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)
<details>
<summary>all...</summary>

* _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (4x / 33.33%)

<details>
<summary> bazel remote cache IO failure (from secondary snippet) (2x / 16.67%) </summary>

<hr/>

**2x**: _2026-05-19 10:30:44 &#43;0000 UTC_: <code>10:45:00: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16759/pull-kubevirt-e2e-k8s-1.35-sig-operator/2056649475119648768#1:build-log.txt%3A4707)
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
<summary> bazel remote cache blob fetch failure (1x / 8.33%) </summary>

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
<summary> transient kube-apiserver body decode noise (1x / 8.33%) </summary>

<hr/>

**1x**: _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)
<details>
<summary>all...</summary>

* _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)
<details><summary>context</summary>
<pre>04:56:31: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
04:56:34: [control-plane-check] kube-scheduler is healthy after 3.507988268s
04:56:35: [control-plane-check] kube-controller-manager is healthy after 4.052522673s
04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
04:56:39: [control-plane-check] kube-apiserver is healthy after 8.502646674s
04:56:39: I0515 00:56:39.582491    1610 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
04:56:39: I0515 00:56:39.584077    1610 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>


</details>

<hr/>
</details>

### internal (1x / 8.33%)

<details>
<summary> kind cluster creation failure (1x / 8.33%) </summary>

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

### needs-investigation (4x / 33.33%)

<details>
<summary> no error snippets (4x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-15 11:34:55 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)
<details>
<summary>all...</summary>

* _2026-05-15 11:34:55 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)

</details>

<hr/>

**1x**: _2026-05-15 08:44:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)
<details>
<summary>all...</summary>

* _2026-05-15 08:44:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)

</details>

<hr/>

**1x**: _2026-05-15 07:29:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)

</details>

<hr/>

**1x**: _2026-05-15 07:29:22 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:22 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)

</details>

<hr/>
</details>

### pr-build (3x / 25.00%)

<details>
<summary> ginkgo test failure marker (3x / 25.00%) </summary>

<hr/>

**1x**: _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)
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


</details>

<hr/>

**2x**: _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)
<details>
<summary>all...</summary>

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


* _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)
<details><summary>context</summary>
<pre>20:14:00:
20:14:00:   There were additional failures detected.  To view them in detail run ginkgo -vv
20:14:00: ------------------------------
20:18:29: • [FAILED] [269.183 seconds]
20:18:29: [rfe_id:393][crit:high][vendor:cnv-qe@redhat.com][level:system][sig-compute] Live Migration with a live-migrate eviction strategy set [ref_id:2293</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (10x / 83.33%)


#### external (4x / 40.00%)

<details>
<summary> bazel remote cache IO failure (from secondary snippet) (2x / 20.00%) </summary>

<hr/>

**2x**: _2026-05-19 10:30:44 &#43;0000 UTC_: <code>10:45:00: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16759/pull-kubevirt-e2e-k8s-1.35-sig-operator/2056649475119648768#1:build-log.txt%3A4707)
<details>
<summary>all...</summary>

* _2026-05-19 15:07:21 &#43;0000 UTC_: <code>15:33:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17833/pull-kubevirt-e2e-k8s-1.34-sig-compute/2056706870336294912#1:build-log.txt%3A4672)

* _2026-05-19 10:30:44 &#43;0000 UTC_: <code>10:45:00: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16759/pull-kubevirt-e2e-k8s-1.35-sig-operator/2056649475119648768#1:build-log.txt%3A4707)

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

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)
<details>
<summary>all...</summary>

* _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)

</details>

<hr/>
</details>

#### needs-investigation (4x / 40.00%)

<details>
<summary> no error snippets (4x / 40.00%) </summary>

<hr/>

**1x**: _2026-05-15 11:34:55 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)
<details>
<summary>all...</summary>

* _2026-05-15 11:34:55 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)

</details>

<hr/>

**1x**: _2026-05-15 08:44:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)
<details>
<summary>all...</summary>

* _2026-05-15 08:44:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)

</details>

<hr/>

**1x**: _2026-05-15 07:29:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)

</details>

<hr/>

**1x**: _2026-05-15 07:29:22 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:22 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)

</details>

<hr/>
</details>

#### internal (1x / 10.00%)

<details>
<summary> kind cluster creation failure (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)
<details>
<summary>all...</summary>

* _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)

</details>

<hr/>
</details>

#### pr-build (1x / 10.00%)

<details>
<summary> ginkgo test failure marker (1x / 10.00%) </summary>

<hr/>

**1x**: _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)
<details>
<summary>all...</summary>

* _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)

</details>

<hr/>
</details>

### release-1.8 (1x / 8.33%)


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

### release-1.6 (1x / 8.33%)


#### pr-build (1x / 100.00%)

<details>
<summary> ginkgo test failure marker (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)
<details>
<summary>all...</summary>

* _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-network (2x / 16.67%)


#### external (2x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-05-19 15:39:19 &#43;0000 UTC_: <code>15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11 Validating nogo output for //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16500/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2056690058336931840#1:build-log.txt%3A414)
<details>
<summary>all...</summary>

* _2026-05-19 15:39:19 &#43;0000 UTC_: <code>15:53:38: ERROR: /root/go/src/kubevirt.io/kubevirt/pkg/monitoring/metrics/common/labels/BUILD.bazel:3:11 Validating nogo output for //pkg/monitoring/metrics/common/labels:go_default_library failed: Failed to fetch blobs because they do not exist remotely.: Missing digest: bde6d7dca14b3f54ffbf70a90243fc32dc05c7f9396548125b1c2e39a6940d2f/65452 for bazel-out/k8-fastbuild/bin/staging/src/kubevirt.io/client-go/log/go_default_library.x</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16500/pull-kubevirt-e2e-k8s-1.35-ipv6-sig-network/2056690058336931840#1:build-log.txt%3A414)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)
<details>
<summary>all...</summary>

* _2026-05-15 04:52:36 &#43;0000 UTC_: <code>04:56:39: I0515 00:56:39.079139    1610 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17771/pull-kubevirt-e2e-k8s-1.34-sig-network/2055146025328316416#1:build-log.txt%3A753)

</details>

<hr/>
</details>

### sig-compute (6x / 50.00%)


#### pr-build (1x / 16.67%)

<details>
<summary> ginkgo test failure marker (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)
<details>
<summary>all...</summary>

* _2026-05-14 18:29:24 &#43;0000 UTC_: <code>20:18:29: • [FAILED] [269.183 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17801/pull-kubevirt-e2e-k8s-1.32-sig-compute-migrations-1.6/2054959179134668800#1:build-log.txt%3A4093)

</details>

<hr/>
</details>

#### needs-investigation (2x / 33.33%)

<details>
<summary> no error snippets (2x / 33.33%) </summary>

<hr/>

**1x**: _2026-05-15 11:34:55 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)
<details>
<summary>all...</summary>

* _2026-05-15 11:34:55 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16257/pull-kubevirt-e2e-k8s-1.35-sig-operator/2055241605253697536)

</details>

<hr/>

**1x**: _2026-05-15 07:29:22 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:22 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.34-sig-operator/2055186770034692096)

</details>

<hr/>
</details>

#### external (2x / 33.33%)

<details>
<summary> bazel remote cache IO failure (from secondary snippet) (2x / 33.33%) </summary>

<hr/>

**2x**: _2026-05-19 10:30:44 &#43;0000 UTC_: <code>10:45:00: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16759/pull-kubevirt-e2e-k8s-1.35-sig-operator/2056649475119648768#1:build-log.txt%3A4707)
<details>
<summary>all...</summary>

* _2026-05-19 15:07:21 &#43;0000 UTC_: <code>15:33:10: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17833/pull-kubevirt-e2e-k8s-1.34-sig-compute/2056706870336294912#1:build-log.txt%3A4672)

* _2026-05-19 10:30:44 &#43;0000 UTC_: <code>10:45:00: ERROR: Build failed. Not running target</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/16759/pull-kubevirt-e2e-k8s-1.35-sig-operator/2056649475119648768#1:build-log.txt%3A4707)

</details>

<hr/>
</details>

#### internal (1x / 16.67%)

<details>
<summary> kind cluster creation failure (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)
<details>
<summary>all...</summary>

* _2026-05-18 10:20:53 &#43;0000 UTC_: <code>10:29:29: ERROR: failed to create cluster: could not find a log line that matches &#34;Reached target .*Multi-User System.*|detected cgroup v1&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17257/pull-kubevirt-e2e-kind-1.34-sev/2056318973946892288#1:build-log.txt%3A731)

</details>

<hr/>
</details>

### sig-storage (4x / 33.33%)


#### needs-investigation (2x / 50.00%)

<details>
<summary> no error snippets (2x / 50.00%) </summary>

<hr/>

**1x**: _2026-05-15 08:44:00 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)
<details>
<summary>all...</summary>

* _2026-05-15 08:44:00 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17351/pull-kubevirt-e2e-k8s-1.34-sig-storage/2055207418115133440)

</details>

<hr/>

**1x**: _2026-05-15 07:29:39 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)
<details>
<summary>all...</summary>

* _2026-05-15 07:29:39 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17802/pull-kubevirt-e2e-k8s-1.35-sig-storage/2055186770911301632)

</details>

<hr/>
</details>

#### pr-build (2x / 50.00%)

<details>
<summary> ginkgo test failure marker (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)
<details>
<summary>all...</summary>

* _2026-05-18 07:52:18 &#43;0000 UTC_: <code>11:09:10: • [FAILED] [549.986 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17829/pull-kubevirt-e2e-k8s-1.35-sig-storage/2056281442882686976#1:build-log.txt%3A9108)

* _2026-05-18 07:45:55 &#43;0000 UTC_: <code>11:09:36: • [FAILED] [570.279 seconds]</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17828/pull-kubevirt-e2e-k8s-1.34-sig-storage-1.8/2056279962301763584#1:build-log.txt%3A9570)

</details>

<hr/>
</details>

Last updated: 2026-05-21 16:23:51
