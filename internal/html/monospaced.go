package html

func init() {
	register(newRegexTransformer("^(?m){{([^\\n}]+)}}", "<tt>$1</tt>"))
}
