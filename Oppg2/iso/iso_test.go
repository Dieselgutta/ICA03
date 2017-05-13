package iso

import "testing"

func TestGreetingExtendedASCII(t *testing.T) {
	greetingEX := GreetingExtendedASCIIex()
	if isEASCII(greetingEX) == false {
		t.Fail()
	}
}

func isEASCII(q2 string) bool {
	for _, i := range q2 {
		if i < 128 {
			return false
		}
	}
	return true
}
