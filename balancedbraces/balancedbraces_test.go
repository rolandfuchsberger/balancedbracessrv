package balancedbraces_test

import (
	"testing"

	"fuchsberger.email/balancedbracessrv/balancedbraces"
)

func TestBalancedBraces(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{"asdf", true},
		{")(}", false},
		{"a(sdf", false},
		{"as(d]f", false},
		{"([])", true},
		{"({[asdf]})", true},
		{"([asdf/]})", false},
		{"({[asdf]}", false},
		{"({[[{(asdf", false},
		{"(", false},
		{"", true},
		{"(/){}[]", true},
		{"(){}[", false},
		{"(){[]", false},
		{"({}[]", false},
	}

	for _, test := range tests {

		if balancedbraces.BalancedBraces(test.input) != test.want {
			t.Errorf("BalancedBraces(%q) = %v", test.input, !test.want)
		}

	}
}
