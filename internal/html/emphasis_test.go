package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmphasis(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`_Hello_`)
	assert.NoError(err)
	assert.Equal("<em>Hello</em>", buf)
	buf, err = Parse("_Hello_\n")
	assert.NoError(err)
	assert.Equal("<em>Hello</em>\n", buf)
}
