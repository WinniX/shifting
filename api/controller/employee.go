package controller

import (
	"github.com/google/uuid"
	apaproxy "github.com/winnix/shifting/api/apaproxy"
	infra "github.com/winnix/shifting/api/infrastructure"
)

func GetUser(ctx *infra.Context) (*infra.HandlerResult, error) {
	strID := ctx.GinCtx.Param("id")
	id, err := uuid.Parse(strID)
	if err != nil {
		ctx.Logger.Sugar().Infof("GetUser: invalid request: %s\n", err.Error())
		return BadRequest("")
	}

	uID := id.String()
	user, found, err := apaproxy.GetUser(ctx, uID)
	if err != nil {
		ctx.Logger.Sugar().Errorf("GetUser: get user with id=%s failed: %s", uID, err.Error())
		return ServerError()
	}

	if !found {
		ctx.Logger.Sugar().Infof("GetUser: user with id=%s was not found\n", uID)
		return NotFound("user not found")
	}

	return Ok(user)
}
