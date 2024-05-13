package mutating

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log/slog"
	"net/http"
)

type WebHook struct {
	mutateConfig *SparkAppConfig
}

func (wh *WebHook) serveHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	_, err := fmt.Fprint(w, "Ok")
	if err != nil {
		return
	}
}

func (wh *WebHook) mutateReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("http method doesnt support, try `POST`")
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	var sparkApp *v1beta2.SparkApplication

	admReview, err := parseAdmRequest(*r)
	if err != nil {
		errS := err.Error()
		slog.Error(errS)
		http.Error(w, errS, http.StatusBadRequest)
		return
	}

	sparkApp, err = parseSparkApp(*admReview)
	if err != nil {
		errS := err.Error()
		slog.Error(errS)
		http.Error(w, errS, http.StatusBadRequest)
		return
	}

	sparkAppPatches := mutateSparkApplication(sparkApp, wh.mutateConfig)

	marshal, err := json.Marshal(sparkAppPatches)
	if err != nil {
		return
	}
	status := metav1.Status{Message: "Success Patch"}
	patchType := v1.PatchTypeJSONPatch

	admResp := v1.AdmissionResponse{
		UID:              admReview.Request.UID,
		Allowed:          true,
		Result:           &status,
		Patch:            marshal,
		PatchType:        &patchType,
		AuditAnnotations: nil,
		Warnings:         nil,
	}

	resp, err := json.Marshal(admResp)
	if err != nil {
		msg := fmt.Sprintf("error marshalling response json: %v", err)
		slog.Error(msg)
		w.WriteHeader(500)
		_, err := w.Write([]byte(msg))
		if err != nil {
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		return
	}
}

func (wh *WebHook) RunWebhookServer(certFile, keyFile string, port uint) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(err)
	}
	slog.Info("Starting webhook server")
	http.HandleFunc("/health", wh.serveHealth)
	http.HandleFunc("/mutate", wh.mutateReview)
	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	if err := server.ListenAndServeTLS("", ""); err != nil {
		panic(err)
	}
}
