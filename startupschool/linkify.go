package startupschool

import "regexp"

var (
	screenNameRE = regexp.MustCompile(`@([a-zA-Z_0-9]+)`)
	urlRE        = regexp.MustCompile(`(https?://)([\.a-zA-Z_0-9/\-]+)`)
)

func linkify(text string) string {
	html := text
	html = urlRE.ReplaceAllString(html, `<a target="_" href="$1$2">$2</a>`)
	html = screenNameRE.ReplaceAllString(html, `<a target="_" href="https://twitter.com/$1">@$1</a>`)
	return html
}
