/*
Copyright © 2023 Juan Wiggenhauser <jgwiggenahuser@gmail.com>
*/
package drain

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	nodeName string
)

var DrainCmd = &cobra.Command{
	Use:   "drain",
	Short: "This command is used for draining a node",
	Long: `The command drain safely drains a Kubernetes node making sure it respects the
	PodDisruptionBudgets defined for every pod scheduled and running on the node`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DrainCmd.Flags().StringVarP(&nodeName, "node", "n", "", "The node to drain")

	if err := DrainCmd.MarkFlagRequired("node"); err != nil {
		fmt.Println(err)
	}
}
