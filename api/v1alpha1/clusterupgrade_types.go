/*
Copyright 2023.

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

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=imagebasedupgrades,shortName=ibu
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ImageBasedUpgrade is the Schema for the ImageBasedUpgrades API
// +operator-sdk:csv:customresourcedefinitions:displayName="Image-based Cluster Upgrade",resources={{Namespace, v1},{Deployment,apps/v1}}
type ImageBasedUpgrade struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImageBasedUpgradeSpec   `json:"spec,omitempty"`
	Status ImageBasedUpgradeStatus `json:"status,omitempty"`
}

// ImageBasedUpgradeSpec defines the desired state of ImageBasedUpgrade
type ImageBasedUpgradeSpec struct {
	Stage            ImageBasedUpgradeStage `json:"stage,omitempty"`
	SeedImageRef     SeedImageRef           `json:"seedImageRef,omitempty"`
	AdditionalImages ConfigMapRef           `json:"additionalImages,omitempty"`
	OADPContent      ConfigMapRef           `json:"oadpContent,omitempty"`
	ExtraManifests   ConfigMapRef           `json:"extraManifests,omitempty"`
	RollbackTarget   string                 `json:"rollbackTarget,omitempty"`
}

// SeedImageRef defines the stage of the image based upgrade
type ImageBasedUpgradeStage string

const (
	//UpgradePrep perform preliminary steps to get the node ready for image based upgrade
	UpgradePrep ImageBasedUpgradeStage = "preUpgrgade"
	//PreUpgrade backup the platform and application artifacts
	PreUpgrade ImageBasedUpgradeStage = "preUpgrgade"
	//Upgrade reboot to the new stateroot and reconfigure the cluster
	Upgrade ImageBasedUpgradeStage = "preUpgrgade"
)

// SeedImageRef defines the seed image and OCP version for the upgrade
type SeedImageRef struct {
	Version string `json:"version,omitempty"`
	Image   string `json:"image,omitempty"`
}

// ConfigMapRef defines a reference to a config map
type ConfigMapRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// ImageBasedUpgradeStatus defines the observed state of ImageBasedUpgrade
type ImageBasedUpgradeStatus struct {
	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="Status"
	ObservedGeneration int64       `json:"observedGeneration,omitempty"`
	StartedAt          metav1.Time `json:"startedAt,omitempty"`
	CompletedAt        metav1.Time `json:"completedAt,omitempty"`
	StateRoots         []StateRoot `json:"stateRoots,omitempty"`
}

// StateRoot defines a list of saved pod states and the running OCP version when they are saved
type StateRoot struct {
	Version string `json:"version,omitempty"`
	// TODO add fields for saved states
}

// +kubebuilder:object:root=true

// ImageBasedUpgradeList contains a list of ImageBasedUpgrade
type ImageBasedUpgradeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageBasedUpgrade `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImageBasedUpgrade{}, &ImageBasedUpgradeList{})
}
