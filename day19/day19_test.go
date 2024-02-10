package day19

import "testing"

func TestTesseractClone(t *testing.T) {

	tt := &tesseract{
		x: ratingMinMax{1, 4000},
		m: ratingMinMax{1, 4000},
		a: ratingMinMax{1, 4000},
		s: ratingMinMax{1, 4000},
	}

	t2 := tt.clone()

	t2.x.min = 1010

	if tt.x.min != 1 && t2.x.min == 1010 {
		t.Errorf("want: 1, got: %d", tt.x.min)
	}
}
