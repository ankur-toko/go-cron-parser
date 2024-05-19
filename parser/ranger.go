package parser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Ranger struct {
	Min         int
	Max         int
	StrToInt    func(string) (int, error)
	IntToString func(int) (string, error)
}

func (r Ranger) Parse(str string) ([]int, error) {
	if strings.Contains(str, ",") {
		parts := strings.Split(str, ",")
		res := []int{}
		for i := 0; i < len(parts); i++ {
			subrange, e := r.Parse(parts[i])
			if e != nil {
				return nil, e
			}
			res = append(res, subrange...)
		}
		return removeDupl(res), nil
	} else if strings.Contains(str, "/") {
		parts := strings.Split(str, "/")
		if len(parts) != 2 {
			return nil, fmt.Errorf(INCORRECT_FORMAT)
		}
		stepUp, e := strconv.Atoi(parts[1])
		if e != nil {
			return nil, e
		}
		if stepUp < 1 {
			return nil, fmt.Errorf(INCORRECT_FORMAT)
		}
		arr, e := r.Parse(parts[0])
		if e != nil {
			return nil, e
		}
		res := []int{}
		for i := 0; i < len(arr); i += stepUp {
			res = append(res, arr[i])
		}
		return res, nil
	} else {
		return r.fanOut(str)
	}
}

func (r Ranger) fanOut(str string) ([]int, error) {
	// processes cron parts not containing `,` and `/`
	if strings.Contains(str, "*") {
		return r.fanOutStar(str)
	} else if strings.Contains(str, "-") {
		return r.fanOutHypen(str)
	}
	num, e := r.StrToInt(str)
	if e != nil || !r.inRange(num) {
		return nil, fmt.Errorf(INVALID_CRON)
	}
	return []int{num}, nil
}

func (r Ranger) fanOutStar(str string) ([]int, error) {
	if str != "*" {
		return nil, fmt.Errorf(INCORRECT_FORMAT)
	}
	return genRange(r.Min, r.Max), nil
}

func (r Ranger) fanOutHypen(str string) ([]int, error) {
	if parts := strings.Split(str, "-"); len(parts) == 2 {
		s, e1 := r.StrToInt(parts[0])
		e, e2 := r.StrToInt(parts[1])
		if e1 != nil || e2 != nil {
			return nil, fmt.Errorf(INCORRECT_FORMAT)
		}
		if !r.inRange(s) || !r.inRange(e) || s > e {
			return nil, fmt.Errorf(INCORRECT_FORMAT)
		}
		return genRange(s, e), nil
	} else {
		return nil, fmt.Errorf(INCORRECT_FORMAT)
	}
}

func genRange(s, e int) []int {
	res := []int{}
	for i := s; i <= e; i++ {
		res = append(res, i)
	}
	return res
}

func (r Ranger) inRange(num int) bool {
	if num >= r.Min && num <= r.Max {
		return true
	}
	return false
}

func removeDupl(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	sort.Ints(arr)
	res := []int{arr[0]}

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			res = append(res, arr[i])
		}
	}
	return res
}
