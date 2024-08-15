package webhook

import (
	"github.com/kubeflow/spark-operator/api/v1beta2"
	v1 "k8s.io/api/core/v1"
)

type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

func addAffinity(sparkApp *v1beta2.SparkApplication, patchValue v1.Affinity, hardPatch bool) []patchOperation {
	var patchOps []patchOperation

	if sparkApp.Spec.Driver.Affinity != nil && sparkApp.Spec.Executor.Affinity != nil {
		if hardPatch {

			patchOps := append(patchOps, patchOperation{Op: "replace", Path: "/spec/driver/affinity", Value: patchValue})
			patchOps = append(patchOps, patchOperation{Op: "replace", Path: "/spec/executor/affinity", Value: patchValue})

			return patchOps
		} else {
			return patchOps
		}
	}

	patchOps = append(patchOps, patchOperation{Op: "add", Path: "/spec/driver/affinity", Value: patchValue})
	patchOps = append(patchOps, patchOperation{Op: "add", Path: "/spec/executor/affinity", Value: patchValue})

	return patchOps

}

func addToleration(sparkApp *v1beta2.SparkApplication, patchValue []v1.Toleration, hardPatch bool) []patchOperation {
	var patchOps []patchOperation

	if sparkApp.Spec.Driver.Tolerations != nil && sparkApp.Spec.Executor.Tolerations != nil {
		if hardPatch {

			patchOps := append(patchOps, patchOperation{Op: "replace", Path: "/spec/driver/toleration", Value: patchValue})
			patchOps = append(patchOps, patchOperation{Op: "replace", Path: "/spec/executor/toleration", Value: patchValue})

			return patchOps
		} else {
			return patchOps
		}
	}

	patchOps = append(patchOps, patchOperation{Op: "add", Path: "/spec/driver/toleration", Value: patchValue})
	patchOps = append(patchOps, patchOperation{Op: "add", Path: "/spec/executor/toleration", Value: patchValue})

	return patchOps

}
