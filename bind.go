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
	"time"

	"github.com/KakashiHatake324/browser-go/browser/exec"
)

// launch the server
func (c *ClientInit) launchServer() error {
	var established bool
	var err error
	reader, writer := io.Pipe()

	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			if c.verbose {
				fmt.Println(scanner.Text())
			}
			if strings.Contains(scanner.Text(), "ERROR") {
				err = errServerConnection
				established = true
			} else if strings.Contains(scanner.Text(), "server is running") {
				established = true
			}
		}
	}()

	exe, err := exec.New("CRI-Go", program)
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
	for {
		if established {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	return err
}
