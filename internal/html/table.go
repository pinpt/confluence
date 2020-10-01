package html

import (
	"regexp"
	"strings"
)

type tableTransformer struct {
	re *regexp.Regexp
}

var _ Transformer = (*tableTransformer)(nil)

func (r *tableTransformer) Transform(in string) (string, error) {
	matches := r.re.FindAllString(in, -1)
	if len(matches) == 0 {
		return in, nil
	}
	toks := make([]string, 0)
	toks = append(toks, "<table>")
	headers := matches[0]
	headersTok := strings.Split(headers, "||")
	headertoks := make([]string, 0)
	headertoks = append(headertoks, "<tr>")
	for _, h := range headersTok {
		headertoks = append(headertoks, "<th>"+strings.TrimSpace(h)+"</th>")
	}
	headertoks = append(headertoks, "</tr>")
	toks = append(toks, strings.Join(headertoks, ""))
	for _, rowtok := range matches[1:] {
		row := make([]string, 0)
		row = append(row, "<tr>")
		rowsTok := strings.Split(rowtok, "|")
		for _, t := range rowsTok {
			row = append(row, "<td>"+t+"</td>")
		}
		row = append(row, "</tr>")
		toks = append(toks, strings.Join(row, ""))
	}
	toks = append(toks, "</table>")
	return strings.Join(toks, "\n"), nil
}

func init() {
	register(&tableTransformer{regexp.MustCompile("(?m)^\\|(.*)\\|(.*)\\|")})
}
