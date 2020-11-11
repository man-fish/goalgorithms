package fenwicktree

type FenWickTree struct {
	size int
	tree []int
}

func New(s int) *FenWickTree {
	return &FenWickTree{
		size: s,
		tree: make([]int, s+1),
	}
}

// TODO
