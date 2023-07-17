package logger

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

type contextKey int

const (
	CorrelationIDKey contextKey = iota
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		l := Get()
		correlationID := xid.New().String()
		ctx := context.WithValue(r.Context(), CorrelationIDKey, correlationID)
		r = r.WithContext(ctx)

		l.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("correlation_id", correlationID)
		})

		w.Header().Add("X-Correlation-ID", correlationID)

		lrw := newLoggingResponseWriter(w)
		r = r.WithContext(l.WithContext(r.Context()))

		defer func() {
			panicVal := recover()
			if panicVal != nil {
				lrw.statusCode = http.StatusInternalServerError
				panic(panicVal)
			}

			l.
				Info().
				Str("method", r.Method).
				Str("url", r.URL.RequestURI()).
				Str("user_agent", r.UserAgent()).
				Int("status_code", lrw.statusCode).
				Dur("elapsed_ms", time.Since(start)).
				Msg("incoming request")
		}()

		next.ServeHTTP(lrw, r)
	})
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}