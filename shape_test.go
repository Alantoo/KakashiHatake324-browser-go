package browsergo_test

import (
	"context"
	"log"
	"testing"

	browsergo "github.com/KakashiHatake324/browser-go"
)

const shapeProxy = ""

func TestShape(t *testing.T) {
	service, err := browsergo.InitBrowser(false, "")
	if err != nil {
		log.Fatal(err)
	}

	instance, err := service.NewService(context.Background(), 220000)
	if err != nil {
		log.Fatal(err)
	}
	defer instance.Close()

	browserOpts := &browsergo.BrowserOpts{
		StartUrl:     "",
		Proxy:        shapeProxy,
		Args:         []browsergo.FlagType{},
		Headless:     true,
		OpenDevtools: false,
		WaitLoad:     false,
	}

	if err := instance.OpenBrowser(browserOpts); err != nil {
		log.Fatal(err)
	}
	proxy, _ := browsergo.FormatProxy(shapeProxy)
	shapeSolver := &browsergo.SolveShape{
		BrowserService: instance,
		Context:        context.TODO(),
		Deadline:       15,
		ScriptUrl:      browsergo.Nordstrom,
		UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
		RequestUrl:     "https://www.nordstrom.com/api/checkout/web/checkout/v1/shopping-bag/items?postalcode=08540",
		RequestMethod:  "POST",
		ProxyString:    proxy,
	}
	shapeHeaders, err := shapeSolver.HandleShape()
	log.Println(shapeHeaders, err)
}
