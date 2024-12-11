package browsergo_test

import (
	"context"
	"log"
	"testing"

	browsergo "github.com/KakashiHatake324/browser-go"
)

const tmptProxy = "31.128.120.186:63690:yvyyhlia:s4XI854jtl"

func TestTmpt(t *testing.T) {
	service, err := browsergo.InitBrowser("browser-go-test", true, false, "")
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
		Proxy:    tmptProxy,
		Args: []browsergo.FlagType{
			browsergo.DisableInProcessStackTraces,
			browsergo.DisableBackgroundMode,
			browsergo.DisableBackgroundNetworking,
			browsergo.DisableBackgroundOccludedWindows,
			browsergo.DisableComponentExtensionsWithBackgroundPages,
			browsergo.NoFirstRun,
			browsergo.NoDefaultBrowserCheck,
			browsergo.RandomWindowSize(),
			browsergo.DisableSync,
		},
		Profile:      "",
		Headless:     true,
		OpenDevtools: false,
		WaitLoad:     false,
	}

	if err := instance.OpenBrowser(browserOpts); err != nil {
		log.Fatal(err)
	}
	proxy, _ := browsergo.FormatProxy(tmptProxy)
	tmptSolver := &browsergo.SolveTmpt{
		BrowserService: instance,
		Context:        context.TODO(),
		Deadline:       45,
		Url:            "https://auth.ticketmaster.com/as/authorization.oauth2?redirect_uri=https%3A%2F%2Fmy.ticketmaster.com%2Fsettings&response_type=code&scope=openid+profile+phone+email+tm&client_id=8bf7204a7e97.web.ticketmaster.us&integratorId=accounts&placementId=settings&lang=en-us&visualPresets=tm&state=authenticatedUser&policySelector=requireEditPhone&intSiteToken=tm-us&updatePhoneToken=eyJhbGciOiJSU0EtT0FFUC0yNTYiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIn0.rgbXwsa4oeWLvvMWKhld3UcL6gE6tBaa41O4FXSIClakf-uAygQbHHzXGgfnLaoG5jcvE9cUiNIA3tvo080A4YLYFePyCCoY-ny0_qRqqlMSGx_tJJYnl8DzDC7QMSWMD8_fgQXWEnwQeVv7Dsg5Z6N0NcTQA_me7HdeoCvMEXKlGWm7qXoyCgKxcEBSb5Iqv8rcxt1ItW33mu8wpbmSKwi7nnt00qgRWupF1-uzSz8DgyQM3uQRcw7CZ6glv2XGcD-7oR2LLZmCOYsma9WooLGc3OZmNFM2jtu4DzvkFIA61D2N8TD7_FGZQij_m1DPE6uzLNKf-7Boh3dtg-bMiA.a3q3XyMoNlqxWeuIBX3HmQ.A7VEuAtkW26q-WQbjehsQZtFTcsOm7m0lGTsMkEHhbL-QFn0LxTHoID-vgLBww3BpOtD5Tspwu7XIizrayB1iFGvBJYCjjOmBk9Um3uq33NVUUv64sIifoxlB2ahtkFFgRX2sq7h3rk8Zdh0ZK4nRg.6Qjrrtmr1CkL-O6W34A-hKLmTq2w98cvbYzz_1y6OyA",
		ProxyString:    proxy,
	}
	tmpt, err := tmptSolver.HandleTmpt()
	log.Println(tmpt, err)
}
