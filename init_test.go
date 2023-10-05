package crigo_test

import (
	"context"
	"log"
	"testing"
	"time"

	crigo "github.com/graph-labs-io/cri-go"
)

func TestInit(t *testing.T) {
	t.Skip()
	service, err := crigo.InitCRI(true, "")
	if err != nil {
		log.Fatal(err)
	}

	defer service.CloseClient()

	instance, err := service.NewService(context.Background(), 220000)
	if err != nil {
		log.Fatal(err)
	}

	defer instance.Close()

	browserOpts := &crigo.BrowserOpts{
		StartUrl:     "https://www.google.com",
		Proxy:        "",
		Args:         []crigo.FlagType{crigo.RandomWindowSize()},
		Headless:     false,
		OpenDevtools: false,
		WaitLoad:     false,
	}

	if err := instance.OpenBrowser(browserOpts); err != nil {
		log.Fatal(err)
	}

	var cookies []*crigo.CRIGoCookies
	cookies = append(cookies, &crigo.CRIGoCookies{
		Name:   "hi",
		Value:  "cookie",
		Domain: ".nike.com",
	})

	if err := instance.SetCookies(cookies); err != nil {
		log.Fatal(err)
	}

	if err := instance.Navigate("https://www.nike.com", true); err != nil {
		log.Fatal(err)
	}

	if err := instance.WaitForElement("[data-var=\"userName\"]"); err != nil {
		log.Fatal(err)
	}

	username, err := instance.Evaluate("document.querySelector('[data-var=\"userName\"]')?.innerHTML")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(username)

	time.Sleep(30 * time.Second)

	if err := instance.Close(); err != nil {
		log.Fatal(err)
	}
}
