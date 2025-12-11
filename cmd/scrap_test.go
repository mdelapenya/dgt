package cmd

import (
	"testing"

	"github.com/mdelapenya/dgt/internal"
	"github.com/stretchr/testify/assert"
)

func TestPlateTaskCreation(t *testing.T) {
	// Test that plateTask struct properly captures task information
	task := plateTask{
		firstChar:    'B',
		initialIndex: 1234,
		secondChar:   5,
		thirdChar:    10,
		untilIndex:   5678,
		untilFirst:   8,
		untilSecond:  12,
		untilThird:   15,
		hasUntil:     true,
	}

	assert.Equal(t, 'B', task.firstChar)
	assert.Equal(t, 1234, task.initialIndex)
	assert.Equal(t, 5, task.secondChar)
	assert.Equal(t, 10, task.thirdChar)
	assert.Equal(t, 5678, task.untilIndex)
	assert.Equal(t, 8, task.untilFirst)
	assert.Equal(t, 12, task.untilSecond)
	assert.Equal(t, 15, task.untilThird)
	assert.True(t, task.hasUntil)
}

func TestProcessFirstCharLogic(t *testing.T) {
	// Test that the character indices are handled correctly
	chars := []rune(internal.Alphabet)
	
	// First character should be 'B' (index 0)
	assert.Equal(t, 'B', chars[0])
	
	// Last character should be 'Z' (index 20)
	assert.Equal(t, 'Z', chars[len(chars)-1])
	
	// Check alphabet length
	assert.Equal(t, 21, len(chars))
}
