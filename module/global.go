package module

import (
	"os"
	"strings"
	"log/slog"
)

// Global variable for user settings
var GlobalForbiddenKeys ForbiddenKeys

func AssignForbiddenKeys() {

	// Get environmental variable
    keys          := os.Getenv("FORBIDDEN_KEYS")
    policy        := os.Getenv("POLICY")
    caseSensitive := true
    if os.Getenv("CASE_SENSITIVE") == "disabled" {
        caseSensitive = false
    }
    // Assign forbidden keys struct
    GlobalForbiddenKeys.CreateForbiddenKeyList(keys, policy, caseSensitive)

	slog.Info(
		"Finished processing forbidden keys.",
		"Forbidden Keys are: ",
		strings.Join(GlobalForbiddenKeys.KeyList, " "),
	)
}


// Object to store for testing
type ForbiddenKeys struct {
    KeyList        []string
    Policy           string
    CaseSensitive    bool
}


// Looks for environmental variables and adds them to
// the key list
func (f *ForbiddenKeys) CreateForbiddenKeyList(keys string, policy string, caseSensitive bool) {

    // Assign values
    f.CaseSensitive = caseSensitive

	// If case sensitive, store all values as lowercase
	if !caseSensitive {
		keys = strings.ToLower(keys)
	}
    f.KeyList       = strings.Split(keys, ",")

    f.Policy        = policy
}