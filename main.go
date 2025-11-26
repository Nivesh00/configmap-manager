package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nivesh00/configmap-admission-webhook.git/global"
	"github.com/Nivesh00/configmap-admission-webhook.git/webhook"
)

func main() {

    // Look for forbidden keys in environmental variables
    // and create a global variable from it
    global.AssignForbiddenKeys()

    http.HandleFunc("/validate-configmap-keys", webhook.HandleValidation)
    http.HandleFunc("/mutate-configmap-keys", webhook.HandleMutation)
    log.Fatal(http.ListenAndServe(":443", nil))
}


