



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-06-25 (2x / 100.00%)


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

<a id="per-error-category"></a>

## per error category [⬆](#top)


### pr-build (1x / 50.00%)

<details>
<summary> panic detected in test output (1x / 50.00%) </summary>

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

### external (1x / 50.00%)

<details>
<summary> podman container removal timeout (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-06-25 06:41:41 &#43;0000 UTC_: <code>07:06:26: Error: cannot remove container 0e0afb9af71a75d9b026ee76ea2ecb8cce792a6dec62cf50a08fee527c34ab5c as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18244/pull-kubevirt-e2e-kind-1.35-sig-compute-arm64/2070034428989542400#1:build-log.txt%3A739)
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


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### main (2x / 100.00%)


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

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (2x / 100.00%)


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

Last updated: 2026-07-02 00:40:06
