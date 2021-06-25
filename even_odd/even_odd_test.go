package even_odd

import (
	"testing"
	//"github.com/edm20627/gopherdojo-studyroom/kadai1/edm20627/even_odd"
)

func TestIsEven(t *testing.T) {
	cases := []struct {
		name     string
		input    int
		expected bool
	}{
		{name: "+odd", input: 5, expected: false},
		{name: "+even", input: 6, expected: true},
		{name: "-odd", input: -5, expected: false},
		{name: "-even", input: -6, expected: true},
		{name: "zero", input: 0, expected: true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := IsEven(c.input); c.expected != actual {
				t.Errorf("want IsOdd(%d) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}
}
