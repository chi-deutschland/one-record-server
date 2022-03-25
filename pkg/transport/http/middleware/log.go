package middleware

import (
	onerecordhttp "github.com/chi-deutschland/one-record-server/pkg/transport/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func LogMiddleware(srvRole string) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lrw := onerecordhttp.NewLoggingResponseWriter(w)
			h.ServeHTTP(lrw, r)
			statusCode := lrw.StatusCode
			logrus.WithFields(logrus.Fields{
				"role": srvRole,
			}).Info(
				r.Host,
				" ",
				r.Method,
				" ",
				r.Proto,
				" | resp: ",
				statusCode,
				" - ",
				http.StatusText(statusCode),
				" | ",
				r.UserAgent())
		})
	}
}
