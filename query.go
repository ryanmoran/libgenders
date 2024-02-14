package libgenders

type Query interface {
	Evaluate(attrs, attrvals map[string]Set, indices Set) Set
}

func ParseQuery(tokens []Token) Query {
	query, _ := parseQuery(tokens)
	return query
}

func parseQuery(tokens []Token) (Query, []Token) {
	if len(tokens) > 0 {
		token, tokens := tokens[0], tokens[1:]

		switch token.Kind {
		case ValueTokenKind:
			return ValueQuery{Expression: token.Text}, tokens

		case ComplementTokenKind:
			query, tokens := parseQuery(tokens)
			return ComplementQuery{Query: query}, tokens

		case UnionTokenKind, IntersectionTokenKind, DifferenceTokenKind:
			right, tokens := parseQuery(tokens)
			left, tokens := parseQuery(tokens)

			switch token.Kind {
			case UnionTokenKind:
				return UnionQuery{Left: left, Right: right}, tokens
			case IntersectionTokenKind:
				return IntersectionQuery{Left: left, Right: right}, tokens
			case DifferenceTokenKind:
				return DifferenceQuery{Left: left, Right: right}, tokens
			}
		}
	}

	return nil, nil
}

type ValueQuery struct {
	Expression string
}

func (vq ValueQuery) Evaluate(attrs, attrvals map[string]Set, _ Set) Set {
	if result, ok := attrvals[vq.Expression]; ok {
		return result
	}
	return attrs[vq.Expression]
}

type UnionQuery struct {
	Left, Right Query
}

func (uq UnionQuery) Evaluate(attrs, attrvals map[string]Set, indices Set) Set {
	left := uq.Left.Evaluate(attrs, attrvals, indices)
	right := uq.Right.Evaluate(attrs, attrvals, indices)
	return left.Union(right)
}

type IntersectionQuery struct {
	Left, Right Query
}

func (iq IntersectionQuery) Evaluate(attrs, attrvals map[string]Set, indices Set) Set {
	left := iq.Left.Evaluate(attrs, attrvals, indices)
	right := iq.Right.Evaluate(attrs, attrvals, indices)
	return left.Intersection(right)
}

type DifferenceQuery struct {
	Left, Right Query
}

func (dq DifferenceQuery) Evaluate(attrs, attrvals map[string]Set, indices Set) Set {
	left := dq.Left.Evaluate(attrs, attrvals, indices)
	right := dq.Right.Evaluate(attrs, attrvals, indices)
	return left.Difference(right)
}

type ComplementQuery struct {
	Query Query
}

func (cq ComplementQuery) Evaluate(attrs, attrvals map[string]Set, indices Set) Set {
	query := cq.Query.Evaluate(attrs, attrvals, indices)
	return indices.Difference(query)
}
