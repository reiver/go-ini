package inisequence

const halfMagic rune = '['

const magic1stRune rune = halfMagic
const magic2ndRune rune = halfMagic

// measured in bytes
const Size1stMagic = len(string(magic1stRune))
const Size2ndMagic = len(string(magic2ndRune))

const magicLen = Size1stMagic + Size2ndMagic

// You can tell whether a line in an INI file is a sequence just based on the first character and the second character.
//
// If the first character is a '[', then it is a section or a sequence.
//
// If the second character is a '[' (in addition to the first character being a '['), then it is a sequence.
func Is1stMagic(r rune) bool {
	return isMagic(r)
}

// You can tell whether a line in an INI file is a sequence just based on the first character and the second character.
//
// If the first character is a '[', then it is a section or a sequence.
//
// If the second character is a '[' (in addition to the first character being a '['), then it is a sequence.
func Is2ndMagic(r rune) bool {
	return isMagic(r)
}

func isMagic(r rune) bool {
	switch r {
	case halfMagic:
		return true
	default:
		return false
	}
}
