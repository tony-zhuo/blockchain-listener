package libs

import (
	"math/rand"
	"time"
)

func RandomInt(n int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	return random.Intn(n)
}
