package queue

import "testing"

func TestPush(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	if q.Peek() != 3 {
		t.Errorf("push failed expeted:%v", 3)
		t.Errorf("			   getted:%v", q.Peek())
	}
}

func TestPoll(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	res := s.Poll()
	if res != 1 {
		t.Errorf("pop failed expeted:%v", 1)
		t.Errorf("			  getted:%v", res)
	}
	s.Poll()
	s.Poll()
	if s.Peek() != nil {
		t.Errorf("pop failed expeted:%v", nil)
		t.Errorf("			  getted:%v", s.Peek())
	}

}
