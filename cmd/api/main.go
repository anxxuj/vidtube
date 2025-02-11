package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/anxxuj/vidtube/internal/env"
	"github.com/anxxuj/vidtube/internal/jsonlog"
	_ "github.com/joho/godotenv/autoload"
)

// Application version
const version = "1.0.0"

type config struct {
	port int
}

type application struct {
	config config
	logger *jsonlog.Logger
}

func main() {
	cfg := config{
		port: env.GetInt("SERVER_PORT", 4000),
	}

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := &application{
		config: cfg,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
	})

	err := srv.ListenAndServe()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
