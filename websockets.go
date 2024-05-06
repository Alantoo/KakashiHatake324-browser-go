package browsergo

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// create the main client to connect with the server
func (s *ClientInit) createMainClient() {
	defer CatchUnhandledError("createMainClient()")
	var err error
	go func() {
		for {
			u := url.URL{Scheme: "ws", Host: fmt.Sprintf("127.0.0.1:%d", s.port), Path: "/"}
			s.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				if s.verbose {
					log.Println("[MAIN CRI-GO ERROR]", err)
				}
				break
			} else {
				log.Println("[MAIN SOCKET CLIENT] CONNECTED")
			main:
				for {
					select {
					case <-s.CTX.Done():
						return
					default:
						if s.conn == nil {
							return
						}
						_, message, err := s.conn.ReadMessage()
						if err != nil {
							if s.verbose {
								if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
									log.Printf("error: %v", err)
								}
								log.Println("[MAIN SOCKET CLIENT]", err)
							}
							break main
						}
						socketMessage := make(map[string]interface{})
						json.Unmarshal(message, &socketMessage)

						switch socketMessage["type"] {
						case "message":
							if s.verbose {
								log.Println("[MAIN SOCKET CLIENT] New Message:", string(message))
							}
						}
					}
				}
			}
			log.Println("[MAIN CRI-GO DISCONNETED] RECONNECTING..")
			time.Sleep(5 * time.Second)
		}
	}()
}

// create a websocket client for each task
func (s *BrowserService) createClient() {
	defer CatchUnhandledError("createClient()")
	var err error
	go func() {
		for {
			u := url.URL{Scheme: "ws", Host: fmt.Sprintf("127.0.0.1:%d", s.client.port), Path: "/"}
			s.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				if s.client.verbose {
					log.Println("[CRI-GO ERROR]", err)
				}
				break
			} else {
			service:
				for {
					select {
					case <-s.CTX.Done():
						return
					default:
						if s.conn == nil {
							return
						}
						_, message, err := s.conn.ReadMessage()
						if err != nil {
							if s.client.verbose {
								if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
									log.Printf("error: %v", err)
								}
								log.Println("[SOCKET CLIENT]", err)
							}
							break service
						}
						socketMessage := make(map[string]interface{})
						json.Unmarshal(message, &socketMessage)

						switch socketMessage["type"] {
						case "count":
						case "message":
							if s.client.verbose {
								log.Println("[SOCKET CLIENT] New Message:", string(message))
							}
							s.sendMessage(socketMessage)
						case "close":
							if s.client.verbose {
								log.Println("[SOCKET CLIENT] New Message:", string(message))
							}
							// close the websocket connection
							if err = s.conn.Close(); err != nil {
								log.Println("[SOCKET CLIENT] Close Error:", err)
							}
							// cancel the context
							s.cancel()
							s.sendMessage(map[string]interface{}{"closed": true})
							return
						case "listener":
							if s.client.verbose {
								log.Println("[SOCKET CLIENT] New Request Listener Message:")
							}
							s.sendListener(socketMessage)
						}
					}
				}
			}
			log.Println("[CRI-GO DISCONNETED] RECONNECTING..")
			time.Sleep(5 * time.Second)
		}
	}()
}

// find an open port
func (s *ClientInit) findPort() error {
	l, close, err := createListener()
	if err != nil {
		return err
	}
	s.port = l.Addr().(*net.TCPAddr).Port
	close()
	return nil
}

// create a listener to find an open port
func createListener() (l net.Listener, close func(), newerr error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	return l, func() {
		_ = l.Close()
	}, nil
}
