package browsergo_test

import (
	"context"
	"log"
	"testing"

	browsergo "github.com/KakashiHatake324/browser-go"
)

const shapeProxy = "p6.mushroomproxy.com:8000:1F7Aq7B16-mushroom-b70bf8ba79!g-us!f-hddv!sid-yLVblzPnMzY:1d7k3jfj3kiuf3f"

func TestShape(t *testing.T) {
	service, err := browsergo.InitBrowser(true, "")
	if err != nil {
		log.Fatal(err)
	}

	instance, err := service.NewService(context.Background(), 400000)
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
		Deadline:       45,
		ScriptUrl:      browsergo.NewBalance,
		UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		RequestUrl:     "https://www.newbalance.com/on/demandware.static/Sites-NBUS-Site/en_US/CheckoutServices-PlaceOrder?termsconditions=undefined&DFReferenceId=",
		RequestMethod:  "POST",
		ProxyString:    proxy,
	}
	shapeHeaders, err := shapeSolver.HandleShape()
	log.Println(shapeHeaders, err)
}
