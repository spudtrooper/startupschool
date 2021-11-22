package startupschool

import "regexp"

var (
	urlRE = regexp.MustCompile(`(https?://)([\.a-zA-Z_0-9/\-]+)`)
)

func linkify(text string) string {
	html := text
	html = urlRE.ReplaceAllString(html, `<a target="_" href="$1$2">$2</a>`)
	return html
}
