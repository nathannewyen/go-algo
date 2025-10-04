package bitmanip

// BitMaskHelper9 provides precomputed bitmasks for common operations
type BitMaskHelper9 struct {
	lowNibbleMask  uint32
	highNibbleMask uint32
	byteMask       uint32
}

// NewBitMaskHelper9 creates helper with standard bitmasks
func NewBitMaskHelper9() *BitMaskHelper9 {
	return &BitMaskHelper9{
		lowNibbleMask:  0x0F,
		highNibbleMask: 0xF0,
		byteMask:       0xFF,
	}
}

// ExtractLowNibble9 returns the lower 4 bits of a byte
func (helper *BitMaskHelper9) ExtractLowNibble9(value uint32) uint32 {
	return value & helper.lowNibbleMask
}

// ExtractHighNibble9 returns the upper 4 bits of a byte
func (helper *BitMaskHelper9) ExtractHighNibble9(value uint32) uint32 {
	return (value & helper.highNibbleMask) >> 4
}
