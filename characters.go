package main

import "math/rand"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+,./;'[]<>?:{}"
const total = 89

func GenerateNextChar() string {
	return string(charset[rand.Intn(total)])
}
