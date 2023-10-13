package browsergo_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	browsergo "github.com/KakashiHatake324/browser-go"
)

func TestShape(t *testing.T) {
	service, err := browsergo.InitBrowser(true, "")
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
		Proxy:        "207.170.85.114:8263:4SR32:UKXTAR3V",
		Args:         []browsergo.FlagType{},
		Headless:     true,
		OpenDevtools: false,
		WaitLoad:     false,
	}

	if err := instance.OpenBrowser(browserOpts); err != nil {
		log.Fatal(err)
	}
	proxy, _ := browsergo.FormatProxy("207.170.85.114:8263:4SR32:UKXTAR3V")
	shapeSolver := &browsergo.SolveShape{
		BrowserService: instance,
		Context:        context.TODO(),
		Deadline:       10,
		ScriptUrl:      fmt.Sprintf("https://www.newbalance.com/on/demandware.static/Sites-NBUS-Site/-/en_US/v%d/js/nb-common.js?single", time.Now().UnixMicro()),
		UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
		RequestUrl:     "https://www.newbalance.com/on/demandware.static/Sites-NBUS-Site/en_US/Cart-AddProduct",
		RequestMethod:  "POST",
		ProxyString:    proxy,
	}
	shapeHeaders, err := shapeSolver.HandleShape()
	log.Println(shapeHeaders, err)
}
