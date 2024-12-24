package deploy

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"jihulab.com/commontool/deployer/internal/config"
)

func healthCheck(p config.Project) error {
	if p.HealthCheck.URL == "" {
		return nil
	}

	timeout := time.Duration(p.HealthCheck.Timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("健康检查超时")
		case <-ticker.C:
			if err := checkEndpoint(p.HealthCheck.URL); err == nil {
				return nil
			}
		}
	}
}

func checkEndpoint(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("健康检查失败，状态码: %d", resp.StatusCode)
	}
	return nil
}
