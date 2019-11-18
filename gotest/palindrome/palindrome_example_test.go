package palindrome

import "fmt"

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("able was I ere I saw elba"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}
