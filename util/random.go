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

// RandomInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {

	return min + rand.Int63n(max-min+1) //max-min
}

// RandomString generates a random string of length n
func RandomString(n int) string {

	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] //generate a number between 0-(k-1_
		sb.WriteByte(c)             //wireByte is a function to add character to string builder
	}

	return sb.String()

}

// RandomOwner generate a random owner name
func RandomOwner() string {

	return RandomString(6)
}

// RandomMoney generate a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generate a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "CAD", "USD", "INR"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
