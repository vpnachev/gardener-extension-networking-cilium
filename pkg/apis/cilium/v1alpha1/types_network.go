// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IdentityAllocationMode selects how identities are shared between cilium
// nodes by setting how they are stored. The options are "crd" or "kvstore".
type IdentityAllocationMode string

const (
	// CRD defines the crd IdentityAllocationMode type.
	CRD IdentityAllocationMode = "crd"
	// KVStore defines the kvstore IdentityAllocationMode type.
	KVStore IdentityAllocationMode = "kvstore"
)

// TunnelMode defines what tunnel mode to use for Cilium.
type TunnelMode string

const (
	// VXLan defines the vxlan tunnel mode
	VXLan TunnelMode = "vxlan"
	// Geneve defines the geneve tunnel mode.
	Geneve TunnelMode = "geneve"
	// Disabled defines the disabled tunnel mode.
	Disabled TunnelMode = "disabled"
)

// KubeProxyReplacementMode defines which mode should kube-proxy run in.
// More infromation here: https://docs.cilium.io/en/v1.7/gettingstarted/kubeproxy-free/
type KubeProxyReplacementMode string

const (
	// Strict defines the strict kube-proxy replacement mode
	Strict KubeProxyReplacementMode = "strict"
	// Probe defines the probe kube-proxy replacement mode
	Probe KubeProxyReplacementMode = "probe"
	// Partial defines the partial kube-proxy replacement mode
	Partial KubeProxyReplacementMode = "partial"
)

// NodePortMode defines how NodePort services are enabled.
type NodePortMode string

const (
	// Hybrid defines the hybrid nodeport mode.
	Hybird NodePortMode = "hybrid"
)

// Store defines the kubernetes storage backend
type Store string

const (
	// Kubernetes defines the kubernetes CRD store type
	Kubernetes Store = "kubernetes"
	// ETCD defines the ETCD store type
	ETCD Store = "etcd"
)

// Hubble enablement for cilium
type Hubble struct {
	// Enabled indicates whether hubble is enabled or not.
	Enabled bool `json:"enabled"`
}

// IPv6 enablement for cilium
type IPv6 struct {
	// Enabled indicates whether IPv6 is enabled or not.
	Enabled bool `json:"enabled"`
}

// Nodeport enablement for cilium
type Nodeport struct {
	// Enabled is used to define whether Nodeport is required or not.
	Enabled bool `json:"nodePortEnabled"`
	// Mode is the mode of NodePort feature
	Mode NodePortMode `json:"nodePortMode"`
}

// KubeProxy configuration for cilium
type KubeProxy struct {
	// Enabled specifies whether kubeproxy is disabled.
	// +optional
	Enabled *bool `json:"disabled"`
	// ServiceHost specify the controlplane node IP Address.
	// +optional
	ServiceHost *string `json:"k8sServiceHost,omitempty"`
	// ServicePort specify the kube-apiserver port number.
	// +optional
	ServicePort *int32 `json:"k8sServicePort,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkConfig is a struct representing the configmap for the cilium
// networking plugin
type NetworkConfig struct {
	metav1.TypeMeta `json:",inline"`
	// Debug configuration to be enabled or not
	// +optional
	Debug *bool `json:"debug,omitempty"`
	// PSPEnabled configuration
	// +optional
	PSPEnabled *bool `json:"psp,omitempty"`
	// KubeProxy configuration to be enabled or not
	// +optional
	KubeProxy *KubeProxy `json:"kubeproxy,omitempty"`
	// Hubble configuration to be enabled or not
	// +optional
	Hubble *Hubble `json:"hubble,omitempty"`
	// TunnelMode configuration, it should be 'vxlan', 'geneve' or 'disabled'
	// +optional
	TunnelMode *TunnelMode `json:"tunnel,omitempty"`
	// Store can be either Kubernetes or etcd.
	// +optional
	Store *Store `json:"store,omitempty"`
	// Enable IPv6
	// +optional
	IPv6 *IPv6 `json:"ipv6,omitempty"`
}
