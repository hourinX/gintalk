package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func GenerateNumericID() (string, error) {
	timestampNano := time.Now().UnixNano()
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to read random bytes: %w", err)
	}

	tsBig := new(big.Int).SetInt64(timestampNano)
	randBig := new(big.Int).SetBytes(randomBytes)

	combinedBig := new(big.Int).Lsh(tsBig, 64)
	combinedBig.Or(combinedBig, randBig)

	// 4. 将大整数转换为字符串
	return combinedBig.String(), nil
}
