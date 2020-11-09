package bitmap

import "testing"

func TestAdd(t *testing.T) {
	bm := New(100)
	bm.Add(10)
	if !bm.Has(10) {
		t.Errorf("wanted %v but get nil", 10)
	}
}
