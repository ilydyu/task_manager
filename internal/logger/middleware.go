package logger

import (
	"fmt"
	"net/http"

	"github.com/ilydyu/task_manager.git/pkg/router"
	"github.com/rs/zerolog/log"
)

func Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := router.WriterWrapper(w)
		next.ServeHTTP(ww, r.WithContext(r.Context()))

		log.Info().
			Str("proto", "http").
			Int("code", ww.Code()).
			Str("method", fmt.Sprintf("%s %s", r.Method, router.ExtractPath(r.Context()))).
			Send()
	}

	return http.HandlerFunc(fn)
}
