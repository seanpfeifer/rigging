package num

// Float is a constraint for floating-point types.
type Float interface {
	~float32 | ~float64
}

// Signed is a constraint for signed integer types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint for unsigned integer types.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint for all integer types.
type Integer interface {
	Signed | Unsigned
}
