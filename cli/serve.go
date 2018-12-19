package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewServeCmd returns serve command.
func NewServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Serves current project",
		Long:  `Serves current Belgrede project: runs a web server and serves produced files`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 0 {
				fmt.Println("Wrong number of params")
				return
			}

			fmt.Println("Serving")
		},
	}
}
