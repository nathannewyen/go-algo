package bitmanip

// BitMaskHelper7 provides precomputed bitmasks for common operations
type BitMaskHelper7 struct {
	lowNibbleMask  uint32
	highNibbleMask uint32
	byteMask       uint32
}

// NewBitMaskHelper7 creates helper with standard bitmasks
func NewBitMaskHelper7() *BitMaskHelper7 {
	return &BitMaskHelper7{
		lowNibbleMask:  0x0F,
		highNibbleMask: 0xF0,
		byteMask:       0xFF,
	}
}

// ExtractLowNibble7 returns the lower 4 bits of a byte
func (helper *BitMaskHelper7) ExtractLowNibble7(value uint32) uint32 {
	return value & helper.lowNibbleMask
}

// ExtractHighNibble7 returns the upper 4 bits of a byte
func (helper *BitMaskHelper7) ExtractHighNibble7(value uint32) uint32 {
	return (value & helper.highNibbleMask) >> 4
}
