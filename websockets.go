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
					log.Println("[MAIN browser-go ERROR]", err)
				}
			} else {
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
							break
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
			log.Println("[MAIN browser-go DISCONNETED] RECONNECTING..")
			time.Sleep(5 * time.Second)
		}
	}()
}

// create a websocket client for each task
func (s *BrowserService) createClient() {
	var err error
	go func() {
		for {
			u := url.URL{Scheme: "ws", Host: fmt.Sprintf("127.0.0.1:%d", s.client.port), Path: "/"}
			s.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				if s.client.verbose {
					log.Println("[browser-go ERROR]", err)
				}
			} else {
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
							break
						}
						socketMessage := make(map[string]interface{})
						json.Unmarshal(message, &socketMessage)

						switch socketMessage["type"] {
						case "message":
							if s.client.verbose {
								log.Println("[SOCKET CLIENT] New Message:", string(message))
							}
							s.sendMessage(socketMessage)
						case "listener":
							if s.client.verbose {
								log.Println("[SOCKET CLIENT] New Request Listener Message:")
							}
							s.sendListener(socketMessage)
						}
					}
				}
			}
			log.Println("[browser-go DISCONNETED] RECONNECTING..")
			time.Sleep(5 * time.Second)
		}
	}()
}

// find an open port
func (s *ClientInit) findPort() {
	l, close := createListener()
	s.port = l.Addr().(*net.TCPAddr).Port
	close()
}

// create a listener to find an open port
func createListener() (l net.Listener, close func()) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	return l, func() {
		_ = l.Close()
	}
}
