package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//func Unpacker(input string) (string, error) {
//
//	var sb strings.Builder
//
//	for i := range input {
//		if !isDigit(input[i]) {
//			sb.WriteRune(rune(input[i]))
//		} else if isDigit(input[i]) {
//			if i == 0 {
//				return "", fmt.Errorf("некорректная строка")
//			}
//			count, _ := strconv.Atoi(string(input[i]))
//			for j := 0; j < count-1; j++ {
//				sb.WriteRune(rune(input[i-1]))
//			}
//
//		}
//	}
//
//	return sb.String(), nil
//}

func Unpacker(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("входная строка пуста")
	}
	var sb strings.Builder
	re := regexp.MustCompile("[0-9]+|[a-zA-Z]")
	matches := re.FindAllString(input, -1)

	for i := range matches {
		if isDigit(matches[i]) {
			if i == 0 {
				return "", errors.New("первый символ строки - цифра")
			}
			count, _ := strconv.Atoi(matches[i])
			for j := 0; j < count; j++ {
				sb.WriteString(matches[i-1])
			}
		} else if !isDigit(matches[i]) {
			if i != len(matches)-1 && !isDigit(matches[i+1]) {
				sb.WriteString(matches[i])
			} else if i == len(matches)-1 {
				sb.WriteString(matches[i])
			}
		}
	}
	return sb.String(), nil
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	s := "wa4bc2d5e"
	fmt.Println(Unpacker(s))
}
