package ascii

import (
	"testing"
	"unicode"
)

func TestGreetingASCII(t *testing.T) {
	a1 := GreetingASCII()
	if isASCII(b) == false {
		t.Errorf("FAIL")
	}
}

func isASCII(a1 string) bool {
	for _, i := range a1 {
		if i > unicode.MaxASCII {
			return false
		}
	}
	return true
}
