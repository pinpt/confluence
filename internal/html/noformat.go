package html

func init() {
	register(newRegexTransformer("^(?sm)\\{noformat\\}(.*)\\{noformat\\}", "<pre>$1</pre>"))
}
