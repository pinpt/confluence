package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBulletList(t *testing.T) {
	assert := assert.New(t)
	content := `
* some
* bullet
* points
`
	buf, err := Parse(content)
	assert.NoError(err)
	assert.Equal("<div class=\"source-jira\"><ul>\n<li>some</li>\n<li>bullet</li>\n<li>points</li>\n</ul></div>", buf)
}

func TestBulletListNested(t *testing.T) {
	assert := assert.New(t)
	content := `
* some
* bullet
** indented
** bullets
* points
`
	buf, err := Parse(content)
	assert.NoError(err)
	assert.Equal("<div class=\"source-jira\"><ul>\n<li>some</li>\n<li>bullet</li>\n<ul>\n<li>indented</li>\n<li>bullets</li>\n</ul>\n<li>points</li>\n</ul></div>", buf)
}

func TestNumberedList(t *testing.T) {
	assert := assert.New(t)
	content := `
# a
# numbered
`
	buf, err := Parse(content)
	assert.NoError(err)
	assert.Equal("<div class=\"source-jira\"><ol>\n<li>a</li>\n<li>numbered</li>\n</ol></div>", buf)
}

func TestMixedList(t *testing.T) {
	assert := assert.New(t)
	content := `
# a
# numbered
#* with
#* nested
#* bullet
# list
`
	buf, err := Parse(content)
	assert.NoError(err)
	assert.Equal("<div class=\"source-jira\"><ol>\n<li>a</li>\n<li>numbered</li>\n<ul>\n<li>with</li>\n<li>nested</li>\n<li>bullet</li>\n</ul>\n<li>list</li>\n</ol></div>", buf)
}

func TestMixedList2(t *testing.T) {
	assert := assert.New(t)
	content := `
* a
* bulleted
*# with
*# nested
*# numbered
* list
`
	buf, err := Parse(content)
	assert.NoError(err)
	assert.Equal("<div class=\"source-jira\"><ul>\n<li>a</li>\n<li>bulleted</li>\n<ol>\n<li>with</li>\n<li>nested</li>\n<li>numbered</li>\n</ol>\n<li>list</li>\n</ul></div>", buf)
}

const listtext = `
*
`

func TestMocklist(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(listtext)
	assert.NoError(err)
	assert.Equal("<div class=\"source-jira\">\n*\n</div>", buf)
}
