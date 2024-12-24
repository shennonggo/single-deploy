package deploy

import (
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"jihulab.com/commontool/deployer/internal/config"
	"jihulab.com/commontool/deployer/internal/utils"
)

func Start(configPath string) error {
	// 加载配置
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return err
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	// 选择项目
	project, err := selectProject(cfg.Projects)
	if err != nil {
		return err
	}

	// 创建spinner
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "🚀 "

	// 执行部署步骤
	steps := GetDeploySteps()
	for _, step := range steps {
		s.Suffix = fmt.Sprintf(" %s...", step.Name)
		s.Start()

		if err := step.Execute(project); err != nil {
			s.Stop()
			return fmt.Errorf("%s失败: %v", step.Name, err)
		}

		s.Stop()
		utils.LogSuccess("%s 完成", step.Name)
	}

	utils.LogSuccess("项目 %s 部署成功！", project.Name)
	return nil
}

func selectProject(projects []config.Project) (config.Project, error) {
	var selected string
	options := make([]string, len(projects))
	projectMap := make(map[string]config.Project)

	for i, p := range projects {
		options[i] = p.Name
		projectMap[p.Name] = p
	}

	prompt := &survey.Select{
		Message: "请选择要部署的项目:",
		Options: options,
	}

	if err := survey.AskOne(prompt, &selected); err != nil {
		return config.Project{}, err
	}

	return projectMap[selected], nil
}
