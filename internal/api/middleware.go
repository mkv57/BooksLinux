package api

import (
	"books/internal/logger"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func Logging(log *slog.Logger) mux.MiddlewareFunc {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			var metod string

			switch r.Method {
			case "GET":
				metod = "отвечаем на запрос"
			case "POST":
				metod = "создаём"
			case "PUT":
				metod = "обновляем"
			case "DELETE":
				metod = "удаляем"
			}

			log.Info("Request",
				slog.String("msg", metod), slog.String("ip", r.RemoteAddr),
			)

			log = log.With(
				slog.String("ip", r.RemoteAddr),
				slog.String("url_path", r.URL.Path),
			)

			ctx := r.Context()

			ctx = logger.NewContext(ctx, log)

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

		})
	}
}
