package html

func init() {
	register(newRegexTransformer("^(?m)\\[([^\\|]+)\\|([^\\n\\]]+)\\]", "<a href=\"$2\">$1</a>"))
	register(newRegexTransformer("^(?m)\\[([^\\n\\]]+)\\]", "<a href=\"$1\">$1</a>"))
}
