package main

import "fmt"

func main() {
	words := []string{"cullp", "ypfyqcljk", "jtuogvscob", "dsriyigc", "fr", "jtuogvscou", "doivwcgxpo", "jtuogvscob", "chuiw", "fqdhyhkvtz", "cszrtrff", "kssjy", "ps", "xvax", "vekbmwwnra"}
	pref := "jtuogvsco"
	fmt.Println(prefixCount(words, pref))
}

func prefixCount(words []string, pref string) int {
	count := 0
	n := len(pref)
	for _, word := range words {
		if len(word) >= n {
			substr := word[:n]
			if pref == substr {
				count++
			}
		}
	}
	return count
}
