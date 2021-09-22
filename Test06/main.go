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

	tx := waf.NewTransaction()
	defer tx.ProcessLogging()

	fmt.Println("goodbye..!!")
}
