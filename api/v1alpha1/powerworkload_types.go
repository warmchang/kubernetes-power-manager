/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type NodeInfo struct {
	// The name of the node associated with these containers and CPUs
	Name   string `json:"name,omitempty"`

	// The containers that are utilizing this workload
	Containers []Container `json:"containers,omitempty"`

	// All of the CPUs accross each container
	CpuIds []int  `json:"cpuIds,omitempty"`
}

// PowerWorkloadSpec defines the desired state of PowerWorkload
type PowerWorkloadSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The name of the workload
	Name string `json:"name"`

	// AllCores determines if the Workload is to be applied to all cores (i.e. use the Default Workload)
	AllCores bool `json:"allCores,omitempty"`

	// Reserved CPUs are the CPUs that have been reserved by Kubelet for use by the Kubernetes admin process
	// This list must match the list in the user's Kubelet configuration
	ReservedCPUs []int `json:"reservedCPUs,omitempty"`

	// The labels signifying the nodes the user wants to use
	PowerNodeSelector map[string]string `json:"powerNodeSelector,omitempty"`

	// Holds the info on the node name and cpu ids for each node
	Node NodeInfo `json:"nodeInfo,omitempty"`

	// PowerProfile is the Profile that this PowerWorkload is based on
	PowerProfile string `json:"powerProfile,omitempty"`
}

// PowerWorkloadStatus defines the observed state of PowerWorkload
type PowerWorkloadStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Shared Cores is the Core List that represents the Shared Cores on the node, only used by a Shared PowerWorkload
	SharedCores []int `json:sharedCores,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// PowerWorkload is the Schema for the powerworkloads API
type PowerWorkload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PowerWorkloadSpec   `json:"spec,omitempty"`
	Status PowerWorkloadStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PowerWorkloadList contains a list of PowerWorkload
type PowerWorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PowerWorkload `json:"items"`
}

/*
func (pw *PowerWorkload) GetNodeList() []string {
	nodeList := make([]string, 0)
	for _, nodeInfo := range pw.Spec.Nodes {
		nodeList = append(nodeList, nodeInfo.Name)
	}

	return nodeList
}
*/

func init() {
	SchemeBuilder.Register(&PowerWorkload{}, &PowerWorkloadList{})
}
