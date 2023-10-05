package browsergo

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"
)

// await for the next message
func (c *CRIService) awaitMessage() (string, EvaluationResponse, []CRICookiesApi, error) {
	var err error
	var message string
	var messageInterface EvaluationResponse
	var cookiesInterface []CRICookiesApi
	c.messageListener = c.receiveMessage()
listen:
	for {
		select {
		case <-c.CTX.Done():
			return "", EvaluationResponse{}, nil, errContextCancelled
		default:
			info, ok := <-c.messageListener
			if ok {
				if c.client.verbose {
					log.Println("[EVENT MESSAGE]", c.uuid, info)
				}
				switch info["service"].(type) {
				case string:
					if info["service"].(string) == c.uuid {
						switch info["message"].(type) {
						case string:
							message = info["message"].(string)
						case map[string]interface{}:
							if info["message"].(map[string]interface{})["value"] != nil {
								switch info["message"].(map[string]interface{})["value"].(type) {
								case bool:
									messageInterface.Value = strconv.FormatBool(info["message"].(map[string]interface{})["value"].(bool))
								case string:
									messageInterface.Value = info["message"].(map[string]interface{})["value"].(string)
								}
							}
							if info["message"].(map[string]interface{})["type"] != nil {
								switch info["message"].(map[string]interface{})["type"].(type) {
								case string:
									messageInterface.Type = info["message"].(map[string]interface{})["type"].(string)
								}
							}
							if info["message"].(map[string]interface{})["description"] != nil {
								switch info["message"].(map[string]interface{})["description"].(type) {
								case string:
									messageInterface.Description = info["message"].(map[string]interface{})["description"].(string)
									if strings.Contains(messageInterface.Description, "SyntaxError") {
										err = errors.New(messageInterface.Description)
									}
								}
							}
							if info["message"].(map[string]interface{})["cookies"] != nil {
								jsonData, _ := json.Marshal(info["message"].(map[string]interface{})["cookies"])
								json.Unmarshal(jsonData, &cookiesInterface)
							}
						default:
							messageInterface = EvaluationResponse{}
						}
						switch info["error"].(bool) {
						case false:
							c.runningBrowser = true
						case true:
							err = errors.New(message)
							break listen
						}
						break listen
					}
				}
			}
		}
	}
	c.removeMessageListener()
	return message, messageInterface, cookiesInterface, err
}

// await for the fetch response
func (c *CRIService) awaitFetchMessage() (CRIGoFetchResponse, error) {
	var err error
	var message string
	var response CRIGoFetchResponse
	c.messageListener = c.receiveMessage()
listen:
	for {
		select {
		case <-c.CTX.Done():
			return CRIGoFetchResponse{}, errContextCancelled
		default:
			info, ok := <-c.messageListener
			if ok {
				if c.client.verbose {
					log.Println("[FETCH MESSAGE]", c.uuid, info)
				}
				switch info["service"].(type) {
				case string:
					if info["service"].(string) == c.uuid {
						switch info["message"].(type) {
						case string:
							message = info["message"].(string)
						default:
							if info["message"].(map[string]interface{})["value"] != nil {
								if info["message"].(map[string]interface{})["value"].(string) != "" {
									err = json.Unmarshal([]byte(info["message"].(map[string]interface{})["value"].(string)), &response)
									if err != nil {
										break
									}
									var decodedString string
									decodedString, err = url.QueryUnescape(decodeBase64(response.Body))
									if err != nil {
										break
									}
									response.Body = decodedString
								}
							}
							if info["message"].(map[string]interface{})["description"] != nil {
								switch info["message"].(map[string]interface{})["description"].(type) {
								case string:
									if info["message"].(map[string]interface{})["description"].(string) != "" {
										err = errors.New(info["message"].(map[string]interface{})["description"].(string))
									}
								}
							}
						}
						switch info["error"].(bool) {
						case false:
							c.runningBrowser = true
						case true:
							err = errors.New(message)
							break listen
						}
						break listen
					}
				}
			}
		}
	}
	c.removeMessageListener()
	return response, err
}

// start the listener function
func (c *CRIService) StartListener(handler func(*InterceptorCommunication)) {
	c.requestListener = c.receiveListener()
	for {
		if !c.listeningToRequests {
			break
		}
		select {
		case <-c.CTX.Done():
			return
		default:
			info, ok := <-c.requestListener
			if ok {
				switch info["service"].(type) {
				case string:
					if info["service"].(string) == c.uuid {
						b, err := json.Marshal(info)
						if err != nil {
							continue
						}
						var handle *InterceptorCommunication
						json.Unmarshal(b, &handle)
						handler(handle)
					}
				}

			}
		}
	}
}
