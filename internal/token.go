package internal

import (
	"fmt"
	"slices"
	"strings"
)

type TokenKind uint8

const (
	SpaceTokenKind TokenKind = iota + 1
	LeftParenTokenKind
	RightParenTokenKind
	ComplementTokenKind
	IntersectionTokenKind
	UnionTokenKind
	DifferenceTokenKind
	ValueTokenKind
)

type Token struct {
	Kind TokenKind
	Text string
}

func NewToken(kind TokenKind, text string) Token {
	return Token{
		Kind: kind,
		Text: text,
	}
}

func Tokenize(query string) ([]Token, error) {
	var (
		buffer strings.Builder
		tokens []Token
	)

	scanner := NewScanner(query)
	for scanner.Len() > 0 {
		s, err := scanner.Next()
		if err != nil {
			return nil, err
		}

		var token Token
		switch s {
		case " ":
			token = NewToken(SpaceTokenKind, s)

		case "||":
			token = NewToken(UnionTokenKind, s)

		case "&&":
			token = NewToken(IntersectionTokenKind, s)

		case "--":
			token = NewToken(DifferenceTokenKind, s)

		case "~":
			token = NewToken(ComplementTokenKind, s)

		case "(":
			token = NewToken(LeftParenTokenKind, s)

		case ")":
			token = NewToken(RightParenTokenKind, s)

		default:
			buffer.WriteString(s)
			continue
		}

		if buffer.Len() > 0 {
			tokens = append(tokens, NewToken(ValueTokenKind, buffer.String()))
			buffer.Reset()
		}

		if token.Kind != SpaceTokenKind {
			tokens = append(tokens, token)
		}
	}

	if buffer.Len() > 0 {
		tokens = append(tokens, NewToken(ValueTokenKind, buffer.String()))
	}

	// NOTE: the following is an implementation of https://en.m.wikipedia.org/wiki/Shunting_yard_algorithm
	// It converts the query in in-fix notation to tokenized Polish notation for the parser to build an AST
	var (
		output    []Token
		operators = Stack{}
	)
	for _, token := range tokens {
		switch token.Kind {
		case ValueTokenKind:
			output = append(output, token)

		case ComplementTokenKind:
			operators.Push(token)

		case UnionTokenKind, IntersectionTokenKind, DifferenceTokenKind:
			for !operators.IsEmpty() && operators.Top().Kind != LeftParenTokenKind {
				output = append(output, operators.Pop())
			}

			operators.Push(token)

		case LeftParenTokenKind:
			operators.Push(token)

		case RightParenTokenKind:
			for !operators.IsEmpty() && operators.Top().Kind != LeftParenTokenKind {
				output = append(output, operators.Pop())
			}

			if operators.IsEmpty() {
				return nil, fmt.Errorf("failed to tokenize query %q: mismatched parentheses", query)
			}

			operators.Pop()

			if !operators.IsEmpty() && operators.Top().Kind == ComplementTokenKind {
				output = append(output, operators.Pop())
			}
		}
	}

	for !operators.IsEmpty() {
		if operators.Top().Kind == LeftParenTokenKind {
			return nil, fmt.Errorf("failed to tokenize query %q: mismatched parentheses", query)
		}

		output = append(output, operators.Pop())
	}

	slices.Reverse(output) // convert from Reverse Polish Notation to Polish Notation

	return output, nil
}
