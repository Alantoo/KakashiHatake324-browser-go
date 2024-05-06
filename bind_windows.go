//go:build windows

package browsergo

import (
	_ "embed"
	"strings"
)

//go:embed browser/exec/browser-solution-win-x64.exe
var program []byte

var pathSeparator = "\\"

func fixPath(path string) string {
	return strings.ReplaceAll(path, "/", pathSeparator)
}
