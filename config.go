package main

import (
	"time"
)

type Config struct {
	Web struct {
		Addr            string        `default:":3000" required:"true"`
		ReadTimeout     time.Duration `default:"5s"`
		WriteTimeout    time.Duration `default:"5s"`
		ShutdownTimeout time.Duration `default:"5s"`
	}
	Debug struct {
		Port     string `default:":30001" required:"true"`
		Metrics  string `default:"/metrics"`
		Profiler string `default:"/debug/pprof/"`
	}
	Elastic struct {
		EnableDebugLogger bool     `envconfig:"APP_ELASTICSEARCH_DEBUG"`
		EnableMetrics     bool     `envconfig:"APP_ELASTICSEARCH_METRICS"`
		Url               []string `envconfig:"APP_ELASTICSEARCH_URL" default:"elastic:9200" required:"true"`
	}
	Log struct {
		Level    string `default:"error"`
		Facility string `default:"finder"`
	}
}
