// Created by Daniel J. Rollins at 2024/04/16 16:24
// leetgo: 1.4.5
// https://leetcode.com/problems/group-anagrams/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type Key [26]byte

func makeKey(str string) (key Key) {
	for _, c := range str {
		key[c-'a']++
	}
	return
}

func groupAnagrams(strs []string) (ans [][]string) {
	groups := make(map[Key][]string, len(strs)/2)
	for _, str := range strs {
		key := makeKey(str)
		groups[key] = append(groups[key], str)
	}

	for _, group := range groups {
		ans = append(ans, group)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	strs := Deserialize[[]string](ReadLine(stdin))
	ans := groupAnagrams(strs)

	fmt.Println("\noutput:", Serialize(ans))
}
