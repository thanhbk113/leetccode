package main

import (
	"fmt"
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	ans := make(map[string][]string)

	for _, str := range strs {
		count := make([]int, 26) // Tạo một mảng đếm tần suất của mỗi ký tự
		for _, char := range str {
			count[char-'a']++
		}

		// Tạo từ khóa từ mảng đếm
		key := ""
		for _, c := range count {
			key += strconv.Itoa(c) + "/"
		}

		ans[key] = append(ans[key], str)
	}

	var grouped [][]string
	for _, group := range ans {
		grouped = append(grouped, group)
	}
	return grouped
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))

}
