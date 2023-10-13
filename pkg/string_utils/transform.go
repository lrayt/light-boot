package string_utils

import "unsafe"

func Byte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Str2Byte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
