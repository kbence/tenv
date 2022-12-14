package commands

import (
	"fmt"
	"os"
	"regexp"

	"github.com/kbence/tenv/internal/tenv"
	"github.com/spf13/cobra"
)

var DomainNameRegExp = regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9-]*[A-Za-z0-9])$`)

func isValidDomainName(domainName string) bool {
	return DomainNameRegExp.MatchString(domainName)
}

func NewSelectProfileCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "select-profile",
		Short:   "Sets the current profile for tsh/tctl",
		Aliases: []string{"select"},
		Args:    cobra.ExactArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			profileNames, err := tenv.GetValidProfileNames()
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot get valid profile names: %s", err)
				return nil, cobra.ShellCompDirectiveError
			}

			return profileNames, cobra.ShellCompDirectiveDefault
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			selectedProfile := args[0]

			if !isValidDomainName(selectedProfile) {
				return fmt.Errorf("'%s' is not a valid profile name", selectedProfile)
			}

			profileNames, err := tenv.GetValidProfileNames()
			if err != nil {
				return err
			}

			for _, profileName := range profileNames {
				if profileName == selectedProfile {
					return tenv.SelectProfile(selectedProfile)
				}
			}

			return fmt.Errorf("profile name")
		},
	}
}
