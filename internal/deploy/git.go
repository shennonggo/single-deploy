package deploy

import (
	"os/exec"

	"jihulab.com/commontool/deployer/internal/config"
	"jihulab.com/commontool/deployer/internal/utils"
)

func pullCode(p config.Project) error {
	if p.GitRepo == "" {
		return nil // 如果没有配置 git 仓库，跳过
	}

	// 检查目录是否存在
	if !utils.DirExists(p.Path) {
		// 如果目录不存在，执行 clone
		return gitClone(p)
	}

	// 如果目录存在，执行 pull
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
