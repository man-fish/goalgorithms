package compare

// Comparable represent elements that can be compared
type Comparable interface {
	// Equal return whether Comparable is equal to c
	Equal(c Comparable) bool
	// Equal return a number among 1, 0, -1,
	// 1  represent bigger than c
	// 0  represent same as c
	// -1 represent smaller then c
	CompareTo(c Comparable) int
}
