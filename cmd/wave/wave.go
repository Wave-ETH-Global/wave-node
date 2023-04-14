package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Wave-ETH-Global/wave-node/app"
	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/router"
	"github.com/google/logger"
	"github.com/labstack/echo/v4"
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
			provideApp,
			provideRouter,
			provideLogging,
		),
		fx.Invoke(
			runApp,
		),
		fx.NopLogger,
	).Run()

}

func runApp(lc fx.Lifecycle, a *app.App, r *echo.Echo, l *logger.Logger) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info("Starting Wave node...")
				return a.Run(ctx, r)
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

func provideApp(cfg *config.Config) *app.App {
	a, err := app.NewApp(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	return a
}

func provideRouter(a *app.App) *echo.Echo {
	r, err := router.NewRouter(a)
	if err != nil {
		logger.Fatal(err)
	}
	return r
}
