/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package check

import (
	"fmt"

	"github.com/JuanWigg/drainkube/cmd/kubernetes"
	"github.com/JuanWigg/drainkube/cmd/util"
	"github.com/spf13/cobra"
	appsv1 "k8s.io/api/apps/v1"
)

type Deploy struct {
	name      string
	selectors map[string]string
}

var (
	deployments     *appsv1.DeploymentList
	deploymentsList []Deploy
)

func checkDeploys() {
	fmt.Println("Checking Deployments...")
	deployments = kubernetes.GetDeploys()
	for _, deploy := range deployments.Items {
		deploymentsList = append(deploymentsList, Deploy{deploy.ObjectMeta.Name, deploy.Spec.Selector.MatchLabels})
	}

	for _, deploy := range deploymentsList {
		fmt.Println("Deployment: ", deploy.name, " Selectors: ", deploy.selectors)
	}

	fmt.Println("Getting pods in the node...")
	podList := kubernetes.GetPods(nodeName)

	fmt.Println("Analyizing pods")
	for _, pod := range podList.Items {
		for _, deploy := range deploymentsList {
			if util.IsMapSubset(pod.ObjectMeta.Labels, deploy.selectors) {
				fmt.Printf("Pod %s is referenced by Deploy %s.\n", pod.ObjectMeta.Name, deploy.name)
			}
		}
	}
}

// deploymentsCmd represents the deployments command
var deploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "Shows the Deployments contained by the node",
	Long:  `Shows what deployments have pods running on the node to check`,
	Run: func(cmd *cobra.Command, args []string) {
		checkDeploys()
	},
}

func init() {
	deploymentsCmd.Flags().StringVarP(&nodeName, "node", "n", "", "The node to check")
	if err := deploymentsCmd.MarkFlagRequired("node"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deploymentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deploymentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
