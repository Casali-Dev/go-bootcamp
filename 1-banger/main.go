package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	msg := os.Args[1]
	l := utf8.RuneCountInString(msg)

	banger := strings.Repeat("!", l)

	s := banger + msg + banger
	s = strings.ToUpper(s)

	fmt.Println(s)
}
