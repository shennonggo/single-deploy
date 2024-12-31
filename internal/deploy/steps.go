package deploy

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/shennonggo/single-deploy/internal/config"
	"github.com/shennonggo/single-deploy/internal/utils"
)

type DeployStep struct {
	Name    string
	Execute func(config.Project) error
}

func GetDeploySteps() []DeployStep {
	return []DeployStep{
		{"Pull Code", pullCode},
		{"Check Environment", checkEnvironment},
		{"Build Project", buildProject},
		{"Start Service", startService},
		{"Health Check", healthCheck},
	}
}

func checkEnvironment(p config.Project) error {
	if !utils.DirExists(p.Path) {
		return fmt.Errorf("project directory does not exist: %s", p.Path)
	}
	return nil
}

func buildProject(p config.Project) error {
	if p.BuildCmd == "" {
		return nil
	}
	cmd := exec.Command("sh", "-c", p.BuildCmd)
	cmd.Dir = p.Path

	// Set up standard output and standard error output
	cmd.Stdout = utils.NewRealTimeWriter(os.Stdout)
	cmd.Stderr = utils.NewRealTimeWriter(os.Stderr)

	return cmd.Run()
}

func startService(p config.Project) error {
	if p.StartCmd == "" {
		return nil
	}
	cmd := exec.Command("sh", "-c", p.StartCmd)
	cmd.Dir = p.Path
	return utils.ExecCommand(cmd)
}
