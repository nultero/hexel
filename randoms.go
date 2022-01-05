package hexel

import (
	"math/rand"
	"time"
)

// Populates seed for calling function, won't have to have this everywhere.
func TimeRandSeed() {
	rand.Seed(time.Now().UnixNano())
}
