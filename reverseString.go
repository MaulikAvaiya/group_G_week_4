package main

func RevString(s string) string {
	byte_s := []rune(s)
	for h, d := 0, len(byte_s)-1; h < d; h, d = h+1, d-1 {
		byte_s[h], byte_s[d] = byte_s[d], byte_s[h]
	}
	return string(byte_s)
}
