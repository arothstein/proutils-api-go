// proutils-api/pwgen/handler.go
//
// API handler for /password-generator

package pwgen

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PwRequest is JSON request format received via POST
type PwRequest struct {
	WordCount  int  `json:"WordCount"`
	AddInt     bool `json:"AddInt"`
	AddSpecial bool `json:"AddSpecial"`
	PwCount    int  `json:"PwCount"`
}

// Passwords is the JSON response record for the password-generator POST request.
type passwordList struct {
	Passwords []string `json:"Passwords"`
}

// HandlePwGen returns a list of generated passwords based on input criteria.
func HandlePwGen(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// Unmarshal
	var req PwRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// variable to hold the generated passwords
	pwSet := make([]string, 1)

	for i := 0; i < req.PwCount; i++ {
		pw, err := NewPassword(req.WordCount, req.AddInt, req.AddSpecial)
		handleError(err)
		pwSet = append(pwSet, pw)
	}

	passwords := passwordList{Passwords: pwSet}
	json.NewEncoder(w).Encode(passwords)
}
