package main

import (
	"fmt"
	"strings"
)

func main() {
	hidup, mati := countAlpha("omama") 

	fmt.Println(hidup, mati)
}

func countAlpha(w string) (x, y int) {
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

	x = len(arrConst)
	y = len(arrVocal)
	return	
}