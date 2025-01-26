package mw

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const sessionIdKey = "ice_session_id"

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrInvalidSession  = errors.New("invalid session")
)

func UseSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		if sess.Get(sessionIdKey) == nil {
			sess.Set(sessionIdKey, uuid.New().String())
			_ = sess.Save()
		}
		sessionId := sess.Get(sessionIdKey)

		c.Set(sessionIdKey, sessionId)
		c.Next()
	}
}

func GetSessionID(c *gin.Context) (string, error) {
	val, ok := c.Get(sessionIdKey)
	if !ok {
		return "", ErrSessionNotFound
	}

	id, ok := val.(string)
	if !ok {
		return "", ErrInvalidSession
	}

	return id, nil
}
