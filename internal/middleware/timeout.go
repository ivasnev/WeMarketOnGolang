package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Заменяем контекст запроса
		c.Request = c.Request.WithContext(ctx)

		// Канал завершения
		done := make(chan struct{}, 1)

		go func() {
			c.Next()
			done <- struct{}{}
		}()

		select {
		case <-ctx.Done():
			// Тайм-аут запроса
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{"error": "request timed out"})
		case <-done:
			// Успешное завершение
		}
	}
}
