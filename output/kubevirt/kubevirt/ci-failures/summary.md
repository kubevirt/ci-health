



<a id="top"></a>

# CI failures for kubevirt/kubevirt

- [per day](#per-day)
- [per error category](#per-error-category)
- [per branch](#per-branch)
- [per SIG](#per-sig)


<a id="per-day"></a>

## per day [⬆](#top)


### 2026-08-04 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)
<details>
<summary>all...</summary>

* _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)

</details>

<hr/>
</details>

### 2026-07-30 (1x / 50.00%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)
<details>
<summary>all...</summary>

* _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### internal (1x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)
<details>
<summary>all...</summary>

* _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
14:02:11: selecting podman as container runtime
14:02:38: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7 is being used by the following container(s): e0a020bd2bd11f60ce26e9525064ea1f4f66c9b590247ed662d259c32a9f9537, 6b2381dd92e851177e8fae6be1d2610d9fd5a53c242b4287982a4e3c899d7cc8: volume is being used
make: *** [Makefile:162: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

### external (1x / 50.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)
<details>
<summary>all...</summary>

* _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)
<details><summary>context</summary>
<pre>01:54:50: Copying config sha256:34e1e89c7cd17b52a316d25aecb8f529ae4b3de71aa90d2482c16fd5233608fd
01:54:50: Writing manifest to image destination
01:54:50: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9 is being used by the following container(s): 6a071776d5ab7e847352faa75fdfc6084d16e607f072b021e3de5c756d8e0253: volume is being used
make: *** [Makefile:177: cluster-down] Error 1
&#43; true
&#43; exit 2
&#43; EXIT_VALUE=2</pre>
</details>


</details>

<hr/>
</details>

<a id="per-branch"></a>

## per branch [⬆](#top)


### release-1.9 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> container image pull failure in context (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)
<details>
<summary>all...</summary>

* _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)

</details>

<hr/>
</details>

### release-1.7 (1x / 50.00%)


#### internal (1x / 100.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)
<details>
<summary>all...</summary>

* _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)

</details>

<hr/>
</details>

<a id="per-sig"></a>

## per SIG [⬆](#top)


### sig-compute (2x / 100.00%)


#### internal (1x / 50.00%)

<details>
<summary> make cluster lifecycle target failure (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)
<details>
<summary>all...</summary>

* _2026-07-30 13:41:27 &#43;0000 UTC_: <code>make: *** [Makefile:162: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18648/pull-kubevirt-e2e-k8s-1.32-sig-compute-1.7/2082821112458121216#1:build-log.txt%3A1056)

</details>

<hr/>
</details>

#### external (1x / 50.00%)

<details>
<summary> container image pull failure in context (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)
<details>
<summary>all...</summary>

* _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)

</details>

<hr/>
</details>

Last updated: 2026-08-06 12:40:44
