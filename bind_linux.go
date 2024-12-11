//go:build linux

package browsergo

import (
	_ "embed"
	"log"
	"runtime"
	"strings"
)

//go:embed browser/exec/browser-solution-linux-x64
var program []byte

//go:embed browser/exec/browser-solution-linux-arm64
var programArm []byte

var pathSeparator = "/"

// get the program depending on the arch
func getProgram() []byte {
	arch := runtime.GOARCH
	log.Println("[INITBrowser] Linux:", arch)
	if arch == "arm64" {
		return programArm
	} else {
		return program
	}
}

func fixPath(path string) string {
	return strings.ReplaceAll(path, "\\", pathSeparator)
}
