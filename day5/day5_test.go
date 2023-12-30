package day5_test

import (
	"testing"

	"github.com/unnamedxaer/aoc2023/day5"
)

func TestSortSeeds(t *testing.T) {

	seeds := [][]int{
		{
			5, -1,
			3, -1,
			9, -1,
			1, -1,
			2, -1,
			8, -1,
			4, -1,
			7, -1,
			6, -1,
		},
		{
			202517468, 131640971, 1553776977, 241828580, 1435322022, 100369067, 2019100043, 153706556, 460203450, 84630899, 3766866638, 114261107, 1809826083, 153144153, 2797169753, 177517156, 2494032210, 235157184, 856311572, 542740109,
		},
	}

	for i := 0; i < len(seeds); i++ {
		out := day5.SortSeeds(seeds[0])

		if len(seeds[i])/2 != len(out) {
			t.Fatalf("expected get %d elements out, got %d", len(seeds[i])/2, len(out))
		}

		prev := out[0]
		for i, v := range out {
			if prev[0] > v[0] {
				t.Fatalf("incorrect order: %v, at pos: %d", out, i)
			}
			if v[0] == 0 {
				t.Fatalf("there should be no zeros, got: %v", out)
			}
		}
	}
}
