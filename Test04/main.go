package main

import (
	"fmt"

	coraza "github.com/jptosso/coraza-waf"
	"github.com/jptosso/coraza-waf/seclang"
)

func initCoraza() *coraza.Waf {

	return coraza.NewWaf()
}

func parseRules(waf *coraza.Waf) {
	parser, _ := seclang.NewParser(waf)
	parser.FromString(`SecAction "id:1,phase:1,deny:403"`)
}

func main() {
	waf := initCoraza()
	parseRules(waf)
	fmt.Println("Hello")

	// Handling an interruption
	tx := waf.NewTransaction()

	// 127.0.0.1:55555 -> 127.0.0.1:80
	tx.ProcessConnection("127.0.0.1", 55555, "127.0.0.1", 80)

	// Request URI waws /some-url?with=args
	tx.ProcessUri("/some-url?with=args", "GET", "1.1")

	// We add some headers
	tx.AddRequestHeader("Host", "somehost.com")
	tx.AddRequestHeader("Cookie", "some-cookie=with-value")
	// Content-Type is important to tell coraza which BodyProcessor must be used
	tx.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")

	// We process phase 1 (Request)
	if it := tx.ProcessRequestHeaders(); it != nil {
		//return processInterruption(it)
	}

	// We add urlencoded POST data
	tx.RequestBodyBuffer.Write([]byte("somepost=data&with=parameters"))
	// We process phase 2 (Request Body)
	if it, _ := tx.ProcessRequestBody(); it != nil {
		//return processInterruption(it)
	}

	fmt.Println("goodbye..")
}
