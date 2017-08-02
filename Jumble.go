/***
Jumble: Given a random string, re-arrange it into legitimate English words from a given Dictionary
e.x. input: "iyanr" output: "rainy"
	input: "aecf" output: ["cafe","face"]
***/

package main

import (
	"flag"
	"fmt"

	"github.com/gengwensu/mylib"
)

var Istr = flag.String("Istr", "aecf", "a random string")

func main() {
	dict := []string{"rainy", "cafe", "face", "listen", "silent"}
	flag.Parse()

	wh := map[string][]string{} //create a haspMap of all words in dict
	for _, w := range dict {
		k := mylib.SortStringByCharacter(w) //use sorted string as key
		wh[k] = append(wh[k], w)
		//	fmt.Printf("add map with key %v with word %v\n", k, w)
	}
	key := mylib.SortStringByCharacter(*Istr)
	fmt.Printf("The results are: %v\n", wh[key])
}
