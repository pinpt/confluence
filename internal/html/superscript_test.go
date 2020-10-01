package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperscript(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`^Hello^`)
	assert.NoError(err)
	assert.Equal("<sup>Hello</sup>", buf)
	buf, err = Parse("^Hello^\n")
	assert.NoError(err)
	assert.Equal("<sup>Hello</sup>\n", buf)
}
