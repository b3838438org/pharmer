package api

import (
	"github.com/appscode/mergo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	ResourceCodeNodeGroup = "ng"
	ResourceKindNodeGroup = "NodeGroup"
	ResourceNameNodeGroup = "nodegroup"
	ResourceTypeNodeGroup = "nodegroups"
)

type NodeGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeGroupSpec   `json:"spec,omitempty"`
	Status            NodeGroupStatus `json:"status,omitempty"`
}

var _ runtime.Object = &NodeGroup{}

type NodeGroupSpec struct {
	Nodes int64 `json:"nodes"`

	// Template describes the nodes that will be created.
	Template NodeTemplateSpec `json:"template" protobuf:"bytes,3,opt,name=template"`
}

// NodeGroupStatus is the most recently observed status of the NodeGroup.
type NodeGroupStatus struct {
	// The number of pods that have labels matching the labels of the pod template of the node group.
	// +optional
	FullyLabeledNodes int64 `json:"fullyLabeledNodes,omitempty" protobuf:"varint,2,opt,name=fullyLabeledNodes"`

	// Nodes is the most recently oberved number of nodes.
	// Deprecated
	Nodes int64 `json:"nodes" protobuf:"varint,1,opt,name=nodes"`

	// The number of ready nodes for this node group.
	// +optional
	// Deprecated
	ReadyNodes int64 `json:"readyNodes,omitempty" protobuf:"varint,4,opt,name=readyNodes"`

	// The number of available nodes (ready for at least minReadySeconds) for this node group.
	// +optional
	// Deprecated
	AvailableNodes int64 `json:"availableNodes,omitempty" protobuf:"varint,5,opt,name=availableNodes"`

	// ObservedGeneration reflects the generation of the most recently observed node group.
	// +optiopharmer-linux-amd64nal
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,3,opt,name=observedGeneration"`

	// Represents the latest available observations of a node group's current state.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []NodeGroupCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
}

type NodeGroupConditionType string

// These are valid conditions of a node group.
const (
	// Available means the node group is available, ie. at least the minimum available
	// nodes required are up and running for at least minReadySeconds.
	NodeGroupAvailable NodeGroupConditionType = "Available"
	// Progressing means the deployment is progressing. Progress for a deployment is
	// considered when a new node group is created or adopted, and when new nodes scale
	// up or old nodes scale down. Progress is not estimated for paused node groups or
	// when progressDeadlineSeconds is not specified.
	NodeGroupProgressing NodeGroupConditionType = "Progressing"
	// ReplicaFailure is added in a deployment when one of its nodes fails to be created
	// or deleted.
	NodeGroupReplicaFailure NodeGroupConditionType = "ReplicaFailure"
)

type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

// NodeGroupCondition describes the state of a deployment at a certain point.
type NodeGroupCondition struct {
	// Type of deployment condition.
	Type NodeGroupConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=NodeGroupConditionType"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`
	// The last time this condition was updated.
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty" protobuf:"bytes,6,opt,name=lastUpdateTime"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,7,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

func (ng *NodeGroup) DeepCopyObject() runtime.Object {
	if ng == nil {
		return ng
	}
	out := new(NodeGroup)
	mergo.MergeWithOverwrite(out, ng)
	return out
}

func (ng NodeGroup) IsMaster() bool {
	_, found := ng.Labels[RoleMasterKey]
	return found
}

func (ng NodeGroup) Role() string {
	if ng.IsMaster() {
		return RoleMaster
	}
	return RoleNode
}

// PodTemplateSpec describes the data a pod should have when created from a template
type NodeTemplateSpec struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	// metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the pod.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec NodeSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// Deprecated, replace with Kubernetes Node
type Node struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeSpec   `json:"spec,omitempty"`
	Status            NodeStatus `json:"status,omitempty"`
}

type IPType string

const (
	IPTypeEphemeral IPType = "Ephemeral"
	IPTypeReserved  IPType = "Reserved"
)

// Deprecated
type NodeSpec struct {
	// Deprecated
	Role string

	SKU            string `json:"sku,omitempty"`
	SpotInstances  bool   `json:"spotInstances,omitempty"`
	DiskType       string `json:"nodeDiskType,omitempty"`
	DiskSize       int64  `json:"nodeDiskSize,omitempty"`
	ExternalIPType IPType `json:"externalIPType,omitempty"`
}

// Deprecated
type NodeStatus struct {
	Phase NodePhase

	Name          string
	ExternalID    string
	PublicIP      string
	PrivateIP     string
	ExternalPhase string
	DiskId        string `json:"diskID,omitempty"`
}

func (n Node) IsMaster() bool {
	_, found := n.Labels[RoleMasterKey]
	return found
}

// InstancePhase is a label for the condition of an Instance at the current time.
// Deprecated
type NodePhase string

const (
	NodeReady   NodePhase = "Ready"
	NodeDeleted NodePhase = "Deleted"
)

type SimpleNode struct {
	metav1.TypeMeta `json:",inline,omitempty"`
	Name            string
	ExternalID      string
	PublicIP        string
	PrivateIP       string
	DiskId          string `json:"diskID,omitempty"`
}

var _ runtime.Object = &SimpleNode{}

func (n *SimpleNode) DeepCopyObject() runtime.Object {
	if n == nil {
		return n
	}
	out := new(SimpleNode)
	mergo.MergeWithOverwrite(out, n)
	return out
}