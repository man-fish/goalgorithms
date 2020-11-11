/*
Package bitmap implements a bit map:
	In computing, a bitmap is a mapping from some domain (for
	example,  a range of integers) to bits. It is also called
	a bit array or bitmap index.
	As a noun, the term "bitmap" is very often used to refer
	to a particular bitmapping application: the pix-map, which
	refers to a map of pixels, where each one may store more
	than two colors, thus using more than one bit per pixel.
	In such a case, the domain in question is the array of pixels
	which constitute a digital graphic output device (a screen or
	monitor). In some contexts, the term bitmap implies one bit
	per pixel, while pixmap is used for images with multiple bits
	per pixel.
	A bitmap is a type of memory organization or image file format
	used to store digital images. The term bitmap comes from the
	computer programming terminology, meaning just a map of bits,
	a spatially mapped array of bits. Now, along with pixmap, it
	commonly refers to the similar concept of a spatially mapped
	array of pixels. Raster images in general may be referred to
	as bitmaps or pixmaps, whether synthetic or photographic, in
	files or memory.
BitMap on Wiki:
	* https://en.wikipedia.org/wiki/Bitmap
*/
package bitmap

// BitMap represent a bit-map datastructure
type BitMap struct {
	// int32 is also ok
	bits []uint32
	size int
}

// New can construct a BitMap and return its ptr
func New(size int) *BitMap {
	return &BitMap{
		bits: make([]uint32, size),
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
