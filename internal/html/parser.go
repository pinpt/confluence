package html

import (
	"fmt"
	"regexp"
	"strings"
)

// SEE https://jira.atlassian.com/secure/WikiRendererHelpAction.jspa?section=all

// Transformer will transform content
type Transformer interface {
	Transform(string) (string, error)
}

type regexTransformer struct {
	re   *regexp.Regexp
	repl string
}

var _ Transformer = (*regexTransformer)(nil)

func (r *regexTransformer) Transform(in string) (string, error) {
	return r.re.ReplaceAllString(in, r.repl), nil
}

func newRegexTransformer(pattern string, repl string) *regexTransformer {
	return &regexTransformer{
		re:   regexp.MustCompile(pattern),
		repl: repl,
	}
}

var transformers = make([]Transformer, 0)

func register(t Transformer) {
	transformers = append(transformers, t)
}

// Parse will parse input and return output
func Parse(buf string) (string, error) {
	var input = buf
	for _, t := range transformers {
		output, err := t.Transform(input)
		if err != nil {
			return "", err
		}
		input = output
	}
	return wrap(input), nil
}

func wrap(output string) string {
	var str string
	paragraphs := strings.Split(output, "\n\n")
	if len(paragraphs) == 1 {
		str = output
	} else {
		for _, p := range paragraphs {
			str += fmt.Sprintf(`<p>%s</p>`, p)
		}
	}
	return fmt.Sprintf(`<div class="source-jira">%s</div>`, str)
}
