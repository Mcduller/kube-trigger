/*
Copyright 2022 The KubeVela Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// TriggerServiceSpec defines the desired state of TriggerService.
type TriggerServiceSpec struct {
	InstanceRef string `json:"instanceRef,omitempty"`
	// Config for kube-trigger
	Triggers []TriggerMeta `json:"triggers"`
}

// +kubebuilder:object:root=true

// TriggerService is the Schema for the kubetriggerconfigs API.
// +kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TriggerService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TriggerServiceSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerServiceList contains a list of TriggerService.
type TriggerServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TriggerService `json:"items"`
}

// TriggerMeta is the meta data of a trigger.
type TriggerMeta struct {
	Source  Source       `json:"source"`
	Filters []FilterMeta `json:"filters"`
	Actions []ActionMeta `json:"actions"`
}

// ActionMeta is what users type in their configurations, specifying what action
// they want to use and what properties they provided.
type ActionMeta struct {
	// Type is the template (identifier) of this action.
	Template string `json:"template"`

	// Properties are user-provided parameters. You should parse it yourself.
	// +kubebuilder:pruning:PreserveUnknownFields
	Properties *runtime.RawExtension `json:"properties,omitempty"`

	// Raw is the raw string representation of this action. Typically, you will
	// not use it. This is for identifying action instances.
	Raw string `json:"raw,omitempty"`
}

// FilterMeta is what users type in their configurations, specifying what filter
// they want to use and what properties they provided.
type FilterMeta struct {
	// Template is the template (identifier) of this filter.
	Template string `json:"template"`

	// Properties are user-provided parameters. You should parse it yourself.
	// +kubebuilder:pruning:PreserveUnknownFields
	Properties *runtime.RawExtension `json:"properties,omitempty"`
}

// Source defines the Source of trigger.
type Source struct {
	Template SourceTemplate `json:"template"`
	// +kubebuilder:pruning:PreserveUnknownFields
	Properties *runtime.RawExtension `json:"properties"`
}

// SourceTemplate is the type of source template
type SourceTemplate string

const (
	// SourceTypeResourceWatcher is the source type for K8sResourceWatcher.
	SourceTypeResourceWatcher SourceTemplate = "k8s-resource-watcher"
	// SourceTypeWebhookTrigger is the source type for WebhookTrigger.
	SourceTypeWebhookTrigger SourceTemplate = "webhook-trigger"
)

// ActionTemplate is the type pf action template
type ActionTemplate string

const (
	// ActionTypeBumpApplicationRevision is the source type for BumpApplicationRevision.
	ActionTypeBumpApplicationRevision ActionTemplate = "bump-application-revision"
	// ActionTypePatchResource is the source type for WebhookTrigger.
	ActionTypePatchResource ActionTemplate = "patch-resource"
)

func init() {
	SchemeBuilder.Register(&TriggerService{}, &TriggerServiceList{})
}
