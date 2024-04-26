// Created by Daniel J. Rollins at 2024/04/16 16:08
// leetgo: 1.4.5
// https://leetcode.com/problems/two-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(nums []int, target int) (ans []int) {
	cache := make(map[int]int, len(nums))

	for i, num := range nums {
		if j, ok := cache[num]; ok {
			return []int{j, i}
		}
		cache[target-num] = i
	}

	return []int{}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
