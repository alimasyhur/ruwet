package utils

import (
	"fmt"
	"time"
)

func StartSpinner(done chan bool, message *string) {
	go func() {
		chars := []string{"|", "/", "-", "\\"}
		i := 0

		for {
			select {
			case <-done:
				fmt.Print("\r\033[K")
				return
			default:
				fmt.Printf("\r%s %s", chars[i%len(chars)], *message)
				time.Sleep(100 * time.Millisecond)
				i++
			}
		}
	}()
}
