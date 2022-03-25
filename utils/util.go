package utils

import (
	"go-baseline/exception"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func StrToInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return result
}

func StrToInt64(value string) int64 {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	return result
}

func IntToStr(value int) string {
	return strconv.Itoa(value)
}

func HashPassword(password string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	exception.PanicIfNeeded(err)

	return string(bytes)

}
