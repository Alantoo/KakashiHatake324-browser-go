//go:build windows

package browsergo

import (
	_ "embed"
	"log"
	"runtime"
	"strings"
)

//go:embed browser/exec/browser-solution-win-x64.exe
var program []byte

var pathSeparator = "\\"

// get the program depending on the arch
func getProgram() []byte {
	arch := runtime.GOARCH
	log.Println("[INITBrowser] Windows:", arch)
	return program
}

func fixPath(path string) string {
	return strings.ReplaceAll(path, "/", pathSeparator)
}
