package crigo

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

// initiate a new instance of cri
func InitCRI(verbose bool, path string) (*ClientInit, error) {
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
		log.Println("[INITCRI] Finding an open port to occupy")
	}
	// initiate the port
	service.findPort()
	if service.verbose {
		log.Println("[INITCRI] Launching driver server on port", service.port)
	}
	// launch node server with the port
	if err := service.launchServer(); err != nil {
		if service.verbose {
			log.Println("[INITCRI] Error Launching Server", err)
		}
		return nil, err
	}
	if service.verbose {
		log.Println("[INITCRI] Launching client and connecting to server..")
	}

	// create the socket client
	service.createMainClient()

	if service.verbose {
		log.Println("[INITCRI] Service created successfully")
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
	var err error
	var message []byte
	if c.verbose {
		log.Println("[CLOSING] Closing Main Client")
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
