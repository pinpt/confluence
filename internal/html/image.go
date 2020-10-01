package html

func init() {
	register(newRegexTransformer("^(?m)\\!([^\\!]+)\\|([^\\n\\]]+)\\!", "<img src=\"$1\" title=\"$2\">"))
	register(newRegexTransformer("^(?m)\\!([^\\n\\!]+)\\!", "<img src=\"$1\">"))
}
