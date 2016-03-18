package nicename

import (
	"math/rand"
	"time"
)

func GeneratePair() string {
	rand.Seed(time.Now().UTC().UnixNano())

	return First[rand.Intn(len(First))] + " " + Second[rand.Intn(len(Second))]

}
