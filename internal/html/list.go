package html

import (
	"regexp"
	"strings"
)

type listTransformer struct {
	re *regexp.Regexp
}

var _ Transformer = (*listTransformer)(nil)

func (r *listTransformer) toTag(sym string) string {
	switch sym {
	case "*", "-":
		return "ul"
	default:
		return "ol"
	}
}

func (r *listTransformer) Transform(in string) (string, error) {
	matches := r.re.FindAllString(in, -1)
	if len(matches) == 0 {
		return in, nil
	}
	toks := make([]string, 0)
	last := matches[0][0:1]
	first := last
	var lastsym string
	toks = append(toks, "<"+r.toTag(last)+">")
	for _, m := range matches {
		sym := m[1:2]
		offset := 2
		if sym != " " {
			// TODO: multi nested
			nextsym := r.toTag(sym)
			if nextsym != lastsym {
				lastsym = r.toTag(sym)
				toks = append(toks, "<"+lastsym+">")
			}
			offset = 3
		} else {
			if lastsym != "" {
				toks = append(toks, "</"+lastsym+">")
				lastsym = ""
			}
		}
		toks = append(toks, "<li>"+m[offset:]+"</li>")
	}
	toks = append(toks, "</"+r.toTag(first)+">")
	return strings.Join(toks, "\n"), nil
}

func init() {
	register(&listTransformer{regexp.MustCompile("(?m)^[\\*\\#\\-]{1,3} (.*)")})
}
