package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	assert := assert.New(t)
	// Original text
	buf, err := Parse(`
||heading 1||heading 2||heading 3||
|col A1|col A2|col A3|
|col B1|col B2|col B3|
	`)
	assert.NoError(err)
	assert.Equal("<div class=\"source-jira\">\n<table>\n<tr><th></th><th>heading 1</th><th>heading 2</th><th>heading 3</th><th></th></tr>\n<tr><td></td><td>col A1</td><td>col A2</td><td>col A3</td><td></td></tr>\n<tr><td></td><td>col B1</td><td>col B2</td><td>col B3</td><td></td></tr>\n</table>\n\t</div>", buf)

	// Only text without another symbols
	buf, err = Parse(`||heading 1||heading 2||heading 3||
|col A1|col A2|col A3|
|col B1|col B2|col B3|`)
	assert.NoError(err)
	assert.Equal("<div class=\"source-jira\"><table>\n<tr><th></th><th>heading 1</th><th>heading 2</th><th>heading 3</th><th></th></tr>\n<tr><td></td><td>col A1</td><td>col A2</td><td>col A3</td><td></td></tr>\n<tr><td></td><td>col B1</td><td>col B2</td><td>col B3</td><td></td></tr>\n</table></div>", buf)

	// Multiple tables, text before, between and after tables
	buf, err = Parse(`
some text before
||heading 1||heading 2||heading 3||
|col A1|col A2|col A3|
|col B1|col B2|col B3|
table duplicate
||heading 1||heading 2||heading 3||
|col A1|col A2|col A3|
|col B1|col B2|col B3|
another table
||heading 4||heading 5||heading 6||
|col A1|col A2|col A3|
|col B1|col B2|col B3|
some text after
	`)
	assert.NoError(err)
	assert.Equal("<div class=\"source-jira\">\nsome text before\n<table>\n<tr><th></th><th>heading 1</th><th>heading 2</th><th>heading 3</th><th></th></tr>\n<tr><td></td><td>col A1</td><td>col A2</td><td>col A3</td><td></td></tr>\n<tr><td></td><td>col B1</td><td>col B2</td><td>col B3</td><td></td></tr>\n</table>\ntable duplicate\n<table>\n<tr><th></th><th>heading 1</th><th>heading 2</th><th>heading 3</th><th></th></tr>\n<tr><td></td><td>col A1</td><td>col A2</td><td>col A3</td><td></td></tr>\n<tr><td></td><td>col B1</td><td>col B2</td><td>col B3</td><td></td></tr>\n</table>\nanother table\n<table>\n<tr><th></th><th>heading 4</th><th>heading 5</th><th>heading 6</th><th></th></tr>\n<tr><td></td><td>col A1</td><td>col A2</td><td>col A3</td><td></td></tr>\n<tr><td></td><td>col B1</td><td>col B2</td><td>col B3</td><td></td></tr>\n</table>\nsome text after\n\t</div>", buf)
}
