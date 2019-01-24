package types

// Hint defines a row or column of hint numbers
type Hint []int

// Equal returns true if two Hints are the same
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
