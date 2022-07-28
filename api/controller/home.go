package controller

import (
	"fmt"

	infra "github.com/winnix/shifting/api/infrastructure"
	"github.com/winnix/shifting/api/service"
)

type homeRequest struct {
	ApaOneToken string `form:"token" validate:"required"`
	AccountCode string `form:"accountCode" validate:"required"`
	PropertyId  string `form:"propertyId" validate:"required"`
	SubjectId   string `form:"subjectId" validate:"required"`
}

func Home(ctx *infra.Context) (*infra.HandlerResult, error) {
	model := homeRequest{}
	if err := ctx.GinCtx.BindQuery(&model); err != nil {
		ctx.Logger.Sugar().Infof("Home: failed binding request model: %s\n", err.Error())
		return BadRequest("")
	}

	if err := ctx.Validator.Struct(model); err != nil {
		ctx.Logger.Sugar().Infof("Home: invalid request model: %s\n", err.Error())
		return BadRequest("")
	}

	accToken, err := service.
		NewAccountTokenService(ctx.Db, model.AccountCode).
		GetByApaleoToken(model.ApaOneToken)
	if err != nil {
		ctx.Logger.Sugar().Infof("Home: apaleo account '%s' is not configured: %s\n", model.AccountCode, err.Error())
		return BadRequest(fmt.Sprintf("apaleo account '%s' is not configured", model.AccountCode))
	}

	fmt.Println(accToken.AccountCode)
	// TODO:
	// - call GET identity/users/{subjectId}
	// - if user.role in [AccountAdmin, PropertyAdmin, SeniorFrontDesk, FrontDesk] - allow

	authToken, err := service.
		NewJwtService(ctx.Config.JwtSigningKey).
		GenerateToken(model.AccountCode, model.PropertyId, model.SubjectId)
	if err != nil {
		ctx.Logger.Sugar().Errorf("failed making auth token: %s\n", err.Error())
		return ServerError()
	}

	return Ok(authToken)
}
