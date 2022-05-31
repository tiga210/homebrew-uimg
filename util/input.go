package util

import (
	"bufio"
	"github.com/howeyc/gopass"
	"os"
)

func GetPassword() (string, error) {
	pass, err := gopass.GetPasswd()
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func GetInput() (string, error) {
	r := bufio.NewReader(os.Stdin)
	d, err := r.ReadBytes('\n')
	d = d[:len(d)-1]
	return string(d), err
}
