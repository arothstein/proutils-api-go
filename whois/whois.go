// proutils-api/whois/whois.go
//
// Logic for getting WHOIS information.

package whois

import (
	"io/ioutil"
	"net/http"

	xj "github.com/basgys/goxml2json"
)

// GetNetwork takes an IP input and queries ARIN API to return WHOIS Network info
func GetNetwork(ip string) []byte {
	resp, err := http.Get("http://whois.arin.net/rest/ip/" + ip)
	if err != nil {
		panic(err)
	}

	// convert xml body to json
	json, err := xj.Convert(resp.Body)
	if err != nil {
		panic(err)
	}

	// convert buffer to []byte
	jsonByte, _ := ioutil.ReadAll(json)
	return jsonByte
}
