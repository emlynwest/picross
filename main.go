package main

import (
	"picross/engine"
	"picross/types"
)

func main() {

	a := types.Grid{
		{1, 1, 1, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1, 1, 0, 0, 0},
		{1, 0, 0, 1, 1, 0, 0, 1, 1},
	}

	for _, r := range a {
		engine.TallyRow(r)
	}
}
