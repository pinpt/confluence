package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoformat(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`{noformat}
preformatted piece of text
so *no* further _formatting_ is done here
{noformat}`)
	assert.NoError(err)
	assert.Equal("<pre>\npreformatted piece of text\nso *no* further _formatting_ is done here\n</pre>", buf)
}
