/*
Copyright © 2023 Juan Wiggenhauser <jgwiggenhauser@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/JuanWigg/drainkube/cmd/check"
	"github.com/JuanWigg/drainkube/cmd/drain"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "drainkube",
	Short: "Drainkube is a tool for safely draining nodes",
	Long: `Use this tool for safely draining your Kubernetes nodes honoring
	the PodDisruptionBudgets made for your pods:

This applications scales the Deployment/HPA for your pod in order to evict it
succesfully making sure the PodDisruptionBudget is being honored.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(drain.DrainCmd)
	rootCmd.AddCommand(check.CheckCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommandPalettes()
}
