package palindrome

import "fmt"

func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		fmt.Println(runes[i], runes[len(runes)-1-i])
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}
