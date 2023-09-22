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

type ActionsInitParameters struct {

	// (String) The group name, which affects when the action runs (for example, when the threshold is met or when the alert is recovered). Each rule type has a list of valid action group names.
	// The group name, which affects when the action runs (for example, when the threshold is met or when the alert is recovered). Each rule type has a list of valid action group names.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// (String) The rule parameters, which differ for each rule type.
	// The parameters for the action, which are sent to the connector.
	Params *string `json:"params,omitempty" tf:"params,omitempty"`
}

type ActionsObservation struct {

	// (String) The group name, which affects when the action runs (for example, when the threshold is met or when the alert is recovered). Each rule type has a list of valid action group names.
	// The group name, which affects when the action runs (for example, when the threshold is met or when the alert is recovered). Each rule type has a list of valid action group names.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// (String) The ID of this resource.
	// The identifier for the connector saved object.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The rule parameters, which differ for each rule type.
	// The parameters for the action, which are sent to the connector.
	Params *string `json:"params,omitempty" tf:"params,omitempty"`
}

type ActionsParameters struct {

	// (String) The group name, which affects when the action runs (for example, when the threshold is met or when the alert is recovered). Each rule type has a list of valid action group names.
	// The group name, which affects when the action runs (for example, when the threshold is met or when the alert is recovered). Each rule type has a list of valid action group names.
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// (String) The ID of this resource.
	// The identifier for the connector saved object.
	// +crossplane:generate:reference:type=github.com/elastic/provider-elasticstack/apis/kibana/v1alpha1.ActionConnector
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("connector_id",true)
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a ActionConnector in kibana to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a ActionConnector in kibana to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// (String) The rule parameters, which differ for each rule type.
	// The parameters for the action, which are sent to the connector.
	// +kubebuilder:validation:Optional
	Params *string `json:"params" tf:"params,omitempty"`
}

type AlertingRuleInitParameters struct {

	// (Block List) An action that runs under defined conditions. (see below for nested schema)
	// An action that runs under defined conditions.
	Actions []ActionsInitParameters `json:"actions,omitempty" tf:"actions,omitempty"`

	// (String) The name of the application or feature that owns the rule.
	// The name of the application or feature that owns the rule.
	Consumer *string `json:"consumer,omitempty" tf:"consumer,omitempty"`

	// (Boolean) Indicates if you want to run the rule on an interval basis.
	// Indicates if you want to run the rule on an interval basis.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The check interval, which specifies how frequently the rule conditions are checked. The interval must be specified in seconds, minutes, hours or days.
	// The check interval, which specifies how frequently the rule conditions are checked. The interval must be specified in seconds, minutes, hours or days.
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// (String) The name of the rule. While this name does not have to be unique, a distinctive name can help you identify a rule.
	// The name of the rule. While this name does not have to be unique, a distinctive name can help you identify a rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// specific notify_when values.
	// Defines how often alerts generate actions. Valid values include: `onActionGroupChange`: Actions run when the alert status changes; `onActiveAlert`: Actions run when the alert becomes active and at each check interval while the rule conditions are met; `onThrottleInterval`: Actions run when the alert becomes active and at the interval specified in the throttle property while the rule conditions are met. NOTE: This is a rule level property; if you update the rule in Kibana, it is automatically changed to use action-specific `notify_when` values.
	NotifyWhen *string `json:"notifyWhen,omitempty" tf:"notify_when,omitempty"`

	// (String) The rule parameters, which differ for each rule type.
	// The rule parameters, which differ for each rule type.
	Params *string `json:"params,omitempty" tf:"params,omitempty"`

	// (String) A UUID v1 or v4 to use instead of a randomly generated ID.
	// A UUID v1 or v4 to use instead of a randomly generated ID.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// (String) The ID of the rule type that you want to call when the rule is scheduled to run. For more information about the valid values, list the rule types using Get rule types API or refer to the Rule types documentation.
	// The ID of the rule type that you want to call when the rule is scheduled to run. For more information about the valid values, list the rule types using [Get rule types API](https://www.elastic.co/guide/en/kibana/master/list-rule-types-api.html) or refer to the [Rule types documentation](https://www.elastic.co/guide/en/kibana/master/rule-types.html).
	RuleTypeID *string `json:"ruleTypeId,omitempty" tf:"rule_type_id,omitempty"`

	// (List of String) A list of tag names that are applied to the rule.
	// A list of tag names that are applied to the rule.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// specific throttle values.
	// Defines how often an alert generates repeated actions. This custom action interval must be specified in seconds, minutes, hours, or days. For example, 10m or 1h. This property is applicable only if `notify_when` is `onThrottleInterval`. NOTE: This is a rule level property; if you update the rule in Kibana, it is automatically changed to use action-specific `throttle` values.
	Throttle *string `json:"throttle,omitempty" tf:"throttle,omitempty"`
}

type AlertingRuleObservation struct {

	// (Block List) An action that runs under defined conditions. (see below for nested schema)
	// An action that runs under defined conditions.
	Actions []ActionsObservation `json:"actions,omitempty" tf:"actions,omitempty"`

	// (String) The name of the application or feature that owns the rule.
	// The name of the application or feature that owns the rule.
	Consumer *string `json:"consumer,omitempty" tf:"consumer,omitempty"`

	// (Boolean) Indicates if you want to run the rule on an interval basis.
	// Indicates if you want to run the rule on an interval basis.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The check interval, which specifies how frequently the rule conditions are checked. The interval must be specified in seconds, minutes, hours or days.
	// The check interval, which specifies how frequently the rule conditions are checked. The interval must be specified in seconds, minutes, hours or days.
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// (String) Date of the last execution of this rule.
	// Date of the last execution of this rule.
	LastExecutionDate *string `json:"lastExecutionDate,omitempty" tf:"last_execution_date,omitempty"`

	// (String) Status of the last execution of this rule.
	// Status of the last execution of this rule.
	LastExecutionStatus *string `json:"lastExecutionStatus,omitempty" tf:"last_execution_status,omitempty"`

	// (String) The name of the rule. While this name does not have to be unique, a distinctive name can help you identify a rule.
	// The name of the rule. While this name does not have to be unique, a distinctive name can help you identify a rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// specific notify_when values.
	// Defines how often alerts generate actions. Valid values include: `onActionGroupChange`: Actions run when the alert status changes; `onActiveAlert`: Actions run when the alert becomes active and at each check interval while the rule conditions are met; `onThrottleInterval`: Actions run when the alert becomes active and at the interval specified in the throttle property while the rule conditions are met. NOTE: This is a rule level property; if you update the rule in Kibana, it is automatically changed to use action-specific `notify_when` values.
	NotifyWhen *string `json:"notifyWhen,omitempty" tf:"notify_when,omitempty"`

	// (String) The rule parameters, which differ for each rule type.
	// The rule parameters, which differ for each rule type.
	Params *string `json:"params,omitempty" tf:"params,omitempty"`

	// (String) A UUID v1 or v4 to use instead of a randomly generated ID.
	// A UUID v1 or v4 to use instead of a randomly generated ID.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// (String) The ID of the rule type that you want to call when the rule is scheduled to run. For more information about the valid values, list the rule types using Get rule types API or refer to the Rule types documentation.
	// The ID of the rule type that you want to call when the rule is scheduled to run. For more information about the valid values, list the rule types using [Get rule types API](https://www.elastic.co/guide/en/kibana/master/list-rule-types-api.html) or refer to the [Rule types documentation](https://www.elastic.co/guide/en/kibana/master/rule-types.html).
	RuleTypeID *string `json:"ruleTypeId,omitempty" tf:"rule_type_id,omitempty"`

	// (String) ID of the scheduled task that will execute the alert.
	// ID of the scheduled task that will execute the alert.
	ScheduledTaskID *string `json:"scheduledTaskId,omitempty" tf:"scheduled_task_id,omitempty"`

	// (String) An identifier for the space. If space_id is not provided, the default space is used.
	// An identifier for the space. If space_id is not provided, the default space is used.
	SpaceID *string `json:"spaceId,omitempty" tf:"space_id,omitempty"`

	// (List of String) A list of tag names that are applied to the rule.
	// A list of tag names that are applied to the rule.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// specific throttle values.
	// Defines how often an alert generates repeated actions. This custom action interval must be specified in seconds, minutes, hours, or days. For example, 10m or 1h. This property is applicable only if `notify_when` is `onThrottleInterval`. NOTE: This is a rule level property; if you update the rule in Kibana, it is automatically changed to use action-specific `throttle` values.
	Throttle *string `json:"throttle,omitempty" tf:"throttle,omitempty"`
}

type AlertingRuleParameters struct {

	// (Block List) An action that runs under defined conditions. (see below for nested schema)
	// An action that runs under defined conditions.
	// +kubebuilder:validation:Optional
	Actions []ActionsParameters `json:"actions,omitempty" tf:"actions,omitempty"`

	// (String) The name of the application or feature that owns the rule.
	// The name of the application or feature that owns the rule.
	// +kubebuilder:validation:Optional
	Consumer *string `json:"consumer,omitempty" tf:"consumer,omitempty"`

	// (Boolean) Indicates if you want to run the rule on an interval basis.
	// Indicates if you want to run the rule on an interval basis.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The check interval, which specifies how frequently the rule conditions are checked. The interval must be specified in seconds, minutes, hours or days.
	// The check interval, which specifies how frequently the rule conditions are checked. The interval must be specified in seconds, minutes, hours or days.
	// +kubebuilder:validation:Optional
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// (String) The name of the rule. While this name does not have to be unique, a distinctive name can help you identify a rule.
	// The name of the rule. While this name does not have to be unique, a distinctive name can help you identify a rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// specific notify_when values.
	// Defines how often alerts generate actions. Valid values include: `onActionGroupChange`: Actions run when the alert status changes; `onActiveAlert`: Actions run when the alert becomes active and at each check interval while the rule conditions are met; `onThrottleInterval`: Actions run when the alert becomes active and at the interval specified in the throttle property while the rule conditions are met. NOTE: This is a rule level property; if you update the rule in Kibana, it is automatically changed to use action-specific `notify_when` values.
	// +kubebuilder:validation:Optional
	NotifyWhen *string `json:"notifyWhen,omitempty" tf:"notify_when,omitempty"`

	// (String) The rule parameters, which differ for each rule type.
	// The rule parameters, which differ for each rule type.
	// +kubebuilder:validation:Optional
	Params *string `json:"params,omitempty" tf:"params,omitempty"`

	// (String) A UUID v1 or v4 to use instead of a randomly generated ID.
	// A UUID v1 or v4 to use instead of a randomly generated ID.
	// +kubebuilder:validation:Optional
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// (String) The ID of the rule type that you want to call when the rule is scheduled to run. For more information about the valid values, list the rule types using Get rule types API or refer to the Rule types documentation.
	// The ID of the rule type that you want to call when the rule is scheduled to run. For more information about the valid values, list the rule types using [Get rule types API](https://www.elastic.co/guide/en/kibana/master/list-rule-types-api.html) or refer to the [Rule types documentation](https://www.elastic.co/guide/en/kibana/master/rule-types.html).
	// +kubebuilder:validation:Optional
	RuleTypeID *string `json:"ruleTypeId,omitempty" tf:"rule_type_id,omitempty"`

	// (String) An identifier for the space. If space_id is not provided, the default space is used.
	// An identifier for the space. If space_id is not provided, the default space is used.
	// +crossplane:generate:reference:type=github.com/elastic/provider-elasticstack/apis/kibana/v1alpha1.Space
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("space_id",true)
	// +kubebuilder:validation:Optional
	SpaceID *string `json:"spaceId,omitempty" tf:"space_id,omitempty"`

	// Reference to a Space in kibana to populate spaceId.
	// +kubebuilder:validation:Optional
	SpaceIDRef *v1.Reference `json:"spaceIdRef,omitempty" tf:"-"`

	// Selector for a Space in kibana to populate spaceId.
	// +kubebuilder:validation:Optional
	SpaceIDSelector *v1.Selector `json:"spaceIdSelector,omitempty" tf:"-"`

	// (List of String) A list of tag names that are applied to the rule.
	// A list of tag names that are applied to the rule.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// specific throttle values.
	// Defines how often an alert generates repeated actions. This custom action interval must be specified in seconds, minutes, hours, or days. For example, 10m or 1h. This property is applicable only if `notify_when` is `onThrottleInterval`. NOTE: This is a rule level property; if you update the rule in Kibana, it is automatically changed to use action-specific `throttle` values.
	// +kubebuilder:validation:Optional
	Throttle *string `json:"throttle,omitempty" tf:"throttle,omitempty"`
}

// AlertingRuleSpec defines the desired state of AlertingRule
type AlertingRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlertingRuleParameters `json:"forProvider"`
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
	InitProvider AlertingRuleInitParameters `json:"initProvider,omitempty"`
}

// AlertingRuleStatus defines the observed state of AlertingRule.
type AlertingRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlertingRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlertingRule is the Schema for the AlertingRules API. Creates or updates a Kibana alerting rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticstack}
type AlertingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.consumer) || has(self.initProvider.consumer)",message="consumer is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.interval) || has(self.initProvider.interval)",message="interval is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.notifyWhen) || has(self.initProvider.notifyWhen)",message="notifyWhen is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.params) || has(self.initProvider.params)",message="params is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ruleTypeId) || has(self.initProvider.ruleTypeId)",message="ruleTypeId is a required parameter"
	Spec   AlertingRuleSpec   `json:"spec"`
	Status AlertingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertingRuleList contains a list of AlertingRules
type AlertingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertingRule `json:"items"`
}

// Repository type metadata.
var (
	AlertingRule_Kind             = "AlertingRule"
	AlertingRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlertingRule_Kind}.String()
	AlertingRule_KindAPIVersion   = AlertingRule_Kind + "." + CRDGroupVersion.String()
	AlertingRule_GroupVersionKind = CRDGroupVersion.WithKind(AlertingRule_Kind)
)

func init() {
	SchemeBuilder.Register(&AlertingRule{}, &AlertingRuleList{})
}
