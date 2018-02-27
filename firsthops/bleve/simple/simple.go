package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/query"
)

const (
	dbPath = "people.bleve"
)

var (
	persons = [][]string{
		{"Bruno", "Sinou"},
		{"John", "Doo"},
		{"Jackie", "Chan"},
		{"Jackie", "Chan Jr."},
		{"Jackie", "Meetoo"},
		{"John", "Difool"},
	}
)

type Person struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	DisplayName string `json:"displayname"`
}

func main() {
	cleanLegacy()
	createIndex()
	testSearches()
}

func testSearches() {
	index, err := bleve.Open(dbPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer index.Close()

	simpleNoResultSearch(index)
	simpleTwoResultSearch(index)
	simpleDNameSearch(index)
	advancedDNameSearch(index)
}

// Various search patterns

func simpleNoResultSearch(index bleve.Index) {
	query := bleve.NewTermQuery("John") // query string terms must be lower case
	request := bleve.NewSearchRequest(query)
	result, err := index.Search(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("simpleNoResultSearch: " + result.String())
}

func simpleTwoResultSearch(index bleve.Index) {
	query := bleve.NewTermQuery("john")
	request := bleve.NewSearchRequest(query)
	result, err := index.Search(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("simpleTwoResultSearch: " + result.String())
}

func simpleDNameSearch(index bleve.Index) {
	query := bleve.NewTermQuery("difool")
	request := prepareRequest(query)
	result, err := index.Search(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("simpleDNameSearch: " + result.String())
}

func advancedDNameSearch(index bleve.Index) {
	phrase := []string{"jackie", "chan"}
	query := bleve.NewPhraseQuery(phrase, "displayname")
	request := prepareRequest(query)
	result, err := index.Search(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("advancedDNameSearch: " + result.String())
}

// Local helpers
func prepareRequest(query query.Query) *bleve.SearchRequest {
	request := bleve.NewSearchRequest(query)
	request.Highlight = bleve.NewHighlightWithStyle("html")
	request.Fields = []string{"firstname", "lastname", "displayname"}
	return request
}

func cleanLegacy() {
	os.RemoveAll(dbPath)
}

func newPerson(fname, lname string) Person {
	dname := fname + " " + lname
	return Person{
		FirstName:   fname,
		LastName:    lname,
		DisplayName: dname,
	}
}

func createIndex() {

	index, err := bleve.Open(dbPath)
	if err == nil { // index already exists, nothing to do
		fmt.Println("Index already exists, nothing to do")
		index.Close()
		return
	}

	mapping := bleve.NewIndexMapping()
	index, err = bleve.New(dbPath, mapping)
	if err != nil {
		log.Fatal(err)
	}

	for i, personStr := range persons {
		person := newPerson(personStr[0], personStr[1])
		err = index.Index("m"+strconv.Itoa(i), person)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Index created, documents indexed")
	index.Close()
}
