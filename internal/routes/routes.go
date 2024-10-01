package routes

import (
	"net/http"
	"vws_api/internal/middleware"
	"vws_api/types"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(vws *types.VWS) {
	r := chi.NewRouter()

	r.Use(middleware.AuthMiddleware(vws.BotToken))

	http.ListenAndServe(":8080", r)
}
