package webhook

import (
	"fmt"
	"github.com/kubeflow/spark-operator/api/v1beta2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mutateSparkApplication(t *testing.T) {
	configFixture := SparkAppConfig{
		FeatureList: FeatureList{
			Toleration: &FeatureMode{
				Enabled:   true,
				HardPatch: true,
			},
			Affinity: &FeatureMode{
				Enabled:   true,
				HardPatch: true,
			},
			HistoryServer:        nil,
			PrometheusMonitoring: nil,
		},
		SparkPatchValue: SparkPatchValue{},
	}
	testSparkApp := v1beta2.SparkApplication{}
	result := mutateSparkApplication(&testSparkApp, &configFixture)
	fmt.Printf("%T\n", result[0])
	assert.NotEqual(t, 0, len(result))
	assert.Equal(t, "add", result[0].Op)
	assert.Equal(t, "/spec/driver/affinity", result[0].Path)
	assert.Equal(t, "add", result[1].Op)
	assert.Equal(t, "/spec/executor/affinity", result[1].Path)
}
