package util

import (
	"math/rand"
	"strings"
	"time"
)

var randomstr = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt generates a random integer  between min , max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//RandomString generates a random string og lenght n
func RandomString(n int) string {
	var s strings.Builder
	l := len(randomstr)
	for i := 0; i < n; i++ {
		index := rand.Intn(l)
		s.WriteByte(randomstr[index])
	}
	return s.String()
}

//RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

//RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(1000, 2000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
