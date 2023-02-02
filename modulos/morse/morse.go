package morse

import (
	"strings"
	// tenemos pergar la direcion donde esta el paquete o copiar donde esta en go
	"github.com/alwindoss/morse"
)

func Palbras(mensaje string) string {
	h := morse.NewHacker()
	morseCode, _ := h.Encode(strings.NewReader(mensaje))

	// Morse Code is: -.-. --- ...- . .-. - / - .... .. ... / - --- / -- --- .-. ... .
	return string(morseCode)
}
