package provider

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
)

var UUIDValidator = stringvalidator.RegexMatches(
	regexp.MustCompile(`(?i)^[a-f0-9]{8}-[a-f0-9]{4}-[1-5][a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`),
	"must be a valid UUID",
)
