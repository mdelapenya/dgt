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

	// Look for the success case: div with border-success class contains the badge type
	result := htmlquery.FindOne(doc, `//div[contains(@class, "border-success")]//p/strong[2]/text()`)
	if result != nil {
		return result.Data
	}

	// Fallback: look for error/not found cases (you may need to adjust this based on the error HTML)
	result = htmlquery.FindOne(doc, `//div[contains(@class, "border-danger")]//p/text()`)
	if result != nil {
		return result.Data
	}

	return "No encontrado"
}
