package main

func romanToInt(s string) int {
	arrRomans := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	var result int

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && arrRomans[string(s[i])] < arrRomans[string(s[i+1])] {
			result += arrRomans[string(s[i+1])] - arrRomans[string(s[i])]
			i++
		} else {
			result += arrRomans[string(s[i])]
		}
	}

}

func main() {
	println(romanToInt("III"))
	println(romanToInt("IV"))
	println(romanToInt("IX"))
	println(romanToInt("LVIII"))
	println(romanToInt("MCMXCIV"))
}
