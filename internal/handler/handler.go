package handler

import (
	"net/http"
	"github.com/wolfogre/qiniuauth/internal/log"
	"strings"
	"go.uber.org/zap"
	"github.com/wolfogre/qiniuauth/internal/judge"
	"fmt"
)

const (
	AUTH_PREFIX = "/auth"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.RemoteAddr[0 : strings.LastIndex(r.RemoteAddr, ":")]
	}
	logger := log.Logger.With(
		"ip", ip,
		"method", r.Method,
		"host", r.Host,
		"url", r.URL.String(),
		"user_agent", r.UserAgent(),
	)

	switch r.Method {
	case "HEAD":
		// /auth-domain_2h10_1d100
		if strings.HasPrefix(r.URL.Path, AUTH_PREFIX) {
			str := strings.TrimPrefix(r.URL.Path,"/")
			idx := strings.Index(str, "/")
			if idx != -1 {
				str = str[:idx]
			}
			strs := strings.Split(str, "_")
			domain := ""
			if splits := strings.Split(strs[0], "-"); len(splits) == 2 {
				domain = splits[1]
			}
			response(logger, w, judge.Judge(domain, strs[1:]), "")
			return
		}
		abort(logger, w, http.StatusNotFound)
	case "GET":
		switch r.URL.Path {
		case "/_status":
			if ok, msg := judge.Status(); ok {
				response(logger, w, http.StatusOK, msg)
			} else {
				response(logger, w, http.StatusInternalServerError, msg)
			}
		default:
			abort(logger, w, http.StatusNotFound)
		}
	default:
		abort(logger, w, http.StatusMethodNotAllowed)
	}
}

func abort(logger *zap.SugaredLogger, w http.ResponseWriter, code int) {
	response(logger, w, code, http.StatusText(code))
}

func response(logger *zap.SugaredLogger, w http.ResponseWriter, code int, content string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	if content == "" {
		fmt.Fprint(w, content)
	} else {
		fmt.Fprintln(w, content)
	}
	logger.With(
		"status_code", code,
	).Info(content)
}