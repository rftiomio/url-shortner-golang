package utils

import (
	"crypto/sha256"
	"github.com/itchyny/base58-go"
	"log"
	"math/big"
	"os"
	"strconv"
)

func generateSha256(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

func generateBase58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortUrl(link string, userId string) string {
	urlHashBytes := generateSha256(link + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := generateBase58Encoded([]byte(strconv.FormatUint(generatedNumber, 10)))
	return finalString[:8]
}
