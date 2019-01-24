package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	mysql "github.com/mdelapenya/dgt/db"
	"github.com/mdelapenya/dgt/parser"
	"github.com/spf13/cobra"
)

const alphabet = "BCDFGHJKLMNPQRSTVWXYZ"
const noSticker = "Sin distintivo"
const notFound = "No se ha encontrado ningun resultado para la matrícula introducida"
const stickerB = "Etiqueta Ambiental B"
const stickerC = "Etiqueta Ambiental C"
const stickerEco = "Etiqueta Ambiental Eco"
const stickerZero = "Etiqueta Ambiental Cero"
const userAgent = "DGT Plates Bot https://github.com/mdelapenya/dgt - " +
	"This bot just gathers info about plates, grouping them by sticker type"

var chars = []rune(alphabet)
var client *http.Client
var groupNoSticker = []string{}
var groupNotFound = []string{}
var groupZero = []string{}
var groupEco = []string{}
var groupC = []string{}
var groupB = []string{}

func init() {
	rootCmd.AddCommand(scrapCmd)

	client = &http.Client{
		Timeout: 30 * time.Second,
	}
}

var scrapCmd = &cobra.Command{
	Use:   "scrap",
	Short: "Scraps all car plates retrieving their ECO sticker",
	Long:  `Scraps all car plates retrieving their ECO sticker, starting in 0000BBB`,
	Run: func(cmd *cobra.Command, args []string) {
		scrapPlates()
	},
}

func createGrouping(plate string, html string) string {
	var sticker string

	if strings.Contains(html, notFound) {
		groupNotFound = append(groupNotFound, plate)
		sticker = notFound
	} else if strings.Contains(html, stickerB) {
		groupB = append(groupB, plate)
		sticker = stickerB
	} else if strings.Contains(html, stickerC) {
		groupC = append(groupC, plate)
		sticker = stickerC
	} else if strings.Contains(html, stickerEco) {
		groupEco = append(groupEco, plate)
		sticker = stickerEco
	} else if strings.Contains(html, stickerZero) {
		groupZero = append(groupZero, plate)
		sticker = stickerZero
	} else if strings.Contains(html, noSticker) {
		groupNoSticker = append(groupNoSticker, plate)
		sticker = noSticker
	}

	if sticker == "" {
		sticker = html
	}

	saveRequest(plate, sticker)

	return sticker
}

func formatNumber(n int) string {
	if n < 10 {
		return fmt.Sprintf("000%d", n)
	}

	if n < 100 {
		return fmt.Sprintf("00%d", n)
	}

	if n < 1000 {
		return fmt.Sprintf("0%d", n)
	}

	return fmt.Sprintf("%d", n)
}

func saveRequest(plate string, sticker string) {
	mysql.InsertPlate(plate, sticker)
}

func processPlate(plate string) string {
	url := fmt.Sprintf("http://www.dgt.es/es/seguridad-vial/distintivo-ambiental/index.shtml?accion=1&matriculahd=&matricula=%s&submit=Comprobar", plate)

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

		return createGrouping(plate, parsedHTML)
	}

	return "Not found"
}

func scrapPlates() {
	plates := []string{}

	for _, c1 := range chars {
		for _, c2 := range chars {
			for _, c3 := range chars {
				for i := 0; i < 10000; i++ {
					var sb strings.Builder
					sb.WriteString(formatNumber(i))
					sb.WriteRune(c1)
					sb.WriteRune(c2)
					sb.WriteRune(c3)

					plates = append(plates, sb.String())
				}
			}
		}
	}

	// Open all urls concurrently using the 'go' keyword:
	for _, plate := range plates {
		sticker := processPlate(plate)
		fmt.Printf("Plate %s is %s.\n", plate, sticker)
	}
}
