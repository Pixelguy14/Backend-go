package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			//Quiere decir que no llego un token
			http.Error(w, "Token not given", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.Replace(authHeader, "Bearer", "", 1)
		//Validamos que el token leido sea valido
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte("PALABRA_SECRETA"), nil
		})
		//ahora validamos el token de verdad
		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
