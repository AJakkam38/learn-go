package main

import (
	"fmt"
	"os"
	"strings"
)

const sentence = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(sentence)
	find := os.Args[1:]

	for _, f := range find {
		for index, word := range words {
			if f == word {
				fmt.Printf("%2d: %q\n", index+1, word)
			}
		}
	}
}
