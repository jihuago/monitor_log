package tools

import "math/rand"

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomString(len int) string {
	bytes := make([]byte, len)
	for i :=0; i < len; i++ {
		bytes[i] = byte(RandomInt(65, 89))
	}

	return string(bytes)
}
