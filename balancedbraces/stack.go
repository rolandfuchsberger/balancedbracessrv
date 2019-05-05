package balancedbraces

type stack struct {
	items []rune
}

func (s *stack) push(n rune) {
	s.items = append(s.items, n)
}

func (s *stack) len() int {
	return len(s.items)
}

func (s *stack) pop() rune {
	if len(s.items) == 0 {
		return 0
	}

	r := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return r
}
