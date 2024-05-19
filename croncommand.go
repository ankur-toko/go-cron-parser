package main

import (
	"fmt"
	"strings"

	"github.com/ankur-toko/go-cron-parser/parser"
)

const INVALID_CRON_WITH_COMMAND string = "error: invalid CRON expression. expected 6 fields"

var paddingLen int = 14

type CronCommand struct {
	cronExp parser.CronExpression
	command string
}

func (c CronCommand) Print() {
	c.cronExp.Print(paddingLen)
	fmt.Printf("%-*v%v\n", paddingLen, "command", c.command)
}

func NewCronCommand(parser parser.ICronParser, str string) (CronCommand, error) {
	input := strings.Split(str, " ")
	if len(input) < 6 {
		return CronCommand{}, fmt.Errorf(INVALID_CRON_WITH_COMMAND)
	}
	exp, err := parser.Parse(strings.Join(input[0:5], " "))
	if err != nil {
		return CronCommand{}, err
	}
	return CronCommand{exp, strings.Join(input[5:], " ")}, nil

}
