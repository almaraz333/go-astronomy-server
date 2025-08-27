package utils

import (
	"log/slog"
	"net/http"
)

func Logging(w http.ResponseWriter, r *http.Request) error {
	slog.Info("", r.Method, r.URL.Path+r.URL.RawQuery)
	return nil
}
