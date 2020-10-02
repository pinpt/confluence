package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleLink(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`[http://jira.atlassian.com]`)
	assert.NoError(err)
	assert.Equal("<a href=\"http://jira.atlassian.com\">http://jira.atlassian.com</a>", buf)
}
func TestNamedLink(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`[Atlassian|http://jira.atlassian.com]`)
	assert.NoError(err)
	assert.Equal("<a href=\"http://jira.atlassian.com\">Atlassian</a>", buf)
}

func TestNamedLinkWithNewLing(t *testing.T) {
	assert := assert.New(t)
	v := "\n\n[https://app.pinpoint.com/issue/7db68cb63ea90f2b/GOLD-367/Individual-Meeting-Hours-dont-match-up-to-team-total|https://app.pinpoint.com/issue/7db68cb63ea90f2b/GOLD-367/Individual-Meeting-Hours-dont-match-up-to-team-total]\n\nCompare that to formatting in: [GOLD-367]\n\n"
	buf, err := Parse(v)
	assert.NoError(err)
	assert.Equal("\n\n<a href=\"https://app.pinpoint.com/issue/7db68cb63ea90f2b/GOLD-367/Individual-Meeting-Hours-dont-match-up-to-team-total\">https://app.pinpoint.com/issue/7db68cb63ea90f2b/GOLD-367/Individual-Meeting-Hours-dont-match-up-to-team-total</a>\n\nCompare that to formatting in: <a href=\"GOLD-367\">GOLD-367</a>\n\n", buf)
}
