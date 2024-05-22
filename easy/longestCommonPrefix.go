package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	lengthStrs := len(strs)
	if lengthStrs == 0 {
		return ""
	}

	minStr := 0

	for i := 0; i < lengthStrs-1; i++ {
		if i+1 > lengthStrs {
			break
		}
		if len(strs[i]) > len(strs[i+1]) {
			minStr = len(strs[i+1])
		}
	}

	lStr := ""

	for i := 0; i < minStr-1; i++ {
		for j := 0; j < lengthStrs-1; j++ {
			if j+1 > lengthStrs {
				break
			}
			if strs[j] == strs[j+1] {
				lStr = lStr + strs[j]
			}
		}
	}

	return lStr

}

func main() {

	fmt.Println("result:", longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
