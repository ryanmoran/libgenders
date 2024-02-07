package libgenders

import (
	"maps"
	"strconv"
	"strings"
)

type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

func (p Parser) Parse(line string) ([]Node, error) {
	line, _, _ = strings.Cut(line, "#")
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return nil, nil
	}

	fields := strings.Fields(line)
	names := p.parseNames(fields[0])
	var attributes map[string]string
	if len(fields) > 1 {
		attributes = p.parseAttrs(fields[1])
	}

	var nodes []Node
	for _, name := range names {
		var attrs map[string]string
		if len(attributes) > 0 {
			attrs = make(map[string]string)
			maps.Copy(attrs, attributes)
		}

		nodes = append(nodes, NewNode(name, attrs))
	}

	return nodes, nil
}

func (p Parser) parseAttrs(field string) map[string]string {
	attributes := make(map[string]string)
	for _, attr := range strings.Split(field, ",") {
		key, value, _ := strings.Cut(attr, "=")
		attributes[key] = value
	}

	return attributes
}

func (p Parser) parseNames(field string) []string {
	var (
		name    string
		inRange bool
		fields  []string
	)

	for _, r := range field {
		if r == ',' && !inRange {
			fields = append(fields, name)
			name = ""
			inRange = false
			continue
		}

		if r == '[' {
			inRange = true
		}

		if r == ']' {
			inRange = false
		}

		name += string(r)
	}

	if len(name) != 0 {
		fields = append(fields, name)
	}

	var names []string
	for _, f := range fields {
		names = append(names, p.parseName(f)...)
	}

	return names
}

func (p Parser) parseName(field string) []string {
	parts := strings.FieldsFunc(field, func(c rune) bool { return c == '[' || c == ']' })
	if len(parts) < 2 {
		return parts
	}

	prefix := parts[0]
	rng := parts[1]

	var suffix string
	if len(parts) == 3 {
		suffix = parts[2]
	}

	var names []string
	for _, index := range p.parseRange(strings.Split(rng, ",")...) {
		names = append(names, prefix+index+suffix)
	}

	return names
}

func (p Parser) parseRange(ranges ...string) []string {
	var elems []string
	for _, rng := range ranges {
		start, end, _ := strings.Cut(rng, "-")

		first, err := strconv.Atoi(start)
		if err != nil {
			panic(err)
		}

		if len(end) == 0 {
			elems = append(elems, strconv.Itoa(first))
			continue
		}

		last, err := strconv.Atoi(end)
		if err != nil {
			panic(err)
		}

		for i := first; i <= last; i++ {
			elems = append(elems, strconv.Itoa(i))
		}
	}

	return elems
}
