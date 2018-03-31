// proutils-api/arin/arin.go
//
// Logic for getting ARIN information.

package arin

import (
	"io/ioutil"
	"net/http"

	xj "github.com/basgys/goxml2json"
)

// GetNetwork takes an IP input and queries ARIN API to return Network info
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
