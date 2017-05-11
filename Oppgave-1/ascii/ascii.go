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
	t1 := []byte("\x48\x65\x6c\x6c\x6f\x20\x3a\x2d\x29")
	for i := 0; i < len(t1); i++ {
		fmt.Printf("%c", t1[i])
	}
	t2 := "\x48\x65\x6c\x6c\x6f\x20\x3a\x2d\x29"
	return t2
}
