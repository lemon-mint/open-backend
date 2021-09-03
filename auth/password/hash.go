package password

import (
	"bytes"
	"crypto/sha512"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/pbkdf2"
)

// BEGIN Default Algorithm
const DefaultAlgorithm = Algorithm_Argon2id

// END Default Algorithm

type AlgID uint32

const (
	Algorithm_Argon2id AlgID = iota
	Algorithm_Argon2i
	Algorithm_PBKDF2_SHA256
)

type Algorithm struct {
	Hasher   func(data, salt []byte) []byte
	Verifier func(data, salt []byte, hash []byte) bool
}

var Algorithms []Algorithm

var _ = func() int {
	Algorithms = make([]Algorithm, 3)
	Algorithms[Algorithm_Argon2id] = Algorithm{
		Hasher: func(data, salt []byte) []byte {
			return argon2.IDKey(data, salt, 24, 64*1024, 4, 32)
		},
		Verifier: func(data, salt []byte, hash []byte) bool {
			h1 := argon2.IDKey(data, salt, 24, 64*1024, 4, 32)
			return bytes.Equal(h1, hash)
		},
	}
	Algorithms[Algorithm_Argon2i] = Algorithm{
		Hasher: func(data, salt []byte) []byte {
			return argon2.Key(data, salt, 24, 64*1024, 4, 32)
		},
		Verifier: func(data, salt []byte, hash []byte) bool {
			h1 := argon2.Key(data, salt, 24, 64*1024, 4, 32)
			return bytes.Equal(h1, hash)
		},
	}
	Algorithms[Algorithm_PBKDF2_SHA256] = Algorithm{
		Hasher: func(data, salt []byte) []byte {
			return pbkdf2.Key(data, salt, 1000000, 32, sha512.New)
		},
		Verifier: func(data, salt []byte, hash []byte) bool {
			h1 := pbkdf2.Key(data, salt, 1000000, 32, sha512.New)
			return bytes.Equal(h1, hash)
		},
	}
	return 0
}()
