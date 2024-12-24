package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"jihulab.com/commontool/deployer/internal/deploy"
	"jihulab.com/commontool/deployer/internal/utils"
)

func main() {
	app := &cli.App{
		Name:  "deployer",
		Usage: "通用项目部署工具",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "configs/deploy-config.json",
				Usage:   "配置文件路径",
			},
		},
		Action: runDeploy,
	}

	if err := app.Run(os.Args); err != nil {
		utils.LogError("部署失败: %v", err)
		os.Exit(1)
	}
}

func runDeploy(c *cli.Context) error {
	return deploy.Start(c.String("config"))
}
