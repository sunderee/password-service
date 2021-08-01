package service

import (
	"math/rand"
	"time"
)

const (
	lowercase string = "abcdefghijklmnopqrstuvwxyz"
	uppercase string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   string = "123456789"
	specials  string = "`~!@#$%^&*())=_-+=.,"
)

type Options struct {
	HasUppercase   bool
	HasNumbers     bool
	HasSpecials    bool
	PasswordLength int
}

func GeneratePassword(options Options) string {
	if options.PasswordLength <= 0 {
		panic("You cannot generate a password with negative or zero length!")
	}

	rand.Seed(time.Now().UnixNano())

	var password string = ""
	var charset string = formCharset(options.HasUppercase, options.HasNumbers, options.HasSpecials)
	for i := 0; i < options.PasswordLength; i++ {
		randomCharacter := charset[rand.Intn(len(charset))]
		password = password + string(randomCharacter)
	}
	return password
}

func formCharset(hasUppercase, hasNumbers, hasSpecials bool) string {
	var charset string = lowercase
	if hasUppercase {
		charset = charset + uppercase
	}
	if hasNumbers {
		charset = charset + numbers
	}
	if hasSpecials {
		charset = charset + specials
	}
	return charset
}
