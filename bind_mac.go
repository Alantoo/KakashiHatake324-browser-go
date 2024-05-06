//go:build !windows

package browsergo

import (
	_ "embed"
	"strings"
)

//go:embed browser/exec/browser-solution-macos-arm64
var program []byte

var pathSeparator = "/"

func fixPath(path string) string {
	return strings.ReplaceAll(path, "\\", pathSeparator)
}
