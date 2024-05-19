package main

import (
	"reflect"
	"testing"

	"github.com/ankur-toko/go-cron-parser/parser"
)

/*
Parse(str string) (CronExpression, error)
*/
type MockParser struct{}

func (m MockParser) Parse(str string) (parser.CronExpression, error) {
	return parser.CronExpression{[]int{0, 15, 30, 45}, []int{0}, []int{1, 15}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, []int{1, 2, 3, 4, 5}}, nil
}

func TestNewCronCommand(t *testing.T) {

	mp := MockParser{}
	cexp, _ := mp.Parse("dummy")

	tests := []struct {
		name    string
		parser  parser.ICronParser
		input   string
		want    CronCommand
		wantErr bool
	}{
		{"command identification", mp, "abc asd asd asd ads program", CronCommand{cexp, "program"}, false},
		{"command identification", mp, "abc asd asd asd ads program with multiple words", CronCommand{cexp, "program with multiple words"}, false},
		{"less than required fields", mp, "abc asd asd", CronCommand{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCronCommand(tt.parser, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCronCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCronCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
