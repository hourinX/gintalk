package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"gin-online-chat-backend/commons/globals"
	"math/big"
	"time"
)

func GenerateNumericID(length int) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("length must be at least 1")
	}
	var timeStr string
	switch {
	case length >= 13:
		timeStr = fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	case length == 10:
		timeStr = fmt.Sprintf("%d", time.Now().Unix())
	default:
		timeStr = fmt.Sprintf("%d", time.Now().Unix())[:length]
	}
	if len(timeStr) < length {
		randomLength := length - len(timeStr)
		randomNum, _ := rand.Int(rand.Reader, big.NewInt(int64(randomLength)))
		timeStr += fmt.Sprintf("%0*d", randomLength, randomNum)
	}
	return timeStr[:length], nil
}

func GeneratorUserCode() (string, error) {
	codeLen, err := rand.Int(rand.Reader, big.NewInt(3))
	if err != nil {
		return "", err
	}
	codeLength := int(codeLen.Int64()) + 8
	firstDigit, err := rand.Int(rand.Reader, big.NewInt(9))
	if err != nil {
		return "", err
	}
	code := string('1' + byte(firstDigit.Int64()))
	for i := 1; i < codeLength; i++ {
		digit, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code += string('0' + byte(digit.Int64()))
	}
	return code, nil
}

func GenerateUniqueUserCodeWithRetry() (string, error) {
	maxAttempts := 3
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		code, err := GeneratorUserCode()
		if err != nil {
			return "", fmt.Errorf("生成失败: %v", err)
		}
		exists, err := globals.IsCodeExistsInUser(code)
		if err != nil {
			return "", fmt.Errorf("数据库查询失败: %v", err)
		}
		if !exists {
			return code, nil
		}
		fmt.Printf("第%d次尝试：代码 %s 已存在，重新生成...\n", attempt, code)
	}
	return "", errors.New("生成冲突次数过多，请稍后再试")
}
