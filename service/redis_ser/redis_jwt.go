package redis_ser

import (
	"context"
	"fast_gin/global"
	"fast_gin/utils/jwts"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func Logout(token string) {
	claims, err := jwts.CheckToken(token)
	if err != nil {
		return
	}
	key := fmt.Sprintf("logout_%s", token)
	sub := claims.ExpiresAt.Sub(time.Now())
	_, err = global.Redis.Set(context.Background(), key, "", sub).Result()
	if err != nil {
		logrus.Error(err)
		return
	}
}

func HasLogout(token string) bool {
	key := fmt.Sprintf("logout_%s", token)
	_, err := global.Redis.Get(context.Background(), key).Result()
	if err != nil {
		return false
	}
	return true

}
