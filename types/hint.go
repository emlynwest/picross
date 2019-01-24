package types

type Hint []int

func (r *Hint) Equal(c *Hint) bool {

	if r == nil && c == nil {
		return true
	}

	if r == nil || c == nil {
		return false
	}

	a := *r
	b := *c

	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}
