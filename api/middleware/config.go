package middleware

import (
	"github.com/gin-gonic/gin"
	cfg "github.com/winnix/shifting/api/config"
)

func Config(config *cfg.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}
