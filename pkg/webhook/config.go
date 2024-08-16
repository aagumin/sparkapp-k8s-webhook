package webhook

import (
	"gopkg.in/yaml.v3"
	v1 "k8s.io/api/core/v1"
	"log"
	"os"
)

type FeatureList struct {
	Toleration           *FeatureMode `yaml:"Toleration,omitempty"`
	Affinity             *FeatureMode `yaml:"Affinity,omitempty"`
	HistoryServer        *FeatureMode `yaml:"HistoryServer,omitempty"`
	PrometheusMonitoring *FeatureMode `yaml:"PrometheusMonitoring,omitempty"`
}

type FeatureMode struct {
	Enabled   bool `yaml:"enabled"`
	HardPatch bool `yaml:"hardPatch"`
}

type SparkPatchValue struct {
	// "ml":Affinity
	AmazmeSparkAffinity            map[string]v1.Affinity     `yaml:"AmazmeSparkAffinity"`
	AmazmeSparkToleration          map[string][]v1.Toleration `yaml:"AmazmeSparkToleration"`
	AmazmeSparkHistoryServerSpark  map[string]string          `yaml:"AmazmeSparkHistoryServerSpark"`
	AmazmeSparkHistoryServerHadoop map[string]string          `yaml:"AmazmeSparkHistoryServerHadoop"`
	AmazmeSparkLabels              map[string]string          `yaml:"AmazmeSparkLabels"`
}

type SparkAppConfig struct {
	FeatureList     FeatureList     `yaml:"FeatureList"`
	SparkPatchValue SparkPatchValue `yaml:"SparkPatchValue"`
}

func GetConf(path string) (SparkAppConfig, error) {
	var fl *SparkAppConfig
	in, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %v", err)
		return *fl, err
	}

	err = yaml.Unmarshal(in, &fl)
	if err != nil {
		log.Fatalf("error: %v", err)
		return *fl, err
	}
	return *fl, nil
}
