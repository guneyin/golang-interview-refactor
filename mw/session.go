package mw

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const sessionIDKey = "ice_session_id"

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrInvalidSession  = errors.New("invalid session")
)

func UseSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		if sess.Get(sessionIDKey) == nil {
			sess.Set(sessionIDKey, uuid.New().String())
			_ = sess.Save()
		}
		sessionID := sess.Get(sessionIDKey)

		c.Set(sessionIDKey, sessionID)
		c.Next()
	}
}

func GetSessionID(c *gin.Context) (string, error) {
	val, ok := c.Get(sessionIDKey)
	if !ok {
		return "", ErrSessionNotFound
	}

	id, ok := val.(string)
	if !ok {
		return "", ErrInvalidSession
	}

	return id, nil
}
