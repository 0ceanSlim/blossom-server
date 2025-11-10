package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nbd-wtf/go-nostr"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run test-upload.go <private-key-hex> <file-hash-sha256>")
		os.Exit(1)
	}

	sk := os.Args[1]
	fileHash := os.Args[2]

	ev := &nostr.Event{
		CreatedAt: nostr.Now(),
		Kind:      24242,
		Tags: nostr.Tags{
			{"t", "upload"},
			{"x", fileHash},
			{"expiration", fmt.Sprintf("%d", time.Now().Add(time.Hour*24).Unix())},
		},
	}
	ev.Sign(sk)

	bytes, err := json.Marshal(ev)
	if err != nil {
		log.Fatal(err)
	}

	b64 := base64.StdEncoding.EncodeToString(bytes)

	fmt.Println(b64)
}
