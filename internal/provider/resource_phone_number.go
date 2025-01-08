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
	_ resource.Resource                = &phoneNumberResource{}
	_ resource.ResourceWithConfigure   = &phoneNumberResource{}
	_ resource.ResourceWithImportState = &phoneNumberResource{}
)

func NewPhoneNumberResource() resource.Resource {
	return &phoneNumberResource{}
}

type phoneNumberResource struct {
	client *api.Client
}

func (r *phoneNumberResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_phone_number"
}

func (r *phoneNumberResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *phoneNumberResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"phone_number_sid": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					UUIDValidator,
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"phone_number": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"voip_carrier_sid": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					UUIDValidator,
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"account_sid": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					UUIDValidator,
				},
			},
			"application_sid": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					UUIDValidator,
				},
			},
		},
	}
}

func (r *phoneNumberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state PhoneNumberResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	provisionRequest := api.ProvisionPhoneNumberReq{
		AccountSid:     uuid.MustParse(plan.AccountSID.ValueString()),
		Number:         plan.PhoneNumber.ValueString(),
		VoipCarrierSid: uuid.MustParse(plan.VoipCarrierSID.ValueString()),
	}

	if !plan.ApplicationSID.IsNull() {
		provisionRequest.ApplicationSid.SetTo(uuid.MustParse(plan.ApplicationSID.ValueString()))
	}

	// Call the API to provision the phone number
	res_provision, err := r.client.ProvisionPhoneNumber(ctx, &provisionRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error provisioning phone number",
			"Could not provision phone number, unexpected error: "+err.Error(),
		)
		return
	}

	added, ok := res_provision.(*api.SuccessfulAdd)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected response type after provisioning phone number",
			fmt.Sprintf("Expected *api.SuccessfulAdd, got %+v", res_provision),
		)
		return
	}

	// API does not return new data so we must fetch it
	res, err := r.client.GetPhoneNumber(ctx, api.GetPhoneNumberParams{
		PhoneNumberSid: added.Sid,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading phone number",
			fmt.Sprintf("Could not read phone number with ID %v due to %+v, %+v "+plan.PhoneNumberSID.ValueString(), err.Error(), res),
		)
		return
	}

	phoneNumber, ok := res.(*api.PhoneNumber)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected response type",
			fmt.Sprintf("Expected *api.PhoneNumber, got %+v", res),
		)
		return
	}

	state.PhoneNumberSID = types.StringValue(added.Sid)
	state.PhoneNumber = types.StringValue(phoneNumber.Number)
	state.VoipCarrierSID = types.StringValue(phoneNumber.VoipCarrierSid.String())
	state.AccountSID = types.StringValue(phoneNumber.AccountSid.Value.String())

	if value, ok := phoneNumber.ApplicationSid.Get(); ok {
		state.ApplicationSID = types.StringValue(value.String())
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *phoneNumberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PhoneNumberResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	phoneNumberSid, err := uuid.Parse(state.PhoneNumberSID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid UUID",
			fmt.Sprintf("Invalid phone number SID %+v", state),
		)
		return
	}

	res, err := r.client.GetPhoneNumber(ctx, api.GetPhoneNumberParams{
		PhoneNumberSid: phoneNumberSid.String(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading phone number",
			fmt.Sprintf("Could not retrieve phone number info: %+v %+v", err.Error(), res),
		)
		return
	}

	if _, ok := res.(*api.GetPhoneNumberNotFound); ok {
		// if not found consider it deleted
		resp.State.RemoveResource(ctx)
		return
	}

	phoneNumber, ok := res.(*api.PhoneNumber)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected response type",
			fmt.Sprintf("Expected *api.PhoneNumber, got %+v", res),
		)
		return
	}

	state.PhoneNumber = types.StringValue(phoneNumber.Number)
	state.VoipCarrierSID = types.StringValue(phoneNumber.VoipCarrierSid.String())
	state.AccountSID = types.StringValue(phoneNumber.AccountSid.Value.String())
	state.ApplicationSID = OptNilUuidToStringType(phoneNumber.ApplicationSid)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *phoneNumberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan PhoneNumberResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updatePhoneNumber := api.UpdatePhoneNumberParams{
		PhoneNumberSid: state.PhoneNumberSID.ValueString(),
	}

	x := api.UpdatePhoneNumberReq{
		AccountSid: uuid.MustParse(plan.AccountSID.ValueString()),
	}

	if !plan.ApplicationSID.IsNull() {
		x.ApplicationSid.SetTo(uuid.MustParse(plan.ApplicationSID.ValueString()))
	} else {
		x.ApplicationSid.SetToNull()
	}

	_, err := r.client.UpdatePhoneNumber(ctx, &x, updatePhoneNumber)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating phone number",
			"Could not update phone number, unexpected error: "+err.Error(),
		)
		return
	}

	// Read the resource back because the API doesn't return anything
	res, err := r.client.GetPhoneNumber(ctx, api.GetPhoneNumberParams{
		PhoneNumberSid: state.PhoneNumberSID.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating phone number",
			fmt.Sprintf("Could not retrieve phone number info %+v %+v", err.Error(), res),
		)
		return
	}

	phoneNumber, ok := res.(*api.PhoneNumber)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected response type",
			fmt.Sprintf("Expected *api.PhoneNumber, got %+v", res),
		)
		return
	}

	state.PhoneNumber = types.StringValue(phoneNumber.Number)
	state.VoipCarrierSID = types.StringValue(phoneNumber.VoipCarrierSid.String())
	state.AccountSID = types.StringValue(phoneNumber.AccountSid.Value.String())
	state.ApplicationSID = OptNilUuidToStringType(phoneNumber.ApplicationSid)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *phoneNumberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PhoneNumberResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.DeletePhoneNumber(ctx, api.DeletePhoneNumberParams{
		PhoneNumberSid: state.PhoneNumberSID.ValueString(),
	})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Phone Number",
			"Could not delete phone number, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *phoneNumberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("phone_number_sid"), req, resp)
}

type PhoneNumberResourceModel struct {
	PhoneNumberSID types.String `tfsdk:"phone_number_sid"`
	PhoneNumber    types.String `tfsdk:"phone_number"`
	VoipCarrierSID types.String `tfsdk:"voip_carrier_sid"`
	AccountSID     types.String `tfsdk:"account_sid"`
	ApplicationSID types.String `tfsdk:"application_sid"`
}
