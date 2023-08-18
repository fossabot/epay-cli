package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPtr(t *testing.T) {
	asserts := assert.New(t)

	asserts.Equal(1, *Ptr(1))
	asserts.Equal("a", *Ptr("a"))
	asserts.Equal(true, *Ptr(true))
	asserts.Equal(1.1, *Ptr(1.1))
}
