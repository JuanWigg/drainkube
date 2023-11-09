package drain

import (
	"fmt"

	"github.com/JuanWigg/drainkube/cmd/kubernetes"
	"github.com/JuanWigg/drainkube/cmd/util"
	v1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
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
				affectedPods = append(affectedPods, Pod{pdb.ObjectMeta.Name, pdb.ObjectMeta.Labels})
			}
		}
	}

}
