package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/database"
	"github.com/Wave-ETH-Global/wave-node/router"
	"github.com/google/logger"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/mcuadros/go-defaults"
	"github.com/pelletier/go-toml/v2"
	"go.uber.org/fx"
)

type AppFlags struct {
	GenerateConfig bool
	VerboseLogging bool
	SyslogLogging  bool
	ConfigPath     string
}

func main() {
	fx.New(
		fx.Provide(
			provideAppFlags,
			provideConfig,
			provideRouter,
			provideLogging,
			database.ProvideDatabase,
		),
		fx.Invoke(
			runRouter,
		),
		fx.NopLogger,
	).Run()

}

func runRouter(lc fx.Lifecycle, r *echo.Echo, l *logger.Logger, cfg *config.Config, db *sqlx.DB) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info("Starting Wave node...")
				port := cfg.HTTP.Port
				addr := cfg.HTTP.Address
				logger.Infof("HTTP server starts listening at %s:%d", addr, port)
				go r.Start(fmt.Sprintf("%s:%d", addr, port))
				return nil
			},
			OnStop: func(ctx context.Context) error {
				l.Close()
				return nil
			},
		},
	)
}

func provideLogging(flags *AppFlags) *logger.Logger {
	return logger.Init("wave", flags.VerboseLogging, flags.SyslogLogging, ioutil.Discard)
}

func provideAppFlags() *AppFlags {
	var flags AppFlags

	flag.StringVar(&flags.ConfigPath, "config", "", "Path to config")
	flag.BoolVar(&flags.GenerateConfig, "genconfig", false, "Generate new config")
	flag.BoolVar(&flags.VerboseLogging, "verbose", true, "Verbose logging")
	flag.BoolVar(&flags.SyslogLogging, "syslog", false, "Log to system logging daemon")
	flag.Parse()

	return &flags
}

func provideConfig(l *logger.Logger, flags *AppFlags) *config.Config {
	if flags.GenerateConfig {
		config := config.Config{}
		defaults.SetDefaults(&config)
		configStr, err := toml.Marshal(config)
		if err != nil {
			l.Fatalf("Cannot generate config! %s", err.Error())
		}
		fmt.Print(string(configStr))
		os.Exit(0)
	}

	if flags.ConfigPath == "" {
		l.Fatal("no config path provided")
	}

	cfg, err := config.NewConfig(flags.ConfigPath)
	if err != nil {
		l.Fatalf("failed to load config: %v", err)
	}

	return &cfg
}

func provideRouter(cfg *config.Config) *echo.Echo {
	r, err := router.NewRouter(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	return r
}
