package main

import (
	"fmt"
	"golang-algorithm/arr"
	"golang-algorithm/leetcode"
)

func main() {
	testTwoSum()
}

func testTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 19
	fmt.Println(leetcode.TwoSum(nums, target))
}

func testFindSubArrayList() {
	arrays := []int{1, 2, 6, 8, 2, 5, 7, 2, 0, 9}
	fmt.Println(arrays)
	k := 14
	p := arr.FindSubArrayList(arrays, k)
	for _, ints := range p {
		for _, i2 := range ints {
			fmt.Printf("%d\t", i2)
		}
		fmt.Println()
	}
}
