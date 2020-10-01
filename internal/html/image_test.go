package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleImage(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`!http://www.host.com/image.gif!`)
	assert.NoError(err)
	assert.Equal("<img src=\"http://www.host.com/image.gif\">", buf)
}

func TestNamedImage(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`!http://www.host.com/image.gif|thumbnail!`)
	assert.NoError(err)
	assert.Equal("<img src=\"http://www.host.com/image.gif\" title=\"thumbnail\">", buf)
}

func TestNamedImageWithAttributes(t *testing.T) {
	t.Skip()
	assert := assert.New(t)
	buf, err := Parse(`!!image.gif|align=right, vspace=4!`)
	assert.NoError(err)
	assert.Equal("<img src=\"image.gif\">", buf)
}
