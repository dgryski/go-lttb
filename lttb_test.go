package lttb

import (
	"reflect"
	"testing"
)

// TestPointSelectionWithZeroArea is a regression test for the
// correct selection of a "b" point in a bucket where all computed
// areas have a value of zero.
func TestPointSelectionWithZeroArea(t *testing.T) {
	data := []Point[float64]{
		{0, 0}, // sentinel value
		{1299456, 116.3707},
		{1300320, 116.3752}, // a
		{1301184, 116.3648}, // b --> Should be selected even when triangle area is zero.
		{1302048, 116.3544}, // c
		{1302912, 116.3328},
		{1306368, 116.3277},
		{1307232, 116.2676},
	}

	want := []Point[float64]{
		{0, 0},
		{1299456, 116.3707},
		{1300320, 116.3752},
		{1301184, 116.3648},
		{1302048, 116.3544},
		{1306368, 116.3277},
		{1307232, 116.2676},
	}

	if have := LTTB(data, 7); !reflect.DeepEqual(have, want) {
		t.Errorf("\nhave %v\nwant %v", have, want)
	}
}
