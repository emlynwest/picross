package types

import "testing"

func TestHint_ArrEqual(t *testing.T) {
	sets := []struct {
		a *Hint
		b *Hint
		r bool
	}{
		{nil, nil, true},
		{&Hint{}, &Hint{}, true},
		{&Hint{1}, &Hint{1}, true},
		{&Hint{1, 2, 3}, &Hint{1, 2, 3}, true},
		{nil, &Hint{1, 2, 3}, false},
		{&Hint{1, 2, 3}, nil, false},
		{&Hint{3, 2, 1}, &Hint{1, 2, 3}, false},
		{&Hint{}, &Hint{1, 2, 3}, false},
		{&Hint{1, 2, 3}, &Hint{}, false},
		{&Hint{1, 2, 3}, &Hint{4, 5, 6}, false},
		{&Hint{1, 2, 3}, &Hint{1}, false},
		{&Hint{1}, &Hint{1, 2, 3}, false},
	}

	for i, s := range sets {
		if res := s.a.Equal(s.b); res != s.r {
			t.Errorf("Case %d failed. Expected %t Got %t", i, s.r, res)
		}
	}
}
