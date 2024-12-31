package main

import (
	"os"

	"github.com/shennonggo/single-deploy/internal/deploy"
	"github.com/shennonggo/single-deploy/internal/utils"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "single-deploy",
		Usage: "Single Project Deployment Tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "configs/deploy-config.json",
				Usage:   "Configuration file path",
			},
		},
		Action: runDeploy,
	}

	if err := app.Run(os.Args); err != nil {
		utils.LogError("Deployment failed: %v", err)
		os.Exit(1)
	}
}

func runDeploy(c *cli.Context) error {
	return deploy.Start(c.String("config"))
}
