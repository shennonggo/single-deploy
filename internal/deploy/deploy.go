package deploy

import (
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/shennonggo/single-deploy/internal/config"
	"github.com/shennonggo/single-deploy/internal/utils"
)

func Start(configPath string) error {
	// Load configuration
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return err
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	// Select project
	project, err := selectProject(cfg.Projects)
	if err != nil {
		return err
	}

	// Create spinner
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "🚀 "

	// Execute deployment steps
	steps := GetDeploySteps()
	for _, step := range steps {
		s.Suffix = fmt.Sprintf(" %s...", step.Name)
		s.Start()

		if err := step.Execute(project); err != nil {
			s.Stop()
			return fmt.Errorf("%s failed: %v", step.Name, err)
		}

		s.Stop()
		utils.LogSuccess("%s completed", step.Name)
	}

	utils.LogSuccess("Project %s deployed successfully!", project.Name)
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
		Message: "Please select the project to deploy:",
		Options: options,
	}

	if err := survey.AskOne(prompt, &selected); err != nil {
		return config.Project{}, err
	}

	return projectMap[selected], nil
}
