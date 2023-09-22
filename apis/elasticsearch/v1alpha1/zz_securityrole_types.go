/*
Copyright Elasticsearch B.V. All rights reserved.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ApplicationsInitParameters struct {

	// (String) The name of the application to which this entry applies.
	// The name of the application to which this entry applies.
	Application *string `json:"application,omitempty" tf:"application,omitempty"`

	// (Set of String) A list of strings, where each element is the name of an application privilege or action.
	// A list of strings, where each element is the name of an application privilege or action.
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// (Set of String) A list resources to which the privileges are applied.
	// A list resources to which the privileges are applied.
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

type ApplicationsObservation struct {

	// (String) The name of the application to which this entry applies.
	// The name of the application to which this entry applies.
	Application *string `json:"application,omitempty" tf:"application,omitempty"`

	// (Set of String) A list of strings, where each element is the name of an application privilege or action.
	// A list of strings, where each element is the name of an application privilege or action.
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// (Set of String) A list resources to which the privileges are applied.
	// A list resources to which the privileges are applied.
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

type ApplicationsParameters struct {

	// (String) The name of the application to which this entry applies.
	// The name of the application to which this entry applies.
	// +kubebuilder:validation:Optional
	Application *string `json:"application" tf:"application,omitempty"`

	// (Set of String) A list of strings, where each element is the name of an application privilege or action.
	// A list of strings, where each element is the name of an application privilege or action.
	// +kubebuilder:validation:Optional
	Privileges []*string `json:"privileges" tf:"privileges,omitempty"`

	// (Set of String) A list resources to which the privileges are applied.
	// A list resources to which the privileges are applied.
	// +kubebuilder:validation:Optional
	Resources []*string `json:"resources" tf:"resources,omitempty"`
}

type FieldSecurityInitParameters struct {

	// (Set of String) List of the fields to which the grants will not be applied.
	// List of the fields to which the grants will not be applied.
	Except []*string `json:"except,omitempty" tf:"except,omitempty"`

	// (Set of String) List of the fields to grant the access to.
	// List of the fields to grant the access to.
	Grant []*string `json:"grant,omitempty" tf:"grant,omitempty"`
}

type FieldSecurityObservation struct {

	// (Set of String) List of the fields to which the grants will not be applied.
	// List of the fields to which the grants will not be applied.
	Except []*string `json:"except,omitempty" tf:"except,omitempty"`

	// (Set of String) List of the fields to grant the access to.
	// List of the fields to grant the access to.
	Grant []*string `json:"grant,omitempty" tf:"grant,omitempty"`
}

type FieldSecurityParameters struct {

	// (Set of String) List of the fields to which the grants will not be applied.
	// List of the fields to which the grants will not be applied.
	// +kubebuilder:validation:Optional
	Except []*string `json:"except,omitempty" tf:"except,omitempty"`

	// (Set of String) List of the fields to grant the access to.
	// List of the fields to grant the access to.
	// +kubebuilder:validation:Optional
	Grant []*string `json:"grant,omitempty" tf:"grant,omitempty"`
}

type SecurityRoleElasticsearchConnectionInitParameters struct {

	// encoded custom Certificate Authority certificate
	// PEM-encoded custom Certificate Authority certificate
	CAData *string `json:"caData,omitempty" tf:"ca_data,omitempty"`

	// (String) Path to a custom Certificate Authority certificate
	// Path to a custom Certificate Authority certificate
	CAFile *string `json:"caFile,omitempty" tf:"ca_file,omitempty"`

	// (String) PEM encoded certificate for client auth
	// PEM encoded certificate for client auth
	CertData *string `json:"certData,omitempty" tf:"cert_data,omitempty"`

	// (String) Path to a file containing the PEM encoded certificate for client auth
	// Path to a file containing the PEM encoded certificate for client auth
	CertFile *string `json:"certFile,omitempty" tf:"cert_file,omitempty"`

	// (Boolean) Disable TLS certificate validation
	// Disable TLS certificate validation
	Insecure *bool `json:"insecure,omitempty" tf:"insecure,omitempty"`

	// (String) Path to a file containing the PEM encoded private key for client auth
	// Path to a file containing the PEM encoded private key for client auth
	KeyFile *string `json:"keyFile,omitempty" tf:"key_file,omitempty"`

	// (String) Username to use for API authentication to Elasticsearch.
	// Username to use for API authentication to Elasticsearch.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SecurityRoleElasticsearchConnectionObservation struct {

	// encoded custom Certificate Authority certificate
	// PEM-encoded custom Certificate Authority certificate
	CAData *string `json:"caData,omitempty" tf:"ca_data,omitempty"`

	// (String) Path to a custom Certificate Authority certificate
	// Path to a custom Certificate Authority certificate
	CAFile *string `json:"caFile,omitempty" tf:"ca_file,omitempty"`

	// (String) PEM encoded certificate for client auth
	// PEM encoded certificate for client auth
	CertData *string `json:"certData,omitempty" tf:"cert_data,omitempty"`

	// (String) Path to a file containing the PEM encoded certificate for client auth
	// Path to a file containing the PEM encoded certificate for client auth
	CertFile *string `json:"certFile,omitempty" tf:"cert_file,omitempty"`

	// (Boolean) Disable TLS certificate validation
	// Disable TLS certificate validation
	Insecure *bool `json:"insecure,omitempty" tf:"insecure,omitempty"`

	// (String) Path to a file containing the PEM encoded private key for client auth
	// Path to a file containing the PEM encoded private key for client auth
	KeyFile *string `json:"keyFile,omitempty" tf:"key_file,omitempty"`

	// (String) Username to use for API authentication to Elasticsearch.
	// Username to use for API authentication to Elasticsearch.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SecurityRoleElasticsearchConnectionParameters struct {

	// (String, Sensitive) API Key to use for authentication to Elasticsearch
	// API Key to use for authentication to Elasticsearch
	// +kubebuilder:validation:Optional
	APIKeySecretRef *v1.SecretKeySelector `json:"apiKeySecretRef,omitempty" tf:"-"`

	// encoded custom Certificate Authority certificate
	// PEM-encoded custom Certificate Authority certificate
	// +kubebuilder:validation:Optional
	CAData *string `json:"caData,omitempty" tf:"ca_data,omitempty"`

	// (String) Path to a custom Certificate Authority certificate
	// Path to a custom Certificate Authority certificate
	// +kubebuilder:validation:Optional
	CAFile *string `json:"caFile,omitempty" tf:"ca_file,omitempty"`

	// (String) PEM encoded certificate for client auth
	// PEM encoded certificate for client auth
	// +kubebuilder:validation:Optional
	CertData *string `json:"certData,omitempty" tf:"cert_data,omitempty"`

	// (String) Path to a file containing the PEM encoded certificate for client auth
	// Path to a file containing the PEM encoded certificate for client auth
	// +kubebuilder:validation:Optional
	CertFile *string `json:"certFile,omitempty" tf:"cert_file,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointsSecretRef *[]v1.SecretKeySelector `json:"endpointsSecretRef,omitempty" tf:"-"`

	// (Boolean) Disable TLS certificate validation
	// Disable TLS certificate validation
	// +kubebuilder:validation:Optional
	Insecure *bool `json:"insecure,omitempty" tf:"insecure,omitempty"`

	// (String, Sensitive) PEM encoded private key for client auth
	// PEM encoded private key for client auth
	// +kubebuilder:validation:Optional
	KeyDataSecretRef *v1.SecretKeySelector `json:"keyDataSecretRef,omitempty" tf:"-"`

	// (String) Path to a file containing the PEM encoded private key for client auth
	// Path to a file containing the PEM encoded private key for client auth
	// +kubebuilder:validation:Optional
	KeyFile *string `json:"keyFile,omitempty" tf:"key_file,omitempty"`

	// (String, Sensitive) Password to use for API authentication to Elasticsearch.
	// Password to use for API authentication to Elasticsearch.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// (String) Username to use for API authentication to Elasticsearch.
	// Username to use for API authentication to Elasticsearch.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SecurityRoleIndicesInitParameters struct {

	// (Boolean) Include matching restricted indices in names parameter. Usage is strongly discouraged as it can grant unrestricted operations on critical data, make the entire system unstable or leak sensitive information.
	// Include matching restricted indices in names parameter. Usage is strongly discouraged as it can grant unrestricted operations on critical data, make the entire system unstable or leak sensitive information.
	AllowRestrictedIndices *bool `json:"allowRestrictedIndices,omitempty" tf:"allow_restricted_indices,omitempty"`

	// (Block List, Max: 1) The document fields that the owners of the role have read access to. (see below for nested schema)
	// The document fields that the owners of the role have read access to.
	FieldSecurity []FieldSecurityInitParameters `json:"fieldSecurity,omitempty" tf:"field_security,omitempty"`

	// (Set of String) A list of indices (or index name patterns) to which the permissions in this entry apply.
	// A list of indices (or index name patterns) to which the permissions in this entry apply.
	Names []*string `json:"names,omitempty" tf:"names,omitempty"`

	// (Set of String) A list of strings, where each element is the name of an application privilege or action.
	// The index level privileges that the owners of the role have on the specified indices.
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// (String) A search query that defines the documents the owners of the role have read access to.
	// A search query that defines the documents the owners of the role have read access to.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type SecurityRoleIndicesObservation struct {

	// (Boolean) Include matching restricted indices in names parameter. Usage is strongly discouraged as it can grant unrestricted operations on critical data, make the entire system unstable or leak sensitive information.
	// Include matching restricted indices in names parameter. Usage is strongly discouraged as it can grant unrestricted operations on critical data, make the entire system unstable or leak sensitive information.
	AllowRestrictedIndices *bool `json:"allowRestrictedIndices,omitempty" tf:"allow_restricted_indices,omitempty"`

	// (Block List, Max: 1) The document fields that the owners of the role have read access to. (see below for nested schema)
	// The document fields that the owners of the role have read access to.
	FieldSecurity []FieldSecurityObservation `json:"fieldSecurity,omitempty" tf:"field_security,omitempty"`

	// (Set of String) A list of indices (or index name patterns) to which the permissions in this entry apply.
	// A list of indices (or index name patterns) to which the permissions in this entry apply.
	Names []*string `json:"names,omitempty" tf:"names,omitempty"`

	// (Set of String) A list of strings, where each element is the name of an application privilege or action.
	// The index level privileges that the owners of the role have on the specified indices.
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// (String) A search query that defines the documents the owners of the role have read access to.
	// A search query that defines the documents the owners of the role have read access to.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type SecurityRoleIndicesParameters struct {

	// (Boolean) Include matching restricted indices in names parameter. Usage is strongly discouraged as it can grant unrestricted operations on critical data, make the entire system unstable or leak sensitive information.
	// Include matching restricted indices in names parameter. Usage is strongly discouraged as it can grant unrestricted operations on critical data, make the entire system unstable or leak sensitive information.
	// +kubebuilder:validation:Optional
	AllowRestrictedIndices *bool `json:"allowRestrictedIndices,omitempty" tf:"allow_restricted_indices,omitempty"`

	// (Block List, Max: 1) The document fields that the owners of the role have read access to. (see below for nested schema)
	// The document fields that the owners of the role have read access to.
	// +kubebuilder:validation:Optional
	FieldSecurity []FieldSecurityParameters `json:"fieldSecurity,omitempty" tf:"field_security,omitempty"`

	// (Set of String) A list of indices (or index name patterns) to which the permissions in this entry apply.
	// A list of indices (or index name patterns) to which the permissions in this entry apply.
	// +kubebuilder:validation:Optional
	Names []*string `json:"names" tf:"names,omitempty"`

	// (Set of String) A list of strings, where each element is the name of an application privilege or action.
	// The index level privileges that the owners of the role have on the specified indices.
	// +kubebuilder:validation:Optional
	Privileges []*string `json:"privileges" tf:"privileges,omitempty"`

	// (String) A search query that defines the documents the owners of the role have read access to.
	// A search query that defines the documents the owners of the role have read access to.
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type SecurityRoleInitParameters struct {

	// (Block Set) A list of application privilege entries. (see below for nested schema)
	// A list of application privilege entries.
	Applications []ApplicationsInitParameters `json:"applications,omitempty" tf:"applications,omitempty"`

	// (Set of String) A list of cluster privileges. These privileges define the cluster level actions that users with this role are able to execute.
	// A list of cluster privileges. These privileges define the cluster level actions that users with this role are able to execute.
	Cluster []*string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	ElasticsearchConnection []SecurityRoleElasticsearchConnectionInitParameters `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (String) An object defining global privileges.
	// An object defining global privileges.
	Global *string `json:"global,omitempty" tf:"global,omitempty"`

	// (Block Set) A list of indices permissions entries. (see below for nested schema)
	// A list of indices permissions entries.
	Indices []SecurityRoleIndicesInitParameters `json:"indices,omitempty" tf:"indices,omitempty"`

	// data.
	// Optional meta-data.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) The name of the role.
	// The name of the role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list of users that the owners of this role can impersonate.
	// A list of users that the owners of this role can impersonate.
	RunAs []*string `json:"runAs,omitempty" tf:"run_as,omitempty"`
}

type SecurityRoleObservation struct {

	// (Block Set) A list of application privilege entries. (see below for nested schema)
	// A list of application privilege entries.
	Applications []ApplicationsObservation `json:"applications,omitempty" tf:"applications,omitempty"`

	// (Set of String) A list of cluster privileges. These privileges define the cluster level actions that users with this role are able to execute.
	// A list of cluster privileges. These privileges define the cluster level actions that users with this role are able to execute.
	Cluster []*string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	ElasticsearchConnection []SecurityRoleElasticsearchConnectionObservation `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (String) An object defining global privileges.
	// An object defining global privileges.
	Global *string `json:"global,omitempty" tf:"global,omitempty"`

	// (String) Internal identifier of the resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block Set) A list of indices permissions entries. (see below for nested schema)
	// A list of indices permissions entries.
	Indices []SecurityRoleIndicesObservation `json:"indices,omitempty" tf:"indices,omitempty"`

	// data.
	// Optional meta-data.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) The name of the role.
	// The name of the role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list of users that the owners of this role can impersonate.
	// A list of users that the owners of this role can impersonate.
	RunAs []*string `json:"runAs,omitempty" tf:"run_as,omitempty"`
}

type SecurityRoleParameters struct {

	// (Block Set) A list of application privilege entries. (see below for nested schema)
	// A list of application privilege entries.
	// +kubebuilder:validation:Optional
	Applications []ApplicationsParameters `json:"applications,omitempty" tf:"applications,omitempty"`

	// (Set of String) A list of cluster privileges. These privileges define the cluster level actions that users with this role are able to execute.
	// A list of cluster privileges. These privileges define the cluster level actions that users with this role are able to execute.
	// +kubebuilder:validation:Optional
	Cluster []*string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	// +kubebuilder:validation:Optional
	ElasticsearchConnection []SecurityRoleElasticsearchConnectionParameters `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (String) An object defining global privileges.
	// An object defining global privileges.
	// +kubebuilder:validation:Optional
	Global *string `json:"global,omitempty" tf:"global,omitempty"`

	// (Block Set) A list of indices permissions entries. (see below for nested schema)
	// A list of indices permissions entries.
	// +kubebuilder:validation:Optional
	Indices []SecurityRoleIndicesParameters `json:"indices,omitempty" tf:"indices,omitempty"`

	// data.
	// Optional meta-data.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) The name of the role.
	// The name of the role.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list of users that the owners of this role can impersonate.
	// A list of users that the owners of this role can impersonate.
	// +kubebuilder:validation:Optional
	RunAs []*string `json:"runAs,omitempty" tf:"run_as,omitempty"`
}

// SecurityRoleSpec defines the desired state of SecurityRole
type SecurityRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityRoleParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SecurityRoleInitParameters `json:"initProvider,omitempty"`
}

// SecurityRoleStatus defines the observed state of SecurityRole.
type SecurityRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityRole is the Schema for the SecurityRoles API. Adds and updates roles in the native realm.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticstack}
type SecurityRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   SecurityRoleSpec   `json:"spec"`
	Status SecurityRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityRoleList contains a list of SecurityRoles
type SecurityRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityRole `json:"items"`
}

// Repository type metadata.
var (
	SecurityRole_Kind             = "SecurityRole"
	SecurityRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityRole_Kind}.String()
	SecurityRole_KindAPIVersion   = SecurityRole_Kind + "." + CRDGroupVersion.String()
	SecurityRole_GroupVersionKind = CRDGroupVersion.WithKind(SecurityRole_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityRole{}, &SecurityRoleList{})
}
