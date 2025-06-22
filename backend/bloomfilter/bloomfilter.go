package bloomfilter

import (
	"log"

	"github.com/bits-and-blooms/bloom/v3"
)

var bloomFilter *bloom.BloomFilter

func InitBloomFilter(expectedItems uint, falsePositiveRate float64) {
	log.Printf("Initializing bloom filter for %d items with %.2f%% false positive rate", expectedItems, falsePositiveRate*100)
	bloomFilter = bloom.NewWithEstimates(expectedItems, falsePositiveRate)
	log.Println("Bloom filter initialized successfully")
}

func IsUsernameInBloom(username string) bool {
	if bloomFilter == nil {
		log.Println("Warning: Bloom filter not initialized, returning false")
		return false
	}
	log.Println("Checking username in bloom filter:", username)
	return bloomFilter.Test([]byte(username))
}

func AddUsernameToBloom(username string) {
	if bloomFilter == nil {
		log.Println("Warning: Bloom filter not initialized, cannot add username")
		return
	}
	log.Println("Adding username to bloom filter:", username)
	bloomFilter.Add([]byte(username))
}
