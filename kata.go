package main

import "fmt"

func main() {
	str := "ada"
	fmt.Println(IsPalindrome(str))

}

func IsPalindrome(str string) bool {
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	if str != string(runes) {
		return false
	}

	return true
}
