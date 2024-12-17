Install
`go get github.com/KakashiHatake324/browser-go`

Initiate the service worker when you start your program

```go
package main

import browsergo "github.com/KakashiHatake324/browser-go"

var (
	BrowserService *browsergo.ClientInit
)

func init() {
    var err error
	BrowserService, err = browsergo.InitBrowser(false)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
    defer service.CloseClient()
}
```

Use the service worker within your application to create browser instance and optionally let the package handle your sessions

```go
    // create an instance with the service and set the timeout
	instance, err := service.NewService(time.Duration(10 * time.Second))
	if err != nil {
		log.Fatal(err)
	}
	// cancel the instance context @can be used to stop tasks
	defer instance.Cancel()
	
	browsergo.CreateSession(sessionName, sessionsPath)
	sessionFlag, _ := browsergo.GetSessionFlag(sessionName, sessionsPath)

	browserOpts := &browsergo.BrowserOpts{
		StartUrl: "",
		Proxy:    proxy,
		Args: []browsergo.FlagType{
			sessionFlag,
			browsergo.RandomWindowSize(),
			browsergo.EnableFeatures([]string{"ReduceUserAgent", "NetworkService", "NetworkServiceInProcess"}),
			browsergo.EnableBlinkFeatures([]string{"IdleDetection"}),
		},
		Headless:     false,
		OpenDevtools: false,
		WaitLoad:     false,
	}

    // opens the browser with the selected options
	if err := instance.OpenBrowser(browserOpts); err != nil {
		log.Fatal(err)
	}
```

Setting cookies

```go
    // initiate the cookies struct
	var cookies []*BrowserGoCookies
    // append new cookies
	cookies = append(cookies, &BrowserGoCookies{
		Name:   "hi",
		Value:  "cookie",
		Domain: ".google.com",
	})

	if err := instance.SetCookies(cookies); err != nil {
		log.Fatal(err)
	}
```

Navigating to a webpage

```go
    // navigate and true to wait for it to load
	if err := instance.Navigate("https://www.github.com", true); err != nil {
		log.Fatal(err)
	}
```

Waiting for an element

```go
    // wait for an element before proceeding, keep in mind the timeout will be active
	if err := instance.WaitForElement("[data-var=\"userName\"]"); err != nil {
		log.Fatal(err)
	}
```

Evaluate javascript and get the results

```go
    // evaluate js and use the returned information [value, description, type]
	username, err := instance.Evaluate("document.querySelector('[data-var=\"userName\"]')?.innerHTML")
	if err != nil {
		log.Fatal(err)
	}
```

Input text with human behavior & Perform click with human behavior

```go
    // first focus on the input element
	if _, err := instance.Evaluate("document.querySelector('[name=\"password\"]').focus()"); err != nil {
		log.Fatal(err)
	}
    // call the input text function that replicates human typing
	if err := instance.InputText(password); err != nil {
		log.Fatal(err)
	}
    // click on a certain dom
	if err := instance.Click("[type=\"submit\"]"); err != nil {
		log.Fatal(err)
	}
```

Using fetch api

```go
    // first create a struct with the request information
	request := &BrowserGoFetchRequest{
		Url:     fmt.Sprintf("https://api.nike.com/product_feed/threads/v2?filter=exclusiveAccess(true,false)&filter=language(en)&filter=marketplace(US)&filter=channelId(%s)&filter=productInfo.merchProduct.styleColor(%s)", nikeWebChannelId, productID),
		Method:  "GET",
		Headers: map[string]interface{}{"accept": "application/json"},
		Body:    "",
		ImmediateAbort: false,
	}
	// then call the fetch api to retrieve the results [fetch will use the page context to make the request with proper tls, cookies, etc..]
	fetchResponse, err := instance.Fetch(request)
	if err != nil {
		log.Fatal(err)
	}
```


Listen to requests and access its information

```go
	// initiate the request listener
  	close, err := instance.RequestListener()
	if err != nil {
		log.Fatal(err)
	}
	defer close()
 	// you can start the listener wherever you want and it will retrieve request and response data 
	// @TODO request intercepts
	go instance.StartListener(func(intercept map[string]interface{}) {
		log.Println(intercept)
	})
```


Close the browser instance

```go
    // call the close function to close the browser
	if err := instance.Close(); err != nil {
		log.Fatal(err)
	}
```
