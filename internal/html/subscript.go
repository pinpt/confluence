package html

func init() {
	register(newRegexTransformer("^(?m)~([^\\n~]+)~", "<sub>$1</sub>"))
}
