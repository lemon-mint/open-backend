package password

import "crypto/rand"

// BEGIN DefaltSaltSize
const DefaultSaltSize = 32

// END DefaltSaltSize

func GetSalt() ([]byte, string) {
	buf := make([]byte, DefaultSaltSize)
	rand.Read(buf)
	return buf, string(buf)
}
