package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeadings(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`h1. Hello`)
	assert.NoError(err)
	assert.Equal("<h1>Hello</h1>", buf)
	buf, err = Parse("h1. Hello\n")
	assert.NoError(err)
	assert.Equal("<h1>Hello</h1>\n", buf)
}
