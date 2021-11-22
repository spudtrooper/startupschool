package startupschool

import (
	"testing"
)

func TestLinkify(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty",
			input: "",
			want:  "",
		},
		{
			name:  "no links",
			input: "blah blah blah",
			want:  "blah blah blah",
		},
		{
			name:  "name link",
			input: "@bob",
			want:  `<a target="_" href="https://twitter.com/bob">@bob</a>`,
		},
		{
			name:  "http link",
			input: "http://foo-bar_baz.com",
			want:  `<a target="_" href="http://foo-bar_baz.com">foo-bar_baz.com</a>`,
		},
		{
			name:  "https link",
			input: "https://foo.com",
			want:  `<a target="_" href="https://foo.com">foo.com</a>`,
		},
		{
			name:  "LiN_K",
			input: "@B_oo00oBB",
			want:  `<a target="_" href="https://twitter.com/B_oo00oBB">@B_oo00oBB</a>`,
		},
		{
			name:  "links",
			input: "@bob https://baz.com @sue http://foo.com",
			want:  `<a target="_" href="https://twitter.com/bob">@bob</a> <a target="_" href="https://baz.com">baz.com</a> <a target="_" href="https://twitter.com/sue">@sue</a> <a target="_" href="http://foo.com">foo.com</a>`,
		},
		{
			name:  "links no space",
			input: "@bob@sue",
			want:  `<a target="_" href="https://twitter.com/bob">@bob</a><a target="_" href="https://twitter.com/sue">@sue</a>`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := linkify(test.input); test.want != got {
				t.Errorf("linkify(%v): want %v, got %v", test.input, test.want, got)
			}
		})
	}
}
