package provider

import (
	api "terraform-provider-jambonz/internal/api/generated"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func OptStringToStringType(s api.OptString) types.String {
	if value, ok := s.Get(); ok {
		return types.StringValue(value)
	} else {
		return types.StringNull()
	}
}

func OptNilStringToStringType(s api.OptNilString) types.String {
	if value, ok := s.Get(); ok {
		return types.StringValue(value)
	} else {
		return types.StringNull()
	}
}

func OptUuidToStringType(s api.OptUUID) types.String {
	if value, ok := s.Get(); ok {
		return types.StringValue(value.String())
	} else {
		return types.StringNull()
	}
}

func OptNilUuidToStringType(s api.OptNilUUID) types.String {
	if value, ok := s.Get(); ok {
		return types.StringValue(value.String())
	} else {
		return types.StringNull()
	}
}
