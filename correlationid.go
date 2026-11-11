package requestwrapper

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

const CorrelationIDHeader = "X-Correlation-ID"

type correlationIDKey struct{}

// CorrelationID attaches the incoming or a generated ID to the request context
// and response headers.
func CorrelationID(next HandlerFunc) HandlerFunc {
	return func(w ResponseWriter, r *http.Request) Response {
		id := r.Header.Get(CorrelationIDHeader)
		if id == "" {
			generated, err := uuid.NewRandom()
			if err != nil {
				log.Printf("generate correlation ID: %v", err)
				return next(w, r)
			}
			id = generated.String()
		}
		w.Header().Set(CorrelationIDHeader, id)
		ctx := context.WithValue(r.Context(), correlationIDKey{}, id)
		return next(w, r.WithContext(ctx))
	}
}

// CorrelationIDFromContext returns the correlation ID, if one is present.
func CorrelationIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(correlationIDKey{}).(string)
	return id, ok
}
