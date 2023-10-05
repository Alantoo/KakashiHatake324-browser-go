//go:build !windows

package crigo

import _ "embed"

//go:embed browser/exec/browser-solution-macos
var program []byte
