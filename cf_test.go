package browsergo_test

import (
	"context"
	"log"
	"strings"
	"testing"
	"time"

	browsergo "github.com/KakashiHatake324/browser-go"
)

func TestST(t *testing.T) {
	service, err := browsergo.InitBrowser(true, "")
	if err != nil {
		log.Fatal(err)
	}

	defer service.CloseClient()

	instance, err := service.NewService(context.Background(), 220000)
	if err != nil {
		log.Fatal(err)
	}

	defer instance.Close()

	browserOpts := &browsergo.BrowserOpts{
		StartUrl:     "",
		Proxy:        proxy,
		Args:         []browsergo.FlagType{},
		Headless:     false,
		OpenDevtools: false,
		WaitLoad:     false,
	}

	if err := instance.OpenBrowser(browserOpts); err != nil {
		log.Fatal(err)
	}

	defer instance.Close()

	if err := instance.Navigate("https://funko.com/", true); err != nil {
		log.Fatal(err)
	}

	for {
		body, err := instance.GetBody()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(body)
		if strings.Contains(body, "class=\"hdr__user\"") {
			break
		}

		time.Sleep(1 * time.Second)
	}

	if err := instance.WaitForElement("[class=\"hdr__user\"]"); err != nil {
		log.Fatal(err)
	}

	cookies, err := instance.GetCookies()
	if err != nil {
		log.Fatal(err)
	}
	for _, cookie := range cookies {
		log.Println(cookie.Name)
		log.Println(cookie.Value)
	}
	instance.Close()
}
