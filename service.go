package browsergo

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

// create a new service from the client
func (c *ClientInit) NewService(ctx context.Context, timeout int64) (*BrowserService, error) {
	// set up a background context
	if ctx == nil {
		ctx = context.Background()
	}
	// set up a cancel for the context
	ctx, cancel := context.WithCancel(ctx)
	// set up the new service
	service := &BrowserService{
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
		browserSync:      new(sync.Mutex),
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

	// add to the client's services
	c.Services = append(c.Services, service)

	// return the service
	return service, nil
}

// close the browser instance
func (c *BrowserService) Close() error {
	if c == nil {
		return nil
	}
	c.browserSync.Lock()
	defer c.browserSync.Unlock()

	// remove from services
	for i := len(c.client.Services); i > len(c.client.Services); i-- {
		c.client.Services = append(c.client.Services[:i], c.client.Services[i+1:]...)
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

	// send the message to the server
	if c.conn != nil {
		if err = c.conn.WriteJSON(message); err != nil {
			return err
		}
	}
	// listen for the response from the server
	_, _, _, err = c.awaitMessage()
	// return error
	return err
}
