package main

import (
	"fmt"
)

type vowelsFinder interface {
	findVowels() []rune
}

type myString string

func (ms myString) findVowels() []rune {
	var vowels []rune

	for _, run := range ms {
		if run == 'a' || run == 'e' || run == 'i' || run == 'o' || run == 'u' {
			vowels = append(vowels, run)

		}
	}

	return vowels
}

func main() {
	name := myString("this is a interface")

	var v vowelsFinder

	v = name

	fmt.Printf("Vowels are : %c\n", v.findVowels())

}
