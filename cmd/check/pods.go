/*
Copyright © 2023 Juan Wiggenhauser <jgwiggenhauser@gmail.com>
*/
package check

import (
	"fmt"

	"github.com/JuanWigg/drainkube/cmd/kubernetes"
	"github.com/spf13/cobra"
)

func checkPods() {
	kubernetes.GetPods(nodeName)
}

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Shows the pod running on the node",
	Long:  `Use this command for checking the pods running on the node.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkPods()
	},
}

func init() {
	podsCmd.Flags().StringVarP(&nodeName, "node", "n", "", "The node to check")
	if err := podsCmd.MarkFlagRequired("node"); err != nil {
		fmt.Println(err)
	}
}
