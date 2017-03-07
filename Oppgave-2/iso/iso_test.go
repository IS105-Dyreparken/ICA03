package iso

import "testing"

func TestGreetingExtendedASCII(t *testing.T) {
	a := GreetingExtendedASCII()
	if isExAscii(a) == false {
		t.Fail()
	}
}

func isExAscii(a string) bool {
	for _, c := range a {
		if (c > 127) && (c < 256) {
			return true
		}
	}
	return false
}
