// proutils-api/whois/handler.go
//
// API handler for /whois

package whois

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HandleWHOISIP returns WHOIS info for provided public IP.
func HandleWHOISIP(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	ip := params.ByName("ip")
	json := GetNetwork(ip)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
