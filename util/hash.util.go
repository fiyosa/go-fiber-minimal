package util

import (
	"encoding/base64"
	"fmt"
	"go-fiber-ddd/config"

	"github.com/speps/go-hashids/v2"
	"golang.org/x/crypto/bcrypt"
)

var Hash hashManager

type hashManager struct {
}

func setupHD() *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = config.Env.APP_SECRET
	hd.MinLength = 10
	id, _ := hashids.NewWithData(hd)
	return id
}

func (hashManager) EncodeId(data int) (string, error) {
	hd := setupHD()
	encode, err := hd.Encode([]int{data})
	if err != nil {
		return "", err
	}
	return encode, nil
}

func (hashManager) DecodeId(data string) (int, error) {
	hd := setupHD()
	decode, err := hd.DecodeWithError(data)
	if err != nil {
		return -1, err
	}
	return decode[0], err
}

func (hashManager) BcryptCreate(data string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(data), 10)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (hashManager) BcryptVerify(check string, original string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(original), []byte(check)); err != nil {
		return false
	}
	return true
}

func (hashManager) EncodeStr(data string) (string, error) {
	result := base64.StdEncoding.EncodeToString([]byte(data))
	return result, nil
}

func (hashManager) DecodeStr(encode string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return "", err
	}
	return string(result), nil
}
