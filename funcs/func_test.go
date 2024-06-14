package funcs

import "testing"

func TestTrimSpaceAndConvInt(t *testing.T) {
	cases := []struct {
		name     string
		str      string
		expected string
	}{
		{name: "valid case", str: "HellO wOrlD", expected: "helloworld"},
		{name: "valid case", str: "   AdMIN1", expected: "adminone"},
		{name: "valid case", str: "14 3   9", expected: "onefourthreenine"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := TrimSpaceAndConvInt(c.str)
			if got != c.expected {
				t.Errorf("expected: %s, but got: %s\n", c.expected, got)
			}
		})
	}
}
