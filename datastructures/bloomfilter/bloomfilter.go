/*
Package bloomfilter implements a bloom filter:
	A **bloom filter** is a space-efficient probabilistic
	data structure designed to test whether an element
	is present in a set. It is designed to be blazingly
	fast and use minimal memory at the cost of potential
	false positives. False positive matches are possible,
	but false negatives are not – in other words, a query
	returns either "possibly in set" or "definitely not in set".

	Bloom proposed the technique for applications where the
	amount of source data would require an impractically large
	amount of memory if "conventional" error-free hashing
	techniques were applied.

False Positives:
	The probability of false positives is determined by
	three factors: the size of the bloom filter, the
	number of hash functions we use, and the number
	of items that have been inserted into the filter.

	The formula to calculate probablity of a false positive is:
		* ( 1 - e <sup>-kn/m</sup> ) <sup>k</sup>
		* `k` = number of hash functions
		* `m` = filter size
		* `n` = number of items inserted
	These variables, `k`, `m`, and `n`, should be picked based
	on how acceptable false positives are. If the values
	are picked and the resulting probability is too high,
	the values should be tweaked and the probability

BitMap on Wiki:
	* https://en.wikipedia.org/wiki/Bloom_filter
*/
package bloomfilter

import (
	"math"

	"github.com/man-fish/goalgorithms/datastructures/bitmap"
)

// BloomFilter is a space-efficient probabilistic data structure designed to test whether an element is present in a set.
type BloomFilter struct {
	m    *bitmap.BitMap
	size int
}

// New is a constructor
func New(n int) *BloomFilter {
	return &BloomFilter{
		m:    bitmap.New(n),
		size: n,
	}
}

// Add adds a element to bloomfilter
func (b *BloomFilter) Add(k string) {
	hslice := b.hashes(k)
	for _, hash := range hslice {
		b.m.Add(hash)
	}
}

// MayHas return whether the key may in the filter
func (b *BloomFilter) MayHas(k string) bool {
	hslice := b.hashes(k)
	for _, hash := range hslice {
		if !b.m.Has(hash) {
			return false
		}
	}
	return true
}

func (b *BloomFilter) hashes(k string) []int {
	return []int{b.hash1(k), b.hash2(k), b.hash3(k)}
}

func (b *BloomFilter) hash1(k string) int {
	hash := 0
	for _, v := range k {
		hash := (hash << 5) + hash + int(v)
		// Convert to 32bit integer
		hash &= hash
		hash = int(math.Abs(float64(hash)))
	}
	return hash % b.size
}

func (b *BloomFilter) hash2(k string) int {
	hash := 5381
	for _, v := range k {
		// hash * 33 + c
		hash = (hash << 5) + hash + int(v)
	}
	return int(math.Abs(float64(hash % b.size)))
}

func (b *BloomFilter) hash3(k string) int {
	hash := 0
	for _, v := range k {
		hash = (hash << 5) - hash + int(v)
		// Convert to 32bit integer
		hash &= hash
	}
	return int(math.Abs(float64(hash % b.size)))
}

// ...
// func (b *BloomFilter) hashk(k string) int
