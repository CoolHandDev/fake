package fake

var lowerLetters = []rune("abcdefghijklmnopqrstuvwxyz")
var upperLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numeric = []rune("0123456789")
var specialChars = []rune(`!'@#$%^&*()_+-=[]{};:",./?`)
var hexDigits = []rune("0123456789abcdef")

func text(atLeast, atMost int, allowLower, allowUpper, allowNumeric, allowSpecial bool) string {
	allowedChars := []rune{}
	if allowLower {
		allowedChars = append(allowedChars, lowerLetters...)
	}
	if allowUpper {
		allowedChars = append(allowedChars, upperLetters...)
	}
	if allowNumeric {
		allowedChars = append(allowedChars, numeric...)
	}
	if allowSpecial {
		allowedChars = append(allowedChars, specialChars...)
	}

	result := []rune{}
	nTimes := r.Intn(atMost-atLeast+1) + atLeast
	for i := 0; i < nTimes; i++ {
		result = append(result, allowedChars[r.Intn(len(allowedChars))])
	}
	return string(result)
}

func Password(atLeast, atMost int, allowUpper, allowNumeric, allowSpecial bool) string {
	return text(atLeast, atMost, true, allowUpper, allowNumeric, allowSpecial)
}

func SimplePassword() string {
	return Password(6, 12, true, true, false)
}

func Color() string {
	return lookup(lang, "colors", true)
}

func hexDigitsStr(n int) string {
	var num []rune
	for i := 0; i < n; i++ {
		num = append(num, hexDigits[r.Intn(len(hexDigits))])
	}
	return string(num)
}

func HexColor() string {
	return hexDigitsStr(6)
}

func ShortHexColor() string {
	return hexDigitsStr(3)
}
