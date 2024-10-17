package main

import (
	"fmt"
	"strings"
)

func count_vowels(str string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count += 1
		}
	}
	return count
}

func main() {
	name := []string{"Samarth", "Kolapkar"}
	var vowels_count int = 0
	for _, str := range name {
		vowels_count = count_vowels(str)
		fmt.Println(str, vowels_count)
	}

}
