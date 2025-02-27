// Code generated by ogen, DO NOT EDIT.

package api

// OperationName is the ogen operation name
type OperationName = string

const (
	AddLimitForAccountOperation                            OperationName = "AddLimitForAccount"
	AddLimitForServiceProviderOperation                    OperationName = "AddLimitForServiceProvider"
	AddSpeechCredentialForSeerviceProviderOperation        OperationName = "AddSpeechCredentialForSeerviceProvider"
	ChangePasswordOperation                                OperationName = "ChangePassword"
	CheckAvailabilityOperation                             OperationName = "CheckAvailability"
	CreateAccountOperation                                 OperationName = "CreateAccount"
	CreateApikeyOperation                                  OperationName = "CreateApikey"
	CreateApplicationOperation                             OperationName = "CreateApplication"
	CreateCallOperation                                    OperationName = "CreateCall"
	CreateCarrierForServiceProviderOperation               OperationName = "CreateCarrierForServiceProvider"
	CreateGoogleCustomVoiceOperation                       OperationName = "CreateGoogleCustomVoice"
	CreateLcrOperation                                     OperationName = "CreateLcr"
	CreateLcrCarrierSetEntryOperation                      OperationName = "CreateLcrCarrierSetEntry"
	CreateLcrForServiceProviderOperation                   OperationName = "CreateLcrForServiceProvider"
	CreateLcrRoutesOperation                               OperationName = "CreateLcrRoutes"
	CreateLeastCostRoutingRoutesAndCarrierEntriesOperation OperationName = "CreateLeastCostRoutingRoutesAndCarrierEntries"
	CreateMessageOperation                                 OperationName = "CreateMessage"
	CreateMsTeamsTenantOperation                           OperationName = "CreateMsTeamsTenant"
	CreateSbcOperation                                     OperationName = "CreateSbc"
	CreateServiceProviderOperation                         OperationName = "CreateServiceProvider"
	CreateSipGatewayOperation                              OperationName = "CreateSipGateway"
	CreateSipRealmOperation                                OperationName = "CreateSipRealm"
	CreateSmppGatewayOperation                             OperationName = "CreateSmppGateway"
	CreateSpeechCredentialOperation                        OperationName = "CreateSpeechCredential"
	CreateUserOperation                                    OperationName = "CreateUser"
	CreateVoipCarrierOperation                             OperationName = "CreateVoipCarrier"
	CreateVoipCarrierFromTemplateOperation                 OperationName = "CreateVoipCarrierFromTemplate"
	CreateVoipCarrierFromTemplateBySPOperation             OperationName = "CreateVoipCarrierFromTemplateBySP"
	DeleteAccountOperation                                 OperationName = "DeleteAccount"
	DeleteApiKeyOperation                                  OperationName = "DeleteApiKey"
	DeleteApplicationOperation                             OperationName = "DeleteApplication"
	DeleteCallOperation                                    OperationName = "DeleteCall"
	DeleteGoogleCustomVoiceOperation                       OperationName = "DeleteGoogleCustomVoice"
	DeleteLeastCostRoutingOperation                        OperationName = "DeleteLeastCostRouting"
	DeleteLeastCostRoutingCarrierSetEntryOperation         OperationName = "DeleteLeastCostRoutingCarrierSetEntry"
	DeleteLeastCostRoutingRouteOperation                   OperationName = "DeleteLeastCostRoutingRoute"
	DeletePhoneNumberOperation                             OperationName = "DeletePhoneNumber"
	DeleteSbcAddressOperation                              OperationName = "DeleteSbcAddress"
	DeleteServiceProviderOperation                         OperationName = "DeleteServiceProvider"
	DeleteSipGatewayOperation                              OperationName = "DeleteSipGateway"
	DeleteSmppGatewayOperation                             OperationName = "DeleteSmppGateway"
	DeleteSpeechCredentialOperation                        OperationName = "DeleteSpeechCredential"
	DeleteSpeechCredentialByAccountOperation               OperationName = "DeleteSpeechCredentialByAccount"
	DeleteTenantOperation                                  OperationName = "DeleteTenant"
	DeleteUserOperation                                    OperationName = "DeleteUser"
	DeleteVoipCarrierOperation                             OperationName = "DeleteVoipCarrier"
	ForgotPasswordOperation                                OperationName = "ForgotPassword"
	GenerateInviteCodeOperation                            OperationName = "GenerateInviteCode"
	GetAccountOperation                                    OperationName = "GetAccount"
	GetAccountApiKeysOperation                             OperationName = "GetAccountApiKeys"
	GetAccountLimitsOperation                              OperationName = "GetAccountLimits"
	GetApplicationOperation                                OperationName = "GetApplication"
	GetCallOperation                                       OperationName = "GetCall"
	GetGoogleCustomVoiceOperation                          OperationName = "GetGoogleCustomVoice"
	GetLeastCostRoutingOperation                           OperationName = "GetLeastCostRouting"
	GetLeastCostRoutingCarrierSetEntryOperation            OperationName = "GetLeastCostRoutingCarrierSetEntry"
	GetLeastCostRoutingRouteOperation                      OperationName = "GetLeastCostRoutingRoute"
	GetMyDetailsOperation                                  OperationName = "GetMyDetails"
	GetPhoneNumberOperation                                OperationName = "GetPhoneNumber"
	GetRecentCallTraceOperation                            OperationName = "GetRecentCallTrace"
	GetRecentCallTraceByAccountOperation                   OperationName = "GetRecentCallTraceByAccount"
	GetRecentCallTraceByCallIdOperation                    OperationName = "GetRecentCallTraceByCallId"
	GetRecentCallTraceBySPOperation                        OperationName = "GetRecentCallTraceBySP"
	GetRegisteredClientOperation                           OperationName = "GetRegisteredClient"
	GetServiceProviderOperation                            OperationName = "GetServiceProvider"
	GetServiceProviderAccountsOperation                    OperationName = "GetServiceProviderAccounts"
	GetServiceProviderCarriersOperation                    OperationName = "GetServiceProviderCarriers"
	GetServiceProviderLcrsOperation                        OperationName = "GetServiceProviderLcrs"
	GetServiceProviderLimitsOperation                      OperationName = "GetServiceProviderLimits"
	GetSipGatewayOperation                                 OperationName = "GetSipGateway"
	GetSmppGatewayOperation                                OperationName = "GetSmppGateway"
	GetSpeechCredentialOperation                           OperationName = "GetSpeechCredential"
	GetSpeechCredentialByAccountOperation                  OperationName = "GetSpeechCredentialByAccount"
	GetStripeCustomerIdOperation                           OperationName = "GetStripeCustomerId"
	GetSubscriptionOperation                               OperationName = "GetSubscription"
	GetTenantOperation                                     OperationName = "GetTenant"
	GetTestDataOperation                                   OperationName = "GetTestData"
	GetUserOperation                                       OperationName = "GetUser"
	GetVoipCarrierOperation                                OperationName = "GetVoipCarrier"
	GetWebhookOperation                                    OperationName = "GetWebhook"
	GetWebhookSecretOperation                              OperationName = "GetWebhookSecret"
	ListAccountsOperation                                  OperationName = "ListAccounts"
	ListAlertsOperation                                    OperationName = "ListAlerts"
	ListAlertsByAccountOperation                           OperationName = "ListAlertsByAccount"
	ListApplicationsOperation                              OperationName = "ListApplications"
	ListCallsOperation                                     OperationName = "ListCalls"
	ListConferencesOperation                               OperationName = "ListConferences"
	ListGoogleCustomVoicesOperation                        OperationName = "ListGoogleCustomVoices"
	ListLeastCostRoutingCarrierSetEntriesOperation         OperationName = "ListLeastCostRoutingCarrierSetEntries"
	ListLeastCostRoutingRoutesOperation                    OperationName = "ListLeastCostRoutingRoutes"
	ListLeastCostRoutingsOperation                         OperationName = "ListLeastCostRoutings"
	ListMsTeamsTenantsOperation                            OperationName = "ListMsTeamsTenants"
	ListPredefinedCarriersOperation                        OperationName = "ListPredefinedCarriers"
	ListPricesOperation                                    OperationName = "ListPrices"
	ListProvisionedPhoneNumbersOperation                   OperationName = "ListProvisionedPhoneNumbers"
	ListQueuesOperation                                    OperationName = "ListQueues"
	ListRecentCallsOperation                               OperationName = "ListRecentCalls"
	ListRecentCallsBySPOperation                           OperationName = "ListRecentCallsBySP"
	ListRegisteredSipUsersOperation                        OperationName = "ListRegisteredSipUsers"
	ListRegisteredSipUsersByUsernameOperation              OperationName = "ListRegisteredSipUsersByUsername"
	ListSbcsOperation                                      OperationName = "ListSbcs"
	ListServiceProvidersOperation                          OperationName = "ListServiceProviders"
	ListSipGatewaysOperation                               OperationName = "ListSipGateways"
	ListSmppGatewaysOperation                              OperationName = "ListSmppGateways"
	ListSmppsOperation                                     OperationName = "ListSmpps"
	ListSpeechCredentialsOperation                         OperationName = "ListSpeechCredentials"
	ListUsersOperation                                     OperationName = "ListUsers"
	ListVoipCarriersOperation                              OperationName = "ListVoipCarriers"
	LoginOperation                                         OperationName = "Login"
	LoginUserOperation                                     OperationName = "LoginUser"
	LogoutUserOperation                                    OperationName = "LogoutUser"
	ManageSubscriptionOperation                            OperationName = "ManageSubscription"
	ProvisionPhoneNumberOperation                          OperationName = "ProvisionPhoneNumber"
	PutTenantOperation                                     OperationName = "PutTenant"
	RegisterUserOperation                                  OperationName = "RegisterUser"
	RetrieveInvoiceOperation                               OperationName = "RetrieveInvoice"
	SendActivationCodeOperation                            OperationName = "SendActivationCode"
	SupportedLanguagesAndVoicesOperation                   OperationName = "SupportedLanguagesAndVoices"
	SupportedLanguagesAndVoicesByAccountOperation          OperationName = "SupportedLanguagesAndVoicesByAccount"
	SynthesizeOperation                                    OperationName = "Synthesize"
	TestSpeechCredentialOperation                          OperationName = "TestSpeechCredential"
	TestSpeechCredentialByAccountOperation                 OperationName = "TestSpeechCredentialByAccount"
	UpdateAccountOperation                                 OperationName = "UpdateAccount"
	UpdateApplicationOperation                             OperationName = "UpdateApplication"
	UpdateCallOperation                                    OperationName = "UpdateCall"
	UpdateGoogleCustomVoiceOperation                       OperationName = "UpdateGoogleCustomVoice"
	UpdateLeastCostRoutingOperation                        OperationName = "UpdateLeastCostRouting"
	UpdateLeastCostRoutingCarrierSetEntryOperation         OperationName = "UpdateLeastCostRoutingCarrierSetEntry"
	UpdateLeastCostRoutingRouteOperation                   OperationName = "UpdateLeastCostRoutingRoute"
	UpdateLeastCostRoutingRoutesAndCarrierEntriesOperation OperationName = "UpdateLeastCostRoutingRoutesAndCarrierEntries"
	UpdatePhoneNumberOperation                             OperationName = "UpdatePhoneNumber"
	UpdateServiceProviderOperation                         OperationName = "UpdateServiceProvider"
	UpdateSipGatewayOperation                              OperationName = "UpdateSipGateway"
	UpdateSmppGatewayOperation                             OperationName = "UpdateSmppGateway"
	UpdateSpeechCredentialOperation                        OperationName = "UpdateSpeechCredential"
	UpdateSpeechCredentialByAccountOperation               OperationName = "UpdateSpeechCredentialByAccount"
	UpdateUserOperation                                    OperationName = "UpdateUser"
	UpdateVoipCarrierOperation                             OperationName = "UpdateVoipCarrier"
	ValidateActivationCodeOperation                        OperationName = "ValidateActivationCode"
	ValidateInviteCodeOperation                            OperationName = "ValidateInviteCode"
)
