package util

import (
	"math/rand"
	"time"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())

}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)

	}
	return sb.String()
}

func RandomOwner() string {
	return randomString(6)

}

func RandomBalance() int64 {
	return RandomInt(0, 100)
}

func RandomCurrency() string{
	currencies := []string{"EURO", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]

}