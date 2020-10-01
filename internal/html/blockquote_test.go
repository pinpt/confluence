package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockquote(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`bq. Hello`)
	assert.NoError(err)
	assert.Equal("<blockquote>Hello</blockquote>", buf)
	buf, err = Parse("bq. Hello\n")
	assert.NoError(err)
	assert.Equal("<blockquote>Hello</blockquote>\n", buf)
}

func TestBlockquoteMultiline(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`{quote}Hello{quote}`)
	assert.NoError(err)
	assert.Equal("<blockquote>Hello</blockquote>", buf)
	buf, err = Parse("{quote}\nHello\n{quote}\n")
	assert.NoError(err)
	assert.Equal("<blockquote>\nHello\n</blockquote>\n", buf)
}
