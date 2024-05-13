package mutating

import (
	"gopkg.in/yaml.v3"
	v1 "k8s.io/api/core/v1"
	"log"
	"os"
)

type FeatureMode struct {
	Enabled   bool `yaml:"Enabled"`
	HardPatch bool `yaml:"softPatch"`
}

type FeatureList struct {
	Toleration           FeatureMode `yaml:"Affinity,omitempty"`
	Affinity             FeatureMode `yaml:"Toleration,omitempty"`
	HistoryServer        FeatureMode `yaml:"HistoryServer,omitempty"`
	PrometheusMonitoring FeatureMode `yaml:"PrometheusMonitoring,omitempty"`
}

type SparkPatchValue struct {
	// "ml":Affinity
	AmazmeSparkToleration          map[string][]v1.Toleration
	AmazmeSparkAffinity            map[string]v1.Affinity
	AmazmeSparkHistoryServerSpark  map[string]string
	AmazmeSparkHistoryServerHadoop map[string]string
	AmazmeSparkLabels              map[string]string
}

type SparkAppConfig struct {
	featureFlag FeatureList     `yaml:"FeatureList"`
	patchValues SparkPatchValue `yaml:"SparkPatchValue"`
}

func (fl *SparkAppConfig) GetConf(path string) *SparkAppConfig {
	in, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(in, &fl)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return fl
}
