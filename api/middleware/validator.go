package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var dateInFuture validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if date.Before(today) {
			return false
		}
	}

	return true
}

func Validator() gin.HandlerFunc {
	validate := validator.New()
	validate.RegisterValidation("dateinfuture", dateInFuture)

	return func(c *gin.Context) {
		c.Set("validator", validate)
		c.Next()
	}
}
