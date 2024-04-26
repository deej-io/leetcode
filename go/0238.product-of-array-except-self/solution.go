// Created by Daniel J. Rollins at 2024/04/16 19:33
// leetgo: 1.4.5
// https://leetcode.com/problems/product-of-array-except-self/

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func init() {
	debug.SetGCPercent(-1)
}

func productExceptSelf(nums []int) []int {
	len := len(nums)
	fwd := make([]int, len)
	bwd := make([]int, len)

	fwd[0] = 1
	bwd[len-1] = 1

	for i := 1; i < len; i++ {
		fwd[i] = fwd[i-1] * nums[i-1]
		bwd[len-1-i] = bwd[len-i] * nums[len-i]
	}

	for i := range nums {
		nums[i] = fwd[i] * bwd[i]
	}

	return nums
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := productExceptSelf(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
