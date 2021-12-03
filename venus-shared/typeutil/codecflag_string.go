// Code generated by "stringer -type=CodecFlag -trimprefix=Codec"; DO NOT EDIT.

package typeutil

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CodecBinary-1]
	_ = x[CodecText-2]
	_ = x[CodecJSON-4]
	_ = x[CodecCbor-8]
	_ = x[CodecMaxLimit-16]
}

const (
	_CodecFlag_name_0 = "BinaryText"
	_CodecFlag_name_1 = "JSON"
	_CodecFlag_name_2 = "Cbor"
	_CodecFlag_name_3 = "MaxLimit"
)

var (
	_CodecFlag_index_0 = [...]uint8{0, 6, 10}
)

func (i CodecFlag) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _CodecFlag_name_0[_CodecFlag_index_0[i]:_CodecFlag_index_0[i+1]]
	case i == 4:
		return _CodecFlag_name_1
	case i == 8:
		return _CodecFlag_name_2
	case i == 16:
		return _CodecFlag_name_3
	default:
		return "CodecFlag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
