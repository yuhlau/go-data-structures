package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceStack(t *testing.T) {
	assert := assert.New(t)

	stack := NewSliceStack()
	assert.Equal(-1, stack.top)
	assert.Equal(0, len(stack.data))
	assert.Equal(SLICESTACK_DEFAULT_CAP, cap(stack.data))
}

func TestSliceStackWithDefaultCap(t *testing.T) {
	assert := assert.New(t)

	stack := NewSliceStackWithDefaultCap(100)
	assert.Equal(-1, stack.top)
	assert.Equal(0, len(stack.data))
	assert.Equal(100, cap(stack.data))
}

func TestPush(t *testing.T) {
	assert := assert.New(t)

	stack := NewSliceStack()

	stack.Push(1) // [1]
	assert.Equal(1, stack.data[0])
	assert.Equal(0, stack.top)

	stack.Push(2)
	stack.Push(3) // [1 2 3]
	assert.Equal(1, stack.data[0])
	assert.Equal(2, stack.data[1])
	assert.Equal(3, stack.data[2])
	assert.Equal(2, stack.top)
}

func TestPop(t *testing.T) {
	assert := assert.New(t)

	stack := NewSliceStack()

	val, err := stack.Pop()
	assert.Nil(val)
	assert.Equal("Stack is empty", err.Error())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3) // [1 2 3]

	val, err = stack.Pop()
	assert.Equal(3, val)
	assert.Nil(err)
	assert.Equal(uint(2), stack.Size())

	stack.Push(5) // [1 2 5]
	val, err = stack.Pop()
	assert.Equal(5, val)
	assert.Nil(err)
	assert.Equal(uint(2), stack.Size())
}

func TestTop(t *testing.T) {
	assert := assert.New(t)

	stack := NewSliceStack()

	val, err := stack.Top()
	assert.Nil(val)
	assert.Equal("Stack is empty", err.Error())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3) // [1 2 3]

	val, err = stack.Top()
	assert.Equal(3, val)
	assert.Nil(err)
	assert.Equal(uint(3), stack.Size())

	stack.Push(5) // [1 2 3 5]
	val, err = stack.Top()
	assert.Equal(5, val)
	assert.Nil(err)
	assert.Equal(uint(4), stack.Size())
}

func TestIsEmpty(t *testing.T) {
	assert := assert.New(t)

	stack := NewSliceStack()
	assert.Equal(true, stack.IsEmpty())

	stack.Push(1)
	assert.Equal(false, stack.IsEmpty())
}

func TestSize(t *testing.T) {
	assert := assert.New(t)

	stack := NewSliceStack()
	assert.Equal(uint(0), stack.Size())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3) // [1 2 3]
	assert.Equal(uint(3), stack.Size())
}
