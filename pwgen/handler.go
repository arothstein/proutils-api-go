// proutils-api/pwgen/handler.go
//
// API handler for /password-generator

package pwgen

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Passwords is the JSON response record for the password-generator POST request.
type passwordList struct {
	Passwords []string `json:"Passwords"`
}

// HandlePwGen returns a list of generated passwords based on input criteria..
func HandlePwGen(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// variable to hold the generated passwords
	pwSet := make([]string, 1)
	wordCount, err := strconv.Atoi(params.ByName("wordCount"))
	handleError(err)
	addInt, err := strconv.ParseBool(params.ByName("addInt"))
	handleError(err)
	addSpecial, err := strconv.ParseBool(params.ByName("addSpecial"))
	handleError(err)
	pwCount, err := strconv.Atoi(params.ByName("pwCount"))

	for i := 0; i < pwCount; i++ {
		pw, err := NewPassword(wordCount, addInt, addSpecial)
		handleError(err)
		pwSet = append(pwSet, pw)
	}

	passwords := passwordList{Passwords: pwSet}
	json.NewEncoder(w).Encode(passwords)
}
