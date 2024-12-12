package browsergo

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	ex "os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/KakashiHatake324/browser-go/browser/exec"
)

// launch the server
func (c *ClientInit) launchServer(name string) error {
	if runtime.GOOS != "linux" {
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
		cmd := exe.Command(strconv.Itoa(c.port), strconv.FormatBool(c.verbose))
		if c.verbose {
			log.Println("[LAUNCH SERVER] Start command sent", cmd.Dir)
		}
		c.closeExe = exe.Close
		cmd.Stdout = writer
		if err := cmd.Start(); err != nil {
			return err
		}
		if c.verbose {
			log.Println("[LAUNCH SERVER] Started exec", cmd.Path)
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
	} else {
		// Path to the binary (adjust the path if needed)
		binaryPath := "/usr/local/bin/linux_browser"

		// Check if the binary exists
		if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
			log.Fatalf("Binary does not exist: %s\n", binaryPath)
		}

		// Prepare the command to run the binary
		cmd := ex.Command(binaryPath)

		// Optional: If the binary needs arguments, add them here
		cmd.Args = []string{strconv.Itoa(c.port), strconv.FormatBool(c.verbose)}

		// Set the command output to be visible in the Go program's output
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run the command
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to execute binary: %v\n", err)
		}

		// If the binary ran successfully
		fmt.Println("Binary executed successfully!")

		return nil
	}
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
