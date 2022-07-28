package infrastructure

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/winnix/shifting/api/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Context struct {
	Db        *gorm.DB
	Config    *config.Config
	Ctx       context.Context
	GinCtx    *gin.Context
	Validator *validator.Validate
	Logger    *zap.Logger
	User      *CurrentUser
}

type CurrentUser struct {
	Id          string
	AccountCode string
	PropertyId  string
}

type HandlerResult struct {
	Status int
	Value  interface{}
}

type HandlerFunc func(*Context) (*HandlerResult, error)
