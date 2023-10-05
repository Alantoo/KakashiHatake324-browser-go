//go:build windows

package browsergo

import _ "embed"

//go:embed browser/exec/browser-solution-win.exe
var program []byte
