package browsergo

import (
	"encoding/json"
	"log"

	uuid "github.com/satori/go.uuid"
)

// open a new browser
func (c *BrowserService) OpenBrowser(opts *BrowserOpts) error {
	// return error if connection is closed
	if c.conn == nil {
		return errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[OPENING BROWSER] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service":  c.uuid,
		"action":   "open",
		"startUrl": opts.StartUrl,
		"args":     opts.Args,
		"proxy":    opts.Proxy,
		"headless": opts.Headless,
		"devtools": opts.OpenDevtools,
		"timeout":  c.timeout,
	}); err != nil {
		return err
	}
	if c.client.verbose {
		log.Println("[OPENING BROWSER] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, _, _, err = c.awaitMessage()

	// return error
	return err
}

// set new browser cookies
func (c *BrowserService) SetCookies(cookies []*BrowserGoCookies) error {
	// return error if connection is closed
	if c.conn == nil {
		return errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[SETTING COOKIES] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "set-cookies",
		"cookies": cookies,
	}); err != nil {
		return err
	}
	if c.client.verbose {
		log.Println("[SETTING COOKIES] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, _, _, err = c.awaitMessage()

	// return error
	return err
}

// navigate to a url and decide to wait or not
func (c *BrowserService) Navigate(url string, wait bool) error {
	// return error if connection is closed
	if c.conn == nil {
		return errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[NAVIGATING] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "navigate",
		"url":     url,
		"wait":    wait,
	}); err != nil {
		return err
	}
	if c.client.verbose {
		log.Println("[NAVIGATING] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, _, _, err = c.awaitMessage()

	// return error
	return err
}

// wait for an element on the page
func (c *BrowserService) WaitForElement(element string) error {
	// return error if connection is closed
	if c.conn == nil {
		return errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[WFE] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "wait_element",
		"element": element,
	}); err != nil {
		return err
	}
	if c.client.verbose {
		log.Println("[WFE] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, _, _, err = c.awaitMessage()

	// return error
	return err
}

// wait for the page to finish loading
func (c *BrowserService) WaitPageLoad() error {
	// return error if connection is closed
	if c.conn == nil {
		return errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[WPL] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "wait_load",
	}); err != nil {
		return err
	}
	if c.client.verbose {
		log.Println("[WPL] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, _, _, err = c.awaitMessage()

	// return error
	return err
}

// evaluate on the page
func (c *BrowserService) Evaluate(js string) (EvaluationResponse, error) {
	// return error if connection is closed
	if c.conn == nil {
		return EvaluationResponse{}, errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[EVALUATING] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "evaluate",
		"js":      js,
	}); err != nil {
		return EvaluationResponse{}, err
	}
	if c.client.verbose {
		log.Println("[EVALUATING] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return EvaluationResponse{}, err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, response, _, err := c.awaitMessage()

	// return error
	return response, err
}

// click an element
func (c *BrowserService) Click(element string) error {
	// return error if connection is closed
	if c.conn == nil {
		return errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[CLICK] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "click",
		"element": element,
	}); err != nil {
		return err
	}
	if c.client.verbose {
		log.Println("[CLICK] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, _, _, err = c.awaitMessage()

	// return error
	return err
}

// input text like a human
func (c *BrowserService) InputText(inputName, text string) error {
	// return error if connection is closed
	if c.conn == nil {
		return errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[INPUTTING TEXT] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "input_text",
		"name":    inputName,
		"text":    text,
	}); err != nil {
		return err
	}
	if c.client.verbose {
		log.Println("[INPUTTING TEXT] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, _, _, err = c.awaitMessage()

	// return error
	return err
}

// input text like a human
func (c *BrowserService) GetFrame(inputName string) (FrameType, error) {
	// return error if connection is closed
	if c.conn == nil {
		return "", errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[GETTING FRAME] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "get_frame",
		"frame":   uuid.NewV4().String(),
		"name":    inputName,
	}); err != nil {
		return "", err
	}
	if c.client.verbose {
		log.Println("[GETTING FRAME] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return "", err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	new_frame, _, _, err := c.awaitMessage()

	// return error
	return FrameType(new_frame), err
}

// input text like a human
func (c *BrowserService) InputTextFrame(frame FrameType, inputName, text string) error {
	// return error if connection is closed
	if c.conn == nil {
		return errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[INPUTTING TEXT IN FRAME] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "input_text_frame",
		"frame":   frame,
		"name":    inputName,
		"text":    text,
	}); err != nil {
		return err
	}
	if c.client.verbose {
		log.Println("[INPUTTING TEXT IN FRAME] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, _, _, err = c.awaitMessage()

	// return error
	return err
}

// get the current page cookies
func (c *BrowserService) GetCookies() ([]BrowserCookiesApi, error) {
	// return error if connection is closed
	if c.conn == nil {
		return nil, errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[GET COOKIES] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "get_cookies",
	}); err != nil {
		return nil, err
	}
	if c.client.verbose {
		log.Println("[GET COOKIES] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return nil, err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, _, response, err := c.awaitMessage()

	// return error
	return response, err
}

// open a new browser
func (c *BrowserService) RequestListener() (func(), error) {
	// return error if connection is closed
	if c.conn == nil {
		return nil, errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[REQUEST LISTENER] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "request_listen",
	}); err != nil {
		return nil, err
	}
	if c.client.verbose {
		log.Println("[REQUEST LISTENER] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return nil, err
	}
	c.browserSync.Unlock()
	// set the listener to on
	c.listeningToRequests = true
	// return error
	return func() {
		c.listeningToRequests = false
		c.removeReceiveListener()
		c.stopListening()
	}, nil
}

// top the listener
func (c *BrowserService) stopListening() {
	// return error if connection is closed
	if c.conn == nil {
		return
	}

	if c.client.verbose {
		log.Println("[REMOVING LISTENER] Initiating...")
	}

	// build the message going to the server
	message, _ := json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "remove_request_listen",
	})
	if c.client.verbose {
		log.Println("[REMOVING LISTENER] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	c.conn.WriteJSON(message)
	c.browserSync.Unlock()
}

// use fetch while on the page
func (c *BrowserService) Fetch(fetch *BrowserGoFetchRequest) (BrowserGoFetchResponse, error) {
	// return error if connection is closed
	if c.conn == nil {
		return BrowserGoFetchResponse{}, errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[USING FETCH] Initiating...")
	}

	jsonHeaders, err := json.Marshal(fetch.Headers)
	if err != nil {
		return BrowserGoFetchResponse{}, err
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service":        c.uuid,
		"action":         "fetch",
		"url":            fetch.Url,
		"method":         fetch.Method,
		"headers":        string(jsonHeaders),
		"body":           fetch.Body,
		"immediateabort": fetch.ImmediateAbort,
	}); err != nil {
		return BrowserGoFetchResponse{}, err
	}
	if c.client.verbose {
		log.Println("[USING FETCH] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return BrowserGoFetchResponse{}, err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	response, err := c.awaitFetchMessage()
	// return error
	return response, err
}

// extra offers more control over the browser such as iframe
func (c *BrowserService) Extra() error {
	// return error if connection is closed
	if c.conn == nil {
		return errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[USING EXTRA] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "extra",
	}); err != nil {
		return err
	}
	if c.client.verbose {
		log.Println("[USING EXTRA] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, _, _, err = c.awaitMessage()

	// return error
	return err
}

// get the body of the current page
func (c *BrowserService) GetBody() (string, error) {
	// return error if connection is closed
	if c.conn == nil {
		return "", errConnectionClosed
	}

	body, err := c.Evaluate("document.querySelector('body').innerHTML")
	return body.Value, err
}

// input text like a human
func (c *BrowserService) SetBody(body string) error {
	// return error if connection is closed
	if c.conn == nil {
		return errConnectionClosed
	}

	var err error
	var message []byte
	if c.client.verbose {
		log.Println("[SETTING BODY] Initiating...")
	}

	// build the message going to the server
	if message, err = json.Marshal(map[string]interface{}{
		"service": c.uuid,
		"action":  "set-body",
		"body":    body,
	}); err != nil {
		return err
	}
	if c.client.verbose {
		log.Println("[SETTING BODY] Sending message")
	}
	c.browserSync.Lock()
	// send the message to the server
	if err = c.conn.WriteJSON(message); err != nil {
		return err
	}
	c.browserSync.Unlock()

	// listen for the response from the server
	_, _, _, err = c.awaitMessage()

	// return error
	return err
}

// randomize the mouse moving around the page to help against detection
func (c *BrowserService) RandomizeMouseMovements() error {

	if err := c.WaitForElement("body"); err != nil {
		return err
	}
	_, err := c.Evaluate(string(generateRandomMouseMovements()))
	return err
}

// randomize the page being scrolled to help again detection
func (c *BrowserService) RandomizeScrollMovements() error {

	if err := c.WaitForElement("body"); err != nil {
		return err
	}
	_, err := c.Evaluate(string(generateRandomScrollMovements()))
	return err
}

// reset the scroll position to the top of the page
func (c *BrowserService) ResetScrollMovement() error {

	if err := c.WaitForElement("body"); err != nil {
		return err
	}
	_, err := c.Evaluate(string(resetScrollMovement()))
	return err
}
