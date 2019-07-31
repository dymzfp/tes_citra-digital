package main

import (
	"fmt"
	"strings"
)

func main() {
	alphabet("oma ma")
}

func alphabet(w string) {
	word := strings.Replace(w, " ", "", -1)
	arrWord := strings.Split(word, "")
	vocal := []string{"a", "i", "u", "e", "o"}
	var arrVocal []string
	var arrConst []string

	for _, value := range arrWord {
		for i, v := range vocal {
			if value == v {
				if len(arrVocal) != 0 {
					for u, y := range arrVocal {
						if value == y {
							break
						} else if u == len(arrVocal) - 1 {
							arrVocal = append(arrVocal, value)
						} 
					}
				} else {
					arrVocal = append(arrVocal, value)
				}
				break
			} else if i == len(vocal) - 1 {
				arrConst = append(arrConst, value)
			}
		}
	}

	fmt.Println("huruf mati :", len(arrConst))
	fmt.Println("huruf hidup :", len(arrVocal))
}