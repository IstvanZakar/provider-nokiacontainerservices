/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// cnaParameters are the configurable fields of a cna.
type cnaParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// cnaObservation are the observable fields of a cna.
type cnaObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A cnaSpec defines the desired state of a cna.
type cnaSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       cnaParameters `json:"forProvider"`
}

// A cnaStatus represents the observed state of a cna.
type cnaStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          cnaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A cna is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nokiacontainerservices}
type cna struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   cnaSpec   `json:"spec"`
	Status cnaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// cnaList contains a list of cna
type cnaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []cna `json:"items"`
}

// cna type metadata.
var (
	cnaKind             = reflect.TypeOf(cna{}).Name()
	cnaGroupKind        = schema.GroupKind{Group: Group, Kind: cnaKind}.String()
	cnaKindAPIVersion   = cnaKind + "." + SchemeGroupVersion.String()
	cnaGroupVersionKind = SchemeGroupVersion.WithKind(cnaKind)
)

func init() {
	SchemeBuilder.Register(&cna{}, &cnaList{})
}
