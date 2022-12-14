package commands

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tenv",
		Short: "Switch between different versions of Teleport",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
	cmd.AddCommand(NewInstallCommand())
	cmd.AddCommand(NewUseCommand())
	cmd.AddCommand(NewLinkCommand())
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewSelectProfileCommand())

	return cmd
}
