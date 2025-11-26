# SIG Failure Reports Summary

This document provides a summary of the SIG failure reports for KubeVirt.

## SIG Compute

*   **Prometheus, Endpoint, Running:** Prometheus endpoint tests are failing because the VMI does not enter the "Running" phase.
*   **Cluster, Profiler, Deadline:** A cluster profiler test failed with an "Internal error encountered: Get ... context deadline exceeded".
*   **CPU, Pinning, Timeout:** CPU pinning tests are timing out.
*   **VMI, Delayed, Timeout:** A test for delayed VMI running states due to Kubernetes client rate limiter changes also timed out.
*   **Live, Migration, Failed:** Live migration tests are failing because the migration state is not finalized as failed or the target pod fails during preparation.
*   **CloudInit, UserData, Timeout:** CloudInit UserData tests are timing out waiting for the VMI to enter the "Running" phase.
*   **MultiQueue, Behavior, Timeout:** MultiQueue behavior tests are timing out.
*   **VMI, Kernel, Boot:** A VMI with external kernel boot test received an unexpected warning event about mounting kernel artifacts.
*   **Live, Migration, Datavolume:** Live migration across namespaces with datavolume disk failed.
*   **Evacuation, Migration, Missing:** Evacuation-triggered migration tests are failing due to missing annotations or migrations not being created.
*   **VMI, Lifecycle, Softreboot:** A VMI lifecycle softreboot test timed out waiting for a reboot record.
*   **CPU, Hotplug, Incorrect:** CPU hotplug tests are failing due to incorrect CPU count after plugging vCPUs.
*   **VM, State, TPM:** A VM state with persistent TPM test failed because a command to a spawned process timed out.
*   **VMI, Lifecycle, Delete:** A VMI lifecycle delete test timed out waiting for the pod to be terminated.

[Source Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-compute-failure-report.html)

## SIG Storage

*   **Live, Migration, Timeout:** Multiple failures related to live migration across namespaces, where migrations are timing out or not being cancelled as expected.
*   **Online, Snapshot, Hotplug:** One test for online VM snapshots with hotplug disks failed because the "AgentConnected" condition was not found.
*   **fstrim, Fedora, Timeout:** A storage test related to `fstrim` on Fedora VMIs timed out.

[Source Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-storage-failure-report.html)

## SIG Network

*   **Interface, State, Migration:** A test for interface state up/down status during migration failed because the "AgentConnected" condition was not found.
*   **Masquerade, Binding, Ping:** A networking test for masquerade binding mechanism failed to ping `google.com` from the VMI.
*   **IPv6, Port, Timeout:** A similar test for IPv6 with a specific port number also timed out.

[Source Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-network-failure-report.html)

## SIG Operator

*   **Update, KubeVirt, Timeout:** Several tests related to updating KubeVirt from a previous release to a target tested release are timing out because one of the KubeVirt control-plane components is not ready.
*   **Deployment, Revert, Timeout:** A test failed because a deployment did not revert to its original state after 120 seconds.

[Source Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-operator-failure-report.html)

## SIG Monitoring

There are no SIG failures to display.

[Source Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-monitoring-failure-report.html)
