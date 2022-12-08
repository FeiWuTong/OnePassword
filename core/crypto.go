package core

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/wumansgy/goEncrypt/aes"
)

type Crypto struct {
	sk []byte
}

func (c *Crypto) KeyGen(key string) {
	h := sha256.New()
	h.Write([]byte(key))
	c.sk = h.Sum(nil)
}

func (c *Crypto) KeyHash() string {
	return hex.EncodeToString(c.sk)
}

func (c *Crypto) Encrypt(text string) (string, error) {
	return aes.AesCbcEncryptBase64([]byte(text), c.sk, nil)
}

func (c *Crypto) Decrypt(text string) (string, error) {
	plainText, err := aes.AesCbcDecryptByBase64(text, c.sk, nil)
	return string(plainText), err
}

// TODO: add a random generator (digit / +letter(including capital) / +symbol )
