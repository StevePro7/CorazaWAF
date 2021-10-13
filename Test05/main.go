package main

import (
	"fmt"
	"log"

	"github.com/jptosso/coraza-waf"
	"github.com/jptosso/coraza-waf/seclang"
)

func initCoraza() *coraza.Waf {

	return coraza.NewWaf()
}

func parseRules(waf *coraza.Waf) {
	parser, _ := seclang.NewParser(waf)
	err := parser.FromString(`SecAction "id:1,phase:1,deny:403"`)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	waf := initCoraza()
	parseRules(waf)
	fmt.Println("Hello")

	// Handling an interruption
	tx := waf.NewTransaction()

	// Parse request
	tx.AddResponseHeader("some", "header")
	// if it := tx.ProcessResponseHeaders(200, "http"); it != nil {
	// 	//return processInterruption(it)
	// }

	// if !tx.IsProcessableResponseBody() {
	// 	// We stream the response to the client
	// 	// sw.WriteStatus(200)
	// 	// sw.Write(res.Body)
	// 	// sw.Close()
	// }

	// Add response data from a string or bytes
	_, err := tx.ResponseBodyBuffer.Write([]byte("some response data"))
	if err != nil {
		log.Fatal(err)
		return
	}
	// Or dump a Response Body buffers into Coraza
	//io.Copy(tx.ResponseBodyBuffer, res.Body)

	// sw.WriteStatus(200)
	// sw.Write(tx.ResponseBodyBuffer.Reader())
	// sw.Close()

	fmt.Println("goodbye..!")
}
