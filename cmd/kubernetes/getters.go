package kubernetes

import (
	"context"
	"fmt"

	"github.com/JuanWigg/drainkube/cmd/util"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPods(nodeName string) *v1.PodList {
	client := util.GetInstance()
	pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{FieldSelector: fmt.Sprintf("spec.nodeName=%s", nodeName)})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the in the node\n", len(pods.Items))
	return pods
}

func GetPdbs() *policyv1.PodDisruptionBudgetList {
	client := util.GetInstance()
	pdbs, err := client.PolicyV1().PodDisruptionBudgets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d PDBs in the in the Cluster\n", len(pdbs.Items))
	return pdbs
}

func GetDeploys() *appsv1.DeploymentList {
	client := util.GetInstance()
	deploys, err := client.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d Deployments in the in the Cluster\n", len(deploys.Items))
	return deploys
}
