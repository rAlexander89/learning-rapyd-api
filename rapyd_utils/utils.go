package rapydUtils

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func GenSalt() (string, error) {
	salt := make([]rune, 12)

	for i := range salt {
		r, err := genRandomRune()
		if err != nil {
			return "", errors.New("error getting next rune")
		}

		salt[i] = *r
	}
	return string(salt), nil
}

func genRandomRune() (*rune, error) {
	charLength := len(letters)
	i, err := rand.Int(rand.Reader, new(big.Int).SetInt64(int64(charLength)))

	if err != nil {
		return nil, errors.New("error getting random char")
	}

	return &letters[i.Int64()], nil
}
