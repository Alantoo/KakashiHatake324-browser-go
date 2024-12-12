package browsergo

import (
	"bufio"
	"context"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
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
		pids, err := getRunningPIDs()
		if err != nil {
			log.Fatalf("Failed to get running PIDs: %v", err)
		}

		select {
		case <-c.CTX.Done():
			return nil
		default:
			if c.verbose {
				log.Println("Running PIDS", strings.Join(pids, ":"))
			}
			time.Sleep(1 * time.Second)
		}
	}
	return err
}

func getRunningPIDs() ([]string, error) {
	var pids []string

	// Read the /proc directory
	files, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		// Check if the directory name is numeric (indicating a PID)
		if file.IsDir() {
			pids = append(pids, file.Name())
		}
	}

	return pids, nil
}
