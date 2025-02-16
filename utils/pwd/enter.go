package pwd

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// GenerateFromPassword 密码加密
func GenerateFromPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("加密密码错误 %s", err)
		return ""
	}
	return string(hashedPassword)
}

// CompareHashAndPassword 密码验证
func CompareHashAndPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
