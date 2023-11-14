package utils

import (
	"os"

	"github.com/mdp/qrterminal/v3"
)

// generate a random hash and send it with URL
// when any handler gets a request it should validate the password
func GenerateQRCode(url string) {
	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 2,
		WithSixel: false,
	}
	qrterminal.GenerateWithConfig(url, config)
}
