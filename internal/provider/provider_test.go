package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//nolint:unparam
func protoV5ProviderFactories() map[string]func() (tfprotov5.ProviderServer, error) {
	return map[string]func() (tfprotov5.ProviderServer, error){
		"randombyte": providerserver.NewProtocol5WithError(New()),
	}
}

func providerVersion101() map[string]resource.ExternalProvider {
	return map[string]resource.ExternalProvider{
		"randombyte": {
			VersionConstraint: "1.0.1",
			Source:            "socolin/randombyte",
		},
	}
}
