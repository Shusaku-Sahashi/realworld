// Code generated by "stringer -type=envType"; DO NOT EDIT.

package config

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[dev-0]
	_ = x[prod-1]
}

const _envType_name = "devprod"

var _envType_index = [...]uint8{0, 3, 7}

func (i envType) String() string {
	if i < 0 || i >= envType(len(_envType_index)-1) {
		return "envType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _envType_name[_envType_index[i]:_envType_index[i+1]]
}