package files

import (
	"os"
	"unicode"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}

	return true
}

func GetFileName(s string) string {
	var n string

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) || unicode.IsLetter(rune(s[i])) || unicode.IsSpace(rune(s[i])) {
			n += string(s[i])
		}
	}

	return n
}
