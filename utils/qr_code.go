package utils

import (
	"github.com/AYehia0/lezer/internal"
)

// generate a random hash and send it with URL
// when any handler gets a request it should validate the password
func GenerateQRCode(url string) {
	internal.New().Get(url).Print()
}
