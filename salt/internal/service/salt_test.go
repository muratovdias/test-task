package service

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestGetSalt(t *testing.T) {
	rand.Seed(time.Now().Unix())
	salt := GetSalt()
	if len(salt) != 12 {
		t.Errorf("expected length 12, instead %d", len(salt))
	}
	str := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, ch := range salt {
		if !strings.ContainsRune(str, ch) {
			t.Errorf("salt contains unacceptable character like '%c'", ch)
		}
	}
}
