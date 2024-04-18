package middleware

import "github.com/gin-gonic/gin"

// Middleware wraps gin HandlerFunc to intercept request.
type Middleware = gin.HandlerFunc
