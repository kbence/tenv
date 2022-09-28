package commands

import (
	"github.com/kbence/tenv/internal/tenv"
	"github.com/spf13/cobra"
)

func NewUseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "use",
		Short:   "Select the version of Teleport to be used",
		Aliases: []string{"u"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return tenv.UseTeleport(args[0])
		},
	}

	return cmd
}
