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
	if result == nil {
		// the plate does not exist
		//*[@id="resultadoBusqueda"]/div/div/p/text()
		result = htmlquery.FindOne(doc, `//div[@id="resultadoBusqueda"]/div/div/p/text()`)
	}

	return result.Data
}
