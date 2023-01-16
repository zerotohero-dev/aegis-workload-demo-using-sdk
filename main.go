/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	"fmt"
	"github.com/zerotohero-dev/aegis-sdk-go/sentry"
	"log"
	"time"
)

func main() {
	for {
		log.Println("fetch")
		d, err := sentry.Fetch()

		if err != nil {
			fmt.Println("Failed to read the secrets file. Will retry in 5 seconds…")
		} else {
			fmt.Println("secret: '", d, "'")
		}

		time.Sleep(5 * time.Second)
	}
}
