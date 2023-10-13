package browsergo

import (
	"context"
	_ "embed"
	"log"
	"strconv"
	"time"

	"github.com/KakashiHatake324/browser-go/browser/exec"
)

// launch the server
func (c *ClientInit) launchServer() error {
	var err error
	exe, err := exec.New("browser-go", program)
	if err != nil {
		return err
	}
	cmd := exe.CommandContext(context.Background(), strconv.Itoa(c.port), strconv.FormatBool(c.verbose))
	c.closeExe = exe.Close
	if err := cmd.Start(); err != nil {
		return err
	}
	if c.verbose {
		log.Println("[LAUNCH SERVER] server launched successfully")
	}
	time.Sleep(5 * time.Second)
	return err
}
