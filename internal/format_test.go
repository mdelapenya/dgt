package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatNumber(t *testing.T) {
	s := FormatNumber(0)
	assert.Equal(t, "0000", s)

	s = FormatNumber(1)
	assert.Equal(t, "0001", s)

	s = FormatNumber(10)
	assert.Equal(t, "0010", s)

	s = FormatNumber(100)
	assert.Equal(t, "0100", s)

	s = FormatNumber(999)
	assert.Equal(t, "0999", s)

	s = FormatNumber(1000)
	assert.Equal(t, "1000", s)

	s = FormatNumber(9999)
	assert.Equal(t, "9999", s)
}
