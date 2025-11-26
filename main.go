package main

import (
	"log"
	"net/http"
	"os"
    "log/slog"

	"github.com/Nivesh00/configmap-admission-webhook.git/module"
)

func main() {

    slog.Info("Starting up program")

    // Look for forbidden keys in environmental variables
    // and create a global variable from it
    err := module.AssignForbiddenKeys()
    if err != nil {
        slog.Error("A fatal problem occured, cannot continue", slog.Any("error", err))
        os.Exit(1)
    }

    slog.Info("Starting server")

    http.HandleFunc("/validate-configmap-keys", module.HandleValidation)
    http.HandleFunc("/mutate-configmap-keys", module.HandleMutation)

    log.Fatal(http.ListenAndServe(":443", nil))
}


