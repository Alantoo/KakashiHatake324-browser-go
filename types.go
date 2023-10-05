package crigo

import (
	"context"
	"sync"

	"github.com/gorilla/websocket"
)

// client information
type ClientInit struct {
	CTX      context.Context
	cancel   context.CancelFunc
	verbose  bool
	closeExe func() error
	conn     *websocket.Conn
	port     int
	sessions string
}

// cri service struct
type CRIService struct {
	CTX                 context.Context
	cancel              context.CancelFunc
	conn                *websocket.Conn
	uuid                string
	client              *ClientInit
	done                chan bool
	messageListener     <-chan map[string]interface{}
	requestListener     <-chan map[string]interface{}
	messageReceivers    []chan map[string]interface{}
	requestReceivers    []chan map[string]interface{}
	messages            chan map[string]interface{}
	requests            chan map[string]interface{}
	listeningToRequests bool
	runningBrowser      bool
	timeout             int64
	browserSync         sync.Mutex
}

// options for opening a new browser
type BrowserOpts struct {
	StartUrl     string
	Proxy        string
	Args         []FlagType
	Headless     bool
	OpenDevtools bool
	WaitLoad     bool
}

// type for frames
type FrameType string

// type for flags
type FlagType string

// randomize movements or scrolls
type RandomizeType string

// cookies struct to import cookies
type CRIGoCookies struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Domain string `json:"domain"`
}

// cookies returned by the api
type CRICookiesApi struct {
	Domain       string  `json:"domain"`
	Expires      float64 `json:"expires"`
	HTTPOnly     bool    `json:"httpOnly"`
	Name         string  `json:"name"`
	Path         string  `json:"path"`
	Priority     string  `json:"priority"`
	SameParty    bool    `json:"sameParty"`
	SameSite     string  `json:"sameSite,omitempty"`
	Secure       bool    `json:"secure"`
	Session      bool    `json:"session"`
	Size         int     `json:"size"`
	SourcePort   int     `json:"sourcePort"`
	SourceScheme string  `json:"sourceScheme"`
	Value        string  `json:"value"`
}

// response from js evaluations
type EvaluationResponse struct {
	Value       string `json:"value"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// request format when making a fetch response
type CRIGoFetchRequest struct {
	Url            string                 `json:"url"`
	Method         string                 `json:"method"`
	Headers        map[string]interface{} `json:"headers"`
	Body           string                 `json:"body,omitempty"`
	ImmediateAbort bool                   `json:"immediateabort,omitempty"`
}

// response format when making a fetch request
type CRIGoFetchResponse struct {
	StatusCode int                    `json:"status"`
	Headers    map[string]interface{} `json:"headers"`
	Body       string                 `json:"body"`
}

// interceptor response struct
type InterceptorCommunication struct {
	Request  interceptorFormat `json:"request,omitempty"`
	Response interceptorFormat `json:"response,omitempty"`
}

// interceptor format for requests/responses
type interceptorFormat struct {
	DocumentUrl string                 `json:"documentUrl,omitempty"`
	Url         string                 `json:"url,omitempty"`
	Method      string                 `json:"method,omitempty"`
	Status      int                    `json:"status,omitempty"`
	StatusText  string                 `json:"statusText,omitempty"`
	Headers     map[string]interface{} `json:"headers,omitempty"`
	Body        string                 `json:"body,omitempty"`
}
