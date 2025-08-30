/*
Copyright 2025 janos.

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

// TLSSecretSyncSpec defines the desired state of TLSSecretSync
type TLSSecretSyncSpec struct {
	// SourceRef points to the source TLS Secret to copy from.
	SourceRef SecretRef `json:"sourceRef"`

	// Targets lists the namespaces that should receive a copy of the Secret.
	Targets TargetSpec `json:"targets"`

	// CopyAnnotations copies annotations from the source Secret to targets.
	// +kubebuilder:default:=true
	CopyAnnotations bool `json:"copyAnnotations,omitempty"`

	// CopyLabels copies labels from the source Secret to targets.
	// +kubebuilder:default:=true
	CopyLabels bool `json:"copyLabels,omitempty"`

	// PruneOnDelete deletes target Secrets when this resource is deleted.
	// +kubebuilder:default:=true
	PruneOnDelete bool `json:"pruneOnDelete,omitempty"`

	// RefreshPolicy controls when synchronization happens.
	// Mode: "OnChange" or "Periodic" (if Periodic, set periodSeconds > 0).
	// +kubebuilder:validation:Optional
	RefreshPolicy *RefreshPolicy `json:"refreshPolicy,omitempty"`
}

type SecretRef struct {
	// Namespace of the source Secret.
	// +kubebuilder:validation:MinLength=1
	Namespace string `json:"namespace"`
	// Name of the source Secret.
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
}

type TargetSpec struct {
	// Target namespaces to copy secrets to
	// +kubebuilder:validation:MinItems=1
	// +listType=atomic
	Namespaces []string `json:"namespaces"`
}

type RefreshPolicy struct {
	// Mode selects sync behavior.
	// +kubebuilder:validation:Enum=OnChange;Periodic
	Mode string `json:"mode"`
	// PeriodSeconds is used only when Mode=Periodic.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Optional
	PeriodSeconds *int64 `json:"periodSeconds,omitempty"`
}

// TLSSecretSyncStatus defines the observed state of TLSSecretSync
type TLSSecretSyncStatus struct {
	// ObservedSourceResourceVersion is the last processed RV of the source Secret.
	// +optional
	ObservedSourceResourceVersion string `json:"observedSourceResourceVersion,omitempty"`

	// Summary provides aggregate counters for sync state.
	// +optional
	Summary *SummaryStatus `json:"summary,omitempty"`

	// Conditions follow K8s conventions (e.g., Ready, Degraded).
	// +kubebuilder:validation:Optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Failures lists only targets currently failing to sync.
	// Entries are unique per namespace and removed on recovery.
	// +kubebuilder:validation:Optional
	// +listType=atomic
	Failures []FailureEntry `json:"failures,omitempty"`
}

// SummaryStatus aggregates sync counts.
type SummaryStatus struct {
	// TotalTargets is how many namespaces are desired.
	// +kubebuilder:validation:Minimum=0
	TotalTargets int32 `json:"totalTargets"`

	// Synced is how many are currently in sync.
	// +kubebuilder:validation:Minimum=0
	Synced int32 `json:"synced"`

	// Failed is how many are currently failing (len(Failures)).
	// +kubebuilder:validation:Minimum=0
	Failed int32 `json:"failed"`

	// LastSyncTime is when the last reconciliation pass finished.
	// +optional
	LastSyncTime *metav1.Time `json:"lastSyncTime,omitempty"`
}

// FailureEntry describes a target namespace that failed to sync.
type FailureEntry struct {
	// Namespace of the failing target.
	// +kubebuilder:validation:MinLength=1
	Namespace string `json:"namespace"`

	// Reason is a short, machine-readable code (e.g., NamespaceMissing, RBACDenied).
	// +kubebuilder:validation:MinLength=1
	Reason string `json:"reason"`

	// Message is a human-readable description of the failure.
	// +optional
	Message string `json:"message,omitempty"`

	// LastErrorTime records when this failure last occurred.
	// +optional
	LastErrorTime *metav1.Time `json:"lastErrorTime,omitempty"`

	// RetryCount is how many consecutive reconcile attempts failed for this namespace.
	// +kubebuilder:validation:Minimum=0
	// +optional
	RetryCount int32 `json:"retryCount,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TLSSecretSync is the Schema for the tlssecretsyncs API
type TLSSecretSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TLSSecretSyncSpec   `json:"spec,omitempty"`
	Status TLSSecretSyncStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TLSSecretSyncList contains a list of TLSSecretSync
type TLSSecretSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TLSSecretSync `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TLSSecretSync{}, &TLSSecretSyncList{})
}
