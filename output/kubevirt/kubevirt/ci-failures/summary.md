



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-08-17 (2x / 50.00%)


#### external (2x / 100.00%)

<details>
<summary> download failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-08-17 13:46:57 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18826/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2089348140795695104#1:build-log.txt%3A248)
<details>
<summary>all...</summary>

* _2026-08-17 13:46:57 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18826/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2089348140795695104#1:build-log.txt%3A248)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-08-17 06:54:40 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18735/pull-kubevirt-e2e-kind-1.36-sev/2089244337438199808#1:build-log.txt%3A1685)
<details>
<summary>all...</summary>

* _2026-08-17 06:54:40 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18735/pull-kubevirt-e2e-kind-1.36-sev/2089244337438199808#1:build-log.txt%3A1685)

</details>

<hr/>
</details>

### 2026-08-15 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-15 07:22:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18819/pull-kubevirt-e2e-kind-1.35-vgpu/2088526644472975360#1:build-log.txt%3A444)
<details>
<summary>all...</summary>

* _2026-08-15 07:22:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18819/pull-kubevirt-e2e-kind-1.35-vgpu/2088526644472975360#1:build-log.txt%3A444)

</details>

<hr/>
</details>

### 2026-08-12 (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-12 16:38:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18172/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2087556212068454400#1:build-log.txt%3A2630)
<details>
<summary>all...</summary>

* _2026-08-12 16:38:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18172/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2087556212068454400#1:build-log.txt%3A2630)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (4x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-08-12 16:38:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18172/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2087556212068454400#1:build-log.txt%3A2630)
<details>
<summary>all...</summary>

* _2026-08-17 06:54:40 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18735/pull-kubevirt-e2e-kind-1.36-sev/2089244337438199808#1:build-log.txt%3A1685)
<details><summary>context</summary>
<pre>07:01:32: INFO: Running command line: bazel-bin/push-alpine-with-test-tooling-container-disk.exe --registry localhost:5000/kubevirt --repository alpine-with-test-tooling-container-disk --tag devel
07:01:40: Error during deploy: deploying images: Patch &#34;http://localhost:5000/v2/kubevirt/alpine-with-test-tooling-container-disk/blobs/uploads/6507ebe3-6186-468b-ad56-502f613a66a0?_state=rN4_2nF7YRUhTWXPps-dXAkz5dp7PtpMsz3WZeWK0Pd7Ik5hbWUiOiJrdWJldmlydC9hbHBpbmUtd2l0aC10ZXN0LXRvb2xpbmctY29udGFpbmVyLWRpc2siLCJVVUlEIjoiNjUwN2ViZTMtNjE4Ni00NjhiLWFkNTYtNTAyZjYxM2E2NmEwIiwiT2Zmc2V0IjowLCJTdGFydGVkQXQiOiIyMDI2LTA4LTE3VDA3OjAxOjM5LjQ3OTUwMjg5NVoifQ%3D%3D&#34;: readfrom tcp 127.0.0.1:59096-&gt;127.0.0.1:5000: unexpected EOF
07:01:41: &#43; rm -f /tmp/kubevirt.deploy.cw8p
make: *** [Makefile:189: cluster-sync] Error 1
&#43; ret=2
&#43; check_for_panics
&#43; set &#43;x</pre>
</details>


* _2026-08-12 16:38:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18172/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2087556212068454400#1:build-log.txt%3A2630)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
16:46:35: selecting podman as container runtime
16:47:10: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated is being used by the following container(s): deeca495cdd3bb06d151d4ddc42a242d7b67c0f385beae0267b981de257c431d: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-08-17 13:46:57 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18826/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2089348140795695104#1:build-log.txt%3A248)
<details>
<summary>all...</summary>

* _2026-08-17 13:46:57 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18826/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2089348140795695104#1:build-log.txt%3A248)
<details><summary>context</summary>
<pre>13:51:36: INFO: 1 process: 1 internal.
13:51:36: ERROR: Build did NOT complete successfully
13:51:36: ERROR: Build failed. Not running target
make: *** [Makefile:39: bazel-build-images] Error 1
&#43; rc=2
&#43; return 2
&#43; ret=2</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-08-15 07:22:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18819/pull-kubevirt-e2e-kind-1.35-vgpu/2088526644472975360#1:build-log.txt%3A444)
<details>
<summary>all...</summary>

* _2026-08-15 07:22:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18819/pull-kubevirt-e2e-kind-1.35-vgpu/2088526644472975360#1:build-log.txt%3A444)
<details><summary>context</summary>
<pre>07:29:28: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
07:29:28: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
07:29:35: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage3412874203/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20260815%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20260815T072930Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=cf5dca3a00a64fb6b9cd358b868f25f8171e0f68d23744f3cee01d80fc56f813&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1786779870~hmac=3f582bb27ae661e4a8f2305613f481d277d263a11d6a9726b311f17835e3cdd4&#34;: EOF)
make: *** [Makefile:174: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (4x / 100.00%)


#### external (4x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-08-12 16:38:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18172/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2087556212068454400#1:build-log.txt%3A2630)
<details>
<summary>all...</summary>

* _2026-08-17 06:54:40 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18735/pull-kubevirt-e2e-kind-1.36-sev/2089244337438199808#1:build-log.txt%3A1685)

* _2026-08-12 16:38:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18172/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2087556212068454400#1:build-log.txt%3A2630)

</details>

<hr/>
</details>
<details>
<summary> download failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-08-17 13:46:57 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18826/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2089348140795695104#1:build-log.txt%3A248)
<details>
<summary>all...</summary>

* _2026-08-17 13:46:57 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18826/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2089348140795695104#1:build-log.txt%3A248)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-08-15 07:22:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18819/pull-kubevirt-e2e-kind-1.35-vgpu/2088526644472975360#1:build-log.txt%3A444)
<details>
<summary>all...</summary>

* _2026-08-15 07:22:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18819/pull-kubevirt-e2e-kind-1.35-vgpu/2088526644472975360#1:build-log.txt%3A444)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (3x / 75.00%)


#### external (3x / 100.00%)

<details>
<summary> download failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-08-17 13:46:57 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18826/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2089348140795695104#1:build-log.txt%3A248)
<details>
<summary>all...</summary>

* _2026-08-17 13:46:57 &#43;0000 UTC_: <code>make: *** [Makefile:39: bazel-build-images] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18826/pull-kubevirt-e2e-kind-1.36-sig-compute-arm64/2089348140795695104#1:build-log.txt%3A248)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-08-15 07:22:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18819/pull-kubevirt-e2e-kind-1.35-vgpu/2088526644472975360#1:build-log.txt%3A444)
<details>
<summary>all...</summary>

* _2026-08-15 07:22:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18819/pull-kubevirt-e2e-kind-1.35-vgpu/2088526644472975360#1:build-log.txt%3A444)

</details>

<hr/>
</details>
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 33.33%) </summary>

<hr/>

**1x**: _2026-08-17 06:54:40 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18735/pull-kubevirt-e2e-kind-1.36-sev/2089244337438199808#1:build-log.txt%3A1685)
<details>
<summary>all...</summary>

* _2026-08-17 06:54:40 &#43;0000 UTC_: <code>make: *** [Makefile:189: cluster-sync] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18735/pull-kubevirt-e2e-kind-1.36-sev/2089244337438199808#1:build-log.txt%3A1685)

</details>

<hr/>
</details>

### sig-network (1x / 25.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-12 16:38:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18172/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2087556212068454400#1:build-log.txt%3A2630)
<details>
<summary>all...</summary>

* _2026-08-12 16:38:47 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18172/pull-kubevirt-e2e-k8s-1.36-sig-network-sriov-emulated/2087556212068454400#1:build-log.txt%3A2630)

</details>

<hr/>
</details>

Last updated: 2026-08-19 12:12:55
