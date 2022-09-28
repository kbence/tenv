package commands

import (
	"github.com/kbence/tenv/internal/tenv"
	"github.com/spf13/cobra"
)

func NewInstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "install",
		Short:   "Install a version of Teleport",
		Aliases: []string{"i"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return tenv.InstallVersion(cmd.Context(), args[0])
		},
	}
}
