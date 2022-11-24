package cmd

import (
	"fmt"
	"strings"

	"github.com/mdelapenya/dgt/internal"
	"github.com/mdelapenya/dgt/scrap"
	"github.com/spf13/cobra"
)

var chars = []rune(internal.Alphabet)

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

func scrapPlate(plate string, persist bool) {
	sticker := scrap.ProcessPlate(plate, persist)
	fmt.Printf("%s: %s\n", plate, sticker)
}

func scrapPlates(fromPlate string) {
	initialIndex, firstChar, secondChar, thirdChar := internal.FromPlate(fromPlate)

	for a := firstChar; a < len(chars); a++ {
		c1 := chars[a]
		for b := secondChar; b < len(chars); b++ {
			c2 := chars[b]
			for c := thirdChar; c < len(chars); c++ {
				c3 := chars[c]
				for i := initialIndex; i < 10000; i++ {
					var sb strings.Builder
					sb.WriteString(internal.FormatNumber(i))
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
