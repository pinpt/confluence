package html

func init() {
	register(newRegexTransformer("^(?m)-([^\\n\\-]+)-", "<s>$1</s>"))
}
