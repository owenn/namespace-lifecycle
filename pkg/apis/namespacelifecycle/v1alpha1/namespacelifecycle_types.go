package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NamespaceLifecycleSpec defines the desired state of NamespaceLifecycle
// +k8s:openapi-gen=true
type NamespaceLifecycleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	ChargeCode string `json:"charge-code"`
	ChargeCodeName string `json:"charge-code-name"`
	ExpiryDate string `json:"expiry-date"`
}

// NamespaceLifecycleStatus defines the observed state of NamespaceLifecycle
// +k8s:openapi-gen=true
type NamespaceLifecycleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NamespaceLifecycle is the Schema for the namespacelifecycles API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type NamespaceLifecycle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespaceLifecycleSpec   `json:"spec,omitempty"`
	Status NamespaceLifecycleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NamespaceLifecycleList contains a list of NamespaceLifecycle
type NamespaceLifecycleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceLifecycle `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NamespaceLifecycle{}, &NamespaceLifecycleList{})
}
