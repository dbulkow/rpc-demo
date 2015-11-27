package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	parts := strings.Split(str, " ")
	out := ""
	for _, s := range parts {
		out = s + " " + out
	}
	return out
}

func main() {
	val := ReverseWords("some string that wants to be backwards")
	fmt.Println(val)
}
