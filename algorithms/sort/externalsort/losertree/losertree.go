package losertree

import (
	"math"
)

const (
	// EOF marks the end of file
	EOF = math.MaxInt32
	// EOS marks the end of a segement in output file
	EOS = 9999
)

// Sourcer represent the leaf of the loser tree and where the data come from
type Sourcer interface {
	// Next returns next num from source with index i (segement)
	// if no nums remain return EOF, All
	// EOF means merge over
	Next(i int) int
}

// Enrty is source from file sys
type Enrty struct {
	// s for segement num with default 1
	S int
	// k for key
	K int
}

// LoserTree implements a min heap datastructure
// it is divided to branch and leaf two parts
// len(branch) = len(leaf) - 1
type LoserTree struct {
	// branch[0] store the winner, branch[1:k] store the losers
	branch []int
	// leaf represent workarea
	Leaf []Enrty
	// size represent workarea count
	size  int
	Input Sourcer
}

// New is the constructor of LoserTree with n data source it has
// n branch, the first leaf store a final loser in compartions
// n+1 leaf, the last leaf store the min value for compare.
func New(n int, s Sourcer) *LoserTree {
	loser := &LoserTree{
		branch: make([]int, n),
		Leaf:   make([]Enrty, n),
		size:   n,
		Input:  s,
	}
	loser.build()
	return loser
}

// build fill workspace
func (t *LoserTree) build() {
	for i := 0; i < t.size; i++ {
		t.branch[i], t.Leaf[i] = 0, Enrty{S: 0, K: 0}
	}

	for i := t.size - 1; i >= 0; i-- {
		t.Leaf[i].K = t.Input.Next(i)
		t.Leaf[i].S = 1
		t.Contest(i)
	}
}

// Winner returns cur competition winner
func (t *LoserTree) Winner() int {
	return t.branch[0]
}

// Contest get the minVal from k sources and fill the tree with losers (larger num)
func (t *LoserTree) Contest(i int) {
	p := (i + t.size) / 2

	for p > 0 {
		pa := t.branch[p]
		if t.Leaf[i].S > t.Leaf[pa].S || (t.Leaf[i].S == t.Leaf[pa].S && t.Leaf[i].K > t.Leaf[pa].K) {
			// i always stores the smallest num
			// t.branch always store the larger num(loser)
			t.branch[p], i = i, t.branch[p]
		}
		p = p / 2
	}
	// store the smallest num
	t.branch[0] = i
}

// // KMerge sort the k sources and write to output
// func (t *LoserTree) KMerge() {
// 	winner := t.branch[0]
// 	for min := t.Leaf[winner].k; min != EOF; {
// 		fmt.Fprintf(t.output, "%v", min)
// 		t.Leaf[winner].k = t.input.Next(winner)
// 		t.Contest(winner)
// 	}
// }

// // ReplaceSelection make k segments to output
// func (t *LoserTree) ReplaceSelection() {
// 	scur, smax := 1, 1
// 	// cur segement id and segemnt max
// 	for scur <= smax {
// 		// make segment to output
// 		t.buildSegement(scur, &smax)
// 		fmt.Fprintf(t.output, "%v", EOS)
// 		// add end of segment
// 		scur = t.Leaf[t.branch[0]].s
// 	}
// }

// func (t *LoserTree) buildSegement(s int, smax *int) {
// 	for t.Leaf[t.branch[0]].s == s {
// 		// get cur minmax id
// 		i := t.branch[0]
// 		// get cur min value of minmax
// 		minmax := t.Leaf[i].k
// 		// write to output
// 		fmt.Fprintf(t.output, "%v", minmax)
// 		if res := t.input.Next(i); res == EOF {
// 			t.Leaf[i].k = EOF
// 			t.Leaf[i].s = *smax + 1
// 		} else {
// 			t.Leaf[i].k = res
// 			if res < minmax {
// 				t.Leaf[i].s = s + 1
// 				*smax = s + 1
// 			} else {
// 				t.Leaf[i].s = s
// 			}
// 		}
// 		t.Contest(i)
// 	}
// }
