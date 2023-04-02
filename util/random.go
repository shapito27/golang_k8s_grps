package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates random integer between min and max.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)

}

// RandomString generates random string of length n.
func RandomString(n int) string {
	var sb strings.Builder
	length := len(alphabet)
	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(length)])
	}
	return sb.String()
}

// RandomOwner generates string of length 6.
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates random amount of money.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates currency code.
func RandomCurrency() string {
	currencies := []string{"USD", "MYR", "AUD"}
	return currencies[rand.Intn(len(currencies))]
}
