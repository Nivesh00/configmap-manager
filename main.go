package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nivesh00/configmap-admission-webhook.git/module"
)

func main() {

    // Look for forbidden keys in environmental variables
    // and create a global variable from it
    module.AssignForbiddenKeys()

    http.HandleFunc("/validate-configmap-keys", module.HandleValidation)
    http.HandleFunc("/mutate-configmap-keys", module.HandleMutation)
    log.Fatal(http.ListenAndServe(":443", nil))
}


