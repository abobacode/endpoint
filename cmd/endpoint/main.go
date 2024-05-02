package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/abobacode/endpoint/config"
	"github.com/abobacode/endpoint/internal/repo"
	"github.com/abobacode/endpoint/internal/service"
	"github.com/abobacode/endpoint/internal/usecase"
	stdout "github.com/abobacode/endpoint/pkg/log"
	"github.com/abobacode/endpoint/pkg/signal"
)

const title = "API For Example"

func main() {
	application := cli.App{
		Name: title,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config-file",
				Required: true,
				Value:    "config.yaml",
				Usage:    "YAML config filepath",
				EnvVars:  []string{"CONFIG_FILE"},
				FilePath: "/srv/lime_secrets/config_file",
			},
			&cli.StringFlag{
				Name:     "bind-address",
				Usage:    "IP и порт сервера, например: 0.0.0.0:3000",
				Required: false,
				Value:    "0.0.0.0:3000",
				EnvVars:  []string{"BIND_ADDRESS"},
			},
			&cli.IntFlag{
				Name:     "listener",
				Usage:    "Unix socket or TCP",
				Required: false,
				Value:    1,
				EnvVars:  []string{"LISTENER"},
			},
		},
		Action: Main,
	}

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Main(ctx *cli.Context) error {
	appContext, cancel := context.WithCancel(ctx.Context)
	defer func() {
		cancel()
		log.Println("app context is canceled, service is down!")
	}()

	cfg, err := config.New(ctx.String("config-file"))
	if err != nil {
		return err
	}

	apiKey := flag.String("api-key", cfg.YouTube.ApiKey, "YouTube API Key")

	point, err := service.New(context.Background(), &service.Options{
		Database: &cfg.Database,
	})
	if err != nil {
		return err
	}

	defer func() {
		point.Shutdown(func(err error) {
			stdout.Warning(err)
		})
		point.Stacktrace()
	}()

	await, stop := signal.Notifier(func() {
		log.Println("Service start shutdown process..")
	})

	pointCase := usecase.NewApp(repo.New(point.Pool))

	go func() {
		if err := pointCase.App(point.Context(), appContext, ctx, apiKey); err != nil {
			stop(err)
		}
	}()

	return await()
}
