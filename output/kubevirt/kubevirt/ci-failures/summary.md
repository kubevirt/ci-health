



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

### 2026-08-03 (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-03 09:30:59 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18291/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2084206350073597952#1:build-log.txt%3A1821)
<details>
<summary>all...</summary>

* _2026-08-03 09:30:59 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18291/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2084206350073597952#1:build-log.txt%3A1821)

</details>

<hr/>
</details>

<a id="per-error-category"></a>

## per error category [⬆](#top)


### external (2x / 100.00%)

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
<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 50.00%) </summary>

<hr/>

**1x**: _2026-08-03 09:30:59 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18291/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2084206350073597952#1:build-log.txt%3A1821)
<details>
<summary>all...</summary>

* _2026-08-03 09:30:59 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18291/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2084206350073597952#1:build-log.txt%3A1821)
<details><summary>context</summary>
<pre>./kubevirtci/cluster-up/down.sh
09:35:23: selecting podman as container runtime
09:35:58: Error response from daemon: volume pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network is being used by the following container(s): 474f3c2a6416a05d693eebc77e9be1b61b355d21948b0dd7999211c5b1210c5e: volume is being used
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

### main (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-03 09:30:59 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18291/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2084206350073597952#1:build-log.txt%3A1821)
<details>
<summary>all...</summary>

* _2026-08-03 09:30:59 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18291/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2084206350073597952#1:build-log.txt%3A1821)

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

**1x**: _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)
<details>
<summary>all...</summary>

* _2026-08-04 01:45:36 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18671/pull-kubevirt-e2e-k8s-1.36-sig-compute-serial-1.9/2084455542964621312#1:build-log.txt%3A604)

</details>

<hr/>
</details>

### sig-network (1x / 50.00%)


#### external (1x / 100.00%)

<details>
<summary> transient kube-apiserver body decode noise (from secondary snippet) (1x / 100.00%) </summary>

<hr/>

**1x**: _2026-08-03 09:30:59 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18291/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2084206350073597952#1:build-log.txt%3A1821)
<details>
<summary>all...</summary>

* _2026-08-03 09:30:59 &#43;0000 UTC_: <code>make: *** [Makefile:177: cluster-down] Error 1</code> [build-log](https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/18291/pull-kubevirt-e2e-k8s-1.36-ipv6-sig-network/2084206350073597952#1:build-log.txt%3A1821)

</details>

<hr/>
</details>

Last updated: 2026-08-07 12:27:12
