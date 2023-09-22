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

type AliasInitParameters struct {

	// (String) Query used to limit documents the alias can access.
	// Query used to limit documents the alias can access.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) Value used to route indexing operations to a specific shard. If specified, this overwrites the routing value for indexing operations.
	// Value used to route indexing operations to a specific shard. If specified, this overwrites the routing value for indexing operations.
	IndexRouting *string `json:"indexRouting,omitempty" tf:"index_routing,omitempty"`

	// (Boolean) If true, the alias is hidden.
	// If true, the alias is hidden.
	IsHidden *bool `json:"isHidden,omitempty" tf:"is_hidden,omitempty"`

	// (Boolean) If true, the index is the write index for the alias.
	// If true, the index is the write index for the alias.
	IsWriteIndex *bool `json:"isWriteIndex,omitempty" tf:"is_write_index,omitempty"`

	// (String) Name of the component template to create.
	// The alias name. Index alias names support date math. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/date-math-index-names.html
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Value used to route indexing and search operations to a specific shard.
	// Value used to route indexing and search operations to a specific shard.
	Routing *string `json:"routing,omitempty" tf:"routing,omitempty"`

	// (String) Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	// Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	SearchRouting *string `json:"searchRouting,omitempty" tf:"search_routing,omitempty"`
}

type AliasObservation struct {

	// (String) Query used to limit documents the alias can access.
	// Query used to limit documents the alias can access.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) Value used to route indexing operations to a specific shard. If specified, this overwrites the routing value for indexing operations.
	// Value used to route indexing operations to a specific shard. If specified, this overwrites the routing value for indexing operations.
	IndexRouting *string `json:"indexRouting,omitempty" tf:"index_routing,omitempty"`

	// (Boolean) If true, the alias is hidden.
	// If true, the alias is hidden.
	IsHidden *bool `json:"isHidden,omitempty" tf:"is_hidden,omitempty"`

	// (Boolean) If true, the index is the write index for the alias.
	// If true, the index is the write index for the alias.
	IsWriteIndex *bool `json:"isWriteIndex,omitempty" tf:"is_write_index,omitempty"`

	// (String) Name of the component template to create.
	// The alias name. Index alias names support date math. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/date-math-index-names.html
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Value used to route indexing and search operations to a specific shard.
	// Value used to route indexing and search operations to a specific shard.
	Routing *string `json:"routing,omitempty" tf:"routing,omitempty"`

	// (String) Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	// Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	SearchRouting *string `json:"searchRouting,omitempty" tf:"search_routing,omitempty"`
}

type AliasParameters struct {

	// (String) Query used to limit documents the alias can access.
	// Query used to limit documents the alias can access.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) Value used to route indexing operations to a specific shard. If specified, this overwrites the routing value for indexing operations.
	// Value used to route indexing operations to a specific shard. If specified, this overwrites the routing value for indexing operations.
	// +kubebuilder:validation:Optional
	IndexRouting *string `json:"indexRouting,omitempty" tf:"index_routing,omitempty"`

	// (Boolean) If true, the alias is hidden.
	// If true, the alias is hidden.
	// +kubebuilder:validation:Optional
	IsHidden *bool `json:"isHidden,omitempty" tf:"is_hidden,omitempty"`

	// (Boolean) If true, the index is the write index for the alias.
	// If true, the index is the write index for the alias.
	// +kubebuilder:validation:Optional
	IsWriteIndex *bool `json:"isWriteIndex,omitempty" tf:"is_write_index,omitempty"`

	// (String) Name of the component template to create.
	// The alias name. Index alias names support date math. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/date-math-index-names.html
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) Value used to route indexing and search operations to a specific shard.
	// Value used to route indexing and search operations to a specific shard.
	// +kubebuilder:validation:Optional
	Routing *string `json:"routing,omitempty" tf:"routing,omitempty"`

	// (String) Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	// Value used to route search operations to a specific shard. If specified, this overwrites the routing value for search operations.
	// +kubebuilder:validation:Optional
	SearchRouting *string `json:"searchRouting,omitempty" tf:"search_routing,omitempty"`
}

type ComponentTemplateElasticsearchConnectionInitParameters struct {

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

type ComponentTemplateElasticsearchConnectionObservation struct {

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

type ComponentTemplateElasticsearchConnectionParameters struct {

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

type ComponentTemplateInitParameters struct {

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	ElasticsearchConnection []ComponentTemplateElasticsearchConnectionInitParameters `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (String) Optional user metadata about the component template.
	// Optional user metadata about the component template.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Name of the component template to create.
	// Name of the component template to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) Template to be applied. It may optionally include an aliases, mappings, or settings configuration. (see below for nested schema)
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template []TemplateInitParameters `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Version number used to manage component templates externally.
	// Version number used to manage component templates externally.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ComponentTemplateObservation struct {

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	ElasticsearchConnection []ComponentTemplateElasticsearchConnectionObservation `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (String) Internal identifier of the resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Optional user metadata about the component template.
	// Optional user metadata about the component template.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Name of the component template to create.
	// Name of the component template to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) Template to be applied. It may optionally include an aliases, mappings, or settings configuration. (see below for nested schema)
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template []TemplateObservation `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Version number used to manage component templates externally.
	// Version number used to manage component templates externally.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ComponentTemplateParameters struct {

	// (Block List, Max: 1, Deprecated) Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead. (see below for nested schema)
	// Elasticsearch connection configuration block. This property will be removed in a future provider version. Configure the Elasticsearch connection via the provider configuration instead.
	// +kubebuilder:validation:Optional
	ElasticsearchConnection []ComponentTemplateElasticsearchConnectionParameters `json:"elasticsearchConnection,omitempty" tf:"elasticsearch_connection,omitempty"`

	// (String) Optional user metadata about the component template.
	// Optional user metadata about the component template.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Name of the component template to create.
	// Name of the component template to create.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) Template to be applied. It may optionally include an aliases, mappings, or settings configuration. (see below for nested schema)
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	// +kubebuilder:validation:Optional
	Template []TemplateParameters `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Version number used to manage component templates externally.
	// Version number used to manage component templates externally.
	// +kubebuilder:validation:Optional
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type TemplateInitParameters struct {

	// (Block Set) Alias to add. (see below for nested schema)
	// Alias to add.
	Alias []AliasInitParameters `json:"alias,omitempty" tf:"alias,omitempty"`

	// (String) Mapping for fields in the index.
	// Mapping for fields in the index.
	Mappings *string `json:"mappings,omitempty" tf:"mappings,omitempty"`

	// modules.html#index-modules-settings
	// Configuration options for the index. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules.html#index-modules-settings
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type TemplateObservation struct {

	// (Block Set) Alias to add. (see below for nested schema)
	// Alias to add.
	Alias []AliasObservation `json:"alias,omitempty" tf:"alias,omitempty"`

	// (String) Mapping for fields in the index.
	// Mapping for fields in the index.
	Mappings *string `json:"mappings,omitempty" tf:"mappings,omitempty"`

	// modules.html#index-modules-settings
	// Configuration options for the index. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules.html#index-modules-settings
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type TemplateParameters struct {

	// (Block Set) Alias to add. (see below for nested schema)
	// Alias to add.
	// +kubebuilder:validation:Optional
	Alias []AliasParameters `json:"alias,omitempty" tf:"alias,omitempty"`

	// (String) Mapping for fields in the index.
	// Mapping for fields in the index.
	// +kubebuilder:validation:Optional
	Mappings *string `json:"mappings,omitempty" tf:"mappings,omitempty"`

	// modules.html#index-modules-settings
	// Configuration options for the index. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules.html#index-modules-settings
	// +kubebuilder:validation:Optional
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`
}

// ComponentTemplateSpec defines the desired state of ComponentTemplate
type ComponentTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ComponentTemplateParameters `json:"forProvider"`
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
	InitProvider ComponentTemplateInitParameters `json:"initProvider,omitempty"`
}

// ComponentTemplateStatus defines the observed state of ComponentTemplate.
type ComponentTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ComponentTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ComponentTemplate is the Schema for the ComponentTemplates API. Creates or updates a component template.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticstack}
type ComponentTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.template) || has(self.initProvider.template)",message="template is a required parameter"
	Spec   ComponentTemplateSpec   `json:"spec"`
	Status ComponentTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComponentTemplateList contains a list of ComponentTemplates
type ComponentTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComponentTemplate `json:"items"`
}

// Repository type metadata.
var (
	ComponentTemplate_Kind             = "ComponentTemplate"
	ComponentTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ComponentTemplate_Kind}.String()
	ComponentTemplate_KindAPIVersion   = ComponentTemplate_Kind + "." + CRDGroupVersion.String()
	ComponentTemplate_GroupVersionKind = CRDGroupVersion.WithKind(ComponentTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&ComponentTemplate{}, &ComponentTemplateList{})
}
