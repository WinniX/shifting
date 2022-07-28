package apaproxy

import (
	"errors"
	"fmt"

	"github.com/winnix/shifting/api/config"
	"github.com/winnix/shifting/api/domain"
	infra "github.com/winnix/shifting/api/infrastructure"
	"github.com/winnix/shifting/api/service"
	"golang.org/x/oauth2"
)

type ApaleoClaims struct {
	AccountCode string `json:"account_code"`
}

func GetOAuthConfig(cfg *config.Config) oauth2.Config {
	authCfg := cfg.OAuth["apaleo"]
	return oauth2.Config{
		ClientID:     authCfg.ClientID,
		ClientSecret: authCfg.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   authCfg.AuthUrl,
			TokenURL:  authCfg.TokenUrl,
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: fmt.Sprintf("%s/auth-callback", cfg.Host),
		Scopes:      authCfg.Scopes,
	}
}

func GetApaleoToken(ctx *infra.Context) (string, error) {
	accountCode := ctx.User.AccountCode
	accToken, err := service.NewAccountTokenService(ctx.Db, accountCode).GetCurrent()
	if err != nil {
		return "", fmt.Errorf("GetApaleoToken: can't get account token for %s: %s", accountCode, err.Error())
	}

	token := oauth2.Token{
		AccessToken:  accToken.AccessToken,
		RefreshToken: accToken.RefreshToken,
		Expiry:       accToken.Expiry,
	}
	oauth2cfg := GetOAuthConfig(ctx.Config)

	tokenSource := oauth2cfg.TokenSource(ctx.Ctx, &token)
	newToken, err := tokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("GetApaleoToken: refresh token failed for %s: %s", accountCode, err.Error())
	}

	if newToken.AccessToken != token.AccessToken {
		err := SaveAccountToken(ctx, accToken, newToken)
		if err != nil {
			return "", fmt.Errorf("GetApaleoToken: update account token failed for %s: %s", accountCode, err.Error())
		}
	}

	return newToken.AccessToken, nil
}

func SaveAccountToken(ctx *infra.Context, base *domain.AccountToken, t *oauth2.Token) error {
	base.AccessToken = t.AccessToken
	base.RefreshToken = t.RefreshToken
	base.Expiry = t.Expiry
	if err := ctx.Db.Save(&base).Error; err != nil {
		return fmt.Errorf("SaveAccountToken: failed saving account token: %s", err.Error())
	}

	return nil
}

func (c *ApaleoClaims) Valid() error {
	if len(c.AccountCode) == 0 {
		return errors.New("account_code claim is empty")
	}

	return nil
}
