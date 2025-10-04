package bitmanip

// BitMaskHelper8 provides precomputed bitmasks for common operations
type BitMaskHelper8 struct {
	lowNibbleMask  uint32
	highNibbleMask uint32
	byteMask       uint32
}

// NewBitMaskHelper8 creates helper with standard bitmasks
func NewBitMaskHelper8() *BitMaskHelper8 {
	return &BitMaskHelper8{
		lowNibbleMask:  0x0F,
		highNibbleMask: 0xF0,
		byteMask:       0xFF,
	}
}

// ExtractLowNibble8 returns the lower 4 bits of a byte
func (helper *BitMaskHelper8) ExtractLowNibble8(value uint32) uint32 {
	return value & helper.lowNibbleMask
}

// ExtractHighNibble8 returns the upper 4 bits of a byte
func (helper *BitMaskHelper8) ExtractHighNibble8(value uint32) uint32 {
	return (value & helper.highNibbleMask) >> 4
}
