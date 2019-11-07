package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"work/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	// log.Println("APIKEY: " + apiKey)
	thesaurus := &thesaurus.BigHuge{APIKEY: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for sysnonyms for "+word, err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms for " + word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
