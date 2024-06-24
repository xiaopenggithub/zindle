package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 生成随机盐
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

// 加盐MD5哈希
func saltedMD5V(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}
	// 将盐和密码结合
	saltedPassword := string(salt) + password

	h := md5.New()
	h.Write([]byte(saltedPassword))
	hashedPassword := h.Sum(nil)

	// 返回盐值和哈希值，通常盐值需要和哈希值一起存储
	return hex.EncodeToString(hashedPassword), nil
}