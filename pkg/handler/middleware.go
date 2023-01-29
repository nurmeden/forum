package handler

import (
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		next.ServeHTTP(w, r)
	})
}

//func welcome(next http.HandlerFunc) http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		c, err := r.Cookie("session_token")
//		fmt.Println(c)
//		if err != nil {
//			if err == http.ErrNoCookie {
//				fmt.Println("Status Unauthorized")
//				w.WriteHeader(http.StatusUnauthorized)
//				return
//			}
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		sessionToken := c.Value
//
//		userSession, exists := sessions[sessionToken]
//		if !exists {
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		fmt.Println(userSession)
//		if userSession.isExpired() {
//			delete(sessions, sessionToken)
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		//w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))
//		next.ServeHTTP(w, r)
//	})
//}
