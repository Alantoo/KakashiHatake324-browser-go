package browsergo

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

// initiate a new instance of browser
func InitBrowser(verbose bool, path string) (*ClientInit, error) {
	// set up a background context
	ctx := context.Background()
	// set up a cancel for the context
	ctx, cancel := context.WithCancel(ctx)
	// init main client
	service := &ClientInit{
		CTX:      ctx,
		cancel:   cancel,
		verbose:  verbose,
		sessions: path,
	}

	if service.verbose {
		log.Println("[INITBrowser] Finding an open port to occupy")
	}
	// initiate the port
	service.findPort()
	if service.verbose {
		log.Println("[INITBrowser] Launching driver server on port", service.port)
	}
	// launch node server with the port
	if err := service.launchServer(); err != nil {
		if service.verbose {
			log.Println("[INITBrowser] Error Launching Server", err)
		}
		return nil, err
	}
	if service.verbose {
		log.Println("[INITBrowser] Launching client and connecting to server..")
	}

	// create the socket client
	service.createMainClient()

	if service.verbose {
		log.Println("[INITBrowser] Service created successfully")
	}
	var err error
	for i := 1; i < 5; i++ {
		if service.conn != nil {
			err = nil
			break
		} else {
			err = errConnectionClosed
		}
		time.Sleep(1 * time.Second)
	}
	return service, err
}

// close the client
func (c *ClientInit) CloseClient() error {
	if c == nil {
		return nil
	}

	c.clientSync.Lock()
	defer c.clientSync.Unlock()

	var err error
	var message []byte
	if c.verbose {
		log.Println("[CLOSING] Closing Main Client, Services:", len(c.Services))
	}

	for i := len(c.Services); i > len(c.Services); i-- {
		c.Services[i].Close()
	}

	if message, err = json.Marshal(map[string]interface{}{
		"action": "kill",
	}); err != nil {
		return err
	}
	if c.conn != nil {
		if err = c.conn.WriteJSON(message); err != nil {
			return err
		}
	}
	if err := c.closeExe(); err != nil {
		return err
	}
	return nil
}
