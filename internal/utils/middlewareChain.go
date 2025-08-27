package utils

import "net/http"

func MiddlewareChain(w http.ResponseWriter, r *http.Request, middleware ...func(w http.ResponseWriter, r *http.Request) error) error {
	for _, mw := range middleware {
		if err := mw(w, r); err != nil {
			return err
		}
	}

	return nil
}
