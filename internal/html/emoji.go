package html

func init() {
	register(newRegexTransformer("^:\\)$", "😃"))
	register(newRegexTransformer("^:\\($", "🙁"))
	register(newRegexTransformer("^:P$", "😛"))
	register(newRegexTransformer("^:D$", "😁"))
	register(newRegexTransformer("^;\\)$", "😉"))
	register(newRegexTransformer("^\\(y\\)$", "👍"))
	register(newRegexTransformer("^\\(n\\)$", "👎"))
	register(newRegexTransformer("^\\(i\\)$", "ℹ️"))
	register(newRegexTransformer("^\\(\\/\\)$", "✅"))
	register(newRegexTransformer("^\\(x\\)$", "❗️"))
	register(newRegexTransformer("^\\(\\!\\)$", "⚠️"))
	register(newRegexTransformer("^\\(\\+\\)$", "➕"))
	register(newRegexTransformer("^\\(\\-\\)$", "➖"))
	register(newRegexTransformer("^\\(\\?\\)$", "❓"))
}
