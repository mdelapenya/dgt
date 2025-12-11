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
	}

	assert.Equal(t, 'B', task.firstChar)
	assert.Equal(t, 1234, task.initialIndex)
	assert.Equal(t, 5, task.secondChar)
	assert.Equal(t, 10, task.thirdChar)
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
