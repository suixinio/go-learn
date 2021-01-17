package main

import (
	"fmt"
	"unicode/utf8"
)

func lenghtOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		fmt.Println(ch)
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	s := "Yes我爱慕课网!" //UTF-8 可变长度
	fmt.Println(len(s))
	fmt.Println("", []byte(s))

	for i, ch := range []byte(s) {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

	fmt.Println(
		lenghtOfNonRepeatingSubStr("abcabcbb"),
		lenghtOfNonRepeatingSubStr("一二三一"),
	)
}
