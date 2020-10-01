package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreaksHR(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`----`)
	assert.NoError(err)
	assert.Equal("<hr>", buf)
	buf, err = Parse("----\n")
	assert.NoError(err)
	assert.Equal("<hr>\n", buf)
}

func TestBreaksMDash(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`---`)
	assert.NoError(err)
	assert.Equal("&mdash;", buf)
	buf, err = Parse("---\n")
	assert.NoError(err)
	assert.Equal("&mdash;\n", buf)
}

func TestBreaksNDash(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`--`)
	assert.NoError(err)
	assert.Equal("&ndash;", buf)
	buf, err = Parse("--\n")
	assert.NoError(err)
	assert.Equal("&ndash;\n", buf)
}

func TestBreaksCombined(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`----\nThis is a big dash ---.\nThis is a small dash --.\n`)
	assert.NoError(err)
	assert.Equal("<hr>\\nThis is a big dash &mdash;.\\nThis is a small dash &ndash;.\\n", buf)
}
