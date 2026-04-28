package cifailures

import "fmt"

var criticalEventReasons = map[string]bool{
	"FailedScheduling": true,
	"FailedMount":      true,
	"FailedCreate":     true,
	"Unhealthy":        true,
	"BackOff":          true,
	"Evicted":          true,
	"OOMKilling":       true,
	"FailedSync":       true,
	"NodeNotReady":     true,
}

func detectPodFailures(pods *K8sPodList) []*K8sFinding {
	if pods == nil {
		return nil
	}
	var findings []*K8sFinding

	for _, pod := range pods.Items {
		// Evicted pods
		if pod.Status.Phase == "Failed" && pod.Status.Reason == "Evicted" {
			findings = append(findings, &K8sFinding{
				Detector:  "pod-evicted",
				Severity:  "warning",
				Kind:      "Pod",
				Namespace: pod.Metadata.Namespace,
				Name:      pod.Metadata.Name,
				Reason:    "Evicted",
				Message:   pod.Status.Message,
			})
			continue
		}

		// Container statuses
		for _, cs := range pod.Status.ContainerStatuses {
			findings = append(findings, detectContainerIssues(&pod, &cs, "container")...)
		}

		// Init container statuses
		for _, cs := range pod.Status.InitContainerStatuses {
			findings = append(findings, detectContainerIssues(&pod, &cs, "init-container")...)
		}

		// Pending + Unschedulable
		if pod.Status.Phase == "Pending" {
			for _, cond := range pod.Status.Conditions {
				if cond.Type == "PodScheduled" && cond.Status == "False" {
					findings = append(findings, &K8sFinding{
						Detector:  "pod-unschedulable",
						Severity:  "error",
						Kind:      "Pod",
						Namespace: pod.Metadata.Namespace,
						Name:      pod.Metadata.Name,
						Reason:    cond.Reason,
						Message:   fmt.Sprintf("Pod stuck in Pending: %s", cond.Message),
					})
				}
			}
		}
	}

	return findings
}

func detectContainerIssues(pod *K8sPod, cs *K8sContainerStatus, containerType string) []*K8sFinding {
	var findings []*K8sFinding

	if cs.State.Waiting != nil {
		switch cs.State.Waiting.Reason {
		case "CrashLoopBackOff":
			findings = append(findings, &K8sFinding{
				Detector:  "pod-crash-loop",
				Severity:  "error",
				Kind:      "Pod",
				Namespace: pod.Metadata.Namespace,
				Name:      pod.Metadata.Name,
				Reason:    "CrashLoopBackOff",
				Message:   fmt.Sprintf("Container %s in CrashLoopBackOff (%d restarts)", cs.Name, cs.RestartCount),
			})
		case "ImagePullBackOff", "ErrImagePull":
			findings = append(findings, &K8sFinding{
				Detector:  "pod-image-pull",
				Severity:  "error",
				Kind:      "Pod",
				Namespace: pod.Metadata.Namespace,
				Name:      pod.Metadata.Name,
				Reason:    cs.State.Waiting.Reason,
				Message:   fmt.Sprintf("Container %s: %s", cs.Name, cs.State.Waiting.Message),
			})
		}
	}

	if cs.LastTerminationState.Terminated != nil && cs.LastTerminationState.Terminated.Reason == "OOMKilled" {
		findings = append(findings, &K8sFinding{
			Detector:  "pod-oom-killed",
			Severity:  "error",
			Kind:      "Pod",
			Namespace: pod.Metadata.Namespace,
			Name:      pod.Metadata.Name,
			Reason:    "OOMKilled",
			Message:   fmt.Sprintf("Container %s was OOMKilled (exit code %d)", cs.Name, cs.LastTerminationState.Terminated.ExitCode),
		})
	}

	if containerType == "init-container" && cs.State.Terminated != nil && cs.State.Terminated.ExitCode != 0 {
		findings = append(findings, &K8sFinding{
			Detector:  "pod-init-failure",
			Severity:  "error",
			Kind:      "Pod",
			Namespace: pod.Metadata.Namespace,
			Name:      pod.Metadata.Name,
			Reason:    fmt.Sprintf("InitContainerFailed:%s", cs.Name),
			Message:   fmt.Sprintf("Init container %s exited with code %d", cs.Name, cs.State.Terminated.ExitCode),
		})
	}

	if cs.RestartCount > 5 && cs.State.Waiting == nil {
		findings = append(findings, &K8sFinding{
			Detector:  "pod-high-restarts",
			Severity:  "warning",
			Kind:      "Pod",
			Namespace: pod.Metadata.Namespace,
			Name:      pod.Metadata.Name,
			Reason:    "HighRestartCount",
			Message:   fmt.Sprintf("Container %s has %d restarts", cs.Name, cs.RestartCount),
		})
	}

	if pod.Status.Phase == "Running" && !cs.Ready && cs.State.Waiting == nil {
		findings = append(findings, &K8sFinding{
			Detector:  "pod-not-ready",
			Severity:  "warning",
			Kind:      "Pod",
			Namespace: pod.Metadata.Namespace,
			Name:      pod.Metadata.Name,
			Reason:    "ContainerNotReady",
			Message:   fmt.Sprintf("Container %s is not ready", cs.Name),
		})
	}

	return findings
}

func detectNodeFailures(nodes *K8sNodeList) []*K8sFinding {
	if nodes == nil {
		return nil
	}
	var findings []*K8sFinding

	for _, node := range nodes.Items {
		for _, cond := range node.Status.Conditions {
			switch cond.Type {
			case "Ready":
				if cond.Status == "False" {
					findings = append(findings, &K8sFinding{
						Detector: "node-not-ready",
						Severity: "error",
						Kind:     "Node",
						Name:     node.Metadata.Name,
						Reason:   cond.Reason,
						Message:  fmt.Sprintf("Node condition Ready is False: %s", cond.Message),
					})
				}
			case "DiskPressure", "MemoryPressure", "PIDPressure":
				if cond.Status == "True" {
					findings = append(findings, &K8sFinding{
						Detector: "node-pressure",
						Severity: "error",
						Kind:     "Node",
						Name:     node.Metadata.Name,
						Reason:   cond.Type,
						Message:  fmt.Sprintf("Node has %s: %s", cond.Type, cond.Message),
					})
				}
			case "NetworkUnavailable":
				if cond.Status == "True" {
					findings = append(findings, &K8sFinding{
						Detector: "node-network-unavailable",
						Severity: "warning",
						Kind:     "Node",
						Name:     node.Metadata.Name,
						Reason:   cond.Reason,
						Message:  fmt.Sprintf("Node network unavailable: %s", cond.Message),
					})
				}
			}
		}
	}

	return findings
}

func detectEventFailures(events *K8sEventList) []*K8sFinding {
	if events == nil {
		return nil
	}

	type eventGroup struct {
		count          int
		sampleMessage  string
		involvedObject K8sObjectReference
	}
	grouped := map[string]*eventGroup{}

	for _, event := range events.Items {
		if event.Type != "Warning" {
			continue
		}
		g, exists := grouped[event.Reason]
		if !exists {
			g = &eventGroup{
				sampleMessage:  event.Message,
				involvedObject: event.InvolvedObject,
			}
			grouped[event.Reason] = g
		}
		g.count += max(event.Count, 1)
	}

	var findings []*K8sFinding
	for reason, g := range grouped {
		severity := "warning"
		if criticalEventReasons[reason] {
			severity = "error"
		}
		findings = append(findings, &K8sFinding{
			Detector:  "event-warning",
			Severity:  severity,
			Kind:      "Event",
			Namespace: g.involvedObject.Namespace,
			Name:      g.involvedObject.Name,
			Reason:    reason,
			Message:   fmt.Sprintf("%d warning event(s) with reason %s", g.count, reason),
			Detail:    g.sampleMessage,
		})
	}

	return findings
}

func detectVMIFailures(vmis *K8sVMIList) []*K8sFinding {
	if vmis == nil {
		return nil
	}
	var findings []*K8sFinding

	for _, vmi := range vmis.Items {
		if vmi.Status.Phase != "Running" && vmi.Status.Phase != "Succeeded" && vmi.Status.Phase != "" {
			severity := "warning"
			if vmi.Status.Phase == "Failed" {
				severity = "error"
			}
			findings = append(findings, &K8sFinding{
				Detector:  "vmi-not-running",
				Severity:  severity,
				Kind:      "VMI",
				Namespace: vmi.Metadata.Namespace,
				Name:      vmi.Metadata.Name,
				Reason:    vmi.Status.Phase,
				Message:   fmt.Sprintf("VMI phase is %s", vmi.Status.Phase),
			})
		}

		if vmi.Status.MigrationState != nil && vmi.Status.MigrationState.Failed {
			findings = append(findings, &K8sFinding{
				Detector:  "vmi-migration-failed",
				Severity:  "error",
				Kind:      "VMI",
				Namespace: vmi.Metadata.Namespace,
				Name:      vmi.Metadata.Name,
				Reason:    "MigrationFailed",
				Message: fmt.Sprintf("VMI migration failed (source: %s, target: %s)",
					vmi.Status.MigrationState.SourceNodeName, vmi.Status.MigrationState.TargetNodeName),
			})
		}

		for _, cond := range vmi.Status.Conditions {
			if cond.Type == "Ready" && cond.Status == "False" {
				findings = append(findings, &K8sFinding{
					Detector:  "vmi-not-ready",
					Severity:  "warning",
					Kind:      "VMI",
					Namespace: vmi.Metadata.Namespace,
					Name:      vmi.Metadata.Name,
					Reason:    cond.Reason,
					Message:   fmt.Sprintf("VMI not ready: %s", cond.Message),
				})
			}
		}
	}

	return findings
}

func detectEtcdIssues(profile *EtcdStorageProfile) []*K8sFinding {
	if profile == nil {
		return nil
	}
	var findings []*K8sFinding

	if profile.FinalTmpfsTotal > 0 {
		usagePct := float64(profile.PeakTmpfsUsed) / float64(profile.FinalTmpfsTotal) * 100

		if usagePct > 90 {
			findings = append(findings, &K8sFinding{
				Detector: "etcd-tmpfs-exhaustion",
				Severity: "error",
				Kind:     "EtcdProfile",
				Name:     "etcd",
				Reason:   "TmpfsExhaustion",
				Message:  fmt.Sprintf("Peak tmpfs usage %.1f%% (%d/%d bytes)", usagePct, profile.PeakTmpfsUsed, profile.FinalTmpfsTotal),
			})
		} else if usagePct > 75 {
			findings = append(findings, &K8sFinding{
				Detector: "etcd-tmpfs-pressure",
				Severity: "warning",
				Kind:     "EtcdProfile",
				Name:     "etcd",
				Reason:   "TmpfsPressure",
				Message:  fmt.Sprintf("Peak tmpfs usage %.1f%% (%d/%d bytes)", usagePct, profile.PeakTmpfsUsed, profile.FinalTmpfsTotal),
			})
		}

		walPct := float64(profile.FinalWALSize) / float64(profile.FinalTmpfsTotal) * 100
		if walPct > 50 {
			findings = append(findings, &K8sFinding{
				Detector: "etcd-large-wal",
				Severity: "warning",
				Kind:     "EtcdProfile",
				Name:     "etcd",
				Reason:   "LargeWAL",
				Message:  fmt.Sprintf("WAL consuming %.1f%% of tmpfs (%d/%d bytes)", walPct, profile.FinalWALSize, profile.FinalTmpfsTotal),
			})
		}
	}

	var totalDBGrowth int64
	for _, r := range profile.Records {
		if r.DeltaDBSize > 0 {
			totalDBGrowth += r.DeltaDBSize
		}
	}
	if totalDBGrowth > 0 {
		findings = append(findings, &K8sFinding{
			Detector: "etcd-db-growth",
			Severity: "info",
			Kind:     "EtcdProfile",
			Name:     "etcd",
			Reason:   "DBGrowth",
			Message:  fmt.Sprintf("Total DB growth across %d specs: %d bytes", profile.TotalSpecs, totalDBGrowth),
		})
	}

	return findings
}

func detectVMIMFailures(vmims *K8sVMIMList) []*K8sFinding {
	if vmims == nil {
		return nil
	}
	var findings []*K8sFinding

	for _, vmim := range vmims.Items {
		if vmim.Status.Phase == "Failed" {
			findings = append(findings, &K8sFinding{
				Detector:  "vmim-failed",
				Severity:  "error",
				Kind:      "VMIM",
				Namespace: vmim.Metadata.Namespace,
				Name:      vmim.Metadata.Name,
				Reason:    "Failed",
				Message:   "VirtualMachineInstanceMigration failed",
			})
		}
	}

	return findings
}
