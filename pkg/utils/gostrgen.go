package utils

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

// GenRandNone .
const GenRandNone = 0

// GenRandLower .
const GenRandLower = 1 << 0

// GenRandUpper .
const GenRandUpper = 1 << 1

// GenRandDigit .
const GenRandDigit = 1 << 2

// GenRandSymbol .
const GenRandSymbol = 1 << 3

// GenRandLowerUpper .
const GenRandLowerUpper = GenRandLower | GenRandUpper

// GenRandLowerDigit .
const GenRandLowerDigit = GenRandLower | GenRandDigit

// GenRandUpperDigit .
const GenRandUpperDigit = GenRandUpper | GenRandDigit

// GenRandLowerUpperDigit .
const GenRandLowerUpperDigit = GenRandLowerUpper | GenRandDigit

// GenRandAll .
const GenRandAll = GenRandLowerUpperDigit | GenRandSymbol

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digit = "0123456789"
const symbols = "~!@#$%^&*()_+-="

// RandStringInit .
func RandStringInit() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandString .
func RandString(size int, set int, include string, exclude string) (string, error) {
	all := include
	if set&GenRandLower > 0 {
		all += lower
	}
	if set&GenRandUpper > 0 {
		all += upper
	}
	if set&GenRandDigit > 0 {
		all += digit
	}
	if set&GenRandSymbol > 0 {
		all += symbols
	}

	lenAll := len(all)
	if len(exclude) >= lenAll {
		return "", errors.New("Too much to exclude")
	}
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		b := all[rand.Intn(lenAll)]
		if strings.Contains(exclude, string(b)) {
			i--
			continue
		}
		buf[i] = b
	}
	return string(buf), nil
}
