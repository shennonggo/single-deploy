package deploy

import (
	"fmt"
	"os/exec"

	"jihulab.com/commontool/deployer/internal/config"
	"jihulab.com/commontool/deployer/internal/utils"
)

type DeployStep struct {
	Name    string
	Execute func(config.Project) error
}

func GetDeploySteps() []DeployStep {
	return []DeployStep{
		{"拉取代码", pullCode},
		{"检查环境", checkEnvironment},
		{"构建项目", buildProject},
		{"启动服务", startService},
		{"健康检查", healthCheck},
	}
}

func checkEnvironment(p config.Project) error {
	if !utils.DirExists(p.Path) {
		return fmt.Errorf("项目目录不存在: %s", p.Path)
	}
	return nil
}

func buildProject(p config.Project) error {
	if p.BuildCmd == "" {
		return nil
	}
	cmd := exec.Command("sh", "-c", p.BuildCmd)
	cmd.Dir = p.Path
	return utils.ExecCommand(cmd)
}

func startService(p config.Project) error {
	if p.StartCmd == "" {
		return nil
	}
	cmd := exec.Command("sh", "-c", p.StartCmd)
	cmd.Dir = p.Path
	return utils.ExecCommand(cmd)
}
