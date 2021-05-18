package myshawdowsockets
import (
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
	"time"
)

const passwordLength = 256

type password [passwordLength]byte

func init() {
	rand.Seed(time.Now().Unix())
}

// Exchange the bytes password to string
func (password *password) Tostring() string {
	return base64.StdEncoding.EncodeToString(password[:])
}

// Exchange string to bytes
func ParsePassword(passwordString string) (*password, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(passwordString))
	if err != nil || len(bs) != passwordLength {
		return nil, errors.New("invalid password")
	}
	password := password{}
	copy(password[:], bs)
	bs = nil
	return &password, nil
}

// Generate password
func RandPassword() string {
	intArr := rand.Perm(passwordLength)
	password := &password{}
	for i, v := range intArr {
		password[i] = byte(v)
		if i == v {
			// cannot repeat the same alp
			return RandPassword()
		}
	}
	return password.Tostring()
}