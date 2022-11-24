package internal

import "fmt"

func FormatNumber(n int) string {
	return fmt.Sprintf("%04d", n)
}
