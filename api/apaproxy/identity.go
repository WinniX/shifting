package apaproxy

import (
	"context"
	"fmt"
	"net/http"

	sw "github.com/winnix/shifting/api/clients/identityclient"
	infra "github.com/winnix/shifting/api/infrastructure"
)

func newIdentityClient() *sw.APIClient {
	cfg := sw.NewConfiguration()
	cfg.Host = "identity.apaleo.com"
	cfg.Scheme = "https"
	return sw.NewAPIClient(cfg)
}

func GetUser(ctx *infra.Context, id string) (*sw.UserModel, bool, error) {
	token, err := GetApaleoToken(ctx)
	if err != nil {
		return nil, false, fmt.Errorf("GetUser: can't get apaleo token: %s", err.Error())
	}

	client := newIdentityClient()
	authCtx := context.WithValue(ctx.Ctx, sw.ContextAccessToken, token)
	req := client.UsersApi.ApiUsersByUserIdGet(authCtx, id)
	user, res, err := req.Execute()
	if err != nil {
		if res.StatusCode != http.StatusNotFound {
			return nil, false, fmt.Errorf("GetUser: apaleo api call failed %s: %s", res.Status, err.Error())
		}

		return nil, false, nil
	}

	return &user, true, nil
}
