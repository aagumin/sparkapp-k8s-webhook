package webhook

import (
	"encoding/json"
	"github.com/kubeflow/spark-operator/api/v1beta2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"testing"
)

func TestParseSparkApp(t *testing.T) {
	mainAppFile := "mainAppFile"
	sampleApp := &v1beta2.SparkApplication{
		TypeMeta: metav1.TypeMeta{APIVersion: "sparkoperator.k8s.io/v1beta2", Kind: "SparkApplication"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sample-app",
			Namespace: "default",
		},
		Spec: v1beta2.SparkApplicationSpec{
			MainApplicationFile: &mainAppFile,
		},
	}

	sampleAppBytes, err := json.Marshal(sampleApp)
	require.NoError(t, err)

	mockReview := &v1.AdmissionReview{
		Request: &v1.AdmissionRequest{
			Object: runtime.RawExtension{
				Raw: sampleAppBytes,
			},
		},
	}

	app, err := parseSparkApp(*mockReview)
	require.NoError(t, err)

	assert.Equal(t, sampleApp.Name, app.Name)
	assert.Equal(t, sampleApp.Namespace, app.Namespace)
	assert.Equal(t, sampleApp.Spec.MainApplicationFile, app.Spec.MainApplicationFile)
}

func TestParseSparkApp_InvalidJSON(t *testing.T) {
	invalidJSON := []byte(`{ "invalid - "data" }`)
	mockReview := &v1.AdmissionReview{
		Request: &v1.AdmissionRequest{
			Object: runtime.RawExtension{
				Raw: invalidJSON,
			},
		},
	}

	value, err := parseSparkApp(*mockReview)
	assert.Nil(t, value)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "failed to unmarshal a SparkApplication from the raw data in the admission request: invalid character 'd' after object key")
}

func TestParseSparkApp_InvalidSparkApp(t *testing.T) {
	invalidJSON := []byte(`{"invalid":"data"}`)
	mockReview := v1.AdmissionReview{
		Request: &v1.AdmissionRequest{
			Object: runtime.RawExtension{
				Raw: invalidJSON,
			},
		},
	}

	// Call the parseSparkApp function with the mock AdmissionReview object
	value, err := parseSparkApp(mockReview)
	assert.Nil(t, value)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Non SparkApplication in object")
}

func TestParseSparkApp_NilAdmissionReview(t *testing.T) {
	// Call the parseSparkApp function with a nil AdmissionReview object
	_, err := parseSparkApp(v1.AdmissionReview{})
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "AdmissionReview has no request")
}

func TestParseSparkApp_NilObject(t *testing.T) {
	// Create a mock AdmissionReview object with a nil Object field
	mockReview := &v1.AdmissionReview{
		Request: &v1.AdmissionRequest{
			Object: runtime.RawExtension{},
		},
	}

	// Call the parseSparkApp function with the mock AdmissionReview object
	sparkApp, err := parseSparkApp(*mockReview)
	assert.Error(t, err)
	assert.Nil(t, sparkApp)
	assert.Equal(t, err.Error(), "AdmissionReview has no object")
}
