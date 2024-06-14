package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// Функция для сортировки букв в слове
func sortedString(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func findAnagramSets(words []string) map[string][]string {
	anagramMap := make(map[string][]string)

	for _, word := range words {
		lowerWord := strings.Map(unicode.ToLower, word)
		sortString := sortedString(lowerWord)
		anagramMap[sortString] = append(anagramMap[sortString], lowerWord)
	}

	result := make(map[string][]string)
	for _, value := range anagramMap {
		if len(value) > 1 {
			sort.Strings(value)
			result[value[0]] = value
		}
	}

	return result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "ток"}

	anagramMap := findAnagramSets(words)

	for key, val := range anagramMap {
		fmt.Printf("%s: %v\n", key, val)
	}
}
