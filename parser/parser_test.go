package parser

import (
	"reflect"
	"testing"
)

func TestCronParser_Parse(t *testing.T) {
	// Only tests the standard parser

	/*
		minute 0 15 30 45
		hour 0
		day of month 1 15
		month 1 2 3 4 5 6 7 8 9 10 11 12
		day of week 1 2 3 4 5
		command /usr/bin/find

	*/

	parser := New()
	tests := []struct {
		name    string
		input   string
		want    CronExpression
		wantErr bool
	}{
		{"basic", "*/15 0 1,15 * 1-5", CronExpression{[]int{0, 15, 30, 45}, []int{0}, []int{1, 15}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, []int{1, 2, 3, 4, 5}}, false},
		{"basic 2", "*/20 0 1,15 * 1-5", CronExpression{[]int{0, 20, 40}, []int{0}, []int{1, 15}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, []int{1, 2, 3, 4, 5}}, false},
		{"sun-mon case", "*/20 0 1,15 * mon-sun", CronExpression{[]int{0, 20, 40}, []int{0}, []int{1, 15}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, []int{0, 1, 2, 3, 4, 5, 6}}, false},

		{"basic error case", "*/15 0 1,15 *", CronExpression{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CronParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CronParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
