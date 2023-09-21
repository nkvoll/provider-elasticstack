/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/elastic/provider-elasticstack/config/kibana_action_connector"
	"github.com/elastic/provider-elasticstack/config/kibana_alerting_rule"
)

const (
	resourcePrefix = "elasticstack"
	modulePath     = "github.com/elastic/provider-elasticstack"
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
		kibana_action_connector.Configure,
		kibana_alerting_rule.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
