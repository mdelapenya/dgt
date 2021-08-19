package parser

import (
	"log"
	"strings"

	htmlquery "github.com/antchfx/htmlquery"
)

// Parse parses an HTML
func Parse(HTML string) string {
	doc, err := htmlquery.Parse(strings.NewReader(HTML))
	if err != nil {
		log.Fatalln("The HTML is not parseable")
	}

	result := htmlquery.FindOne(doc, `//div[@id="resultadoBusqueda"]/div/div/p/strong/text()`)

	return result.Data
}
