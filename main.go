package main

import (
	"fmt"
	"os"

	"github.com/ankur-toko/go-cron-parser/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`cron-parser: expression is missing`)
		return
	}
	cparser := parser.New()
	cronCommand, err := NewCronCommand(cparser, os.Args[1])
	if err == nil {
		cronCommand.Print()
	} else {
		fmt.Print(err)
	}
}
