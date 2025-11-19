package core

import (
	"log"
	"time"
)

func expireSample() float32 {
	var limit int = 20
	var expiredCount int = 0

	for k, v := range store {
		if v.ExpiresAt != -1 {
			limit--

			// if the key is expired
			if v.ExpiresAt <= time.Now().UnixMilli() {
				delete(store, k)
				expiredCount++
			}
		}

		// once we iterated to 20 keys that have some expiration set we break the loop

		if limit == 0 {
			break
		}
	}
	return float32(expiredCount) / float32(20.0)
}

// Delete all the expired keys - the active way
func DeleteExpiredKeys() {
	for {
		frac := expireSample()
		// if the sample has less tahn 25% keys expired we break the loop

		if frac < 0.25 {
			break
		}
	}

	log.Println("deleted the expired but undeleted keys ", len(store))
}
