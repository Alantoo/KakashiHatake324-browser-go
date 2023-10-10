//go:build windows

package browsergo

import _ "embed"

//go:embed browser/exec/browser-solution-win-x64.exe
var program []byte
