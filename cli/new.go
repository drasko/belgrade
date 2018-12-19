package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewNewCmd returns new command.
func NewNewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new <project_name>",
		Short: "Creates new project",
		Long:  `Creates new Belgrede project: generates boilerplate code structure with correct imports`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 1 {
				fmt.Println("Wrong number of params")
				return
			}

			fmt.Printf("Create new project %s\n", args[0])
		},
	}
}
