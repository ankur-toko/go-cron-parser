package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func IntToString() func(int) (string, error) {
	return func(a int) (string, error) {
		return strconv.Itoa(a), nil
	}
}

func StringToInt() func(string) (int, error) {
	return strconv.Atoi
}

func StringToIntMonth() func(string) (int, error) {
	m := map[string]int{}
	m["jan"] = 1
	m["feb"] = 2
	m["mar"] = 3
	m["apr"] = 4
	m["may"] = 5
	m["jun"] = 6
	m["jul"] = 7
	m["aug"] = 8
	m["sep"] = 9
	m["oct"] = 10
	m["nov"] = 11
	m["dec"] = 12

	return func(str string) (int, error) {
		v, e := strconv.Atoi(str)
		if e == nil {
			return v, nil
		}
		v, ok := m[strings.ToLower(str[0:3])]
		if !ok {
			return -1, fmt.Errorf("unknown month")
		} else {
			return v, nil
		}
	}
}

func StringToIntDay() func(string) (int, error) {
	m := map[string]int{}
	m["mon"] = 0
	m["tue"] = 1
	m["wed"] = 2
	m["thu"] = 3
	m["fri"] = 4
	m["sat"] = 5
	m["sun"] = 6

	return func(str string) (int, error) {
		v, e := strconv.Atoi(str)
		if e == nil {
			return v, nil
		}
		v, ok := m[strings.ToLower(str[0:3])]
		if !ok {
			return -1, fmt.Errorf("unknown month")
		} else {
			return v, nil
		}
	}
}

func IntToStringMonth() func(int) (string, error) {
	m := map[int]string{}
	m[1] = "jan"
	m[2] = "feb"
	m[3] = "mar"
	m[4] = "apr"
	m[5] = "may"
	m[6] = "jun"
	m[7] = "jul"
	m[8] = "aug"
	m[9] = "sep"
	m[10] = "oct"
	m[11] = "nov"
	m[12] = "dec"

	return func(i int) (string, error) {
		return m[i], nil
	}
}

func IntToStringDay() func(int) (string, error) {
	m := map[int]string{}

	m[0] = "mon"
	m[1] = "tue"
	m[2] = "wed"
	m[3] = "thu"
	m[4] = "fri"
	m[5] = "sat"
	m[6] = "sun"

	return func(i int) (string, error) {
		return m[i], nil
	}
}
