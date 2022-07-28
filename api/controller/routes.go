package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/winnix/shifting/api/config"
	infra "github.com/winnix/shifting/api/infrastructure"
	"github.com/winnix/shifting/api/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Ok(val interface{}) (*infra.HandlerResult, error) {
	return &infra.HandlerResult{
		Status: http.StatusOK,
		Value:  val,
	}, nil
}

func BadRequest(message string) (*infra.HandlerResult, error) {
	return &infra.HandlerResult{
		Status: http.StatusBadRequest,
		Value:  message,
	}, nil
}

func NotFound(message string) (*infra.HandlerResult, error) {
	return &infra.HandlerResult{
		Status: http.StatusNotFound,
		Value:  message,
	}, nil
}

func ServerError() (*infra.HandlerResult, error) {
	return nil, errors.New("something went wrong")
}

func getContext(c *gin.Context) *infra.Context {
	accountCode := c.GetString("accountCode")
	propertyId := c.GetString("propertyId")
	userId := c.GetString("subjectId")

	var user *infra.CurrentUser = nil
	if len(accountCode) > 0 && len(propertyId) > 0 && len(userId) > 0 {
		user = &infra.CurrentUser{
			Id:          userId,
			AccountCode: accountCode,
			PropertyId:  propertyId,
		}
	}

	return &infra.Context{
		Db:        c.MustGet("dsn").(*gorm.DB),
		Config:    c.MustGet("config").(*config.Config),
		Ctx:       context.Background(),
		GinCtx:    c,
		Validator: c.MustGet("validator").(*validator.Validate),
		Logger:    c.MustGet("logger").(*zap.Logger),
		User:      user,
	}
}

func wrapper(f infra.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := getContext(c)

		res, err := f(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else if res != nil {
			c.JSON(res.Status, gin.H{"result": res.Value})
		}
	}
}

func SetRoutes(r *gin.Engine, cfg *config.Config) {
	r.GET("/install", wrapper(Install))
	r.GET("/auth-callback", wrapper(AuthCallback))

	r.GET("/home", wrapper(Home))

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired(cfg))
	{
		authorized.GET("/employee/:id", wrapper(GetUser))
		authorized.POST("/schedule", wrapper(CreateSchedule))

		authorized.GET("/week/:date", wrapper(GetWeekByDate))
		authorized.GET("/shift/:id", wrapper(GetShift))
	}
}
