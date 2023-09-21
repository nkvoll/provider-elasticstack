/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ApmAvailabilityIndicatorInitParameters struct {

	// (String)
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String)
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String)
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// (String)
	TransactionName *string `json:"transactionName,omitempty" tf:"transaction_name,omitempty"`

	// (String)
	TransactionType *string `json:"transactionType,omitempty" tf:"transaction_type,omitempty"`
}

type ApmAvailabilityIndicatorObservation struct {

	// (String)
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String)
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String)
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// (String)
	TransactionName *string `json:"transactionName,omitempty" tf:"transaction_name,omitempty"`

	// (String)
	TransactionType *string `json:"transactionType,omitempty" tf:"transaction_type,omitempty"`
}

type ApmAvailabilityIndicatorParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment" tf:"environment,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Index *string `json:"index" tf:"index,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Service *string `json:"service" tf:"service,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	TransactionName *string `json:"transactionName" tf:"transaction_name,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	TransactionType *string `json:"transactionType" tf:"transaction_type,omitempty"`
}

type ApmLatencyIndicatorInitParameters struct {

	// (String)
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String)
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String)
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// (Number)
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// (String)
	TransactionName *string `json:"transactionName,omitempty" tf:"transaction_name,omitempty"`

	// (String)
	TransactionType *string `json:"transactionType,omitempty" tf:"transaction_type,omitempty"`
}

type ApmLatencyIndicatorObservation struct {

	// (String)
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String)
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String)
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// (Number)
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// (String)
	TransactionName *string `json:"transactionName,omitempty" tf:"transaction_name,omitempty"`

	// (String)
	TransactionType *string `json:"transactionType,omitempty" tf:"transaction_type,omitempty"`
}

type ApmLatencyIndicatorParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment" tf:"environment,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Index *string `json:"index" tf:"index,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Service *string `json:"service" tf:"service,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	TransactionName *string `json:"transactionName" tf:"transaction_name,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	TransactionType *string `json:"transactionType" tf:"transaction_type,omitempty"`
}

type GoodInitParameters struct {

	// (String)
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation,omitempty"`

	// (String)
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Number)
	From *float64 `json:"from,omitempty" tf:"from,omitempty"`

	// (Number)
	To *float64 `json:"to,omitempty" tf:"to,omitempty"`
}

type GoodObservation struct {

	// (String)
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation,omitempty"`

	// (String)
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Number)
	From *float64 `json:"from,omitempty" tf:"from,omitempty"`

	// (Number)
	To *float64 `json:"to,omitempty" tf:"to,omitempty"`
}

type GoodParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Aggregation *string `json:"aggregation" tf:"aggregation,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Field *string `json:"field" tf:"field,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	From *float64 `json:"from,omitempty" tf:"from,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	To *float64 `json:"to,omitempty" tf:"to,omitempty"`
}

type HistogramCustomIndicatorInitParameters struct {

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Good []GoodInitParameters `json:"good,omitempty" tf:"good,omitempty"`

	// (String)
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String)
	TimestampField *string `json:"timestampField,omitempty" tf:"timestamp_field,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Total []TotalInitParameters `json:"total,omitempty" tf:"total,omitempty"`
}

type HistogramCustomIndicatorObservation struct {

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Good []GoodObservation `json:"good,omitempty" tf:"good,omitempty"`

	// (String)
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String)
	TimestampField *string `json:"timestampField,omitempty" tf:"timestamp_field,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Total []TotalObservation `json:"total,omitempty" tf:"total,omitempty"`
}

type HistogramCustomIndicatorParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Good []GoodParameters `json:"good" tf:"good,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Index *string `json:"index" tf:"index,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	TimestampField *string `json:"timestampField,omitempty" tf:"timestamp_field,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Total []TotalParameters `json:"total" tf:"total,omitempty"`
}

type KqlCustomIndicatorInitParameters struct {

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Good *string `json:"good,omitempty" tf:"good,omitempty"`

	// (String)
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String)
	TimestampField *string `json:"timestampField,omitempty" tf:"timestamp_field,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Total *string `json:"total,omitempty" tf:"total,omitempty"`
}

type KqlCustomIndicatorObservation struct {

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Good *string `json:"good,omitempty" tf:"good,omitempty"`

	// (String)
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String)
	TimestampField *string `json:"timestampField,omitempty" tf:"timestamp_field,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Total *string `json:"total,omitempty" tf:"total,omitempty"`
}

type KqlCustomIndicatorParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Good *string `json:"good,omitempty" tf:"good,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Index *string `json:"index" tf:"index,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	TimestampField *string `json:"timestampField,omitempty" tf:"timestamp_field,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Total *string `json:"total,omitempty" tf:"total,omitempty"`
}

type MetricCustomIndicatorGoodInitParameters struct {

	// (String)
	Equation *string `json:"equation,omitempty" tf:"equation,omitempty"`

	// (Block List, Min: 1) (see below for nested schema)
	Metrics []MetricsInitParameters `json:"metrics,omitempty" tf:"metrics,omitempty"`
}

type MetricCustomIndicatorGoodObservation struct {

	// (String)
	Equation *string `json:"equation,omitempty" tf:"equation,omitempty"`

	// (Block List, Min: 1) (see below for nested schema)
	Metrics []MetricsObservation `json:"metrics,omitempty" tf:"metrics,omitempty"`
}

type MetricCustomIndicatorGoodParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Equation *string `json:"equation" tf:"equation,omitempty"`

	// (Block List, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Metrics []MetricsParameters `json:"metrics" tf:"metrics,omitempty"`
}

type MetricCustomIndicatorInitParameters struct {

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Good []MetricCustomIndicatorGoodInitParameters `json:"good,omitempty" tf:"good,omitempty"`

	// (String)
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String)
	TimestampField *string `json:"timestampField,omitempty" tf:"timestamp_field,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Total []MetricCustomIndicatorTotalInitParameters `json:"total,omitempty" tf:"total,omitempty"`
}

type MetricCustomIndicatorObservation struct {

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Good []MetricCustomIndicatorGoodObservation `json:"good,omitempty" tf:"good,omitempty"`

	// (String)
	Index *string `json:"index,omitempty" tf:"index,omitempty"`

	// (String)
	TimestampField *string `json:"timestampField,omitempty" tf:"timestamp_field,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Total []MetricCustomIndicatorTotalObservation `json:"total,omitempty" tf:"total,omitempty"`
}

type MetricCustomIndicatorParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Good []MetricCustomIndicatorGoodParameters `json:"good" tf:"good,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Index *string `json:"index" tf:"index,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	TimestampField *string `json:"timestampField,omitempty" tf:"timestamp_field,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Total []MetricCustomIndicatorTotalParameters `json:"total" tf:"total,omitempty"`
}

type MetricCustomIndicatorTotalInitParameters struct {

	// (String)
	Equation *string `json:"equation,omitempty" tf:"equation,omitempty"`

	// (Block List, Min: 1) (see below for nested schema)
	Metrics []TotalMetricsInitParameters `json:"metrics,omitempty" tf:"metrics,omitempty"`
}

type MetricCustomIndicatorTotalObservation struct {

	// (String)
	Equation *string `json:"equation,omitempty" tf:"equation,omitempty"`

	// (Block List, Min: 1) (see below for nested schema)
	Metrics []TotalMetricsObservation `json:"metrics,omitempty" tf:"metrics,omitempty"`
}

type MetricCustomIndicatorTotalParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Equation *string `json:"equation" tf:"equation,omitempty"`

	// (Block List, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Metrics []TotalMetricsParameters `json:"metrics" tf:"metrics,omitempty"`
}

type MetricsInitParameters struct {

	// (String)
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation,omitempty"`

	// (String)
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) The name of the SLO.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type MetricsObservation struct {

	// (String)
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation,omitempty"`

	// (String)
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) The name of the SLO.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type MetricsParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Aggregation *string `json:"aggregation" tf:"aggregation,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Field *string `json:"field" tf:"field,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) The name of the SLO.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type ObjectiveInitParameters struct {

	// (Number)
	Target *float64 `json:"target,omitempty" tf:"target,omitempty"`

	// (Number)
	TimesliceTarget *float64 `json:"timesliceTarget,omitempty" tf:"timeslice_target,omitempty"`

	// (String)
	TimesliceWindow *string `json:"timesliceWindow,omitempty" tf:"timeslice_window,omitempty"`
}

type ObjectiveObservation struct {

	// (Number)
	Target *float64 `json:"target,omitempty" tf:"target,omitempty"`

	// (Number)
	TimesliceTarget *float64 `json:"timesliceTarget,omitempty" tf:"timeslice_target,omitempty"`

	// (String)
	TimesliceWindow *string `json:"timesliceWindow,omitempty" tf:"timeslice_window,omitempty"`
}

type ObjectiveParameters struct {

	// (Number)
	// +kubebuilder:validation:Optional
	Target *float64 `json:"target" tf:"target,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	TimesliceTarget *float64 `json:"timesliceTarget,omitempty" tf:"timeslice_target,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	TimesliceWindow *string `json:"timesliceWindow,omitempty" tf:"timeslice_window,omitempty"`
}

type SLOInitParameters struct {

	// (Block List, Max: 1) (see below for nested schema)
	ApmAvailabilityIndicator []ApmAvailabilityIndicatorInitParameters `json:"apmAvailabilityIndicator,omitempty" tf:"apm_availability_indicator,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	ApmLatencyIndicator []ApmLatencyIndicatorInitParameters `json:"apmLatencyIndicator,omitempty" tf:"apm_latency_indicator,omitempty"`

	// (String) An occurrences budgeting method uses the number of good and total events during the time window. A timeslices budgeting method uses the number of good slices and total slices during the time window. A slice is an arbitrary time window (smaller than the overall SLO time window) that is either considered good or bad, calculated from the timeslice threshold and the ratio of good over total events that happened during the slice window. A budgeting method is required and must be either occurrences or timeslices.
	// An `occurrences` budgeting method uses the number of good and total events during the time window. A `timeslices` budgeting method uses the number of good slices and total slices during the time window. A slice is an arbitrary time window (smaller than the overall SLO time window) that is either considered good or bad, calculated from the timeslice threshold and the ratio of good over total events that happened during the slice window. A budgeting method is required and must be either occurrences or timeslices.
	BudgetingMethod *string `json:"budgetingMethod,omitempty" tf:"budgeting_method,omitempty"`

	// (String) A description for the SLO.
	// A description for the SLO.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Optional group by field to use to generate an SLO per distinct value.
	// Optional group by field to use to generate an SLO per distinct value.
	GroupBy *string `json:"groupBy,omitempty" tf:"group_by,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	HistogramCustomIndicator []HistogramCustomIndicatorInitParameters `json:"histogramCustomIndicator,omitempty" tf:"histogram_custom_indicator,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	KqlCustomIndicator []KqlCustomIndicatorInitParameters `json:"kqlCustomIndicator,omitempty" tf:"kql_custom_indicator,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	MetricCustomIndicator []MetricCustomIndicatorInitParameters `json:"metricCustomIndicator,omitempty" tf:"metric_custom_indicator,omitempty"`

	// (String) The name of the SLO.
	// The name of the SLO.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) The target objective is the value the SLO needs to meet during the time window. If a timeslices budgeting method is used, we also need to define the timesliceTarget which can be different than the overall SLO target. (see below for nested schema)
	// The target objective is the value the SLO needs to meet during the time window. If a timeslices budgeting method is used, we also need to define the timesliceTarget which can be different than the overall SLO target.
	Objective []ObjectiveInitParameters `json:"objective,omitempty" tf:"objective,omitempty"`

	// (Block List, Max: 1) The default settings should be sufficient for most users, but if needed, these properties can be overwritten. (see below for nested schema)
	// The default settings should be sufficient for most users, but if needed, these properties can be overwritten.
	Settings []SettingsInitParameters `json:"settings,omitempty" tf:"settings,omitempty"`

	// (Block List, Min: 1, Max: 1) Currently support calendarAligned and rolling time windows. Any duration greater than 1 day can be used: days, weeks, months, quarters, years. Rolling time window requires a duration, e.g. 1w for one week, and type: rolling. SLOs defined with such time window, will only consider the SLI data from the last duration period as a moving window. Calendar aligned time window requires a duration, limited to 1M for monthly or 1w for weekly, and type: calendarAligned. (see below for nested schema)
	// Currently support `calendarAligned` and `rolling` time windows. Any duration greater than 1 day can be used: days, weeks, months, quarters, years. Rolling time window requires a duration, e.g. `1w` for one week, and type: `rolling`. SLOs defined with such time window, will only consider the SLI data from the last duration period as a moving window. Calendar aligned time window requires a duration, limited to `1M` for monthly or `1w` for weekly, and type: `calendarAligned`.
	TimeWindow []TimeWindowInitParameters `json:"timeWindow,omitempty" tf:"time_window,omitempty"`
}

type SLOObservation struct {

	// (Block List, Max: 1) (see below for nested schema)
	ApmAvailabilityIndicator []ApmAvailabilityIndicatorObservation `json:"apmAvailabilityIndicator,omitempty" tf:"apm_availability_indicator,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	ApmLatencyIndicator []ApmLatencyIndicatorObservation `json:"apmLatencyIndicator,omitempty" tf:"apm_latency_indicator,omitempty"`

	// (String) An occurrences budgeting method uses the number of good and total events during the time window. A timeslices budgeting method uses the number of good slices and total slices during the time window. A slice is an arbitrary time window (smaller than the overall SLO time window) that is either considered good or bad, calculated from the timeslice threshold and the ratio of good over total events that happened during the slice window. A budgeting method is required and must be either occurrences or timeslices.
	// An `occurrences` budgeting method uses the number of good and total events during the time window. A `timeslices` budgeting method uses the number of good slices and total slices during the time window. A slice is an arbitrary time window (smaller than the overall SLO time window) that is either considered good or bad, calculated from the timeslice threshold and the ratio of good over total events that happened during the slice window. A budgeting method is required and must be either occurrences or timeslices.
	BudgetingMethod *string `json:"budgetingMethod,omitempty" tf:"budgeting_method,omitempty"`

	// (String) A description for the SLO.
	// A description for the SLO.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Optional group by field to use to generate an SLO per distinct value.
	// Optional group by field to use to generate an SLO per distinct value.
	GroupBy *string `json:"groupBy,omitempty" tf:"group_by,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	HistogramCustomIndicator []HistogramCustomIndicatorObservation `json:"histogramCustomIndicator,omitempty" tf:"histogram_custom_indicator,omitempty"`

	// side.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	KqlCustomIndicator []KqlCustomIndicatorObservation `json:"kqlCustomIndicator,omitempty" tf:"kql_custom_indicator,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	MetricCustomIndicator []MetricCustomIndicatorObservation `json:"metricCustomIndicator,omitempty" tf:"metric_custom_indicator,omitempty"`

	// (String) The name of the SLO.
	// The name of the SLO.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) The target objective is the value the SLO needs to meet during the time window. If a timeslices budgeting method is used, we also need to define the timesliceTarget which can be different than the overall SLO target. (see below for nested schema)
	// The target objective is the value the SLO needs to meet during the time window. If a timeslices budgeting method is used, we also need to define the timesliceTarget which can be different than the overall SLO target.
	Objective []ObjectiveObservation `json:"objective,omitempty" tf:"objective,omitempty"`

	// (Block List, Max: 1) The default settings should be sufficient for most users, but if needed, these properties can be overwritten. (see below for nested schema)
	// The default settings should be sufficient for most users, but if needed, these properties can be overwritten.
	Settings []SettingsObservation `json:"settings,omitempty" tf:"settings,omitempty"`

	// (String) An identifier for the space. If space_id is not provided, the default space is used.
	// An identifier for the space. If space_id is not provided, the default space is used.
	SpaceID *string `json:"spaceId,omitempty" tf:"space_id,omitempty"`

	// (Block List, Min: 1, Max: 1) Currently support calendarAligned and rolling time windows. Any duration greater than 1 day can be used: days, weeks, months, quarters, years. Rolling time window requires a duration, e.g. 1w for one week, and type: rolling. SLOs defined with such time window, will only consider the SLI data from the last duration period as a moving window. Calendar aligned time window requires a duration, limited to 1M for monthly or 1w for weekly, and type: calendarAligned. (see below for nested schema)
	// Currently support `calendarAligned` and `rolling` time windows. Any duration greater than 1 day can be used: days, weeks, months, quarters, years. Rolling time window requires a duration, e.g. `1w` for one week, and type: `rolling`. SLOs defined with such time window, will only consider the SLI data from the last duration period as a moving window. Calendar aligned time window requires a duration, limited to `1M` for monthly or `1w` for weekly, and type: `calendarAligned`.
	TimeWindow []TimeWindowObservation `json:"timeWindow,omitempty" tf:"time_window,omitempty"`
}

type SLOParameters struct {

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	ApmAvailabilityIndicator []ApmAvailabilityIndicatorParameters `json:"apmAvailabilityIndicator,omitempty" tf:"apm_availability_indicator,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	ApmLatencyIndicator []ApmLatencyIndicatorParameters `json:"apmLatencyIndicator,omitempty" tf:"apm_latency_indicator,omitempty"`

	// (String) An occurrences budgeting method uses the number of good and total events during the time window. A timeslices budgeting method uses the number of good slices and total slices during the time window. A slice is an arbitrary time window (smaller than the overall SLO time window) that is either considered good or bad, calculated from the timeslice threshold and the ratio of good over total events that happened during the slice window. A budgeting method is required and must be either occurrences or timeslices.
	// An `occurrences` budgeting method uses the number of good and total events during the time window. A `timeslices` budgeting method uses the number of good slices and total slices during the time window. A slice is an arbitrary time window (smaller than the overall SLO time window) that is either considered good or bad, calculated from the timeslice threshold and the ratio of good over total events that happened during the slice window. A budgeting method is required and must be either occurrences or timeslices.
	// +kubebuilder:validation:Optional
	BudgetingMethod *string `json:"budgetingMethod,omitempty" tf:"budgeting_method,omitempty"`

	// (String) A description for the SLO.
	// A description for the SLO.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Optional group by field to use to generate an SLO per distinct value.
	// Optional group by field to use to generate an SLO per distinct value.
	// +kubebuilder:validation:Optional
	GroupBy *string `json:"groupBy,omitempty" tf:"group_by,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	HistogramCustomIndicator []HistogramCustomIndicatorParameters `json:"histogramCustomIndicator,omitempty" tf:"histogram_custom_indicator,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	KqlCustomIndicator []KqlCustomIndicatorParameters `json:"kqlCustomIndicator,omitempty" tf:"kql_custom_indicator,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	MetricCustomIndicator []MetricCustomIndicatorParameters `json:"metricCustomIndicator,omitempty" tf:"metric_custom_indicator,omitempty"`

	// (String) The name of the SLO.
	// The name of the SLO.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) The target objective is the value the SLO needs to meet during the time window. If a timeslices budgeting method is used, we also need to define the timesliceTarget which can be different than the overall SLO target. (see below for nested schema)
	// The target objective is the value the SLO needs to meet during the time window. If a timeslices budgeting method is used, we also need to define the timesliceTarget which can be different than the overall SLO target.
	// +kubebuilder:validation:Optional
	Objective []ObjectiveParameters `json:"objective,omitempty" tf:"objective,omitempty"`

	// (Block List, Max: 1) The default settings should be sufficient for most users, but if needed, these properties can be overwritten. (see below for nested schema)
	// The default settings should be sufficient for most users, but if needed, these properties can be overwritten.
	// +kubebuilder:validation:Optional
	Settings []SettingsParameters `json:"settings,omitempty" tf:"settings,omitempty"`

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

	// (Block List, Min: 1, Max: 1) Currently support calendarAligned and rolling time windows. Any duration greater than 1 day can be used: days, weeks, months, quarters, years. Rolling time window requires a duration, e.g. 1w for one week, and type: rolling. SLOs defined with such time window, will only consider the SLI data from the last duration period as a moving window. Calendar aligned time window requires a duration, limited to 1M for monthly or 1w for weekly, and type: calendarAligned. (see below for nested schema)
	// Currently support `calendarAligned` and `rolling` time windows. Any duration greater than 1 day can be used: days, weeks, months, quarters, years. Rolling time window requires a duration, e.g. `1w` for one week, and type: `rolling`. SLOs defined with such time window, will only consider the SLI data from the last duration period as a moving window. Calendar aligned time window requires a duration, limited to `1M` for monthly or `1w` for weekly, and type: `calendarAligned`.
	// +kubebuilder:validation:Optional
	TimeWindow []TimeWindowParameters `json:"timeWindow,omitempty" tf:"time_window,omitempty"`
}

type SettingsInitParameters struct {

	// (String)
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (String)
	SyncDelay *string `json:"syncDelay,omitempty" tf:"sync_delay,omitempty"`
}

type SettingsObservation struct {

	// (String)
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (String)
	SyncDelay *string `json:"syncDelay,omitempty" tf:"sync_delay,omitempty"`
}

type SettingsParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	SyncDelay *string `json:"syncDelay,omitempty" tf:"sync_delay,omitempty"`
}

type TimeWindowInitParameters struct {

	// (String)
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TimeWindowObservation struct {

	// (String)
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TimeWindowParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type TotalInitParameters struct {

	// (String)
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation,omitempty"`

	// (String)
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Number)
	From *float64 `json:"from,omitempty" tf:"from,omitempty"`

	// (Number)
	To *float64 `json:"to,omitempty" tf:"to,omitempty"`
}

type TotalMetricsInitParameters struct {

	// (String)
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation,omitempty"`

	// (String)
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) The name of the SLO.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TotalMetricsObservation struct {

	// (String)
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation,omitempty"`

	// (String)
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) The name of the SLO.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TotalMetricsParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Aggregation *string `json:"aggregation" tf:"aggregation,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Field *string `json:"field" tf:"field,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) The name of the SLO.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type TotalObservation struct {

	// (String)
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation,omitempty"`

	// (String)
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Number)
	From *float64 `json:"from,omitempty" tf:"from,omitempty"`

	// (Number)
	To *float64 `json:"to,omitempty" tf:"to,omitempty"`
}

type TotalParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Aggregation *string `json:"aggregation" tf:"aggregation,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Field *string `json:"field" tf:"field,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	From *float64 `json:"from,omitempty" tf:"from,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	To *float64 `json:"to,omitempty" tf:"to,omitempty"`
}

// SLOSpec defines the desired state of SLO
type SLOSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SLOParameters `json:"forProvider"`
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
	InitProvider SLOInitParameters `json:"initProvider,omitempty"`
}

// SLOStatus defines the observed state of SLO.
type SLOStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SLOObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SLO is the Schema for the SLOs API. Creates or updates a Kibana SLO.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticstack}
type SLO struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.budgetingMethod) || has(self.initProvider.budgetingMethod)",message="budgetingMethod is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || has(self.initProvider.description)",message="description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objective) || has(self.initProvider.objective)",message="objective is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timeWindow) || has(self.initProvider.timeWindow)",message="timeWindow is a required parameter"
	Spec   SLOSpec   `json:"spec"`
	Status SLOStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SLOList contains a list of SLOs
type SLOList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SLO `json:"items"`
}

// Repository type metadata.
var (
	SLO_Kind             = "SLO"
	SLO_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SLO_Kind}.String()
	SLO_KindAPIVersion   = SLO_Kind + "." + CRDGroupVersion.String()
	SLO_GroupVersionKind = CRDGroupVersion.WithKind(SLO_Kind)
)

func init() {
	SchemeBuilder.Register(&SLO{}, &SLOList{})
}
