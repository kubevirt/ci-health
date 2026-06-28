



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-06-25 (2x / 22.22%)


#### pr-build (1x / 50.00%)

<details>
<summary> panic detected in test output (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> podman container removal timeout (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)

</details>

<hr/>
</details>

### 2026-06-24 (4x / 44.44%)


#### external (4x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (2x / 50.00%) </summary>

<hr/>

**2x**: _2026-06-24 15:02:05 &#43;0000 UTC_: <code>15:09:33: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-storage/2069797359105085440#1:build-log.txt%3A1151)
<details>
<summary>all...</summary>

* _2026-06-24 15:02:13 &#43;0000 UTC_: <code>15:08:09: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-network/2069797358933118976#1:build-log.txt%3A1190)

* _2026-06-24 15:02:05 &#43;0000 UTC_: <code>15:09:33: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-storage/2069797359105085440#1:build-log.txt%3A1151)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache connection timeout (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-06-24 15:04:26 &#43;0000 UTC_: <code>21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-operator/2069797359335772160#1:build-log.txt%3A1069)
<details>
<summary>all...</summary>

* _2026-06-24 15:04:26 &#43;0000 UTC_: <code>21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-operator/2069797359335772160#1:build-log.txt%3A1069)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-06-24 15:01:21 &#43;0000 UTC_: <code>15:06:56: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs/BUILD.bazel:3:11 Validating nogo output for //vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs:go_default_library failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2069797358828261376#1:build-log.txt%3A1255)
<details>
<summary>all...</summary>

* _2026-06-24 15:01:21 &#43;0000 UTC_: <code>15:06:56: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs/BUILD.bazel:3:11 Validating nogo output for //vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs:go_default_library failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2069797358828261376#1:build-log.txt%3A1255)

</details>

<hr/>
</details>

### 2026-06-23 (1x / 11.11%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-23 10:49:37 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18133/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2069371909820125184)
<details>
<summary>all...</summary>

* _2026-06-23 10:49:37 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18133/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2069371909820125184)

</details>

<hr/>
</details>

### 2026-06-22 (1x / 11.11%)


#### external (1x / 100.00%)

<details>
<summary> podman container removal timeout (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-22 13:43:52 &#43;0000 UTC_: <code>14:25:50: Error: cannot remove container c87a5eaed63d75196f4c26360200079808f2a288908100e56da4e8c477b994fa as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18211/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2069038953901592576#1:build-log.txt%3A4436)
<details>
<summary>all...</summary>

* _2026-06-22 13:43:52 &#43;0000 UTC_: <code>14:25:50: Error: cannot remove container c87a5eaed63d75196f4c26360200079808f2a288908100e56da4e8c477b994fa as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18211/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2069038953901592576#1:build-log.txt%3A4436)

</details>

<hr/>
</details>

### 2026-06-21 (1x / 11.11%)


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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (7x / 77.78%)

<details>
<summary> bazel remote cache blob fetch failure (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-06-24 15:02:05 &#43;0000 UTC_: <code>15:09:33: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-storage/2069797359105085440#1:build-log.txt%3A1151)
<details>
<summary>all...</summary>

* _2026-06-24 15:02:13 &#43;0000 UTC_: <code>15:08:09: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-network/2069797358933118976#1:build-log.txt%3A1190)
<details><summary>context</summary>
<pre>15:08:07: INFO: Build option --define has changed, discarding analysis cache.
15:08:08: INFO: Analyzed 7 targets (20 packages loaded, 14806 targets configured).
15:08:08: INFO: Found 7 targets...
15:08:09: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
15:08:09: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: c90df6922f119c9a6419c01352467ff22dfac26652a3d0cbb125b4e70c849115/21088 for bazel-out/k8-fastbuild/bin/vendor/github.com/onsi/ginkgo/v2/formatter/go_default_library.x
15:08:09: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4e17963ff53f78aaa1f723ff42cfb6c454cb1d9c70f229fa24826d57dd3c0410/181224 for bazel-out/k8-fastbuild/bin/vendor/github.com/onsi/ginkgo/v2/types/go_default_library.x
15:08:09: INFO: Elapsed time: 1.710s, Critical Path: 0.65s</pre>
</details>


* _2026-06-24 15:02:05 &#43;0000 UTC_: <code>15:09:33: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-storage/2069797359105085440#1:build-log.txt%3A1151)
<details><summary>context</summary>
<pre>15:09:31: INFO: Build option --define has changed, discarding analysis cache.
15:09:32: INFO: Analyzed 7 targets (20 packages loaded, 14806 targets configured).
15:09:32: INFO: Found 7 targets...
15:09:33: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:
15:09:33: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: c90df6922f119c9a6419c01352467ff22dfac26652a3d0cbb125b4e70c849115/21088 for bazel-out/k8-fastbuild/bin/vendor/github.com/onsi/ginkgo/v2/formatter/go_default_library.x
15:09:33: com.google.devtools.build.lib.remote.common.CacheNotFoundException: Missing digest: 4e17963ff53f78aaa1f723ff42cfb6c454cb1d9c70f229fa24826d57dd3c0410/181224 for bazel-out/k8-fastbuild/bin/vendor/github.com/onsi/ginkgo/v2/types/go_default_library.x
15:09:33: INFO: Elapsed time: 2.684s, Critical Path: 0.55s</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (2x / 22.22%) </summary>

<hr/>

**2x**: _2026-06-22 13:43:52 &#43;0000 UTC_: <code>14:25:50: Error: cannot remove container c87a5eaed63d75196f4c26360200079808f2a288908100e56da4e8c477b994fa as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18211/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2069038953901592576#1:build-log.txt%3A4436)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)
<details><summary>context</summary>
<pre>07:06:26: Command Output: time=&#34;2026-06-25T07:06:19Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container kind-1.35-control-plane in 10 seconds, resorting to SIGKILL&#34;
07:06:26: time=&#34;2026-06-25T07:06:21Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container kind-1.35-worker in 10 seconds, resorting to SIGKILL&#34;
07:06:26: Error: cannot remove container bef9fdc1edf782f21c2ad7d5d7c703d5bd4e05e110d7c0928819fcb4babcb18d as it could not be stopped: given PID did not die within timeout
07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout
07:06:26:
07:06:26: Stack Trace:
07:06:26: sigs.k8s.io/kind/pkg/errors.WithStack</pre>
</details>


* _2026-06-22 13:43:52 &#43;0000 UTC_: <code>14:25:50: Error: cannot remove container c87a5eaed63d75196f4c26360200079808f2a288908100e56da4e8c477b994fa as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18211/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2069038953901592576#1:build-log.txt%3A4436)
<details><summary>context</summary>
<pre>14:25:50: Command Output: time=&#34;2026-06-22T14:25:19Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container kind-1.34-control-plane in 10 seconds, resorting to SIGKILL&#34;
14:25:50: time=&#34;2026-06-22T14:25:21Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container kind-1.34-worker in 10 seconds, resorting to SIGKILL&#34;
14:25:50: Error: cannot remove container 2e214a939c9010e27b066e79dde0e3e0d9fcb283538430d19a44b7f30a67c317 as it could not be stopped: given PID did not die within timeout
14:25:50: Error: cannot remove container c87a5eaed63d75196f4c26360200079808f2a288908100e56da4e8c477b994fa as it could not be stopped: given PID did not die within timeout
&#43; exit 2
&#43; EXIT_VALUE=2
&#43; set &#43;o xtrace</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-24 15:01:21 &#43;0000 UTC_: <code>15:06:56: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs/BUILD.bazel:3:11 Validating nogo output for //vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs:go_default_library failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2069797358828261376#1:build-log.txt%3A1255)
<details>
<summary>all...</summary>

* _2026-06-24 15:01:21 &#43;0000 UTC_: <code>15:06:56: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs/BUILD.bazel:3:11 Validating nogo output for //vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs:go_default_library failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2069797358828261376#1:build-log.txt%3A1255)
<details><summary>context</summary>
<pre>15:06:56: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs/BUILD.bazel:3:11: Running nogo on //vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs:go_default_library failed: Exec failed due to IOException: 2 errors during bulk transfer:
15:06:56: java.net.UnknownHostException: bazel-cache.kubevirt-prow
15:06:56: java.net.UnknownHostException: bazel-cache.kubevirt-prow
15:06:56: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs/BUILD.bazel:3:11 Validating nogo output for //vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs:go_default_library failed: Exec failed due to IOException: 2 errors during bulk transfer:
15:06:56: java.net.UnknownHostException: bazel-cache.kubevirt-prow
15:06:56: java.net.UnknownHostException: bazel-cache.kubevirt-prow
15:06:56: INFO: Elapsed time: 1.524s, Critical Path: 0.29s</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> bazel remote cache connection timeout (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-24 15:04:26 &#43;0000 UTC_: <code>21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-operator/2069797359335772160#1:build-log.txt%3A1069)
<details>
<summary>all...</summary>

* _2026-06-24 15:04:26 &#43;0000 UTC_: <code>21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-operator/2069797359335772160#1:build-log.txt%3A1069)
<details><summary>context</summary>
<pre>21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080
21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080
21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080
21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080
21:56:35: INFO: From GoCompilePkg vendor/libguestfs.org/libnbd/go_default_library.a:
21:56:35: In file included from closures.go:32,
21:56:35:                  from /tmp/rules_go_work-961419814/cgo/kubevirt.io/kubevirt/vendor/libguestfs.org/libnbd/_cgo_export.c:4:</pre>
</details>


</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 11.11%) </summary>

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

### needs-investigation (1x / 11.11%)

<details>
<summary> no error snippets (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-23 10:49:37 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18133/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2069371909820125184)
<details>
<summary>all...</summary>

* _2026-06-23 10:49:37 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18133/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2069371909820125184)

</details>

<hr/>
</details>

### pr-build (1x / 11.11%)

<details>
<summary> panic detected in test output (1x / 11.11%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)
<details><summary>context</summary>
<pre>&#43; set &#43;x

================================
ERROR: Found panic in test output
Files:
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816/artifacts/pods/0_kubevirt_virt-handler-dqq6m-virt-handler.log
https://storage.googleapis.com/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816/artifacts/pods/0_kubevirt_virt-handler-dqq6m-virt-handler_previous.log</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (8x / 88.89%)


#### external (6x / 75.00%)

<details>
<summary> bazel remote cache blob fetch failure (2x / 25.00%) </summary>

<hr/>

**2x**: _2026-06-24 15:02:05 &#43;0000 UTC_: <code>15:09:33: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-storage/2069797359105085440#1:build-log.txt%3A1151)
<details>
<summary>all...</summary>

* _2026-06-24 15:02:13 &#43;0000 UTC_: <code>15:08:09: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-network/2069797358933118976#1:build-log.txt%3A1190)

* _2026-06-24 15:02:05 &#43;0000 UTC_: <code>15:09:33: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-storage/2069797359105085440#1:build-log.txt%3A1151)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache connection timeout (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-06-24 15:04:26 &#43;0000 UTC_: <code>21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-operator/2069797359335772160#1:build-log.txt%3A1069)
<details>
<summary>all...</summary>

* _2026-06-24 15:04:26 &#43;0000 UTC_: <code>21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-operator/2069797359335772160#1:build-log.txt%3A1069)

</details>

<hr/>
</details>
<details>
<summary> podman container removal timeout (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)
<details>
<summary>all...</summary>

* _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-06-24 15:01:21 &#43;0000 UTC_: <code>15:06:56: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs/BUILD.bazel:3:11 Validating nogo output for //vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs:go_default_library failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2069797358828261376#1:build-log.txt%3A1255)
<details>
<summary>all...</summary>

* _2026-06-24 15:01:21 &#43;0000 UTC_: <code>15:06:56: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs/BUILD.bazel:3:11 Validating nogo output for //vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs:go_default_library failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2069797358828261376#1:build-log.txt%3A1255)

</details>

<hr/>
</details>

#### pr-build (1x / 12.50%)

<details>
<summary> panic detected in test output (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)

</details>

<hr/>
</details>

#### needs-investigation (1x / 12.50%)

<details>
<summary> no error snippets (1x / 12.50%) </summary>

<hr/>

**1x**: _2026-06-23 10:49:37 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18133/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2069371909820125184)
<details>
<summary>all...</summary>

* _2026-06-23 10:49:37 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18133/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2069371909820125184)

</details>

<hr/>
</details>

### release-1.7 (1x / 11.11%)


#### external (1x / 100.00%)

<details>
<summary> podman container removal timeout (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-22 13:43:52 &#43;0000 UTC_: <code>14:25:50: Error: cannot remove container c87a5eaed63d75196f4c26360200079808f2a288908100e56da4e8c477b994fa as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18211/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2069038953901592576#1:build-log.txt%3A4436)
<details>
<summary>all...</summary>

* _2026-06-22 13:43:52 &#43;0000 UTC_: <code>14:25:50: Error: cannot remove container c87a5eaed63d75196f4c26360200079808f2a288908100e56da4e8c477b994fa as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18211/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2069038953901592576#1:build-log.txt%3A4436)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (6x / 66.67%)


#### external (4x / 66.67%)

<details>
<summary> podman container removal timeout (2x / 33.33%) </summary>

<hr/>

**2x**: _2026-06-22 13:43:52 &#43;0000 UTC_: <code>14:25:50: Error: cannot remove container c87a5eaed63d75196f4c26360200079808f2a288908100e56da4e8c477b994fa as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18211/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2069038953901592576#1:build-log.txt%3A4436)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)

* _2026-06-22 13:43:52 &#43;0000 UTC_: <code>14:25:50: Error: cannot remove container c87a5eaed63d75196f4c26360200079808f2a288908100e56da4e8c477b994fa as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18211/pull-kubevirt-e2e-k8s-1.34-sig-compute-arm64-1.7/2069038953901592576#1:build-log.txt%3A4436)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache connection timeout (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-06-24 15:04:26 &#43;0000 UTC_: <code>21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-operator/2069797359335772160#1:build-log.txt%3A1069)
<details>
<summary>all...</summary>

* _2026-06-24 15:04:26 &#43;0000 UTC_: <code>21:56:35: io.netty.channel.ConnectTimeoutException: connection timed out: bazel-cache.kubevirt-prow/172.30.114.100:8080</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-operator/2069797359335772160#1:build-log.txt%3A1069)

</details>

<hr/>
</details>
<details>
<summary> container image pull failure in context (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)
<details>
<summary>all...</summary>

* _2026-06-21 22:38:31 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18202/pull-kubevirt-e2e-kind-1.35-vgpu/2068825787292717056#1:build-log.txt%3A517)

</details>

<hr/>
</details>

#### pr-build (1x / 16.67%)

<details>
<summary> panic detected in test output (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)
<details>
<summary>all...</summary>

* _2026-06-25 06:41:50 &#43;0000 UTC_: <code>ERROR: Found panic in test output</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.34-sev/2070034457527586816#1:build-log.txt%3A5211)

</details>

<hr/>
</details>

#### needs-investigation (1x / 16.67%)

<details>
<summary> no error snippets (1x / 16.67%) </summary>

<hr/>

**1x**: _2026-06-23 10:49:37 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18133/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2069371909820125184)
<details>
<summary>all...</summary>

* _2026-06-23 10:49:37 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18133/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2069371909820125184)

</details>

<hr/>
</details>

### sig-storage (1x / 11.11%)


#### external (1x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-06-24 15:02:05 &#43;0000 UTC_: <code>15:09:33: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-storage/2069797359105085440#1:build-log.txt%3A1151)
<details>
<summary>all...</summary>

* _2026-06-24 15:02:05 &#43;0000 UTC_: <code>15:09:33: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-storage/2069797359105085440#1:build-log.txt%3A1151)

</details>

<hr/>
</details>

### sig-network (2x / 22.22%)


#### external (2x / 100.00%)

<details>
<summary> bazel remote cache blob fetch failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-24 15:02:13 &#43;0000 UTC_: <code>15:08:09: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-network/2069797358933118976#1:build-log.txt%3A1190)
<details>
<summary>all...</summary>

* _2026-06-24 15:02:13 &#43;0000 UTC_: <code>15:08:09: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/command/BUILD.bazel:3:11: GoCompilePkg vendor/github.com/onsi/ginkgo/v2/ginkgo/command/go_default_library.a failed: Failed to fetch blobs because they do not exist remotely.: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.34-sig-network/2069797358933118976#1:build-log.txt%3A1190)

</details>

<hr/>
</details>
<details>
<summary> bazel remote cache IO failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-24 15:01:21 &#43;0000 UTC_: <code>15:06:56: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs/BUILD.bazel:3:11 Validating nogo output for //vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs:go_default_library failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2069797358828261376#1:build-log.txt%3A1255)
<details>
<summary>all...</summary>

* _2026-06-24 15:01:21 &#43;0000 UTC_: <code>15:06:56: ERROR: /root/go/src/kubevirt.io/kubevirt/vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs/BUILD.bazel:3:11 Validating nogo output for //vendor/github.com/onsi/ginkgo/v2/ginkgo/automaxprocs:go_default_library failed: Exec failed due to IOException: 2 errors during bulk transfer:</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18109/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2069797358828261376#1:build-log.txt%3A1255)

</details>

<hr/>
</details>

Last updated: 2026-06-28 15:45:57
