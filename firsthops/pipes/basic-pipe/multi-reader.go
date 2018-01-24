package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// concat with a for loop
	fmt.Println("## With for loop")
	header := strings.NewReader("<msg>")
	body := strings.NewReader("hello")
	footer := strings.NewReader("</msg>")
	for _, r := range []io.Reader{header, body, footer} {
		_, err := io.Copy(os.Stdout, r)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println()

	// concat with a multireader
	fmt.Println("\n## With multireader")
	header = strings.NewReader("<msg>")
	body = strings.NewReader("hello")
	footer = strings.NewReader("</msg>")
	r := io.MultiReader(header, body, footer)
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
	fmt.Println()
}
