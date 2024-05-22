package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	checkS := make(map[byte]byte, 0)
	checkT := make(map[byte]byte, 0)
	nS, nT := len(s), len(t)

	if nS != nT {
		return false
	}

	for i := 0; i < nS; i++ {
		if valS, okS := checkS[s[i]]; okS {
			if valS != t[i] {
				return false
			}
		} else {
			if _, okT := checkT[t[i]]; okT {
				return false
			}
			checkS[s[i]] = t[i]
			checkT[t[i]] = s[i]
		}
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	a := isIsomorphic(s, t)
	fmt.Println("🚀 ~ funcmain ~ a : ", a)
}
