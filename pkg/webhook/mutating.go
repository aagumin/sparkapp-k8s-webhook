package webhook

import (
	"log/slog"

	"github.com/kubeflow/spark-operator/api/v1beta2"
)

func mutateSparkApplication(sparkApp *v1beta2.SparkApplication, cfg *SparkAppConfig) []patchOperation {
	// all webhook
	var result []patchOperation

	annotations := sparkApp.GetAnnotations()
	teamId := annotations["example.spark.app"]
	if teamId == "" || annotations == nil {
		slog.Debug("No teamId found in annotations, skipping affinity and toleration. Set `example.spark.app` annotations")
		return result
	}
	if cfg.FeatureList.Affinity.Enabled && cfg.FeatureList.Toleration.Enabled {
		result = append(result, addAffinity(sparkApp, cfg.SparkPatchValue.StandSparkAffinity[teamId], cfg.FeatureList.Affinity.HardPatch)...)
		result = append(result, addToleration(sparkApp, cfg.SparkPatchValue.StandSparkToleration[teamId], cfg.FeatureList.Toleration.HardPatch)...)
		slog.Info("Add node selection rules")
	}

	return result
}
