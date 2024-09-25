package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/httplog/v2"
	"github.com/go-viper/mapstructure/v2"
	"github.com/knadh/koanf/parsers/toml/v2"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/optician/meeting-room-booking/internal"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample().Sugar()

	configurator := koanf.New(".")
	parser := toml.Parser()

	// os.Args because there is only one command line argument
	var configPath string
	if len(os.Args) < 2 {
		logger.Fatal("not enough arguments: run 'main -- /cfg/path'")
	} else {
		configPath = os.Args[2]
	}

	if err := configurator.Load(file.Provider(configPath), parser); err != nil {
		logger.Fatalf("error loading config: %v", err)
	}

	var config internal.Config
	unmarshalConf := koanf.UnmarshalConf{
		DecoderConfig: &mapstructure.DecoderConfig{
			ErrorUnset:  true,
			ErrorUnused: true,
			Result:      &config,
		},
	}
	if err := configurator.UnmarshalWithConf("", &config, unmarshalConf); err != nil {
		logger.Fatalf("error loading config: %v", err)
	}
	logger.Infof("config: %v", configurator.Raw())

	httpLogger := httplog.NewLogger("httplog", httplog.Options{
		// JSON:             true,
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})
	router := internal.Make(httpLogger, logger, &config)
	http.ListenAndServe(":3000", router)
}
