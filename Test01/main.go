package main

import (
	"fmt"
	"github.com/jptosso/coraza-waf/seclang"
	"log"

	"github.com/jptosso/coraza-waf"
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
	fmt.Println("Hello there.!")
}
