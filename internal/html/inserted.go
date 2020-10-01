package html

func init() {
	register(newRegexTransformer("^(?m)\\+([^\\n\\+]+)\\+", "<ins>$1</ins>"))
}
