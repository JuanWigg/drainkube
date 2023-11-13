/*
Copyright © 2023 Juan Wiggenhauser <jgwiggenahuser@gmail.com>
*/
package drain

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	nodeName         string
	affectedPdbs     []PDB
	affectedPods     []Pod
	affectedDeploys  []Deploy
	affectedRollouts []Rollout
	affectedHPAs     []HPA
)

func drainNode() {
	// Get affected resources
	fmt.Println("Getting affected PDBs...")
	getAffectedPdbs()
	fmt.Println("The affected PDBs are:")
	fmt.Println(affectedPdbs)
	fmt.Println("Getting affected Rollouts...")
	getAffectedRollouts()
	fmt.Println("The affected Rollouts are:")
	fmt.Println(affectedRollouts)
	fmt.Println("Getting affected Deploys...")
	getAffectedDeploys()
	fmt.Println("The affected Deploys are:")
	fmt.Println(affectedDeploys)
	fmt.Println("Getting affected HPAs...")
	getAffectedHpas()
	fmt.Println("The affected HPAs are:")
	fmt.Println(affectedHPAs)

	// Start the process
}

var DrainCmd = &cobra.Command{
	Use:   "drain",
	Short: "This command is used for draining a node",
	Long: `The command drain safely drains a Kubernetes node making sure it respects the
	PodDisruptionBudgets defined for every pod scheduled and running on the node`,
	Run: func(cmd *cobra.Command, args []string) {
		drainNode()
	},
}

func init() {
	DrainCmd.Flags().StringVarP(&nodeName, "node", "n", "", "The node to drain")

	if err := DrainCmd.MarkFlagRequired("node"); err != nil {
		fmt.Println(err)
	}
}
