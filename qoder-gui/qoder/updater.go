package qoder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	GitHubRepo    = "70548887/QoderSessionManager"
	GitHubAPIBase = "https://api.github.com/repos/" + GitHubRepo
)

type ReleaseInfo struct {
	TagName string  `json:"tag_name"`
	Name    string  `json:"name"`
	Body    string  `json:"body"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
	Size               int64  `json:"size"`
}

// CheckUpdate 检查是否有新版本
func CheckUpdate() (*ReleaseInfo, bool, error) {
	resp, err := http.Get(GitHubAPIBase + "/releases/latest")
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, false, fmt.Errorf("GitHub API returned %d", resp.StatusCode)
	}

	var release ReleaseInfo
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, false, err
	}

	remoteVer := strings.TrimPrefix(release.TagName, "v")
	localVer := strings.TrimPrefix(Version, "v")

	if remoteVer != localVer && remoteVer > localVer {
		return &release, true, nil
	}
	return &release, false, nil
}

// DoUpdate 执行升级
func DoUpdate(release *ReleaseInfo) error {
	assetName := getAssetName()
	var downloadURL string
	for _, asset := range release.Assets {
		if asset.Name == assetName {
			downloadURL = asset.BrowserDownloadURL
			break
		}
	}
	if downloadURL == "" {
		return errors.New("未找到适用于当前平台的安装包")
	}

	execPath, err := os.Executable()
	if err != nil {
		return err
	}
	execPath, _ = filepath.EvalSymlinks(execPath)

	tmpPath := execPath + ".new"
	if err := downloadFile(tmpPath, downloadURL); err != nil {
		os.Remove(tmpPath)
		return fmt.Errorf("下载失败: %w", err)
	}

	backupPath := execPath + ".bak"
	os.Remove(backupPath)
	if err := os.Rename(execPath, backupPath); err != nil {
		os.Remove(tmpPath)
		return fmt.Errorf("备份当前版本失败: %w", err)
	}

	if err := os.Rename(tmpPath, execPath); err != nil {
		os.Rename(backupPath, execPath)
		return fmt.Errorf("替换失败，已回滚: %w", err)
	}

	os.Remove(backupPath)
	return nil
}

func getAssetName() string {
	exe, _ := os.Executable()
	baseName := filepath.Base(exe)
	name := strings.TrimSuffix(baseName, ".exe")
	return fmt.Sprintf("%s-%s-%s.exe", name, runtime.GOOS, runtime.GOARCH)
}

func downloadFile(destPath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
