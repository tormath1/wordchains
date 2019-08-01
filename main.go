package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type node struct {
	parents []node
	word    string
}

func findShortestChain(start, end string, univ, chain *[]string) {
	toVisit := []node{
		node{
			word:    start,
			parents: nil,
		},
	}
	var visited []string
	var neigh []string

	for i := 0; i < len(toVisit); i++ {
		if strings.Compare(toVisit[i].word, end) == 0 {
			pprint(toVisit[i].parents, chain)
			return
		}
		getNeighbors(toVisit[i].word, univ, &neigh)
		for _, n := range neigh {
			if !contains(n, &visited) {
				visited = append(visited, n)
				no := node{word: n, parents: []node{toVisit[i]}}
				toVisit = append(toVisit, no)
			}
		}
		neigh = nil
	}
}

func pprint(chain []node, res *[]string) {
	if chain == nil {
		return
	}
	*res = append(*res, chain[0].word)
	pprint(chain[0].parents, res)
}

func contains(el string, arr *[]string) bool {
	for _, a := range *arr {
		if strings.Compare(el, a) == 0 {
			return true
		}
	}
	return false
}

// getNeighbors return the word arounds the given word
// with 1 character of difference
func getNeighbors(word string, universe, neigh *[]string) {
	for _, w := range *universe {
		if isNeighbor(w, word) {
			*neigh = append(*neigh, w)
		}
	}
}

// preProcessData will remove words who are
// not the same size as the inputs
func preProcessData(start string, words, res *[]string) {
	l := len(start)

	for _, w := range *words {
		if len(w) == l {
			*res = append(*res, w)
		}
	}
}

// isWordInTheUniverse this function seems to be directly
// fetch from the stargate. Seriously, we return a boolean
// in order to check if start and end word are in the universe.
func isWordInTheUniverse(word string, universe *[]string) bool {
	for _, w := range *universe {
		if strings.Compare(word, w) == 0 {
			return true
		}
	}
	return false
}

// IsNeighbor return true if the two words
// have exactly 1 different character
func isNeighbor(w, word string) bool {
	var diff int
	for i := 0; i < len(w); i++ {
		if w[i] != word[i] {
			diff++
		}
	}
	return diff == 1
}

func main() {
	f, err := os.Open("./wordlist.txt")

	start := "ruby"
	end := "code"
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

	preProcessData(start, &wordsList, &words)

	if !isWordInTheUniverse(start, &words) || !isWordInTheUniverse(end, &words) {
		log.Fatal("start or stop word is not in the universe")
	}

	var chain []string
	findShortestChain(start, end, &words, &chain)
	log.Println(chain)
}
