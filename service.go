package crigo

import (
	"context"
	"encoding/json"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

// create a new service from the client
func (c *ClientInit) NewService(ctx context.Context, timeout int64) (*CRIService, error) {
	// set up a background context
	if ctx == nil {
		ctx = context.Background()
	}
	// set up a cancel for the context
	ctx, cancel := context.WithCancel(ctx)
	// set up the new service
	service := &CRIService{
		CTX:              ctx,
		cancel:           cancel,
		uuid:             uuid.NewV4().String(),
		client:           c,
		done:             make(chan bool),
		timeout:          timeout,
		messages:         make(chan map[string]interface{}),
		requests:         make(chan map[string]interface{}),
		messageReceivers: nil,
		requestReceivers: nil,
	}

	go service.listenMessage()
	go service.listenRequests()

	// create the socket client
	service.createClient()

	for i := 1; i < 5; i++ {
		if service.conn != nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	// if there is no connection return error
	if service.conn == nil {
		return nil, errConnectionClosed
	}

	// return the service
	return service, nil
}

// close the browser instance
func (c *CRIService) Close() error {
	if c == nil {
		return nil
	}
	// cancel the context
	c.cancel()

	// return error if connection is closed
	if c.conn == nil {
		return c.closeChannels()
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[CLOSING] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "close",
	}); err != nil {
		return err
	}
	if c.client.verbose {
		log.Println("[CLOSING] Sending message")
	}

	c.browserSync.Lock()
	// send the message to the server
	if c.conn != nil {
		if err = c.conn.WriteJSON(message); err != nil {
			return err
		}
	}
	c.browserSync.Unlock()

	// close the websocket connection
	c.conn = nil

	// return error
	return c.closeChannels()
}
