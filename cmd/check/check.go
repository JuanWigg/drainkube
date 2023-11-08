/*
Copyright © 2023 Juan Wiggenhauser <jgwiggenhauser@gmail.com>
*/
package check

import (
	"github.com/spf13/cobra"
)

var (
	nodeName string
)

// CheckCmd represents the check command
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Use the check command to check different information before doing a drain",
	Long: `You can use this command for checking for example the pods running on a node, 
affected PodDisruptionBudgets, affected Deployments or affected HPAs`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	CheckCmd.AddCommand(podsCmd)
	CheckCmd.AddCommand(pdbsCmd)
	CheckCmd.AddCommand(deploymentsCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
