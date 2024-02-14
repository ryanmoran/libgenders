package internal

import "strings"

type Scanner struct {
	reader *strings.Reader
}

func NewScanner(s string) Scanner {
	return Scanner{reader: strings.NewReader(s)}
}

func (s Scanner) Len() int {
	return s.reader.Len()
}

func (s Scanner) Next() (string, error) {
	buffer := make([]byte, 2)
	n, err := s.reader.Read(buffer)
	if err != nil {
		return "", err
	}

	switch string(buffer) {
	case "&&", "||", "--":
		return string(buffer), nil
	}

	if n == 2 {
		err = s.reader.UnreadByte()
		if err != nil {
			return "", err
		}
	}

	return string(buffer[:1]), nil
}
