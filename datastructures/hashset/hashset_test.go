package hashset

import "testing"

func TestAdd(t *testing.T) {
	s := New("string")
	s.Add("foo")
	if ok, _ := s.Contain("foo"); !ok {
		t.Errorf("wanted has: %v, but nil", "foo")
	}
}

func TestClear(t *testing.T) {
	s := New("string")
	s.Add("foo")
	s.Clear()
	if ok, _ := s.Contain("foo"); ok {
		t.Errorf("want nil, get: %v", "foo")
	}
}
