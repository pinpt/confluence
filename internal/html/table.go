package html

import (
	"regexp"
	"strings"
)

type tableTransformer struct {
	re *regexp.Regexp
}

var _ Transformer = (*tableTransformer)(nil)
var getTableFields = regexp.MustCompile("(?m)^\\|(.*)\\|(.*)\\|")

func (r *tableTransformer) Transform(in string) (string, error) {
	tables := r.re.FindAllString(in + "\n", -1)
	if len(tables) == 0 {
		return in, nil
	}
	for _, table := range tables {
		table = strings.TrimRight(table, "\n")
		matches := getTableFields.FindAllString(table, -1)
		if len(matches) == 0 {
			continue
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
		in = strings.Replace(in, table, strings.Join(toks, "\n"), -1)
	}
	return in, nil
}

func init() {
	register(&tableTransformer{regexp.MustCompile("(?sm)(\\|\\|.*?\\|\\|\\s*(?:\\|.*?\\|\\n)+)")})
}
