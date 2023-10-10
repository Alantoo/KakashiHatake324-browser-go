package browsergo

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
)

// Returns a random number between the minimum and the maximum paraneters
func RandomInt(min, max int) int {
	if max-min == 0 {
		return 0
	}
	return min + rand.Intn(max-min)
}

// Decode a base64 encoded string
func decodeBase64(encoded string) string {
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	return string(decoded)
}

// Format proxy from ip:port:user:pass
func FormatProxy(proxyString string) (string, error) {
	var proxy string
	var err error
	proxySplit := strings.Split(proxyString, ":")
	if len(proxySplit) == 4 {
		proxy = "http://" + proxySplit[2] + ":" + proxySplit[3] + "@" + proxySplit[0] + ":" + proxySplit[1]
	} else if len(proxySplit) == 2 {
		proxy = "http://" + proxySplit[0] + ":" + proxySplit[1]
	} else {
		err = errors.New("error formatting proxies")
	}
	return proxy, err
}
