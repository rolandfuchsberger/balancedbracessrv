package balancedbraces

import (
	"testing"
)

func TestStack(t *testing.T) {

	initst := "asd€èf"
	st := stack{
		items: []rune(initst),
	}

	pushRune := 't'
	st.push(pushRune)

	if string(st.items) != initst+string(pushRune) {
		t.Errorf("st=%q; st.push('t'); st=%q", initst, string(st.items))
	}

	stackString := string(st.items)

	r := []rune(stackString)
	for i := len(r) - 1; i >= 0; i-- {
		pop := st.pop()
		if pop != r[i] {
			t.Errorf("st.pop() = %q, expected: %q", pop, r[i])
		}
	}

	if st.len() != 0 {
		t.Errorf("st.len() = %v", st.len())
	}

	if p := (&stack{}).pop(); p != 0 {
		t.Errorf("stack.pop() = %q on empty stack", p)
	}

}
