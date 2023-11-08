/*
Copyright © 2023 Juan Wiggenhauser <jgwiggenhauser@gmail.com>
*/
package check

import (
	"context"
	"fmt"

	"github.com/JuanWigg/drainkube/cmd/util"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PDB struct {
	name      string
	selectors map[string]string
}

var (
	podDisruptionBudgets []PDB
)

func checkPdbs() {
	fmt.Println("Checking PodDisruptionBudgets...")
	client := util.GetInstance()
	pdbs, err := client.PolicyV1().PodDisruptionBudgets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d PDBs in the Cluster\n", len(pdbs.Items))
	for _, pdb := range pdbs.Items {
		fmt.Printf("Getting PDB %s Selector\n", pdb.ObjectMeta.Name)
		podDisruptionBudgets = append(podDisruptionBudgets, PDB{pdb.ObjectMeta.Name, pdb.Spec.Selector.MatchLabels})
	}

	for _, pdb := range podDisruptionBudgets {
		fmt.Printf("PDB: %s, Selectors: %s\n", pdb.name, pdb.selectors)
	}

	fmt.Println("Getting pods in the node...")
	podList := getPods()

	fmt.Println("Analyizing pods")
	for _, pod := range podList.Items {
		for _, pdb := range podDisruptionBudgets {
			if util.IsMapSubset(pod.ObjectMeta.Labels, pdb.selectors) {
				fmt.Printf("Pod %s is referenced by PDB %s.\n", pod.ObjectMeta.Name, pdb.name)
			}
		}
	}
}

// pdbsCmd represents the pdbs command
var pdbsCmd = &cobra.Command{
	Use:   "pdbs",
	Short: "Shows the PDBs affected by the node",
	Long: `This commands shows the PodDisruptionBudgets that are affected by the node
	because the node has Pods related to the PodDisruptionBuget running.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkPdbs()
	},
}

func init() {
	pdbsCmd.Flags().StringVarP(&nodeName, "node", "n", "", "The node to check")
	if err := pdbsCmd.MarkFlagRequired("node"); err != nil {
		fmt.Println(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pdbsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pdbsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
