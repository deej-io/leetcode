// Created by Daniel J. Rollins at 2024/04/16 16:47
// leetgo: 1.4.5
// https://leetcode.com/problems/top-k-frequent-elements/

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type occurences struct {
	num   int
	count int
}

func topKFrequent(nums []int, k int) (ans []int) {
	ans = make([]int, 0 ,k)
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	orderedCounts := make([]occurences, 0, len(counts))
	for num, count := range counts {
		orderedCounts = append(orderedCounts, occurences{num: num, count: count})
	}
	slices.SortFunc(orderedCounts, func(a occurences, b occurences) int {
		return b.count - a.count
	})
	for i := 0; i < k && i < len(orderedCounts); i++ {
		ans = append(ans, orderedCounts[i].num)
	}
	return 
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := topKFrequent(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
