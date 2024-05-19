package parser

import "fmt"

type CronExpression struct {
	Minute []int
	Hour   []int
	Day    []int
	Month  []int
	DOW    []int
}

func (c CronExpression) Print(padding int) {
	fmt.Printf("%-*v%v\n", padding, "minutes", c.Minute)
	fmt.Printf("%-*v%v\n", padding, "hours", c.Hour)
	fmt.Printf("%-*v%v\n", padding, "days", c.Day)
	fmt.Printf("%-*v%v\n", padding, "month", c.Month)
	fmt.Printf("%-*v%v\n", padding, "day of week", c.DOW)
}
