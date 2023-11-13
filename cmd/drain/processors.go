package drain

import (
	"fmt"

	"github.com/JuanWigg/drainkube/cmd/kubernetes"
	"github.com/JuanWigg/drainkube/cmd/util"
	v1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	allPdbs *policyv1.PodDisruptionBudgetList
	allPods *v1.PodList
)

func getAffectedPdbs() {
	allPdbs = kubernetes.GetPdbs()
	allPods = kubernetes.GetPods(nodeName)

	for _, pod := range allPods.Items {
		for _, pdb := range allPdbs.Items {
			if util.IsMapSubset(pod.ObjectMeta.Labels, pdb.Spec.Selector.MatchLabels) {
				fmt.Printf("Pod %s is referenced by PDB %s.\n", pod.ObjectMeta.Name, pdb.ObjectMeta.Name)
				affectedPdbs = append(affectedPdbs, PDB{pdb.ObjectMeta.Name, pdb.Spec.Selector.MatchLabels})
				affectedPods = append(affectedPods, Pod{pdb.ObjectMeta.Name, pod.ObjectMeta.Labels})
			}
		}
	}
}

func getAffectedDeploys() {
	allDeploys := kubernetes.GetDeploys()

	for _, deploy := range allDeploys.Items {
		for _, pod := range affectedPods {
			if util.IsMapSubset(deploy.Spec.Selector.MatchLabels, pod.labels) {
				affectedDeploys = append(affectedDeploys, Deploy{deploy.ObjectMeta.Name, deploy.Spec.Selector.MatchLabels})
			}
		}
	}
}

func getAffectedRollouts() {
	allRollouts := kubernetes.GetRollouts()
	allPods = kubernetes.GetPods(nodeName)

	for _, rollout := range allRollouts.Items {
		for _, pod := range affectedPods {
			rolloutSelector, found, err := unstructured.NestedStringMap(rollout.Object, "spec", "selector", "matchLabels")
			if err != nil || !found {
				fmt.Printf("Selector not found for rollout %s: error=%s", rollout.GetName(), err)
				continue
			}
			if util.IsMapSubset(pod.labels, rolloutSelector) {
				affectedRollouts = append(affectedRollouts, Rollout{rollout.GetName(), rolloutSelector})
			}
		}
	}
}

func getAffectedHpas() {
	allHPAs := kubernetes.GetHPAs()

	for _, hpa := range allHPAs.Items {
		for _, deploy := range affectedDeploys {
			if deploy.name == hpa.Spec.ScaleTargetRef.Name {
				affectedHPAs = append(affectedHPAs, HPA{hpa.ObjectMeta.Name, hpa.Spec.ScaleTargetRef.Name, hpa.Spec.ScaleTargetRef.Kind})
			}
		}
		for _, rollout := range affectedRollouts {
			if hpa.Spec.ScaleTargetRef.Name == rollout.name {
				affectedHPAs = append(affectedHPAs, HPA{hpa.ObjectMeta.Name, hpa.Spec.ScaleTargetRef.Name, hpa.Spec.ScaleTargetRef.Kind})
			}
		}
	}
}
