/*
Package hashset implement a set datasruct:
 	In computer science, a set is an abstract data type that can store unique values,
	without any particular order. It is a computer implementation of the mathematical
	concept of a finite set. Unlike most other collection types, rather than retrieving
	a specific element from a set, one typically tests a value for membership in a set.
WikiPage:
	* https://en.wikipedia.org/wiki/Set_(abstract_data_type)
*/
package hashset

import (
	"errors"
	"reflect"
)

// HashSet is hashset implemented by HashMap
type HashSet struct {
	set  map[interface{}]interface{}
	size int
	t    string
}

// New is construct func of Hashset
func New(dataType string) *HashSet {
	return &HashSet{
		set:  make(map[interface{}]interface{}),
		t:    dataType,
		size: 0,
	}
}

// Size returns ele nums of hashset
func (s *HashSet) Size() int {
	return s.size
}

// T returns element type of hashset
func (s *HashSet) T() string {
	return s.t
}

// All returns all eles
func (s *HashSet) All() []interface{} {
	data := make([]interface{}, 0)
	for v := range s.set {
		data = append(data, v)
	}
	return data
}

// Add ele to set
func (s *HashSet) Add(data interface{}) error {
	err := s.checkT(data)
	if err != nil {
		return err
	}
	_, ok := s.set[data]
	if ok {
		return errors.New("ele exist")
	}
	s.set[data] = nil
	s.size++
	return nil
}

func (s *HashSet) checkT(data interface{}) error {
	if data == nil {
		return errors.New("data should not be nil")
	}

	t := reflect.TypeOf(data).String()
	if t != s.t {
		return errors.New("unsupported type")
	}
	return nil
}

// Remove ele from set
func (s *HashSet) Remove(data interface{}) error {
	err := s.checkT(data)
	if err != nil {
		return err
	}
	_, ok := s.set[data]
	if ok {
		delete(s.set, data)
		s.size--
		return nil
	}
	return errors.New("key not exist")
}

// Contain return whether the set contains centain value
func (s *HashSet) Contain(data interface{}) (bool, error) {
	err := s.checkT(data)
	if err != nil {
		return false, err
	}
	_, ok := s.set[data]
	if ok {
		return true, nil
	} else {
		return false, nil
	}
}

// Clear empty set
func (s *HashSet) Clear() {
	s.set = make(map[interface{}]interface{})
	s.size = 0
}
