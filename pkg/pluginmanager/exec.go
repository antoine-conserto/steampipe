//go:build darwin || linux
// +build darwin linux

package pluginmanager

import (
	"os/exec"
	"syscall"

	"github.com/turbot/steampipe/pkg/constants"
	"github.com/turbot/steampipe/pkg/filepaths"
)

func getPluginManagerCmd(steampipeExecutablePath string) *exec.Cmd {
	// note: we assume the install dir has been assigned to file_paths.SteampipeDir
	// - this is done both by the FDW and Steampipe
	pluginManagerCmd := exec.Command(steampipeExecutablePath,
		"plugin-manager",
		"--"+constants.ArgInstallDir, filepaths.SteampipeDir)

	// set attributes on the command to ensure the process is not shutdown when its parent terminates
	pluginManagerCmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	return pluginManagerCmd
}
