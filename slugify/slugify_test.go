package slugify

import "testing"

func TestSlugify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"German/Nordic", "München & København!", "muenchen-and-koebenhavn"},
		{"Icelandic", "Reykjavík & Þingvellir", "reykjavik-and-thingvellir"},
		{"French/Spanish", "L'été au bord du Niágara", "l-ete-au-bord-du-niagara"},
		{"Polish", "Zażółć gęślą jaźń", "zazolc-gesla-jazn"},
		{"Punctuation Soup", "Go -- Lang --- !!Slug!!", "go-lang-slug"},
		{"Currency/Symbols", "Buy for $50 & get 10% off", "buy-for-50-dollar-and-get-10-percent-off"},
		{"Emoji/Whitespace", " I love Go! 🚀 ", "i-love-go"},
		{"Numbers Only", "123 456!!!", "123-456"},
		{"Empty/Invalid", " @#$% ", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Slugify(tt.input)
			if actual != tt.expected {
				t.Errorf("Slugify(%q) = %q; want %q", tt.input, actual, tt.expected)
			}
		})
	}
}
