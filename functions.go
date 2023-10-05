package crigo

import (
	"encoding/base64"
	"math/rand"
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
