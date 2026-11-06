





<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-11-04 (1x / 14.29%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-04 07:13:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)
<details>
<summary>all...</summary>

* _2026-11-04 07:13:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)




</details>

<hr/>
</details>

### 2026-11-03 (4x / 57.14%)


#### internal (4x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (3x / 75.00%) </summary>

<hr/>

**3x**: _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)
<details>
<summary>all...</summary>

* _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4057)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:30:29: selecting podman as container runtime
12:32:15: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-network-1.7 is being used by the following container(s): 437e990e2003afcfa461cefe139c5de3abcd06e6208ef33c9e77cb53b350a614, d7746eb1230afd44770adaf9ee2a421590a3d28144b9974b1f14b2ea60f6311e: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:36:37: selecting podman as container runtime
12:36:52: Error response from daemon: cannot remove container 09207706028f593ad43cb2e4bf5ea652f5970c6efdf3a476d3f2ee015494902d as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:58:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3943)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>12:15:21: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
12:15:25: [control-plane-check] kube-scheduler is healthy after 3.842462827s
12:15:26: [control-plane-check] kube-controller-manager is healthy after 5.068241346s
12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:15:30: [control-plane-check] kube-apiserver is healthy after 9.002895106s
12:15:30: I1103 07:15:30.875756    1605 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
12:15:30: I1103 07:15:30.878733    1605 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3949)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:27:34: selecting podman as container runtime
12:27:49: Error response from daemon: cannot remove container 4d2bdd9539b64bcf49fcea3d1d98ac72c4d24a6d0e4c0180ba3525a5a560ec7c as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> taint not found during cluster setup (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)
<details>
<summary>all...</summary>

* _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found
12:24:20: &#43;&#43;&#43; grep -i worker
12:24:20: &#43;&#43;&#43; _get_nodes
12:24:20: &#43;&#43;&#43; awk &#39;{print $1}&#39;</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1541)

<details><summary>context</summary>
<pre>12:29:21: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
12:29:21: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
12:29:49: &#43; rm -f /tmp/kubevirt.deploy.oDzW
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1550)

<details><summary>context</summary>
<pre>12:29:50: using podman due to KIND_EXPERIMENTAL_PROVIDER
12:29:50: enabling experimental podman provider
12:29:50: Deleting cluster &#34;vgpu&#34; ...
12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1568)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1569)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1570)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1571)

<details><summary>context</summary>
<pre>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>

### 2026-11-02 (2x / 28.57%)


#### internal (1x / 50.00%)

<details>
<summary> configured node not found during provisioning (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A810)

<details><summary>context</summary>
<pre>12:19:44: [control-plane-check] kube-controller-manager is healthy after 8.488191ms
12:19:44: [control-plane-check] kube-scheduler is healthy after 10.387052ms
12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A812)

<details><summary>context</summary>
<pre>12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:47.090991    1901 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
12:19:46: [control-plane-check] kube-apiserver is healthy after 2.005710556s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1156)

<details><summary>context</summary>
<pre>12:20:04: 
12:20:04: &#43; wait 69
12:20:53: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:20:53: selecting podman as container runtime
12:20:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): 6a6a2973d141a1448dc31509ecfd824ede38d3435e163c24d8b9cba7915f160b: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> container image pull failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)

<details><summary>context</summary>
<pre>09:16:19: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
09:16:19: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### internal (5x / 71.43%)

<details>
<summary> make cluster lifecycle target failure (3x / 42.86%) </summary>

<hr/>

**3x**: _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)
<details>
<summary>all...</summary>

* _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)
<details><summary>context</summary>
<pre>12:30:01: 2026/11/03 12:30:01 existing blob: sha256:05daf2a1d5cd22bf20bf774e8599d99ee37f8edccf2588aaa86f21aaec88c88d
12:30:02: 2026/11/03 12:30:02 localhost:36275/kubevirt/kv-network-passt-binding-cni:devel: digest: sha256:0d6e045da432423816b2bc90af03985fc7c0432ef9ca321aa853181e18ba8a66 size: 701
12:30:17: &#43; rm -f /tmp/kubevirt.deploy.jwbQ
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4057)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:30:29: selecting podman as container runtime
12:32:15: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-network-1.7 is being used by the following container(s): 437e990e2003afcfa461cefe139c5de3abcd06e6208ef33c9e77cb53b350a614, d7746eb1230afd44770adaf9ee2a421590a3d28144b9974b1f14b2ea60f6311e: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)
<details><summary>context</summary>
<pre>12:35:36: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
12:35:36: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
12:36:36: &#43; rm -f /tmp/kubevirt.deploy.1MC8
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:36:37: selecting podman as container runtime
12:36:52: Error response from daemon: cannot remove container 09207706028f593ad43cb2e4bf5ea652f5970c6efdf3a476d3f2ee015494902d as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:58:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3943)
<details><summary>context</summary>
<pre>12:27:25: 2026/11/03 12:27:25 existing blob: sha256:cc7f351ef2b73318cc7d1d133ee9c9dc6913543797cd8f8b294edade3700892e
12:27:25: 2026/11/03 12:27:25 localhost:46415/kubevirt/kv-network-passt-binding-cni:devel: digest: sha256:0d6e045da432423816b2bc90af03985fc7c0432ef9ca321aa853181e18ba8a66 size: 701
12:27:34: &#43; rm -f /tmp/kubevirt.deploy.hT3Q
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>12:15:21: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
12:15:25: [control-plane-check] kube-scheduler is healthy after 3.842462827s
12:15:26: [control-plane-check] kube-controller-manager is healthy after 5.068241346s
12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:15:30: [control-plane-check] kube-apiserver is healthy after 9.002895106s
12:15:30: I1103 07:15:30.875756    1605 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
12:15:30: I1103 07:15:30.878733    1605 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3949)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:27:34: selecting podman as container runtime
12:27:49: Error response from daemon: cannot remove container 4d2bdd9539b64bcf49fcea3d1d98ac72c4d24a6d0e4c0180ba3525a5a560ec7c as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> configured node not found during provisioning (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)
<details><summary>context</summary>
<pre>12:19:49: Warning: would violate PodSecurity &#34;restricted:latest&#34;: host namespaces (hostNetwork=true), allowPrivilegeEscalation != false (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.allowPrivilegeEscalation=false), unrestricted capabilities (containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.capabilities.drop=[&#34;ALL&#34;]; container &#34;kube-flannel&#34; must not include &#34;NET_ADMIN&#34;, &#34;NET_RAW&#34; in securityContext.capabilities.add), restricted volume types (volumes &#34;run&#34;, &#34;cni-plugin&#34;, &#34;cni&#34;, &#34;xtables-lock&#34; use restricted volume type &#34;hostPath&#34;), runAsNonRoot != true (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.runAsNonRoot=true), seccompProfile (pod or containers &#34;install-cni-plugin&#34;, &#34;install-cni&#34;, &#34;kube-flannel&#34; must set securityContext.seccompProfile.type to &#34;RuntimeDefault&#34; or &#34;Localhost&#34;)
12:19:49: daemonset.apps/kube-flannel-ds created
12:19:49: time=&#34;2026-11-02T12:19:49Z&#34; level=info msg=&#34;[node 1]: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-&#34;
12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found
12:19:49:
12:19:49: ===== 21fb777d537fc8ba8050063cc17395841f375e1ae321c53c470f287880e4f2da ====
12:19:49: &#43; NUM_NODES=2</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A810)

<details><summary>context</summary>
<pre>12:19:44: [control-plane-check] kube-controller-manager is healthy after 8.488191ms
12:19:44: [control-plane-check] kube-scheduler is healthy after 10.387052ms
12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A812)

<details><summary>context</summary>
<pre>12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:47.090991    1901 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
12:19:46: [control-plane-check] kube-apiserver is healthy after 2.005710556s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1156)

<details><summary>context</summary>
<pre>12:20:04: 
12:20:04: &#43; wait 69
12:20:53: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:20:53: selecting podman as container runtime
12:20:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): 6a6a2973d141a1448dc31509ecfd824ede38d3435e163c24d8b9cba7915f160b: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> taint not found during cluster setup (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)
<details>
<summary>all...</summary>

* _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)
<details><summary>context</summary>
<pre>12:24:20: &#43;&#43; for node in ${master_nodes[@]}
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/master:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/master:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found
12:24:20: &#43;&#43;&#43; grep -i worker
12:24:20: &#43;&#43;&#43; _get_nodes
12:24:20: &#43;&#43;&#43; awk &#39;{print $1}&#39;</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1541)

<details><summary>context</summary>
<pre>12:29:21: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
12:29:21: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
12:29:49: &#43; rm -f /tmp/kubevirt.deploy.oDzW
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1550)

<details><summary>context</summary>
<pre>12:29:50: using podman due to KIND_EXPERIMENTAL_PROVIDER
12:29:50: enabling experimental podman provider
12:29:50: Deleting cluster &#34;vgpu&#34; ...
12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1568)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1569)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1570)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1571)

<details><summary>context</summary>
<pre>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>

### external (1x / 14.29%)

<details>
<summary> container image pull failure (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)
<details><summary>context</summary>
<pre>09:16:19: Getting image source signatures
09:16:19: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
09:16:19: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers</pre>
</details>



<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)

<details><summary>context</summary>
<pre>09:16:19: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
09:16:19: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

### needs-investigation (1x / 14.29%)

<details>
<summary> no error snippets (1x / 14.29%) </summary>

<hr/>

**1x**: _2026-11-04 07:13:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)
<details>
<summary>all...</summary>

* _2026-11-04 07:13:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)




</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### release-1.8 (1x / 14.29%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-04 07:13:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)
<details>
<summary>all...</summary>

* _2026-11-04 07:13:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)




</details>

<hr/>
</details>

### release-1.7 (4x / 57.14%)


#### internal (4x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (3x / 75.00%) </summary>

<hr/>

**3x**: _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)
<details>
<summary>all...</summary>

* _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4057)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:30:29: selecting podman as container runtime
12:32:15: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-network-1.7 is being used by the following container(s): 437e990e2003afcfa461cefe139c5de3abcd06e6208ef33c9e77cb53b350a614, d7746eb1230afd44770adaf9ee2a421590a3d28144b9974b1f14b2ea60f6311e: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:36:37: selecting podman as container runtime
12:36:52: Error response from daemon: cannot remove container 09207706028f593ad43cb2e4bf5ea652f5970c6efdf3a476d3f2ee015494902d as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



* _2026-11-03 11:58:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3943)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>12:15:21: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
12:15:25: [control-plane-check] kube-scheduler is healthy after 3.842462827s
12:15:26: [control-plane-check] kube-controller-manager is healthy after 5.068241346s
12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:15:30: [control-plane-check] kube-apiserver is healthy after 9.002895106s
12:15:30: I1103 07:15:30.875756    1605 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
12:15:30: I1103 07:15:30.878733    1605 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3949)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:27:34: selecting podman as container runtime
12:27:49: Error response from daemon: cannot remove container 4d2bdd9539b64bcf49fcea3d1d98ac72c4d24a6d0e4c0180ba3525a5a560ec7c as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> taint not found during cluster setup (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)
<details>
<summary>all...</summary>

* _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found
12:24:20: &#43;&#43;&#43; grep -i worker
12:24:20: &#43;&#43;&#43; _get_nodes
12:24:20: &#43;&#43;&#43; awk &#39;{print $1}&#39;</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1541)

<details><summary>context</summary>
<pre>12:29:21: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
12:29:21: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
12:29:49: &#43; rm -f /tmp/kubevirt.deploy.oDzW
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1550)

<details><summary>context</summary>
<pre>12:29:50: using podman due to KIND_EXPERIMENTAL_PROVIDER
12:29:50: enabling experimental podman provider
12:29:50: Deleting cluster &#34;vgpu&#34; ...
12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1568)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1569)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1570)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1571)

<details><summary>context</summary>
<pre>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>

### main (2x / 28.57%)


#### internal (1x / 50.00%)

<details>
<summary> configured node not found during provisioning (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A810)

<details><summary>context</summary>
<pre>12:19:44: [control-plane-check] kube-controller-manager is healthy after 8.488191ms
12:19:44: [control-plane-check] kube-scheduler is healthy after 10.387052ms
12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A812)

<details><summary>context</summary>
<pre>12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:47.090991    1901 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
12:19:46: [control-plane-check] kube-apiserver is healthy after 2.005710556s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1156)

<details><summary>context</summary>
<pre>12:20:04: 
12:20:04: &#43; wait 69
12:20:53: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:20:53: selecting podman as container runtime
12:20:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): 6a6a2973d141a1448dc31509ecfd824ede38d3435e163c24d8b9cba7915f160b: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> container image pull failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)

<details><summary>context</summary>
<pre>09:16:19: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
09:16:19: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-monitoring (1x / 14.29%)


#### needs-investigation (1x / 100.00%)

<details>
<summary> no error snippets (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-04 07:13:51 &#43;0000 UTC_:  _(no match in error message grep)_ [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)
<details>
<summary>all...</summary>

* _2026-11-04 07:13:51 &#43;0000 UTC_: [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19748/pull-kubevirt-e2e-k8s-1.34-sig-monitoring-1.8/2117877790450651136)




</details>

<hr/>
</details>

### sig-network (1x / 14.29%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)
<details>
<summary>all...</summary>

* _2026-11-03 12:09:09 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4051)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.33-sig-network-1.7/2117587001577312256#1:build-log.txt%3A4057)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:30:29: selecting podman as container runtime
12:32:15: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.33-sig-network-1.7 is being used by the following container(s): 437e990e2003afcfa461cefe139c5de3abcd06e6208ef33c9e77cb53b350a614, d7746eb1230afd44770adaf9ee2a421590a3d28144b9974b1f14b2ea60f6311e: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

### sig-storage (1x / 14.29%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)
<details>
<summary>all...</summary>

* _2026-11-03 11:59:44 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1217)


<details>
<summary>secondary evidence</summary>


* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.32-sig-storage-1.7/2117586999056535552#1:build-log.txt%3A1223)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:36:37: selecting podman as container runtime
12:36:52: Error response from daemon: cannot remove container 09207706028f593ad43cb2e4bf5ea652f5970c6efdf3a476d3f2ee015494902d as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>

### sig-compute (4x / 57.14%)


#### internal (3x / 75.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-03 11:58:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3943)
<details>
<summary>all...</summary>

* _2026-11-03 11:58:41 &#43;0000 UTC_: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3943)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A798)

<details><summary>context</summary>
<pre>12:15:21: [control-plane-check] Checking kube-scheduler at https://127.0.0.1:10259/livez
12:15:25: [control-plane-check] kube-scheduler is healthy after 3.842462827s
12:15:26: [control-plane-check] kube-controller-manager is healthy after 5.068241346s
12:15:30: I1103 07:15:30.373979    1605 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:15:30: [control-plane-check] kube-apiserver is healthy after 9.002895106s
12:15:30: I1103 07:15:30.875756    1605 kubeconfig.go:657] ensuring that the ClusterRoleBinding for the kubeadm:cluster-admins Group exists
12:15:30: I1103 07:15:30.878733    1605 kubeconfig.go:730] creating the ClusterRoleBinding for the kubeadm:cluster-admins Group by using super-admin.conf</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-k8s-1.34-windows2016-1.7/2117586987438313472#1:build-log.txt%3A3949)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:27:34: selecting podman as container runtime
12:27:49: Error response from daemon: cannot remove container 4d2bdd9539b64bcf49fcea3d1d98ac72c4d24a6d0e4c0180ba3525a5a560ec7c as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> configured node not found during provisioning (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)
<details>
<summary>all...</summary>

* _2026-11-02 12:10:13 &#43;0000 UTC_: <code>12:19:49: Error from server (NotFound): nodes &#34;node01&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A875)


<details>
<summary>secondary evidence</summary>


* <strong>diagnostic</strong>: <code>12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A810)

<details><summary>context</summary>
<pre>12:19:44: [control-plane-check] kube-controller-manager is healthy after 8.488191ms
12:19:44: [control-plane-check] kube-scheduler is healthy after 10.387052ms
12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding</pre>
</details>



* <strong>diagnostic</strong>: <code>12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A812)

<details><summary>context</summary>
<pre>12:19:45: I1102 07:19:45.589042    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: forbidden: User &#34;kubernetes-admin&#34; cannot get path &#34;/livez&#34;: RBAC: clusterrole.rbac.authorization.k8s.io &#34;cluster-admin&#34; not found
12:19:45: I1102 07:19:46.088210    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:45: I1102 07:19:46.088280    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:46.586901    1901 request.go:1500] &#34;Body was not decodable (unable to check for Status)&#34; err=&#34;couldn&#39;t get version/kind; json parse error: json: cannot unmarshal array into Go value of type struct { APIVersion string \&#34;json:\\\&#34;apiVersion,omitempty\\\&#34;\&#34;; Kind string \&#34;json:\\\&#34;kind,omitempty\\\&#34;\&#34; }&#34;
12:19:46: I1102 07:19:46.586947    1901 wait.go:283] kube-apiserver check failed at https://192.168.66.101:6443/livez: an error on the server (&#34;[&#43;]ping ok\n[&#43;]log ok\n[&#43;]loopback-serving-certificate ok\n[&#43;]etcd ok\n[&#43;]poststarthook/start-apiserver-admission-initializer ok\n[&#43;]poststarthook/generic-apiserver-start-informers ok\n[&#43;]poststarthook/priority-and-fairness-config-consumer ok\n[&#43;]poststarthook/priority-and-fairness-filter ok\n[&#43;]poststarthook/storage-object-count-tracker-hook ok\n[&#43;]poststarthook/start-apiextensions-informers ok\n[&#43;]poststarthook/start-apiextensions-controllers ok\n[&#43;]poststarthook/crd-informer-synced ok\n[&#43;]poststarthook/start-system-namespaces-controller ok\n[&#43;]poststarthook/peer-endpoint-reconciler-controller ok\n[&#43;]poststarthook/start-cluster-authentication-info-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-controller ok\n[&#43;]poststarthook/start-kube-apiserver-identity-lease-garbage-collector ok\n[&#43;]poststarthook/storage-readiness ok\n[&#43;]poststarthook/start-legacy-token-tracking-controller ok\n[&#43;]poststarthook/start-service-ip-repair-controllers ok\n[-]poststarthook/rbac/bootstrap-roles failed: reason withheld\n[&#43;]poststarthook/scheduling/bootstrap-system-priority-classes ok\n[&#43;]poststarthook/priority-and-fairness-config-producer ok\n[&#43;]poststarthook/bootstrap-controller ok\n[&#43;]poststarthook/start-kubernetes-service-cidr-controller ok\n[&#43;]poststarthook/aggregator-reload-proxy-client-cert ok\n[&#43;]poststarthook/start-kube-aggregator-informers ok\n[&#43;]poststarthook/apiservice-status-local-available-controller ok\n[&#43;]poststarthook/apiservice-status-remote-available-controller ok\n[&#43;]poststarthook/apiservice-registration-controller ok\n[&#43;]poststarthook/apiservice-discovery-controller ok\n[&#43;]poststarthook/kube-apiserver-autoregistration ok\n[&#43;]autoregister-completion ok\n[&#43;]poststarthook/apiservice-openapi-controller ok\n[&#43;]poststarthook/apiservice-openapiv3-controller ok\nlivez check failed&#34;) has prevented the request from succeeding
12:19:46: I1102 07:19:47.090991    1901 uploadconfig.go:109] [upload-config] Uploading the kubeadm ClusterConfiguration to a ConfigMap
12:19:46: [control-plane-check] kube-apiserver is healthy after 2.005710556s</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1156)

<details><summary>context</summary>
<pre>12:20:04: 
12:20:04: &#43; wait 69
12:20:53: error executing kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-: failed to execute command: kubectl --kubeconfig=/etc/kubernetes/admin.conf taint nodes node01 node-role.kubernetes.io/control-plane:NoSchedule-
make: *** [Makefile:179: cluster-up] Error 1
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:182: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19671/pull-kubevirt-e2e-k8s-1.36-windows2016/2117227317414596608#1:build-log.txt%3A1184)

<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
12:20:53: selecting podman as container runtime
12:20:54: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-windows2016 is being used by the following container(s): 6a6a2973d141a1448dc31509ecfd824ede38d3435e163c24d8b9cba7915f160b: volume is being used
make: *** [Makefile:182: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>



</details>

<hr/>
</details>
<details>
<summary> taint not found during cluster setup (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)
<details>
<summary>all...</summary>

* _2026-11-03 11:58:13 &#43;0000 UTC_: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1231)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1234)

<details><summary>context</summary>
<pre>12:24:20: error: taint &#34;node-role.kubernetes.io/master:NoSchedule&#34; not found
12:24:20: &#43;&#43; _kubectl taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: &#43;&#43; /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubectl --kubeconfig=/home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/_ci-configs/kind-1.33-vgpu/.kubeconfig taint nodes vgpu-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
12:24:20: error: taint &#34;node-role.kubernetes.io/control-plane:NoSchedule&#34; not found
12:24:20: &#43;&#43;&#43; grep -i worker
12:24:20: &#43;&#43;&#43; _get_nodes
12:24:20: &#43;&#43;&#43; awk &#39;{print $1}&#39;</pre>
</details>



* <strong>secondary</strong>: <code>make: *** [Makefile:174: cluster-sync] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1541)

<details><summary>context</summary>
<pre>12:29:21: INFO: 2 processes: 1 internal, 1 processwrapper-sandbox.
12:29:21: INFO: Running command line: bazel-bin/example-guest-agent-copier /root/go/src/kubevirt.io/kubevirt/_out/cmd/example-guest-agent/example-guest-agent
12:29:49: &#43; rm -f /tmp/kubevirt.deploy.oDzW
make: *** [Makefile:174: cluster-sync] Error 125
&#43; ret=2
&#43; make cluster-down
./kubevirtci/cluster-up/down.sh</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1550)

<details><summary>context</summary>
<pre>12:29:50: using podman due to KIND_EXPERIMENTAL_PROVIDER
12:29:50: enabling experimental podman provider
12:29:50: Deleting cluster &#34;vgpu&#34; ...
12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1553)

<details><summary>context</summary>
<pre>12:30:05: ERROR: failed to delete cluster &#34;vgpu&#34;: failed to delete nodes: command &#34;podman rm -f -v vgpu-control-plane&#34; failed with error: exit status 125
12:30:05:
12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125</pre>
</details>



* <strong>cleanup</strong>: <code>12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1555)

<details><summary>context</summary>
<pre>12:30:05: Command Output: time=&#34;2026-11-03T12:30:00Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2</pre>
</details>



* <strong>cleanup</strong>: <code>make: *** [Makefile:162: cluster-down] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1556)

<details><summary>context</summary>
<pre>12:30:05: Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
12:30:16: time=&#34;2026-11-03T12:30:16Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
12:30:21: Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
make: *** [Makefile:162: cluster-down] Error 125
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1568)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container vgpu-registry in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1569)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal SIGTERM failed to stop container kind_ardinghelli in 10 seconds, resorting to SIGKILL&#34;
time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1570)

<details><summary>context</summary>
<pre>time=&#34;2026-11-03T12:30:37Z&#34; level=warning msg=&#34;StopSignal (37) failed to stop container vgpu-control-plane in 10 seconds, resorting to SIGKILL&#34;
Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================</pre>
</details>



* <strong>cleanup</strong>: <code>Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19679/pull-kubevirt-e2e-kind-1.33-vgpu-1.7/2117586987538976768#1:build-log.txt%3A1571)

<details><summary>context</summary>
<pre>Error: cannot remove container 4f7211fdcc8ed1b910923a3924940fcc458d25a6fed2b02e52d1e2cbd23ee13a as it could not be stopped: given PID did not die within timeout
Error: cannot remove container a36ad1c88a4eea2b70150f21f545f73a12677acfe04c73d9438c57cd48a9187e as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 93db881196418c335d7dddf3e68a49081e5edabb7e2fbb1b5951a47e2bb44d32 as it could not be stopped: given PID did not die within timeout
Error: cannot remove container 72eafda1844867523f5b03b3475882bf0503b5e2a40a687ee167e0079b746163 as it could not be stopped: given PID did not die within timeout
/usr/local/bin/runner.sh: line 50: wait: pid 1201 is not a child of this shell
================================================================================
Done cleaning up after podman in container.</pre>
</details>


</details>



</details>

<hr/>
</details>

#### external (1x / 25.00%)

<details>
<summary> container image pull failure (1x / 25.00%) </summary>

<hr/>

**1x**: _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)
<details>
<summary>all...</summary>

* _2026-11-02 09:09:03 &#43;0000 UTC_: <code>09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A433)


<details>
<summary>secondary evidence</summary>


* <strong>secondary</strong>: <code>make: *** [Makefile:179: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/19158/pull-kubevirt-e2e-kind-1.37-vgpu/2117182035628920832#1:build-log.txt%3A434)

<details><summary>context</summary>
<pre>09:16:19: Copying blob sha256:7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a
09:16:19: Copying blob sha256:5c1b9e8d7bf7b758fa84807a6bce45e4af333e1ddd566b5972550b6fcfbed9b8
09:16:29: Error: unable to copy from source docker://quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1: writing blob: storing blob to file &#34;/var/tmp/container_images_storage1952500010/1&#34;: happened during read: unexpected EOF (while reconnecting: Get &#34;https://cdn01.quay.io/quayio-production-s3/sha256/7b/7be39e07f3a362cb350d0d0b5a554e29829fce73909cb105a3eafeaa5677068a?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIATAAF2YHTGR23ZTE6%2F20261102%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20261102T091622Z&amp;X-Amz-Expires=600&amp;X-Amz-SignedHeaders=host&amp;X-Amz-Signature=f26a3c01ef9ab733f4ebeedea9783fc7ddb2e2abced3b97fc20cb4612b4846e2&amp;region=us-east-1&amp;namespace=phoracek&amp;repo_name=lspci&amp;akamai_signature=exp=1793611882~hmac=2cb0f9c8a45db3c5098bfeb2133eeb14a79ef7b10b4f00e01d28605c9c2ee013&#34;: EOF)
make: *** [Makefile:179: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>



</details>

<hr/>
</details>

Last updated: 2026-11-06 00:17:11
