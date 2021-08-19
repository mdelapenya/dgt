package scrap

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	mysql "github.com/mdelapenya/dgt/db"
	"github.com/mdelapenya/dgt/parser"
)

// NotFound string representing a not found plate
const NotFound = "No se ha encontrado ningún resultado para la matrícula introducida"

const noSticker = "Sin distintivo"
const stickerB = "Etiqueta Ambiental B Amarilla"
const stickerC = "Etiqueta Ambiental C Verde"
const stickerEco = "Etiqueta Ambiental Eco"
const stickerZero = "Etiqueta Ambiental 0"
const userAgent = "DGT Plates Bot https://github.com/mdelapenya/dgt - " +
	"This bot just gathers info about plates, grouping them by sticker type"

var client *http.Client
var groupNoSticker = []string{}
var groupNotFound = []string{}
var groupZero = []string{}
var groupEco = []string{}
var groupC = []string{}
var groupB = []string{}

// ProcessPlate fetches plate information from DGT web site, using scrapping techniques
func ProcessPlate(plate string, persist bool) string {
	url := fmt.Sprintf("https://sede.dgt.gob.es/es/vehiculos/distintivo-ambiental/?accion=1&matriculahd=&matricula=%s&submit=Consultar", plate)

	// Create and modify HTTP request before sending
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", userAgent)

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%s: Error while reading remote service", plate)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		htmlResult := string(bodyBytes)
		parsedHTML := parser.Parse(htmlResult)

		return createGrouping(plate, parsedHTML, persist)
	}

	return "Not found"
}

func init() {
	client = &http.Client{
		Timeout: 30 * time.Second,
	}
}

func createGrouping(plate string, html string, persist bool) string {
	var sticker string

	if strings.Contains(html, NotFound) {
		groupNotFound = append(groupNotFound, plate)
		sticker = "❌ " + NotFound
	} else if strings.Contains(html, stickerB) {
		groupB = append(groupB, plate)
		sticker = "🟡 " + stickerB
	} else if strings.Contains(html, stickerC) {
		groupC = append(groupC, plate)
		sticker = "🟢 " + stickerC
	} else if strings.Contains(html, stickerEco) {
		groupEco = append(groupEco, plate)
		sticker = "🟣 " + stickerEco
	} else if strings.Contains(html, stickerZero) {
		groupZero = append(groupZero, plate)
		sticker = "🔵 " + stickerZero
	} else if strings.Contains(html, noSticker) {
		groupNoSticker = append(groupNoSticker, plate)
		sticker = "⚪️ " + noSticker
	}

	if sticker == "" {
		sticker = html
	}

	if persist {
		saveRequest(plate, sticker)
	}

	return sticker
}

func saveRequest(plate string, sticker string) {
	mysql.InsertPlate(plate, sticker)
}
