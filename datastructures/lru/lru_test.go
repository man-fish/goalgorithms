package lru

import (
	"reflect"
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("foo", String("bar"))
	if v, ok := lru.Get("foo"); !ok || string(v.(String)) != "bar" {
		t.Fatalf("cache hit %v failed", "foo")
	}
	if _, ok := lru.Get("bar"); ok {
		t.Fatalf("cache miss %v failed", "bar")
	}
}

func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := "v1", "v2", "v3"
	mem := len(v1 + v2 + k1 + k2)
	lru := New(int64(mem), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("v1"); ok {
		t.Fatalf("cache miss %v failed", "v1")
	}
}

func TestOnEvicted(t *testing.T) {
	evictedKeys := make([]string, 0)
	evictedHandler := func(key string, value Value) {
		evictedKeys = append(evictedKeys, key)
	}
	lru := New(int64(4), evictedHandler)
	lru.Add("k1", String("v1"))
	lru.Add("k2", String("v2"))
	lru.Add("k3", String("v3"))

	expected := []string{"k1", "k2"}
	if !reflect.DeepEqual(expected, evictedKeys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expected)
	}
}
