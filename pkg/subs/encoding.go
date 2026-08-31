package subs

import "unicode/utf8"

func IsValidUTF8(b []byte) bool {
	return utf8.Valid(b)
}
