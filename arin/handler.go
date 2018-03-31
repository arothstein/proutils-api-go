// proutils-api/arin/handler.go
//
// API handler for /arin

package arin

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HandleARIN returns ARIN info for provided public IP.
func HandleARIN(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	ip := params.ByName("ip")
	json := GetNetwork(ip)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
