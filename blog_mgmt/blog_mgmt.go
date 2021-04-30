package main

import (
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	filename := "/giu/chamberlain/books/summary/README"
	f, err := ioutil.ReadFile(filename + ".md")
	if err != nil {
		log.Println(err.Error())
	}

	content := template.HTML(blackfriday.Run(f))
	convertHtml := strings.ReplaceAll((string)(content), ".md", ".html")

	log.Print(convertHtml)

	saveErr := ioutil.WriteFile(filename + ".html", []byte(convertHtml), 0660)
	if saveErr != nil {
		log.Print(saveErr.Error())
	}
}