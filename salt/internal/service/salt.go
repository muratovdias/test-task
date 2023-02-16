package service

import (
	"math/rand"
	"time"
)

func GetSalt() string {
	str := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	salt := make([]byte, 12)
	rand.Seed(time.Now().Unix())
	for i := range salt {
		salt[i] = str[rand.Intn(len(str))]
	}
	return string(salt)
}
