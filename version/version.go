package version

import (
	"fmt"
	"github.com/alimy/echo-music/api/v1"
	"github.com/alimy/echo-music/cmd"
	"github.com/spf13/cobra"
)

var (
	Version = "0.0.0"

	// GitHash Value will be set during build
	GitHash = "Not provided"

	// BuildTime Value will be set during build
	BuildTime = "Not provided"

	// version sub-command
	versionCmd = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s (APIVersion:%s)\nBuildTime:%s\nBuildGitSHA:%s\n",
				Version, api.ApiVersion, BuildTime, GitHash)
		},
	}
)

func init() {
	// Re
	cmd.Register(versionCmd)
}
