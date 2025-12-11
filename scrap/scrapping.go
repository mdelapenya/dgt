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

const (
	invalidID = iota
	noStickerID
	stickerBID
	stickerCID
	stickerEcoID
	stickerZeroID
	notFound
)

// NotFound string representing a not found plate
const NotFound = "No se ha encontrado ningún resultado para la matrícula introducida"

const noSticker = "Sin distintivo"
const stickerB = "Distintivo Ambiental B."
const stickerC = "Distintivo Ambiental C."
const stickerEco = "Distintivo Ambiental ECO."
const stickerZero = "Distintivo Ambiental 0."
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
func ProcessPlate(plate string, persist bool) (string, error) {
	url := fmt.Sprintf("https://sede.dgt.gob.es/es/vehiculos/informacion-de-vehiculos/distintivo-ambiental/index.html?matricula=%s", plate)

	// Create and modify HTTP request before sending
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", userAgent)

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%s: Error while reading remote service", plate)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		htmlResult := string(bodyBytes)
		parsedHTML := parser.Parse(htmlResult)

		stickerID, sticker := createGrouping(plate, parsedHTML)

		if persist {
			saveRequest(plate, stickerID)
		}

		return sticker, nil
	}

	return "Not found", fmt.Errorf("%s: plate not found", plate)
}

func init() {
	client = &http.Client{
		Timeout: 30 * time.Second,
	}
}

func createGrouping(plate string, html string) (int, string) {
	var stickerID int
	var stickerDesc string

	if strings.Contains(html, NotFound) {
		groupNotFound = append(groupNotFound, plate)
		stickerID = notFound
		stickerDesc = "❌ " + NotFound
	} else if strings.Contains(html, stickerB) {
		groupB = append(groupB, plate)
		stickerID = stickerBID
		stickerDesc = "🟡 " + stickerB
	} else if strings.Contains(html, stickerC) {
		groupC = append(groupC, plate)
		stickerID = stickerCID
		stickerDesc = "🟢 " + stickerC
	} else if strings.Contains(html, stickerEco) {
		groupEco = append(groupEco, plate)
		stickerID = stickerEcoID
		stickerDesc = "🟣 " + stickerEco
	} else if strings.Contains(html, stickerZero) {
		groupZero = append(groupZero, plate)
		stickerID = stickerZeroID
		stickerDesc = "🔵 " + stickerZero
	} else if strings.Contains(html, noSticker) {
		stickerID = noStickerID
		groupNoSticker = append(groupNoSticker, plate)
		stickerDesc = "⚪️ " + noSticker
	}

	if stickerDesc == "" {
		stickerDesc = html
	}

	return stickerID, stickerDesc
}

func saveRequest(plate string, stickerID int) {
	mysql.InsertPlate(plate, stickerID)
}
