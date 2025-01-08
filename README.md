# Jambonz Terraform Provider
[![Terraform Registry](https://img.shields.io/github/v/release/hkrutzer/terraform-provider-jambonz?color=5e4fe3&label=Terraform%20Registry&logo=terraform&sort=semver)](https://registry.terraform.io/providers/hkrutzer/jambonz/latest)

_Unofficial_ Terraform provider for [jambonz](https://jambonz.org/)

- [Documentation](https://registry.terraform.io/providers/hkrutzer/jambonz/latest/docs)

## Using the provider

```hcl
# Specify the provider and version
terraform {
  required_providers {
    jambonz = {
      source  = "hkrutzer/jambonz"
      version = "~> 0.1.0"
    }
  }
}

# Configure the provider
provider "jambonz" {
  endpoint = "https://jambonz.cloud/api/v1"
  api_key  = "ff6dad71-6015-4a1a-af22-1dc97f71d3b1"
}

# Get an account
data "jambonz_account" "my_account" {
  account_sid = "187e88ad-c55e-48c2-b87d-cb41f4d68967"
}

# Create a phone number
resource "jambonz_phone_number" "my_number" {
  phone_number = "+1234567890"
  voip_carrier_sid = "da27ee1a-a0dc-4126-bd46-6dbb9afad942"
  account_sid = data.jambonz_account.my_account.account_sid
}
```
