package internal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromPlate(t *testing.T) {
	tests := []struct {
		plate         string
		expectedIndex int
		expectedC1    int
		expectedC2    int
		expectedC3    int
	}{
		{plate: "0000BCD", expectedIndex: 0, expectedC1: 0, expectedC2: 1, expectedC3: 2},
		{plate: "1234FGH", expectedIndex: 1234, expectedC1: 3, expectedC2: 4, expectedC3: 5},
		{plate: "4567JKL", expectedIndex: 4567, expectedC1: 6, expectedC2: 7, expectedC3: 8},
		{plate: "8901MNP", expectedIndex: 8901, expectedC1: 9, expectedC2: 10, expectedC3: 11},
		{plate: "7620QRS", expectedIndex: 7620, expectedC1: 12, expectedC2: 13, expectedC3: 14},
		{plate: "8804TVW", expectedIndex: 8804, expectedC1: 15, expectedC2: 16, expectedC3: 17},
		{plate: "2943XYZ", expectedIndex: 2943, expectedC1: 18, expectedC2: 19, expectedC3: 20},
		// non existing characters
		{plate: "0000AÁÀ", expectedIndex: 0, expectedC1: -1, expectedC2: -1, expectedC3: -1},
		{plate: "0000EÉÈ", expectedIndex: 0, expectedC1: -1, expectedC2: -1, expectedC3: -1},
		{plate: "0000IÍÌ", expectedIndex: 0, expectedC1: -1, expectedC2: -1, expectedC3: -1},
		{plate: "0000OÓÒ", expectedIndex: 0, expectedC1: -1, expectedC2: -1, expectedC3: -1},
		{plate: "0000UÚÙ", expectedIndex: 0, expectedC1: -1, expectedC2: -1, expectedC3: -1},
		{plate: "0000ÑÇß", expectedIndex: 0, expectedC1: -1, expectedC2: -1, expectedC3: -1},
	}

	for _, test := range tests {
		i, c1, c2, c3 := FromPlate(test.plate)

		assert.Equal(t, test.expectedIndex, i)
		assert.Equal(t, test.expectedC1, c1)
		assert.Equal(t, test.expectedC2, c2)
		assert.Equal(t, test.expectedC3, c3)

		i, c1, c2, c3 = FromPlate(strings.ToLower(test.plate))

		assert.Equal(t, test.expectedIndex, i)
		assert.Equal(t, test.expectedC1, c1)
		assert.Equal(t, test.expectedC2, c2)
		assert.Equal(t, test.expectedC3, c3)
	}
}
