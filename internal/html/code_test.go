package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`{code:xml}
	<test>
		 <another tag="attribute"/>
	</test>
{code}`)
	assert.NoError(err)
	assert.Equal("<code language=\"xml\">\n\t<test>\n\t\t <another tag=\"attribute\"/>\n\t</test>\n</code>", buf)
}
