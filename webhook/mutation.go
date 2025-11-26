package webhook

import (
	"encoding/json"
    "fmt"
    "log/slog"
    "net/http"

	"sigs.k8s.io/controller-runtime/pkg/webhook"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func HandleMutation(w http.ResponseWriter, r *http.Request) {
	
}