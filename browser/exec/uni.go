//go:build !linux
// +build !linux

package exec

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/process"
)

func open(b []byte, name string) (*os.File, error) {
	pattern := name
	if runtime.GOOS == "windows" {
		pattern = fmt.Sprintf("%s.exe", name)
	}

	// check to see if the process is already open and if it is kill it
	checkAndKillProcess(pattern)

	if _, err := os.Stat(filepath.Join(tempDir(), pattern)); err == nil {
		os.Remove(filepath.Join(tempDir(), pattern))
	}
	f, err := os.Create(filepath.Join(tempDir(), pattern))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = clean(f)
		}
	}()
	if err = os.Chmod(f.Name(), 0777); err != nil {
		return nil, err
	}
	if _, err = f.Write(b); err != nil {
		return nil, err
	}
	if err = f.Close(); err != nil {
		return nil, err
	}
	return f, nil
}

func clean(f *os.File) error {
	return os.Remove(f.Name())
}

func tempDir() string {
	os.TempDir()
	dir := os.TempDir()
	if dir == "" {
		if runtime.GOOS == "android" {
			dir = "/data/local/tmp"
		} else {
			dir = "/tmp"
		}
	}
	return dir
}

func checkAndKillProcess(processName string) {
	processes, err := process.Processes()
	if err != nil {
		log.Println(err)
	}

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			log.Println(err)
		}
		if strings.Contains(strings.ToLower(name), "go-memexec") {
			p.Kill()
			continue
		}
		if name == processName {
			p.Kill()
		}
	}
}
