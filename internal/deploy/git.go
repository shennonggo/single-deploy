package deploy

import (
	"os/exec"

	"github.com/shennonggo/single-deploy/internal/config"
	"github.com/shennonggo/single-deploy/internal/utils"
)

func pullCode(p config.Project) error {
	if p.GitRepo == "" {
		return nil // Skip if git repository is not configured
	}

	// Check if directory exists
	if !utils.DirExists(p.Path) {
		// If directory doesn't exist, execute clone
		return gitClone(p)
	}

	// If directory exists, execute pull
	return gitPull(p)
}

func gitClone(p config.Project) error {
	cmd := exec.Command("git", "clone", "-b", p.GitBranch, p.GitRepo, p.Path)
	return utils.ExecCommand(cmd)
}

func gitPull(p config.Project) error {
	cmd := exec.Command("git", "pull", "origin", p.GitBranch)
	cmd.Dir = p.Path
	return utils.ExecCommand(cmd)
}
