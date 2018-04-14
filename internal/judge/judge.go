package judge

import (
	"time"
	"fmt"

	"github.com/google/uuid"
	"github.com/wolfogre/qiniu-download/internal/dao"
	"github.com/wolfogre/qiniu-download/internal/log"
)

// 一个 IP 1 小时最多下载 10 次，暂时写死
func GenToken(ip string) string {
	count, err := dao.Incr(ip, time.Hour)
	if err != nil {
		log.Logger.Error(err)
		log.ChangeStatus(err)
		return ""
	}

	if count > 10 {
		err := fmt.Errorf("%v download too frequently", ip)
		log.Logger.Error(err)
		log.ChangeStatus(err)
		return ""
	}

	token := uuid.New().String()
	if token == "" {
		err := fmt.Errorf("token is empty")
		log.Logger.Error(err)
		log.ChangeStatus(err)
		return ""
	}

	err = dao.PutToken(ip, token)
	if err != nil {
		log.Logger.Error(err)
		log.ChangeStatus(err)
		return ""
	}

	return token
}


func VerifyToken(token string) bool {
	result, err := dao.GetDeleteToken(token)
	if err != nil {
		log.Logger.Error(err)
		log.ChangeStatus(err)
		return false
	}
	return result
}