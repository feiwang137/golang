/*
@Description: xxxx
@Author: feiwang
@Date: 2023/4/23 23:04
*/
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	t.Log("start testing")
	result := add(1, 2)
	assert.Equal(t, result, 3)
}
