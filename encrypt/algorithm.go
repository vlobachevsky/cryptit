package encrypt

func Nimbus(str string) string {
	encryptedString := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode + 3)
		encryptedString += character
	}
	return encryptedString
}