package html

func init() {
	register(newRegexTransformer("^(?s)\\{color:(\\w+)\\}(.*)\\{color\\}", "<span style=\"color:$1\">$2</span>"))
}
