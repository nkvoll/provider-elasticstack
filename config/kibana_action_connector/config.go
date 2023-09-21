package kibana_action_connector

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("elasticstack_kibana_action_connector", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "elasticstack"
		r.ShortGroup = "kibana_action_connector"
	})
}
