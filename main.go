package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// PreProcessData will remove words who are
// not the same size as the inputs
func PreProcessData(start string, words, res *[]string) {
	l := len(start)

	for _, w := range *words {
		if len(w) == l {
			*res = append(*res, w)
		}
	}
}

// IsWordInTheUniverse this function seems to be directly
// fetch from the stargate. Seriously, we return a boolean
// in order to check if start and end word are in the universe.
func IsWordInTheUniverse(word string, universe *[]string) bool {
	for _, w := range *universe {
		if strings.Compare(word, w) == 0 {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("./wordlist.txt")

	start := "dog"
	end := "cat"
	if len(start) != len(end) {
		log.Fatal("start and end words must have the same size")
	}
	if err != nil {
		log.Fatalf("unable to open wordlist: %v", err)
	}
	defer f.Close()
	var wordsList []string
	var words []string

	s := bufio.NewScanner(f)

	for s.Scan() {
		wordsList = append(wordsList, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	PreProcessData(start, &wordsList, &words)

	if !IsWordInTheUniverse(start, &words) || !IsWordInTheUniverse(end, &words) {
		log.Fatal("start or stop word is not in the universe")
	}
}
