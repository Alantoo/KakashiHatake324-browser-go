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
func (c *ClientInit) launchServer(name string) error {
	if name == "" {
		name = "browser-go"
	}
	var err error
	exe, err := exec.New(name, getProgram())
	if err != nil {
		return err
	}
	if c.verbose {
		log.Println("[LAUNCH SERVER] Executable generated")
	}
	reader, writer := io.Pipe()
	scanner := bufio.NewScanner(reader)
	cmd := exe.CommandContext(context.Background(), strconv.Itoa(c.port), strconv.FormatBool(c.verbose))
	if c.verbose {
		log.Println("[LAUNCH SERVER] Start command sent")
	}
	c.closeExe = exe.Close
	cmd.Stdout = writer
	if err := cmd.Start(); err != nil {
		return err
	}
	if c.verbose {
		log.Println("[LAUNCH SERVER] Waiting for server to load..")
	}
	var loaded bool
	go func() {
		for scanner.Scan() {
			if c.verbose {
				fmt.Println(scanner.Text())
			}
			if !loaded {
				if strings.Contains(scanner.Text(), "ERROR") {
					if c.verbose {
						log.Println("[LAUNCH SERVER] error launching server:", err)
					}
					err = errServerConnection
					loaded = true
				} else if strings.Contains(scanner.Text(), "server is running") {
					if c.verbose {
						log.Println("[LAUNCH SERVER] server launched successfully")
					}
					loaded = true
				}
			}
		}
	}()

	for !loaded {
		select {
		case <-c.CTX.Done():
			return nil
		default:
			time.Sleep(1 * time.Second)
		}
	}
	return err
}
