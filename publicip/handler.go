// proutils-api/publicip/handler.go
//
// API handler for /public-ip

package publicip

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HandlePublicIP returns the requesting client's public IP based on the
// X-Forward-For header value.
func HandlePublicIP(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ip := GetIP(r)

	ipinfo := IPInfo{IP: ip}
	json.NewEncoder(w).Encode(ipinfo)
}
