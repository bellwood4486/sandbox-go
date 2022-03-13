package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

func printHeader(r *http.Request) {
	h := r.Header
	ks := make([]string, 0, len(h))
	for k := range h {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	for _, k := range ks {
		fmt.Printf("> %s: %s\n", k, r.Header.Get(k))
	}
}

func validMACSHA256(message, messageMAC []byte) bool {
	// FIXME: Never hardcode the token into your app!
	mac := hmac.New(sha256.New, []byte("MYSECRET"))
	mac.Write(message)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal(messageMAC, []byte(expectedMAC))
}

func validMAC(h func() hash.Hash, message, messageMAC []byte) bool {
	// FIXME: Never hardcode the token into your app!
	mac := hmac.New(h, []byte("MYSECRET"))
	mac.Write(message)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal(messageMAC, []byte(expectedMAC))
}

func main() {
	http.HandleFunc("/github", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("called from github")
		printHeader(r)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		hmacsha256 := r.Header.Get("X-Hub-Signature-256")
		hmacsha256 = strings.TrimLeft(hmacsha256, "sha256=")
		fmt.Println(validMACSHA256(body, []byte(hmacsha256)))

	})
	http.HandleFunc("/twitter", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("called from twitter")
		printHeader(r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
