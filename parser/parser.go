package parser

import (
	"fmt"
	"strings"
)

const CRON_SEGMENTS int = 5

type CronParser struct {
	Rangers [CRON_SEGMENTS]Ranger
}

type ICronParser interface {
	Parse(str string) (CronExpression, error)
}

func New() ICronParser {
	return CronParser{
		[CRON_SEGMENTS]Ranger{
			{0, 59, StringToInt(), IntToString()},
			{0, 23, StringToInt(), IntToString()},
			{1, 31, StringToInt(), IntToString()},
			{1, 12, StringToIntMonth(), IntToStringMonth()},
			{0, 6, StringToIntDay(), IntToStringDay()},
		},
	}
}

func (parser CronParser) Parse(str string) (CronExpression, error) {
	parts := strings.Split(str, " ")
	if len(parts) != CRON_SEGMENTS {
		return CronExpression{}, fmt.Errorf(INVALID_CRON)
	}
	res := [][]int{}
	for i := 0; i < CRON_SEGMENTS; i++ {
		r, e := parser.Rangers[i].Parse(parts[i])
		if e != nil {
			return CronExpression{}, e
		}
		res = append(res, r)
	}

	cronExp := CronExpression{
		res[0],
		res[1],
		res[2],
		res[3],
		res[4],
	}

	return cronExp, nil
}
