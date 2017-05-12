package iso

import "fmt"

func IterateOverASCIIStringLiteral(s string) {
	for i := 0; i < len(s); i++ {
		a := s[i]
		fmt.Printf("%x %q %b \n", a, a, a)
	}
}

// Kode for Oppgave 2b
const x = "\x22\x53\x61\x6c\x75\x74\x2c\x20\xe7\x61\x20\x76\x61\x20\xb0\x2d\x29\x20\x80\x35\x30\x22"

//const a = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func GreetingExtendedASCII() string {
	b := []byte(x)
	for i := 0; i < len(b); i++ {
		fmt.Printf("%c", b[i])
	}
	return x
}
