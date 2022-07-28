package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/winnix/shifting/api/config"
	"github.com/winnix/shifting/api/controller"
	"github.com/winnix/shifting/api/middleware"
	"go.uber.org/zap"
)

type options struct {
	config string
}

func initOptions() options {
	var opts options
	flag.StringVar(&opts.config, "config", "", "Path to config")
	flag.Parse()
	return opts
}

func main() {
	opts := initOptions()
	logger, err := newLoggerConfig().Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error configuring logging: %v", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Sugar().Infow("Args", "Args", os.Args)

	cfg, err := config.LoadConfig(opts.config)
	if err != nil {
		logger.Fatal("Can't read config", zap.Error(err), zap.String("configPath", opts.config))
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Database(logger, cfg))
	router.Use(middleware.Logger(logger))
	router.Use(middleware.Validator())
	router.Use(middleware.Config(cfg))

	controller.SetRoutes(router, cfg)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.Port),
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func newLoggerConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}
