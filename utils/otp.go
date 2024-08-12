package utils

import (
	"math/rand"
)

func GenerateOTP(length int) string {
	numbers := "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		otp[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(otp)
}
