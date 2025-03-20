package utils

import (
	"math/rand"
	"strings"
)


func GetRandomStr(e int) string {
	const charset = "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678"
	var result strings.Builder
	for i := 0; i < e; i++ {
		index := rand.Intn(len(charset))
		result.WriteByte(charset[index])
	}
	return result.String()
}
