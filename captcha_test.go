package browsergo_test

import (
	"context"
	"log"
	"testing"

	browsergo "github.com/KakashiHatake324/browser-go"
)

const captchaProxy = "31.128.120.186:63690:yvyyhlia:s4XI854jtl"

func TestCaptcha(t *testing.T) {
	service, err := browsergo.InitBrowser("browser-go-test", true, true, "")
	if err != nil {
		log.Fatal(err)
	}
	instance, err := service.NewService(context.Background(), 400000)
	if err != nil {
		log.Fatal(err)
	}
	defer instance.Close()

	browserOpts := &browsergo.BrowserOpts{
		StartUrl: "",
		Proxy:    captchaProxy,
		Args:     []browsergo.FlagType{
			/*
				browsergo.DisableInProcessStackTraces,
				browsergo.DisableBackgroundMode,
				browsergo.DisableBackgroundNetworking,
				browsergo.DisableBackgroundOccludedWindows,
				browsergo.DisableComponentExtensionsWithBackgroundPages,
				browsergo.NoFirstRun,
				browsergo.NoDefaultBrowserCheck,
				browsergo.RandomWindowSize(),
				browsergo.DisableSync,
			*/
		},
		Headless:     true,
		OpenDevtools: false,
		WaitLoad:     false,
		ForceChrome:  true,
	}

	if err := instance.OpenBrowser(browserOpts); err != nil {
		log.Fatal(err)
	}
	captchaSolver := &browsergo.SolveCaptcha{
		BrowserService: instance,
		Context:        context.TODO(),
		Deadline:       20,
		Url:            "https://identity.ticketmaster.com/sign-in?integratorId=prd1741.iccp&placementId=mytmlogin&redirectUri=https://www.ticketmaster.com/",
		SiteKey:        "6LdWxZEkAAAAAIHtgtxW_lIfRHlcLWzZMMiwx9E1",
		Action:         "Event",
	}
	tmpt, err := captchaSolver.HandleCaptcha()
	log.Println(tmpt, err)
}
