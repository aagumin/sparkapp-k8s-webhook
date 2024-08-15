package webhook

import (
	"github.com/kubeflow/spark-operator/api/v1beta2"
	"log/slog"
)

func mutateSparkApplication(sparkApp *v1beta2.SparkApplication, cfg *SparkAppConfig) []patchOperation {
	// all webhook
	var result []patchOperation

	annotations := sparkApp.Annotations
	teamId := annotations["amazme.spark.app"]

	if cfg.FeatureList.Affinity.Enabled && cfg.FeatureList.Toleration.Enabled {
		result = append(result, addAffinity(sparkApp, cfg.SparkPatchValue.AmazmeSparkAffinity[teamId], cfg.FeatureList.Affinity.HardPatch)...)
		result = append(result, addToleration(sparkApp, cfg.SparkPatchValue.AmazmeSparkToleration[teamId], cfg.FeatureList.Toleration.HardPatch)...)
		slog.Info("Add node selection rules")
	}

	return result
}
