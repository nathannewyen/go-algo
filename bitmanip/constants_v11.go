package bitmanip

// BitMaskHelper11 provides precomputed bitmasks for common operations
type BitMaskHelper11 struct {
	lowNibbleMask  uint32
	highNibbleMask uint32
	byteMask       uint32
}

// NewBitMaskHelper11 creates helper with standard bitmasks
func NewBitMaskHelper11() *BitMaskHelper11 {
	return &BitMaskHelper11{
		lowNibbleMask:  0x0F,
		highNibbleMask: 0xF0,
		byteMask:       0xFF,
	}
}

// ExtractLowNibble11 returns the lower 4 bits of a byte
func (helper *BitMaskHelper11) ExtractLowNibble11(value uint32) uint32 {
	return value & helper.lowNibbleMask
}

// ExtractHighNibble11 returns the upper 4 bits of a byte
func (helper *BitMaskHelper11) ExtractHighNibble11(value uint32) uint32 {
	return (value & helper.highNibbleMask) >> 4
}
