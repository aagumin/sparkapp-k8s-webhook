package webhook

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
	v1 "k8s.io/api/core/v1"
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
	StandSparkAffinity            map[string]v1.Affinity     `yaml:"StandSparkAffinity"`
	StandSparkToleration          map[string][]v1.Toleration `yaml:"StandSparkToleration"`
	StandSparkHistoryServerSpark  map[string]string          `yaml:"StandSparkHistoryServerSpark"`
	StandSparkHistoryServerHadoop map[string]string          `yaml:"StandSparkHistoryServerHadoop"`
	StandSparkLabels              map[string]string          `yaml:"StandSparkLabels"`
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
