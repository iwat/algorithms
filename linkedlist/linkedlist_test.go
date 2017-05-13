package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	head := Create("a", "b", "c", "d", "e", "f", "g", "h")

	assert.Equal(t, head.Value, "a")
	head = head.Next
	assert.Equal(t, head.Value, "b")
	head = head.Next
	assert.Equal(t, head.Value, "c")
	head = head.Next
	assert.Equal(t, head.Value, "d")
	head = head.Next
	assert.Equal(t, head.Value, "e")
	head = head.Next
	assert.Equal(t, head.Value, "f")
	head = head.Next
	assert.Equal(t, head.Value, "g")
	head = head.Next
	assert.Equal(t, head.Value, "h")
	head = head.Next
	assert.Nil(t, head)
}
