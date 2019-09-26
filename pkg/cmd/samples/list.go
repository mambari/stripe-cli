package samples

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"

	"github.com/stripe/stripe-cli/pkg/samples"
	"github.com/stripe/stripe-cli/pkg/validators"
)

// ListCmd prints a list of all the available sample projects that users can
// generate
type ListCmd struct {
	Cmd *cobra.Command
}

// NewListCmd creates and returns a list command for samples
func NewListCmd() *ListCmd {
	listCmd := &ListCmd{}
	listCmd.Cmd = &cobra.Command{
		Use:   "list",
		Args:  validators.NoArgs,
		Short: "list available Stripe samples",
		Long:  `A list of available Stripe Sample integrations`,
		Run:   listCmd.runListCmd,
	}

	return listCmd
}

func (lc *ListCmd) runListCmd(cmd *cobra.Command, args []string) {
	fmt.Println("A list of available Stripe Sample integrations:")
	fmt.Println()

	names := samples.Names()
	sort.Strings(names)

	for _, name := range names {
		fmt.Println(samples.List[name].BoldName())
		fmt.Println(samples.List[name].Description)
		fmt.Println(fmt.Sprintf("Repo: %s", samples.List[name].URL))
		fmt.Println()
	}
}
