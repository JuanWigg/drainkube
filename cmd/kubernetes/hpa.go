package kubernetes

import (
	"context"

	"github.com/JuanWigg/drainkube/cmd/util"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func DisableHPADownscaling(name string, namespace string) {
	hpaClient := util.GetInstance().AutoscalingV2().HorizontalPodAutoscalers(namespace)

	payload := `{"spec": {"behavior": {"scaleDown": {"selectPolicy": "Disabled"}}}}`

	_, err := hpaClient.Patch(context.TODO(), name, types.MergePatchType, []byte(payload), v1.PatchOptions{})
	if err != nil {
		panic(err.Error())
	}
}

func DefaultHPADownscaling(name string, namespace string) {
	hpaClient := util.GetInstance().AutoscalingV2().HorizontalPodAutoscalers(namespace)

	payload := `{"spec": {"behavior": {"scaleDown": {"selectPolicy": "Max"}}}}`

	_, err := hpaClient.Patch(context.TODO(), name, types.MergePatchType, []byte(payload), v1.PatchOptions{})
	if err != nil {
		panic(err.Error())
	}
}
