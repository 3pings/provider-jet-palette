/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CredentialsObservation struct {
}

type CredentialsParameters struct {

	// +kubebuilder:validation:Required
	CredentialType *string `json:"credentialType" tf:"credential_type,omitempty"`

	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// +kubebuilder:validation:Optional
	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type HelmObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HelmParameters struct {

	// +kubebuilder:validation:Required
	Credentials []CredentialsParameters `json:"credentials" tf:"credentials,omitempty"`

	// +kubebuilder:validation:Required
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// +kubebuilder:validation:Required
	IsPrivate *bool `json:"isPrivate" tf:"is_private,omitempty"`
}

// HelmSpec defines the desired state of Helm
type HelmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HelmParameters `json:"forProvider"`
}

// HelmStatus defines the observed state of Helm.
type HelmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HelmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Helm is the Schema for the Helms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palettejet}
type Helm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HelmSpec   `json:"spec"`
	Status            HelmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HelmList contains a list of Helms
type HelmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Helm `json:"items"`
}

// Repository type metadata.
var (
	Helm_Kind             = "Helm"
	Helm_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Helm_Kind}.String()
	Helm_KindAPIVersion   = Helm_Kind + "." + CRDGroupVersion.String()
	Helm_GroupVersionKind = CRDGroupVersion.WithKind(Helm_Kind)
)

func init() {
	SchemeBuilder.Register(&Helm{}, &HelmList{})
}
