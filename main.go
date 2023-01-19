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
			fmt.Println(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		if d.Data == "" {
			fmt.Println("no secret yet… will check again later.")
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf(
			"secret: updated: %s, created: %s, value: %s\n",
			d.Updated, d.Created, d.Data,
		)
		time.Sleep(5 * time.Second)
	}
}
