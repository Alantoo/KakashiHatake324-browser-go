package browsergo_test

import (
	"context"
	"log"
	"testing"

	browsergo "github.com/KakashiHatake324/browser-go"
)

const tmptProxy = "31.128.120.186:63690:yvyyhlia:s4XI854jtl"

func TestTmpt(t *testing.T) {
	service, err := browsergo.InitBrowser("browser-go-test", true, true, "")
	if err != nil {
		log.Fatal(err)
	}
	instance, err := service.NewService(context.Background(), 400000)
	if err != nil {
		log.Fatal(err)
	}
	defer instance.Close()

	proxy, _ := browsergo.DeFormatProxy(tmptProxy)
	tmptSolver := &browsergo.SolveTmpt{
		BrowserService: instance,
		Context:        context.TODO(),
		Deadline:       11,
		Url:            "https://identity.ticketmaster.com/sign-in?integratorId=prd1741.iccp&placementId=mytmlogin&redirectUri=https://www.ticketmaster.com/",
		ProxyString:    proxy,
		ForceChrome:    true,
	}
	tmpt, err := tmptSolver.HandleTmpt()
	log.Println(tmpt, err)
}
