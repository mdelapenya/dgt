package cmd

import (
	"fmt"
	"strings"

	"github.com/mdelapenya/dgt/scrap"
	"github.com/spf13/cobra"
)

const alphabet = "BCDFGHJKLMNPQRSTVWXYZ"

var chars = []rune(alphabet)

var persist bool
var plate string

func init() {
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

		scrapPlates()
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

func scrapPlates() {
	for _, c1 := range chars {
		for _, c2 := range chars {
			for _, c3 := range chars {
				for i := 0; i < 10000; i++ {
					var sb strings.Builder
					sb.WriteString(formatNumber(i))
					sb.WriteRune(c1)
					sb.WriteRune(c2)
					sb.WriteRune(c3)

					scrapPlate(sb.String(), persist)
				}
			}
		}
	}
}
