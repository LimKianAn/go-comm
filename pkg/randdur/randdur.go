package randdur // random duration

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Sec(max int) time.Duration { // [1:max] seconds
	randN := rand.Intn(max) + 1 // Intn returns [0:max)
	return time.Duration(randN) * time.Second
}
