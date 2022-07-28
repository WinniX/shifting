package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/winnix/shifting/api/config"
	"github.com/winnix/shifting/api/domain"
	"github.com/winnix/shifting/api/utils/gormzap"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database(logger *zap.Logger, cfg *config.Config) gin.HandlerFunc {
	dbLogger := gormzap.New(zap.L())
	db, err := gorm.Open(
		postgres.Open(cfg.DSN),
		&gorm.Config{Logger: dbLogger})
	if err != nil {
		logger.Fatal("Can't open DB connection", zap.Error(err))
	}

	RunMigrations(db, logger)

	return func(c *gin.Context) {
		c.Set("dsn", db)
		c.Next()
	}
}

func RunMigrations(db *gorm.DB, logger *zap.Logger) {
	err := db.AutoMigrate(&domain.Shift{}, &domain.AuthRequest{}, &domain.AccountToken{})
	if err != nil {
		logger.Fatal("Can't run DB migrations", zap.Error(err))
	}
}
