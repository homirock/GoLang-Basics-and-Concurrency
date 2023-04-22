package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	newStr := ""

	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		newStr += string(word1[i]) + string(word2[j])
		i++
		j++
	}
	for i < len(word1) {
		newStr += string(word1[i])
		i++
	}

	for j < len(word2) {
		newStr += string(word2[j])
		j++
	}
	return newStr
}

func main() {
	fmt.Println(mergeAlternately("pqr", "stv"))
}
