package pbr

// Bias is the minimum distance unit.
// Applying bias provides more robust processing of geometry.
const Bias = 1e-6

// Pixel elements are stored in specific offsets.
// These constants allow easy access, eg `someFloat64Array[i + Blue]`
const (
	Red      = 0
	Green    = 1
	Blue     = 2
	Count    = 3
	Noise    = 4
	Elements = 5
)
