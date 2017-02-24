package ascii

import "fmt"

func IterateOverASCIIStringLiteral(s string) {
	for i := 0; i < len(s); i++ {
		a := s[i]
		fmt.Printf("%x %q %b\n", a, a, a)

		//fmt.Printf("%x ", a)
		//fmt.Printf("%c ", a)
		//fmt.Printf("%b\n ", a)
	}
}

func GreetingASCII() string {
	b := []byte("\x48\x65\x6c\x6c\x6f\x20\x3a\x2d\x29")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%c", b[i])
	}
	c := "\x48\x65\x6c\x6c\x6f\x20\x3a\x2d\x29"
	return c
}
