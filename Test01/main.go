package main

import (
	"fmt"

	coraza "github.com/jptosso/coraza-waf"
	"github.com/jptosso/coraza-waf/seclang"
)

func initCoraza() {

	coraza.NewWaf()
}

func parseRules(waf *coraza.Waf) {
	parser, _ := seclang.NewParser(waf)
	parser.FromString(`SecAction "id:1,phase:1,deny:403"`)
}

func main() {
	initCoraza()
	fmt.Println("Hello there.!")
}
