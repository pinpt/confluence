package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCitation(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`??Hello??`)
	assert.NoError(err)
	assert.Equal("<cite>Hello</cite>", buf)
	buf, err = Parse("??Hello??\n")
	assert.NoError(err)
	assert.Equal("<cite>Hello</cite>\n", buf)
}
