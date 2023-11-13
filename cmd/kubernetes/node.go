package kubernetes

import (
	"context"
	"encoding/json"

	"github.com/JuanWigg/drainkube/cmd/util"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type patchBoolValue struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value bool   `json:"value"`
}

type patchStringValue struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

func CordonNode(nodeName string) {
	nodesClient := util.GetInstance().CoreV1().Nodes()

	payload := []patchBoolValue{{
		Op:    "replace",
		Path:  "/spec/unschedulable",
		Value: true,
	}}

	payloadBytes, _ := json.Marshal(payload)

	_, err := nodesClient.Patch(context.TODO(), nodeName, types.JSONPatchType, payloadBytes, v1.PatchOptions{})
	if err != nil {
		panic(err.Error())
	}
}
