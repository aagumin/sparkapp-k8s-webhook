package webhook

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kubeflow/spark-operator/api/v1beta2"
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"log/slog"
	"net/http"
)

// WebHook represents a webhook server for webhook SparkApplication objects.
type WebHook struct {
	MutateConfig *SparkAppConfig
}

// serveHealth serves the "/health" endpoint, which returns a 204 No Content status code and the string "Ok".
// It is used to indicate that the webhook server is running and healthy.
//
// Parameters:
//   - w: The http.ResponseWriter to which the response will be written.
//   - r: The http.Request that triggered the webhook server to serve the "/health" endpoint.
//
// Returns:
//   - nil if the "/health" endpoint is successfully served, otherwise an error.
func (wh *WebHook) serveHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	_, err := fmt.Fprint(w, "Ok")
	if err != nil {
		slog.Info("Health success")
		return
	}
}

// mutateReview handles the "/mutate" endpoint, which mutates the provided SparkApplication object according to the MutateConfig.
// It returns a JSON-encoded AdmissionResponse containing the patched SparkApplication object.
//
// Parameters:
//   - w: The http.ResponseWriter to which the response will be written.
//   - r: The http.Request that triggered the webhook server to serve the "/mutate" endpoint.
//
// Returns:
//   - nil if the "/mutate" endpoint is successfully served, otherwise an error.
func (wh *WebHook) mutateReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("http method doesnt support, try `POST`")
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	var sparkApp *v1beta2.SparkApplication

	admReview, err := parseAdmRequest(*r)
	slog.Debug(fmt.Sprintf("Successfully parsed adm review request %s", admReview))
	if err != nil {
		errS := err.Error()
		slog.Error(errS)
		http.Error(w, errS, http.StatusBadRequest)
		return
	}

	sparkApp, err = parseSparkApp(*admReview)
	slog.Debug("Successfully parsed sparkApp from adm request")

	if err != nil {
		errS := err.Error()
		slog.Error(errS)
		http.Error(w, errS, http.StatusBadRequest)
		return
	}

	sparkAppPatches := mutateSparkApplication(sparkApp, wh.MutateConfig)
	slog.Debug(fmt.Sprintf("Successfully patch spark app - %s", sparkAppPatches))
	var patchStatus metav1.Status
	marshal, err := json.Marshal(sparkAppPatches)

	if err != nil {
		patchStatus = metav1.Status{Message: err.Error()}
	} else {
		patchStatus = metav1.Status{Message: "successfully patched spark app"}
	}

	patchType := v1.PatchTypeJSONPatch

	admResp := v1.AdmissionResponse{
		UID:       admReview.Request.UID,
		Allowed:   true,
		Result:    &patchStatus,
		Patch:     marshal,
		PatchType: &patchType,
	}

	admReview.Response = &admResp
	resp, err := json.Marshal(admReview)

	slog.Debug(fmt.Sprintf("Successfully marshaled admission response %s", string(resp)))

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

// RunWebhookServer starts a webhook server that listens on the specified port and serves the "/health" and "/mutate" endpoints.
// It uses the provided TLS certificates for secure communication.
// The server's error log is set to the provided logger.
//
// Parameters:
//   - certFile: The path to the TLS certificate file.
//   - keyFile: The path to the TLS key file.
//   - port: The port number on which the server should listen.
//   - logger: The logger to use for the server's error log.
//
// Returns:
//   - nil if the server starts successfully, otherwise an error.
func (wh *WebHook) RunWebhookServer(certFile, keyFile string, port uint, logger *log.Logger) {
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
		ErrorLog: logger,
	}

	if err := server.ListenAndServeTLS("", ""); err != nil {
		panic(err)
	}
}
