/*
Copyright © 2023 Juan Wiggenhauser <jgwiggenhauser@gmail.com>
*/
package check

import (
	"context"
	"fmt"

	"github.com/JuanWigg/drainkube/cmd/util"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func checkPods() {
	client := util.GetInstance()
	pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{FieldSelector: fmt.Sprintf("spec.nodeName=%s", nodeName)})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the in the node\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Println(pod.ObjectMeta.Name)
	}

}

func getPods() *v1.PodList {
	client := util.GetInstance()
	pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{FieldSelector: fmt.Sprintf("spec.nodeName=%s", nodeName)})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the in the node\n", len(pods.Items))
	return pods
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
