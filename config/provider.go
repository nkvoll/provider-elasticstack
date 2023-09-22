/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "elasticstack"
	modulePath     = "github.com/elastic/provider-elasticstack"

	elasticsearchShortGroup = "elasticsearch"
	fleetShortGroup         = "fleet"
	kibanaShortGroup        = "kibana"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("elasticstack.k8s.elastic.co"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		func(p *ujconfig.Provider) {
			p.AddResourceConfigurator("elasticstack_kibana_action_connector", func(r *ujconfig.Resource) {
				r.ShortGroup = kibanaShortGroup

				// Mark secrets as sensitive -- this should probably be changed in the elasticstack provider?
				r.TerraformResource.Schema["secrets"].Sensitive = true

				r.References["space_id"] = ujconfig.Reference{
					Type:      "github.com/elastic/provider-elasticstack/apis/kibana/v1alpha1.Space",
					Extractor: "github.com/upbound/upjet/pkg/resource.ExtractParamPath(\"space_id\",true)",
				}
			})
			p.AddResourceConfigurator("elasticstack_kibana_alerting_rule", func(r *ujconfig.Resource) {
				r.ShortGroup = kibanaShortGroup

				r.References["actions.id"] = ujconfig.Reference{
					Type:      "github.com/elastic/provider-elasticstack/apis/kibana/v1alpha1.ActionConnector",
					Extractor: "github.com/upbound/upjet/pkg/resource.ExtractParamPath(\"connector_id\",true)",
				}
				r.References["space_id"] = ujconfig.Reference{
					Type:      "github.com/elastic/provider-elasticstack/apis/kibana/v1alpha1.Space",
					Extractor: "github.com/upbound/upjet/pkg/resource.ExtractParamPath(\"space_id\",true)",
				}
			})
			p.AddResourceConfigurator("elasticstack_kibana_slo", func(r *ujconfig.Resource) {
				r.ShortGroup = kibanaShortGroup

				r.References["space_id"] = ujconfig.Reference{
					Type:      "github.com/elastic/provider-elasticstack/apis/kibana/v1alpha1.Space",
					Extractor: "github.com/upbound/upjet/pkg/resource.ExtractParamPath(\"space_id\",true)",
				}
			})
			p.AddResourceConfigurator("elasticstack_kibana_space", func(r *ujconfig.Resource) {
				r.ShortGroup = kibanaShortGroup

				r.References["space_id"] = ujconfig.Reference{
					Type:      "github.com/elastic/provider-elasticstack/apis/kibana/v1alpha1.Space",
					Extractor: "github.com/upbound/upjet/pkg/resource.ExtractParamPath(\"space_id\",true)",
				}
			})
		},
		func(p *ujconfig.Provider) {
			p.AddResourceConfigurator("elasticstack_elasticsearch_cluster_settings", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_component_template", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_data_stream", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_enrich_policy", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_index", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_index_lifecycle", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_index_template", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_ingest_pipeline", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_logstash_pipeline", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_script", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_security_api_key", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_security_role", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_security_role_mapping", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_security_system_user", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_security_user", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_snapshot_lifecycle", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_snapshot_repository", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_transform", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})
			p.AddResourceConfigurator("elasticstack_elasticsearch_watch", func(r *ujconfig.Resource) {
				r.ShortGroup = elasticsearchShortGroup
			})

		},
		func(p *ujconfig.Provider) {
			p.AddResourceConfigurator("elasticstack_fleet_agent_policy", func(r *ujconfig.Resource) {
				r.ShortGroup = fleetShortGroup
			})
			p.AddResourceConfigurator("elasticstack_fleet_output", func(r *ujconfig.Resource) {
				r.ShortGroup = fleetShortGroup
			})
			p.AddResourceConfigurator("elasticstack_fleet_server_host", func(r *ujconfig.Resource) {
				r.ShortGroup = fleetShortGroup
			})
		},
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
