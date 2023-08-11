package tests_controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestAdd(t *testing.T) {
	assert.Equal(t, Add(1,1), 2)
}