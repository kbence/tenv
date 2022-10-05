package commands

import (
	"github.com/kbence/tenv/internal/tenv"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List installed teleport versions",
		Aliases: []string{"ls", "l"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return tenv.ListInstalledVersions()
		},
	}

	return cmd
}
