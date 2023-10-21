package browsergo

import (
	"compress/gzip"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/andybalholm/brotli"
	uuid "github.com/satori/go.uuid"
)

type SolveShape struct {
	*BrowserService
	// seconds to stop the deadline
	Context       context.Context
	Cancel        func()
	Deadline      int
	ScriptUrl     ShapeSite
	UserAgent     string
	RequestUrl    string
	RequestMethod string
	ProxyString   string
}

// solve shape with a shape request
func (c *SolveShape) HandleShape() (map[string][]string, error) {
	c.Context, c.Cancel = context.WithDeadline(context.TODO(), time.Now().Add(time.Duration(c.Deadline)*time.Second))
	defer c.Cancel()
	headers := make(map[string][]string)
	var err error
	select {
	case <-c.Context.Done():
		err = errors.New("deadline has exceeded so the context cancelled")
		return headers, err
	case <-c.CTX.Done():
		err = errors.New("main context was cancelled")
		return headers, err
	default:
		var complete bool
		domainUrl, err := url.Parse(c.RequestUrl)
		if err != nil {
			return headers, err
		}

		if !c.shapeLoaded {
			if err := c.Navigate(fmt.Sprintf("https://%s", domainUrl.Host), false); err != nil {
				return headers, err
			}

			time.Sleep(1 * time.Second)

			if err := c.SetBody(shapeBody); err != nil {
				return headers, err
			}

			script, err := c.makeScriptRequest()
			if err != nil {
				return headers, err
			}

			script = strings.ReplaceAll(script, "u=\"/", fmt.Sprintf("u=\"https://%s/", domainUrl.Host))
			fullScript := fmt.Sprintf("(async() => { %s })()", script)
			if _, err = c.Evaluate(fullScript); err != nil {
				return headers, err
			}

			c.shapeLoaded = true
		}
		close, err := c.RequestListener()
		if err != nil {
			return headers, err
		}
		defer close()

		go c.StartListener(
			func(intercept *InterceptorCommunication) {
				select {
				case <-c.Context.Done():
					err = errors.New("deadline has exceeded so the context cancelled")
					return
				case <-c.CTX.Done():
					err = errors.New("main context was cancelled")
					return
				default:
					if intercept.Request.Url == c.RequestUrl && intercept.Request.Method == c.RequestMethod {
						for k, v := range intercept.Request.Headers {
							if strings.HasSuffix(strings.ToLower(k), "-a") ||
								strings.HasSuffix(strings.ToLower(k), "-b") ||
								strings.HasSuffix(strings.ToLower(k), "-c") ||
								strings.HasSuffix(strings.ToLower(k), "-d") ||
								strings.HasSuffix(strings.ToLower(k), "-e") ||
								strings.HasSuffix(strings.ToLower(k), "-f") ||
								strings.HasSuffix(strings.ToLower(k), "-z") ||
								strings.HasSuffix(strings.ToLower(k), "-a0") ||
								strings.HasSuffix(strings.ToLower(k), "-z0") {
								headers[strings.ToLower(k)] = []string{v.(string)}
							}
						}
						if len(headers) != 0 {
							complete = true
							return
						}
					}
				}
			},
		)

		time.Sleep(3 * time.Second)

		request := &BrowserGoFetchRequest{
			Url:            c.RequestUrl,
			Method:         c.RequestMethod,
			Headers:        map[string]interface{}{},
			Body:           "",
			ImmediateAbort: true,
		}
		_, err = c.Fetch(request)
		if err != nil {
			return headers, err
		}

		for !complete && err == nil {
			select {
			case <-c.Context.Done():
				err = errors.New("deadline has exceeded so the context cancelled")
				return headers, err
			case <-c.CTX.Done():
				err = errors.New("main context was cancelled")
				return headers, err
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
	return headers, err
}

func (c *SolveShape) makeScriptRequest() (string, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	if c.ProxyString != "" {
		proxyParsed, err := url.Parse(c.ProxyString)
		if err != nil {
			return "", errors.New("error solving shape: parsing proxy")
		}
		transport.Proxy = http.ProxyURL(proxyParsed)
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   6 * time.Second,
	}

	req, err := http.NewRequestWithContext(c.Context, "GET", string(c.ScriptUrl), nil)
	if err != nil {
		return "", errors.New("error solving shape: http.NewRequestWithContext()")
	}
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "none")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", c.UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("error solving shape: client.Do(req)")
	} else {
		if resp.StatusCode != 200 {
			return "", errors.New("error solving shape: proxy banned")
		}
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", errors.New("error solving shape: io.ReadAll(resp.Body)")
		}
		if strings.Join(resp.Header["Content-Encoding"], "") == "gzip" {
			rdata := strings.NewReader(string(bodyText))
			r, err := gzip.NewReader(rdata)
			if err != nil {
				return "", errors.New("error solving shape: gzip")
			}
			bodyText, err = io.ReadAll(r)
			if err != nil {
				return "", errors.New("error solving shape: io.ReadAll(r)")
			}
		} else if strings.Join(resp.Header["Content-Encoding"], "") == "br" {
			rdata := strings.NewReader(string(bodyText))
			r := brotli.NewReader(rdata)
			bodyText, err = io.ReadAll(r)
			if err != nil {
				return "", errors.New("error solving shape: br")
			}
		}
		bodyString := string(bodyText)
		return bodyString, nil
	}
}

type ShapeSite string

var (
	Nordstrom  = ShapeSite("https://www.nordstrom.com/mwp/integration/ns_common.js?async")
	NewBalance = ShapeSite(fmt.Sprintf("https://www.newbalance.com/on/demandware.static/Sites-NBUS-Site/-/en_US/v%d/js/nb-common.js?single", time.Now().UnixMicro()))
)

var shapeBody = fmt.Sprintf(`<html><head><title>Shape</title></head><body>%s</body></html>`, uuid.NewV4().String())
