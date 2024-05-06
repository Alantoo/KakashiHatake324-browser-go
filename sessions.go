package browsergo

import (
	"fmt"
	"os"
	"strings"
)

// create a session with a folder name
func (s *ClientInit) CreateSession(name string) error {
	name = strings.ToLower(name)
	sessionPath := fixPath(fmt.Sprintf("%s%sprofile-%s", s.sessions, pathSeparator, name))
	if _, err := os.Stat(sessionPath); os.IsNotExist(err) {
		if err := os.Mkdir(sessionPath, os.ModePerm); err != nil {
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
	sessionPath := fixPath(fmt.Sprintf("%s%sprofile-%s", s.sessions, pathSeparator, name))
	if err := os.RemoveAll(sessionPath); os.IsNotExist(err) {
		return err
	}
	return nil
}

// get a session string to pass through the args settings
func (s *ClientInit) GetSessionName(name string) (string, error) {
	name = strings.ToLower(name)
	sessionPath := fixPath(fmt.Sprintf("%s%sprofile-%s", s.sessions, pathSeparator, name))
	if _, err := os.Stat(sessionPath); os.IsNotExist(err) {
		return sessionPath, nil
	} else {
		return sessionPath, nil
	}
}
