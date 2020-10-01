package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`
||heading 1||heading 2||heading 3||
|col A1|col A2|col A3|
|col B1|col B2|col B3|
	`)
	assert.NoError(err)
	assert.Equal("<table>\n<tr><th></th><th>heading 1</th><th>heading 2</th><th>heading 3</th><th></th></tr>\n<tr><td></td><td>col A1</td><td>col A2</td><td>col A3</td><td></td></tr>\n<tr><td></td><td>col B1</td><td>col B2</td><td>col B3</td><td></td></tr>\n</table>", buf)
}
