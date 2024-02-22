// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
)

// NodePortStrategyApplyConfiguration represents an declarative configuration of the NodePortStrategy type for use
// with apply.
type NodePortStrategyApplyConfiguration struct {
	Protocol *v1.IngressControllerProtocol `json:"protocol,omitempty"`
}

// NodePortStrategyApplyConfiguration constructs an declarative configuration of the NodePortStrategy type for use with
// apply.
func NodePortStrategy() *NodePortStrategyApplyConfiguration {
	return &NodePortStrategyApplyConfiguration{}
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *NodePortStrategyApplyConfiguration) WithProtocol(value v1.IngressControllerProtocol) *NodePortStrategyApplyConfiguration {
	b.Protocol = &value
	return b
}
