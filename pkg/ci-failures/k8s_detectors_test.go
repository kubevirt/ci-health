package cifailures

import (
	"testing"
)

func TestDetectPodFailures(t *testing.T) {
	tests := []struct {
		name             string
		pods             *K8sPodList
		expectedCount    int
		expectedDetector string
		expectedReason   string
	}{
		{
			name:          "nil pods",
			pods:          nil,
			expectedCount: 0,
		},
		{
			name:          "empty pod list",
			pods:          &K8sPodList{},
			expectedCount: 0,
		},
		{
			name: "healthy pod",
			pods: &K8sPodList{Items: []K8sPod{{
				Status: K8sPodStatus{
					Phase: "Running",
					ContainerStatuses: []K8sContainerStatus{{
						Name:  "app",
						Ready: true,
					}},
				},
			}}},
			expectedCount: 0,
		},
		{
			name: "CrashLoopBackOff",
			pods: &K8sPodList{Items: []K8sPod{{
				Metadata: K8sObjectMeta{Name: "test-pod", Namespace: "default"},
				Status: K8sPodStatus{
					Phase: "Running",
					ContainerStatuses: []K8sContainerStatus{{
						Name:         "app",
						RestartCount: 12,
						State: K8sContainerState{
							Waiting: &K8sContainerStateWaiting{Reason: "CrashLoopBackOff"},
						},
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "pod-crash-loop",
			expectedReason:   "CrashLoopBackOff",
		},
		{
			name: "OOMKilled",
			pods: &K8sPodList{Items: []K8sPod{{
				Metadata: K8sObjectMeta{Name: "oom-pod", Namespace: "ns"},
				Status: K8sPodStatus{
					Phase: "Running",
					ContainerStatuses: []K8sContainerStatus{{
						Name:  "app",
						Ready: true,
						LastTerminationState: K8sContainerState{
							Terminated: &K8sContainerStateTerminated{
								Reason:   "OOMKilled",
								ExitCode: 137,
							},
						},
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "pod-oom-killed",
			expectedReason:   "OOMKilled",
		},
		{
			name: "ImagePullBackOff",
			pods: &K8sPodList{Items: []K8sPod{{
				Metadata: K8sObjectMeta{Name: "img-pod", Namespace: "ns"},
				Status: K8sPodStatus{
					Phase: "Pending",
					ContainerStatuses: []K8sContainerStatus{{
						Name: "app",
						State: K8sContainerState{
							Waiting: &K8sContainerStateWaiting{
								Reason:  "ImagePullBackOff",
								Message: "Back-off pulling image",
							},
						},
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "pod-image-pull",
			expectedReason:   "ImagePullBackOff",
		},
		{
			name: "ErrImagePull",
			pods: &K8sPodList{Items: []K8sPod{{
				Metadata: K8sObjectMeta{Name: "img-pod", Namespace: "ns"},
				Status: K8sPodStatus{
					Phase: "Pending",
					ContainerStatuses: []K8sContainerStatus{{
						Name: "app",
						State: K8sContainerState{
							Waiting: &K8sContainerStateWaiting{Reason: "ErrImagePull"},
						},
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "pod-image-pull",
			expectedReason:   "ErrImagePull",
		},
		{
			name: "Pending unschedulable",
			pods: &K8sPodList{Items: []K8sPod{{
				Metadata: K8sObjectMeta{Name: "pending-pod", Namespace: "ns"},
				Status: K8sPodStatus{
					Phase: "Pending",
					Conditions: []K8sPodCondition{{
						Type:    "PodScheduled",
						Status:  "False",
						Reason:  "Unschedulable",
						Message: "0/3 nodes are available",
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "pod-unschedulable",
			expectedReason:   "Unschedulable",
		},
		{
			name: "Evicted pod",
			pods: &K8sPodList{Items: []K8sPod{{
				Metadata: K8sObjectMeta{Name: "evicted-pod", Namespace: "ns"},
				Status: K8sPodStatus{
					Phase:   "Failed",
					Reason:  "Evicted",
					Message: "The node was low on resource: memory.",
				},
			}}},
			expectedCount:    1,
			expectedDetector: "pod-evicted",
			expectedReason:   "Evicted",
		},
		{
			name: "Init container failure",
			pods: &K8sPodList{Items: []K8sPod{{
				Metadata: K8sObjectMeta{Name: "init-pod", Namespace: "ns"},
				Status: K8sPodStatus{
					Phase: "Init:Error",
					InitContainerStatuses: []K8sContainerStatus{{
						Name: "init",
						State: K8sContainerState{
							Terminated: &K8sContainerStateTerminated{
								ExitCode: 1,
								Reason:   "Error",
							},
						},
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "pod-init-failure",
		},
		{
			name: "High restart count",
			pods: &K8sPodList{Items: []K8sPod{{
				Metadata: K8sObjectMeta{Name: "restart-pod", Namespace: "ns"},
				Status: K8sPodStatus{
					Phase: "Running",
					ContainerStatuses: []K8sContainerStatus{{
						Name:         "app",
						Ready:        true,
						RestartCount: 10,
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "pod-high-restarts",
			expectedReason:   "HighRestartCount",
		},
		{
			name: "Pod running but container not ready",
			pods: &K8sPodList{Items: []K8sPod{{
				Metadata: K8sObjectMeta{Name: "notready-pod", Namespace: "ns"},
				Status: K8sPodStatus{
					Phase: "Running",
					ContainerStatuses: []K8sContainerStatus{{
						Name:  "app",
						Ready: false,
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "pod-not-ready",
			expectedReason:   "ContainerNotReady",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findings := detectPodFailures(tt.pods)
			if len(findings) != tt.expectedCount {
				t.Fatalf("expected %d findings, got %d: %+v", tt.expectedCount, len(findings), findings)
			}
			if tt.expectedCount > 0 {
				if findings[0].Detector != tt.expectedDetector {
					t.Errorf("expected detector %q, got %q", tt.expectedDetector, findings[0].Detector)
				}
				if tt.expectedReason != "" && findings[0].Reason != tt.expectedReason {
					t.Errorf("expected reason %q, got %q", tt.expectedReason, findings[0].Reason)
				}
			}
		})
	}
}

func TestDetectNodeFailures(t *testing.T) {
	tests := []struct {
		name             string
		nodes            *K8sNodeList
		expectedCount    int
		expectedDetector string
		expectedReason   string
	}{
		{
			name:          "nil nodes",
			nodes:         nil,
			expectedCount: 0,
		},
		{
			name: "healthy node",
			nodes: &K8sNodeList{Items: []K8sNode{{
				Metadata: K8sObjectMeta{Name: "node1"},
				Status: K8sNodeStatus{
					Conditions: []K8sNodeCondition{
						{Type: "Ready", Status: "True"},
						{Type: "DiskPressure", Status: "False"},
						{Type: "MemoryPressure", Status: "False"},
						{Type: "PIDPressure", Status: "False"},
					},
				},
			}}},
			expectedCount: 0,
		},
		{
			name: "NotReady node",
			nodes: &K8sNodeList{Items: []K8sNode{{
				Metadata: K8sObjectMeta{Name: "node1"},
				Status: K8sNodeStatus{
					Conditions: []K8sNodeCondition{{
						Type:    "Ready",
						Status:  "False",
						Reason:  "KubeletNotReady",
						Message: "container runtime is down",
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "node-not-ready",
			expectedReason:   "KubeletNotReady",
		},
		{
			name: "MemoryPressure",
			nodes: &K8sNodeList{Items: []K8sNode{{
				Metadata: K8sObjectMeta{Name: "node1"},
				Status: K8sNodeStatus{
					Conditions: []K8sNodeCondition{{
						Type:   "MemoryPressure",
						Status: "True",
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "node-pressure",
			expectedReason:   "MemoryPressure",
		},
		{
			name: "DiskPressure",
			nodes: &K8sNodeList{Items: []K8sNode{{
				Metadata: K8sObjectMeta{Name: "node1"},
				Status: K8sNodeStatus{
					Conditions: []K8sNodeCondition{{
						Type:   "DiskPressure",
						Status: "True",
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "node-pressure",
			expectedReason:   "DiskPressure",
		},
		{
			name: "NetworkUnavailable",
			nodes: &K8sNodeList{Items: []K8sNode{{
				Metadata: K8sObjectMeta{Name: "node1"},
				Status: K8sNodeStatus{
					Conditions: []K8sNodeCondition{{
						Type:    "NetworkUnavailable",
						Status:  "True",
						Reason:  "NoRouteCreated",
						Message: "route not created",
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "node-network-unavailable",
			expectedReason:   "NoRouteCreated",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findings := detectNodeFailures(tt.nodes)
			if len(findings) != tt.expectedCount {
				t.Fatalf("expected %d findings, got %d: %+v", tt.expectedCount, len(findings), findings)
			}
			if tt.expectedCount > 0 {
				if findings[0].Detector != tt.expectedDetector {
					t.Errorf("expected detector %q, got %q", tt.expectedDetector, findings[0].Detector)
				}
				if tt.expectedReason != "" && findings[0].Reason != tt.expectedReason {
					t.Errorf("expected reason %q, got %q", tt.expectedReason, findings[0].Reason)
				}
			}
		})
	}
}

func TestDetectEventFailures(t *testing.T) {
	tests := []struct {
		name             string
		events           *K8sEventList
		expectedCount    int
		expectedSeverity string
	}{
		{
			name:          "nil events",
			events:        nil,
			expectedCount: 0,
		},
		{
			name: "Normal events ignored",
			events: &K8sEventList{Items: []K8sEvent{{
				Type:   "Normal",
				Reason: "Pulled",
				Count:  1,
			}}},
			expectedCount: 0,
		},
		{
			name: "Critical warning event",
			events: &K8sEventList{Items: []K8sEvent{{
				Type:           "Warning",
				Reason:         "FailedScheduling",
				Message:        "0/3 nodes are available",
				Count:          5,
				InvolvedObject: K8sObjectReference{Kind: "Pod", Name: "test", Namespace: "ns"},
			}}},
			expectedCount:    1,
			expectedSeverity: "error",
		},
		{
			name: "Non-critical warning event",
			events: &K8sEventList{Items: []K8sEvent{{
				Type:           "Warning",
				Reason:         "SomeOtherReason",
				Message:        "something happened",
				Count:          1,
				InvolvedObject: K8sObjectReference{Kind: "Pod", Name: "test", Namespace: "ns"},
			}}},
			expectedCount:    1,
			expectedSeverity: "warning",
		},
		{
			name: "Event with zero count treated as 1",
			events: &K8sEventList{Items: []K8sEvent{{
				Type:           "Warning",
				Reason:         "BackOff",
				Message:        "back-off restarting",
				Count:          0,
				InvolvedObject: K8sObjectReference{Kind: "Pod", Name: "test", Namespace: "ns"},
			}}},
			expectedCount:    1,
			expectedSeverity: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findings := detectEventFailures(tt.events)
			if len(findings) != tt.expectedCount {
				t.Fatalf("expected %d findings, got %d: %+v", tt.expectedCount, len(findings), findings)
			}
			if tt.expectedCount > 0 && tt.expectedSeverity != "" {
				if findings[0].Severity != tt.expectedSeverity {
					t.Errorf("expected severity %q, got %q", tt.expectedSeverity, findings[0].Severity)
				}
			}
		})
	}
}

func TestDetectVMIFailures(t *testing.T) {
	tests := []struct {
		name             string
		vmis             *K8sVMIList
		expectedCount    int
		expectedDetector string
		expectedSeverity string
	}{
		{
			name:          "nil vmis",
			vmis:          nil,
			expectedCount: 0,
		},
		{
			name: "Running VMI",
			vmis: &K8sVMIList{Items: []K8sVMI{{
				Status: K8sVMIStatus{Phase: "Running"},
			}}},
			expectedCount: 0,
		},
		{
			name: "Succeeded VMI",
			vmis: &K8sVMIList{Items: []K8sVMI{{
				Status: K8sVMIStatus{Phase: "Succeeded"},
			}}},
			expectedCount: 0,
		},
		{
			name: "Failed VMI",
			vmis: &K8sVMIList{Items: []K8sVMI{{
				Metadata: K8sObjectMeta{Name: "vmi1", Namespace: "ns"},
				Status:   K8sVMIStatus{Phase: "Failed"},
			}}},
			expectedCount:    1,
			expectedDetector: "vmi-not-running",
			expectedSeverity: "error",
		},
		{
			name: "Scheduling VMI",
			vmis: &K8sVMIList{Items: []K8sVMI{{
				Metadata: K8sObjectMeta{Name: "vmi1", Namespace: "ns"},
				Status:   K8sVMIStatus{Phase: "Scheduling"},
			}}},
			expectedCount:    1,
			expectedDetector: "vmi-not-running",
			expectedSeverity: "warning",
		},
		{
			name: "Failed migration",
			vmis: &K8sVMIList{Items: []K8sVMI{{
				Metadata: K8sObjectMeta{Name: "vmi1", Namespace: "ns"},
				Status: K8sVMIStatus{
					Phase: "Running",
					MigrationState: &K8sMigrationState{
						Failed:         true,
						SourceNodeName: "node1",
						TargetNodeName: "node2",
					},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "vmi-migration-failed",
			expectedSeverity: "error",
		},
		{
			name: "VMI not ready condition",
			vmis: &K8sVMIList{Items: []K8sVMI{{
				Metadata: K8sObjectMeta{Name: "vmi1", Namespace: "ns"},
				Status: K8sVMIStatus{
					Phase: "Running",
					Conditions: []K8sVMICondition{{
						Type:    "Ready",
						Status:  "False",
						Reason:  "GuestNotRunning",
						Message: "guest agent not running",
					}},
				},
			}}},
			expectedCount:    1,
			expectedDetector: "vmi-not-ready",
			expectedSeverity: "warning",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findings := detectVMIFailures(tt.vmis)
			if len(findings) != tt.expectedCount {
				t.Fatalf("expected %d findings, got %d: %+v", tt.expectedCount, len(findings), findings)
			}
			if tt.expectedCount > 0 {
				if findings[0].Detector != tt.expectedDetector {
					t.Errorf("expected detector %q, got %q", tt.expectedDetector, findings[0].Detector)
				}
				if findings[0].Severity != tt.expectedSeverity {
					t.Errorf("expected severity %q, got %q", tt.expectedSeverity, findings[0].Severity)
				}
			}
		})
	}
}

func TestDetectEtcdIssues(t *testing.T) {
	tests := []struct {
		name             string
		profile          *EtcdStorageProfile
		expectedCount    int
		expectedDetector string
		expectedSeverity string
	}{
		{
			name:          "nil profile",
			profile:       nil,
			expectedCount: 0,
		},
		{
			name: "healthy profile",
			profile: &EtcdStorageProfile{
				TotalSpecs:      10,
				FinalTmpfsTotal: 1024 * 1024 * 1024,
				PeakTmpfsUsed:   100 * 1024 * 1024,
				FinalWALSize:    50 * 1024 * 1024,
			},
			expectedCount: 0,
		},
		{
			name: "tmpfs exhaustion over 90%",
			profile: &EtcdStorageProfile{
				TotalSpecs:      10,
				FinalTmpfsTotal: 512 * 1024 * 1024,
				PeakTmpfsUsed:   480 * 1024 * 1024, // ~93.75%
				FinalWALSize:    50 * 1024 * 1024,
			},
			expectedCount:    1,
			expectedDetector: "etcd-tmpfs-exhaustion",
			expectedSeverity: "error",
		},
		{
			name: "tmpfs pressure between 75% and 90%",
			profile: &EtcdStorageProfile{
				TotalSpecs:      10,
				FinalTmpfsTotal: 512 * 1024 * 1024,
				PeakTmpfsUsed:   420 * 1024 * 1024, // ~82%
				FinalWALSize:    50 * 1024 * 1024,
			},
			expectedCount:    1,
			expectedDetector: "etcd-tmpfs-pressure",
			expectedSeverity: "warning",
		},
		{
			name: "large WAL over 50% of tmpfs",
			profile: &EtcdStorageProfile{
				TotalSpecs:      10,
				FinalTmpfsTotal: 512 * 1024 * 1024,
				PeakTmpfsUsed:   100 * 1024 * 1024,
				FinalWALSize:    300 * 1024 * 1024, // ~58.6%
			},
			expectedCount:    1,
			expectedDetector: "etcd-large-wal",
			expectedSeverity: "warning",
		},
		{
			name: "DB growth across specs",
			profile: &EtcdStorageProfile{
				TotalSpecs:      3,
				FinalTmpfsTotal: 1024 * 1024 * 1024,
				PeakTmpfsUsed:   100 * 1024 * 1024,
				FinalWALSize:    50 * 1024 * 1024,
				Records: []EtcdStorageSpecRecord{
					{SpecName: "spec1", DeltaDBSize: 1000},
					{SpecName: "spec2", DeltaDBSize: -500},
					{SpecName: "spec3", DeltaDBSize: 2000},
				},
			},
			expectedCount:    1,
			expectedDetector: "etcd-db-growth",
			expectedSeverity: "info",
		},
		{
			name: "zero tmpfs total avoids division by zero",
			profile: &EtcdStorageProfile{
				TotalSpecs:      1,
				FinalTmpfsTotal: 0,
				PeakTmpfsUsed:   100,
			},
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findings := detectEtcdIssues(tt.profile)
			if len(findings) != tt.expectedCount {
				t.Fatalf("expected %d findings, got %d: %+v", tt.expectedCount, len(findings), findings)
			}
			if tt.expectedCount > 0 {
				if findings[0].Detector != tt.expectedDetector {
					t.Errorf("expected detector %q, got %q", tt.expectedDetector, findings[0].Detector)
				}
				if findings[0].Severity != tt.expectedSeverity {
					t.Errorf("expected severity %q, got %q", tt.expectedSeverity, findings[0].Severity)
				}
				if findings[0].Kind != "EtcdProfile" {
					t.Errorf("expected kind %q, got %q", "EtcdProfile", findings[0].Kind)
				}
			}
		})
	}
}

func TestDetectVMIMFailures(t *testing.T) {
	tests := []struct {
		name          string
		vmims         *K8sVMIMList
		expectedCount int
	}{
		{
			name:          "nil vmims",
			vmims:         nil,
			expectedCount: 0,
		},
		{
			name: "Succeeded VMIM",
			vmims: &K8sVMIMList{Items: []K8sVMIM{{
				Status: K8sVMIMStatus{Phase: "Succeeded"},
			}}},
			expectedCount: 0,
		},
		{
			name: "Failed VMIM",
			vmims: &K8sVMIMList{Items: []K8sVMIM{{
				Metadata: K8sObjectMeta{Name: "mig1", Namespace: "ns"},
				Status:   K8sVMIMStatus{Phase: "Failed"},
			}}},
			expectedCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findings := detectVMIMFailures(tt.vmims)
			if len(findings) != tt.expectedCount {
				t.Fatalf("expected %d findings, got %d: %+v", tt.expectedCount, len(findings), findings)
			}
			if tt.expectedCount > 0 {
				if findings[0].Detector != "vmim-failed" {
					t.Errorf("expected detector %q, got %q", "vmim-failed", findings[0].Detector)
				}
			}
		})
	}
}
