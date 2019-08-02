package main

import (
	"fmt"
	"strings"
)

func main() {
	sortWord("omama")
}

func sortWord(w string) {
	arrWord := strings.Split(strings.Replace(w, " ", "", -1), "")
	vocal := []string{"a", "i", "u", "e", "o"}
	var life, dead []string
	var hasil string
	for _, value := range arrWord {
		for i, v := range vocal {
			if value == v {
				life = append(life, value)
				break
			} else if i == len(vocal) - 1 {
				dead = append(dead, value)
			}
		}
	}
	alphabet := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	for _, y := range alphabet {
		for _, x := range life {
			if y == x {
				hasil += x
			}
		}
	}
	for _, y := range alphabet {
		for _, x := range dead {
			if y == x {
				hasil += x
			}
		}
	}
	fmt.Println(hasil)
}