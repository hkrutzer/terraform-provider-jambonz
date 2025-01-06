package api

import (
	"context"
	api "terraform-provider-jambonz/internal/api/generated"
)

type ApiTokenSource struct {
	token string
}

func (m *ApiTokenSource) BearerAuth(ctx context.Context, operationName api.OperationName) (api.BearerAuth, error) {
	return api.BearerAuth{Token: m.token}, nil
}

func NewClient(_ context.Context, baseUrl string, apiKey string) (*api.Client, error) {
	x := ApiTokenSource{token: apiKey}
	return api.NewClient(baseUrl, &x)
}
