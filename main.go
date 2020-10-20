package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"net"
	"os"
	"regexp"
)

func main() {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(fmt.Sprintf("failed to read random bytes: %s\n", err))
	}

	r := bufio.NewScanner(os.Stdin)

	iprgx := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+|[0-9a-f:]+)`)

	for r.Scan() {
		line := r.Text()
		out := iprgx.ReplaceAllStringFunc(line, func(in string) string {
			ip := net.ParseIP(in)
			if ip == nil {
				return in
			}

			sum := sha256.Sum256([]byte(fmt.Sprintf("%s:%s", randomBytes, ip.String())))
			ip[len(ip)-1] = sum[0]

			return ip.String()
		})

		fmt.Println(out)
	}

}
