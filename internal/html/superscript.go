package html

func init() {
	register(newRegexTransformer("^(?m)\\^([^\\n\\^]+)\\^", "<sup>$1</sup>"))
}
