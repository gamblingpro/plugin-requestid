package plugin_requestid

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

const defaultHeader = "X-Request-ID"

type Config struct {
	HeaderName string
	Enabled    bool
}

func CreateConfig() *Config {
	return &Config{
		HeaderName: defaultHeader,
		Enabled:    true,
	}
}

func New(ctx context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.Enabled && r.Header.Get(config.HeaderName) == "" {
			id := uuid.New()
			id.MarshalBinary()
			r.Header.Add(config.HeaderName, id.String())
			w.Header().Add(config.HeaderName, id.String())
		}

		next.ServeHTTP(w, r)
	}), nil
}
