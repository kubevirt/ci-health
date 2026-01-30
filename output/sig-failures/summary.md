# SIG Failure Summary

This document provides a summary of the SIG failure reports.

## SIG Compute

*   **VMI, Migration, Timeout:** The test "Migration should fail if target pod fails during target preparation" timed out, indicating issues with VMI migration state finalization.
*   **VMI, Configuration, Timeout:** Tests related to CPU pinning and node-labeller timed out or failed due to control-plane components not being ready.
*   **K8S, API, Connection-Refused:** A significant number of tests failed with "connect: connection refused" errors when attempting to interact with Kubernetes API endpoints. This affected tests in Infrastructure, VSOCK, ContainerDisk, InstancetypeReferencePolicy, ValidatingAdmissionPolicy, and VirtualMachine categories.
*   **VMI, Security, Timeout:** Tests for virt-launcher securityContext and SELinux types failed with timeouts or TLS handshake timeouts.
*   **VMI, ReplicaSet, Print:** A test for server-side printing of VirtualMachineInstanceReplicaSet returned an incorrect replica count.
*   **VMI, Lifecycle, Timeout:** A softreboot test timed out.

[Source Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-compute-failure-report.html)

## SIG Storage

*   **VMI, Migration, Timeout:** The test "[sig-storage] Hotplug [storage-req] VMI migration should allow live migration with attached hotplug volumes containerDisk VMI" timed out while waiting for volume statuses to be in the Ready phase.
*   **VMI, Migration, Canceled:** Tests for canceling live migration across namespaces by deleting the migration resource failed because the migration did not complete as expected.

[Source Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-storage-failure-report.html)

## SIG Network

*   **VMI, SRIOV, Timeout:** The test "SRIOV VMI connected to single SRIOV network memory hotplug Should successfully reattach host-device" timed out and failed to match expected network interface configurations.
*   **VMI, SRIOV, Timeout:** The test "[sig-network] SRIOV nic-hotplug a running VM can hotplug a network interface" timed out due to missing or extra network interface elements.
*   **VMI, Subdomain, Timeout:** The test "[sig-network] Subdomain with a headless service given VMI should have the expected FQDN with Masquerade binding without subdomain" failed with a timeout error.

[Source Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-network-failure-report.html)

## SIG Operator

*   No SIG failures to display.

[Source Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-operator-failure-report.html)

## SIG Monitoring

*   No SIG failures to display.

[Source Report](https://storage.googleapis.com/kubevirt-prow/reports/sig-failure-reports/sig-monitoring-failure-report.html)