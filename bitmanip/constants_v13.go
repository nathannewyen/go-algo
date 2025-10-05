package bitmanip

// BitMaskHelper13 provides precomputed bitmasks for common operations
type BitMaskHelper13 struct {
	lowNibbleMask  uint32
	highNibbleMask uint32
	byteMask       uint32
}

// NewBitMaskHelper13 creates helper with standard bitmasks
func NewBitMaskHelper13() *BitMaskHelper13 {
	return &BitMaskHelper13{
		lowNibbleMask:  0x0F,
		highNibbleMask: 0xF0,
		byteMask:       0xFF,
	}
}

// ExtractLowNibble13 returns the lower 4 bits of a byte
func (helper *BitMaskHelper13) ExtractLowNibble13(value uint32) uint32 {
	return value & helper.lowNibbleMask
}

// ExtractHighNibble13 returns the upper 4 bits of a byte
func (helper *BitMaskHelper13) ExtractHighNibble13(value uint32) uint32 {
	return (value & helper.highNibbleMask) >> 4
}
