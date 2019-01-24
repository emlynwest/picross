package engine

import (
	"picross/types"
)

func TallyRow(r *types.Row) *types.Hint {
	c := types.Hint{}

	count := 0

	for _, v := range *r {
		if v > 0 {
			count++
		}

		if v == 0 && count > 0 {
			c = append(c, count)
			count = 0
		}
	}

	if count > 0 {
		c = append(c, count)
	}

	return &c
}
