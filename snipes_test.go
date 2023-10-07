package browsergo_test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"

	browsergo "github.com/KakashiHatake324/browser-go"
)

const proxys = "http://aFsAq7B16-mushroom-b70bf8ba79!g-us!f-hddv!sid-gASmCJOgfOs:ad7k3jfj3kiuf3f@p2.mushroomproxy.com:8000"

func TestWalmart(t *testing.T) {
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
		StartUrl: "",
		//Proxy:        proxys,
		Args:         []browsergo.FlagType{},
		Headless:     false,
		OpenDevtools: false,
		WaitLoad:     false,
	}

	if err := instance.OpenBrowser(browserOpts); err != nil {
		log.Fatal(err)
	}

	defer instance.Close()

	close, err := instance.RequestListener()
	if err != nil {
		log.Fatal(err)
	}

	defer close()

	go instance.StartListener(func(intercept *browsergo.InterceptorCommunication) {
		if intercept.Request.Url == "https://collector-pxu6b0qd2s.px-cloud.net/assets/js/bundle" {
			//log.Println("IN THE LISTENER", intercept.Request)
		}
		if strings.Contains(intercept.Response.Url, "pxu6b0qd2s") {
			var res map[string]interface{}
			json.Unmarshal([]byte(intercept.Response.Body), &res)
			switch l := res["ob"].(type) {
			case string:
				decoded := DecodeResponse(l)
				log.Println("IN THE LISTENER", decoded)
			}
		}
	})

	if err := instance.Navigate("https://www.walmart.com/blocked", true); err != nil {
		log.Fatal(err)
	}

	if body, err := instance.GetBody(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(body)
	}

	time.Sleep(5 * time.Second)

	if err := instance.Extra(); err != nil {
		log.Fatal(err)
	}

	/*
		frame, err := instance.GetFrame("[title=\"creditCardIframeForm\"]")
		if err != nil {
			log.Fatal(err)
		}
	*/

	script := `return JSON.stringify({
		"_u6b0qd2Shandler": window._u6b0qd2Shandler,
		})`
	fullScript := fmt.Sprintf("(() => { %s })()", script)

	if resp, err := instance.Evaluate(fullScript); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
	time.Sleep(1 * time.Minute)
}

const siteHtml = `
<html lang="en">

<head>
    <title>Robot or human?</title>
    <meta name="viewport" content="width=device-width">
    <style>
    #sign-in-widget a,
    #sign-in-widget a:active,
    #sign-in-widget a:hover {
        color: #000
    }

    #sign-in-widget h1 {
        font-weight: 500;
        font-size: 20px;
        font-size: 1.25rem;
        letter-spacing: -.6px;
        margin: 1px auto
    }

    @media (min-width:30em) {
        #sign-in-widget h1 {
            margin-top: 24px;
            font-size: 24px;
            font-size: 1.5rem
        }
    }

    #sign-in-widget {
        font-family: BogleWeb, Helvetica Neue, Helvetica, Arial, sans-serif
    }

    #sign-in-widget * {
        box-sizing: border-box
    }

    #sign-in-widget .text-right {
        text-align: right
    }

    @font-face {
        font-family: NewYorkIcons;
        src: url(6255ed72d86ece856725a2d80878bce6.eot);
        font-weight: 400;
        font-style: normal
    }

    @font-face {
        font-family: NewYorkIcons;
        src: url(fd827841624904d4b8f51d20174fa3a4.woff2) format("woff2"), url(5b38b158833d0265af2b1c1093e489bd.woff) format("woff"), url(2ebae25dcb1bb39acbac9cffd8f10b15.ttf) format("truetype");
        font-weight: 400;
        font-style: normal
    }

    @keyframes spin-clockwise {
        0% {
            transform: rotate(0)
        }

        to {
            transform: rotate(1turn)
        }
    }

    @keyframes zoom-inverse {
        0% {
            opacity: 0;
            transform: scale(3)
        }

        33% {
            opacity: 1
        }

        to {
            transform: scale(1)
        }
    }

    @keyframes blinker {
        0% {
            opacity: 1
        }

        to {
            opacity: 0
        }
    }

    @keyframes slide-from-bottom {
        0% {
            transform: translateY(100%)
        }

        to {
            transform: translateY(0)
        }
    }

    @keyframes slide-from-top {
        0% {
            transform: translateY(0)
        }

        to {
            transform: translateY(100%)
        }
    }

    @keyframes fade-in-opacity {
        0% {
            opacity: 0
        }

        to {
            opacity: 1
        }
    }

    .elc-icon-spark:before {
        font-family: NewYorkIcons;
        display: inline;
        speak: none;
        content: "\E935"
    }

    .message {
        text-align: left
    }

    .message.active {
        display: block
    }

    .divider {
        margin: 24px 0 8px;
        text-align: center
    }

    @media (min-width:30em) {
        .divider {
            margin-top: 40px
        }
    }

    #sign-in-widget {
        text-align: center
    }

    #sign-in-widget p {
        margin-bottom: 0
    }

    #sign-in-widget .header-logo {
        text-decoration: none;
        font-size: 36px;
        font-size: 2.25rem
    }

    @media (min-width:30em) {
        #sign-in-widget .message.active+h1 {
            margin-top: 32px
        }
    }

    #sign-in-widget .elc-icon {
        font-family: NewYorkIcons
    }

    .u-sentenceCase {
        display: inline-block;
        text-decoration: inherit;
        text-transform: lowercase
    }

    .u-sentenceCase:first-letter {
        text-transform: uppercase
    }

    .u-sentenceCase--no-transform {
        text-transform: none
    }

    .message {
        -moz-osx-font-smoothing: grayscale;
        -webkit-font-smoothing: antialiased;
        display: none;
        padding: 8px 10px;
        border: 1px solid;
        color: #414042;
        font-size: 14px;
        font-size: .875rem
    }

    .message.active {
        display: inline-block
    }

    .spark {
        color: #ffc400
    }

    .sign-in-widget>* {
        display: none
    }

    [data-active="0"] .sign-in-widget>:first-child,
    [data-active="1"] .sign-in-widget>:nth-child(2),
    [data-active="10"] .sign-in-widget>:nth-child(11),
    [data-active="11"] .sign-in-widget>:nth-child(12),
    [data-active="12"] .sign-in-widget>:nth-child(13),
    [data-active="13"] .sign-in-widget>:nth-child(14),
    [data-active="14"] .sign-in-widget>:nth-child(15),
    [data-active="15"] .sign-in-widget>:nth-child(16),
    [data-active="16"] .sign-in-widget>:nth-child(17),
    [data-active="17"] .sign-in-widget>:nth-child(18),
    [data-active="18"] .sign-in-widget>:nth-child(19),
    [data-active="2"] .sign-in-widget>:nth-child(3),
    [data-active="3"] .sign-in-widget>:nth-child(4),
    [data-active="4"] .sign-in-widget>:nth-child(5),
    [data-active="5"] .sign-in-widget>:nth-child(6),
    [data-active="6"] .sign-in-widget>:nth-child(7),
    [data-active="7"] .sign-in-widget>:nth-child(8),
    [data-active="8"] .sign-in-widget>:nth-child(9),
    [data-active="9"] .sign-in-widget>:nth-child(10) {
        display: block
    }

    [data-active] {
        min-height: 100%;
        position: relative
    }

    [data-active]>* {
        margin-right: auto;
        margin-left: auto
    }

    @media (min-width:240px) {
        [data-active]>* {
            max-width: 304px;
            margin-right: auto;
            margin-left: auto
        }
    }

    @media (min-width:1024px) {
        [data-active]>* {
            max-width: 320px;
            margin-right: auto;
            margin-left: auto
        }
    }

    body,
    html {
        height: 100%;
        margin: 0
    }

    #sign-in-widget .header-logo {
        margin-top: 2px;
        display: inherit
    }

    @media (min-width:30em) {
        #sign-in-widget .header-logo {
            margin-top: 32px
        }
    }

    #sign-in-widget h1~.sign-in-widget {
        padding-bottom: 200px
    }

    .lite-footer {
        width: 100%;
        max-width: 100% !important
    }

    .lite-footer>hr {
        margin: 0;
        width: 100%;
        border-color: #e6e7e8;
        border-style: solid
    }

    @media (min-width:48em) {
        .lite-footer {
            position: absolute;
            bottom: 0;
            margin-top: 200px
        }
    }

    .lite-footer .main-container {
        padding: 16px;
        margin-bottom: 24px;
        font-size: 14px;
        font-size: .875rem;
        display: -ms-flexbox;
        display: flex;
        -ms-flex-direction: column;
        flex-direction: column;
        -ms-flex-pack: center;
        justify-content: center;
        -ms-flex-line-pack: center;
        align-content: center;
        width: 100%;
        max-width: 100% !important
    }

    @media (min-width:64em) {
        .lite-footer .main-container {
            -ms-flex-direction: row;
            flex-direction: row;
            -ms-flex-pack: justify;
            justify-content: space-between;
            -ms-flex-line-pack: start;
            align-content: flex-start
        }
    }

    .lite-footer .main-container .left-section {
        text-align: center
    }

    .lite-footer .main-container .left-section>a {
        display: block;
        margin-bottom: 16px
    }

    @media (min-width:48em) {
        .lite-footer .main-container .left-section {
            text-align: left
        }

        .lite-footer .main-container .left-section>a {
            display: inline;
            margin-right: 16px
        }
    }

    .lite-footer .main-container .right-section {
        text-align: center;
        font-size: 12px;
        font-size: .75rem
    }

    .lite-footer .main-container .right-section>p {
        margin: 0
    }

    @media (min-width:48em) {
        .lite-footer .main-container .right-section {
            text-align: left
        }

        .lite-footer .main-container .right-section>p {
            margin-top: 16px
        }
    }

    @media (min-width:64em) {
        .lite-footer .main-container .right-section>p {
            margin: 0
        }
    }

    /* upload and replace urls with new CDN URL*/
    @font-face {
        font-family: NewYorkIcons;
        src: url(https://i5.wal.co/dfw/63fd9f59-161f/f0d7612f-54fb-4910-a9c7-20c65d03f878/v1/6255ed72d86ece856725a2d80878bce6.eot);
        font-weight: 400;
        font-style: normal
    }

    @font-face {
        font-family: NewYorkIcons;
        src: url(https://i5.wal.co/dfw/63fd9f59-161f/f0d7612f-54fb-4910-a9c7-20c65d03f878/v1/fd827841624904d4b8f51d20174fa3a4.woff2) format("woff2"), url(https://i5.wal.co/dfw/63fd9f59-161f/f0d7612f-54fb-4910-a9c7-20c65d03f878/v1/5b38b158833d0265af2b1c1093e489bd.woff) format("woff"), url(https://i5.wal.co/dfw/63fd9f59-161f/f0d7612f-54fb-4910-a9c7-20c65d03f878/v1/2ebae25dcb1bb39acbac9cffd8f10b15.ttf) format("truetype");
        font-weight: 400;
        font-style: normal
    }

    @font-face {
        font-family: "BogleWeb";
        src: url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Bold.eot");
        src: url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Bold.eot?#iefix") format("embedded-opentype"),
            url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Bold.woff2") format("woff2"),
            url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Bold.woff") format("woff"),
            url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Bold.ttf") format("truetype");
        font-weight: 700;
        font-style: normal;
    }

    @font-face {
        font-family: "BogleWeb";
        src: url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Bold.eot");
        src: url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Bold.eot?#iefix") format("embedded-opentype"),
            url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Bold.woff2") format("woff2"),
            url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Bold.woff") format("woff"),
            url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Bold.ttf") format("truetype");
        font-weight: 600;
        font-style: normal;
    }

    @font-face {
        font-family: "BogleWeb";
        src: url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Regular.eot");
        src: url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Regular.eot?#iefix") format("embedded-opentype"),
            url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Regular.woff2") format("woff2"),
            url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Regular.woff") format("woff"),
            url("https://i5.walmartimages.com/dfw/63fd9f59-a78c/fcfae9b6-2f69-4f89-beed-f0eeb4237946/v1/BogleWeb_subset-Regular.ttf") format("truetype");
        font-weight: 400;
        font-style: normal;
    }
    </style>
    <script>
    function getUrlVars() {
        var vars = {};
        var parts = window.location.href.replace(/[?&]+([^=&]+)=([^&]*)/gi, function(m, key, value) {
            vars[key] = value;
        });
        return vars;
    }

    function getUrlParam(parameter, defaultvalue) {
        var urlparameter = defaultvalue;
        if (window.location.href.indexOf(parameter) > -1) {
            urlparameter = getUrlVars()[parameter];
        }
        return urlparameter;
    }

    </script>
</head>

<body>
    <div>
        <div data-active="0" id="sign-in-widget">
            <a class="header-logo" href="/" aria-atomic="true" aria-label="Walmart. Save Money. Live Better. Home Page">
                <span class="spark elc-icon elc-icon-spark elc-icon-36"></span>
            </a>
            <h1 class="heading heading-d sign-in-widget">
                Robot or human?
            </h1>
            <div class="sign-in-widget">
                <div class="re-captcha">
                    <p class="bot-message" id=message>Activate and hold the button to confirm that you’re human. Thank You!</p>
                    <div id="px-captcha" style="margin:16px;margin-bottom: 32px; margin-top: 32px; align-content:center; align-items: center;"></div>
                </div>
            </div>
            <script>
            window._pxAppId = 'PXu6b0qd2S';
            window._pxJsClientSrc = '/px/' + window._pxAppId + '/init.js';
            window._pxFirstPartyEnabled = true;
            window._pxHostUrl = '/px/' + window._pxAppId + '/xhr';
            window._pxreCaptchaTheme = 'light';
            window._PXETnJ2Y5H = {
                challenge: {
                    view: {
                         textFont: "BogleWeb, Helvetica Neue, Helvetica, Arial, sans-serif"
                   }
                }
            };

            var hc = getUrlParam('g', 'b');
            var alt = hc
            if (alt=='a') {
                document.getElementById('message').innerHTML = '<p>Check the box to confirm that you’re human. Thank You!</p>';
            }
            var captchajs = "/px/" + window._pxAppId + "/captcha/captcha.js?a=c&m=0&g=" + hc
            </script>
            <script id="blockScript"></script>
            <script>
            document.getElementById('blockScript').src = captchajs;
            </script>
            <div class="lite-footer">
                <hr aria-hidden="true" class="divider">
                <div class="main-container">
                    <div class="left-section">
                        <a class="" href="https://help.walmart.com/app/answers/detail/a_id/8">
                            <span class="u-sentenceCase u-sentenceCase--no-transform">Terms of Use</span>
                        </a>
                        <a class="" href="https://corporate.walmart.com/privacy-security">
                            <span class="u-sentenceCase u-sentenceCase--no-transform">Privacy Policy</span>
                        </a>
                        <a class="" href="/account/api/ccpa-intake?native=false&amp;type=sod">
                            <span class="u-sentenceCase u-sentenceCase--no-transform">Do Not Sell My Personal Information</span>
                        </a>
                        <a class="" href="/account/api/ccpa-intake?native=false&amp;type=access">
                            <span class="u-sentenceCase u-sentenceCase--no-transform">Request My Personal Information</span>
                        </a>
                    </div>
                    <div class="right-section">
                        <p class="copy-base-ny">©2023 Walmart Stores, Inc.</p>
                    </div>
                </div>
            </div>
        </div>
</body>

</html>
`

func DecodeResponse(data string) []string {
	return decodeJoe(data, "98")
}

func decodeJoe(t string, n string) []string {
	a := base64Decode(t)
	i, _ := strconv.Atoi(extractDigits(n))
	return strings.Split(xorCipher(a, i%128), "~~~~")
}

func base64Decode(t string) string {
	bytes, _ := base64.StdEncoding.DecodeString(t)

	return string(bytes)
}

func extractDigits(t string) string {
	digits := ""
	for i := 0; i < len(t); i++ {
		charCode := t[i]
		if charCode >= '0' && charCode <= '9' {
			digits += string(charCode)
		}
	}
	return digits
}

func xorCipher(t string, n int) string {
	ciphered := ""
	for _, char := range t {
		ciphered += string(rune(n ^ int(char)))
	}
	return ciphered
}
