package util

import (
	"encoding/base64"
	"fmt"
	"go-fiber-minimal/config"

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

// 1 => "hj"
func (hashManager) EncodeId(data uint) (string, error) {
	hd := setupHD()
	encode, err := hd.Encode([]int{int(data)})
	if err != nil {
		return "", err
	}
	return encode, nil
}

// "hj" => 1
func (hashManager) DecodeId(data string) (uint, error) {
	hd := setupHD()
	decode, err := hd.DecodeWithError(data)
	if err != nil {
		return 0, err
	}
	return uint(decode[0]), err
}

// "password" => "$2a$10$X0X0X0X0X0X0"
func (hashManager) BcryptCreate(data string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(data), 10)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// ("password", "$2a$10$X0X0X0X0X0X0") => true
func (hashManager) BcryptVerify(check string, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(check)); err != nil {
		return false
	}
	return true
}

// "password" => "cGFzc3dvcmQ="
func (hashManager) EncodeStr(data string) (string, error) {
	result := base64.StdEncoding.EncodeToString([]byte(data))
	return result, nil
}

// "cGFzc3dvcmQ=" => "password"
func (hashManager) DecodeStr(encode string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return "", err
	}
	return string(result), nil
}
