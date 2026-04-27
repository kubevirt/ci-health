package cifailures

import (
	"fmt"
	"time"
)

// K8sObjectMeta captures the metadata fields we inspect.
type K8sObjectMeta struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	Labels            map[string]string `json:"labels,omitempty"`
	CreationTimestamp string            `json:"creationTimestamp,omitempty"`
}

// K8sPodList is a lightweight representation of a core/v1 PodList.
type K8sPodList struct {
	Items []K8sPod `json:"items"`
}

type K8sPod struct {
	Metadata K8sObjectMeta `json:"metadata"`
	Spec     K8sPodSpec    `json:"spec"`
	Status   K8sPodStatus  `json:"status"`
}

type K8sPodSpec struct {
	NodeName       string         `json:"nodeName,omitempty"`
	InitContainers []K8sContainer `json:"initContainers,omitempty"`
	Containers     []K8sContainer `json:"containers,omitempty"`
}

type K8sContainer struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type K8sPodStatus struct {
	Phase                 string               `json:"phase"`
	Reason                string               `json:"reason,omitempty"`
	Message               string               `json:"message,omitempty"`
	Conditions            []K8sPodCondition    `json:"conditions,omitempty"`
	ContainerStatuses     []K8sContainerStatus `json:"containerStatuses,omitempty"`
	InitContainerStatuses []K8sContainerStatus `json:"initContainerStatuses,omitempty"`
}

type K8sPodCondition struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type K8sContainerStatus struct {
	Name                 string            `json:"name"`
	Ready                bool              `json:"ready"`
	RestartCount         int               `json:"restartCount"`
	State                K8sContainerState `json:"state,omitzero"`
	LastTerminationState K8sContainerState `json:"lastTerminationState,omitzero"`
}

type K8sContainerState struct {
	Waiting    *K8sContainerStateWaiting    `json:"waiting,omitempty"`
	Running    *K8sContainerStateRunning    `json:"running,omitempty"`
	Terminated *K8sContainerStateTerminated `json:"terminated,omitempty"`
}

type K8sContainerStateWaiting struct {
	Reason  string `json:"reason"`
	Message string `json:"message,omitempty"`
}

type K8sContainerStateRunning struct {
	StartedAt string `json:"startedAt,omitempty"`
}

type K8sContainerStateTerminated struct {
	ExitCode   int    `json:"exitCode"`
	Reason     string `json:"reason,omitempty"`
	Message    string `json:"message,omitempty"`
	Signal     int    `json:"signal,omitempty"`
	StartedAt  string `json:"startedAt,omitempty"`
	FinishedAt string `json:"finishedAt,omitempty"`
}

// K8sNodeList is a lightweight representation of a core/v1 NodeList.
type K8sNodeList struct {
	Items []K8sNode `json:"items"`
}

type K8sNode struct {
	Metadata K8sObjectMeta `json:"metadata"`
	Status   K8sNodeStatus `json:"status"`
}

type K8sNodeStatus struct {
	Conditions []K8sNodeCondition `json:"conditions,omitempty"`
}

type K8sNodeCondition struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

// K8sEventList is a lightweight representation of a core/v1 EventList.
type K8sEventList struct {
	Items []K8sEvent `json:"items"`
}

type K8sEvent struct {
	Metadata       K8sObjectMeta      `json:"metadata"`
	InvolvedObject K8sObjectReference `json:"involvedObject"`
	Reason         string             `json:"reason"`
	Message        string             `json:"message"`
	Type           string             `json:"type"`
	Count          int                `json:"count"`
	LastTimestamp  string             `json:"lastTimestamp,omitempty"`
	FirstTimestamp string             `json:"firstTimestamp,omitempty"`
}

type K8sObjectReference struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

// K8sVMIList is a lightweight representation of a KubeVirt VirtualMachineInstanceList.
type K8sVMIList struct {
	Items []K8sVMI `json:"items"`
}

type K8sVMI struct {
	Metadata K8sObjectMeta `json:"metadata"`
	Status   K8sVMIStatus  `json:"status"`
}

type K8sVMIStatus struct {
	Phase          string             `json:"phase"`
	Reason         string             `json:"reason,omitempty"`
	Conditions     []K8sVMICondition  `json:"conditions,omitempty"`
	MigrationState *K8sMigrationState `json:"migrationState,omitempty"`
}

type K8sVMICondition struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type K8sMigrationState struct {
	Completed      bool   `json:"completed"`
	Failed         bool   `json:"failed"`
	TargetNodeName string `json:"targetNode,omitempty"`
	SourceNodeName string `json:"sourceNode,omitempty"`
}

// K8sVMIMList is a lightweight representation of a KubeVirt VirtualMachineInstanceMigrationList.
type K8sVMIMList struct {
	Items []K8sVMIM `json:"items"`
}

type K8sVMIM struct {
	Metadata K8sObjectMeta `json:"metadata"`
	Status   K8sVMIMStatus `json:"status"`
}

type K8sVMIMStatus struct {
	Phase string `json:"phase"`
}

// K8sSnapshot identifies a cluster state capture point: a parallel process
// number and a failure count within that process.
type K8sSnapshot struct {
	Process      string `yaml:"process"`
	FailureCount int    `yaml:"failure_count"`
}

func (s K8sSnapshot) String() string {
	return fmt.Sprintf("process=%s/failure=%d", s.Process, s.FailureCount)
}

// K8sFinding represents a single detected problem in a k8s object.
type K8sFinding struct {
	Detector  string      `yaml:"detector"`
	Severity  string      `yaml:"severity"`
	Kind      string      `yaml:"kind"`
	Namespace string      `yaml:"namespace,omitempty"`
	Name      string      `yaml:"name"`
	Reason    string      `yaml:"reason"`
	Message   string      `yaml:"message"`
	Detail    string      `yaml:"detail,omitempty"`
	Snapshot  K8sSnapshot `yaml:"snapshot"`
}

// K8sAnalysisResult is the top-level output for analyze-k8s.
type K8sAnalysisResult struct {
	ProwJobURL string        `yaml:"prow_job_url"`
	JobName    string        `yaml:"job_name"`
	BuildID    int           `yaml:"build_id"`
	Started    time.Time     `yaml:"started"`
	Finished   time.Time     `yaml:"finished"`
	Snapshots  []K8sSnapshot `yaml:"snapshots"`
	Findings   []*K8sFinding `yaml:"findings"`
	Summary    K8sSummary    `yaml:"summary"`
}

// K8sSummary provides aggregate counts by kind, severity, and detector.
type K8sSummary struct {
	TotalFindings int            `yaml:"total_findings"`
	ByKind        map[string]int `yaml:"by_kind"`
	BySeverity    map[string]int `yaml:"by_severity"`
	ByDetector    map[string]int `yaml:"by_detector"`
}
