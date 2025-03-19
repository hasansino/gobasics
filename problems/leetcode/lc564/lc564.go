package lc564

import (
	"fmt"
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {
	if n == "1" {
		return "0"
	}

	length := len(n)
	half := (length + 1) / 2
	prefix, _ := strconv.Atoi(n[:half])

	candidates := []int{
		getPalindrome(prefix, length%2 == 0),
		getPalindrome(prefix-1, length%2 == 0),
		getPalindrome(prefix+1, length%2 == 0),
		int(math.Pow10(length-1)) - 1,
		int(math.Pow10(length)) + 1,
	}

	fmt.Printf("candidates => %v \n", candidates)

	original, _ := strconv.Atoi(n)
	closest := -1

	for _, candidate := range candidates {
		if candidate == original {
			continue
		}
		if closest == -1 || abs(candidate-original) < abs(closest-original) || (abs(candidate-original) == abs(closest-original) && candidate < closest) {
			closest = candidate
		}
	}

	return strconv.Itoa(closest)
}

func getPalindrome(prefix int, even bool) int {
	s := strconv.Itoa(prefix)
	r := reverseString(s)
	if even {
		s += r
	} else {
		s += r[1:]
	}
	result, _ := strconv.Atoi(s)
	return result
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
