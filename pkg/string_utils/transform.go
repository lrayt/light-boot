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

func Str2ByteWithPrefix(prefix, s string) []byte {
	s = prefix + s
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
