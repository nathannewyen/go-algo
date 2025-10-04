package bitmanip

// BitMaskHelper12 provides precomputed bitmasks for common operations
type BitMaskHelper12 struct {
	lowNibbleMask  uint32
	highNibbleMask uint32
	byteMask       uint32
}

// NewBitMaskHelper12 creates helper with standard bitmasks
func NewBitMaskHelper12() *BitMaskHelper12 {
	return &BitMaskHelper12{
		lowNibbleMask:  0x0F,
		highNibbleMask: 0xF0,
		byteMask:       0xFF,
	}
}

// ExtractLowNibble12 returns the lower 4 bits of a byte
func (helper *BitMaskHelper12) ExtractLowNibble12(value uint32) uint32 {
	return value & helper.lowNibbleMask
}

// ExtractHighNibble12 returns the upper 4 bits of a byte
func (helper *BitMaskHelper12) ExtractHighNibble12(value uint32) uint32 {
	return (value & helper.highNibbleMask) >> 4
}
