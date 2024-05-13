package mutating

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
	v1 "k8s.io/api/admission/v1"
	"net/http"
)

func parseAdmRequest(r http.Request) (*v1.AdmissionReview, error) {
	if r.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("Content-Type: %q should be %q",
			r.Header.Get("Content-Type"), "application/json")
	}

	bodybuf := new(bytes.Buffer)
	_, err := bodybuf.ReadFrom(r.Body)
	if err != nil {
		// TODO Error
		return nil, err
	}
	body := bodybuf.Bytes()

	if len(body) == 0 {
		return nil, fmt.Errorf("admission request body is empty")
	}

	var a v1.AdmissionReview

	if err := json.Unmarshal(body, &a); err != nil {
		return nil, fmt.Errorf("could not parse admission review request: %v", err)
	}

	if a.Request == nil {
		return nil, fmt.Errorf("admission review can't be used: Request field is nil")
	}

	return &a, nil
}

func parseSparkApp(review v1.AdmissionReview) (*v1beta2.SparkApplication, error) {
	var app v1beta2.SparkApplication

	if err := json.Unmarshal(review.Request.Object.Raw, &app); err != nil {
		return nil, fmt.Errorf("failed to unmarshal a SparkApplication from the raw data in the admission request: %v", err)

	}
	return &app, nil
}
