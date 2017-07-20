/***
Jumble: Given a random string, re-arrange it into legitimate English words from a given Dictionary
e.x. input: "iyanr" output: "rainy"
	input: "aecf" output: ["cafe","face"]
***/

package main

import (
	"flag"
	"fmt"
	"sort"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}

var Istr = flag.String("Istr", "aecf", "a random string")

func main() {
	dict := []string{"rainy", "cafe", "face", "listen", "silent"}
	flag.Parse()

	wh := map[string][]string{} //create a haspMap of all words in dict
	for _, w := range dict {
		k := SortStringByCharacter(w) //use sorted string as key
		wh[k] = append(wh[k], w)
		//	fmt.Printf("add map with key %v with word %v\n", k, w)
	}
	key := SortStringByCharacter(*Istr)
	fmt.Printf("The results are: %v\n", wh[key])
}
