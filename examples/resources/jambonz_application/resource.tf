data "jambonz_account" "my_account" {
  account_sid = "b42f0f47-3972-4361-a2a4-e69cf0e1e8c3"
}

resource "jambonz_application" "my_application" {
  name             = "My Test Application"
  account_sid      = data.jambonz_account.my_account
  record_all_calls = false

  call_hook = {
    url    = "https://example.com/calls"
    method = "POST"
    # Optional auth
    username = "user"
    password = "pass"
  }

  call_status_hook = {
    url    = "https://example.com/status"
    method = "POST"
  }

  # This block is required but the URL can be empty if no
  # hook is used.
  messaging_hook = {
    url    = ""
    method = "POST"
  }
}
