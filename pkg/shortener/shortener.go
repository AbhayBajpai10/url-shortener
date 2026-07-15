package shortener

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

// take charset->pick rand chrs-> return string
func GenerateShortCode(length int) string {
	rand.Seed(time.Now().UnixMicro())
	code:= make([]byte, length);
	for i:= range code {
		code[i]=charset[rand.Intn(len(charset))]
	}
	return string(code);
}
