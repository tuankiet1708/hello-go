package palindrome

import "testing"

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("able was I ere I saw elba")
	}
}
