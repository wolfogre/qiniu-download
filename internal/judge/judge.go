package judge

import (
	"net/http"
	"github.com/wolfogre/qiniuauth/internal/dao"
	"github.com/wolfogre/qiniuauth/internal/log"
	"fmt"
)

const (
	ALLOW   = http.StatusOK
	REJECT  = http.StatusForbidden
	UNKNOWN = http.StatusBadRequest
)

var (
	status = "ok"
)

func Status() (bool, string) {
	return status == "ok", status
}

func Judge(domain string, limits []string) int {
	if domain == "" {
		return UNKNOWN
	}
	lms := parseLimits(limits)
	if lms == nil {
		return UNKNOWN
	}
	for _, v := range lms {
		logger := log.Logger.With(
			"domain", domain,
			"limit_second", v.Second,
			"limit_count", v.Count,
		)
		count, err := dao.Incr(domain, v.Second)
		if err != nil {
			status = fmt.Sprintf("redis error: %v", err)
			logger.Error(status)
			return REJECT
		}
		if count > v.Count {
			status = fmt.Sprintf("%v out of limit (%v, %v): %v", domain, v.Second, v.Count, count)
			logger.Error(status)
			return REJECT
		}
	}
	return ALLOW
}

