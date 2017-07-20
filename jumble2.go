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
	"strings"
)

var Istr = flag.String("Istr", "aecf", "a random string")

func main() {
	dict := []string{"rainy", "cafe", "face", "listen", "silent"}
	flag.Parse()

	wh := map[string][]string{} //create a haspMap of all words in dict
	for _, w := range dict {
		k := SortString(w) //use sorted string as key
		wh[k] = append(wh[k], w)
		//	fmt.Printf("add map with key %v with word %v\n", k, w)
	}
	key := SortString(*Istr)
	fmt.Printf("The results are: %v\n", wh[key])
}
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
