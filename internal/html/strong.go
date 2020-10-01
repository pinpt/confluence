package html

func init() {
	register(newRegexTransformer("^(?m)\\*([^\\n\\*]+)\\*", "<strong>$1</strong>"))
}
