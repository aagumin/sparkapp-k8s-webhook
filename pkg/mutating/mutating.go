package mutating

import (
	"github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
	"log/slog"
)

func mutateSparkApplication(sparkApp *v1beta2.SparkApplication, cfg *SparkAppConfig) []patchOperation {
	// all mutating
	var result []patchOperation

	annotations := sparkApp.Annotations
	teamId := annotations["amazme.spark.app"]

	if cfg.FeatureList.Affinity.Enabled && cfg.FeatureList.Toleration.Enabled {
		result = append(result, addAffinity(sparkApp, cfg.SparkPatchValue.AmazmeSparkAffinity[teamId], cfg.FeatureList.Affinity.HardPatch)...)
		result = append(result, addToleration(sparkApp, cfg.SparkPatchValue.AmazmeSparkToleration[teamId], cfg.FeatureList.Affinity.HardPatch)...)
		slog.Info("Add node selector rules")
	}

	return result
}
