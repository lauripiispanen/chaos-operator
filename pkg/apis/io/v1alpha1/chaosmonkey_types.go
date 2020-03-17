package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ChaosMonkeySpec defines the desired state of ChaosMonkey
type ChaosMonkeySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// ChaosMonkeyStatus defines the observed state of ChaosMonkey
type ChaosMonkeyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChaosMonkey is the Schema for the chaosmonkeys API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=chaosmonkeys,scope=Namespaced
type ChaosMonkey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChaosMonkeySpec   `json:"spec,omitempty"`
	Status ChaosMonkeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChaosMonkeyList contains a list of ChaosMonkey
type ChaosMonkeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChaosMonkey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChaosMonkey{}, &ChaosMonkeyList{})
}
