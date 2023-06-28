package main

import (
	"fmt"
	"strings"
)

func contarAlgarismos(s string) map[string]int {

	frequenciaAlgarismos := make(map[string]int)

	strings.Fields(s)

	for _, ranS := range strings.Fields(s) {

		frequenciaAlgarismos [ranS] += 1

	}

	return frequenciaAlgarismos

}

func main() {

	s := "Toma"

	fmt.Println(contarAlgarismos(s))

}
