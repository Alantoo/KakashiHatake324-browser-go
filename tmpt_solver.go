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
	Deadline    int
	Url         string
	ProxyString string
}

// solve shape with a shape request
func (c *SolveTmpt) HandleTmpt() (string, error) {
	c.Context, c.Cancel = context.WithDeadline(context.TODO(), time.Now().Add(time.Duration(c.Deadline)*time.Second))
	defer c.Cancel()
	var err error
	select {
	case <-c.Context.Done():
		err = errors.New("deadline has exceeded so the context cancelled")
		return "", err
	case <-c.CTX.Done():
		err = errors.New("main context was cancelled")
		return "", err
	default:
		if err := c.Navigate(c.Url, false); err != nil {
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
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
	return "", err
}
