package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "v1.0.0"

func ConfigCMD(superCMD *cobra.Command) {

	cmd := &cobra.Command{
		Use:     "version",
		Short:   "show version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Version)
		},
	}

	superCMD.AddCommand(cmd)
}