package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonospaced(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`{{Hello}}`)
	assert.NoError(err)
	assert.Equal("<tt>Hello</tt>", buf)
	buf, err = Parse("{{Hello}}\n")
	assert.NoError(err)
	assert.Equal("<tt>Hello</tt>\n", buf)
}
