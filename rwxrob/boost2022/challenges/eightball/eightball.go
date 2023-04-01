package eightball

import (
	"math/rand"
	"time"
)

var Responses = [...]string{
	"Yes",
	"No",
	"Maybe",
	"Why not?",
}

func Respond() string {
	rand.Seed(time.Now().UnixNano())
	return Responses[rand.Intn(len(Responses))]
}
