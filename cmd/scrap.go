package cmd

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"

	"github.com/mdelapenya/dgt/internal"
	"github.com/mdelapenya/dgt/scrap"
	"github.com/spf13/cobra"
)

var chars = []rune(internal.Alphabet)

var persist bool
var plate string
var from string
var until string

func init() {
	scrapCmd.Flags().StringVarP(&from, "from", "F", "", "Plate where to scrap from")
	scrapCmd.Flags().StringVarP(&until, "until", "U", "", "Plate where to scrap until (included)")
	scrapCmd.Flags().BoolVarP(&persist, "persist", "p", false, "If the result will be persisted in a data store")
	scrapCmd.Flags().StringVarP(&plate, "plate", "P", "", "Plate to scrap. It will ignore the 'persist' flag")

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

		scrapPlates(from, until)
	},
}

func scrapPlate(plate string, persist bool) error {
	sticker, err := scrap.ProcessPlate(plate, persist)
	if err != nil {
		return err
	}

	fmt.Printf("%s: %s\n", plate, sticker)
	return nil
}

// plateTask represents a task for processing all plates starting with a specific character
type plateTask struct {
	firstChar    rune
	initialIndex int
	secondChar   int
	thirdChar    int
}

func scrapPlates(fromPlate string, untilPlate string) {
	initialIndex, firstChar, secondChar, thirdChar := internal.FromPlate(fromPlate)

	// Parse the until plate to determine stopping conditions
	var uFirstChar int
	hasUntil := untilPlate != ""
	if hasUntil {
		_, uFirstChar, _, _ = internal.FromPlate(untilPlate)
	}

	// Create a pool of workers based on the number of CPUs
	numWorkers := runtime.NumCPU()

	// Channel for distributing work to goroutines
	tasks := make(chan plateTask, len(chars))
	
	// WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(tasks, &wg, persist, untilPlate)
	}

	// Distribute tasks: one task per first character
	for a := firstChar; a < len(chars); a++ {
		// Skip characters beyond the until character
		if hasUntil && a > uFirstChar {
			break
		}

		task := plateTask{
			firstChar:    chars[a],
			initialIndex: initialIndex,
			secondChar:   secondChar,
			thirdChar:    thirdChar,
		}
		tasks <- task
		
		// Reset indices after the first character
		initialIndex = 0
		secondChar = 0
		thirdChar = 0
	}

	// Close the tasks channel to signal workers that no more tasks will be sent
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()
}

// worker processes plateTask items from the tasks channel
func worker(tasks <-chan plateTask, wg *sync.WaitGroup, persist bool, untilPlate string) {
	defer wg.Done()

	for task := range tasks {
		processFirstChar(task, persist, untilPlate)
	}
}

// processFirstChar processes all plates starting with a specific first character
func processFirstChar(task plateTask, persist bool, untilPlate string) {
	c1 := task.firstChar
	initialIndex := task.initialIndex
	secondChar := task.secondChar
	thirdChar := task.thirdChar

	for b := secondChar; b < len(chars); b++ {
		c2 := chars[b]
		for c := thirdChar; c < len(chars); c++ {
			continueProcessing := processPlates(initialIndex, c1, c2, c, persist, untilPlate)
			if !continueProcessing {
				return
			}

			initialIndex = 0
			thirdChar = 0
		}
		secondChar = 0
	}
}

func processPlate(number int, c1 rune, c2 rune, c3 rune, persist bool) {
	var sb strings.Builder

	sb.WriteString(internal.FormatNumber(number))
	sb.WriteRune(c1)
	sb.WriteRune(c2)
	sb.WriteRune(c3)

	err := scrapPlate(sb.String(), persist)
	if err != nil {
		log.Fatal(err)
	}
}

// processPlates processes all the plates from the given initial index, until the given until plate
// It will return true if the outer process should continue, or false if it should stop
func processPlates(initialIndex int, c1 rune, c2 rune, thirdChar int, persist bool, untilPlate string) bool {
	c3 := chars[thirdChar]
	
	// Parse the until plate once if provided
	var shouldCheckUntil bool
	var uInitialIndex, uFirstChar, uSecondChar, uThirdChar int
	if untilPlate != "" {
		shouldCheckUntil = true
		uInitialIndex, uFirstChar, uSecondChar, uThirdChar = internal.FromPlate(untilPlate)
	}
	
	for i := initialIndex; i < 10000; i++ {
		processPlate(i, c1, c2, c3, persist)

		// if the plate is the until plate, stop the process
		if shouldCheckUntil && i == uInitialIndex && c1 == chars[uFirstChar] && c2 == chars[uSecondChar] && c3 == chars[uThirdChar] {
			return false
		}
	}

	return true
}
