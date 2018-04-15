package handler

import (
	"net/http"
	"strings"
	"fmt"

	"go.uber.org/zap"
	"github.com/wolfogre/qiniu-download/internal/log"
	"github.com/wolfogre/qiniu-download/internal/judge"
)

const (
	AUTH_PREFIX = "/auth"
	ALLOW   = http.StatusOK
	REJECT  = http.StatusForbidden
)

type Handler struct {
	cdnAddr string
}

func NewHandler(cdnAddr string) *Handler {
	return &Handler{
		cdnAddr: cdnAddr,
	}
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
		if r.URL.Path == AUTH_PREFIX {
			if judge.VerifyToken(r.URL.Query().Get("token")) {
				response(logger, w, ALLOW, "")
			} else {
				response(logger, w, REJECT, "")
			}
			return
		}
		abort(logger, w, http.StatusMethodNotAllowed)
	case "GET":
		switch r.URL.Path {
		case "/_status":
			if ok, msg := log.Status(); ok {
				response(logger, w, http.StatusOK, msg)
			} else {
				response(logger, w, http.StatusInternalServerError, msg)
			}
		default:
			token := judge.GenToken(ip)
			if token == "" {
				response(logger, w, REJECT, "you download too frequently, please wait for 1 hour then retry.")
				log.ChangeStatus(fmt.Errorf("%v download too frequently", ip))
				return
			}
			redirectUrl := fmt.Sprintf("%v%v?token=%v", h.cdnAddr, r.URL.Path, token)
			http.Redirect(w, r, redirectUrl, http.StatusFound)
			logger.With(
				"status_code", http.StatusFound,
				"token", token,
				"url", r.URL.String(),
				"redirect_url", redirectUrl,
			).Info("redirect")
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