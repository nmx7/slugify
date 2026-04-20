package slugify

import (
	"regexp"
	"strings"
	"unicode"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Global replacer for high performance.
// Includes common expansions for German, Nordic, and Slavic characters.
var charExpansion = strings.NewReplacer(
    // German
    "ä", "ae", "ö", "oe", "ü", "ue",
    "Ä", "ae", "Ö", "oe", "Ü", "ue",
    "ß", "ss",
    // Nordic & Scandinavian
    "æ", "ae", "Æ", "ae",
    "ø", "oe", "Ø", "oe", // 'oe' is common for Ø, but 'o' is also used
    "å", "a",  "Å", "a",
    // Icelandic
    "ð", "d",  "Ð", "d",
    "þ", "th", "Þ", "th",
    // Polish & Slavic
    "ł", "l",  "Ł", "l",
    "ć", "c",  "Ć", "c",
    "ś", "s",  "Ś", "s",
    "ź", "z",  "Ź", "z",
    "ż", "z",  "Ż", "z",
    // Cyrillic Transliteration (Common chars)
    "а", "a", "б", "b", "в", "v", "г", "g", "д", "d",
    "е", "e", "ё", "yo", "ж", "zh", "з", "z", "и", "i",
    "й", "j", "к", "k", "л", "l", "м", "m", "н", "n",
    "о", "o", "п", "p", "р", "r", "с", "s", "т", "t",
    "у", "u", "ф", "f", "х", "kh", "ц", "ts", "ч", "ch",
    "ш", "sh", "щ", "shch", "ы", "y", "э", "e", "ю", "yu", "я", "ya",
    // Symbols & Currencies
    "&", " and ",
    "@", " at ",
    "€", " euro ",
    "£", " pound ",
    "$", " dollar ",
    "%", " percent ",
)

var nonAlphaNumeric = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func Slugify(s string) string {
	s = charExpansion.Replace(s)

	// Normalize to NFD and strip remaining accents (like é, ñ, etc.)
	// This covers everything not explicitly in the replacer map.
	t := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		norm.NFC,
	)

	s, _,_ = transform.String(t,s)

	// Replace non-alphanumeric characters with spaces
	s = nonAlphaNumeric.ReplaceAllString(s, " ")

	s = strings.ToLower(s)

	// Trim leading/trailing spaces and replace internal spaces with hyphens
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "-")

	return s
}
