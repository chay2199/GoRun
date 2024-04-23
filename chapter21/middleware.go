import (
	"context"
	"net/http"
)

type contextKey int

const TraceKey contextKey = 1

func AddTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if traceID := r.Header.Get("X-Cloud-Trace-Context"); traceID != "" {
			ctx = context.WithValue(ctx, TraceKey, traceID)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}