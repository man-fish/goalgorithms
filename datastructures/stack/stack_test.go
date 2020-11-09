package stack

import "testing"

func TestPush(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Peek() != 3 {
		t.Errorf("push failed expeted:%v", 3)
		t.Errorf("			   getted:%v", s.Peek())
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	res := s.Pop()
	if res != 3 {
		t.Errorf("pop failed expeted:%v", 3)
		t.Errorf("			  getted:%v", res)
	}
	s.Pop()
	s.Pop()
	if s.Peek() != nil {
		t.Errorf("pop failed expeted:%v", nil)
		t.Errorf("			  getted:%v", s.Peek())
	}
}
