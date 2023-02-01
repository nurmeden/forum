package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	// (BEGIN) The code from this point is the same as the first part of the `Welcome` route
	fmt.Println("refreshing")
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			ErrorHandler(w, http.StatusUnauthorized)
			return
		}
		ErrorHandler(w, http.StatusUnauthorized)
		return
	}
	sessionToken := c.Value

	userSession, exists := sessions[sessionToken]
	if !exists {
		ErrorHandler(w, http.StatusUnauthorized)
		return
	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		ErrorHandler(w, http.StatusUnauthorized)
		return
	}
	// (END) The code until this point is the same as the first part of the `Welcome` route

	// If the previous session is valid, create a new session token for the current user
	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the user whom it represents
	sessions[newSessionToken] = session{
		Username: userSession.Username,
		expiry:   expiresAt,
	}

	// Delete the older session token
	delete(sessions, sessionToken)

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
	http.Redirect(w, r, "/", 302)
}
