package bitmanip

// BitMaskHelper10 provides precomputed bitmasks for common operations
type BitMaskHelper10 struct {
	lowNibbleMask  uint32
	highNibbleMask uint32
	byteMask       uint32
}

// NewBitMaskHelper10 creates helper with standard bitmasks
func NewBitMaskHelper10() *BitMaskHelper10 {
	return &BitMaskHelper10{
		lowNibbleMask:  0x0F,
		highNibbleMask: 0xF0,
		byteMask:       0xFF,
	}
}

// ExtractLowNibble10 returns the lower 4 bits of a byte
func (helper *BitMaskHelper10) ExtractLowNibble10(value uint32) uint32 {
	return value & helper.lowNibbleMask
}

// ExtractHighNibble10 returns the upper 4 bits of a byte
func (helper *BitMaskHelper10) ExtractHighNibble10(value uint32) uint32 {
	return (value & helper.highNibbleMask) >> 4
}
