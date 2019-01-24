package engine

import (
	"github.com/emlynwest/picross/types"
	"testing"
)

func TestTallyRow(t *testing.T) {
	sets := []struct {
		row    *types.Row
		result *types.Hint
	}{
		{&types.Row{1, 1, 1, 0, 0, 0, 1, 1, 1}, &types.Hint{3, 3}},
		{&types.Row{0, 0, 0, 1, 1, 1, 0, 0, 0}, &types.Hint{3}},
		{&types.Row{1, 0, 0, 1, 1, 0, 0, 1, 1}, &types.Hint{1, 2, 2}},
		{&types.Row{0, 0, 0, 0, 0, 0, 0, 0, 0}, &types.Hint{}},
		{&types.Row{1, 1, 1, 1, 1, 1, 1, 1, 1}, &types.Hint{9}},
	}

	for i, s := range sets {
		if r := TallyRow(s.row); !r.Equal(s.result) {
			t.Errorf("Case %d failed. Want %d Got %d", i, s.result, r)
		}
	}
}
