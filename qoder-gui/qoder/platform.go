package qoder

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// GetQoderBasePath 返回 Qoder IDE 数据存储的基路径
func GetQoderBasePath() string {
	homeDir, _ := os.UserHomeDir()
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(homeDir, "Library", "Application Support", "Qoder")
	case "windows":
		if appData := os.Getenv("APPDATA"); appData != "" {
			return filepath.Join(appData, "Qoder")
		}
		return filepath.Join(homeDir, "AppData", "Roaming", "Qoder")
	case "linux":
		return filepath.Join(homeDir, ".config", "Qoder")
	default:
		return filepath.Join(homeDir, ".qoder")
	}
}

// GetDefaultBackupDir 返回默认备份目录
func GetDefaultBackupDir() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, "Documents", "qoder-backups")
}

// OpenBrowser 跨平台打开浏览器
func OpenBrowser(url string) {
	switch runtime.GOOS {
	case "darwin":
		exec.Command("open", url).Start()
	case "windows":
		exec.Command("cmd", "/c", "start", "", url).Start()
	case "linux":
		exec.Command("xdg-open", url).Start()
	}
}
