package ascii

import (
	"testing"
	"unicode"
)

func TestGreetingASCII(t *testing.T) {
	a := GreetingASCII()
	if isAscii(a) == false {
		t.Fail()
	}
}

// Original fra https://play.golang.org/p/hnZzfnbXeF
func isAscii(a string) bool {
	for _, c := range a {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}
