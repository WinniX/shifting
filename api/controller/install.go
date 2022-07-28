package controller

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	apaproxy "github.com/winnix/shifting/api/apaproxy"
	"github.com/winnix/shifting/api/domain"
	infra "github.com/winnix/shifting/api/infrastructure"
	"github.com/winnix/shifting/api/service"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func Install(ctx *infra.Context) (*infra.HandlerResult, error) {
	authReq := domain.AuthRequest{}
	if err := ctx.Db.Create(&authReq).Error; err != nil {
		ctx.Logger.Sugar().Errorf("failed creating auth request: %s\n", err.Error())
		return ServerError()
	}

	oauthCfg := apaproxy.GetOAuthConfig(ctx.Config)
	authUrl := oauthCfg.AuthCodeURL(authReq.State.String(), oauth2.AccessTypeOffline)
	ctx.GinCtx.Redirect(http.StatusTemporaryRedirect, authUrl)

	return nil, nil
}

func AuthCallback(ctx *infra.Context) (*infra.HandlerResult, error) {
	state := ctx.GinCtx.Query("state")
	stateU, err := uuid.Parse(state)
	if err != nil {
		ctx.Logger.Sugar().Infof("AuthCallback: received invalid state: %s\n", err.Error())
		return BadRequest("")
	}

	code := ctx.GinCtx.Query("code")
	if len(code) == 0 {
		ctx.Logger.Sugar().Infof("AuthCallback: code is missing\n")
		return BadRequest("")
	}

	authReq := domain.AuthRequest{}
	err = ctx.Db.Where(&domain.AuthRequest{State: stateU}).First(&authReq).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.Logger.Sugar().Infof("AuthCallback: state not found\n")
		return BadRequest("")
	}

	cfg := apaproxy.GetOAuthConfig(ctx.Config)
	token, err := cfg.Exchange(ctx.Ctx, code)
	if err != nil {
		ctx.Logger.Sugar().Infof("AuthCallback: code exchange failed: %s\n", err.Error())
		return BadRequest("")
	}

	claims := apaproxy.ApaleoClaims{}
	jwt.ParseWithClaims(token.AccessToken, &claims, nil)
	if err = claims.Valid(); err != nil {
		ctx.Logger.Sugar().Errorf("error parsing access token: %s\n", err.Error())
		return ServerError()
	}

	accToken, err := service.
		NewAccountTokenService(ctx.Db, claims.AccountCode).
		GetCurrent()
	// if record already exists - we'll update it, otherwise - create
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.Logger.Sugar().Errorf("error getting account token: %s\n", err.Error())
		return ServerError()
	}

	accToken.AccountCode = claims.AccountCode
	accToken.ApaleoOneToken = uuid.New().String()
	if err := apaproxy.SaveAccountToken(ctx, accToken, token); err != nil {
		ctx.Logger.Sugar().Errorf("failed saving account token: %s\n", err.Error())
		return ServerError()
	}

	// a bit lame
	ctx.User = &infra.CurrentUser{
		AccountCode: accToken.AccountCode,
	}

	if err := apaproxy.SaveIntegration(ctx, accToken.ApaleoOneToken); err != nil {
		ctx.Logger.Sugar().Errorf("failed saving ui integration: %s\n", err.Error())
		return ServerError()
	}

	return Ok("Successfully connected to apaleo")
}
