package handlers

import (
	"net/http"

	"github.com/almaraz333/go-astronomy-server/internal/utils"
)

func Test(w http.ResponseWriter, r *http.Request) {
	utils.MiddlewareChain(w, r, utils.Logging, func(w http.ResponseWriter, r *http.Request) error {
		return utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"test": "Working!"})
	})
}
