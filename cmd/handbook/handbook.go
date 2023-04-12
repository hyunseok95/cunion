package main

import (
	"math/rand"
	"time"

	app "github.com/hyunseok95/cunion/pkg/handbook"
	util "github.com/hyunseok95/cunion/pkg/util"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func main() {
	util.Trap(app.Run())
}
