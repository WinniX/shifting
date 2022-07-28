package apaproxy

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	sw "github.com/winnix/shifting/api/clients/integrationclient"
	infra "github.com/winnix/shifting/api/infrastructure"
)

const UI_INTEGRATION_LABEL = "Shifting"
const UI_INTEGRATION_SOURCETYPE = "Public"
const UI_INTEGRATION_TARGET = "PropertyMenuApps"

func newIntegrationClient() *sw.APIClient {
	cfg := sw.NewConfiguration()
	cfg.Host = "integration.apaleo.com"
	cfg.Scheme = "https"
	return sw.NewAPIClient(cfg)
}

func addIntegration(c *sw.APIClient, authCtx context.Context, homeUrl string) error {
	req := c.UiIntegrationsApi.IntegrationUiIntegrationsByTargetPost(authCtx, UI_INTEGRATION_TARGET)
	model := sw.NewCreateUiIntegrationModel(UI_INTEGRATION_LABEL, homeUrl, UI_INTEGRATION_SOURCETYPE)
	req = req.Body(*model)
	_, _, err := req.Execute()
	if err != nil {
		return fmt.Errorf("add integration failed: %s", err.Error())
	}

	return nil
}

func SaveIntegration(ctx *infra.Context, homeToken string) error {
	token, err := GetApaleoToken(ctx)
	if err != nil {
		return fmt.Errorf("SaveIntegration: can't get apaleo token: %s", err.Error())
	}

	client := newIntegrationClient()
	authCtx := context.WithValue(ctx.Ctx, sw.ContextAccessToken, token)
	items, res, err := client.UiIntegrationsApi.
		IntegrationUiIntegrationsByTargetGet(authCtx, UI_INTEGRATION_TARGET).
		Execute()
	if err != nil && res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("SaveIntegration: apaleo api call failed %s: %s", res.Status, err.Error())
	}

	homeUrl := fmt.Sprintf("%s/home?token=%s", ctx.Config.Host, homeToken)
	if items.Count == 0 {
		if err := addIntegration(client, authCtx, homeUrl); err != nil {
			return fmt.Errorf("SaveIntegration: apaleo api call failed: %s", err.Error())
		}

		return nil
	}

	for _, item := range items.UiIntegrations {
		if strings.HasPrefix(item.SourceUrl, homeUrl) {
			return nil
		}
	}

	if err := addIntegration(client, authCtx, homeUrl); err != nil {
		return fmt.Errorf("SaveIntegration: apaleo api call failed: %s", err.Error())
	}

	return nil
}
