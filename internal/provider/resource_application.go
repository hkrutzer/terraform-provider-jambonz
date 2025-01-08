package provider

import (
	"context"
	"fmt"

	api "terraform-provider-jambonz/internal/api/generated"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &applicationResource{}
	_ resource.ResourceWithConfigure   = &applicationResource{}
	_ resource.ResourceWithImportState = &applicationResource{}
)

func NewApplicationResource() resource.Resource {
	return &applicationResource{}
}

type applicationResource struct {
	client *api.Client
}

func (r *applicationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application"
}

func (r *applicationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *applicationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	webhookSchema := map[string]schema.Attribute{
		"webhook_sid": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			Validators: []validator.String{
				UUIDValidator,
			},
		},
		"url": schema.StringAttribute{
			Required: true,
		},
		"method": schema.StringAttribute{
			Required: true,
			// TODO Default: stringdefault.StaticString("POST"),
		},
		"username": schema.StringAttribute{
			Optional: true,
		},
		"password": schema.StringAttribute{
			Optional:  true,
			Sensitive: true,
		},
	}

	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"application_sid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"account_sid": schema.StringAttribute{
				Required: true,
			},
			"call_hook": schema.SingleNestedAttribute{
				Required:   true,
				Attributes: webhookSchema,
			},
			"call_status_hook": schema.SingleNestedAttribute{
				Required:   true,
				Attributes: webhookSchema,
			},
			"messaging_hook": schema.SingleNestedAttribute{
				Required:   true,
				Attributes: webhookSchema,
			},
			"record_all_calls": schema.BoolAttribute{
				Required: true,
			},
		},
	}
}

func (r *applicationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplicationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	recordAllCalls := api.CreateApplicationReqRecordAllCalls0
	if plan.RecordAllCalls.ValueBool() {
		recordAllCalls = api.CreateApplicationReqRecordAllCalls1
	}

	applicationReq := api.CreateApplicationReq{
		Name:           plan.Name.ValueString(),
		AccountSid:     uuid.MustParse(plan.AccountSID.ValueString()),
		CallHook:       toWebhook(plan.CallHook),
		CallStatusHook: toWebhook(plan.CallStatusHook),
		MessagingHook:  toWebhook(plan.MessagingHook),
		RecordAllCalls: recordAllCalls,
	}

	// Call the API to create the application
	resCreateApplication, err := r.client.CreateApplication(ctx, &applicationReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating application",
			"Could not create application, unexpected error: "+err.Error(),
		)
		return
	}

	added, ok := resCreateApplication.(*api.SuccessfulAdd)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected response type after creating account",
			fmt.Sprintf("Expected *api.SuccessfulAdd, got %+v from %+v", resCreateApplication, applicationReq),
		)
		return
	}

	resApplication, err := r.client.GetApplication(ctx, api.GetApplicationParams{
		ApplicationSid: added.Sid,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading application",
			"Could not read application, unexpected error: "+err.Error(),
		)
		return
	}

	application, ok := resApplication.(*api.Application)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected response type",
			fmt.Sprintf("Expected *api.Application, got %+v", resApplication),
		)
		return
	}

	// Set the state with the created application
	state := toApplicationResourceModel(application)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *applicationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplicationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resApplication, err := r.client.GetApplication(ctx, api.GetApplicationParams{
		ApplicationSid: state.ApplicationSID.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading application",
			"Could not read application, unexpected error: "+err.Error(),
		)
		return
	}

	if _, ok := resApplication.(*api.GetApplicationNotFound); ok {
		// if not found consider it deleted
		resp.State.RemoveResource(ctx)
		return
	}

	application, ok := resApplication.(*api.Application)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected response type",
			fmt.Sprintf("Expected *api.Application, got %+v", resApplication),
		)
		return
	}

	// Update the state with the fetched application
	state = toApplicationResourceModel(application)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *applicationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan ApplicationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	identifier := api.UpdateApplicationParams{
		ApplicationSid: state.ApplicationSID.ValueString(),
	}

	updateBody := api.UpdateApplicationReq{
		Name:           plan.Name.ValueString(),
		AccountSid:     uuid.MustParse(plan.AccountSID.ValueString()),
		CallHook:       toWebhook(plan.CallHook),
		CallStatusHook: toWebhook(plan.CallStatusHook),
		MessagingHook:  toWebhook(plan.MessagingHook),
	}

	updateBody.SetRecordAllCalls(api.UpdateApplicationReqRecordAllCalls0)
	if plan.RecordAllCalls.ValueBool() {
		updateBody.SetRecordAllCalls(api.UpdateApplicationReqRecordAllCalls1)
	}

	_, err := r.client.UpdateApplication(ctx, &updateBody, identifier)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating application",
			"Could not update application, unexpected error: "+err.Error(),
		)
		return
	}

	// Read the resource back because the API doesn't return anything
	getRes, err := r.client.GetApplication(ctx, api.GetApplicationParams(identifier))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating application",
			"Could not fetch application, unexpected error: "+err.Error(),
		)
		return
	}
	application, ok := getRes.(*api.Application)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected response type",
			fmt.Sprintf("Expected *api.Application, got %+v", getRes),
		)
		return
	}

	// Update the state with the fetched application
	state = toApplicationResourceModel(application)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *applicationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplicationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.DeleteApplication(ctx, api.DeleteApplicationParams{
		ApplicationSid: state.ApplicationSID.ValueString(),
	})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting application",
			"Could not delete application, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *applicationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("application_sid"), req, resp)
}

func getNilStringFromValue(value types.String) api.NilString {
	if value.IsNull() {
		return api.NilString{Null: true}
	}
	return api.NilString{
		Value: value.ValueString(),
		Null:  false,
	}
}

func toApplicationResourceModel(application *api.Application) ApplicationResourceModel {
	call_hook := toWebhookModel(&application.CallHook)
	call_status_hook := toWebhookModel(&application.CallStatusHook)
	messaging_hook := toWebhookModel(&application.MessagingHook)

	model := ApplicationResourceModel{
		ApplicationSID: types.StringValue(application.ApplicationSid.String()),
		Name:           types.StringValue(application.Name),
		AccountSID:     types.StringValue(application.AccountSid.String()),
		CallHook:       &call_hook,
		CallStatusHook: &call_status_hook,
		MessagingHook:  &messaging_hook,
		RecordAllCalls: types.BoolValue(application.RecordAllCalls == 1),
	}

	return model
}

func toWebhook(wh *WebhookModel) api.Webhook {
	// TODO Is there a creator thing?
	webhookSid := api.OptUUID{
		Set: false,
	}

	if !wh.WebhookSID.IsUnknown() && !wh.WebhookSID.IsNull() {
		webhookSid = api.NewOptUUID(uuid.MustParse(wh.WebhookSID.String()))
	}

	return api.Webhook{
		WebhookSid: webhookSid,
		URL:        wh.URL.ValueString(),
		Method:     api.WebhookMethod(wh.Method.ValueString()),
		Username:   getNilStringFromValue(wh.Username),
		Password:   getNilStringFromValue(wh.Password),
	}
}

func toWebhookModel(webhook *api.Webhook) WebhookModel {
	model := WebhookModel{
		URL:    types.StringValue(webhook.URL),
		Method: types.StringValue(string(webhook.Method)),
	}

	if webhook.WebhookSid.IsSet() {
		model.WebhookSID = types.StringValue(webhook.WebhookSid.Value.String())
	}

	if !webhook.Username.IsNull() {
		model.Username = types.StringValue(webhook.Username.Value)
	}

	if !webhook.Password.IsNull() {
		model.Password = types.StringValue(webhook.Password.Value)
	}

	return model
}

type ApplicationResourceModel struct {
	ApplicationSID types.String  `tfsdk:"application_sid"`
	Name           types.String  `tfsdk:"name"`
	AccountSID     types.String  `tfsdk:"account_sid"`
	CallHook       *WebhookModel `tfsdk:"call_hook"`
	CallStatusHook *WebhookModel `tfsdk:"call_status_hook"`
	MessagingHook  *WebhookModel `tfsdk:"messaging_hook"`
	RecordAllCalls types.Bool    `tfsdk:"record_all_calls"`
}
