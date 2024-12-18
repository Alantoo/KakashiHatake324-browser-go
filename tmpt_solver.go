package browsergo

import (
	"context"
	"errors"
	"time"
)

type SolveTmpt struct {
	*BrowserService
	/****** The main context */
	Context context.Context
	/****** The cancel function */
	Cancel func()
	/****** Deadline in seconds for the job */
	Deadline int64
	/****** The url the tmpt comes from */
	Url string
	/****** The proxy string in ip:port:user:pass format */
	ProxyString string
	/****** MacOs devices are defaulted to opera but we can force chrome */
	ForceChrome bool
}

// solve shape with a shape request
func (c *SolveTmpt) HandleTmpt() (string, error) {
	var err error
	browserOpts := &BrowserOpts{
		Proxy:   c.ProxyString,
		Profile: "shape_gen",
		Args: []FlagType{
			Incognito,
			DisableAutomations,
			EnableLowEndMode,
			ForceDeviceScaleFactor1,
			DisableInProcessStackTraces,
			NoSandbox,
			DisableTranslate,
			NoFirstRun,
			NoDefaultBrowserCheck,
			RandomWindowSize(),
			DisableFeatures([]string{"LayoutNG", "PreloadMediaEngagementData", "MediaEngagementBypassAutoplayPolicies"}),
		},
		Headless:     true,
		OpenDevtools: false,
		WaitLoad:     false,
		ForceChrome:  c.ForceChrome,
		tmpt:         false,
	}

	if err := c.OpenBrowser(browserOpts); err != nil {
		return "", err
	}
	c.Context, c.Cancel = context.WithDeadline(c.CTX, time.Now().Add(time.Duration(c.Deadline)*time.Second))
	timing := time.NewTimer(time.Duration(c.Deadline) * 60 * time.Second)
	defer c.Cancel()
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

		if err := c.ClearCookies(); err != nil {
			return "", err
		}

		time.Sleep(100 * time.Millisecond)
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
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
	return "", err
}
