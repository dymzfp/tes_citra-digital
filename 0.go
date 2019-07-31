package main

import (
	"fmt"
	"strings"
)

func alphabet(w string) {
	arrWord := strings.Split(w, "")

	fmt.Println(count(arrWord))
}

func count(s []string) int {
	return len(s)
}

func main() {
	alphabet("omama")
}