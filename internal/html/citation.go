package html

func init() {
	register(newRegexTransformer("^(?m)\\?\\?([^\\n\\?]+)\\?\\?", "<cite>$1</cite>"))
}
