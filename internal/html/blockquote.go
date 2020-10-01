package html

func init() {
	register(newRegexTransformer("^(?m)bq\\.\\s+([^\\n]+)", "<blockquote>$1</blockquote>"))
	register(newRegexTransformer("^(?sm)\\{quote\\}(.*)\\{quote\\}", "<blockquote>$1</blockquote>"))
}
