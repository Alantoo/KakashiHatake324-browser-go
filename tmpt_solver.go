package browsergo

import (
	"context"
	"errors"
	"time"
)

type SolveTmpt struct {
	*BrowserService
	// seconds to stop the deadline
	Context     context.Context
	Cancel      func()
	Deadline    int64
	Url         string
	ProxyString string
}

// solve shape with a shape request
func (c *SolveTmpt) HandleTmpt() (string, error) {
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
		if err := c.Navigate(c.Url, false); err != nil {
			return "", err
		}
		time.Sleep(250 * time.Millisecond)
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
				cookies, err := c.GetCookies()
				if err != nil {
					return "", err
				}
				for _, c := range cookies {
					if c.Name == "tmpt" {
						return c.Value, nil
					}
				}
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}
	return "", err
}
