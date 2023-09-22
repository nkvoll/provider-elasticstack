/*
Copyright Elasticsearch B.V. All rights reserved.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import resource "github.com/crossplane/crossplane-runtime/pkg/resource"

// GetItems of this AgentPolicyList.
func (l *AgentPolicyList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this OutputList.
func (l *OutputList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this ServerHostList.
func (l *ServerHostList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}
