package bloomfilter

import "testing"

func TestAdd(t *testing.T) {
	f := New(100)
	f.Add("foo")
	if !f.MayHas("foo") {
		t.Errorf("wanted %v but get nil", "foo")
	}
}
