package test

import (
	"algorithm022/Week_02/demo1"
	"algorithm022/Week_02/demo2"
	"algorithm022/Week_02/demo4"
	"algorithm022/Week_02/demo8"
	"fmt"
	"testing"
)

func TestDemo1(t *testing.T) {
	ss := "anagram"
	tt := "nagaram"
	res := demo1.Run(ss, tt)
	fmt.Println(res)
}

func TestDemo2(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	var res []int
	res = demo2.Run(nums, target)
	fmt.Println(res)

	res = demo2.Run2(nums, target)
	fmt.Println(res)
}

func TestDemo4(t *testing.T) {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	demo4.Run(a)
}

func TestDemo8(t *testing.T) {
	res := demo8.Run(1)
	fmt.Println(res)
}
