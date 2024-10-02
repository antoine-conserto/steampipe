//go:build windows
// +build windows

package pluginmanager

import (
	"os/exec"

	"github.com/turbot/steampipe/pkg/constants"
	"github.com/turbot/steampipe/pkg/filepaths"
)

func getPluginManagerCmd(steampipeExecutablePath string) *exec.Cmd {
	// note: we assume the install dir has been assigned to file_paths.SteampipeDir
	// - this is done both by the FDW and Steampipe
	pluginManagerCmd := exec.Command(steampipeExecutablePath,
		"plugin-manager",
		"--"+constants.ArgInstallDir, filepaths.SteampipeDir)

	return pluginManagerCmd
}
