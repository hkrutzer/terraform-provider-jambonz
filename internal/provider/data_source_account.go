package provider

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	api "terraform-provider-jambonz/internal/api/generated"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &accountDataSource{}
	_ datasource.DataSourceWithConfigure = &accountDataSource{}
)

func NewAccountDataSource() datasource.DataSource {
	return &accountDataSource{}
}

type accountDataSource struct {
	client *api.Client
}

func (d *accountDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account"
}

func (d *accountDataSource) Schema(_ context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_sid": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					UUIDValidator,
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"sip_realm": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"registration_hook": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"webhook_sid": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							UUIDValidator,
						},
					},
					"url": schema.StringAttribute{
						Computed: true,
					},
					"method": schema.StringAttribute{
						Computed: true,
					},
					"username": schema.StringAttribute{
						Computed: true,
					},
					"password": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"device_calling_application_sid": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					UUIDValidator,
				},
			},
			"service_provider_sid": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					UUIDValidator,
				},
			},
		},
	}
}

func (d *accountDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *accountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data accountDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	account_sid := uuid.MustParse(data.AccountSID.ValueString())

	res, err := d.client.GetAccount(ctx, api.GetAccountParams{
		AccountSid: account_sid,
	})

	if err != nil {
		resp.Diagnostics.AddError("Failed to fetch Account", err.Error())
		return
	}

	switch p := res.(type) {
	case *api.Account:
		data.AccountSID = types.StringValue(p.AccountSid.String())
		data.Name = types.StringValue(p.Name)
		data.SipRealm = OptNilStringToStringType(p.SipRealm)
		data.DeviceCallingApplicationSID = OptNilUuidToStringType(p.DeviceCallingApplicationSid)
		data.ServiceProviderSID = types.StringValue(p.ServiceProviderSid.String())
		if registration_hook, ok := p.RegistrationHook.Get(); ok {
			data.RegistrationHook.WebhookSID = OptUuidToStringType(registration_hook.WebhookSid)
			data.RegistrationHook.URL = types.StringValue(registration_hook.URL)
			data.RegistrationHook.Method = types.StringValue(string(registration_hook.Method))
			data.RegistrationHook.Username = OptStringToStringType(registration_hook.Username)
			data.RegistrationHook.Password = OptStringToStringType(registration_hook.Password)
		}
	default:
		resp.Diagnostics.AddError("Failed to fetch Account", fmt.Sprintf("%+v", p))
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

type accountDataSourceModel struct {
	AccountSID                  types.String  `tfsdk:"account_sid"`
	Name                        types.String  `tfsdk:"name"`
	SipRealm                    types.String  `tfsdk:"sip_realm"`
	RegistrationHook            *webhookModel `tfsdk:"registration_hook"`
	DeviceCallingApplicationSID types.String  `tfsdk:"device_calling_application_sid"`
	ServiceProviderSID          types.String  `tfsdk:"service_provider_sid"`
}

type webhookModel struct {
	WebhookSID types.String `tfsdk:"webhook_sid"`
	URL        types.String `tfsdk:"url"`
	Method     types.String `tfsdk:"method"`
	Username   types.String `tfsdk:"username"`
	Password   types.String `tfsdk:"password"`
}
