package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mdelapenya/dgt/scrap"
	"github.com/spf13/cobra"
)

const alphabet = "BCDFGHJKLMNPQRSTVWXYZ"

var chars = []rune(alphabet)

var persist bool
var plate string
var from string

func init() {
	scrapCmd.Flags().StringVarP(&from, "from", "F", "", "Plate where to scrap from")
	scrapCmd.Flags().BoolVarP(&persist, "persist", "p", false, "If the result will be persisted in a data store")
	scrapCmd.Flags().StringVarP(&plate, "plate", "P", "", "Plate to scrap")

	rootCmd.AddCommand(scrapCmd)
}

var scrapCmd = &cobra.Command{
	Use:   "scrap",
	Short: "Scraps all car plates retrieving their ECO sticker",
	Long:  `Scraps all car plates retrieving their ECO sticker, starting in 0000BBB`,
	Run: func(cmd *cobra.Command, args []string) {
		if plate != "" {
			scrapPlate(plate, false)
			return
		}

		scrapPlates(from)
	},
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

func scrapPlate(plate string, persist bool) {
	sticker := scrap.ProcessPlate(plate, persist)
	fmt.Printf("%s: %s\n", plate, sticker)
}

func scrapPlates(fromPlate string) {
	initialIndex := 0
	firstChar := 0
	secondChar := 0
	thirdChar := 0

	if fromPlate != "" {
		plateNumber, err := strconv.Atoi(string(fromPlate[0:4]))
		if err == nil {
			initialIndex = plateNumber
		}

		lowerCaseAlphabet := strings.ToLower(alphabet)

		firstChar = strings.IndexRune(lowerCaseAlphabet, rune(fromPlate[4]))
		secondChar = strings.IndexRune(lowerCaseAlphabet, rune(fromPlate[5]))
		thirdChar = strings.IndexRune(lowerCaseAlphabet, rune(fromPlate[6]))
	}

	for a := firstChar; a < len(chars); a++ {
		c1 := chars[a]
		for b := secondChar; b < len(chars); b++ {
			c2 := chars[b]
			for c := thirdChar; c < len(chars); c++ {
				c3 := chars[c]
				for i := initialIndex; i < 10000; i++ {
					var sb strings.Builder
					sb.WriteString(formatNumber(i))
					sb.WriteRune(c1)
					sb.WriteRune(c2)
					sb.WriteRune(c3)

					scrapPlate(sb.String(), persist)
				}
				initialIndex = 0
				thirdChar = 0
			}
			secondChar = 0
		}
		firstChar = 0
	}
}
