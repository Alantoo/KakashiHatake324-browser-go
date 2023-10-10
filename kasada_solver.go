package browsergo

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type SolveKasada struct {
	*ClientInit
	// seconds to stop the deadline
	Context     context.Context
	Deadline    int
	UserAgent   string
	RequestUrl  string
	KpsdkST     int64
	ProxyString string
	XKpsdkCt    string
	XKpsdkCd    string
	XKpsdkV     string
}

// solve shape with a shape request
func (c *SolveKasada) HandleKasada() error {
	var err error

	if c.KpsdkST != 0 {
		cd := NewCDGenerator(c.KpsdkST)
		cd.baseCDGen()
		c.XKpsdkCd = cd.generateJSON()
		return nil
	}

	c.XKpsdkCd = ""
	c.XKpsdkCt = ""

	if strings.Contains(c.ProxyString, "http://") {
		if c.ProxyString, err = DeFormatProxy(c.ProxyString); err != nil {
			return err
		}
	}
	ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(time.Duration(c.Deadline)*time.Second))
	defer cancel()

	instance, err := c.NewService(ctx, 220000)
	if err != nil {
		return err
	}

	browserOpts := &BrowserOpts{
		StartUrl: "",
		Proxy:    c.ProxyString,
		Args: []FlagType{
			SetUserAgent(c.UserAgent),
		},
		Headless:     true,
		OpenDevtools: false,
		WaitLoad:     false,
	}

	if err := instance.OpenBrowser(browserOpts); err != nil {
		return err
	}

	close, err := instance.RequestListener()
	if err != nil {
		log.Fatal(err)
	}

	defer close()

	defer instance.Close()

	go instance.StartListener(func(intercept *InterceptorCommunication) {
		select {
		case <-c.Context.Done():
			err = errors.New("main context was cancelled")
			return
		case <-c.CTX.Done():
			err = errors.New("main context was cancelled")
			return
		case <-ctx.Done():
			err = errors.New("deadline has exceeded so the context cancelled")
			return
		default:
			if _, ok := intercept.Request.Headers["x-kpsdk-v"]; ok {
				c.XKpsdkV = intercept.Request.Headers["x-kpsdk-v"].(string)
			}

			if _, ok := intercept.Response.Headers["x-kpsdk-st"]; ok && intercept.Response.Status == 429 {
				err = errors.New("blocked")
			} else {
				if _, ok := intercept.Response.Headers["x-kpsdk-st"]; ok {
					kpst, _ := strconv.Atoi(intercept.Response.Headers["x-kpsdk-st"].(string))
					c.KpsdkST = int64(kpst)
					c.XKpsdkCt = intercept.Response.Headers["x-kpsdk-ct"].(string)
				}
			}
		}
	})

	if err := instance.Navigate(c.RequestUrl, false); err != nil {
		return err
	}

	for c.KpsdkST == 0 && err == nil {
		select {
		case <-c.Context.Done():
			err = errors.New("main context was cancelled")
			return err
		case <-c.CTX.Done():
			err = errors.New("main context was cancelled")
			return err
		case <-ctx.Done():
			err = errors.New("deadline has exceeded so the context cancelled")
			return err
		default:
			body, errs := instance.GetBody()
			if errs != nil {
				err = errs
			}
			if strings.Contains(body, "The Chromium Authors") {
				err = errors.New("blocked")
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}
	if err != nil {
		log.Println("error genertating kasada headers")
		return err
	}
	cd := NewCDGenerator(c.KpsdkST)
	cd.baseCDGen()
	c.XKpsdkCd = cd.generateJSON()
	return nil
}

type CDGenerator struct {
	ID       string
	Answers  []int
	ST       int64
	WorkTime int64
	RST      int64
	Delta    int
}

func NewCDGenerator(kpsdkST int64) *CDGenerator {
	cd := &CDGenerator{
		ST:       kpsdkST,
		WorkTime: time.Now().UnixNano() / int64(time.Millisecond),
	}
	cd.ID = cd.generateUUID()
	delta := RandomInt(50, 150)
	cd.RST = cd.ST + int64(delta)
	cd.Delta = (delta - 1) * -1
	return cd
}

func (cd *CDGenerator) generateUUID() string {
	uuid := make([]byte, 16)
	_, _ = rand.Read(uuid)
	return fmt.Sprintf("%x", uuid)
}

func (cd *CDGenerator) generateJSON() string {
	result := map[string]interface{}{
		"workTime": cd.WorkTime,
		"id":       cd.ID,
		"answers":  cd.Answers,
		"d":        cd.Delta,
		"rst":      cd.RST,
		"st":       cd.ST,
		"duration": fmt.Sprintf("%.1f", math.Floor(1)*9+1),
	}
	jsonData, _ := json.Marshal(result)
	return string(jsonData)
}

func (cd *CDGenerator) byteToHex(bArr []byte) string {
	cArr := "0123456789abcdef"
	var cArr2 string
	for _, b2 := range bArr {
		cArr2 += string(cArr[(b2&240)>>4]) + string(cArr[b2&15])
	}
	return cArr2
}

func (cd *CDGenerator) checkChallenge(strVal string) float64 {
	var b float64
	for _, digit := range strVal[:13] {
		b = b*16 + float64(digit) - 48
	}
	j := b + 1
	return 4.503599627370496e15 / j
}

func (cd *CDGenerator) SHA256Digest(message string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(message))
	return hasher.Sum(nil)
}

func (cd *CDGenerator) baseCDGen() {
	cd.generateCD("tp-v2-input", 10, 2)
}

func (cd *CDGenerator) generateCD(message string, i, i2 int) {
	d := float64(i) / float64(i2)
	digest2 := cd.SHA256Digest(fmt.Sprintf("%s, %d, %s", message, cd.WorkTime, cd.ID))
	i3 := 0
	for i3 < i2 {
		i4 := 1
		for {
			digest := cd.SHA256Digest(fmt.Sprintf("%d, %s", i4, cd.byteToHex(digest2)))
			e := cd.byteToHex(digest)
			if cd.checkChallenge(e) >= d {
				break
			}
			i4++
		}
		cd.Answers = append(cd.Answers, i4)
		i3++
		digest2 = cd.SHA256Digest(fmt.Sprintf("%d, %s", i4, cd.byteToHex(digest2)))
	}
}
