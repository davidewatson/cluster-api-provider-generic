/*
Copyright 2019 The Kubernetes Authors.

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
	"k8s.io/apimachinery/pkg/runtime"
)

// MachineCreateRequest defines a Machine creation request
type MachineCreateRequest struct {
	// MachineID is a unique identifier created by the Actuator to allow for idempotent
	// behavior. This can be used for tag-based lookups within the provider’s external APIs
	// or can be used in a way similar to AWS Client Tokens
	// Note: The Machine object UID should not be used for this value, since it does not
	// persist after pivoting cluster-api resources to a different management cluster. An
	// initial naive approach would be to leverage a concatenation of namespace/machineName or
	// namespace/clusterName/machineName, but longer term a more safe approach should
	// be used.
	MachineID string `json:"machineID"`

	// ProviderSpec details provider-specific configuration to use during machine creation
	// +optional
	ProviderSpec *runtime.RawExtension `json:"providerSpec,omitempty"`
}

// MachineCreateResponse defines a Machine creation response
type MachineCreateResponse struct {
	// ProviderID is the ID of the machine provided by the provider. This value
	// is intended to be used by the Actuator to populate the ProviderID field on
	// the Machine Object.
	// +optional
	ProviderID *string `json:"providerID,omitempty"`

	IPAddress string                 `json:"ipAddress"`
	Hostname  string                 `json:"hostname"`
	SSHConfig SSHConfig              `json:"sshConfig"`
	Status    ProviderResponseStatus `json:"status"`
}

// SSHConfig specifies everything needed to ssh to a host
type SSHConfig struct {
	// The IP or hostname used to SSH to the machine
	Host string `json:"host"`

	// The used to SSH to the machine
	Port int `json:"port"`

	// The SSH public keys of the machine
	PublicKeys []string `json:"publicKeys"`

	// The SSH private key used to SSH to the machine
	PrivateKey string `json:"privateKey"`
}

// ProviderResponseStatus
type ProviderResponseStatus struct {
	Success bool `json:"success"`

	// RetryableError
	// +optional
	RetryableError *bool `json:"retryableError,omitempty"`

	// ErrorMessage
	// +optional
	ErrorMessage string `json:"errorMessage,omitempty"`

	// ErrorReason
	// +optional
	ErrorReason string `json:"errorReason,omitempty"`
}

// MachineDeleteRequest defines a Machine deletion request
type MachineDeleteRequest struct {
	// MachineID is a unique identifier created by the Actuator to allow for idempotent
	// behavior. This should be the same value that was used for the MachineCreateRequest
	MachineID string `json:"machineID"`

	IPAddress string    `json:"ipAddress"`
	Hostname  string    `json:"hostname"`
	SSHConfig SSHConfig `json:"sshConfig"`

	// ProviderSpec details provider-specific configuration to use during machine creation
	// +optional
	ProviderSpec *runtime.RawExtension `json:"providerSpec,omitempty"`
}

// MachineDeleteResponse defines a Machine deletion request
type MachineDeleteResponse struct {
	Status ProviderResponseStatus `json:"status"`
}
