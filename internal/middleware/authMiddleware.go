package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"
	"vws_api/types"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

//AuthMiddleware - валидация инитдаты пользователя
func AuthMiddleware(botToken string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authParts := strings.Split(r.Header.Get("Authorization"), " ")
			if len(authParts) < 2 {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}

			authType := authParts[0]
			authData := authParts[1]

			switch authType {
			case "tma":

				if err := initdata.Validate(authData, botToken, time.Hour*1); err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
					return
				}

				initData, err := initdata.Parse(authData)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
					return
				}

				ctx := context.WithValue(r.Context(), types.InitDataKey, initData)
				next.ServeHTTP(w, r.WithContext(ctx))

			default:
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}

		})
	}
}
