package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const paragraphBody = `Look at this bug in stable on Pinpoint customer:

[https://app.pinpoint.com/issue/7db68cb63ea90f2b/GOLD-367/Individual-Meeting-Hours-dont-match-up-to-team-total|https://app.pinpoint.com/issue/7db68cb63ea90f2b/GOLD-367/Individual-Meeting-Hours-dont-match-up-to-team-total]

Compare that to formatting in: [GOLD-367]

Looks like regressed again. 🎉`

const adfOutput = `<div class="source-jira"><p>Look at this bug in stable on Pinpoint customer:</p><p><a href="https://app.pinpoint.com/issue/7db68cb63ea90f2b/GOLD-367/Individual-Meeting-Hours-dont-match-up-to-team-total">https://app.pinpoint.com/issue/7db68cb63ea90f2b/GOLD-367/Individual-Meeting-Hours-dont-match-up-to-team-total</a></p><p>Compare that to formatting in: <a href="GOLD-367">GOLD-367</a></p><p>Looks like regressed again. 🎉</p></div>`

func TestParagraph(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(paragraphBody)
	assert.NoError(err)
	assert.Equal(adfOutput, buf)
}
