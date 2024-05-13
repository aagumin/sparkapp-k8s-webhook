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

	if cfg.featureFlag.Affinity.Enabled && cfg.featureFlag.Toleration.Enabled {
		result = append(result, addAffinity(sparkApp, cfg.patchValues.AmazmeSparkAffinity[teamId], cfg.featureFlag.Affinity.HardPatch)...)
		result = append(result, addToleration(sparkApp, cfg.patchValues.AmazmeSparkToleration[teamId], cfg.featureFlag.Affinity.HardPatch)...)
		slog.Info("Add node selector rules")
	}

	return result
}
