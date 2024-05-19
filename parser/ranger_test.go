package parser

import (
	"reflect"
	"strconv"
	"testing"
)

func TestRanger_Parse(t *testing.T) {
	r := Ranger{2, 9, strconv.Atoi, func(i int) (string, error) { return strconv.Itoa(i), nil }}
	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{"* case", "*", []int{2, 3, 4, 5, 6, 7, 8, 9}, false},
		{"- case", "2-5", []int{2, 3, 4, 5}, false},
		{"* with /", "*/1", []int{2, 3, 4, 5, 6, 7, 8, 9}, false},
		{"* with /2", "*/2", []int{2, 4, 6, 8}, false},
		{"* with /3", "*/3", []int{2, 5, 8}, false},
		{"* with /30", "*/30", []int{2}, false},
		{"* with /-1", "*/-1", nil, true},

		{"- with /", "2-5/1", []int{2, 3, 4, 5}, false},
		{"- with /2", "2-5/2", []int{2, 4}, false},
		{"- with /3", "2-5/3", []int{2, 5}, false},
		{"- with /30", "2-5/30", []int{2}, false},
		{"- with /-1", "2-5/-1", nil, true},

		{"* with ,", "*,*", []int{2, 3, 4, 5, 6, 7, 8, 9}, false},
		{"* and - with ,", "2-5,*", []int{2, 3, 4, 5, 6, 7, 8, 9}, false},
		{"- with , overlapping", "2-5,3-6", []int{2, 3, 4, 5, 6}, false},
		{"- with , and no overlapping", "2-4,7-9", []int{2, 3, 4, 7, 8, 9}, false},
		{"- with , and no overlapping and sort order test", "7-9,2-4", []int{2, 3, 4, 7, 8, 9}, false},

		{"* with , and /", "*,*/2", []int{2, 3, 4, 5, 6, 7, 8, 9}, false},
		{"* and - with , /", "2-5/2,*", []int{2, 3, 4, 5, 6, 7, 8, 9}, false},
		{"- with , overlapping and /2", "2-5/2,3-6", []int{2, 3, 4, 5, 6}, false},
		{"- with , and no overlapping", "2-4/2,7-9/2", []int{2, 4, 7, 9}, false},

		{"incorrect division case", "2-5/1-2", nil, true},
		{"incorrect division case", "*/1-2", nil, true},
		{"incorrect range format", "asf/1", nil, true},
		{"division case", "*/1-2", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ranger.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ranger.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDupl(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"empty case",
			args{
				[]int{},
			},
			[]int{},
		},
		{
			"single element",
			args{
				[]int{1},
			},
			[]int{1},
		},
		{
			"case1: with dupl",
			args{
				[]int{1, 1},
			},
			[]int{1},
		}, {
			"case2: with some dupl",
			args{
				[]int{1, 2, 1},
			},
			[]int{1, 2},
		}, {
			"case3: with no dupl",
			args{
				[]int{1, 2, 3},
			},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDupl(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDupl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRanger_inRange(t *testing.T) {
	r := Ranger{
		5, 10, StringToInt(), IntToString(),
	}

	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			"out of range",
			1,
			false,
		},
		{
			"within range",
			6,
			true,
		},
		{
			"out of range right side",
			11,
			false,
		},
		{
			"border 1",
			5,
			true,
		},
		{
			"border 2",
			10,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := r.inRange(tt.input); got != tt.want {
				t.Errorf("Ranger.inRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genRange(t *testing.T) {
	tests := []struct {
		name string
		l    int
		r    int
		want []int
	}{
		{"case 1", 1, 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"case 2", 8, 10, []int{8, 9, 10}},
		{"case 3", 8, 8, []int{8}},
		{"case 4", 10, 8, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genRange(tt.l, tt.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRanger_fanOutStar(t *testing.T) {
	r := Ranger{5, 10, StringToInt(), IntToString()}

	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{"base case", "*", []int{5, 6, 7, 8, 9, 10}, false},
		{"false case", "*1", nil, true},
		{"false case", "-", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.fanOutStar(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ranger.fanOutStar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ranger.fanOutStar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRanger_fanOut(t *testing.T) {
	r := Ranger{2, 5, strconv.Atoi, func(i int) (string, error) { return strconv.Itoa(i), nil }}

	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{"* case", "*", []int{2, 3, 4, 5}, false},
		{"- case", "3-4", []int{3, 4}, false},
		{"- case", "2-5", []int{2, 3, 4, 5}, false},
		{"- case incorrect range", "3-10", nil, true},
		{"- case incorrect range", "1-5", nil, true},
		{"invalid format", "-1*5", nil, true},
		{"invalid format", "-*", nil, true},
		{"invalid format", "1-*", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.fanOut(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ranger.fanOut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ranger.fanOut() = %v, want %v", got, tt.want)
			}
		})
	}
}
