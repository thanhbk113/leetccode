package main

// google interview question
func repeatedCharacter(s string) byte {
	seen_letters := [256]bool{} // it take O(1) space comlexity

	for _, c := range s {
		if seen_letters[c] {
			return byte(c)
		}
		seen_letters[c] = true
	}
	return 0
}
