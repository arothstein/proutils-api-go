// proutils-api/pwgen/handler.go
//
// API handler for /password-generator

package pwgen

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HandlePwGen returns a list of generated passwords based on input criteria..
func HandlePwGen(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	wordCount := params.ByName("wordCount")
	pw := Generate(wordCount)
	pwList := Passwords{PW: pw}
	json.NewEncoder(w).Encode(pwList)
}
