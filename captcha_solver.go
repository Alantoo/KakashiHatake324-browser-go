package browsergo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/KakashiHatake324/mockjs"
)

type SolveCaptcha struct {
	*BrowserService
	// seconds to stop the deadline
	Context  context.Context
	Cancel   func()
	Deadline int64
	Url      string
	SiteKey  string
	Action   string
}

// solve shape with a shape request
func (c *SolveCaptcha) HandleCaptcha() (string, error) {
	c.Context, c.Cancel = context.WithDeadline(c.BrowserService.CTX, time.Now().Add(time.Duration(c.Deadline)*time.Second))
	timing := time.NewTimer(time.Duration(c.Deadline) * 60 * time.Second)
	defer c.Cancel()
	var err error
	select {
	case <-c.Context.Done():
		err = errors.New("deadline has exceeded so the context cancelled")
		return "", err
	case <-c.CTX.Done():
		err = errors.New("main context was cancelled")
		return "", err
	case <-c.BrowserService.CTX.Done():
		err = errors.New("main context was cancelled")
		return "", err
	default:
		time.Sleep(1000 * time.Millisecond)
		if err := c.Navigate("https://www.ticketmaster.com/event/0000617D0901855C", false); err != nil {
			return "", err
		}

		if err := c.SetBody(fmt.Sprintf(`<html>
        <head>
          <title>Recaptcha Harvester</title>
        </head>
          <script>console.log('getting started');</script>
        <body>
          <div>
            <script src="https://www.google.com/recaptcha/enterprise.js?render=6LdWxZEkAAAAAIHtgtxW_lIfRHlcLWzZMMiwx9E1" defer=""></script>
            <script>
            grecaptcha.enterprise.ready(function () {
              grecaptcha.enterprise.execute("%s",{
                  action: "%s"
              }).then(function (token) {
              document.getElementById("g-recaptcha-response").value = token;
              }).catch((error) => {
                console.log(error);
              });
            });
            </script>
            <input id="g-recaptcha-response" name="g-recaptcha-response" value="">
          </div>
        </body>
      </html>`, c.SiteKey, c.Action)); err != nil {
			return "", err
		}

		var complete bool
		for !complete {
			select {
			case <-c.Context.Done():
				err = errors.New("deadline has exceeded so the context cancelled")
				return "", err
			case <-c.CTX.Done():
				err = errors.New("main context was cancelled")
				return "", err
			case <-c.BrowserService.CTX.Done():
				err = errors.New("main context was cancelled")
				return "", err
			case <-timing.C:
				err = errors.New("main context was cancelled")
				return "", err
			default:
				token, err := c.Evaluate("document.querySelector('[name=\"g-recaptcha-response\"]').value")
				if err != nil {
					log.Println(err)
				}
				log.Println(c.GetBody())
				if token.Value != "" {
					return mockjs.InitWindow().EncodeURIComponent(token.Value), nil
				}
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}
	return "", err
}
