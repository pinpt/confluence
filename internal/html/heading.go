package html

func init() {
	register(newRegexTransformer("^(?m)^h([1-6])\\.\\s+([^\\n]+)", "<h$1>$2</h$1>"))
}
