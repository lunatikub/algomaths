package main

import (
	"flag"
	"fmt"

	common "github.com/lunatikub/AlgoMathAndCo/common"
)

type options struct {
	dict string
	key  int
}

func getOptions() *options {
	opts := new(options)
	flag.StringVar(&opts.dict, "dict", "", "Dictionnary input")
	flag.IntVar(&opts.key, "key", 3, "Scecret key (shift parameter)")
	flag.Parse()
	return opts
}

const (
	alphaSZ = 26
)

func cipherLetter(letter, key int) int {
	return (letter + key) % alphaSZ
}

func cipherWord(word string) string {
	for _, c := range word {
		fmt.Println(c - 'a')
	}
	return word
}

func main() {
	opts := getOptions()
	D := common.NewDict(opts.dict)
	D.ToDot()
}
