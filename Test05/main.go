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

	// Add response data from a strint or bytes
	tx.ResponseBodyBuffer.Write([]byte("some response data"))
	// Or dump a REsponse Body buffer into Coraza
	//io.Copy(tx.ResponseBodyBuffer, res.Body)

	// sw.WriteStatus(200)
	// sw.Write(tx.REsponseBodyBuffer.Reader())
	// sw.Close()

	fmt.Println("goodbye..!")
}
