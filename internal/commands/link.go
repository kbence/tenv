package commands

import (
	"github.com/kbence/tenv/internal/tenv"
	"github.com/spf13/cobra"
)

func NewLinkCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link",
		Short: "Creates Teleport commands as symbolic links pointing to the binary",
		RunE: func(cmd *cobra.Command, args []string) error {
			force, err := cmd.Flags().GetBool("force")
			if err != nil {
				return err
			}

			return tenv.CreateLinks(force)
		},
	}
	cmd.Flags().BoolP("force", "f", false, "Ignore existing binaries")

	return cmd
}
