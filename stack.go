package libgenders

type Stack struct {
	tokens []Token
}

func (s *Stack) Push(token Token) {
	s.tokens = append(s.tokens, token)
}

func (s *Stack) Pop() Token {
	token := s.Top()
	s.tokens = s.tokens[:len(s.tokens)-1]
	return token
}

func (s *Stack) Top() Token {
	return s.tokens[len(s.tokens)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.tokens) == 0
}
