package day22

import (
	"testing"
)

func TestCanFall(t *testing.T) {

	bricks := []*brick{
		{
			id: 0,
			p1: vector{0, 0, 1},
			p2: vector{0, 2, 1},
		},
		{
			id: 1,
			p2: vector{0, 2, 3},
			p1: vector{0, 0, 2},
		},
		{
			id: 2,
			p2: vector{2, 2, 3},
			p1: vector{2, 2, 2},
		},
		{
			id: 3,
			p2: vector{1, 1, 3},
			p1: vector{2, 2, 2},
		},
		{
			id: 4,
			p2: vector{0, 3, 2},
			p1: vector{2, 4, 2},
		},
		{
			id: 5,
			p2: vector{0, 2, 2},
			p1: vector{2, 4, 2},
		},
		{
			id: 6,
			p2: vector{0, 0, 10},
			p1: vector{10, 10, 11},
		},
	}
	b := bricks[1]
	want := false

	got := canFall(bricks, b)

	if got != want {
		t.Errorf("for: %s in: %v, want: %v, got: %v", b, bricks, want, got)
	}

	b = bricks[0]
	want = false

	got = canFall(bricks, b)

	if got != want {
		t.Errorf("for: %s in: %v, want: %v, got: %v", b, bricks, want, got)
	}

	b = bricks[2]
	want = true

	got = canFall(bricks, b)

	if got != want {
		t.Errorf("for: %s in: %v, want: %v, got: %v", b, bricks, want, got)
	}

	b = bricks[3]
	want = true

	got = canFall(bricks, b)

	if got != want {
		t.Errorf("for: %s in: %v, want: %v, got: %v", b, bricks, want, got)
	}

	b = bricks[4]
	want = true

	got = canFall(bricks, b)

	if got != want {
		t.Errorf("for: %s in: %v, want: %v, got: %v", b, bricks, want, got)
	}

	b = bricks[5]
	want = false

	got = canFall(bricks, b)

	if got != want {
		t.Errorf("for: %s in: %v, want: %v, got: %v", b, bricks, want, got)
	}

	b = bricks[6]
	want = true

	got = canFall(bricks, b)

	if got != want {
		t.Errorf("for: %s in: %v, want: %v, got: %v", b, bricks, want, got)
	}
}

func TestCanFall_AllPlacedAround(t *testing.T) {
	bricks := []*brick{
		&brick{
			id: 0,
			p1: vector{2, 2, 2},
			p2: vector{3, 3, 2},
		},
		//
		&brick{
			id: 1,
			p1: vector{0, 0, 1},
			p2: vector{1, 1, 1},
		},
		&brick{
			id: 2,
			p1: vector{2, 1, 1},
			p2: vector{5, 1, 1},
		},
		&brick{
			id: 3,
			p1: vector{4, 1, 1},
			p2: vector{5, 2, 1},
		},
		&brick{
			id: 4,
			p1: vector{4, 4, 1},
			p2: vector{4, 3, 1},
		},
		&brick{
			id: 5,
			p1: vector{6, 4, 1},
			p2: vector{0, 6, 1},
		},
		&brick{
			id: 6,
			p1: vector{0, 3, 1},
			p2: vector{1, 3, 1},
		},
		&brick{
			id: 7,
			p1: vector{0, 3, 1},
			p2: vector{0, 3, 1},
		},
		&brick{
			id: 8,
			p1: vector{1, 2, 1},
			p2: vector{1, 2, 1},
		},
	}
	b := bricks[0]

	want := true
	got := canFall(bricks, b)

	if got != want {
		t.Errorf("wanted: %v, got %v", want, got)
	}
}

func TestIsSupporting_AllPlacedAround(t *testing.T) {
	bricks := []*brick{
		&brick{
			id: 0,
			p1: vector{2, 2, 0},
			p2: vector{3, 3, 0},
		},
		//
		&brick{
			id: 1,
			p1: vector{0, 0, 1},
			p2: vector{1, 1, 1},
		},
		&brick{
			id: 2,
			p1: vector{2, 1, 1},
			p2: vector{5, 1, 1},
		},
		&brick{
			id: 3,
			p1: vector{4, 1, 1},
			p2: vector{5, 2, 1},
		},
		&brick{
			id: 4,
			p1: vector{4, 4, 1},
			p2: vector{4, 3, 1},
		},
		&brick{
			id: 5,
			p1: vector{6, 4, 1},
			p2: vector{0, 6, 1},
		},
		&brick{
			id: 6,
			p1: vector{0, 3, 1},
			p2: vector{1, 3, 1},
		},
		&brick{
			id: 7,
			p1: vector{0, 3, 1},
			p2: vector{0, 3, 1},
		},
		&brick{
			id: 8,
			p1: vector{1, 2, 1},
			p2: vector{1, 2, 1},
		},
	}
	b1 := bricks[0]

	for i := 1; i < len(bricks); i++ {
		b2 := bricks[i]

		want := false
		got := isSupporting(b1, b2)

		if got != want {
			t.Errorf("for b1: %v, b2: %v wanted: %v, got %v", want, got, b1, b2)
		}
	}
}

func TestIsSupporting_AllSupportedAround(t *testing.T) {
	bricks := []*brick{
		&brick{
			id: 0,
			p1: vector{2, 2, 0},
			p2: vector{3, 3, 0},
		},
		//
		&brick{
			id: 1,
			p1: vector{0, 0, 1},
			p2: vector{2, 2, 1},
		},
		&brick{
			id: 2,
			p1: vector{2, 1, 1},
			p2: vector{5, 2, 1},
		},
		&brick{
			id: 3,
			p1: vector{3, 1, 1},
			p2: vector{5, 2, 1},
		},
		&brick{
			id: 4,
			p1: vector{4, 4, 1},
			p2: vector{3, 3, 1},
		},
		&brick{
			id: 5,
			p1: vector{6, 3, 1},
			p2: vector{0, 6, 1},
		},
		&brick{
			id: 6,
			p1: vector{0, 3, 1},
			p2: vector{2, 3, 1},
		},
		&brick{
			id: 7,
			p1: vector{2, 3, 1},
			p2: vector{3, 3, 1},
		},
		&brick{
			id: 8,
			p1: vector{2, 2, 1},
			p2: vector{1, 2, 1},
		},
		&brick{
			id: 9,
			p1: vector{3, 2, 1},
			p2: vector{3, 2, 1},
		},
		&brick{
			id: 9,
			p1: vector{2, 3, 1},
			p2: vector{2, 3, 1},
		},
	}
	b1 := bricks[0]

	for i := 1; i < len(bricks); i++ {
		b2 := bricks[i]

		want := true
		got := isSupporting(b1, b2)

		if got != want {
			t.Errorf("for b1: %v, b2: %v wanted: %v, got %v", want, got, b1, b2)
		}
	}
}

func TestIsSupporting(t *testing.T) {
	b1 := &brick{
		id: 1,
		p1: vector{2, 2, 1},
		p2: vector{3, 3, 2},
	}

	b2 := &brick{
		id: 2,
		p1: vector{2, 2, 3},
		p2: vector{3, 3, 2},
	}

	want := true
	got := isSupporting(b1, b2)

	if got != want {
		t.Errorf("for b1: %v, b2: %v wanted: %v, got %v", want, got, b1, b2)
	}
}
