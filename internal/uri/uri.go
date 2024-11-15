package uri

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

func Execute(action, vault, param, targetFolder string) error {
	encodedVault := url.PathEscape(vault)
	encodedParam := url.PathEscape(param)

	var paramName string
	switch action {
	case "search":
		paramName = "query"
	default:
		paramName = "file"

		// Optionally add targetFolder if provided
		if targetFolder != "" {
			encodedFolder := url.PathEscape(targetFolder)
			encodedParam = fmt.Sprintf("%s/", encodedFolder) + encodedParam
		}
	}

	// Start building the URI with the base parameters
	uri := fmt.Sprintf("obsidian://%s?vault=%s&%s=%s",
		action, encodedVault, paramName, encodedParam)

	fmt.Println(uri)

	return openURI(uri)
}

func openURI(uri string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", uri)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", uri)
	default: // Linux and others
		cmd = exec.Command("xdg-open", uri)
	}

	return cmd.Run()
}
