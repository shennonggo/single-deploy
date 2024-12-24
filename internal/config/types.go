package config

type Project struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Type        string `json:"type"`
	GitRepo     string `json:"gitRepo"`
	GitBranch   string `json:"gitBranch"`
	BuildCmd    string `json:"buildCmd"`
	StartCmd    string `json:"startCmd"`
	HealthCheck struct {
		URL     string `json:"url"`
		Timeout int    `json:"timeout"`
	} `json:"healthCheck"`
}

type Config struct {
	Projects []Project `json:"projects"`
}
