package rt

import "emacs/lisp"

// ToBool coerces Emacs Lisp value to canonical boolean
// representation ("nil" or "t" symbol).
func ToBool(o lisp.Object) bool {
	return lisp.Not(lisp.Not(o))
}

// BytesToStr converts slice of bytes to string.
func BytesToStr(slice *Slice) string {
	if slice.offset == 0 {
		return arrayToStr(slice.data)
	}
	return arrayToStr(
		substring(slice.data, slice.offset, slice.offset+slice.len),
	)
}

// StrToBytes converts string to slice of bytes.
func StrToBytes(s string) *Slice {
	return ArrayToSlice(strToArray(s))
}
