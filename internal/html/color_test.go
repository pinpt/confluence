package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`{color:red}Hello{color}`)
	assert.NoError(err)
	assert.Equal("<span style=\"color:red\">Hello</span>", buf)
	buf, err = Parse("{color:red}\nHello\n{color}\n")
	assert.NoError(err)
	assert.Equal("<span style=\"color:red\">\nHello\n</span>\n", buf)
}
