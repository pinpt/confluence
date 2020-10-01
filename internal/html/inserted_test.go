package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInserted(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`+Hello+`)
	assert.NoError(err)
	assert.Equal("<ins>Hello</ins>", buf)
	buf, err = Parse("+Hello+\n")
	assert.NoError(err)
	assert.Equal("<ins>Hello</ins>\n", buf)
}
