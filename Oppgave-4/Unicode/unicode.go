package unicode

var langTrans string

// Kode for Oppgave 4a
func Translate(expression, language string) string {
	if language == "jp" && expression == "nord og sør" {
		langTrans := "\xE2\x80\x9C" + expression + "\xE2\x80\x9D" + " på japansk er " + "\xE2\x80\x9C\xE5\x8C\x97\xE3\x81\xA8\xE5\x8D\x97\xE2\x80\x9D\n"
		return langTrans
	}
	if language == "is" && expression == "nord og sør" {
		langTrans := "\xE2\x80\x9C" + expression + "\xE2\x80\x9D" + " på islandsk er " + "\xE2\x80\x9C\x6E\x6F\x72\xC3\xB0\x75\x72\x20\x6F\x67\x20\x73\x75\xC3\xB0\x75\x72\xE2\x80\x9D\n"
		return langTrans
	}
	return langTrans
}
