package commands

import (
	"github.com/kbence/tenv/internal/tenv"
	"github.com/spf13/cobra"
)

func getInstallValidArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	versions, _ := tenv.GetInstalledVersions()
	return tenv.StringSliceFilter(
		versions.StringSlice(),
		tenv.FilterStringStartsWith(toComplete),
	), cobra.ShellCompDirectiveDefault
}

func NewUseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "use",
		Short:             "Select the version of Teleport to be used",
		Aliases:           []string{"u"},
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: getInstallValidArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return tenv.UseTeleport(args[0])
		},
	}

	return cmd
}
