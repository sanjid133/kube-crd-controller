package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	SecDbLabel = "secret.data.name"
	KindSecDb  = "SecDb"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecDb is a specification for a SecDb resource
type SecDb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SecDbSpec `json:"spec"`
}

// SecDbSpec is the spec for a SecDb resource
type SecDbSpec struct {
	SecretType string     `json:"secretType"`
	Info       []InfoSpec `json:"info"`
}

// InfoSpec is the item for a SecDbSpec resource
type InfoSpec struct {
	SecretName string            `json:"secretName"`
	Data       map[string]string `json:"data"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecDbList is a list of SecDb resources
type SecDbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SecDb `json:"items"`
}
