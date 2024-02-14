package libgenders

import "sort"

type Set []int

func (s Set) Union(o Set) Set {
	sort.Ints(s)
	sort.Ints(o)

	var union Set
	i := 0
	j := 0
	for {
		if i == len(s) {
			union = append(union, o[j:]...)
			break
		}

		if j == len(o) {
			union = append(union, s[i:]...)
			break
		}

		if s[i] == o[j] {
			i++
			continue
		}

		if s[i] < o[j] {
			union = append(union, s[i])
			i++
		} else {
			union = append(union, o[j])
			j++
		}
	}

	return union
}

func (s Set) Intersection(o Set) Set {
	sort.Ints(s)
	sort.Ints(o)

	var intersection Set
	i := 0
	j := 0
	for {
		if i == len(s) {
			break
		}

		if j == len(o) {
			break
		}

		if s[i] == o[j] {
			intersection = append(intersection, s[i])
			i++
			j++
			continue
		}

		if s[i] < o[j] {
			i++
		} else {
			j++
		}
	}

	return intersection
}

func (s Set) Difference(o Set) Set {
	sort.Ints(s)
	sort.Ints(o)

	var diff []int
	i := 0
	j := 0
	for i < len(s) {
		if j < len(o) && s[i] == o[j] {
			i++
			j++
			continue
		}

		diff = append(diff, s[i])
		i++
	}

	return diff
}
