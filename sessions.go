package browsergo

import (
	"fmt"
	"os"
	"strings"
)

// create a session with a folder name
func (s *ClientInit) CreateSession(name string) error {
	name = strings.ToLower(name)
	if _, err := os.Stat(fmt.Sprintf("%s/_profile-%s", s.sessions, name)); os.IsNotExist(err) {
		if err := os.Mkdir(fmt.Sprintf("%s/_profile-%s", s.sessions, name), os.ModePerm); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return errSessionExists
	}
}

// create a session with a folder name
func (s *ClientInit) DeleteSession(name string) error {
	name = strings.ToLower(name)
	if err := os.RemoveAll(fmt.Sprintf("%s/_profile-%s", s.sessions, name)); os.IsNotExist(err) {
		return err
	}
	return nil
}

// get a session string to pass through the args settings
func (s *ClientInit) GetSessionFlag(name string) (FlagType, error) {
	name = strings.ToLower(name)
	if _, err := os.Stat(fmt.Sprintf("%s/_profile-%s", s.sessions, name)); os.IsNotExist(err) {
		return "", errSessionDoesntExists
	} else {
		return FlagType(fmt.Sprintf("--user-data-dir=%s/_profile-%s", s.sessions, name)), nil
	}
}
