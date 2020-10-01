package html

func init() {
	register(newRegexTransformer("^(?sm)\\{code:(\\w+)\\}(.*)\\{code\\}", "<code language=\"$1\">$2</code>"))
}
