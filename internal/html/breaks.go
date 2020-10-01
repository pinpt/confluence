package html

func init() {
	register(newRegexTransformer("(?m)----", "<hr>"))
	register(newRegexTransformer("(?m)---", "&mdash;"))
	register(newRegexTransformer("(?m)--", "&ndash;"))
}
