package main

import (
	"crypto"
	"math/rand"
	"os"
	"time"

	"fmt"

	"github.com/containerd/containerd/pkg/hasher"
	"github.com/hyunseok95/cunion/cmd"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
	crypto.RegisterHash(crypto.SHA256, hasher.NewSHA256)
}

func main() {
    args := os.Args
    if len(args) < 2 {
        fmt.Println("Please provide a command")
        os.Exit(1)
    }
    command := args[1]
    fmt.Println("Executing command:", command)

    switch command {
    case "containerd":
        cmd.Containerd()
    case "test":
        cmd.Test()
    default:
        fmt.Println("Unknown command:", command)
        os.Exit(1)
    }
}
