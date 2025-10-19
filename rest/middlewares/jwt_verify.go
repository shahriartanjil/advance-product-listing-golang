package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"ecommere.com/config"
	"ecommere.com/utility"
)

type ctxKey string

const userKey ctxKey = "user"

// GetUserFromContext returns the JWT payload set by Auth().
func GetUserFromContext(ctx context.Context) (utility.Payload, bool) {
	u, ok := ctx.Value(userKey).(utility.Payload)
	return u, ok
}

// Auth verifies a Bearer JWT created by utility.CreateJwt and injects the user payload into context.
// Example: mux.Handle("GET /products", Auth(secret)(http.HandlerFunc(handlers.GetProducts)))
func Auth(secret string) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			auth := r.Header.Get("Authorization")
			if !strings.HasPrefix(auth, "Bearer ") {
				utility.SendError(w, http.StatusUnauthorized, "missing or invalid Authorization header")
				return
			}
			token := strings.TrimPrefix(auth, "Bearer ")
			parts := strings.Split(token, ".")
			if len(parts) != 3 {
				utility.SendError(w, http.StatusUnauthorized, "invalid token format")
				return
			}

			headerB64, payloadB64, signatureB64 := parts[0], parts[1], parts[2]

			// Recompute signature: base64url(header) + "." + base64url(payload)
			mac := hmac.New(sha256.New, []byte(secret))
			mac.Write([]byte(headerB64))
			mac.Write([]byte("."))
			mac.Write([]byte(payloadB64))
			expectedSig := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(mac.Sum(nil))

			if !hmac.Equal([]byte(signatureB64), []byte(expectedSig)) {
				utility.SendError(w, http.StatusUnauthorized, "invalid token signature")
				return
			}

			// Decode payload (base64url, no padding)
			payloadBytes, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(payloadB64)
			if err != nil {
				utility.SendError(w, http.StatusUnauthorized, "invalid payload encoding")
				return
			}

			var user utility.Payload
			if err := json.Unmarshal(payloadBytes, &user); err != nil {
				utility.SendError(w, http.StatusUnauthorized, "invalid payload data")
				return
			}

			// Store user in context and continue
			ctx := context.WithValue(r.Context(), userKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// AuthFromConfig creates the Auth middleware using the secret in config.
func AuthFromConfig() func(http.Handler) http.Handler {
	secret := config.GetConfig().JwtSecretKey
	return Auth(secret) // your existing Auth(secret) middleware
}
