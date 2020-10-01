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
