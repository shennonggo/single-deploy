package utils

import "fmt"

func LogError(format string, args ...interface{}) {
	fmt.Printf("❌ "+format+"\n", args...)
}

func LogSuccess(format string, args ...interface{}) {
	fmt.Printf("✅ "+format+"\n", args...)
}

func LogInfo(format string, args ...interface{}) {
	fmt.Printf("ℹ️ "+format+"\n", args...)
}
