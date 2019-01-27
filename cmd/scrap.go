package cmd

import (
	"fmt"
	"strings"

	"github.com/mdelapenya/dgt/scrap"
	"github.com/spf13/cobra"
)

const alphabet = "BCDFGHJKLMNPQRSTVWXYZ"

var chars = []rune(alphabet)

func init() {
	rootCmd.AddCommand(scrapCmd)
}

var scrapCmd = &cobra.Command{
	Use:   "scrap",
	Short: "Scraps all car plates retrieving their ECO sticker",
	Long:  `Scraps all car plates retrieving their ECO sticker, starting in 0000BBB`,
	Run: func(cmd *cobra.Command, args []string) {
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
		sticker := scrap.ProcessPlate(plate)
		fmt.Printf("Plate %s is %s.\n", plate, sticker)
	}
}
