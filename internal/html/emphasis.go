package html

func init() {
	register(newRegexTransformer("^(?m)\\_([^\\n\\_]+)\\_", "<em>$1</em>"))
}
