package browsergo

import (
	"bufio"
	"context"
	_ "embed"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/KakashiHatake324/browser-go/browser/exec"
)

// launch the server
func (c *ClientInit) launchServer() error {
	var err error
	reader, writer := io.Pipe()
	scanner := bufio.NewScanner(reader)
	exe, err := exec.New("browser-go", program)
	if err != nil {
		return err
	}
	cmd := exe.CommandContext(context.Background(), strconv.Itoa(c.port), strconv.FormatBool(c.verbose))
	c.closeExe = exe.Close
	cmd.Stdout = writer
	if err := cmd.Start(); err != nil {
		return err
	}
	if c.verbose {
		log.Println("[LAUNCH SERVER] server launched successfully")
	}
	for scanner.Scan() {
		if c.verbose {
			fmt.Println(scanner.Text())
		}
		if strings.Contains(scanner.Text(), "ERROR") {
			err = errServerConnection
			break
		} else if strings.Contains(scanner.Text(), "server is running") {
			break
		}
	}
	return err
}
