package store

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestInitializeStore(t *testing.T) {
	assert.Equal(t, testStoreService.redisClient != nil, true)
}

func TestStoreAndRetrieval(t *testing.T) {
	initialLink := "https://google.com"
	shortUrl := "wAxRY1zL"

	SaveUrlMapping(shortUrl, initialLink)

	retrievedUrl := RetrieveInitialUrl(shortUrl)
	log.Println("-------Received Url------->>", retrievedUrl)
	assert.Equal(t, initialLink, retrievedUrl)
}
