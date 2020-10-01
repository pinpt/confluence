package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscript(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`~Hello~`)
	assert.NoError(err)
	assert.Equal("<sub>Hello</sub>", buf)
	buf, err = Parse("~Hello~\n")
	assert.NoError(err)
	assert.Equal("<sub>Hello</sub>\n", buf)
}
