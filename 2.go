package main

import (
	"fmt"
	"strings"
	"encoding/json"
	"net/http"
)

type Hasil struct {
	Hidup 	int 	`json:"Hidup"`
	Mati 	int 	`josn:"Mati"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	http.HandleFunc("/count", HandleCount)

	fmt.Println("Server running")
	http.ListenAndServe(":1234", nil)
}

func HandleCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		wordInput := r.FormValue("input")
		if wordInput == "" {
			http.Error(w, "Input not found", http.StatusBadRequest)
			return
		}

		h, m := countAlpha(wordInput)
		data  := Hasil{h, m}

		json.NewEncoder(w).Encode(data)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
	return
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