// proutils-api/publicip/publicip.go
//
// Logic for getting client public IP.

package publicip

import (
	"net"
	"net/http"
	"strings"
)

// IPInfo is the JSON response record for the public-ip GET request.
type IPInfo struct {
	IP string `json:"IP"`
}

// GetIP returns a user's public facing IP address (IPv4 OR IPv6).
// It will return the IP address in plain text.
func GetIP(r *http.Request) string {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// Grabs the first IP address in the X-Forwarded-For header
	// list bbecause this is always the *origin* IP address, which
	// is the *true* IP of the user. For more information on this, see the
	// Wikipedia page: https://en.wikipedia.org/wiki/X-Forwarded-For

	ip := net.ParseIP(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]).String()
	return ip
}
