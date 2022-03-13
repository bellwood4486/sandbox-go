package main

import (
	"crypto/hmac"
	"crypto/sha1"
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
	return validMAC(sha256.New, message, messageMAC)
}

func validMACSHA1(message, messageMAC []byte) bool {
	return validMAC(sha1.New, message, messageMAC)
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
		hmacsha256 := strings.TrimPrefix(r.Header.Get("X-Hub-Signature-256"), "sha256=")
		fmt.Printf("sha256: %v\n", validMACSHA256(body, []byte(hmacsha256)))

		hmacsha1 := strings.TrimPrefix(r.Header.Get("X-Hub-Signature"), "sha1=")
		fmt.Printf("sha1: %v\n", validMACSHA1(body, []byte(hmacsha1)))

	})
	http.HandleFunc("/twitter", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("called from twitter")
		printHeader(r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
