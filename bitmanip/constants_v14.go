package bitmanip

// BitMaskHelper14 provides precomputed bitmasks for common operations
type BitMaskHelper14 struct {
	lowNibbleMask  uint32
	highNibbleMask uint32
	byteMask       uint32
}

// NewBitMaskHelper14 creates helper with standard bitmasks
func NewBitMaskHelper14() *BitMaskHelper14 {
	return &BitMaskHelper14{
		lowNibbleMask:  0x0F,
		highNibbleMask: 0xF0,
		byteMask:       0xFF,
	}
}

// ExtractLowNibble14 returns the lower 4 bits of a byte
func (helper *BitMaskHelper14) ExtractLowNibble14(value uint32) uint32 {
	return value & helper.lowNibbleMask
}

// ExtractHighNibble14 returns the upper 4 bits of a byte
func (helper *BitMaskHelper14) ExtractHighNibble14(value uint32) uint32 {
	return (value & helper.highNibbleMask) >> 4
}
