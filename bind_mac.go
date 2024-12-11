//go:build darwin

package browsergo

import (
	_ "embed"
	"strings"
)

//go:embed browser/exec/browser-solution-macos-arm64
var program []byte

var pathSeparator = "/"

// get the program depending on the arch
func getProgram() []byte {
	return program
}

func fixPath(path string) string {
	return strings.ReplaceAll(path, "\\", pathSeparator)
}
