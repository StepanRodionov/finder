package main

import (
	"context"
	"finder/internal/log"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"syscall"
)

var (
	envPrefix = "FINDER"
	appName   = "finder"
	gitTag    = "[not set]"
)

func main() {
	var config Config
	if err := envconfig.Process(envPrefix, &config); err != nil {
		panic(err)
	}

	logger, err := log.New(config.Log.Level, config.Log.Facility, gitTag)
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg := elasticsearch.Config{
		Addresses:         config.Elastic.Url,
		EnableMetrics:     config.Elastic.EnableMetrics,
		EnableDebugLogger: config.Elastic.EnableDebugLogger,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	storage := storage{
		logger: logger,
		ec:     es,
	}
	// TODO - fill dictionaries
	//if err := storage.Load(ctx, config.Storage.Index, config.Storage.UpdateInterval); err != nil {
	//	panic(err)
	//}

	errors := make(chan error, 1)

	server := http.Server{
		Addr:         config.Web.Addr,
		Handler:      createHTTPHandler(config, &storage),
		ReadTimeout:  config.Web.ReadTimeout,
		WriteTimeout: config.Web.WriteTimeout,
	}
	go func() {
		logger.Info("httpserver.started", zap.String("port", server.Addr))
		errors <- server.ListenAndServe()
	}()

	select {
	case err = <-errors:
		if err != nil {
			logger.Error("server.start", zap.Error(err))
		}
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), config.Web.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error("server.shutdown", zap.Error(err))
			if err := server.Close(); err != nil {
				logger.Error("server.close", zap.Error(err))
			}
		}
	}
}
