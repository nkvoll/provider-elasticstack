package kibana_alerting_rule

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("elasticstack_kibana_alerting_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "elasticstack"
		r.ShortGroup = "kibana_alerting_rule"

		// This resource need the repository in which branch would be created
        // as an input. And by defining it as a reference to Repository
        // object, we can build cross resource referencing. See
        // repositoryRef in the example in the Testing section below.
        r.References["action_connector"] = config.Reference{
            Type: "github.com/elastic/provider-elasticstack/apis/repository/v1alpha1.KibanaActionConnector",
        }
	})
}
