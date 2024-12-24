package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ExecCommand(cmd *exec.Cmd) error {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("执行命令失败: %v\n输出: %s\n错误: %s",
			err, stdout.String(), stderr.String())
	}

	LogInfo(stdout.String())
	return nil
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
