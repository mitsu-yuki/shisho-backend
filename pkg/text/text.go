package text

import "unicode"

func IsKatakana(s string) bool {
	for _, c := range s {
		if !unicode.In(c, unicode.Katakana) {
			return false
		}
	}
	return true
}

