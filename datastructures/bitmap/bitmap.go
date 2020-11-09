package bitmap

// BitMap represent a bit-map datastructure
type BitMap struct {
	bits []int32
	size int
}

// New can construct a BitMap and return its ptr
func New(size int) *BitMap {
	return &BitMap{
		bits: make([]int32, size),
		size: size*32 - 1,
	}
}

// Add add num to BitMap
func (m *BitMap) Add(num int) {
	if num > m.size {
		return
	}
	i32Offset := num / 32
	bitOffset := num % 32
	bits := m.bits[i32Offset]
	bits = bits | (1 << bitOffset)
	m.bits[i32Offset] = bits
}

// Has returns whether a num is in BitMap
func (m *BitMap) Has(num int) bool {
	if num > m.size {
		return false
	}
	i32Offset := num / 32
	bitOffset := num % 32
	bit := m.bits[i32Offset] & (1 << bitOffset)
	return bit != 0
}
