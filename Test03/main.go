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
	fmt.Println("Hello there.!..!!")

	// Handling an interruption
	tx := waf.NewTransaction()
	// Add some variables and process some phases
	if it := tx.Interruption; it != nil {
		switch it.Action {
		case "deny":
			//rw.WriteStatus(it.Status)
			//rw.Write([]byte("some error message"))
		}
	}

}
