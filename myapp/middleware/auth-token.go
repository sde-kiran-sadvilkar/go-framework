package middleware

import "net/http"

func (m *Middleware) AuthToken(w http.ResponseWriter, r *http.Request) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := m.Models.Tokens.AuthenticateToken(r)
		if err != nil {
			var payload struct {
				Message string `json:"message"`
				Error   bool   `json:"error"`
			}

			payload.Error = true
			payload.Message = err.Error()

			_ = m.Core.WriteJson(w, http.StatusUnauthorized, payload)

		}
	})

}
