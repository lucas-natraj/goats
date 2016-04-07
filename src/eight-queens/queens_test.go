package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMethod(t *testing.T) {
	fmt.Println("Yo")
	assert.Equal(1, 1)
}
