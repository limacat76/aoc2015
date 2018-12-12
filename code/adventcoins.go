package code

import (
	"crypto/md5"
	"fmt"
	"strings"
)

//IsItCoin sees if the string in input is good for a coin
func IsItCoin(input string) bool {
	return strings.HasPrefix(input, "00000")
}

//IsItCoin sees if the string in input is good for a coin with 6 leading zeroes
func IsItCoin6(input string) bool {
	return strings.HasPrefix(input, "000000")
}

//CalculateHash calculates the hash of key in current iteration
func CalculateHash(key string, iteration int) string {
	y := fmt.Sprintf("%s%d", key, iteration)
	x := fmt.Sprintf("%x", md5.Sum([]byte(y)))
	return x
}
