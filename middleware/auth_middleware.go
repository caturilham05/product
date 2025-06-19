package middleware

import (
	"caturilham05/product/helper"
	"caturilham05/product/model/web"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	accessToken := request.Header.Get("Authorization")

	if request.Method == http.MethodPost && request.URL.Path == "/api/products" {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	// List path yang tidak perlu autentikasi
	allowedPaths := []string{
		"/api/login",
		"/swagger/",
	}

	// Cek apakah request termasuk dalam daftar allowedPaths
	for _, path := range allowedPaths {
		if strings.HasPrefix(request.URL.Path, path) {
			middleware.Handler.ServeHTTP(writer, request)
			return
		}
	}

	if accessToken == "" {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		writer.WriteHeader(http.StatusUnauthorized)
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	_, err := helper.ValidateToken(accessToken)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: err.Error(),
		}
		writer.WriteHeader(http.StatusUnauthorized)
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	middleware.Handler.ServeHTTP(writer, request)
}
