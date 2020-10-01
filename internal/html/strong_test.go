package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrong(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`*Hello*`)
	assert.NoError(err)
	assert.Equal("<strong>Hello</strong>", buf)
	buf, err = Parse("*Hello*\n")
	assert.NoError(err)
	assert.Equal("<strong>Hello</strong>\n", buf)
}
