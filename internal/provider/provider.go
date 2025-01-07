// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	api_client "terraform-provider-jambonz/internal/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure JambonzProvider satisfies various provider interfaces.
var _ provider.Provider = &JambonzProvider{}

// JambonzProvider defines the provider implementation.
type JambonzProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// jambonzProviderModel describes the provider data model.
type JambonzProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	ApiKey   types.String `tfsdk:"api_key"`
}

func (p *JambonzProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "jambonz"
	resp.Version = p.version
}

func (p *JambonzProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				Description: "Jambonz API endpoint. May also be provided via the JAMBONZ_ENDPOINT environment variable.",
				Required:    true,
			},
			"api_key": schema.StringAttribute{
				Description: "Jambonz API key. May also be provided via the JAMBONZ_API_KEY environment variable.",
				Required:    true,
			},
		},
	}
}

func (p *JambonzProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config JambonzProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value
	if config.Endpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown Jambonz API endpoint",
			"The provider cannot create the Jambonz API client as there is an unknown configuration value for the Jambonz API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the JAMBONZ_HOST environment variable.",
		)
	}

	if config.ApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown Jambonz API endpoint",
			"The provider cannot create the Jambonz API client as there is an unknown configuration value for the Jambonz API key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the JAMBONZ_API_KEY environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := os.Getenv("JAMBONZ_ENDPOINT")
	api_key := os.Getenv("JAMBONZ_API_KEY")

	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}
	if !config.ApiKey.IsNull() {
		api_key = config.ApiKey.ValueString()
	}

	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Missing Jambonz API Host",
			"The provider cannot create the Jambonz API client as there is a missing or empty value for the Jambonz API host. "+
				"Set the host value in the configuration or use the JAMBONZ_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if api_key == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing Jambonz API key",
			"The provider cannot create the Jambonz API client as there is a missing or empty value for the Jambonz API key. "+
				"Set the host value in the configuration or use the JAMBONZ_API_KEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	client, err := api_client.NewClient(ctx, endpoint, api_key)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create Jambonz API Client",
			"An unexpected error occurred when creating the Jambonz API client. "+
				"Jambonz Client Error: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *JambonzProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewPhoneNumberResource,
	}
}

func (p *JambonzProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAccountDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &JambonzProvider{
			version: version,
		}
	}
}
