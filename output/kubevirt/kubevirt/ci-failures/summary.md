



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-05-05 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-05 10:04:34 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17704/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2051603690623602688#1:build-log.txt%3A553)
<details>
<summary>all...</summary>

* _2026-05-05 10:04:34 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17704/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2051603690623602688#1:build-log.txt%3A553)

</details>

<hr/>
</details>

### 2026-05-04 (1x / 50.00%)


#### pr-build (1x / 100.00%)

<details>
<summary> go test exit FAIL (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-04 18:36:02 &#43;0000 UTC_: <code>19:48:34: FAIL</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17692/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.6/2051369950953607168#1:build-log.txt%3A81318)
<details>
<summary>all...</summary>

* _2026-05-04 18:36:02 &#43;0000 UTC_: <code>19:48:34: FAIL</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17692/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.6/2051369950953607168#1:build-log.txt%3A81318)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### pr-build (1x / 50.00%)

<details>
<summary> go test exit FAIL (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-05-04 18:36:02 &#43;0000 UTC_: <code>19:48:34: FAIL</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17692/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.6/2051369950953607168#1:build-log.txt%3A81318)
<details>
<summary>all...</summary>

* _2026-05-04 18:36:02 &#43;0000 UTC_: <code>19:48:34: FAIL</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17692/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.6/2051369950953607168#1:build-log.txt%3A81318)
<details><summary>context</summary>
<pre>19:48:34: Ran 3 of 1433 Specs in 2825.199 seconds
19:48:34: FAIL! -- 2 Passed | 1 Failed | 3 Pending | 1427 Skipped
19:48:34: --- FAIL: TestTests (2828.62s)
19:48:34: FAIL
19:48:34:
19:48:34: Ginkgo ran 1 suite in 47m12.41912815s
19:48:34:</pre>
</details>


</details>

<hr/>
</details>

### external (1x / 50.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-05-05 10:04:34 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17704/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2051603690623602688#1:build-log.txt%3A553)
<details>
<summary>all...</summary>

* _2026-05-05 10:04:34 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17704/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2051603690623602688#1:build-log.txt%3A553)
<details><summary>context</summary>
<pre>10:26:42: Trying to pull quay.io/kubevirtci/gocli:2602190951-056f8f79...
10:28:28: Error: unable to copy from source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: initializing source docker://quay.io/kubevirtci/gocli:2602190951-056f8f79: pinging container registry quay.io: Get &#34;https://quay.io/v2/&#34;: net/http: TLS handshake timeout
10:28:28: /home/prow/go/src/github.com/kubevirt/kubevirt/kubevirtci/cluster-up/up.sh: line 48: pop_var_context: head of shell_variables not a function context
make: *** [Makefile:173: cluster-up] Error 125
&#43;&#43; collect_debug_logs
&#43;&#43; local containers
&#43;&#43;&#43; determine_cri_bin</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### release-1.8 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-05 10:04:34 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17704/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2051603690623602688#1:build-log.txt%3A553)
<details>
<summary>all...</summary>

* _2026-05-05 10:04:34 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17704/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2051603690623602688#1:build-log.txt%3A553)

</details>

<hr/>
</details>

### release-1.6 (1x / 50.00%)


#### pr-build (1x / 100.00%)

<details>
<summary> go test exit FAIL (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-04 18:36:02 &#43;0000 UTC_: <code>19:48:34: FAIL</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17692/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.6/2051369950953607168#1:build-log.txt%3A81318)
<details>
<summary>all...</summary>

* _2026-05-04 18:36:02 &#43;0000 UTC_: <code>19:48:34: FAIL</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17692/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.6/2051369950953607168#1:build-log.txt%3A81318)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-05 10:04:34 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17704/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2051603690623602688#1:build-log.txt%3A553)
<details>
<summary>all...</summary>

* _2026-05-05 10:04:34 &#43;0000 UTC_: <code>make: *** [Makefile:173: cluster-up] Error 125</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17704/pull-kubevirt-e2e-k8s-1.33-sig-compute-1.8/2051603690623602688#1:build-log.txt%3A553)

</details>

<hr/>
</details>

### sig-performance (1x / 50.00%)


#### pr-build (1x / 100.00%)

<details>
<summary> go test exit FAIL (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-05-04 18:36:02 &#43;0000 UTC_: <code>19:48:34: FAIL</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17692/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.6/2051369950953607168#1:build-log.txt%3A81318)
<details>
<summary>all...</summary>

* _2026-05-04 18:36:02 &#43;0000 UTC_: <code>19:48:34: FAIL</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17692/pull-kubevirt-e2e-k8s-1.31-sig-performance-1.6/2051369950953607168#1:build-log.txt%3A81318)

</details>

<hr/>
</details>

Last updated: 2026-05-07 12:44:42
