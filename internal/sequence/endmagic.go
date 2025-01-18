package inisequence

const halfEndMagic rune = ']'

const endMagic1stRune rune = halfEndMagic
const endMagic2ndRune rune = halfEndMagic

// measured in bytes
const size1stEndMagic = len(string(endMagic1stRune))
const size2ndEndMagic = len(string(endMagic2ndRune))

const endMagicLen = size1stEndMagic + size2ndEndMagic
