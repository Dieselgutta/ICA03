package ascii

import "testing"

func TestGreetingsASCII(t *testing.T) {
	ascii := GreetingASCII()
	if !(isASCII(ascii)) {
		t.Fail()
	}
}

func isASCII(s string) bool {
	for _, C := range s {
		if C > 127 {
			return false
		}
	}
	return true
}
