package external

import (
	"fmt"
	"io"

	"github.com/man-fish/goalgorithms/algorithms/sort/externalsort/losertree"
)

// Spliter is a big file Spliter
type Spliter struct {
	t      *losertree.LoserTree
	output io.Writer
}

// NewSpliter is the constructor of spliter
func NewSpliter(n int, s losertree.Sourcer) *Spliter {
	spt := new(Spliter)
	spt.t = losertree.New(n, s)
	return spt
}

// ReplaceSelection make k segments to output
func (s *Spliter) ReplaceSelection() {
	scur, smax := 1, 1
	// cur segement id and segemnt max
	for scur <= smax {
		// make segment to output
		s.buildSegement(scur, &smax)
		fmt.Fprintf(s.output, "%v", losertree.EOS)
		// add end of segment
		scur = s.t.Leaf[s.t.Winner()].S
	}
}

func (s *Spliter) buildSegement(scur int, smax *int) {
	for w := s.t.Winner(); s.t.Leaf[w].S == scur; {
		// get cur minmax id
		i := w
		// get cur min value of minmax
		minmax := s.t.Leaf[i].K
		// write to output
		fmt.Fprintf(s.output, "%v", minmax)
		if res := s.t.Input.Next(i); res == losertree.EOF {
			s.t.Leaf[i].K = losertree.EOF
			s.t.Leaf[i].S = *smax + 1
		} else {
			s.t.Leaf[i].K = res
			if res < minmax {
				s.t.Leaf[i].S = scur + 1
				*smax = scur + 1
			} else {
				s.t.Leaf[i].S = scur
			}
		}
		s.t.Contest(i)
	}
}

// Combiner implement to merge k segments
type Combiner struct {
	t      *losertree.LoserTree
	output io.Writer
}

// NewCombiner is the constructor of Combiner
func NewCombiner(n int, s losertree.Sourcer) *Combiner {
	c := new(Combiner)
	c.t = losertree.New(n, s)
	return c
}

// KMerge sort the k sources and write to output
func (c *Combiner) KMerge() {
	winner := c.t.Winner()
	for min := c.t.Leaf[winner].K; min != losertree.EOF; {
		fmt.Fprintf(c.output, "%v", min)
		c.t.Leaf[winner].K = c.t.Input.Next(winner)
		c.t.Contest(winner)
	}
}
