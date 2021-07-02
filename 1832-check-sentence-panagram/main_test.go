package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	sentence := "thequickbrownfoxjumpsoverthelazydog"

	result := checkIfPangram(sentence)

	assert.True(t, result)
}

func Test2(t *testing.T) {
	sentence := "leetcode"

	result := checkIfPangram(sentence)

	assert.False(t, result)
}

